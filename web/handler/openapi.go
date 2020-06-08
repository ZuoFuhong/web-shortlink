package handler

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"net/http"
	"web-shortlink/config"
	"web-shortlink/internal/defs"
	"web-shortlink/internal/errs"
	"web-shortlink/internal/storage"
)

type openAPI struct {
	storage storage.Storage
}

var OpenAPI = new(openAPI)

func init() {
	conf := config.LoadConf()
	OpenAPI.storage = storage.NewRedisCli(conf.Redis.Addr, conf.Redis.Pwd, conf.Redis.Db)
}

// 创建短地址
func (h *openAPI) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	req := new(defs.ShortenReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		panic(errs.ParameterError)
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		panic(errs.ParameterError)
	}

	url, err := h.storage.Shorten(req.Url, req.ExpirationInMinutes)
	if err != nil {
		panic(err)
	}

	resp := defs.ShortenResp{ShortUrl: url, LongUrl: req.Url}
	bytes, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	_, _ = w.Write(bytes)
}

// 查询短地址
func (h *openAPI) GetShortlinkInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortUrl := vars["url"]
	info, err := h.storage.ShortlinkInfo(shortUrl)
	if err != nil {
		panic(errs.NewHttpErr(errs.OpenAPI, err.Error()))
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	_, _ = w.Write([]byte(info))
}

// 307重定向
func (h *openAPI) Redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortUrl := vars["url"]
	longUrl, err := h.storage.Unshorten(shortUrl)
	if err != nil {
		panic(errs.NewHttpErr(errs.OpenAPI, err.Error()))
	}
	http.Redirect(w, r, longUrl, http.StatusTemporaryRedirect)
}
