package defs

type ShortenReq struct {
	Url                 string `json:"url" validate:"required"`
	ExpirationInMinutes int64  `json:"expiration_in_minutes" validate:"required"`
}

type ShortenResp struct {
	ShortUrl string `json:"shortUrl"`
	LongUrl  string `json:"longUrl"`
}
