package main

import (
	"fmt"
	"net/http"
	"os"

	"example.com/tui-import-repro/assets"
	"example.com/tui-import-repro/ui/pages"
	"github.com/a-h/templ"
	"github.com/joho/godotenv"
)

func main() {
	initDotEnv()

	mux := http.NewServeMux()
	setupAssetsRoutes(mux)
	mux.Handle("GET /", templ.Handler(pages.Landing()))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8091"
	}

	addr := ":" + port
	fmt.Printf("Server is running on http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		panic(err)
	}
}

func initDotEnv() {
	_ = godotenv.Load()
}

func setupAssetsRoutes(mux *http.ServeMux) {
	isDevelopment := os.Getenv("GO_ENV") != "production"

	assetHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isDevelopment {
			w.Header().Set("Cache-Control", "no-store")
		}

		var fs http.Handler
		if isDevelopment {
			fs = http.FileServer(http.Dir("./assets"))
		} else {
			fs = http.FileServer(http.FS(assets.Assets))
		}

		fs.ServeHTTP(w, r)
	})

	mux.Handle("GET /assets/", http.StripPrefix("/assets/", assetHandler))
}
