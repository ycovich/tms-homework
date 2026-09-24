#!/bin/bash
# Запуск Ubuntu для ДЗ (macOS / Linux).
#   ./start.sh ivan           — первый запуск; ivan — твоё имя латиницей
#   ./start.sh                — вернуться в контейнер, все изменения на месте
#   ./start.sh --reset ivan   — удалить контейнер и начать с нуля
set -e
NAME=linux-quest
IMAGE=ghcr.io/ycovich/linux-quest

if [ "$1" = "--reset" ]; then
  docker rm -f "$NAME" >/dev/null 2>&1 || true
  shift
fi

STUDENT=$(echo "$1" | tr '[:upper:]' '[:lower:]')

# Контейнер уже есть: возвращаемся в него, но только если он создан под тем же именем
if EXISTING=$(docker container inspect -f '{{.Config.Labels.quest_student}}' "$NAME" 2>/dev/null); then
  [ "$EXISTING" != "<no value>" ] || EXISTING=""
  if [ -n "$STUDENT" ] && [ -n "$EXISTING" ] && [ "$STUDENT" != "$EXISTING" ]; then
    echo "Контейнер уже создан для пользователя $EXISTING."
    echo "Вернуться в него: ./start.sh"
    echo "Начать с нуля под именем $STUDENT: ./start.sh --reset $STUDENT"
    exit 1
  fi
  exec docker start -ai "$NAME"
fi

if [[ ! $STUDENT =~ ^[a-z][a-z0-9_-]{0,30}$ ]]; then
  echo "Укажи своё имя латиницей, например:  ./start.sh ivan"
  exit 1
fi

# Образ скачивается один раз, потом только обновляется.
# Если обновить не вышло (например, нет интернета), работаем со старым образом
if docker image inspect "$IMAGE" >/dev/null 2>&1; then
  docker pull -q "$IMAGE" >/dev/null 2>&1 || true
else
  echo "Скачиваю образ, это займёт пару минут..."
  docker pull "$IMAGE"
fi
exec docker run -it --name "$NAME" -h tms-lab --label quest_student="$STUDENT" -e STUDENT="$STUDENT" "$IMAGE"
