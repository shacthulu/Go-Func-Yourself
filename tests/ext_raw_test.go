package gfytests

import (
	"testing"
)

//switch this repo.  The Python code is missing a parentheses.
func TestRawPy(t *testing.T) {
	submission := Submission{
		SubmissionType:     "raw",
		CodeType:           "py",
		GitRepo:            "",
		EntryPointFileName: "",
		Code:               "print(\"hello\")",
		DownloadURL:        "",
	}
	testSkeleton(submission, t, "Raw Python")
}

func TestRawJs(t *testing.T) {
	submission := Submission{
		SubmissionType:     "raw",
		CodeType:           "js",
		GitRepo:            "",
		EntryPointFileName: "",
		Code:               "console.log(\"hello\")",
		DownloadURL:        "",
	}
	testSkeleton(submission, t, "RawJavascript")
}

func TestRawGo(t *testing.T) {
	submission := Submission{
		SubmissionType:     "raw",
		CodeType:           "go",
		GitRepo:            "",
		EntryPointFileName: "",
		Code:               "package main\nimport \"fmt\"\nfunc main(){\nfmt.Println(\"hello\")\n}",
		DownloadURL:        "",
	}
	testSkeleton(submission, t, "Raw Go")
}
