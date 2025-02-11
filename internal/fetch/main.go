package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/kendfss/but"
	"github.com/kendfss/fopa/internal"
)

func main() {
	response := but.Mustv(http.Get(internal.SourceURL))
	defer response.Body.Close()
	but.Must(os.MkdirAll(filepath.Dir(internal.SourceFilePath), os.ModePerm|os.ModeDir))
	sourceFile := but.Mustv(os.Create(internal.SourceFilePath))
	defer sourceFile.Close()
	_ = but.Mustv(io.Copy(sourceFile, response.Body))
}
