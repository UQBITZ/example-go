// example-go: servidor HTTP mínimo, só biblioteca padrão, publicado na UQBITZ.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

// Marcador que o ensaio confere no corpo da resposta; muda a cada versão.
const marker = "uqbitz-example-go-v2"

var page = template.Must(template.New("page").Parse(`<!doctype html>
<html lang="pt-BR">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>example-go</title>
</head>
<body>
<main>
<h1>Go na UQBITZ</h1>
<p>Servidor <code>net/http</code> compilado no build e servido numa microVM.</p>
<ul>
<li>Versão do Go: {{.GoVersion}}</li>
<li>Hora do servidor: {{.Now}}</li>
</ul>
<p id="marker">{{.Marker}}</p>
</main>
</body>
</html>
`))

type pageData struct {
	GoVersion string
	Now       string
	Marker    string
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		data := pageData{
			GoVersion: runtime.Version(),
			Now:       time.Now().UTC().Format(time.RFC3339),
			Marker:    marker,
		}
		if err := page.Execute(w, data); err != nil {
			log.Printf("render: %v", err)
		}
	})
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	addr := "0.0.0.0:" + port
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
