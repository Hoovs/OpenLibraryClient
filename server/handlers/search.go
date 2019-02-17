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

// Internal helper function to catch error from writing back to user.
func write(w http.ResponseWriter, header int, body []byte, l *zap.Logger) {
	w.WriteHeader(header)
	if _, err := w.Write(body); err != nil {
		l.Error(err.Error())
	}
}

// SearchHandler handles a GET request and calls s.BaseSearchUrl with the passed in
// query parameters.
func (s *SearchHandler) SearchHandler(w http.ResponseWriter, r *http.Request) {
	getUrl, err := s.createLibraryURL(r.URL.Query())
	if err != nil {
		write(w, http.StatusBadRequest, []byte(err.Error()), s.Logger)
		return
	}
	s.Logger.Info("Calling remote client", zap.String("url", getUrl))

	req, err := http.NewRequest("GET", getUrl, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		write(w, http.StatusBadRequest, []byte(err.Error()), s.Logger)
		return
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			s.Logger.Error(err.Error())
		}
	}()

	v, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		write(w, http.StatusInternalServerError, []byte("unable to read body"), s.Logger)
		return
	}
	write(w, http.StatusOK, v, s.Logger)
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
