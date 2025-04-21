#!/bin/bash

source spinner.sh

# Цвета и символы
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # no colour
CHECK="${GREEN}✓${NC}"
CROSS="${RED}✗${NC}"
ARROW="${YELLOW}➔${NC}"

# Функция для поиска и форматированного вывода *_test.go файлов
find_test_files() {
    local test_files
    
    # Поиск файлов только в текущем каталоге (без рекурсии) и сохранение в переменную с разделителем \n
    test_files=$(find . -maxdepth 1 -type f -name '*_test.go' | sort | tr '\n' '\0' | xargs -0 -I{} realpath --relative-to=. "{}")
    
    # Возвращаем результат
    echo "$test_files"
}

# Пример использования:
print_test_files_list() {
    local files_list=$(find_test_files)
    local count=$(echo "$files_list" | wc -l)
    
    if [[ $count -eq 0 ]]; then
        echo "⛔ No *_test.go files found in current directory and subdirectories"
        return 1
    fi
    
    echo "📋 Found $count test files:"
    echo "┌───────────────────────────────────────"
    echo "$files_list" | awk '{print "│ " $0}'
    echo "└───────────────────────────────────────"
}

# Вызываем функцию получения списка тестов
echo -e "${ARROW} Running tests:"
print_test_files_list

(
    # не использовать кэш
    go clean -testcache -cache

    # Запуск теста с проверкой результата
    if go test -v -timeout 30s ./...; then
        echo -e "${CHECK} Success"
    else
        echo -e "${CROSS} ${RED}Failed${NC}"
        exit 1 # Выход с ошибкой если тест упал
    fi
) &
spinner Preparing...
