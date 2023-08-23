package helpers

import (
	"bytes"
	"context"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"strings"
)

func DecodeImgFromBase64(ctx context.Context, imgBase64 string) (decoded []byte, contentType string, err error) {
	var dec []byte

	b64data := imgBase64[strings.IndexByte(imgBase64, ',')+1:]

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(b64data))
	m, formatString, err := image.Decode(reader)
	if err != nil {
		log.Println("Error while image.Decode(reader)", err)
		return dec, "", err
	}

	bounds := m.Bounds()
	log.Println(m, bounds, formatString)

	buff := new(bytes.Buffer)
	switch formatString {
	case "png":
		contentType = "image/png"
		err = png.Encode(buff, m)
		if err != nil {
			log.Println("failed to create buffer", err)
		}

	case "jpeg", "jpg":
		contentType = "image/jpeg"
		err = jpeg.Encode(buff, m, &jpeg.Options{})
		if err != nil {
			log.Println("failed to create buffer", err)
		}

	default:
		log.Println("Unsuppotred file type")
		return dec, "", err
	}
	return buff.Bytes(), contentType, nil
}

func EncodeImgToBase64(ctx context.Context, imgInBytes []byte) (encoded string) {
	mimeType := http.DetectContentType(imgInBytes)
	switch mimeType {
	case "image/jpeg":
		encoded += "data:image/jpeg;base64,"
	case "image/png":
		encoded += "data:image/png;base64,"
	}
	encoded += base64.StdEncoding.EncodeToString(imgInBytes)
	return encoded
}
