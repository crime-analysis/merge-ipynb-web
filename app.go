package merge

import (
	"io"
	"mime/multipart"
	"net/http"
	"sync"

	"github.com/crime-analysis/merge-ipynb"
)

const _24K = (1 << 20) * 24

func init() {
	http.HandleFunc("/", app)
}

func process(files []io.Reader, list []*multipart.FileHeader, ch chan error) {
	wg := sync.WaitGroup{}
	wg.Add(len(list))

	for i, fileheader := range list {
		i := i
		fileheader := fileheader

		go func() {
			defer wg.Done()
			file, err := fileheader.Open()
			files[i] = file
			ch <- err
		}()
	}

	close(ch)
}

// Usage:
// curl \
// > -F "f=@hw3_questions.ipynb" \
// > -F "f=@hw3_p1.ipynb" \
// > http://localhost:8080
func app(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(_24K)

	for key, list := range r.MultipartForm.File {
		// only look for the "f" file array
		if key == "f" {
			files := make([]io.Reader, len(list))
			ch := make(chan error, len(list))
			process(files, list, ch)

			for err := range ch {
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}

			if err := merge.Merge(w, files...); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}
