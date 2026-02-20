# Nom du binaire
BINARY_NAME=tmp/main

.PHONY: all build run clean dev

# Commande par défaut utilisée par Air
build:
	@echo "🏗️  Génération Templ..."
	@templ generate
	@echo "🎨 Compilation Tailwind v4..."
	@./node_modules/.bin/tailwindcss -i ./static/css/input.css -o ./static/css/style.css
	@echo "🐹 Build Go..."
	@go build -o $(BINARY_NAME) ./cmd/web/main.go

# Lancer Air pour le live-reload
dev:
	@air

# Nettoyage
clean:
	@rm -rf tmp
	@rm -f static/css/style.css