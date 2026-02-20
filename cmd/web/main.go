package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	// Remplace par le nom exact de ton module dans go.mod
	"github.com/bearded-nomad/web-starter/internal/templates/pages"
)

func main() {
	mux := http.NewServeMux()

	// 1. On sert le dossier static pour que le navigateur trouve le style.css
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// 2. On branche la page d'accueil sur l'URL racine "/"
	mux.Handle("/", templ.Handler(pages.Home()))

	// 3. On allume le moteur
	fmt.Println("🚀 Serveur démarré sur http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Erreur fatale du serveur :", err)
	}
}
