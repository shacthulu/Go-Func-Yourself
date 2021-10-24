package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestInjectToDockerfile(t *testing.T) {
	dockerfile := "./test_dockerfile"
	defer os.Remove(dockerfile)

	injected := false

	dockerfile_contents, err := ioutil.ReadFile(dockerfile)
	if err != nil {
		t.Errorf("Error reading Dockerfile: %s", err)
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

	dockerfile_contents = []byte(strings.Join(dockerfile_lines, "\n"))

	err = ioutil.WriteFile(dockerfile, dockerfile_contents, 0644)
	if err != nil {
		t.Errorf("Error writing Dockerfile: %s", err)
	}

	dockerfile_contents, err = ioutil.ReadFile(dockerfile)
	if err != nil {
		t.Errorf("Error reading Dockerfile: %s", err)
	}

	if !bytes.Contains(dockerfile_contents, []byte("COPY ./stdin_reader /app/stdin_reader")) {
		t.Errorf("Dockerfile does not contain expected line: %s", dockerfile_contents)
	}
}

func TestRewriteDockerfile(t *testing.T) {
	dockerfile := "./test_dockerfile"
	defer os.Remove(dockerfile)

	injected := false

	dockerfile_contents, err := ioutil.ReadFile(dockerfile)
	if err != nil {
		t.Errorf("Error reading Dockerfile: %s", err)
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

	dockerfile_contents = []byte(strings.Join(dockerfile_lines, "\n"))

	err = ioutil.WriteFile(dockerfile, dockerfile_contents, 0644)
	if err != nil {
		t.Errorf("Error writing Dockerfile: %s", err)
	}

	dockerfile_contents, err = ioutil.ReadFile(dockerfile)
	if err != nil {
		t.Errorf("Error reading Dockerfile: %s", err)
	}

	if !bytes.Contains(dockerfile_contents, []byte("COPY ./stdin_reader /app/stdin_reader")) {
		t.Errorf("Dockerfile does not contain expected line: %s", dockerfile_contents)
	}

	if !bytes.Contains(dockerfile_contents, []byte("ENTRYPOINT [\"/app/stdin_reader\"]")) {
		t.Errorf("Dockerfile does not contain expected line: %s", dockerfile_contents)
	}
}
