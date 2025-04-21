import os
import re
import argparse
from typing import Dict, List, Tuple, Optional, Set
from pathlib import Path

# TODO: Draft, release in next versions... It will be gendecomposes and its tests
# FIXME: field names (some where we need UPPERCASE some where CamelCase and so on), propper mapping
# v2

# Пример работы:
# Для файла с классами X86OpMem, CsX86Encoding, ArmRegisters 
# скрипт автоматически определит префиксы: "X86", "Cs", "Arm" и будет корректно обрабатывать все эти типы.

# Использование:
# python genstruct.py /path/to/python/files -o /output/directory/for/generated_*.go

import os
import re
import argparse
from typing import Dict, List, Tuple, Optional, Set
from pathlib import Path

class TypeConverter:
    """Класс для конвертации типов Python в Go"""
    
    PY_TO_GO_TYPES = {
        'ctypes.c_uint': 'uint',
        'ctypes.c_uint8': 'uint8',
        'ctypes.c_uint16': 'uint16',
        'ctypes.c_uint32': 'uint32',
        'ctypes.c_uint64': 'uint64',
        'ctypes.c_int': 'int',
        'ctypes.c_int8': 'int8',
        'ctypes.c_int16': 'int16',
        'ctypes.c_int32': 'int32',
        'ctypes.c_int64': 'int64',
        'ctypes.c_bool': 'bool',
        'ctypes.c_float': 'float32',
        'ctypes.c_double': 'float64',
        'ctypes.c_char': 'byte',
        'ctypes.c_char_p': 'string',
        'ctypes.c_void_p': 'unsafe.Pointer',
    }
    
    def __init__(self, class_prefixes: Set[str]):
        self.class_prefixes = class_prefixes
    
    def convert(self, py_type: str) -> str:
        """Конвертирует Python тип в Go тип"""
        py_type = py_type.strip()
        
        # Обработка массивов
        if '[' in py_type and ']' in py_type:
            base_type = py_type.split('[')[0].strip()
            length = py_type.split('[')[1].split(']')[0].strip()
            return f'[{length}]{self.convert(base_type)}'
        
        # Обработка пользовательских типов (по обнаруженным префиксам)
        for prefix in self.class_prefixes:
            if py_type.startswith(prefix):
                return f'*{py_type}'
        
        # Обработка стандартных типов
        return self.PY_TO_GO_TYPES.get(py_type, py_type)

class PythonToGoConverter:
    """Основной класс для конвертации Python классов в Go структуры"""
    
    def __init__(self):
        self.current_file = ""
        self.imports = set()
        self.class_prefixes = set()
    
    def process_directory(self, input_dir: str, output_dir: str):
        """Обрабатывает все Python файлы в директории, пропуская *_const.py"""
        input_path = Path(input_dir)
        output_path = Path(output_dir)
        output_path.mkdir(exist_ok=True)
        
        # TODO: excluede __inint__.py
        # Получаем список файлов, исключая *_const.py
        py_files = [f for f in input_path.glob('*.py') if not f.name.endswith('_const.py')]
        
        if not py_files:
            print(f"No Python files found in {input_dir} (excluding *_const.py)")
            return
        
        # Сначала собираем все префиксы классов
        self._collect_class_prefixes(py_files)
        
        # Затем конвертируем файлы
        for py_file in py_files:
            self.current_file = py_file.stem
            go_file = output_path / f'{self.current_file}.go'
            
            with open(py_file, 'r', encoding='utf-8') as f:
                python_code = f.read()
            
            go_code = self.convert_file(python_code)
            
            with open(go_file, 'w', encoding='utf-8') as f:
                f.write(go_code)
            
            print(f'Processed: {py_file} -> {go_file}')
    
    def _collect_class_prefixes(self, py_files: List[Path]):
        """Собирает префиксы классов из всех файлов (кроме *_const.py)"""
        for py_file in py_files:
            if py_file.name.endswith('_const.py'):
                continue
                
            with open(py_file, 'r', encoding='utf-8') as f:
                python_code = f.read()
            
            # Ищем все классы в файле
            class_pattern = re.compile(r'class\s+([A-Za-z][A-Za-z0-9_]*)')
            for match in class_pattern.finditer(python_code):
                class_name = match.group(1)
                # Определяем префикс (первые символы до первой строчной буквы)
                prefix = ''
                for char in class_name:
                    if char.isupper():
                        prefix += char
                    else:
                        break
                if prefix and len(prefix) > 1:  # Игнорируем однобуквенные префиксы
                    self.class_prefixes.add(prefix)
        
        print(f"Detected class prefixes: {', '.join(self.class_prefixes)}")
    
    def convert_file(self, python_code: str) -> str:
        """Конвертирует содержимое Python файла в Go код"""
        type_converter = TypeConverter(self.class_prefixes)
        classes = self._extract_classes(python_code)
        if not classes:
            return ""
        
        go_code = 'package main\n\n'
        
        # Добавляем необходимые импорты
        if self.imports:
            go_code += 'import (\n'
            for imp in sorted(self.imports):
                go_code += f'\t"{imp}"\n'
            go_code += ')\n\n'
        
        # Генерируем структуры
        for class_name, class_info in classes.items():
            go_code += self._generate_go_struct(class_name, class_info, type_converter)
        
        # Генерируем методы
        for class_name, class_info in classes.items():
            if class_info['methods']:
                go_code += self._generate_go_methods(class_name, class_info, type_converter)
        
        return go_code
    
    def _extract_classes(self, python_code: str) -> Dict[str, Dict]:
        """Извлекает классы из Python кода"""
        classes = {}
        class_pattern = re.compile(
            r'class\s+(\w+)\s*\(([^)]*)\)\s*:\s*_fields_\s*=\s*\(([\s\S]*?)\)\s*(?:\n\s*@|\n\s*def|\n\s*class|\n\s*$)',
            re.MULTILINE
        )
        method_pattern = re.compile(
            r'@property\s*\n\s*def\s+(\w+)\s*\(self\)\s*:\s*return\s+([^\n]+)',
            re.MULTILINE
        )
        
        for class_match in class_pattern.finditer(python_code):
            class_name, bases, fields_str = class_match.groups()
            parent_class = bases.split(',')[0].strip() if bases else 'object'
            
            fields = self._parse_fields(fields_str)
            methods = []
            
            # Извлекаем методы из тела класса
            class_body_start = class_match.end()
            next_class_match = class_pattern.search(python_code, class_body_start)
            class_body_end = next_class_match.start() if next_class_match else len(python_code)
            class_body = python_code[class_body_start:class_body_end]
            
            for method_match in method_pattern.finditer(class_body):
                method_name, return_expr = method_match.groups()
                methods.append((method_name, return_expr.strip()))
            
            classes[class_name] = {
                'parent': parent_class,
                'fields': fields,
                'methods': methods
            }
        
        return classes
    
    def _parse_fields(self, fields_str: str) -> List[Tuple[str, str]]:
        """Парсит строку с полями класса"""
        fields = []
        field_pattern = re.compile(r"\(\s*'([^']+)'\s*,\s*([^)]+)\s*\)")
        
        for field_match in field_pattern.finditer(fields_str):
            name, py_type = field_match.groups()
            fields.append((name.strip(), py_type.strip()))
            
            # Добавляем необходимые импорты
            if 'ctypes.c_void_p' in py_type:
                self.imports.add('unsafe')
        
        return fields
    
    def _generate_go_struct(self, class_name: str, class_info: Dict, type_converter: TypeConverter) -> str:
        """Генерирует Go структуру"""
        go_code = f'type {class_name} struct {{\n'
        
        # Обрабатываем наследование (только для не-ctypes родителей)
        parent = class_info['parent']
        if parent not in ('ctypes.Structure', 'ctypes.Union', 'object'):
            go_code += f'\t{parent}\n'
        
        # Добавляем поля
        for field_name, py_type in class_info['fields']:
            go_type = type_converter.convert(py_type)
            go_field_name = self._snake_to_camel(field_name)
            go_code += f'\t{go_field_name} {go_type} `json:"{field_name}"`\n'
        
        go_code += '}\n\n'
        return go_code
    
    def _generate_go_methods(self, class_name: str, class_info: Dict, type_converter: TypeConverter) -> str:
        """Генерирует Go методы"""
        go_code = ''
        
        for method_name, return_expr in class_info['methods']:
            # Конвертируем выражение return
            go_return_expr = return_expr \
                .replace('self.', 'x.') \
                .replace('value.', 'x.Value.') \
                .replace('mem.', 'x.Mem.')
            
            # Определяем тип возвращаемого значения
            return_type = self._infer_return_type(return_expr, type_converter)
            
            go_method_name = self._snake_to_camel(method_name)
            go_code += f'func (x *{class_name}) {go_method_name}() {return_type} {{\n'
            go_code += f'\treturn {go_return_expr}\n'
            go_code += '}\n\n'
        
        return go_code
    
    def _infer_return_type(self, return_expr: str, type_converter: TypeConverter) -> str:
        """Определяет тип возвращаемого значения по выражению return"""
        parts = return_expr.split('.')
        if len(parts) < 2:
            return 'interface{}'
        
        base = parts[0]
        if base == 'self':
            return 'interface{}'
        
        # Пытаемся определить тип из выражения
        if parts[1].startswith(('imm', 'reg', 'disp', 'scale')):
            return 'int64'
        elif any(parts[1].startswith(prefix.lower()) for prefix in self.class_prefixes):
            # Если возвращается объект с известным префиксом
            return f'*{parts[1].split(".")[0].title()}'
        elif parts[1].startswith(('prefix', 'opcode')):
            return '[]uint8'
        elif 'size' in parts[1] or 'offset' in parts[1]:
            return 'uint8'
        elif 'bool' in parts[1].lower():
            return 'bool'
        
        return 'interface{}'
    
    @staticmethod
    def _snake_to_camel(snake_str: str) -> str:
        """Конвертирует snake_case в CamelCase"""
        parts = snake_str.split('_')
        return ''.join(part.title() for part in parts)

def main():
    parser = argparse.ArgumentParser(description='Convert Python ctypes classes to Go structs (skips *_const.py)')
    parser.add_argument('input_dir', help='Directory with Python files to convert')
    parser.add_argument('-o', '--output', help='Output directory for Go files', default='./output')
    args = parser.parse_args()
    
    converter = PythonToGoConverter()
    converter.process_directory(args.input_dir, args.output)

if __name__ == '__main__':
    main()
