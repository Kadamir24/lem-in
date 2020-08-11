package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Ants struct {
	path   []string
	ID     int
	Nroom  int
	finish bool
}

type Room struct {
	busy bool
}

type St struct {
	start *Vertex
	end   *Vertex
}
type Vertex struct {
	visited    bool
	value      string
	neighbours []*Vertex
	parent     *Vertex
	// ID         int
}

type Combination [][]string

var AllComb []Combination

func NewVertex(value string) *Vertex {
	return &Vertex{
		value: value,
	}
}

func (v *Vertex) connect(vertex *Vertex) {
	v.neighbours = append(v.neighbours, vertex)
	vertex.neighbours = append(vertex.neighbours, v)
}

type Graph struct{}

var st St

var array []string

var array2 [][]string

// func (g *Graph) dfs(vertex, end *Vertex) {

// 	if vertex.visited {

// 		return
// 	}
// 	array = append(array, *vertex)
// 	vertex.visited = true
// 	fmt.Println(vertex)
// 	if vertex == st.end {
// 		vertex.visited = false
// 		return
// 	}
// 	for _, v := range vertex.neighbours {
// 		g.dfs(v)
// 	}
// }

func (g *Graph) dfs(vertex, end *Vertex) [][]string {
	// fmt.Println(vertex)
	vertex.visited = true

	array = append(array, vertex.value)

	if vertex == end {

		// fmt.Println(array)
		copy := make([]string, len(array))
		for i := range array {
			copy[i] = array[i]
		}

		array2 = append(array2, copy)

		// fmt.Println(array2)
		// fmt.Println(array2)
		// return array2
	} else {
		for _, v := range vertex.neighbours {
			if !v.visited {
				g.dfs(v, end)
			}
		}

	}

	array = array[:len(array)-1]
	vertex.visited = false

	// fmt.Println("chet:", array)
	//array = array[1:]

	return array2

}

func bfs(vertex, end *Vertex) ([]*Vertex, bool) {

	q := []*Vertex{}
	// q.New()
	p := []*Vertex{}
	visited := make(map[*Vertex]bool)

	q = append(q, vertex)

	for {
		if len(q) == 0 {
			return q, false

		}
		node := q[0]
		q = q[1:]
		if node == end {
			return getPath(p, st.start, st.end), true

		}
		visited[node] = true
		near := node.neighbours

		for i := 0; i < len(near); i++ {
			j := near[i]

			if !visited[j] {
				j.parent = node
				p = append(p, j)
				q = append(q, j)

				visited[j] = true

			}
		}
	}
}

func getPath(parent []*Vertex, start, end *Vertex) []*Vertex {
	path := []*Vertex{}

	for i := len(parent) - 1; i >= 0; i-- {
		if parent[i] != start {
			path = append(path, parent[i].parent)

		}
	}
	return path
}

func Min(paths [][]string) []string {
	min := paths[0]
	var temp []string
	for i := 0; i < len(paths)-1; i++ {
		for j := i + 1; j < len(paths); j++ {
			if len(paths[i]) < len(paths[j]) {
				temp = paths[i]
				if len(min) > len(temp) {
					min = temp
				}
			}
		}

	}
	return min
}

func IsCross(stack [][]string, two []string, start, end string) bool {

	for _, lol2 := range stack {
		for _, lol := range lol2 {
			for _, kek := range two {
				if lol == kek && lol != start && lol != end {
					return true
				}
			}

		}

	}

	return false

}

func IsUnique(target []string, all [][]string, start, end string) bool {

	for _, unit1 := range all {
		for _, unit2 := range unit1 {
			for _, lol := range target {
				if unit2 == lol && lol != start && lol != end {
					return false
				}
			}
		}
	}
	return true

}

func CrossCheck(allpaths [][]string, start, end string) []Combination {

	var overstack []Combination
	var stack Combination

	// x := Min(allpaths)

	// stack = append(stack, x)

	if len(allpaths) == 1 {
		for _, i := range allpaths {
			stack = append(stack, i)
			overstack = append(overstack, stack)
			return overstack
		}

	}

	for i := 0; i < len(allpaths); i++ {
		stack = append(stack, allpaths[i])
		for j := 1; j < len(allpaths); j++ {

			if !IsCross(stack, allpaths[j], start, end) {

				// fmt.Println("START")
				// if len(allpaths[i]) <= len(allpaths[j]) {

				// } else {
				// 	allpaths[i] = allpaths[j]

				// }
				if IsUnique(allpaths[j], stack, start, end) {

					// fmt.Println("LOL")
					stack = append(stack, allpaths[j])
					// allpaths[j] = []string{}

					// allpaths[i] = []string{}
				}

			}

		}
		overstack = append(overstack, stack)
		stack = Combination{}
	}

	return overstack

}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("ERROR: invalid data format")
		return
	}
	arg := args[0]
	var nums string
	var ants int
	file, err := ioutil.ReadFile(arg)
	if err != nil {
		fmt.Println("ERROR: invalid data format")
		return
	}
	str := strings.Split(string(file), "\n")
	var g []Vertex
	var start, end string
	for i, v := range str {

		// if len(v) == 1 {
		// 	nums = v
		// 	ants, err = strconv.Atoi(nums)
		// 	if err != nil {
		// 		fmt.Println("Salamaleikum")
		// 	}
		// }

		nums = str[0]

		ants, err = strconv.Atoi(nums)
		if err != nil {
			fmt.Println("ERROR: invalid data format")
			return
		}

		str2 := strings.Split(v, "-")
		// str3 := strings.Split(v, " ")

		if len(str2) == 2 {
			b := false
			q := false
			for _, k := range g {

				if k.value == string(str2[0]) {

					b = true
				}
				if k.value == string(str2[1]) {
					q = true
				}
			}
			if !b {
				g = append(g, *NewVertex(string(str2[0])))

			}
			if !q {
				g = append(g, *NewVertex(string(str2[1])))

			}

		}
		//  else {
		// 	if strings.HasPrefix(v, "##start") {
		// 		start = string(str[i+1][0])
		// 	} else if strings.HasPrefix(v, "##end") {
		// 		end = string(str[i+1][0])
		// 	}
		// }

		if strings.HasPrefix(v, "##start") {
			s1 := strings.Split(str[i+1], " ")
			start = string(s1[0])

		} else if strings.HasPrefix(v, "##end") {
			s2 := strings.Split(str[i+1], " ")
			end = string(s2[0])
		}

	}

	for _, v := range str {
		str2 := strings.Split(v, "-")
		if len(str2) == 2 {
			for i, k := range g {
				if k.value == string(str2[0]) {

					for j, l := range g {
						if l.value == string(str2[1]) {

							g[i].connect(&g[j])
						}
					}

				}
			}

		}
	}

	for i := range g {

		if g[i].value == start {
			st.start = &g[i]
			for _, v := range g[i].neighbours {
				if v.value == end {
					printStartEnd(ants, end, file)
					return
				}
			}
		} else if g[i].value == end {
			st.end = &g[i]

		}

	}

	gr := Graph{}
	allpaths := gr.dfs(st.start, st.end)
	sortedPath(allpaths)
	combs := (CrossCheck(allpaths, start, end))

	opt := OptimalPath(combs, ants)
	sortedPath(opt)

	var res []Res
	queue := make([]int, len(opt))
	f := 1
	i := 0
	for f <= ants {

		tmp := opt[i]
		// count := 0

		r := Res{
			ant:      f,
			antsPath: tmp,
			ind:      1,
		}

		res = append(res, r)
		queue[i]++
		// for _, v := range res {

		// 	if isSame(opt[i], v.antsPath) {
		// 		count++

		// 	}
		// }

		sum := len(tmp) + queue[i]
		if len(opt) > i+1 {
			// count2 := 0
			// for _, v := range res {

			// 	if isSame(opt[i+1], v.antsPath) {
			// 		count2++

			// 	}
			// }
			if sum > len(opt[i+1])+queue[i+1] {

				i++
			}

		} else {
			i = 0
		}
		f++

	}

	d := 1
	done := 0
	p := 0

	fmt.Println(string(file), "\n")
	for done != 1 {

		for p < ants {
			done = move_ant(p, res)
			p++
		}

		fmt.Println()

		d++
		p = 0
	}
}

func sortedPath(all [][]string) {
	for i := 0; i < len(all); i++ {
		for j := i; j < len(all); j++ {
			if len(all[i]) > len(all[j]) {
				all[i], all[j] = all[j], all[i]
			}
		}
	}
}

func printStartEnd(amount int, end string, file []byte) {
	fmt.Println(string(file), "\n")

	for i := 1; i <= amount; i++ {
		fmt.Printf("L%d-%s ", i, end)
	}
	fmt.Println()
}

func move_ant(c int, res []Res) int {
	done := make(map[*Res]bool)
	for j, v := range res[c].antsPath {
		if !done[&res[c]] && !isBusy(v, res) && j == res[c].ind {
			fmt.Print("L", res[c].ant, "-", v, " ")
			if j != len(res[c].antsPath)-1 {
				res[c].busy = v
			} else {
				res[c].busy = ""
			}

			res[c].ind++
			done[&res[c]] = true

			if c == len(res)-1 && j == len(res[c].antsPath)-1 {
				return 1
			}

		}

	}
	return 0

}

func isBusy(v string, res []Res) bool {

	for _, k := range res {
		if k.busy == v {
			return true
		}
	}
	return false
}

func isSame(i, k []string) bool {
	max := k
	if len(i) > len(k) {
		max = i
	}
	for j := 0; j < len(max); j++ {
		if i[j] != k[j] {
			return false
		}
	}
	return true
}

type Res struct {
	ant      int
	antsPath []string
	busy     string
	ind      int
}

//priehali

func OptimalPath(combs []Combination, ants int) Combination {
	var answer Combination
	var indexoptimal []int
	nants := 0
	//counter := 0
	for _, paths := range combs {
		nants = ants
		max := Max(paths)
		x := len(max)
		// for x != len(max) {
		// 	x++
		// 	ants--
		// }
		columns := len(paths)
		area := x * columns

		area2 := 0

		for _, lol := range paths {
			area2 += len(lol)

		}

		for area2 != area && nants > 0 {
			area2++
			nants--
		}

		temp := nants / columns
		tempost := nants % columns
		x = x + temp

		if tempost > 0 {
			if tempost < area2 {
				x = x + 1
				// } else {
				// 	tempost
				// }
			}
			// fmt.Println("Posle plusa:", x)

			// indexoptimal = append(indexoptimal, x)

		}

		indexoptimal = append(indexoptimal, x)

	}

	indmin := MinInd(indexoptimal)

	for i, res := range combs {
		if i == indmin {
			answer = res
		}
	}
	return answer

}

func Max(paths [][]string) []string {
	max := paths[0]
	var temp []string
	for i := 0; i < len(paths)-1; i++ {
		for j := i + 1; j < len(paths); j++ {
			if len(paths[i]) < len(paths[j]) {
				temp = paths[j]
				if len(max) < len(temp) {
					max = temp
				}
			}
		}

	}
	return max
}

func MinInd(arr []int) int {
	index := 0
	counter := 0
	// min := arr[0]
	// var temp int

	for i := 0; i < len(arr)-1; i++ {
		// counter = 0
		for j := i + 1; j < len(arr); j++ {
			if arr[i] <= arr[j] {
				// fmt.Println(arr[i], arr[j])
				// min = arr[j]
				index = i
				// fmt.Println(index)
			} else {
				// index = j
				// counter = 0
				break
				// fmt.Println("j", index)
			}
			counter++
			if counter == len(arr)-1 {
				// fmt.Println("RETURNIM:", index, counter)
				return index
			}
		}

	}

	return index

}
