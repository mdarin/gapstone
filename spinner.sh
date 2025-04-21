#!/bin/bash

# INFO: example
# Использование в сценариях:
# source spinner.sh  # если функция в отдельном файле
#
# echo "Начало выполнения скрипта"
# # Запускаем команду в фоне и сразу spinner
# (sleep 3; echo "Действие 1 завершено") & spinner "Выполнение действия 1"
#
# # Другая команда
# (sleep 2; echo "Действие 2 завершено") & spinner "Выполнение действия 2"
#
# echo "Скрипт завершен"

# Функция spinner с Unicode-анимацией
spinner() {
    local pid=$!
    local delay=0.1
    local spinstr='⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏' #'⣾⣽⣻⢿⡿⣟⣯⣷' 
    local text="$1"

    # Убедимся, что терминал поддерживает UTF-8
    export LC_ALL=en_US.UTF-8
    export LANG=en_US.UTF-8

    # Если текст не передан, используем текст по умолчанию
    [[ -z "$text" ]] && text="Processing..."

    # Сохраняем текущую позицию курсора
    tput sc

    while kill -0 "$pid" 2>/dev/null; do
        local temp="${spinstr#?}"
        printf "\r%s %s" "${spinstr:0:1}" "${text}"
        spinstr="$temp${spinstr%"$temp"}"
        sleep "$delay"
        tput rc
    done

    # Очищаем строку после завершения
    printf "\r\033[K"
}

# Пример использования:
# (долгая_команда) & spinner "Описание процесса"

# Демонстрация работы:
# (
#     sleep 3 # Имитация долгой операции
# ) &
# spinner "Загрузка данных..."
