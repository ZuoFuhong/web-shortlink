package interfaces

import (
	"net/http"
	"web-shortlink/app/domain/service"
	"web-shortlink/app/infra/rds"
)

type ShortLinkServiceImpl struct {
	short service.IShortLinkService
}

func InitializeService() *ShortLinkServiceImpl {
	rdsInfra := rds.NewRedisInfra()
	shortLinkService := service.NewShortLinkService(rdsInfra)
	return &ShortLinkServiceImpl{
		short: shortLinkService,
	}
}

func (s *ShortLinkServiceImpl) Ping(w http.ResponseWriter, r *http.Request) {
	Ok(w, "ok")
}
