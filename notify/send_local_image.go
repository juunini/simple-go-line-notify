package notify

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

// SendLocalImage : send line notify simple text with upload local image(jpg, png only)
func SendLocalImage(accessToken, message, imagePath string) (err error) {
	body, contentType, err := makeMultipartBody(message, imagePath)

	req, err := http.NewRequest(
		"POST",
		lineNotifyApiURL,
		&body,
	)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", contentType)

	err = sendToLineServer(req, accessToken)
	return
}

func makeMultipartBody(message, imagePath string) (body bytes.Buffer, contentType string, err error) {
	writer := multipart.NewWriter(&body)

	if err = writeBodyInMessage(writer, message); err != nil {
		return
	}
	if err = writeBodyInImageFile(writer, imagePath); err != nil {
		return
	}
	writer.Close()

	contentType = writer.FormDataContentType()
	return
}

func writeBodyInMessage(writer *multipart.Writer, message string) (err error) {
	var messageWriter io.Writer
	messageWriter, err = writer.CreateFormField("message")
	if err != nil {
		return
	}
	if _, err = io.Copy(messageWriter, strings.NewReader(message)); err != nil {
		return
	}
	return
}

func writeBodyInImageFile(writer *multipart.Writer, imagePath string) (err error) {
	var imageFile *os.File
	imageFile, err = os.Open(imagePath)
	if err != nil {
		return
	}

	var imageWriter io.Writer
	imageWriter, err = writer.CreateFormFile("imageFile", imageFile.Name())
	if err != nil {
		return
	}
	if _, err = io.Copy(imageWriter, imageFile); err != nil {
		return
	}
	imageFile.Close()
	return
}
