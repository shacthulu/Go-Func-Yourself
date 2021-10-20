package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// THIS SET OF TESTS IS NOT READY YET!!!!!!!!!!!!!!!!!!!!!!
// TODO: Atomic Tests

// TestDownloadFile tests the DownloadFile function with a simple image
func TestDownloadFile(t *testing.T) {
	url := "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"
	filepath := "google.png"
	err := DownloadFile(filepath, url)
	if err != nil {
		t.Fatal(err)
	}
}

// TestUnzip tests the Unzip function with a local zip file
func TestUnzip(t *testing.T) {
	var filenames []string
	filenames, err := Unzip("test.zip", "test")
	if err != nil {
		t.Errorf("Error unzipping file: %s", err)
	}
	if len(filenames) != 2 {
		t.Errorf("Error unzipping file: %s", err)
	}
}

//TestFindFilePath tests the FindFilePath function
func TestFindFilePath(t *testing.T) {
	pathWithFilename, pathWithoutFilename, _ := findFilePath("main.go")
	if pathWithFilename != "/app/main.go" {
		t.Error("pathWithFilename is not correct")
	}
	if pathWithoutFilename != "/app" {
		t.Error("pathWithoutFilename is not correct")
	}
}

//TestCreateSymLink tests the CreateSymLink function
func TestCreateSymLink(t *testing.T) {
	err := createSymLink("/tmp/test", "/tmp/test_link")
	if err != nil {
		t.Errorf("createSymLink: error creating link: %s\n", err)
	}
	println("createSymLink: symlink created successfully")
}

//TestGetGitRepo tests GetGitRepo function
func TestGetGitRepo(t *testing.T) {
	err := getGitRepo("https://github.com/james-bowman/nlp.git", "/tmp/test")
	if err != nil {
		t.Errorf("TestGetGitRepo: error getting repo: %s\n", err)
	}
}

//TestInstallRequirementsPy tests InstallRequirementsPy function
func TestInstallRequirementsPy(t *testing.T) {
	err := installRequirementsPy("/app/requirements.txt")
	if err != nil {
		t.Errorf("TestInstallRequirementsPy: error installing requirements: %s\n", err)
	}
}

// TestInstallRequirementsGo tests the InstallRequirementsGo function
func TestInstallRequirementsGo(t *testing.T) {
	err := installRequirementsGo("/app/requirements.txt", "/app")
	if err != nil {
		t.Errorf("TestInstallRequirementsGo: error installing requirements: %s\n", err)
	}
}

// TestInstallRequirementsNPM tests installRequirementsNPM function
func TestInstallRequirementsNPM(t *testing.T) {
	err := installRequirementsNPM("/app/requirements.txt")
	if err != nil {
		t.Errorf("TestInstallRequirementsNPM: error installing requirements: %s\n", err)
	}
}

// TestSetEnvVars tests the SetEnvVars function
func TestSetEnvVars(t *testing.T) {
	err := setEnvVars("/app/env.txt")
	if err != nil {
		t.Errorf("TestSetEnvVars: error setting environment variables: %s\n", err)
	}
}

// TestAppendEnvVarsProfile tests the appendEnvVars function
func TestAppendEnvVarsProfile(t *testing.T) {
	err := appendEnvVarsProfile("/app/env.txt")
	if err != nil {
		t.Errorf("TestAppendEnvVarsProfile: error appending environment variables to .profile: %s\n", err)
	}
}

// TestWriteCodeFile function tests the WriteCodeFile function
func TestWriteCodeFile(t *testing.T) {
	err := writeCodeFile("print('hello world')", "py")
	if err != nil {
		t.Errorf("TestWriteCodeFile: error writing code file: %s\n", err)
	}
}

//TestInitFunctionRaw tests the raw submission type
func TestInitFunctionRaw(t *testing.T) {
	println("TestInitFunctionRaw: Testing raw submission")
	//create a request body
	requestBody := map[string]string{
		"submissionType": "raw",
		"code":           "print(\"Hello World\")",
		"codeType":       "go",
	}
	//convert the request body to json
	jsonRequestBody, err := json.Marshal(requestBody)
	if err != nil {
		println("TestInitFunctionRaw: Error marshalling json: " + err.Error())
		t.Fail()
	}
	//create a request
	req, err := http.NewRequest("POST", "/apiv1/init", bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		println("TestInitFunctionRaw: Error creating request: " + err.Error())
		t.Fail()
	}
	//create a response recorder
	rr := httptest.NewRecorder()
	//serve the request
	InitFunction(rr, req)
	//check the status code
	if status := rr.Code; status != http.StatusOK {
		println("TestInitFunctionRaw: handler returned wrong status code: got " + string(rune(status)) + " want " + string(rune(http.StatusOK)))
		t.Fail()
	}
	//check the response body
	expected := `{"message":"Function created successfully"}`
	if rr.Body.String() != expected {
		println("TestInitFunctionRaw: handler returned unexpected body: got " + rr.Body.String() + " want " + expected)
		t.Fail()
	}
}

//TestInitFunctionGit tests the git submission type
func TestInitFunctionGit(t *testing.T) {
	println("TestInitFunctionGit: Testing git submission")
	//create a request body
	requestBody := map[string]string{
		"submissionType":     "git",
		"gitRepo":            "https://github.com/james-jones/test-function.git",
		"entryPointFileName": "main.go",
	}
	//convert the request body to json
	jsonRequestBody, err := json.Marshal(requestBody)
	if err != nil {
		println("TestInitFunctionGit: Error marshalling json: " + err.Error())
		t.Fail()
	}
	//create a request
	req, err := http.NewRequest("POST", "/apiv1/init", bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		println("TestInitFunctionGit: Error creating request: " + err.Error())
		t.Fail()
	}
	//create a response recorder
	rr := httptest.NewRecorder()
	//serve the request
	InitFunction(rr, req)
	//check the status code
	if status := rr.Code; status != http.StatusOK {
		println("TestInitFunctionGit: handler returned wrong status code: got " + string(rune(status)) + " want " + string(rune(http.StatusOK)))
		t.Fail()
	}
	//check the response body
	expected := `{"message":"Function created successfully"}`
	if rr.Body.String() != expected {
		println("TestInitFunctionGit: handler returned unexpected body: got " + rr.Body.String() + " want " + expected)
		t.Fail()
	}
}

//TestInitFunctionFile tests the file submission type
func TestInitFunctionFile(t *testing.T) {
	println("TestInitFunctionFile: Testing file submission")
	//create a request body
	requestBody := map[string]string{
		"submissionType":     "file",
		"entryPointFileName": "main.go",
		"downloadURL":        "https://raw.githubusercontent.com/james-jones/test-function/master/main.go",
	}
	//convert the request body to json
	jsonRequestBody, err := json.Marshal(requestBody)
	if err != nil {
		println("TestInitFunctionFile: Error marshalling json: " + err.Error())
		t.Fail()
	}
	//create a request
	req, err := http.NewRequest("POST", "/apiv1/init", bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		println("TestInitFunctionFile: Error creating request: " + err.Error())
		t.Fail()
	}
	//create a response recorder
	rr := httptest.NewRecorder()
	//serve the request
	InitFunction(rr, req)
	//check the status code
	if status := rr.Code; status != http.StatusOK {
		println("TestInitFunctionFile: handler returned wrong status code: got " + string(rune(status)) + " want " + string(rune(http.StatusOK)))
		t.Fail()
	}
	//check the response body
	expected := `{"message":"Function created successfully"}`
	if rr.Body.String() != expected {
		println("TestInitFunctionFile: handler returned unexpected body: got " + rr.Body.String() + " want " + expected)
		t.Fail()
	}
}

//TestInitFunctionArchive tests the archive submission type
func TestInitFunctionArchive(t *testing.T) {
	println("TestInitFunctionArchive: Testing archive submission")
	//create a request body
	requestBody := map[string]string{
		"submissionType":     "archive",
		"entryPointFileName": "main.go",
		"downloadURL":        "https://github.com/james-jones/test-function/archive/master.zip",
	}
	//convert the request body to json
	jsonRequestBody, err := json.Marshal(requestBody)
	if err != nil {
		println("TestInitFunctionArchive: Error marshalling json: " + err.Error())
		t.Fail()
	}
	//create a request
	req, err := http.NewRequest("POST", "/apiv1/init", bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		println("TestInitFunctionArchive: Error creating request: " + err.Error())
		t.Fail()
	}
	//create a response recorder
	rr := httptest.NewRecorder()
	//serve the request
	InitFunction(rr, req)
	//check the status code
	if status := rr.Code; status != http.StatusOK {
		println("TestInitFunctionArchive: handler returned wrong status code: got " + string(rune(status)) + " want " + string(rune(http.StatusOK)))
		t.Fail()
	}
	//check the response body
	expected := `{"message":"Function created successfully"}`
	if rr.Body.String() != expected {
		println("TestInitFunctionArchive: handler returned unexpected body: got " + rr.Body.String() + " want " + expected)
		t.Fail()
	}
}

//TestInitFunctionInvalidSubmissionType tests the invalid submission type
func TestInitFunctionInvalidSubmissionType(t *testing.T) {
	println("TestInitFunctionInvalidSubmissionType: Testing invalid submission type")
	//create a request body
	requestBody := map[string]string{
		"submissionType": "invalid",
		"code":           "print(\"Hello World\")",
		"codeType":       "go",
	}
	//convert the request body to json
	jsonRequestBody, err := json.Marshal(requestBody)
	if err != nil {
		println("TestInitFunctionInvalidSubmissionType: Error marshalling json: " + err.Error())
		t.Fail()
	}
	//create a request
	req, err := http.NewRequest("POST", "/apiv1/init", bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		println("TestInitFunctionInvalidSubmissionType: Error creating request: " + err.Error())
		t.Fail()
	}
	//create a response recorder
	rr := httptest.NewRecorder()
	//serve the request
	InitFunction(rr, req)
	//check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		println("TestInitFunctionInvalidSubmissionType: Wrong status code. Got: " + strconv.Itoa(status) + ", Expected: " + strconv.Itoa(http.StatusBadRequest))
		t.Fail()
	}
}

// TestPullDependenciesGo tests pulling dependencies of a go.mod file
func TestPullDependenciesGo(t *testing.T) {
	//given
	codeType := "go"
	requirementsFile := "go.mod"
	entryPath := "/app/main/"
	//when
	err := pullDependencies(codeType, requirementsFile, entryPath)
	//then
	assert.Nil(t, err)
}

// TestPullDependenciesPy tests pulling dependencies of a requirements.txt file
func TestPullDependenciesPy(t *testing.T) {
	//given
	codeType := "py"
	requirementsFile := "requirements.txt"
	entryPath := "/app/main/"
	//when
	err := pullDependencies(codeType, requirementsFile, entryPath)
	//then
	assert.Nil(t, err)
}

// TestPullDependenciesJs tests pulling dependencies of a package.json file
func TestPullDependenciesJs(t *testing.T) {
	//given
	codeType := "js"
	requirementsFile := "package.json"
	entryPath := "/app/main/"
	//when
	err := pullDependencies(codeType, requirementsFile, entryPath)
	//then
	assert.Nil(t, err)
}
