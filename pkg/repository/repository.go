package repository

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"path/filepath"

	"github.com/pvaass/repo-analyzer/pkg/repository/platforms"
)

// Repository is the only exported struct in the repository package.
// It provides a way to interact with different git hosting platforms.
type Repository struct {
	URI      string
	Files    map[string][]platforms.File
	Platform platforms.Platform
}

// List returns a list of files from a path in a repository
// and stores them in Repository for later use, avoiding multiple api requests.
func (r *Repository) List(path string) []platforms.File {
	elems, ok := r.Files[path]
	if !ok {
		elems = r.Platform.FileList(path)
		r.Files[path] = elems
	}
	return r.Files[path]
}

// File returns the contents of a file from the repository.
// When no such file exists, it returns error
func (r *Repository) File(name string) ([]byte, error) {
	path := filepath.Dir(name)
	elems, ok := r.Files[path]
	if !ok {
		elems = r.Platform.FileList(path)
		r.Files[path] = elems
	}

	for index, file := range r.Files[path] {
		if name == file.Path {
			if len(file.Content) == 0 {
				file.Content = download(file.DownloadURI)
				r.Files[path][index] = file
			}
			return file.Content, nil
		}
	}

	return []byte{}, errors.New("Could not find file " + name + " in remote repository")
}

// New configures and returns a new Repository
func New(platform platforms.Platform, uri string) Repository {
	r := new(Repository)
	platform.SetURI(uri)
	r.URI = uri
	r.Files = make(map[string][]platforms.File)
	r.Platform = platform

	return *r
}

func download(url string) []byte {
	log.Println("Downloading " + url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panic("Http read error", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic("Body read error", err)
	}
	return body
}
