package entity

type ShortLinkInfo struct {
	URL                 string `json:"url"`
	ExpirationInMinutes int32  `json:"expiration_in_minutes"`
	CreatedAt           int64  `json:"created_at"`
}
