package main

import (
	"fmt"
	"strings"
)

func sourceGopher(downstream chan string) {
	for _, v := range []string{
		"hello world",
		"a bad apple",
		"goodbye all",
		"rats getting larger",
		"cats getting thicker",
		"cats getting thicker",
	} {
		downstream <- v
	}
	close(downstream)
}

func removeIdentical(upstream, downstream chan string) {
	var previous string
	for item := range upstream {
		if item != previous {
			downstream <- item
		}
		previous = item
	}
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream)
}

func splitWords(upstream, downstream chan string) {
	for item := range upstream {
		for word := range strings.SplitSeq(item, " ") {
			downstream <- word
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}

type StringStep = func(upstream, downstream chan string)
type StringPipeline struct {
	StartStep func(downstream chan string)
	Steps     []StringStep
	EndStep   func(upstream chan string)
}

func (p StringPipeline) Run() {
	leftChannel := make(chan string)
	go p.StartStep(leftChannel)
	for _, step := range p.Steps {
		rightChannel := make(chan string)
		go step(leftChannel, rightChannel)
		leftChannel = rightChannel
	}
	p.EndStep(leftChannel)

}

func main() {
	pipeline := StringPipeline{
		StartStep: sourceGopher,
		Steps: []StringStep{
			removeIdentical,
			filterGopher,
			splitWords,
		},
		EndStep: printGopher,
	}
	pipeline.Run()
}
