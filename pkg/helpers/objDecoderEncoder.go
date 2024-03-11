package helpers

import (
	"b2b/m/pkg/errors"
	"bytes"
	"context"
	"crypto/md5"
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

	//bounds := m.Bounds()
	//log.Println(m, bounds, formatString)

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
		err = jpeg.Encode(buff, m, &jpeg.Options{Quality: 100})
		if err != nil {
			log.Println("failed to create buffer", err)
		}
		// pngImageBytes, err := ToPng(m)
		// if err != nil {
		// 	log.Println("failed to create buffer", err)
		// }
		// contentType = "image/png"
		// err = png.Encode(buff, m)
		// if err != nil {
		// 	log.Println("failed to create buffer", err)
		// }

	case "gif":
		contentType = "image/gif"
		err = png.Encode(buff, m)
		if err != nil {
			log.Println("failed to create buffer", err)
		}

	default:
		log.Println("Unsuppotred file type")
		return dec, "", err
	}
	return buff.Bytes(), contentType, nil
}

func CheckSum(s string) int {
	result := 0
	array := md5.Sum([]byte(s))
	for _, v := range array {
		result += int(v)
	}
	return result
}

func EncodeImgToBase64(ctx context.Context, imgInBytes []byte) (encoded string) {
	mimeType := http.DetectContentType(imgInBytes)
	switch mimeType {
	case "image/jpeg":
		encoded += "data:image/jpeg;base64,"
	case "image/png":
		encoded += "data:image/png;base64,"
	case "image/gif":
		encoded += "data:image/gif;base64,"
	}
	encoded += base64.StdEncoding.EncodeToString(imgInBytes)
	return encoded
}

// ToPng converts an image to png
func ToPng(imageBytes []byte) ([]byte, error) {
	contentType := http.DetectContentType(imageBytes)

	switch contentType {
	case "image/png":
	case "image/jpeg":
		img, err := jpeg.Decode(bytes.NewReader(imageBytes))
		if err != nil {
			log.Println(err, "unable to decode jpeg")
			return nil, errors.UnableToDecodeJpeg
		}

		buf := new(bytes.Buffer)
		if err := png.Encode(buf, img); err != nil {
			log.Println(err, "unable to encode png")
			return nil, errors.UnableToEncodePng
		}

		return buf.Bytes(), nil
	}
	log.Println("unable to convert to png", contentType)

	return nil, errors.UnknownImgFormat
}
