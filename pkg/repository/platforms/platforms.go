package platforms

import "errors"

type Platforms struct {
	collection []Platform
}

type Platform interface {
	SupportsURI(uri string) bool
	FileList(path string) []File
	SetURI(uri string)
}

type File struct {
	Name        string
	Path        string
	DownloadURI string
	Content     []byte
}

func (ps Platforms) ForURI(uri string) (Platform, error) {
	for _, platform := range ps.collection {
		if platform.SupportsURI(uri) {
			return platform, nil
		}
	}
	return ps.collection[0], errors.New("No supported platforms for uri \"" + uri + "\".")
}

func (ps *Platforms) Add(p ...Platform) {
	for _, platform := range p {
		ps.collection = append(ps.collection, platform)
	}
}
