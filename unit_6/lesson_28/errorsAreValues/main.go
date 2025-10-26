package main

import (
	"fmt"
	"io"
	"os"
)

type safeWriter struct {
	w   io.Writer
	err error
}

func (sw *safeWriter) writeLn(s string) {
	if sw.err != nil {
		return
	}
	_, sw.err = fmt.Fprintln(sw.w, s)
}

func proverbs(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	sw := safeWriter{w: file}
	sw.writeLn("Errors are values.")
	sw.writeLn("Don't just check errors, handle them gracefully.")
	sw.writeLn("Don't panic.")
	sw.writeLn("Make the zero value useful.")
	sw.writeLn("The bigger the interface, the weaker the abstraction.")
	sw.writeLn("interface{} says nothing.")
	sw.writeLn("Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.")
	sw.writeLn("Documentation is for users.")
	sw.writeLn("A little copying is better than a little dependency.")
	sw.writeLn("Clear is better than clever.")
	sw.writeLn("Concurrency is not parallelism.")
	sw.writeLn("Don't communicate by sharing memory, share memory by communicating.")
	sw.writeLn("Channels orchestrate; mutexes serialize.")
	return sw.err
}

func main() {
	err := proverbs("proverbs.txt")
	if err != nil {
		fmt.Println(err)
	}
}
