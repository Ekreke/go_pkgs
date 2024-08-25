package examples

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func IOCopyExp() {
	r := strings.NewReader("some io.Reader stream to be read")
	/*
		A successful Copy returns err == nil, not err == EOF. Because Copy is defined to read from src until EOF, it does not treat an EOF from Read as an error to be reported.
	*/
	// stdout is a file , these code below will send data to stdout , this file will be output to a terminal , it's usually is users's terminal of current programme

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

func IOCopyBufExp() {
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	buf := make([]byte, 1)
	// buf is used here...
	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
	}

	// ... reused here also. No need to allocate an extra buffer.
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	}
}

func IOCopyNExp() {
	r := strings.NewReader("some io.Reader stream to be read")
	if _, err := io.CopyN(os.Stdout, r, 4); err != nil {
		log.Fatal(err)
	}
}

func IOPipeExp() {
	// this pipe is based on channel
	r, w := io.Pipe()
	go func() {
		fmt.Fprint(w, "some io.Reader stream to be read")
		w.Close()
	}()
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

func IOReadAllExp() {
	r := strings.NewReader("some io.Reader stream to be read")
	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)
}

func IOReadAtLeastExp() {
	r := strings.NewReader("some io.Reader stream to be read")
	b := make([]byte, 4)
	if _, err := io.ReadAtLeast(r, b, 4); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)
}

func IOReadCloserExp() {
	r := strings.NewReader("some io.Reader stream to be read")
	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)
}
