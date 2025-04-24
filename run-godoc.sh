#!/bin/bash

# * После запуска godoc будет доступен по адресу: 
# * http://localhost:6060/pkg/github.com/mdarin/gapstone 
# (name in go.mod)

# Запуск с поддержкой модулей
godoc -http=:6060



# Полезные команды

# Команда	Описание
# godoc fmt	Документация пакета fmt
# godoc -src fmt Printf	Исходный код Printf
# godoc -ex sync Mutex	Примеры для Mutex