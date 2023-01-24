package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	f, err := os.Create("out/draft.txt")
	ck(err)
	defer f.Close()
	w := bufio.NewWriter(f)

	// Essay begins here.

	RandLineWriter(w, "data/mainresult.txt")
	InputLineWriter(w, "State the main results")
	RandLineWriter(w, "data/novelty.txt")
	InputLineWriter(w, `State new, interesting, unique 
or surprising ideas. For example
they could be new diffculties resolved here
that are not present in previous work`)

	w.Flush()
}

func ck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// RandLineWriter furnishes a random line given a file.
func RandLineWriter(w *bufio.Writer, filename string) {
	f, err := os.Open(filename)
	var b string
	var sb strings.Builder

	ck(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	linenumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		if r.Intn(linenumber) == 0 {
			b = line
		}
		linenumber++
	}

	sb.WriteString(b)
	sb.WriteRune('\n')
	w.WriteString(sb.String())
}

// InputLineWriter applies line from input prompt.
func InputLineWriter(w *bufio.Writer, prompt string) {
	b := bufio.NewReader(os.Stdin)
	fmt.Println(prompt + ": ")
	s, err := b.ReadString('\n')
	ck(err)
	w.WriteString(s)
}
