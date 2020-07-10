package common

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/SunMaybo/jewel-crawler/common/sign"
	"github.com/SunMaybo/jewel-crawler/common/uuid"
	"net/url"
	"strings"
)

func GenerateRandomID() string {
	_uuid := uuid.NewV4()
	return strings.ReplaceAll(_uuid.String(), "-", "")
}

func Signature(obj interface{}) string {
	signature := sign.Encode(obj)
	hash := sha1.New()
	hash.Write([]byte(signature))
	return hex.EncodeToString(hash.Sum(nil))
}

func SignatureMap(data map[string]string) string {
	p := url.Values{}
	for s, s2 := range data {
		p.Add(s, s2)
	}
	hash := sha1.New()
	hash.Write([]byte(p.Encode()))
	return hex.EncodeToString(hash.Sum(nil))
}
