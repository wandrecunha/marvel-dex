package main

import (
	"embed"
	"fabricioveronez/app-go/handler"
	"io/fs"
	"net/http"
)

//go:embed static/*
var staticFS embed.FS

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.Index)
	mux.HandleFunc("/find", handler.Find)

	// Configuração para servir arquivos estáticos
	staticFiles, _ := fs.Sub(staticFS, "static")

	// Adiciona o prefixo '/static/' ao caminho
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFiles))))

	http.ListenAndServe(":5000", mux)
}
