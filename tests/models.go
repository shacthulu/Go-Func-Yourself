package gfytests

type Submission struct {
	SubmissionType     string `json:"submissionType"`
	CodeType           string `json:"codeType"`
	GitRepo            string `json:"gitRepo"`
	EntryPointFileName string `json:"entryPointFileName"`
	Code               string `json:"code"`
	DownloadURL        string `json:"downloadURL"`
}
