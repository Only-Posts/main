# main
Основной проект

## CMD
Хранит в себе main.go основа приложение

## Config
Конфиг приложения, туда записываются значения для доступа к другим сервисам

## Internal 
Самописные пакеты для упрощения работы

## Если не работает бд, запускать внутри контейнера
        
goose -dir "./migrations" postgres "postgres://postgres:postgres@db:5432/postgres?sslmode=disable" up
        
Если установлен make и go
        
make migrate
        
