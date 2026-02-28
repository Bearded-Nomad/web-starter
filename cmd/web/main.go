package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/bearded-nomad/web-starter/internal/templates/pages"
)

func main() {
	mux := http.NewServeMux()

	// 1. Fichiers Statiques
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// 2. SEO Dynamique (S'adapte à TOUS les domaines)
	mux.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "User-agent: *\nAllow: /\n\nSitemap: https://%s/sitemap.xml", r.Host)
	})

	mux.HandleFunc("/sitemap.xml", SitemapHandler)

	// 3. Pages
	mux.Handle("/", templ.Handler(pages.Home()))
	mux.Handle("/services", templ.Handler(pages.ServicesSection()))
	// mux.Handle("/contact", templ.Handler(pages.Contact())) // Exemple pour plus tard

	// 4. Lancement
	fmt.Println("🚀 Serveur prêt sur le port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}

func SitemapHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml")

	// On récupère le domaine actuel (ex: garage.maxlesscode.fr ou client.com)
	domain := r.Host

	// On définit les pages du site (tu pourras en ajouter ici plus tard)
	pages := []string{"", "contact", "mentions-legales"}

	// Construction du XML
	fmt.Fprint(w, `<?xml version="1.0" encoding="UTF-8"?>`)
	fmt.Fprint(w, `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`)

	for _, page := range pages {
		url := fmt.Sprintf("https://%s/%s", domain, page)
		fmt.Fprintf(w, `
   <url>
      <loc>%s</loc>
      <changefreq>monthly</changefreq>
      <priority>%s</priority>
   </url>`, url, getPriority(page))
	}

	fmt.Fprint(w, `</urlset>`)
}

// Petite fonction helper pour les priorités SEO
func getPriority(page string) string {
	if page == "" {
		return "1.0"
	}
	return "0.8"
}
