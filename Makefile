run:
	@echo "🏗️ Génération des fichiers Templ..."
	@templ generate
	@echo "🎨 Compilation ultra-rapide du CSS (Tailwind v4)..."
	@npx @tailwindcss/cli -i ./static/css/input.css -o ./static/css/style.css
	@echo "🚀 Démarrage du serveur Go..."
	@go run cmd/web/main.go
dev:
	@air

build:
	@templ generate
	@npx @tailwindcss/cli -i ./static/css/input.css -o ./static/css/style.css
	@go build -o site-bin ./cmd/web/main.go
