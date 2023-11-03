# переменный цыетов
YELLOW := \033[1;33m
GREEN := \033[1;32m
RESET := \033[0m

# Компилятор Go
GO := go

# Путь к пакету приложения
PACKAGE := ./cmd

define INFO
	@echo "$(GREEN)[INFO]: $(1)$(RESET)"
endef

define WARN
	@echo "$(YELLOW)[WARN]: $(1)$(RESET)"
endef

.PHONY: get
get:
	@$(call INFO, "Установка зависимостей...")
	@go mod tidy

.PHONY: build
build:
	@$(call INFO, "Сборка проекта...")
	$(GO) build -o myapp $(PACKAGE)

.PHONY: format
format:
	@$(call INFO, "Форматирование кода...")
	@go fmt ./...

.PHONY: vet
vet:
	@$(call INFO, "Проверка кода с помощью go vet...")
	@go vet ./...
	@$(call INFO, "Проверка кода с помощью go vet закончена")
	@echo "-------------------------------------"

.PHONY: test
test:
	@$(call INFO, "Запуск тестов...")
	@$(GO) test ./...

.PHONY: clean
clean:
	@$(call INFO, "Очистка...")
	@rm -f myapp

.PHONY: run
run:
	@$(call INFO, "Запуск приложения...")
	@$(GO) run $(PACKAGE)
	@echo "-------------------------------------"
