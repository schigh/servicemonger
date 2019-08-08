package b64

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"os"
)

// RenderBase64ForImage will open an image file
// and return its base64 encoded value
func RenderBase64ForImage(imgPath string) (string, error) {
	f, err := os.Open(imgPath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return "", err
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	if err := png.Encode(&b, img); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}
