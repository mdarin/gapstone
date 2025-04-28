#!/bin/bash

# Цвета для вывода
GREEN='\033[32m'
GRAY='\033[90m'
RED='\033[31m'
NC='\033[0m' # No Color

# Основная функция
main() {
    local input_path="$1"

    if [ -d "$input_path" ]; then
        process_directory "$input_path"
    elif [ -f "$input_path" ]; then
        process_file "$input_path"
    else
        echo "Error: $input_path is not a valid file or directory"
        exit 1
    fi
}

# toCamelCase (первое слово с маленькой буквы)
toCamelCase() {
    echo "$1" | sed -E 's/(^|_|-|\.|\s)([a-z])/\U\2/g' | sed -E 's/^([A-Z])/\l\1/'
}

# toPascalCase (все слова с заглавной буквы)
toPascalCase() {
    echo "$1" | sed -E 's/(^|_|-|\.|\s)([a-z])/\U\2/g'
}

# функция toSnakeCase преобразует строки в snake_case
toSnakeCase() {
    echo "$1" |
        sed -E 's/([a-z0-9])([A-Z])/\1_\2/g;s/[^[:alnum:]_]+/_/g;s/_+/_/g;s/^_|_$//g' |
        tr '[:upper:]' '[:lower:]'
}

# Удаляет указанный префикс из строки (по умолчанию "test_")
# Параметры:
#   $1 - входная строка
#   $2 - необязательный: кастомный префикс (по умолчанию "test_")
# Возвращает:
#   - строку без префикса (если он был найден)
#   - оригинальную строку (если префикс не найден)
#   - пустую строку (если входная строка пустая)
remove_prefix() {
    local input="${1:-}"
    local prefix="${2:-test_}"
    local preserve_empty="${3:-false}" # Опция сохранения пустого результата

    # Проверка на пустую входную строку
    if [[ -z "$input" ]]; then
        if [[ "$preserve_empty" == "true" ]]; then
            echo ""
        else
            echo "Ошибка: Входная строка пуста" >&2
        fi
        return 1
    fi

    # Проверка на пустой префикс
    if [[ -z "$prefix" ]]; then
        echo "Ошибка: Префикс не может быть пустым" >&2
        return 2
    fi

    # Удаление префикса с проверкой
    if [[ "$input" == "$prefix"* ]]; then
        local result="${input#"$prefix"}"

        # Дополнительная проверка после удаления
        if [[ -z "$result" && "$preserve_empty" != "true" ]]; then
            echo "Ошибка: Результат пуст после удаления префикса" >&2
            return 3
        fi

        echo "$result"
    else
        echo "$input"
    fi
}

# Обработка директории
process_directory() {
    local dir="$1"
    echo "Processing directory: $dir"

    while IFS= read -r -d '' file; do
        # TODO: handle errors
        process_file "$file"
    done < <(find "$dir" -type f -name "*.c" -print0)
}

# Обработка одного файла
process_file() {
    local input_file="$1"
    local script_name="$0"
    local current_date=$(date +"%Y-%m-%d %H:%M:%S")

    # Временные файлы
    local define_file platform_file
    define_file=$(mktemp)
    platform_file=$(mktemp)

    local tmp=$(basename "$input_file")
    tmp=$(remove_prefix "$tmp")
    local resource_name=$(toCamelCase "${tmp%.*}")
    local resource_file="$(toSnakeCase "$resource_name")_resource.go"

    # Извлечение данных
    extract_defines "$input_file" "$define_file"
    extract_platforms "$input_file" "$platform_file" "$resource_name"

    # Вывод результатов
    printf '%s\n\n%s\n\n' "/*
Gapstone is a Go binding for the Capstone disassembly library.
For examples,try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Ben Nagy
	License: BSD style - see LICENSE file for details
    (c) 2013 COSEINC. All Rights Reserved.

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
    Command: ${script_name} ${input_file}
    Created at: ${current_date}

*/" "package gapstone" >"$resource_file"

    printf "//=== Found #define values in $input_file ===\n\n" >>"$resource_file"
    printf "var %sCodes = map[string]string{\n" "$resource_name" >>"$resource_file"
    [ -s "$define_file" ] && cat "$define_file" >>"$resource_file" || echo -e "${GRAY}-${NC} Processing file: ${input_file} ${GRAY}No #define found${NC}"
    printf "}\n\n" >>"$resource_file"

    printf "//=== Found platform entries in $input_file ===\n\n" >>"$resource_file"
    printf "var %sPlatforms = []platform {\n" "$resource_name" >>"$resource_file"
    [ -s "$platform_file" ] && cat "$platform_file" >>"$resource_file" || echo -e "${GRAY}-${NC} Processing file: ${input_file} ${GRAY}No platform entries found${NC}"
    printf "}\n\n" >>"$resource_file"

    echo -e "${GREEN}✓${NC} Processing file: $input_file -> $resource_file"

    # Очистка
    rm -f "$define_file" "$platform_file"
}

# Извлечение #define
extract_defines() {
    awk '
    # функция:
    # \xX → \xX0 (дополняет)
    # \xXX → \xXX (оставляет как есть)
    # Корректно обрабатывает смешанные последовательности
    # Работает в любой версии mawk
    #  Сохраняет все не-hex символы без изменений
    function pad_hex(input,    output, i, len, char, state, hex_chars) {
        output = ""
        len = length(input)
        state = 0  # 0=normal, 1=found backslash, 2=found x, 3=first hex digit
        hex_chars = ""
        
        for (i = 1; i <= len; i++) {
            char = substr(input, i, 1)
            
            if (state == 0 && char == "\\") {
                state = 1
                hex_chars = "\\"
                continue
            }
            
            if (state == 1) {
                if (char == "x") {
                    state = 2
                    hex_chars = hex_chars "x"
                } else {
                    output = output hex_chars char
                    state = 0
                    hex_chars = ""
                }
                continue
            }
            
            if (state == 2) {
                if (char ~ /[0-9a-fA-F]/) {
                    state = 3
                    hex_chars = hex_chars char
                } else {
                    output = output hex_chars char
                    state = 0
                    hex_chars = ""
                }
                continue
            }
            
            if (state == 3) {
                if (char ~ /[0-9a-fA-F]/) {
                    # Полная hex последовательность \xXX
                    output = output hex_chars char
                    state = 0
                    hex_chars = ""
                } else if (char == "\\") {
                    # Началась новая последовательность - дополняем текущую
                    output = output hex_chars "0"
                    state = 1
                    hex_chars = "\\"
                } else {
                    # Только один hex символ - дополняем нулём
                    output = output hex_chars "0" char
                    state = 0
                    hex_chars = ""
                }
                continue
            }
            
            # Обычный символ
            output = output char
        }
        
        # Обработка незавершённой последовательности в конце строки
        if (state == 3) {
            # Был только один hex символ - дополняем
            output = output hex_chars "0"
        }
        else if (state > 0) {
            # Незавершённая последовательность - добавляем как есть
            output = output hex_chars
        }
        
        return output
    }

    # проверка hex-формата
    function is_hex_string(str,    parts, n, i) {
        # Удаляем все пробелы и кавычки для проверки
        clean_str = str
        gsub(/[[:space:]"]/, "", clean_str)
        
        # Должна состоять ТОЛЬКО из \xXX последовательностей
        if (clean_str ~ /^(\\x[0-9a-fA-F]{2})+$/) {
            return 1
        }
        
        # Дополнительная проверка каждой последовательности
        n = split(clean_str, parts, /\\x/)
        for (i = 2; i <= n; i++) {
            if (length(parts[i]) != 2 || parts[i] !~ /[0-9a-fA-F]{2}/) {
                return 1
            }
        }
        
        return 0
    }
    
    function join(strings, n) {
        result = ""
        for (i = 1; i <= n; i++) {
            result = result strings[i]
        }
        return result
    }

    function concat(str1, str2) {
        return str1 str2
    }

    function toCamelCase(str,    i, n, parts, result, sep, word) {
        # Определяем допустимые разделители (подчеркивание, дефис, пробел, точка)
        sep = "[_\\-\\. ]+"
        
        # Разбиваем строку на части по разделителям
        n = split(str, parts, sep)
        if (n == 0) return str  # Если нет разделителей, возвращаем исходную строку
        
        # Обрабатываем первую часть (все символы в нижний регистр)
        word = parts[1]
        if (word != "") {
            result = tolower(substr(word, 1, 1))
            if (length(word) > 1) {
                result = result tolower(substr(word, 2))
            }
        }
        
        # Обрабатываем остальные части
        for (i = 2; i <= n; i++) {
            word = parts[i]
            if (word == "") continue  # Пропускаем пустые части
            
            # Если часть начинается с цифры - добавляем как есть
            if (word ~ /^[0-9]/) {
                result = result word
            } 
            # Иначе - первая буква заглавная, остальные строчные
            else {
                result = result toupper(substr(word, 1, 1))
                if (length(word) > 1) {
                    result = result tolower(substr(word, 2))
                }
            }
        }
        
        # Удаляем возможные оставшиеся разделители
        gsub(sep, "", result)
        return result
    }

    function toPascalCase(str,    camel) {
        camel = toCamelCase(str)
        return toupper(substr(camel, 1, 1)) substr(camel, 2)
    }

    BEGIN {
        in_define = 0
        define_name = ""
        define_value = ""
    }
    
    # Начало многострочного определения
    /^[[:space:]]*#[[:space:]]*define[[:space:]]+/ && /\\[[:space:]]*$/ {
        in_define = 1
        # Извлекаем имя и начало значения
        sub(/^[[:space:]]*#[[:space:]]*define[[:space:]]+/, "")
        define_name = $1
        sub(/^[^[:space:]]+[[:space:]]*/, "")
        define_value = $0
        # Удаляем обратный слеш в конце строки и лишние пробелы
        sub(/\\[[:space:]]*$/, "", define_value)
        gsub(/[[:space:]]+/, " ", define_value)  # Нормализуем пробелы
        next
    }
    
    # Продолжение многострочного определения
    in_define && /\\[[:space:]]*$/ {
        # Добавляем строку к значению (без обратного слеша)
        line = $0
        sub(/\\[[:space:]]*$/, "", line)
        gsub(/[[:space:]]+/, " ", line)  # Нормализуем пробелы
        define_value = define_value line  # Добавляем с пробелом вместо переноса
        next
    }
    
    # Конец многострочного определения
    in_define {
        # Добавляем последнюю строку
        line = $0
        gsub(/[[:space:]]+/, " ", line)  # Нормализуем пробелы
        define_value = define_value line
        # Удаляем комментарии
        gsub(/\/\/.*|\/\*.*\*\//, "", define_value)
        # Удаляем лишние пробелы
        gsub(/^[[:space:]]+|[[:space:]]+$/, "", define_value)
        gsub(/[[:space:]]+/, " ", define_value)  # Заменяем множественные пробелы на один
        gsub(/" "/, "", define_value) 
        # Строгая проверка - должны быть только hex-последовательности с пробелами
        if (is_hex_string(define_value)) {
            # Выводим результат в одну строку
            print "//"  define_name "=" define_value
            printf("\"%s\":%s,\n", define_name, pad_hex(define_value))
        }
        in_define = 0
        define_name = ""
        define_value = ""
        next
    }
    
    # Однострочные определения
    /^[[:space:]]*#[[:space:]]*define[[:space:]]+/ {
        # Удаляем #define и лишние пробелы
        sub(/^[[:space:]]*#[[:space:]]*define[[:space:]]+/, "")
        # Разделяем имя и значение
        name = $1
        sub(/^[^[:space:]]+[[:space:]]*/, "")
        value = $0
        # Удаляем возможные комментарии
        sub(/\/\/.*$/, "", value)
        sub(/\/\*.*\*\//, "", value)
        # Нормализуем пробелы
        gsub(/[[:space:]]+/, " ", value)
        gsub(/^[[:space:]]+|[[:space:]]+$/, "", value)
        # Строгая проверка - должны быть только hex-последовательности с пробелами
        if (is_hex_string(value)) {
            # Выводим в формате имя=значение
            value1 = pad_hex(value)
            print "//" name "=" value1
            printf("\"%s\":%s,\n", name, value1)
        }
    }
    ' "$1" >"$2"
}

# Извлечение структуры platform
extract_platforms() {
    awk -v codes="$3" '
    function sub_or_default(pattern, replacement, default,    str) {
        str = $0  # Work on a copy
        if (sub(pattern, replacement, str)) {
            return str
        } else {
            return default
        }
    }

    function clean_field(field) {
        # Удаляем пробелы и приведение типов
        gsub(/^[[:space:]]*(\([^)]*\))?[[:space:]]*/, "", field)
        gsub(/[[:space:]]*$/, "", field)
        return field
    }

    BEGIN {
        in_platform_array = 0
        in_entry = 0
        entry = ""
        entry_start_line = 0
    }
    
    # Начало блока platform
    /struct[[:space:]]+platform[[:space:]]+platforms\[.*\]/ {
        in_platform_array = 1
        next
    }
    
    # Конец блока platform
    in_platform_array && /\};/ {
        in_platform_array = 0
        next
    }
    
    # Обработка внутри блока platform
    in_platform_array {
        # Пропускаем строки с директивами препроцессора
        if (/^[[:space:]]*#/) next
        
        # Находим начало записи
        if (/{/) {
            in_entry = 1
            entry = ""
            entry_start_line = NR
            opt_count = 0
            delete opts_type
            delete opts_value
            next
        }
        
        # Собираем многострочные записи
        if (in_entry) {
            # Удаляем лишние пробелы в начале строки
            sub(/^[[:space:]]+/, "")
            entry = entry " " $0
        }
        
        # Конец записи
        if (/}/ && in_entry) {
            in_entry = 0
            
            # Очищаем запись от комментариев
            gsub(/\/\/.*/, "", entry)
            gsub(/\/\*.*\*\//, "", entry)
            
            # Удаляем лишние пробелы и переносы строк
            gsub(/[[:space:]]+/, " ", entry)
            gsub(/,[[:space:]]+/, ",", entry)
            
            # Удаляем фигурные скобки
            gsub(/[{}]/, "", entry)
            
            # Разделяем поля
            n = split(entry, fields, ",")
            
            # Извлекаем основные поля (минимум 5 полей)
            if (n >= 5) {
                arch = clean_field(fields[1])
                mode = clean_field(fields[2])
                code = clean_field(fields[3])
                size = clean_field(fields[4])
                comment = clean_field(fields[5])
                
                # Обработка опций (поля 6+)
                for (i = 6; i <= n; i += 2) {
                    if (i+1 <= n) {
                        opt_type = clean_field(fields[i])
                        opt_value = clean_field(fields[i+1])
                        if (opt_type != "" && opt_value != "") {
                            opts_type[++opt_count] = opt_type
                            opts_value[opt_count] = opt_value
                        }
                    }
                }
                
                # Выводим запись
                printf("{\n")
                print "// -------------------------------------------"
                print "// Platform Entry (line " entry_start_line "):"
                print "// Architecture: " arch
                print "// Mode: " mode
                print "// Code: " code
                # print "// Size: " size
                print "// Comment: " comment
                
                # Выводим опции, если они есть
                if (opt_count > 0) {
                    print "//  Options:"
                    for (i = 1; i <= opt_count; i++) {
                        print "//    " opts_type[i] ": " opts_value[i]
                    }
                }
                # print ""
                # генерируем golang код
              
                printf("\tarch: %s,\n", arch)
                printf("\tmode: %s,\n", mode)
                printf("\tcode: %sCodes[\"%s\"],\n", codes, code)
                printf("\tcomment: %s,\n", comment)
                printf("\toptions: []option{\n")
                printf("\t\t{CS_OPT_DETAIL, CS_OPT_ON},\n")
                # Для опций:
                for (i = 1; i <= opt_count; i++) {
                    printf("{%s, %s},\n", opts_type[i], opts_value[i])
                }
                printf("\t},\n")
              
                printf("},\n")



               
            }
        }
    }
    ' "$1" >"$2"
}

# Проверка аргументов
if [ $# -ne 1 ]; then
    echo "Usage: $0 <input_file.c or directory>"
    exit 1
fi

# Запуск основной функции
main "$1"
