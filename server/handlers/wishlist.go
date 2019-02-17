package handlers

import (
	"encoding/json"
	"github.com/Hoovs/OpenLibraryClient/server/db"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

type WishListHandler struct {
	Logger *zap.Logger
	Db     *db.DB
}

func (wh *WishListHandler) PostWishListHandler(w http.ResponseWriter, r *http.Request) {
	wh.Logger.Info("WishListHandler for post called")
	bodyRow := &db.WishListRow{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		write(w, http.StatusBadRequest, []byte(err.Error()), wh.Logger)
		return
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			wh.Logger.Error(err.Error())
		}
	}()

	err = json.Unmarshal(body, bodyRow)
	if err != nil {
		write(w, http.StatusBadRequest, []byte(err.Error()), wh.Logger)
		return
	}

	err = wh.Db.InsertRow(*bodyRow)
	if err != nil {
		write(w, http.StatusBadRequest, []byte(err.Error()), wh.Logger)
		return
	}
}

func (wh *WishListHandler) GetWishListHandler(w http.ResponseWriter, r *http.Request) {
	wh.Logger.Info("WishListHandler for get called")
	vars := mux.Vars(r)
	idStr := vars["wishListId"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		write(w, http.StatusBadRequest, []byte("Unable to parse wish list id from request"), wh.Logger)
		return
	}

	row, err := wh.Db.GetWishList(id)
	if err != nil {
		write(w, http.StatusBadRequest, []byte("Unable to fetch row"), wh.Logger)
		return
	}

	req, err := http.NewRequest("GET", "http://localhost:8080/search?q="+row.BookTitle, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		write(w, http.StatusBadRequest, []byte(err.Error()), wh.Logger)
		return
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			wh.Logger.Error(err.Error())
		}
	}()

	v, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		write(w, http.StatusInternalServerError, []byte("unable to read body"), wh.Logger)
		return
	}
	write(w, http.StatusOK, v, wh.Logger)
	b, err := json.Marshal(row)
	w.Write(b)
	return
}

func (wh *WishListHandler) DeleteWishListHandler(w http.ResponseWriter, r *http.Request) {
	wh.Logger.Info("WishListHandler for delete called")
	vars := mux.Vars(r)
	idStr := vars["wishListId"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		write(w, http.StatusBadRequest, []byte("Unable to parse wish list id from request"), wh.Logger)
		return
	}

	if err := wh.Db.DeleteWishList(id); err != nil {
		wh.Logger.Error(err.Error())
		write(w, http.StatusBadRequest, []byte("Unable to delete from wish list"), wh.Logger)
		return
	}
	write(w, http.StatusNoContent, nil, wh.Logger)
	return
}
