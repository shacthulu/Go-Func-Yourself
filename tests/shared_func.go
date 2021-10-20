package gfytests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func testSkeleton(submission Submission, t *testing.T, origin string) {
	jsonStr, err := json.Marshal(submission)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(string(jsonStr)))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Failed %s: malformed response: %s", origin, err)
	}
	if resp.StatusCode == 200 {
		t.Logf("Successful %s: Response is a 200 -> %s, body -> %s", origin, resp.Status, string(body))
	}
	if resp.StatusCode != 200 {
		t.Errorf("Failed %s: Response is not a 200 -> %s, body -> %s", origin, resp.Status, string(body))
	}
	defer resp.Body.Close()
	fmt.Println(string(body))
	//clean()
}

// TestClean cleans up all the files in the /app/ directory.
func TestClean(t *testing.T) {
	dir := "/app/"

	//get the list of files and folders in the current directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	//iterate over the files and folders
	for _, file := range files {
		//if the file is a folder, remove it recursively
		if file.IsDir() {
			fmt.Println("Removing folder:", file.Name())
			os.RemoveAll(filepath.Join(dir, file.Name()))
		} else {
			//if the file is a file, remove it
			fmt.Println("Removing file:", file.Name())
			os.Remove(filepath.Join(dir, file.Name()))
		}
	}
}
