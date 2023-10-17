package mimetype

import "net/http"

const undetectedMIMEType = "application/octet-stream"

func GetMIMEtype(content []byte) string {
	mimeType := http.DetectContentType(content)
	if mimeType == undetectedMIMEType {
		return "undetectable"
	}

	return mimeType
}

func Contains(bannedMIMETypes []string, mimeType string) bool {
	for _, str := range bannedMIMETypes {
		if str == mimeType {
			return true
		}
	}

	return false
}
