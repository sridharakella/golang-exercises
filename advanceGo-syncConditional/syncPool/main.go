package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// TODO: create pool of bytes.Buffers which can be reused.

// whole below code is added
var pool = sync.Pool{
	New: func() interface{} {
		fmt.Println("creating new buffer")
		return new(bytes.Buffer)
	},
}

func log(w io.Writer, val string) {

	b := pool.Get().(*bytes.Buffer) // lines are added
	b.Reset()                       // lines are added

	b.WriteString(time.Now().Format("15:04:05"))
	b.WriteString(" : ")
	b.WriteString(val)
	b.WriteString("\n")

	w.Write(b.Bytes())
	pool.Put(b) // lines are added
}

func main() {
	log(os.Stdout, "debug-string1")
	time.Sleep(10 * time.Second)
	log(os.Stdout, "debug-string2")
}
