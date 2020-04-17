package main

import (
	"html/template"
	"log"
	"net/http"
)

type Handler struct {
	Tpl *template.Template
}

func main() {

	log.Println("Starting service...")

	h := Handler{
		Tpl: template.Must(template.ParseGlob("tpl/*")),
	}

	http.HandleFunc("/", h.Index)
	http.HandleFunc("/news/", h.News)

	log.Println("Listening port :8000...")
	http.ListenAndServe(":8000", nil)

}

func (h *Handler) News(w http.ResponseWriter, r *http.Request) {
	type d struct {
		Title, H1 string
		List      []string
	}

	data := d{
		Title: `Сайт компании "Рога и копыта" | Новости`,
		H1:    `Новости`,
		List: []string{
			`Первая новость`,
			`Потом вторая`,
			`И еще одна`,
		},
	}

	err := h.Tpl.ExecuteTemplate(w, "news.tpl", data)
	if err != nil {
		log.Fatal()
	}
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	type d struct {
		Title, H1, Text string
	}

	data := d{
		Title: `Сайт компании "Рога и копыта" | Главная страница`,
		H1:    `Главная страница`,
		Text:  `Немного текста о компании`,
	}

	err := h.Tpl.ExecuteTemplate(w, "index.tpl", data)
	if err != nil {
		log.Fatal()
	}
}
