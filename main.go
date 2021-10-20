package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mholt/archiver/v3"
	"gopkg.in/src-d/go-git.v4"
)

// DownloadFile will download a the file at url into the destination filepath
func DownloadFile(filepath string, url string) error {
	fmt.Printf("DownloadFile: Downloading File: %s to %s\n", url, filepath)
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("DownloadFile: Error downloading: %s\n", err)
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		fmt.Printf("DownloadFile: Error creating file: %s\n", err)
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("DownloadFile: Error writing file: %s\n", err)
	}
	return err
}

//Unzip takes an arbitrary archive and decompresses it to the destination directiory
func Unzip(src string, destination string) ([]string, error) {
	fmt.Printf("Unzip: Unzipping file: %s to %s\n", src, destination)
	var filenames []string

	// unzipping the archive to the destination
	err := archiver.Unarchive(src, destination)
	if err != nil {
		fmt.Printf("Unzip: Error unzipping file: %s\n", err)
		return filenames, err
	}

	// iterate through the files and append them to the slice
	files, err := filepath.Glob(filepath.Join(destination, "*"))
	if err != nil {
		fmt.Printf("Unzip: Error reading detination: %s\n", err)
	}
	for _, file := range files {
		filenames = append(filenames, file)
		fmt.Printf("Unzip: appended file to array %s\n", file)
	}
	return filenames, err
}

//find the full path entry entryPointFileName starting in the /app/ directory and return it in the format /app/dir1/dir2/main.go
func findFilePath(entryPointFileName string) (pathWithFilename string, pathWithoutFilename string, err error) {
	fmt.Printf("findFilePath: finding the path to the entry point file: %s\n", entryPointFileName)
	fullNameAndPath := ""
	// looking through the /app/ directory to find the entryPointFileName
	err = filepath.Walk("/app/",
		// walk it, grrrl
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() == entryPointFileName {
				fmt.Printf("findFilePath: Entrypoint file found at: %s\n", entryPointFileName)
				fullNameAndPath = path
			}
			return nil
		})
	if err != nil {
		fmt.Printf("findFilePath: Entrypoint file found at: %s\n", entryPointFileName)
		return pathWithFilename, pathWithoutFilename, err
	}
	if entryPointFileName == "" {
		fmt.Printf("findFilePath: Entrypoint file not found\n")
		return "", "", fmt.Errorf("Entrypoint file not found")
	}
	pathWithFilename = fullNameAndPath
	//TODO: fix any error that could occur if the entrypointfilename is also a directory name in the path to said entrypointfilename
	pathWithoutFilename = strings.Replace(pathWithFilename, entryPointFileName, "", 1)
	fmt.Printf("findFilePath: entry point file found at: %s with the parent directory: %s\n", pathWithFilename, pathWithoutFilename)
	return pathWithFilename, pathWithoutFilename, err
}

//Create a symlink
func createSymLink(src string, dst string) error {
	println("createSymLink " + src + " ->" + dst)
	err := os.Symlink(src, dst)
	if err != nil {
		fmt.Printf("createSymLink: error creating link: %s\n", err)
	}
	println("createSymLink: symlink created successfully")
	return err
}

//clone a git repo to a dest directory. src should end in .git
func getGitRepo(src string, dst string) error {
	println("getGitRepo " + src + " and placing it at " + dst)
	_, err := git.PlainClone(dst, false, &git.CloneOptions{
		URL: src,
	})
	if err != nil {
		fmt.Printf("getGitRepo: error getting repo: %s\n", err)
		return err
	}
	return nil
}

//Install requirements from a Python requirements file.
func installRequirementsPy(requirementsFile string) error {
	println("installRequirementsPy: install requirements from " + requirementsFile)
	cmd := exec.Command("pip3", "install", "-r", requirementsFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("installRequirementsPy: error installing requirements: %s\n", err)
		return err
	}
	return nil
}

//Install requirements from a go.mod file
func installRequirementsGo(requirementsFile string, entryPath string) error {
	println("installRequirementsPy: install requirements from " + requirementsFile + " at " + entryPath)
	cmd := exec.Command("go", "get", "-u", "-v", "entryPath/...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("installRequirementsGo: error installing requirements: %s\n", err)
		return err
	}
	return nil
}

//Install requirements from a package.json file
func installRequirementsNPM(requirementsFile string) error {
	println("installRequirementsNPM: install requirements from " + requirementsFile)
	cmd := exec.Command("npm", "install")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("installRequirementsNPM: error installing requirements: %s\n", err)
		return err
	}
	return nil
}

//Set environment variables listed in var=value format (one per line) in a file named /app/.env
func setEnvVars(envFile string) error {
	println("setEnvVars: setting environment variables from " + envFile)
	file, err := os.Open(envFile)
	if err != nil {
		fmt.Printf("setEnvVars: error opening file: %s\n", err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "=") {
			parts := strings.Split(line, "=")
			err = os.Setenv(parts[0], parts[1])
		}
	}
	if err != nil {
		fmt.Printf("setEnvVars: error setting environment variables: %s\n", err)
		return err
	}
	return nil
}

//Add environment variables listed in var=value format (one per line) in to the ~/.profile file
func appendEnvVarsProfile(envFile string) error {
	println("appendEnvVarsProfile: appending environment variables to .profile from " + envFile)
	file, err := os.Open(envFile)
	if err != nil {
		fmt.Printf("appendEnvVarsProfile: error opening environment file: %s\n", err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "=") {
			parts := strings.Split(line, "=")
			f, err := os.OpenFile(".profile", os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				fmt.Printf("appendEnvVarsProfile: error opening .profile: %s\n", err)
				return err
			}
			defer f.Close()
			if _, err = f.WriteString("export " + parts[0] + "=" + parts[1] + "\n"); err != nil {
				if err != nil {
					fmt.Printf("appendEnvVarsProfile: error writing to .profile: %s\n", err)
				}
				return err
			}
		}
	}
	return nil
}

// write raw string to code file based on the given codeType
func writeCodeFile(rawCode string, codeType string) (err error) {
	println("writeCodeFile: executing with codeType: ", codeType)
	bytesWritten := 0
	switch codeType {
	//python case
	case "py":
		println("writeCodeFile: writing code file for codeType: ", codeType)
		f, err := os.OpenFile("/app/main.py", os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			println("writeCodeFile: could not open app/main.py for writing: ", err)
			return err
		}
		defer f.Close()
		if bytesWritten, err = f.WriteString(rawCode); err != nil {
			println("writeCodeFile: could not write code to file: ", err)
			return err
		}
		if err := f.Close(); err != nil {
			println("writeCodeFile: could not close file: ", err)
			return err
		}
		fmt.Printf("writeCodeFile: successfully wrote Python code file with %d bytes\n", bytesWritten)
		return nil
	//go case
	case "go":
		println("writeCodeFile: writing code file for codeType: ", codeType)
		f, err := os.OpenFile("/app/main.go", os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			println("writeCodeFile: could not open app/main.go for writing: ", err)
			return err
		}
		defer f.Close()
		if bytesWritten, err = f.WriteString(rawCode); err != nil {
			println("writeCodeFile: could not write code to file: ", err)
			return err
		}
		if err := f.Close(); err != nil {
			println("writeCodeFile: could not close file: ", err)
			return err
		}
		fmt.Printf("writeCodeFile: successfully wrote go code file with %d bytes\n", bytesWritten)
		return nil
	//node case
	case "js":
		println("writeCodeFile: writing code file for codeType: ", codeType)
		f, err := os.OpenFile("/app/main.js", os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			println("writeCodeFile: could not open /app/main.js for writing: ", err)
			return err
		}
		defer f.Close()
		if bytesWritten, err = f.WriteString(rawCode); err != nil {
			println("writeCodeFile: could not write code to file: ", err)
			return err
		}
		if err := f.Close(); err != nil {
			println("writeCodeFile: could not close file: ", err)
			return err
		}
		fmt.Printf("writeCodeFile: successfully wrote js code file with %d bytes\n", bytesWritten)
		return nil
	//default case
	default:
		fmt.Printf("writeCodeFile: invalid code type: %s, type must be one of js, py, or go", codeType)
		return errors.New("Invalid code type: " + codeType)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/apiv1/init", InitFunction)
	http.Handle("/", r)
	//TODO: remove hardcoded values
	println("Starting Server at 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//InitFunction is the handler for the /apiv1/init endpoint
func InitFunction(w http.ResponseWriter, r *http.Request) {
	println("InitFunction: REST call triggered, function creation started")
	//get the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("InitFunction: failed to read request body: %s\n", err)
		http.Error(w, "unable to init function", http.StatusInternalServerError)
		panic(err)
	}
	//convert the request body to a map
	var requestBody map[string]string
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		fmt.Printf("InitFunction: failed to read unmarshall json %s\n", err)
		http.Error(w, "unable to init function", http.StatusInternalServerError)
		panic(err)
	}
	//get the submission type from the request body.  Can be one of raw, git, file, archive
	submissionType := requestBody["submissionType"]
	//get the code from the request body. ignored if submissionType isn't raw
	code := requestBody["code"]
	//get the codeType from the request body. Javascript, Go, or Python is valid for now
	codeType := requestBody["codeType"]
	//get the entryPointFileName from the request body.  This will be executed with the input parameters when invoked. Ignored if submissionType is raw or file.
	entryPointFileName := requestBody["entryPointFileName"]
	//get the gitRepo from the request body. Only for git submissionType.
	gitRepo := requestBody["gitRepo"]
	//TODO: authenticated github
	downloadURL := requestBody["downloadURL"]
	switch submissionType {
	//raw submission type
	case "raw":
		println("InitFunction: Processing raw submission")
		//write the code to the main.go file
		err = writeCodeFile(code, codeType)
		if err != nil {
			println("InitFunction: Exiting InitFunction due to Error")
			http.Error(w, "unable to init function", http.StatusInternalServerError)
			return
		}
		if codeType == "go" {
			//compile the code
			println("InitCommand: Compiling go code")
			cmd := exec.Command("go", "build", "-o", "/app/main", "/app/main.go")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				println("InitCommand: Compiling go code: " + err.Error())
				http.Error(w, "unable to init function", http.StatusInternalServerError)
				return
			}
		}
		w.WriteHeader(http.StatusOK)
	//git submission type
	case "git":
		//clone the git repo
		println("InitFunction: Processing git repo submission")
		err = getGitRepo(gitRepo, "/app/")
		if err != nil {
			println("InitFunction: Exiting InitFunction due to Error")
			http.Error(w, "unable to init function", http.StatusInternalServerError)
			return
		}
		mainFileFullPath, mainPath, err := findFilePath(entryPointFileName)
		if err != nil {
			println("InitFunction: Exiting InitFunction due to Error")
			http.Error(w, "unable to init function", http.StatusInternalServerError)
			return
		}
		//set the env vars
		err = setEnvVars(mainPath + ".env")
		if err != nil {
			println("InitFunction: No environment variables set")
		}
		err = appendEnvVarsProfile(mainPath + ".env")
		if err != nil {
			println("InitFunction: No environment variables appended to profile")
		}
		//build the code
		if codeType == "go" {
			println("InitCommand: Compiling go code")
			cmd := exec.Command("go", "build", "-o", "/app/main", mainFileFullPath)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				println("InitCommand: Error compiling go code: " + err.Error())
				http.Error(w, "unable to init function", http.StatusInternalServerError)
				return
			}
			mainFileFullPath = "/app/main"
			err = createSymLink(mainFileFullPath, "/app/main")
			if err != nil {
				println("InitCommand: Error creating symlink: " + err.Error())
				http.Error(w, "unable to init function", http.StatusInternalServerError)
				return
			}
		}
		if codeType != "go" {
			err = createSymLink(mainFileFullPath, "/app/main."+codeType)
			if err != nil {
				println("InitCommand: Error creating symlink: \n" + err.Error())
				http.Error(w, "unable to init function", http.StatusInternalServerError)
				return
			}
		}
		w.WriteHeader(http.StatusOK)
	//file submission type
	case "file":
		//find the full path entry entryPointFileName starting in the /app/ directory and return it in the format /app/dir1/dir2/main.go
		println("InitFunction: Processing file submission")
		err = DownloadFile("/app/main."+codeType, downloadURL)
		if err != nil {
			println("InitFunction: Exiting InitFunction due to Error")
			http.Error(w, "unable to init function", http.StatusInternalServerError)
			return
		}

		//run the compiler
		if codeType == "go" {
			println("InitCommand: Compiling go code")
			cmd := exec.Command("go", "build", "-o", "/app/main", "/app/main.go")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				println("InitCommand: Compiling go code: " + err.Error())
				http.Error(w, "unable to init function", http.StatusInternalServerError)
				return
			}
		}
		w.WriteHeader(http.StatusOK)
	//archive submission type
	case "archive":
		println("InitFunction: Processing archive submission")
		//Download the archive
		err = DownloadFile("/app/tmp.zip", downloadURL)
		if err != nil {
			println("InitFunction: Exiting InitFunction due to Error")
			http.Error(w, "unable to init function", http.StatusInternalServerError)
			return
		}
		_, err = Unzip("/app/tmp.zip", "/app")
		if err != nil {
			println("InitFunction: Exiting InitFunction due to Error")
			http.Error(w, "unable to init function", http.StatusInternalServerError)
			return
		}
		mainFileFullPath, mainPath, err := findFilePath(entryPointFileName)
		if err != nil {
			println("InitFunction: Exiting InitFunction due to Error")
			http.Error(w, "unable to init function", http.StatusInternalServerError)
			return
		}
		//set the env vars
		setEnvVars(mainPath + ".env")
		if err != nil {
			println("InitFunction: No environment variables set")
		}
		appendEnvVarsProfile(mainPath + ".env")
		if err != nil {
			println("InitFunction: No environment variables appended to profile")
		}
		//build the code
		if codeType == "go" {
			println("InitCommand: Compiling go code")
			cmd := exec.Command("go", "build", "-o", "/app/main", mainFileFullPath)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				println("InitCommand: Compiling go code: " + err.Error())
				http.Error(w, "unable to init function", http.StatusInternalServerError)
				return
			}
			err = createSymLink(mainFileFullPath, "/app/main")
			if err != nil {
				println("InitCommand: Error creating symlink: " + err.Error())
				http.Error(w, "unable to init function", http.StatusInternalServerError)
				return
			}
		}
		if codeType != "go" {
			createSymLink(mainFileFullPath, "/app/main."+codeType)
			if err != nil {
				println("InitCommand: Error creating symlink: " + err.Error())
				http.Error(w, "unable to init function", http.StatusInternalServerError)
				return
			}
		}
		w.WriteHeader(http.StatusOK)
	//default case
	default:
		print("invalid submissionType: " + submissionType)
		if err != nil {
			println("InitCommand: Error: " + err.Error())
			http.Error(w, "unable to init function", http.StatusInternalServerError)
			return
		}
	}
}

//TODO: Integrate pullDependencies
func pullDependencies(codeType string, requirementsFile string, entryPath string) error {
	fmt.Printf("PullDependencies for codeType %s, requirementsFile %s, entryPath %s", codeType, requirementsFile, entryPath)
	switch codeType {
	case "go":
		println("PullDependencies for codeType go")
		err := installRequirementsGo(requirementsFile, entryPath)
		return err
	case "py":
		println("PullDependencies for codeType py")
		err := installRequirementsPy(requirementsFile)
		return err
	case "js":
		println("PullDependencies for codeType js")
		err := installRequirementsNPM(requirementsFile)
		return err
	}
	return nil
}

/*TODO:
- add error handling
- add logging
- add a /healthz endpoint
- add a /readyz endpoint
- add security
- fix camelCase
- start the function server
- end the init server
- go clean
*/

// Unit Tests:
//go test -v -cover -coverprofile=coverage.out
//go tool cover -html=coverage.out
