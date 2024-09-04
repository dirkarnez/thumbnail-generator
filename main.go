package main

import (
	"os"
	"log"
	"net/http"
	"io"
	"path/filepath"
	thumbnail "github.com/prplecake/go-thumbnail"
)

func main() {
	testImageURL := "https://upload.wikimedia.org/wikipedia/commons/5/56/Donald_Trump_official_portrait.jpg"
	resp, err := http.Get(testImageURL)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	gen := thumbnail.NewGenerator(thumbnail.Generator{Scaler: "CatmullRom"})
	i, err := gen.NewImageFromByteArray(data)
	if err != nil {
		log.Fatal(err)
	}

	dest := filepath.Base(testImageURL)

	thumbBytes, err := gen.CreateThumbnail(i)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(dest, thumbBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

	