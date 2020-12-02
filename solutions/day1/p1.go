package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println(Solve())
	fmt.Println(Solve2())
}

func Solve() string {
	file, err := os.Open("input/1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	m := make(map[int]bool, 0)

	for scanner.Scan() {
		line := scanner.Text()
		v, _ := strconv.ParseInt(line, 10, 32)
		t := int(2020 - v)
		_, found := m[t]
		if found {
			answer := t * int(v)
			return strconv.FormatInt(int64(answer), 10)
		} else {
			m[int(v)] = true
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	panic("not found")
}

type Pair struct {
	a, b int
}

func Solve2() string {
	file, err := os.Open("input/1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ns := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		v, _ := strconv.ParseInt(line, 10, 32)
		ns = append(ns, int(v))
	}

	m := make(map[int]int, 0)

	for i := 0; i < len(ns)-1; i++ {
		a := ns[i]
		for j := i + 1; j < len(ns); j++ {
			b := ns[j]
			t := 2020 - a - b
			m[t] = a * b
		}
	}

	for i:=0;i<len(ns);i++ {
		t := ns[i]
		f, found := m[t]
		if found {
			answer := t * f
			return strconv.FormatInt(int64(answer), 10)
		}
	}

	panic("not found")
}
