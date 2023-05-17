package main

import (
	"container/heap"
	"html/template"
	"log"
	"net/http"
	"time"
)

var (
	A = &Node{ID: "A"}
	B = &Node{ID: "B"}
	C = &Node{ID: "C"}
	D = &Node{ID: "D"}
	E = &Node{ID: "E"}
	F = &Node{ID: "F"}
	G = &Node{ID: "G"}
	H = &Node{ID: "H"}
	I = &Node{ID: "I"}
	J = &Node{ID: "J"}
	K = &Node{ID: "K"}
)

func initializeGraph() map[string]*Node {
	graph := make(map[string]*Node)
	graph["A"] = A
	graph["B"] = B
	graph["C"] = C
	graph["D"] = D
	graph["E"] = E
	graph["F"] = F
	graph["G"] = G
	graph["H"] = H
	graph["I"] = I
	graph["J"] = J
	graph["K"] = K

	A.Children = map[*Node]int{
		B: 2,
		C: 4,
		D: 6,
	}
	B.Children = map[*Node]int{
		A: 2,
		C: 3,
		E: 5,
	}
	C.Children = map[*Node]int{
		A: 4,
		B: 3,
		D: 1,
		E: 6,
		F: 8,
	}
	D.Children = map[*Node]int{
		A: 6,
		C: 1,
		F: 7,
	}
	E.Children = map[*Node]int{
		B: 5,
		C: 6,
		G: 4,
		H: 9,
	}
	F.Children = map[*Node]int{
		C: 8,
		D: 7,
		G: 3,
		H: 2,
		I: 5,
	}
	G.Children = map[*Node]int{
		E: 4,
		F: 3,
		I: 1,
	}
	H.Children = map[*Node]int{
		E: 9,
		F: 2,
		I: 6,
	}
	I.Children = map[*Node]int{
		F: 5,
		G: 1,
		H: 6,
		J: 2,
		K: 4,
	}
	J.Children = map[*Node]int{
		I: 2,
		K: 3,
	}
	K.Children = map[*Node]int{
		I: 4,
		J: 3,
	}

	return graph
}

var graph = initializeGraph()

type ShortestPath struct {
	StartLocation string
	EndLocation   string
	Path          []string
	Weight        int
	Duration      time.Duration
}

type Node struct {
	ID       string
	Children map[*Node]int
}

type Item struct {
	node     *Node
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*Item)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func Dijkstra(graph map[string]*Node, start, end string) ([]string, int) {
	distances := make(map[string]int)
	previous := make(map[string]*Node)
	pq := make(PriorityQueue, 0)

	for id := range graph {
		distances[id] = 1e9
	}
	distances[start] = 0
	heap.Push(&pq, &Item{node: graph[start], priority: distances[start]})

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Item).node

		if current.ID == end {
			break
		}

		for neighbor, weight := range current.Children {
			distance := distances[current.ID] + weight
			if distance < distances[neighbor.ID] {
				distances[neighbor.ID] = distance
				previous[neighbor.ID] = current
				heap.Push(&pq, &Item{node: neighbor, priority: distance})
			}
		}
	}

	path := make([]string, 0)
	current := graph[end]
	for current != nil {
		path = append([]string{current.ID}, path...)
		current = previous[current.ID]
	}

	return path, distances[end]
}

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

func main() {
	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", shortestPathHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
