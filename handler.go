package main

import (
	"html/template"
	"net/http"
	"time"
)

func shortestPathHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderTemplate(w, "index.html", nil)
		return
	}

	start := r.FormValue("start")
	end := r.FormValue("end")

	startTime := time.Now()
	path, weight := Dijkstra(graph, start, end)
	endTime := time.Now()
	duration := endTime.Sub(startTime)

	sp := ShortestPath{
		StartLocation: start,
		EndLocation:   end,
		Path:          path,
		Weight:        weight,
		Duration:      duration,
	}

	renderTemplate(w, "index.html", sp)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmpl = "templates/" + tmpl
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
