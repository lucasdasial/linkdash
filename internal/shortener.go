package internal

import "encoding/base64"

func GetShort(value string) string {

	return base64.StdEncoding.EncodeToString([]byte(value))
}
