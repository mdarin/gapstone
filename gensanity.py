"""
Как использовать скрипт:

Запустите его с указанием входной и выходной директорий:
python gensanity.py /path/to/constants_dir /output/dir --command "./gensanity.py /path/to/constants_dir"
Особенности работы скрипта:

Поиск констант:
Ищет файлы *_constants.go в указанной директории
Извлекает константы по четырем шаблонам:
CS_ARCH_<ARCH>
<ARCH>_REG_ENDING
<ARCH>_INS_ENDING
<ARCH>_GRP_ENDING

Сопоставление констант:
Для каждой архитектуры (например, ARM64) находит соответствующие константы REG_ENDING, INS_ENDING и GRP_ENDING
Создает структуру данных, связывающую эти константы

Генерация Go файла:
Создает файл sanity_checks.go с правильным форматом
Включает в заголовок команду, которая использовалась для генерации
Добавляет текущую дату и время
Сортируе архитектуры для стабильного вывода
"""

import os
import re
import argparse
from datetime import datetime
from pathlib import Path
from typing import Dict, List, Set, Optional

def find_architectures_in_main_file(main_file: Path) -> Set[str]:
    """Извлекает архитектуры из capstone_constants.go"""
    if not main_file.exists():
        print(f"\033[33mWarning: Main constants file {main_file} not found!\033[0m")
        return set()

    with open(main_file, 'r', encoding='utf-8') as f:
        content = f.read()

    arch_pattern = re.compile(r'CS_ARCH_([A-Z0-9]+)')
    architectures = set()
    
    for match in arch_pattern.finditer(content):
        architectures.add(match.group(1))
    
    print(f"Found architectures in {main_file}: {', '.join(architectures)}")
    return architectures

def find_related_constants(directory: Path, architectures: Set[str]) -> Dict[str, Dict[str, str]]:
    """Ищет связанные константы во всех *_constants.go файлах"""
    constants = {
        'REG_ENDING': {},
        'INS_ENDING': {},
        'GRP_ENDING': {},
    }

    pattern_map = {
        'REG_ENDING': re.compile(r'([A-Z0-9]+)_REG_ENDING'),
        'INS_ENDING': re.compile(r'([A-Z0-9]+)_INS_ENDING'),
        'GRP_ENDING': re.compile(r'([A-Z0-9]+)_GRP_ENDING'),
    }

    for const_file in directory.glob('*_constants.go'):
        if const_file.name == 'capstone_constants.go':
            continue  # Уже обработали этот файл

        with open(const_file, 'r', encoding='utf-8') as f:
            content = f.read()

        for const_type, pattern in pattern_map.items():
            for match in pattern.finditer(content):
                arch = match.group(1)
                if arch in architectures:
                    full_const = f'{arch}_{const_type}'
                    constants[const_type][arch] = full_const
                    print(f"Found {full_const} in {const_file}")

    return constants

def generate_go_file(architectures: Set[str], 
                   constants: Dict[str, Dict[str, str]],
                   output_dir: Path, 
                   command: str):
    """Генерирует файл sanity_checks.go"""
    # Сортируем архитектуры для стабильного вывода
    sorted_archs = sorted(architectures)

    # Генерируем содержимое файла
    content = f"""/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples,try reading the *_test.go files.

    Library Author: Nguyen Anh Quynh
    Binding Author: Ben Nagy
    License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
    Command: {command}
    Created at: {datetime.now().strftime("%Y-%m-%d %H:%M:%S")}

*/

package gapstone

// WARN
// Maintain the expected version and sanity checks manually, so we can verify
// against the installed C lib. Not foolproof, but should save 90% of accidents
const expectedMaj = 5
const expectedMin = 0

type sanityCheck struct {{
    insMax int
    regMax int
    grpMax int
}}

type sanityChecks map[int]sanityCheck

func (s *sanityChecks) Maj() int {{ return expectedMaj }}
func (s *sanityChecks) Min() int {{ return expectedMin }}

// WARN
// Remember the all the constants CONST are direct refs to C.CONST, so in
// combination with these we should be _fairly_ sure we're getting the
// disassembly capstone expects to provide.
// * <ARCH>_GRP_ENDING || <ARCH>_INS_ENDING || <ARCH>_REG_ENDING
var checks = sanityChecks{{
"""

    for arch in sorted_archs:
        reg = constants['REG_ENDING'].get(arch)
        ins = constants['INS_ENDING'].get(arch)
        grp = constants['GRP_ENDING'].get(arch)

        if not all([reg, ins, grp]):
            print(f"\033[33mWarning: Incomplete constants for architecture {arch}\033[0m")
            continue

        content += f"\tCS_ARCH_{arch}: sanityCheck{{\n"
        content += f"\t\tregMax: {reg},\n"
        content += f"\t\tinsMax: {ins},\n"
        content += f"\t\tgrpMax: {grp},\n"
        content += "\t},\n"

    content += "}\n"

    # Записываем файл
    output_path = output_dir / "sanity_checks.go"
    with open(output_path, 'w', encoding='utf-8') as f:
        f.write(content)

def main():
    parser = argparse.ArgumentParser(description='Generate sanity_checks.go from *_constants.go files')
    parser.add_argument('input_dir', help='Directory with *_constants.go files to scan', default='.')
    parser.add_argument('output_dir', help='Output directory for sanity_checks.go', default='.')
    parser.add_argument('--command', help='Command to include in header', default='./gensanity.py')
    args = parser.parse_args()

    input_path = Path(args.input_dir)
    output_path = Path(args.output_dir)
    output_path.mkdir(exist_ok=True)

    # 1. Сначала обрабатываем основной файл
    main_file = input_path / 'capstone_constants.go'
    architectures = find_architectures_in_main_file(main_file)

    if not architectures:
        print("\033[31mError: No architectures found in capstone_constants.go!\033[0m")
        return

    # 2. Затем ищем связанные константы в других файлах
    constants = find_related_constants(input_path, architectures)

    # 3. Генерируем итоговый файл
    generate_go_file(architectures, constants, output_path, args.command)

    print(f"\n\033[32mGenerated sanity_checks.go in {output_path}\033[0m")

if __name__ == '__main__':
    main()
