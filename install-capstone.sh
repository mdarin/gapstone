#!/bin/bash

# * Установка Capstone Engine 5.0
set -e  # Прерывать выполнение при ошибках

# Цвета и символы для красивого вывода
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[0;33m'
GRAY='\033[90m'
NC='\033[0m' # No Color
CHECK="${GREEN}✓${NC}"
CROSS="${RED}✗${NC}"
INFO="${YELLOW}ℹ${NC}"

# Проверяем, нужно ли использовать sudo
SUDO=''
if ((EUID != 0)); then
    SUDO='sudo'
fi

# URL для загрузки
CAPSTONE_VERSION="5.0.6" # TODO: chose version manualy here
CAPSTONE_URL="https://github.com/capstone-engine/capstone/releases/download/${CAPSTONE_VERSION}/capstone-${CAPSTONE_VERSION}.tar.xz"
TEMP_DIR=$(mktemp -d)
INSTALL_DIR="/usr/local"  # Стандартный каталог для установки библиотек # TODO: chose dir if needed here

# Функция для вывода сообщений с отступом
status() {
    echo -e " ${1} ${2}"
}

# Функция для проверки команд
check_command() {
    if command -v "$1" &> /dev/null; then
        status "$CHECK" "Команда $1 доступна"
        return 0
    else
        status "$CROSS" "Команда $1 не найдена"
        return 1
    fi
}

# Убедимся, что установлены необходимые зависимости
status "$INFO" "Проверка зависимостей..."
if ! check_command "curl"; then
    status "$INFO" "Установка curl..."
    $SUDO apt-get install -y curl
fi

if ! check_command "tar"; then
    status "$INFO" "Установка tar..."
    $SUDO apt-get install -y tar
fi

# Скачиваем и распаковываем Capstone
status "$INFO" "Загрузка Capstone..."
curl -L "$CAPSTONE_URL" -o "$TEMP_DIR/capstone.tar.xz" && {
    status "$CHECK" "Файл успешно загружен"
} || {
    status "$CROSS" "Ошибка загрузки!"
    exit 1
}

status "$INFO" "Распаковка архива..."
tar -xf "$TEMP_DIR/capstone.tar.xz" -C "$TEMP_DIR" && {
    status "$CHECK" "Архив успешно распакован"
} || {
    status "$CROSS" "Ошибка распаковки!"
    exit 1
}

# Copy sources to root gapstone directory to further parsing
cp -r "$TEMP_DIR"/capstone-* /workspaces/gapstone/capstone

# Переходим в директорию с исходниками
cd "$TEMP_DIR"/capstone-* && {
    status "$CHECK" "Директория с исходниками найдена"
} || {
    status "$CROSS" "Не удалось найти директорию с исходниками!"
    exit 1
}

# Компиляция с поддержкой нужных архитектур
status "$INFO" "Компиляция Capstone..."
if CAPSTONE_ARCHS="aarch64 x86" ./make.sh; then
    status "$CHECK" "Компиляция успешно завершена"
else
    status "$CROSS" "Ошибка компиляции!"
    exit 1
fi

# Установка
status "$INFO" "Установка в систему..."
if $SUDO ./make.sh install; then
    status "$CHECK" "Установка завершена успешно"
else
    status "$CROSS" "Ошибка установки!"
    exit 1
fi

# Обновляем кэш динамических библиотек
status "$INFO" "Обновление кэша библиотек..."
$SUDO ldconfig && {
    status "$CHECK" "Кэш библиотек обновлен"
} || {
    status "$CROSS" "Ошибка обновления кэша!"
}

# Очистка
status "$INFO" "Очистка временных файлов..."
rm -rf "$TEMP_DIR" && {
    status "$CHECK" "Временные файлы удалены"
} || {
    status "$CROSS" "Ошибка при очистке!"
}

echo -e "\n${GREEN}✔ Успешно установлено!${NC}"
CAPSTONE_LIB=$(ldconfig -p | grep libcapstone)
echo -e "${YELLOW}Версия: ${CAPSTONE_VERSION}${NC}"
echo -e "${GRAY}Библиотека:${CAPSTONE_LIB}${NC}"
