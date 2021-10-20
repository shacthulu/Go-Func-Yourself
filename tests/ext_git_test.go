package gfytests

import (
	"testing"
)

//switch this repo.  The Python code is missing a parentheses.
func TestGitPy(t *testing.T) {
	submission := Submission{
		SubmissionType:     "git",
		CodeType:           "py",
		GitRepo:            "https://github.com/hcs/bootcamp-python.git",
		EntryPointFileName: "hello_soln.py",
		Code:               "",
		DownloadURL:        "",
	}
	testSkeleton(submission, t, "Python Git")
}

func TestGitJs(t *testing.T) {
	submission := Submission{
		SubmissionType:     "git",
		CodeType:           "js",
		GitRepo:            "https://github.com/PravallikaManchinisetti/helloworld.js.git",
		EntryPointFileName: "helloworld.js",
		Code:               "",
		DownloadURL:        "",
	}
	testSkeleton(submission, t, "JavaScript Git")
}

func TestGitGo(t *testing.T) {
	submission := Submission{
		SubmissionType:     "git",
		CodeType:           "go",
		GitRepo:            "https://github.com/go-training/helloworld.git",
		EntryPointFileName: "main.go",
		Code:               "",
		DownloadURL:        "",
	}
	testSkeleton(submission, t, "Go Git")
}
