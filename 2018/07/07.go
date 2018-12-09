package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("December 7th, 2018")

	// 	input := []byte(`Step C must be finished before step A can begin.
	// Step C must be finished before step F can begin.
	// Step A must be finished before step B can begin.
	// Step A must be finished before step D can begin.
	// Step B must be finished before step E can begin.
	// Step D must be finished before step E can begin.
	// Step F must be finished before step E can begin.`)
	input := []byte(`Step Y must be finished before step J can begin.
Step C must be finished before step L can begin.
Step L must be finished before step X can begin.
Step H must be finished before step R can begin.
Step R must be finished before step X can begin.
Step I must be finished before step B can begin.
Step N must be finished before step Q can begin.
Step F must be finished before step X can begin.
Step K must be finished before step G can begin.
Step G must be finished before step P can begin.
Step A must be finished before step S can begin.
Step O must be finished before step D can begin.
Step M must be finished before step W can begin.
Step Q must be finished before step J can begin.
Step X must be finished before step E can begin.
Step U must be finished before step V can begin.
Step Z must be finished before step D can begin.
Step P must be finished before step W can begin.
Step S must be finished before step J can begin.
Step J must be finished before step T can begin.
Step W must be finished before step T can begin.
Step V must be finished before step B can begin.
Step B must be finished before step T can begin.
Step D must be finished before step T can begin.
Step E must be finished before step T can begin.
Step I must be finished before step Z can begin.
Step X must be finished before step D can begin.
Step Q must be finished before step D can begin.
Step S must be finished before step T can begin.
Step R must be finished before step W can begin.
Step O must be finished before step V can begin.
Step C must be finished before step Q can begin.
Step C must be finished before step S can begin.
Step S must be finished before step E can begin.
Step A must be finished before step D can begin.
Step V must be finished before step T can begin.
Step K must be finished before step B can begin.
Step B must be finished before step D can begin.
Step V must be finished before step E can begin.
Step N must be finished before step M can begin.
Step Z must be finished before step T can begin.
Step L must be finished before step A can begin.
Step K must be finished before step Z can begin.
Step F must be finished before step J can begin.
Step M must be finished before step U can begin.
Step Z must be finished before step P can begin.
Step R must be finished before step E can begin.
Step M must be finished before step X can begin.
Step O must be finished before step E can begin.
Step K must be finished before step V can begin.
Step U must be finished before step D can begin.
Step K must be finished before step T can begin.
Step F must be finished before step W can begin.
Step I must be finished before step U can begin.
Step Z must be finished before step S can begin.
Step H must be finished before step D can begin.
Step O must be finished before step P can begin.
Step B must be finished before step E can begin.
Step X must be finished before step U can begin.
Step A must be finished before step J can begin.
Step Y must be finished before step V can begin.
Step U must be finished before step T can begin.
Step G must be finished before step B can begin.
Step U must be finished before step W can begin.
Step H must be finished before step W can begin.
Step G must be finished before step J can begin.
Step X must be finished before step Z can begin.
Step L must be finished before step R can begin.
Step Q must be finished before step X can begin.
Step I must be finished before step O can begin.
Step J must be finished before step E can begin.
Step N must be finished before step D can begin.
Step C must be finished before step B can begin.
Step I must be finished before step W can begin.
Step P must be finished before step J can begin.
Step D must be finished before step E can begin.
Step L must be finished before step J can begin.
Step R must be finished before step J can begin.
Step N must be finished before step A can begin.
Step F must be finished before step O can begin.
Step Y must be finished before step Q can begin.
Step L must be finished before step F can begin.
Step Q must be finished before step U can begin.
Step O must be finished before step T can begin.
Step Z must be finished before step E can begin.
Step Y must be finished before step K can begin.
Step G must be finished before step A can begin.
Step Q must be finished before step E can begin.
Step V must be finished before step D can begin.
Step F must be finished before step K can begin.
Step C must be finished before step E can begin.
Step F must be finished before step A can begin.
Step X must be finished before step B can begin.
Step G must be finished before step U can begin.
Step C must be finished before step H can begin.
Step Y must be finished before step W can begin.
Step R must be finished before step Z can begin.
Step W must be finished before step D can begin.
Step C must be finished before step T can begin.
Step H must be finished before step M can begin.
Step O must be finished before step Q can begin.`)

	lines := strings.Split(string(input), "\n")
	steps := map[string]step{}
	for i := 0; i < len(lines); i++ {
		dep := strings.Split(lines[i], " ")[1]
		stp := strings.Split(lines[i], " ")[7]

		if _, ok := steps[stp]; !ok {
			// s:=fmt.Sprintf("%x", stp)
			t, _ := strconv.ParseInt(fmt.Sprintf("%x", stp), 16, 64)
			t -= 4
			steps[stp] = step{
				name: stp,
				time: int(t),
			}
		}
		if _, ok := steps[dep]; !ok {
			t, _ := strconv.ParseInt(fmt.Sprintf("%x", stp), 16, 64)
			t -= 4
			steps[dep] = step{
				name: dep,
				time: int(t),
			}
		}

		stepstruct := steps[stp]
		stepstruct.deps = append(stepstruct.deps, dep)
		steps[stp] = stepstruct
	}

	todos := steps
	avalaible := []string{}
	executed := []string{}
	e := 0
	keepGoing := true
	for keepGoing {
		keepGoing = false
		for name, step := range todos {
			if len(step.deps) == 0 {
				avalaible = append(avalaible, name)
				delete(todos, name)
			}
		}
		sort.Strings(avalaible)
		if len(avalaible) > 0 {
			executed = append(executed, avalaible[0])
			for name, step := range todos {
				for _, ex := range executed {
					for j := 0; j < len(step.deps); j++ {
						if avalaible[0] == step.deps[j] {
							step.deps = append(step.deps[:j], step.deps[j+1:]...)
							j--
							keepGoing = true
						}
					}
				}
				steps[name] = step
			}
			avalaible = avalaible[1:]
		}
		e++
	}
	res := ""
	for _, name := range executed {
		res += name
	}
	fmt.Printf("Result (part 1): %s\n", res)

	todos = steps
	avalaible = []string{}
	executed = []string{}
	e = 0
	keepGoing = true
	timeline := map[int][5]struct {
		ongoing  string
		timeleft int
	}{}
	s := 0
	done := []string{}
	for keepGoing {
		keepGoing = false
		for name, step := range todos {
			if len(step.deps) == 0 {
				avalaible = append(avalaible, name)
				delete(todos, name)
			}
		}
	}
	res = ""
	for _, name := range executed {
		res += name
	}
	fmt.Printf("Result (part 2): %s\n", res)
}

type step struct {
	name string
	deps []string
	time int
}
