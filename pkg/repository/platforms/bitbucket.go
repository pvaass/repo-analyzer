package platforms

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

type bitBucketContent struct {
	Node        string   `json:"node"`
	Path        string   `json:"path"`
	Directories []string `json:"directories"`
	Files       []struct {
		Size         int       `json:"size"`
		Path         string    `json:"path"`
		Timestamp    time.Time `json:"timestamp"`
		Utctimestamp string    `json:"utctimestamp"`
		Revision     string    `json:"revision"`
	} `json:"files"`
}

type BitBucket struct {
	repoOwner string
	repoName  string
}

func (g *BitBucket) SetURI(uri string) {
	uri = strings.Replace(uri, "https://", "", -1)
	uri = strings.Replace(uri, "http://", "", -1)

	parts := strings.Split(uri, "/")

	g.repoOwner = parts[1]
	g.repoName = parts[2]
}

func (BitBucket) SupportsURI(uri string) bool {
	return strings.Contains(uri, "bitbucket.org")
}

func (g BitBucket) getContentResponse(path string) *http.Response {
	if g.repoOwner == "" || g.repoName == "" {
		panic("call #SetURI before #FileList: no URI set.")
	}

	if path == "." {
		path = ""
	}
	req, err := http.NewRequest("GET", "https://api.bitbucket.org/1.0/repositories/"+g.repoOwner+"/"+g.repoName+"/src/master/"+path, nil)
	if err != nil {
		log.Panic("Invalid request", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panic("Http read error", err)
	}

	return resp
}

func (g BitBucket) parseContentResponse(resp *http.Response) bitBucketContent {
	var c bitBucketContent
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic("Body read error", err)
	}

	if resp.StatusCode >= 300 {
		if resp.StatusCode == 404 {
			return c
		}
		log.Panic("Invalid request")
	}

	err = json.Unmarshal([]byte(body), &c)
	if err != nil {
		log.Panic("Invalid Json Decode", err)
	}

	return c
}

func (g BitBucket) FileList(path string) []File {
	var fileList []File

	resp := g.getContentResponse(path)
	defer resp.Body.Close()

	content := g.parseContentResponse(resp)

	for _, element := range content.Files {
		_, fileName := filepath.Split(element.Path)
		url := "https://api.bitbucket.org/1.0/repositories/" + g.repoOwner + "/" + g.repoName + "/raw/master/" + element.Path
		fileList = append(fileList, File{Name: fileName, DownloadURI: url, Path: element.Path})
	}
	return fileList
}
