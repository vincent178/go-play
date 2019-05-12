package main

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures": {"discrete math"},
	"databases":       {"data structures"},
	"discrete math":   {"intro to programming"},
	"formal language": {"discrete math"},
	"networks":        {"operating systems"},
	"operating systems": {
		"data structures",
		"computer organization",
	},
	"programming languages": {
		"data structures",
		"computer organization",
	},
}

func main() {
}

func topoSort(map[string][]string) []string {

	seen := map[string]bool

	for key, value := range prereqs {
	}
}
