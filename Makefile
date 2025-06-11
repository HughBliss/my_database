# Makefile

# --- Переменные ---

# Находим все ваши схемы автоматически по названиям поддиректорий в ./schema
# Результат будет: "auth example"
SCHEMAS := $(notdir $(wildcard ./schema/*))

# --- Основные цели для пользователя ---

# .PHONY означает, что эти цели не являются файлами. Это команды.
.PHONY: all generate clean help

# 'make all' или просто 'make' будет запускать генерацию для всех схем.
all: generate

# 'make generate' - основная команда для генерации.
# Она зависит от целей для каждой конкретной схемы.
# Например, если SCHEMAS="auth example", то эта команда будет эквивалентна 'make generate-auth generate-example'
generate: $(addprefix generate-,$(SCHEMAS))

# 'make clean' удаляет все сгенерированные директории.
clean:
	@echo "Cleaning up generated ent files..."
	@rm -rf ./pkg/gen/db*

# --- Правила генерации ---

# Это "шаблонное правило". Оно работает для любой цели, которая начинается с "generate-".
# Например, для 'make generate-auth', переменная $* будет равна "auth".
# Команда запускается, если изменились файлы в ./schema/$* или сам генератор.
generate-%:
	@echo "==> Generating ent schema for '$*'..."
	@go run ./cmd/generate/main.go -s $*

# --- Справка ---

# 'make help' покажет красивые подсказки.
help:
	@echo "Available commands:"
	@echo "  make generate         - Generate ent code for all schemas ($(SCHEMAS))"
	@echo "  make clean            - Remove all generated code"
	@echo "  make help             - Show this help message"