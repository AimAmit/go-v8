package server

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func HandlePost(r *http.Request, key string) (err error) {

	file, m, err := r.FormFile("code")
	if err != nil {
		return
	}

	err = os.MkdirAll(fmt.Sprintf("code/%s", key), 0777)
	if err != nil {
		return
	}

	reader, err := zip.NewReader(file, m.Size)
	if err != nil {
		return
	}

	var fileNames []string
	for _, f := range reader.File {

		open, err := f.Open()
		if err != nil {
			continue
		}

		//return err
		//log.Println(f.Name)
		//if strings.Contains(f.Name, "._") {
		//	continue
		//}

		openedFile, err := os.OpenFile(fmt.Sprintf("code/%s/%s", key, f.Name), os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			continue
			//return err
		}

		if err != nil {
			return err
		}

		if _, err = io.Copy(openedFile, open); err != nil {
			return err
		}

		fileNames = append(fileNames, f.Name)
	}

	err = Bundle(key)
	log.Println("bundle: ", err)

	return
}
