package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/mattheath/base62"
	"github.com/spaolacci/murmur3"
	"time"
	"web-shortlink/app/domain/entity"
	"web-shortlink/app/infra/rds"
	"web-shortlink/pkg/utils"
)

const (
	// ShortlinkKey mapping the shortlink to the url
	ShortlinkKey = "shortlink:%s:url"
	// URLHashKey mapping the hash of the url to the shortlink
	URLHashKey = "urlhash:%s:url"
	// ShortlinkDetailKey mapping the shortlink to the detail of url
	ShortlinkDetailKey = "shortlink:%s:detail"
)

type IShortLinkService interface {
	Shorten(ctx context.Context, url string, exp int32) (string, error)

	ShortLinkInfo(ctx context.Context, eid string) (*entity.ShortLinkInfo, error)

	UnShorten(ctx context.Context, eid string) (string, error)
}

type ShortLinkService struct {
	rds *rds.Infra
}

func NewShortLinkService(rdsInfra *rds.Infra) IShortLinkService {
	return &ShortLinkService{
		rds: rdsInfra,
	}
}

// Shorten convert url to shortlink
func (r *ShortLinkService) Shorten(ctx context.Context, url string, exp int32) (string, error) {
	// convert url to sha1 hash
	hv := utils.ToSha1(url)
	// fetch it if the url is cached
	d, err := r.rds.Get(ctx, fmt.Sprintf(URLHashKey, hv))
	if err == redis.Nil {
		// not exists，nothing to do
	} else if err != nil {
		return "", err
	} else {
		return d, nil
	}
	eid := generateEid(url)
	// store the url against this encoded id
	if err = r.rds.Set(ctx, fmt.Sprintf(ShortlinkKey, eid), url, time.Minute*time.Duration(exp)); err != nil {
		return "", err
	}
	// store the url against the hash of it
	if err = r.rds.Set(ctx, fmt.Sprintf(URLHashKey, hv), eid, time.Minute*time.Duration(exp)); err != nil {
		return "", err
	}
	shortLinkBytes, err := json.Marshal(&entity.ShortLinkInfo{
		URL:                 url,
		ExpirationInMinutes: exp,
		CreatedAt:           time.Now().Unix(),
	})
	if err != nil {
		return "", err
	}
	// store the url detail against this encoded id
	if err = r.rds.Set(ctx, fmt.Sprintf(ShortlinkDetailKey, eid), shortLinkBytes, time.Minute*time.Duration(exp)); err != nil {
		return "", nil
	}
	return eid, nil
}

// ShortLinkInfo returns the detail of the shortlink
func (r *ShortLinkService) ShortLinkInfo(ctx context.Context, eid string) (*entity.ShortLinkInfo, error) {
	cacheVal, err := r.rds.Get(ctx, fmt.Sprintf(ShortlinkDetailKey, eid))
	if err == redis.Nil {
		return nil, errors.New("unknown short URL")
	}
	if err != nil {
		return nil, err
	}
	shortInfo := new(entity.ShortLinkInfo)
	if err := json.Unmarshal([]byte(cacheVal), shortInfo); err != nil {
		return nil, errors.New("invalid short URL")
	}
	return shortInfo, nil
}

// UnShorten convert shortlink to url
func (r *ShortLinkService) UnShorten(ctx context.Context, eid string) (string, error) {
	url, err := r.rds.Get(ctx, fmt.Sprintf(ShortlinkKey, eid))
	if err == redis.Nil {
		return "", errors.New("unknown short URL")
	} else if err != nil {
		return "", err
	} else {
		return url, nil
	}
}

func generateEid(url string) string {
	seed := uint32(time.Now().Unix())
	for {
		v := murmur3.Sum32WithSeed([]byte(url), seed)
		str := base62.EncodeInt64(int64(v))
		if str != "" {
			return str
		}
		seed++
	}
}
