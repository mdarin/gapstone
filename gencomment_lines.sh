#!/bin/bash

# If you see someting like this on run the command go test ./...
# ./riscv_constants.go:471:22: could not determine kind of name for C.RISCV_GRP_ISRV32C
# ./riscv_constants.go:472:23: could not determine kind of name for C.RISCV_GRP_ISRV32CF
# ./riscv_constants.go:474:22: could not determine kind of name for C.RISCV_GRP_ISRV64A
# ./riscv_constants.go:475:22: could not determine kind of name for C.RISCV_GRP_ISRV64C
# ./riscv_constants.go:476:22: could not determine kind of name for C.RISCV_GRP_ISRV64D
# try to fix it by this script

# set -eux

# не забываем установить команду 'file'
# sudo apt update && sudo apt install -y file

echo 'Commenting out unused stuff...'

# Цвета для вывода
GREEN='\033[32m'
GRAY='\033[90m'
RED='\033[31m'
NC='\033[0m' # No Color

# Функция для замены с проверкой
perform_sed() {
    local file="$1"
    local pattern="$2"
    local comment="$3"
    local temp_file
    local changed=false

    temp_file=$(mktemp)

    # Строгий шаблон для точного совпадения:
    # 1. \b - граница слова (или [^[:alnum:]_] для POSIX)
    # 2. Точное совпадение pattern
    # 3. Проверка, что после нет букв/цифр/подчеркивания
    strict_pattern="(^[[:space:]]*($PATTERNS))"

    if sed -E "s@${strict_pattern}@${comment}\1@" "$file" >"$temp_file" 2>/dev/null; then
        if ! cmp -s "$file" "$temp_file"; then
            mv "$temp_file" "$file"
            changed=true
        else
            rm "$temp_file"
        fi
    else
        echo -e "${RED}✗${NC} $file (sed error)" >&2
        return 1
    fi

    if $changed; then
        echo -e "${GREEN}✓${NC} $file"
    else
        echo -e "${GRAY}-${NC} $file (no matches)"
    fi
}

# Указываем каталог для поиска (по умолчанию — текущая папка)
DIR="${1:-.}"

# Строки, которые нужно закомментировать (через | для regex)
PATTERNS=""
while IFS= read -r line; do
    pattern=$(awk -F'C\\.' '{print $2}' <<<"$line" | awk -F'[^A-Za-z0-9_]' '{print $1}')
    [[ ! "$pattern" =~ [^[:space:]] ]] && continue
    [ -z "$PATTERNS" ] && PATTERNS="$pattern" || PATTERNS+="|$pattern"
done < <(go test ./... 2>&1)

if [-e "$DIR" ] 

echo "Поиск в: $DIR"
echo "Шаблоны: $PATTERNS"

[ -z "$PATTERNS" ] && echo "Ошибка: не предоставлены шаблоны для исключения" && exit 1

# Символ комментария
COMMENT="//"

# Поиск файлов и замена строк
while IFS= read -r -d '' file; do
    # file="./riscv_constants.go"
    # echo "Проверка файла: $file"
    # Проверяем, является ли файл текстовым (избегаем бинарных файлов)
    # if file "$file" | grep -q "text"; then
    #     echo "Обработка файла: $file"
    # Заменяем строки, содержащие PATTERNS, на закомментированные
    # sed -i -E "s@(.*$PATTERNS.*)@$COMMENT \1@g" "$file"
    #  sed -i -E "s@(^[[:space:]]*$PATTERNS[[:space:]][^[[:space:]]]*)@$COMMENT \1@" "$file" # поставить коментарий однкратно(одно совпадение в строке)
    [ -f "$file" ] || continue
    perform_sed "$file" "$PATTERNS" "$COMMENT"
    # else
    #     echo "Пропуск файла (не текст): $file"
    # fi
done < <(find "$DIR" -maxdepth 1 -type f ! -path "$DIR/$0" -name '*_constants.go' -print0) # не трогать этот скрипт

echo "Готово! Строки закомментированы."
