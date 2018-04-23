package save

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

func Image(imageurl string) {
	purl, err := url.Parse(imageurl)
	filepath := "./img/" + path.Base(purl.Path)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		response, e := http.Get(imageurl)
		if e != nil {
			log.Fatal(e)
		}

		defer response.Body.Close()

		//open a file for writing
		file, err := os.Create(filepath)
		if err != nil {
			log.Fatal(err)
		}
		// Use io.Copy to just dump the response body to the file. This supports huge files
		_, err = io.Copy(file, response.Body)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		fmt.Println("Success!")
		return
	}

	fmt.Println("Image Exists!")

}

func ImagePathFromURL(imageurl string) string {
	purl, err := url.Parse(imageurl)

	if err != nil {
		log.Fatal(err)
	}

	return "./img/" + path.Base(purl.Path)
}
