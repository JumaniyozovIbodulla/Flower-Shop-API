package pkg

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func SendSMS(phone string, otp int) error {

	url := "https://notify.eskiz.uz/api/message/sms/send"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("mobile_phone", phone)
	_ = writer.WriteField("message", "Bu Eskiz dan test")
	_ = writer.WriteField("from", "4546")
	err := writer.Close()

	if err != nil {
		fmt.Println(err)
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjMxOTU0ODksImlhdCI6MTcyMDYwMzQ4OSwicm9sZSI6InRlc3QiLCJzaWduIjoiYzNmYTg1Y2U3ODIwMWU4ZGMxMTI2ZTMyMzQwZDJmYTdlOGJlNjk5MWVhNDAxN2IyOTYyYmJiNGE3OTkwOTlhOCIsInN1YiI6Ijc3ODcifQ.f0oEqT6uzeDFK1bHzZfRwsXlHZOmBSn5kcTAlOQjHY8"))
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
  
	fmt.Println(string(body))
	return nil
}
