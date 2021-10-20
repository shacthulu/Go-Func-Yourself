package gfytests

import (
	"testing"
)

//switch this repo.  The Python code is missing a parentheses.
func TestFilePy(t *testing.T) {
	submission := Submission{
		SubmissionType:     "file",
		CodeType:           "py",
		GitRepo:            "",
		EntryPointFileName: "",
		Code:               "",
		DownloadURL:        "https://gist.githubusercontent.com/igorzel/3792811/raw/078ab633270b2a80c476e97ae94cad5aa0e9fd6f/helloworld.py",
	}
	testSkeleton(submission, t, "Python File")
}

func TestFileJs(t *testing.T) {
	submission := Submission{
		SubmissionType:     "file",
		CodeType:           "js",
		GitRepo:            "",
		EntryPointFileName: "",
		Code:               "",
		DownloadURL:        "https://gist.githubusercontent.com/modalsoul/3868393/raw/0d504c31b183e4efc52dbaea14ced04c69d6656c/helloworld.js",
	}
	testSkeleton(submission, t, "JavaScript File")
}

func TestFileGo(t *testing.T) {
	submission := Submission{
		SubmissionType:     "file",
		CodeType:           "go",
		GitRepo:            "",
		EntryPointFileName: "",
		Code:               "",
		DownloadURL:        "https://raw.githubusercontent.com/go-training/helloworld/master/main.go",
	}
	testSkeleton(submission, t, "Go File")
}
