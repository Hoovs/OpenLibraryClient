package handlers

import (
	"errors"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
)

// SearchHandler wraps calling BaseSearchUrl and returns the body back
type SearchHandler struct {
	Logger        *zap.Logger
	BaseSearchUrl string
}

// SearchHandler handles a GET request and calls s.BaseSearchUrl with the passed in
// query parameters
func (s *SearchHandler) SearchHandler(w http.ResponseWriter, r *http.Request) {
	getUrl, err := s.createLibraryURL(r.URL.Query())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	s.Logger.Info("Calling remote client", zap.String("url", getUrl))

	req, err := http.NewRequest("GET", getUrl, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	defer resp.Body.Close()

	v, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to read body"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(v)
	return
}

// Takes in the query string from the request and url encodes it for calling.
func (s *SearchHandler) createLibraryURL(v url.Values) (string, error) {
	q := v.Get("q")
	if q == "" {
		return "", errors.New("must specify a q= parameter")
	}

	s.Logger.Info("Library URL", zap.String("q", q))
	u := &url.URL{Path: s.BaseSearchUrl + q}
	return u.String()[2:], nil
}
