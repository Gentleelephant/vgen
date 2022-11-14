package encode

import "encoding/base64"

func EncodeBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
