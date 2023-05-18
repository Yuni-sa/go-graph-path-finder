package main

type Node struct {
	ID       string
	Children map[*Node]int
}

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
	var graph = map[string]*Node{
		"A": A,
		"B": B,
		"C": C,
		"D": D,
		"E": E,
		"F": F,
		"G": G,
		"H": H,
		"I": I,
		"J": J,
		"K": K,
	}

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
