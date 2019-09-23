package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json") // json 化する元のデータ
	source := map[string]string{"Hello": "World"}

	gzipWriter := gzip.NewWriter(w)
	gzipWriter.Header.Name = "test.log"

	multiWriter := io.MultiWriter(os.Stdout, gzipWriter)

	encoder := json.NewEncoder(multiWriter)
	encoder.SetIndent("", " ")
	encoder.Encode(source)

	gzipWriter.Flush() // TODO: needed?
	gzipWriter.Close()
}

// ここにコードを書く
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
