package interfaces

import "strings"

type ShortenReq struct {
	Url                 string `json:"url"`                   // 长链
	ExpirationInMinutes int32  `json:"expiration_in_minutes"` // 有效期，单位：分钟
}

func (r *ShortenReq) CheckRequiredParam() bool {
	if (strings.HasPrefix(r.Url, "http://") || strings.HasPrefix(r.Url, "https://")) && r.ExpirationInMinutes > 0 {
		return true
	}
	return false
}

type ShortenResp struct {
	ShortUrl string `json:"shortUrl"`
	LongUrl  string `json:"longUrl"`
}
