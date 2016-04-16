package main

import (
	"fmt"
  "os"
	"strings"

	"gopkg.in/pipe.v2"
)

func main() {
	defer ExitHandler()

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "nothing to do\n")
		Exit(2)
	}

  pipes := make([]pipe.Pipe, 0, 255)
	pipes = append(pipes, make_base(os.Args[1]))

  for _, arg := range os.Args[2:] {
		pipes = append(pipes, make_delta(arg))
  }

	pipeline := pipe.Line(pipes...)
	fmt.Fprintf(os.Stderr, "Dump: %v\n", pipeline)
/*
	p := pipe.Line(
		pipe.Exec("df"),
		pipe.Filter(func(line []byte) bool {
			return bytes.HasSuffix(line, []byte(" /boot"))
		}),
		pipe.Tee(b),
		pipe.WriteFile("boot.txt", 0644),
	)
	err := pipe.Run(p)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
*/
}

func make_base(filename string) pipe.Pipe {
	if strings.HasSuffix(filename, ".tar.gz") {
		f, err := os.OpenFile(filename, os.O_RDONLY, 0777)

		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v (opening %v)\n", err, filename)
			Exit(1)
		}

		fmt.Fprintf(os.Stderr, "Adding base file:  %v\n", filename)
		return pipe.Read(f)
	}
}

func make_delta(filename string) pipe.Pipe {
	fmt.Fprintf(os.Stderr, "Adding delta file: %v\n", filename)
	return pipe.ChDir("/")
}
