@echo off
echo Шаг 1: Поднимаем docker-compose
docker-compose --file data/docker/docker-compose.yaml up -d

echo Шаг 2: Ждём немного, чтобы сервисы успели стартовать (опционально)
timeout /t 5 >nul

echo Шаг 3: Компилируем тесты
go test -c ./cmd -o ./bin/test.exe

echo Шаг 4: Запускаем тесты
.\bin\test.exe -test.v

echo Шаг 5: Выключаем конты
docker-compose --file data/docker/docker-compose.yaml down