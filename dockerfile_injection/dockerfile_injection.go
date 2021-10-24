package main

// Go

// Go application to take in an existing dockerfile, insert a line that copies ./stdin_reader from the local system to /app/stdin_reader in the finished container, and print the new Dockerfile to stdout.  Replace the initial Dockerfile with the new Dockerfile.
//
// Usage:
//   inject_to_dockerfile -dockerfile ./Dockerfile
//   - The COPY command should be changed to COPY ./stdin_reader /app/stdin_reader
//   - The ENTRYPOINT command should be changed to ENTRYPOINT ["./stdin_reader"]
//
// The final Dockerfile should be written to stdout.
//
// Usage:
//   rewrite_dockerfile -dockerfile ./Dockerfile

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dockerfile := flag.String("dockerfile", "", "Dockerfile to inject into")
	flag.Parse()
	injected := false

	if *dockerfile == "" {
		fmt.Println("-dockerfile is required")
		os.Exit(1)
	}

	dockerfile_contents, err := ioutil.ReadFile(*dockerfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dockerfile_lines := strings.Split(string(dockerfile_contents), "\n")

	for i, line := range dockerfile_lines {
		if strings.Contains(line, "COPY") {
			dockerfile_lines = append(dockerfile_lines[:i], append([]string{
				"COPY ./stdin_reader /app/stdin_reader",
			}, dockerfile_lines[i:]...)...)
			injected = true
			break
		}
	}

	if injected == false {

		dockerfile_lines = append(dockerfile_lines, "COPY ./stdin_reader /app/stdin_reader")
	}

	for i, line := range dockerfile_lines {
		if strings.Contains(line, "ENTRYPOINT") {
			dockerfile_lines = append(dockerfile_lines[:i], append([]string{
				"ENTRYPOINT [\"/app/stdin_reader\"]",
			}, dockerfile_lines[i:]...)...)
			break
		}
	}

	fmt.Println(strings.Join(dockerfile_lines, "\n"))
}
