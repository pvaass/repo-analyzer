package platforms

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type gitHubContents []struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Sha         string `json:"sha"`
	Size        int    `json:"size"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	GitURL      string `json:"git_url"`
	DownloadURL string `json:"download_url"`
	Type        string `json:"type"`
	Links       struct {
		Self string `json:"self"`
		Git  string `json:"git"`
		HTML string `json:"html"`
	} `json:"_links"`
}

type GitHub struct {
	Token     string
	repoOwner string
	repoName  string
}

func (g *GitHub) SetURI(uri string) {
	uri = strings.Replace(uri, "https://", "", -1)
	uri = strings.Replace(uri, "http://", "", -1)

	parts := strings.Split(uri, "/")

	g.repoOwner = parts[1]
	g.repoName = parts[2]
}

func (GitHub) SupportsURI(uri string) bool {
	return strings.Contains(uri, "github")
}

func (g GitHub) getContentResponse(path string) *http.Response {
	if g.repoOwner == "" || g.repoName == "" {
		panic("call GitHub#SetURI before GitHub#FileList: no URI set.")
	}
	req, err := http.NewRequest("GET", "https://api.github.com/repos/"+g.repoOwner+"/"+g.repoName+"/contents/"+path, nil)
	if err != nil {
		panic(err)
	}

	if len(g.Token) > 0 {
		req.Header.Set("Authorization", "token "+g.Token)
	}
	req.Header.Set("Accept", "application/vnd.github.v3.raw")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panic("Http read error", err)
	}

	return resp
}

func (g GitHub) parseContentResponse(resp *http.Response) gitHubContents {
	var c gitHubContents
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic("Body read error", err)
	}

	if resp.StatusCode >= 300 {
		var ghError struct {
			Message string `json:"message"`
		}
		err3 := json.Unmarshal([]byte(body), &ghError)
		if err3 != nil {
			log.Panic("Invalid Json Decode", err3)
		}
		return c
	}

	err = json.Unmarshal([]byte(body), &c)
	if err != nil {
		log.Panic("Invalid Json Decode", err)
	}

	return c
}

func (g GitHub) FileList(path string) []File {
	var fileList []File

	resp := g.getContentResponse(path)
	defer resp.Body.Close()

	content := g.parseContentResponse(resp)

	for _, element := range content {
		fileList = append(fileList, File{Name: element.Name, DownloadURI: element.DownloadURL, Path: element.Path})
	}
	return fileList
}
