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

TEST_NAME="^${1-}$"

echo -e "${ARROW} Running test ${BLUE}${TEST_NAME}${NC}"
(
    # не использовать кэш
    go clean -testcache -cache

    # Запуск теста с проверкой результата
    if go test -v -timeout 30s -run "${TEST_NAME}"; then
        echo -e "${CHECK} Success test ${BLUE}${TEST_NAME}${NC}"
    else
        echo -e "${CROSS} ${RED}Failed${NC} test ${BLUE}${TEST_NAME}${NC}"
        exit 1 # Выход с ошибкой если тест упал
    fi
) &
spinner Preparing...
