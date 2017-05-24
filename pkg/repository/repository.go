package repository

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pvaass/repo-analyzer/pkg/repository/platforms"
)

type Repository struct {
	URI   string
	Files []platforms.File
}

func (r Repository) FileNames() []string {
	var names []string
	for _, file := range r.Files {
		names = append(names, file.Name)
	}
	return names
}

func (r *Repository) File(name string) []byte {
	for index, file := range r.Files {
		if name == file.Name {
			if len(file.Content) == 0 {
				file.Content = download(file.DownloadURI)
				r.Files[index] = file
			}
			return file.Content
		}
	}
	panic("No such file")
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

func New(platform platforms.Platform, uri string) Repository {
	r := new(Repository)
	platform.SetURI(uri)
	r.URI = uri
	r.Files = platform.FileList("/")

	return *r
}
