package main

import (
	"container/heap"
	"time"
)

type ShortestPath struct {
	StartLocation string
	EndLocation   string
	Path          []string
	Weight        int
	Duration      time.Duration
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
