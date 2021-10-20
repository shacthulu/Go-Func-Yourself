package gfytests

import (
	"testing"
)

//switch this repo.  The Python code is missing a parentheses.
func TestArchivePy(t *testing.T) {
	submission := Submission{
		SubmissionType:     "archive",
		CodeType:           "py",
		GitRepo:            "",
		EntryPointFileName: "hello_soln.py",
		Code:               "",
		DownloadURL:        "https://github.com/hcs/bootcamp-python/archive/refs/heads/master.zip",
	}
	testSkeleton(submission, t, "Python Archive")
}

func TestArchiveJs(t *testing.T) {
	submission := Submission{
		SubmissionType:     "archive",
		CodeType:           "js",
		GitRepo:            "",
		EntryPointFileName: "helloworld.js",
		Code:               "",
		DownloadURL:        "https://github.com/PravallikaManchinisetti/helloworld.js/archive/refs/heads/master.zip",
	}
	testSkeleton(submission, t, "Javascript Archive")
}

func TestArchiveGo(t *testing.T) {
	submission := Submission{
		SubmissionType:     "archive",
		CodeType:           "go",
		GitRepo:            "",
		EntryPointFileName: "main.go",
		Code:               "",
		DownloadURL:        "https://github.com/go-training/helloworld/archive/refs/heads/master.zip",
	}
	testSkeleton(submission, t, "Go Archive")
}
