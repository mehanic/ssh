#!/bin/bash

USER_NAME=""

if [ ! -c /dev/ptmx ]; then
    echo "Помилка: /dev/ptmx не знайдено."
    exit 1
fi

sudo chown $USER_NAME:tty /dev/ptmx

sudo chmod 660 /dev/ptmx

PTS_DIR="/dev/pts"
if [ ! -d "$PTS_DIR" ]; then
    echo "Створюємо каталог $PTS_DIR"
    sudo mkdir -p $PTS_DIR
    sudo mount -t devpts devpts $PTS_DIR
fi

echo "PTY налаштовано для користувача $USER_NAME"
echo "Тепер ви можете запускати Go SSH-сервер без root."
