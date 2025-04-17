#!/bin/bash

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

# Обработка директории
process_directory() {
    local dir="$1"
    echo "Processing directory: $dir"
    
    while IFS= read -r -d '' file; do
        process_file "$file"
    done < <(find "$dir" -type f -name "*.c" -print0)
}

# Обработка одного файла
process_file() {
    local input_file="$1"
    echo -e "\nProcessing file: $input_file"
    
    # Временные файлы
    local define_file platform_file
    define_file=$(mktemp)
    platform_file=$(mktemp)
    
    # Извлечение данных
    extract_defines "$input_file" "$define_file"
    extract_platforms "$input_file" "$platform_file"
    
    # Вывод результатов
    echo "=== Found #define values in $input_file ==="
    [ -s "$define_file" ] && cat "$define_file" || echo "No #define found"
    
    echo -e "\n=== Found platform entries in $input_file ==="
    [ -s "$platform_file" ] && cat "$platform_file" || echo "No platform entries found"
    
    # Очистка
    rm -f "$define_file" "$platform_file"
}

# Извлечение #define
extract_defines() {
    awk '
    /#define/ {
        # Удаляем #define и лишние пробелы
        sub(/^[[:space:]]*#define[[:space:]]+/, "")
        # Разделяем имя и значение
        split($0, parts, /[[:space:]]+/)
        name = parts[1]
        sub(name "[[:space:]]*", "")
        value = $0
        # Удаляем возможные комментарии
        sub(/\/\/.*$/, "", value)
        sub(/\/\*.*\*\//, "", value)
        gsub(/^[[:space:]]+|[[:space:]]+$/, "", value)
        # Выводим в формате имя=значение
        print name "=" value
    }
    ' "$1" > "$2"
}

# Извлечение структуры platform
extract_platforms() {
    awk '
    # Начало блока platform
    /struct platform platforms\[\]/,/\};/ {
        # Пропускаем строки с #ifdef/#endif
        if ($0 ~ /^[[:space:]]*#/) next
        
        # Находим начало записи
        if ($0 ~ /{[^}]*$/) {
            in_entry = 1
            entry = ""
        }
        
        # Собираем многострочные записи
        if (in_entry) {
            # Удаляем лишние пробелы в начале строки
            sub(/^[[:space:]]*/, "")
            entry = entry $0
        }
        
        # Конец записи
        if ($0 ~ /}[^}]*$/) {
            in_entry = 0
            # Очищаем запись от комментариев
            gsub(/\/\/.*/, "", entry)
            gsub(/\/\*.*\*\//, "", entry)
            # Удаляем лишние пробелы и переносы строк
            gsub(/[[:space:]]+/, " ", entry)
            gsub(/, /, ",", entry)
            # Удаляем фигурные скобки
            gsub(/[{}]/, "", entry)
            # Разделяем поля
            split(entry, fields, ",")
            
            # Извлекаем значения полей
            arch = fields[1]
            mode = fields[2]
            code = fields[3]
            sub(/^[[:space:]]*/, "", code)
            sub(/[[:space:]]*$/, "", code)
            size = fields[4]
            comment = fields[5]
            
            # Удаляем приведение типа (unsigned char *)
            sub(/\(unsigned char \*\)/, "", code)
            
            # Выводим запись
            print "Platform Entry:"
            print "  Architecture: " arch
            print "  Mode: " mode
            print "  Code: " code
            print "  Size: " size
            print "  Comment: " comment
            print ""
        }
    }
    ' "$1" > "$2"
}

# Проверка аргументов
if [ $# -ne 1 ]; then
    echo "Usage: $0 <input_file.c or directory>"
    exit 1
fi

# Запуск основной функции
main "$1"
