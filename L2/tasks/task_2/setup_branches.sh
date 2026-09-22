#!/usr/bin/env bash
# Создаёт от текущей ветки две новые ветки — vowels и consonants — и в каждой
# по-своему переписывает main.go. Возвращает репозиторий на исходную ветку.
set -euo pipefail

if ! git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
  echo "Ошибка: текущая папка не является git-репозиторием." >&2
  exit 1
fi

if [ -n "$(git status --porcelain)" ]; then
  echo "Ошибка: есть незакоммиченные изменения. Закоммить или отмени их перед запуском скрипта." >&2
  exit 1
fi

if [ ! -f main.go ]; then
  echo "Ошибка: main.go не найден в текущей папке." >&2
  exit 1
fi

BASE_BRANCH=$(git symbolic-ref --short HEAD)
echo "Базовая ветка: $BASE_BRANCH"

# --- vowels ---
git checkout -b vowels "$BASE_BRANCH"
cat > main.go <<'EOF'
package main

import "fmt"

func main() {
	fmt.Println("a group of people appointed to investigate, report or decide to perform action over some matter is called") // COMMITTEE

	printAnswer()
	fmt.Println()
}

func printAnswer() {
	fmt.Print("O")
	fmt.Print("I")
	fmt.Print("E")
	fmt.Print("E")
}
EOF
git add main.go
git commit -m "Add vowels" --quiet
git checkout "$BASE_BRANCH" --quiet

# --- consonants ---
git checkout -b consonants "$BASE_BRANCH"
cat > main.go <<'EOF'
package main

import "fmt"

func main() {
	fmt.Println("a group of people appointed to investigate, report or decide to perform action over some matter is called") // COMMITTEE

	printAnswer()
	fmt.Println()
}

func printAnswer() {
	fmt.Print("C")
	fmt.Print("M")
	fmt.Print("M")
	fmt.Print("T")
	fmt.Print("T")
}
EOF
git add main.go
git commit -m "Add consonants" --quiet
git checkout "$BASE_BRANCH" --quiet

echo ""
echo "Готово. Созданы ветки: vowels, consonants (от $BASE_BRANCH)."
echo "Слей их в $BASE_BRANCH так, чтобы вывод программы был корректным."
