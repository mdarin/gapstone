#!/bin/bash

#FIXME: struct names and empty fields...


# Цвета для вывода
GREEN='\033[32m'
GRAY='\033[90m'
RED='\033[31m'
BLUE='\033[34m'
NC='\033[0m' # No Color

# Функция для преобразования типов ctypes в Go типы
map_type() {
    local py_type="$1"

    case "$py_type" in
    c_uint) echo "uint32" ;;
    c_int) echo "int32" ;;
    c_uint8) echo "uint8" ;;
    c_int8) echo "int8" ;;
    c_uint16) echo "uint16" ;;
    c_int16) echo "int16" ;;
    c_uint32) echo "uint32" ;;
    c_int32) echo "int32" ;;
    c_uint64) echo "uint64" ;;
    c_int64) echo "int64" ;;
    c_bool) echo "bool" ;;
    c_char) echo "byte" ;;
    c_float) echo "float32" ;;
    c_double) echo "float64" ;;
    *Array*) # Обработка массивов
        local inner_type=$(echo "$py_type" | sed -n 's/.*Array(\([^)]*\)).*/\1/p')
        local array_size=$(echo "$py_type" | sed -n 's/.*Array([^,]*, *\([^)]*\)).*/\1/p')
        echo "[$array_size]$(map_type "$inner_type")"
        ;;
    *Structure)
        * # Вложенные структуры
        echo "$py_type" | sed -n 's/.*(\([^)]*\)).*/\1/p'
        ;;
    *Union*) echo "interface{} // Union" ;;
    *) echo "interface{} // Unknown type: $py_type" ;;
    esac
}

# Функция для преобразования snake_case в CamelCase
to_camel_case() {
    echo "$1" | awk '{
        $0 = tolower($0)
        gsub(/_/, " ", $0)
        gsub(/(^| )([a-z])/, "\\1\\u&", $0)
        gsub(/ /, "", $0)
        print $0
    }'
}

# Функция обработки файла
process_file() {
    local python_file="$1"
    local output_dir="${2:-.}"
    local go_file="${output_dir}/$(basename "${python_file%.*}")_gen.go"
    local package_name="gapstone" #$(basename "${python_file%.*}")

    echo -e "${BLUE}Processing:${NC} $python_file -> $go_file"

    awk -v package="$package_name" '
    function map_type(py_type,    arr, element_type, array_size) {
        # Базовые типы
        if (py_type ~ /^c_uint[0-9]*$/) {
            if (py_type == "c_uint") return "uint32"
            return "uint" substr(py_type, 7)
        }
        if (py_type ~ /^c_int[0-9]*$/) {
            if (py_type == "c_int") return "int32"
            return "int" substr(py_type, 6)
        }
        if (py_type == "c_bool") return "bool"
        if (py_type == "c_char") return "byte"
        if (py_type == "c_float") return "float32"
        if (py_type == "c_double") return "float64"
        
        # Массивы
        if (py_type ~ /^Array\(.*\)$/) {
            sub(/^Array\(/, "", py_type)
            sub(/\)$/, "", py_type)
            split(py_type, arr, /,\s*/)
            element_type = map_type(arr[1])
            array_size = arr[2]
            return "[" array_size "]" element_type
        }
        
        # Структуры
        if (py_type ~ /^Structure\(.*\)$/) {
            sub(/^Structure\(/, "", py_type)
            sub(/\)$/, "", py_type)
            return py_type  # Возвращаем имя структуры
        }
        
        # Объединения
        if (py_type ~ /^Union\(.*\)$/) return "interface{} // Union"
        
        return "interface{} // Unknown type: " py_type
    }

    function to_camel_case(s,    words, result, i, n) {
        gsub(/_/, " ", s)
        n = split(s, words, " ")
        result = ""
        for (i = 1; i <= n; i++) {
            if (words[i] != "") {
                result = result toupper(substr(words[i], 1, 1)) tolower(substr(words[i], 2))
            }
        }
        return result
    }

    BEGIN {
        print "// Code generated from " FILENAME "; DO NOT EDIT."
        print "//go:generate bash ./convert.sh " FILENAME
        print ""
        print "package " package
        print ""
        print "import ("
        print "    \"unsafe\""
        print "    \"reflect\""
        print ")"
        print ""
        struct_count = 0
    }

    /^class.*Structure\)/ {
        if (struct_count++ > 0) print ""
        struct_name = $2
        print "// " struct_name " mirrors the C structure"
        print "type " struct_name " struct {"
        in_struct = 1
        next
    }

    in_struct && /_fields_/ {
        in_fields = 1
        next
    }

    in_fields && /\(\x27[^\x27]+\x27,\s*ctypes\.[^\)]+\)/ {
        # Извлекаем имя поля
        field_start = index($0, "\x27") + 1
        field_end = index(substr($0, field_start), "\x27") + field_start - 1
        field_name = substr($0, field_start, field_end - field_start)
        
        # Извлекаем тип поля
        type_start = index($0, "ctypes.") + 7
        type_end = index(substr($0, type_start), ")") + type_start - 1
        field_type = substr($0, type_start, type_end - type_start)
        
        # Очищаем тип от возможных пробелов в конце
        sub(/\s*$/, "", field_type)
        
        # Преобразование
        camel_name = to_camel_case(field_name)
        go_type = map_type(field_type)
        
        # Вывод поля структуры
        printf "    %-20s %-30s", camel_name, go_type
        if (go_type ~ /Unknown/) printf "// Original: " field_type
        printf "\n"
        next
    }

    in_fields && /\)/ {
        in_fields = 0
        next
    }

    in_struct && /\)/ {
        # Добавляем методы для удобства работы
        print "}"
        print ""
        print "// Size returns the size of the structure in bytes"
        print "func (s *" struct_name ") Size() int {"
        print "    return int(unsafe.Sizeof(*s))"
        print "}"
        print ""
        print "// ReflectType returns the reflect.Type of the structure"
        print "func (s *" struct_name ") ReflectType() reflect.Type {"
        print "    return reflect.TypeOf(*s)"
        print "}"
        in_struct = 0
        next
    }

    /^class.*Union\)/ {
        if (struct_count++ > 0) print ""
        union_name = $2
        print "// " union_name " mirrors the C union"
        print "type " union_name " struct {"
        print "    data [8]byte // Union fields - use appropriate getter methods"
        print "}"
        print ""
        print "// Size returns the size of the union in bytes"
        print "func (u *" union_name ") Size() int {"
        print "    return len(u.data)"
        print "}"
        next
    }

    END {
        if (struct_count == 0) {
            print "// No structures found in the input file"
        } else {
            print "// Total structures processed:", struct_count
        }
    }
    ' "$python_file" >"$go_file"

    # Форматируем результат с помощью gofmt
    if command -v gofmt >/dev/null; then
        if ! gofmt -w "$go_file"; then
            echo -e "${RED}✗${NC} Error formatting $go_file"
            return 1
        fi
        echo -e "${GREEN}✓${NC} Successfully generated $go_file"
    else
        echo -e "${GRAY}⚠${NC} gofmt not found, skipping formatting"
    fi
}

# Основной цикл обработки файлов
main() {
    local input_dir="${@:-.}"
    local output_dir="./generated"
    mkdir -p "$output_dir"

    if [ $# -eq 1 ]; then
        echo -e "${BLUE}Processing all .py files in ${input_dir} directory${NC}"
        while IFS= read -r -d '' python_file; do
            [ -f "$python_file" ] || continue
            if ! process_file "$python_file" "$output_dir"; then
                echo -e "${RED}Error processing $python_file${NC}" >&2
            fi

            # Если вам нужно исключить и другие шаблоны имен файлов, можно добавить дополнительные условия:
            # ! -name '*_const.py' ! -name '*_test.py' ! -name '*_mock.py'
        done < <(find "$input_dir" -maxdepth 1 -type f ! -name '*_const.py' -name '*.py' -print0)
    else
        for python_file in "$@"; do
            [ -f "$python_file" ] || {
                echo -e "${RED}Error: File $python_file not found${NC}" >&2
                continue
            }
            if ! process_file "$python_file" "$output_dir"; then
                echo -e "${RED}Error processing $python_file${NC}" >&2
            fi
        done
    fi

    echo -e "${GREEN}✓ Conversion completed successfully${NC}"
    echo -e "Generated files are in ${BLUE}$output_dir${NC}"
}

main "$@"
