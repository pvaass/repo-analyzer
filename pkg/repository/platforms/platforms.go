package platforms

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

func (ps Platforms) ForURI(uri string) Platform {
	for _, platform := range ps.collection {
		if platform.SupportsURI(uri) {
			return platform
		}
	}
	panic("No platform!")
}

func (ps *Platforms) Add(p Platform) {
	ps.collection = append(ps.collection, p)
}
