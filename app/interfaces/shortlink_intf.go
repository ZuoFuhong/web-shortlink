package interfaces

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"web-shortlink/errcode"
	"web-shortlink/pkg/log"
)

// CreateShortUrl 创建短地址
func (s *ShortLinkServiceImpl) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	req := new(ShortenReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	if !req.CheckRequiredParam() {
		Error(w, errcode.BadRequestParam, "无效的参数")
		return
	}
	shortUrl, err := s.short.Shorten(r.Context(), req.Url, req.ExpirationInMinutes)
	if err != nil {
		Error(w, errcode.InternalServerError, err.Error())
		return
	}
	resp := ShortenResp{
		LongUrl:  req.Url,
		ShortUrl: shortUrl,
	}
	Ok(w, resp)
}

// GetShortLinkInfo 查询短地址
func (s *ShortLinkServiceImpl) GetShortLinkInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eid := vars["eid"]
	log.InfoContextf(r.Context(), "eid: %s", eid)
	shortInfo, err := s.short.ShortLinkInfo(r.Context(), eid)
	if err != nil {
		Error(w, errcode.InternalServerError, err.Error())
		return
	}
	Ok(w, shortInfo)
}

// Redirect 307 重定向
func (s *ShortLinkServiceImpl) Redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eid := vars["eid"]
	log.InfoContextf(r.Context(), "eid: %s", eid)
	originUrl, err := s.short.UnShorten(r.Context(), eid)
	if err != nil {
		Error(w, errcode.InternalServerError, err.Error())
		return
	}
	http.Redirect(w, r, originUrl, http.StatusTemporaryRedirect)
}
