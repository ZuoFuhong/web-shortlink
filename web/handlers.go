package web

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"net/http"
	"web-shortlink/defs"
	"web-shortlink/storage"
)

var s storage.Storage

func init() {
	// todo: 读取配置文件
	s = storage.NewRedisCli("47.98.199.80:6379", "myredis", 0)
}

func createShortUrl(w http.ResponseWriter, r *http.Request) {
	req := &defs.ShortenReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		sendErrorResponse(w, defs.Response{HttpSC: http.StatusInternalServerError, Result: defs.Result{Code: "007", ErrMsg: err.Error()}})
		return
	}
	defer r.Body.Close()

	s, err := s.Shorten(req.Url, req.ExpirationInMinutes)
	if err != nil {
		sendErrorResponse(w, defs.ErrorStorageError)
		return
	}
	sendNormalResponse(w, defs.Result{Code: "0", ErrMsg: "OK", Data: defs.ShortenResp{ShortUrl: s, LongUrl: req.Url}}, http.StatusOK)
}

func getShortlinkInfo(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.URL.Query().Get("shortUrl")
	info, err := s.ShortlinkInfo(shortUrl)
	if err != nil {
		sendErrorResponse(w, defs.Response{HttpSC: http.StatusInternalServerError, Result: defs.Result{Code: "007", ErrMsg: err.Error()}})
	} else {
		urlDetail := &storage.URLDetail{}
		_ = json.Unmarshal([]byte(info.(string)), urlDetail)
		sendNormalResponse(w, defs.Result{Code: "0", ErrMsg: "OK", Data: urlDetail}, http.StatusOK)
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortUrl := vars["shortUrl"]
	longUrl, err := s.Unshorten(shortUrl)
	if err != nil {
		sendErrorResponse(w, defs.Response{HttpSC: http.StatusInternalServerError, Result: defs.Result{Code: "007", ErrMsg: err.Error()}})
	} else {
		http.Redirect(w, r, longUrl, http.StatusTemporaryRedirect)
	}
}
