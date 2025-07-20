package pdfmanager

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

func SendToMergeService(pdfs [][]byte, servicePath string) ([]byte, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	for i, file := range pdfs {
		fw, _ := writer.CreateFormFile("files", fmt.Sprintf("file%d.pdf", i+1))
		fw.Write(file)
	}

	writer.Close()

	req, _ := http.NewRequest("POST", servicePath, &buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
