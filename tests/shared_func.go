package gfytests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
}
