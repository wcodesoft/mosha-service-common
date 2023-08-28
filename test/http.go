package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
)

func ExecuteRequest(req *http.Request, handler http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

func JsonReaderFactory(in interface{}) io.Reader {
	buf := bytes.NewBuffer(nil)
	_ = json.NewEncoder(buf).Encode(in)
	return buf
}
