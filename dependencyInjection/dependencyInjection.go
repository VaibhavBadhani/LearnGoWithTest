package dependencyinjection

import (
	"bytes"
	"fmt"
	"io"
	"time"
)

// func Greet(writer *bytes.Buffer, name string) {
// 	fmt.Printf("Hello, %s", name)
// }

func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func Countdown(writer io.Writer, start int) {
	for i := 3; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(writer, i)
	}
	fmt.Fprint(writer, "Go!")
}
