#!/bin/bash
# set -eux

# ! Unused, but works good

# скрипт для генерации такого же префикса
generate_prefix() {
    local script_name="$0"
    local input_file="$1"
    local from_spec="$2"
    local current_date=$(date +"%Y-%m-%d %H:%M:%S")

    cat <<END
/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples, try reading the *_test.go files.

    Library Author: Nguyen Anh Quynh
    Binding Author: Ben Nagy
    License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
    Command: $script_name $input_file
    Created at: $current_date
    From SPEC: $from_spec

*/

package gapstone

// #cgo LDFLAGS: -lcapstone
// #cgo freebsd CFLAGS: -I/usr/local/include
// #cgo freebsd LDFLAGS: -L/usr/local/lib
// #include <stdlib.h>
// #include <capstone/capstone.h>
import "C"

END
}

# Пример использования:
# generate_prefix "input_argument" > output_file.go

# Функция для обработки файла с дизассемблированным кодом
process_spec_file() {
    local file="$1"
    local current_platform=""
    local current_code=""
    local current_disasm=""
    local in_disasm=false

    while IFS= read -r line; do
        # Определяем начало нового блока
        if [[ $line =~ ^\*{3,} ]]; then
            # Если у нас есть накопленные данные - выводим их
            if [[ -n $current_platform && -n $current_code && -n $current_disasm ]]; then
                print_record "$current_platform" "$current_code" "$current_disasm"
            fi
            # Сбрасываем переменные для нового блока
            current_platform=""
            current_code=""
            current_disasm=""
            in_disasm=false
            continue
        fi

        # Парсим платформу
        if [[ $line =~ ^Platform:[[:space:]]*(.*)$ ]]; then
            current_platform="${BASH_REMATCH[1]}"
            continue
        fi

        # Парсим код
        if [[ $line =~ ^Code:[[:space:]]*(.*)$ ]]; then
            current_code="${BASH_REMATCH[1]}"
            continue
        fi

        # Определяем начало секции дизассемблирования
        if [[ $line =~ ^Disasm:$ ]]; then
            in_disasm=true
            continue
        fi

        # Собираем дизассемблированный код
        if $in_disasm; then
            current_disasm+="$line"$'\n'
        fi
    done <"$file"

    # Обработка последнего блока, если файл не заканчивается разделителем
    if [[ -n $current_platform && -n $current_code && -n $current_disasm ]]; then
        print_record "$current_platform" "$current_code" "$current_disasm"
    fi
}

# Вспомогательная функция для вывода записи
print_record() {
    local platform="$1"
    local code="$2"
    local disasm="$3"

    echo "=== Record ==="
    echo "Platform: $platform"
    echo "Code: $code"
    # echo "Disasm:"
    # echo "$disasm"
}

# Пример использования функции
if [[ $# -eq 0 ]]; then
    echo "Usage: $0 <asm_file>"
    exit 1
fi

# функция, которая обрабатывает строки, оставляя только буквы, цифры и заменяя пробелы на подчёркивания
sanitize_string() {
    # Проверка на пустую входную строку
    [[ -z "$1" ]] && {
        echo ""
        return 0
    }

    local str="$1"

    # 1. Заменяем все пробелы и табы на подчеркивания
    str="${str//[[:space:]]/_}"

    # 2. Удаляем все символы, кроме букв, цифр, подчеркиваний
    str=$(echo "$str" | tr -cd '[:alnum:]_')

    # 3. Удаляем повторяющиеся подчеркивания и дефисы
    str=$(echo "$str" | sed -E 's/_+/_/g; s/-+/-/g')

    # 4. Удаляем подчеркивания и дефисы в начале и конце строки
    str=${str#[-_]} # В начале
    str=${str%[-_]} # В конце

    # 5. Опционально: приводим к нижнему регистру
    # str=$(echo "$str" | tr '[:upper:]' '[:lower:]')

    # 6. Проверка на пустую строку после всех преобразований
    [[ -z "$str" ]] && str="empty"

    echo "$str"
    return 0
}

# функция, которая обрабатывает строку с шестнадцатеричными значениями и форматирует их через запятую, проверяя первый и последний элементы
format_hex_values() {
    local input="$1"
    local separator="$2"
    local result=""

    # Удаляем кавычки и лишние пробелы в начале/конце
    input=$(echo "$input" | sed -e 's/^"//' -e 's/"$//' -e 's/^[[:space:]]*//' -e 's/[[:space:]]*$//')

    # Разбиваем строку по пробелам
    IFS=' ' read -ra values <<<"$input"

    # Обрабатываем каждый элемент
    for ((i = 0; i < ${#values[@]}; i++)); do
        # Проверяем, что значение не пустое
        if [[ -n "${values[i]}" ]]; then
            # Добавляем значение без запятой(separator) перед первым элементом
            if [[ -z "$result" ]]; then
                result="${values[i]}"
            else
                result+="${separator}${values[i]}"
            fi
        fi
    done

    echo "$result"
}

# функция с проверками для извлечения названия файла из строки вида ./sparc.SPEC
extract_spec_name() {
    [[ -z "$1" ]] && {
        echo "Ошибка: Пустая входная строка" >&2
        return 1
    }
    [[ "$1" != *".SPEC" ]] && {
        echo "Ошибка: Неверный формат. Ожидается <filename>.SPEC" >&2
        return 1
    }

    local filename=$(basename "$1" .SPEC)
    [[ -n "$filename" ]] && echo "$filename" || {
        echo "Ошибка: Не удалось извлечь имя файла" >&2
        return 1
    }
}

# Выводим результат в требуемом формате
generate_output() {
    local input_file="$1"
      local array_name="$2"  # Имя переменной массива
    local separator="${3:-}"           # Опциональный разделитель, по умолчанию пустой

        # Создаем локальную ссылку на массив
    declare -n platform_codes_ref="$array_name"

    # 1. Генерация имени выходного файла
    local output_file_name=$(extract_spec_name "$input_file")
    output_file="${output_file_name}_codes.go" || return 1
    echo "$output_file"

    # 2. Генерация префикса
    generate_prefix "$PWD" "$input_file" >"$output_file"

    # 3. Пустая строка-разделитель
    echo "" >>"$output_file"
    echo "const(" >>"$output_file"

    # 4. Итерация по платформам
    local first_iteration=true
    for platform in "${!platform_codes_ref[@]}"; do
        # Для всех элементов кроме первого добавляем разделитель
        if ! $first_iteration && [[ -n "$separator" ]]; then
            echo "$separator" >>"$output_file"
        fi
        first_iteration=false

        # Санитизация и форматирование
        local sanitized_platform formatted_codes
        sanitized_platform=$(sanitize_string "$platform") || continue
        formatted_codes=$(format_hex_values "${platform_codes_ref[$platform]}") || continue

        # Вывод результата
        printf "code%s__%s = \"%s\"\n" "$sanitized_platform" "$output_file_name" "$formatted_codes" >>"$output_file"
    done

    echo ")" >>"$output_file"

    # 5. Завершающая пустая строка
    echo "" >>"$output_file"
}

process_asm_records() {
    local input_file="$1"
    declare -A platform_codes # Ассоциативный массив для хранения кодов по платформам

    # Переменные для хранения текущей записи
    local current_platform=""
    local current_code=""

    while IFS= read -r line; do
        # Определяем начало новой записи
        if [[ "$line" == "=== Record ===" ]]; then
            # Если есть предыдущая запись - сохраняем
            if [[ -n "$current_platform" && -n "$current_code" ]]; then
                platform_codes["$current_platform"]+="$current_code"
            fi
            current_platform=""
            current_code=""
            continue
        fi

        # Извлекаем название платформы
        if [[ "$line" =~ ^Platform:[[:space:]]+(.*)$ ]]; then
            current_platform="${BASH_REMATCH[1]}"
            continue
        fi

        # Извлекаем код
        if [[ "$line" =~ ^Code:[[:space:]]+(.*)$ ]]; then
            # Удаляем "0x" из каждого байта и преобразуем в \x формат
            current_code="$(echo "${BASH_REMATCH[1]}" | sed 's/0x/\\x/g')"
            continue
        fi
    done < <(process_spec_file "$input_file")

    # Добавляем последнюю запись
    if [[ -n "$current_platform" && -n "$current_code" ]]; then
        platform_codes["$current_platform"]+="$current_code"
    fi

    # Выводим результат в требуемом формате
    # echo "$(extract_spec_name $input_file)_codes.go"
    # echo "$(generate_prefix "$PWD")"
    # echo ""
    # separator=""
    # for platform in "${!platform_codes[@]}"; do
    #     sanitized_platform=$(sanitize_string "$platform")
    #     formatted_codes=$(format_hex_values "${platform_codes[$platform]}" "$separator")
    #     echo "code${sanitized_platform} = \"${formatted_codes}\""
    # done
    # echo ""
    generate_output "$input_file" platform_codes
}

process_spec_files() {
    # local files=()
    while IFS= read -r -d '' file; do
        # files+=("$file")
        # echo "$file"
        # process_spec_file "$file"
        process_asm_records "$file"
    done < <(find "$1" -maxdepth 1 -type f -name "*.SPEC" -print0)
    # printf '%s\n' "${files[@]}"
}

process_spec_files $1

# Использование:
# spec_files=($(get_spec_files .))
# printf 'spec_files: %s\n', "${spec_files[@]}"

# # Пример использования
# if [[ $# -eq 0 ]]; then
#     echo "Usage: $0 <input_file>"
#     exit 1
# fi

# process_asm_records "$1"
