package pkg

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
)

func GenerateOTP() int {
	return rand.Intn(900000) + 100000
}

func NullStringToString(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}

func UploadFileToTelegraph(fileHeader *multipart.FileHeader) (string, error) {
	src, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	part, err := writer.CreateFormFile("file", fileHeader.Filename)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(part, src)
	if err != nil {
		return "", err
	}
	writer.Close()

	request, err := http.NewRequest("POST", "https://telegra.ph/upload", &buffer)
	if err != nil {
		return "", err
	}
	request.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result []map[string]string
	err = json.Unmarshal(body, &result)
	if err != nil || len(result) == 0 || result[0]["src"] == "" {
		return "", errors.New("invalid response from telegraph")
	}

	return "https://telegra.ph" + result[0]["src"], nil
}
