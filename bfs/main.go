package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type farm struct {
	ants_number int
	rooms       map[string][]int
	start       map[string][]int
	end         map[string][]int
	links       map[string][]string
}

func main() {
	var myFarm farm
	scenario := false
	myFarm.Read("test.txt")
	bfs := BFS(myFarm, scenario)
	bfsn := BFS(myFarm, !scenario)
	fmt.Println("not scenario :", bfsn)
	ants := Ants(myFarm, BFS(myFarm, scenario), BFS(myFarm, !scenario))
	fmt.Println(scenario, !scenario)

	fmt.Printf("\nall sorted paths from start to end with scenario: %s\n", bfs)
	fmt.Println("Place all Ants on there path: ", ants)
	// MoveAnts(myFarm, ants)
	PrintAnts(myFarm, ants)

	// fmt.Println(Ants(myFarm, BFS(myFarm)))
	// fmt.Println("number of ants is : ", myFarm.ants_number)
	// fmt.Println("rooms are : ", myFarm.rooms)
	// fmt.Println("start is : ", myFarm.start)
	// fmt.Println("end is : ", myFarm.end)
	// fmt.Println("links are : ", myFarm.links)
	// fmt.Println("adjacent is : ", Graph(myFarm))
}

func (myFarm *farm) Read(filename string) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Println("error reading", err)
	}
	content := strings.Split(string(bytes), "\n")

	myFarm.rooms = make(map[string][]int)
	myFarm.start = make(map[string][]int)
	myFarm.end = make(map[string][]int)
	myFarm.links = make(map[string][]string)

	var st, en int
	number, err := strconv.Atoi(content[0])
	if err != nil {
		log.Println("couldn't convert", err)
	}
	myFarm.ants_number = number

	for index := range content {
		if strings.TrimSpace(content[index]) == "##start" {
			st++
			if index+1 <= len(content)-1 {
				split := strings.Split(strings.TrimSpace(content[index+1]), " ")
				x, err := strconv.Atoi(split[1])
				y, err2 := strconv.Atoi(split[2])
				if err == nil && err2 == nil {
					myFarm.start[split[0]] = []int{x, y}
				}

			}

		} else if strings.TrimSpace(content[index]) == "##end" {
			en++
			if index+1 <= len(content)-1 {
				split := strings.Split(strings.TrimSpace(content[index+1]), " ")
				x, err := strconv.Atoi(split[1])
				y, err2 := strconv.Atoi(split[2])
				if err == nil && err2 == nil {
					myFarm.end[split[0]] = []int{x, y}
				}

			}
		} else if strings.Contains(content[index], "-") {
			split := strings.Split(strings.TrimSpace(content[index]), "-")
			if len(split) == 2 {
				myFarm.links[split[0]] = append(myFarm.links[split[0]], split[1])
			}
		} else if strings.Count(content[index], " ") == 2 {
			split := strings.Split(strings.TrimSpace(content[index]), " ")
			if len(split) == 3 {
				x, err := strconv.Atoi(split[1])
				y, err2 := strconv.Atoi(split[2])
				if err == nil || err2 == nil {
					myFarm.rooms[split[0]] = []int{x, y}
				}
			}
		} else if (strings.HasPrefix(strings.TrimSpace(content[index]), "#") || strings.HasPrefix(strings.TrimSpace(content[index]), "L")) && (strings.TrimSpace(content[index]) != "##start" && strings.TrimSpace(content[index]) != "##end") {
			continue
		}
	}
	if en != 1 || st != 1 {
		log.Println("rooms setup is incorrect", err)
	}
}

func Graph(farm farm, scenario bool) map[string][]string {
	adjacent := make(map[string][]string)
	var Start string
	for key := range farm.start {
		Start = key
	}

	for room := range farm.rooms {
		adjacent[room] = []string{}
	}
	for room, links := range farm.links {
		for _, link := range links {
			if scenario {
				if room == Start {
					adjacent[room] = append([]string{link}, adjacent[room]...)
					adjacent[link] = append([]string{room}, adjacent[link]...)
				} else {

					adjacent[room] = append(adjacent[room], link)
					adjacent[link] = append(adjacent[link], room)
				}
			} else {

				adjacent[room] = append(adjacent[room], link)
				adjacent[link] = append(adjacent[link], room)
			}
		}
	}

	return adjacent
}

func BFS(myFarm farm, scenario bool) [][]string {
	adjacent := Graph(myFarm, scenario)
	var Queue []string
	var endd string
	start := myFarm.start
	end := myFarm.end
	var Sorted [][]string

	for key := range start {
		Visited := make(map[string]bool)
		Visited[key] = true
		for key := range end {
			endd = key
		}
		i := 0
		for _, adj := range adjacent[key] {
			var ar []string
			if !Visited[adj] {

				Queue = append(Queue, adj)
				Visited[adj] = true
			}

			var current string
			Parents := make(map[string]string)

			finEnd := false
			for len(Queue) > 0 {

				current = Queue[0]

				ar = append(ar, current)
				for _, v := range Queue {
					if v == endd {
						current = v
						finEnd = true
						Queue = []string{}
						break
					}
				}

				fmt.Println("\nCurrent = ", current)
				Visited[current] = true
				if !finEnd {
					if len(Queue) == 1 {
						Queue = []string{}
					} else {
						Queue = Queue[1:]
					}
				}

				fmt.Println("Queue after removing current = ", Queue)
				// Visited[current] = true
				if current == endd {
					Visited[current] = true
					// for _, v := range Queue {
					// 	if !isUsed[v] {
					// 		Visited[v] = false
					// 	}
					// }
					fmt.Println("Queue at current=endd is: ", Queue)
					Queue = []string{}
					// AllPaths = append(AllPaths, Queue)

					break
				}

				for _, link := range adjacent[current] {
					// Visited[current] = true
					// if link == endd {
					// 	Queue = append(Queue, link)
					// 	// Visited[link] = true
					// 	Parents[link] = current
					// 	break
					// }

					if !Visited[link] {

						Queue = append(Queue, link)
						fmt.Println("Queue after adding link = ", Queue)
						// Visited[link] = true
						Parents[link] = current
					}
				}
			}

			if !Visited[endd] {
				fmt.Print("\n No path found to end room \n")
				continue
			}
			Visited[endd] = false
			path := []string{endd}
			current = endd

			for Parents[current] != "" {
				current = Parents[current]
				// isUsed[current] = true
				path = append([]string{current}, path...)
			}
			fmt.Println("array of current: ", ar)
			c := 0
			for i, v := range ar {
				for _, e := range path {
					if v != e {
						c++
						fmt.Printf("\nc N° %d is : %d\n", i, c)
					}
				}
				if c == len(path) {
					fmt.Printf("\nif c = 0 N° %d is : %d\n", i, c)

					Visited[v] = false
				}
				c = 0
			}

			path = append([]string{key}, path...)
			fmt.Printf("\nPath N° %v is : %v\n", i, path)
			i++
			Sorted = append(Sorted, path)
		}
		break
	}

	// fmt.Println("\nAllPaths: ", AllPaths)
	Sorted = SortPath(Sorted)
	// fmt.Printf("\nall sorted paths from start to end: %v\n", Sorted)
	return Sorted
}

func SortPath(Paths [][]string) [][]string {
	if len(Paths) <= 1 {
		return Paths
	}
	pivot := Paths[len(Paths)-1]
	var less, greater [][]string
	for _, v := range Paths[:len(Paths)-1] {
		if len(v) <= len(pivot) {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}
	return append(append(SortPath(less), pivot), SortPath(greater)...)
}

func Ants(myFarm farm, path1, path2 [][]string) [][]string {
	ants := myFarm.ants_number
	paths := [][]string{}
	fmt.Println("num of ants is :", ants)

	// for i := 0; i < len(paths); i++ {
	i := 0
	j := 1
	for j <= ants {
		for k := 0; k < len(path1); k++ {
			if i < len(path1) {
				if len(path1[i]) > len(path1[k]) {
					i = k
				}
			} else {
				i = 0
			}
			if k == len(path1)-1 {
				path1[i] = append(path1[i], "L"+strconv.Itoa(j))
				i = 0
			}
		}
		j++
	}

	fmt.Println("Path1 = ", path1)

	m := 0
	n := 1
	for n <= ants {
		for k := 0; k < len(path2); k++ {
			if m < len(path2) {
				if len(path2[m]) > len(path2[k]) {
					m = k
				}
			} else {
				m = 0
			}
			if k == len(path2)-1 {
				path2[m] = append(path2[m], "L"+strconv.Itoa(n))
				m = 0
			}
		}
		n++
	}

	fmt.Println("Path2 = ", path2)

	if len(path1[len(SortPath(path1))-1]) <= len(path2[len(SortPath(path2))-1]) {
		for _, v := range path1 {
			paths = append(paths, v)
		}
	} else {
		for _, v := range path2 {
			paths = append(paths, v)
		}
	}

	// }

	// k := 0
	// for i := ants; i > 0; i-- {
	// 	for j := 1; j < len(paths); j++ {
	// 		if k < len(paths) {
	// 			if len(paths[k]) <= len(paths[j]) {
	// 				paths[k] = append(paths[k], "L"+strconv.Itoa(i))
	// 				k++
	// 				break
	// 			} else {
	// 				k = 0
	// 			}
	// 		} else {
	// 			k = 0
	// 		}
	// 	}
	// }

	return paths
}

func MoveAnts(myFarm farm, paths [][]string, scenario bool) {
	for i := 0; i < len(paths); i++ {
		k := len(paths[i]) - 1
		for j := 1; j < len(paths[i]); j++ {

			// if paths[i][j] == "end" {
			// 	fmt.Print(paths[i][k] + "-" + paths[i][j] + " ")
			// 	break
			// }

			fmt.Print(paths[i][k] + "-" + paths[i][j] + " ")
			// if i == len(paths)-1 {
			// 	k--
			// }
			break
		}

	}

	var a, b []string

	all := [][][]string{}
	g := Ants(myFarm, BFS(myFarm, false), BFS(myFarm, true))
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			if strings.HasPrefix(g[i][j], "L") {
				b = append(b, g[i][j])
			} else if j != 0 {
				a = append(a, g[i][j])
			}
		}
		all = append(all, [][]string{a, b})

		a = []string{}
		b = []string{}
	}
	fmt.Print("all paths separed: ", all)

	var RoomsArray [][][]string
	var ArrayElem [][]string
	var Elem []string

	for i := 0; i < len(all); i++ {
		for j := len(all[i][1]) - 1; j >= 0; j-- {
			for k := 0; k < len(all[i][0]); k++ {
				Elem = append(Elem, all[i][1][j]+"-"+all[i][0][k])
				for l := j; l < j; l++ {
					Elem = append([]string{"zz"}, Elem...)
				}
			}
			ArrayElem = append(ArrayElem, Elem)
			Elem = []string{}

		}
		RoomsArray = append(RoomsArray, ArrayElem)
		ArrayElem = [][]string{}
	}

	fmt.Print("\nants in Rooms: ", RoomsArray)

	// for i := 0; i < len(all); i++ {
	// 	k := 0
	// 	for j := len(all[i])-1; j >=0; j--{
	// 		fmt.Print(all[i][j][len(all[i][j])-1-k]+"-"+)

	// 	}

	// }
	AfterPrint := [][]string{}

	for i := 0; i < len(RoomsArray); i++ {
		for j := 0; j < len(RoomsArray[i]); j++ {
			AfterPrint = append(AfterPrint, RoomsArray[i][j])
		}
	}
	fmt.Println("\n After Print: ", AfterPrint)

	for _, v := range AfterPrint {
		for i := 0; i < len(v); i++ {
			fmt.Print(v[i])
		}
	}
}

func PrintAnts(myfarm farm, paths [][]string) {
	var p, a, text []string
	var lines [][]string
	var all [][]string
	fmt.Println("\n\npaths are :", paths)
	for i := 0; i < len(paths); i++ {
		for j := 1; j < len(paths[i]); j++ {
			if strings.HasPrefix(paths[i][j], "L") {
				a = append(a, paths[i][j])
			} else {
				p = append(p, paths[i][j])
			}
		}
		for l := 0; l < len(a); l++ {
			for x := 0; x < len(p); x++ {
				text = append(text, a[l]+"-"+p[x]+" ")
			}
			lines = append(lines, text)
			text = []string{}
		}
		fmt.Println("\nlines are : ", lines)
		var print [][]string
		for i := range lines {
			space := []string{}
			if i != 0 {
				for n := 0; n < i; n++ {
					space = append(space, "")
				}
				lines[i] = append(space, lines[i]...)
				print = append(print, lines[i])
			} else {
				print = append(print, lines[i])
			}
			// fmt.Println("l0", i, lines[i][0])
			// fmt.Println("l1", i, lines[i][1])
			// fmt.Println("l2", i, lines[i][2])
		}

		all = append(all, print...)
		a = []string{}
		p = []string{}
		lines = [][]string{}

	}
	fmt.Println("all is : ", all)

	for len(all) > 0 {
		for i := 0; i < len(all); i++ {
			fmt.Print(all[i][0])
			// fmt.Println("all elem", all[i])
			// fmt.Println("all minus elem", all[i][1:])

			all[i] = all[i][1:]

			if len(all[i]) == 0 {
				all = append(all[:i], all[i+1:]...)
				i--
			}
		}
		fmt.Println()
	}
}
