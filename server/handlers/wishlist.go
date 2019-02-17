package handlers

import (
	"encoding/json"
	"github.com/Hoovs/OpenLibraryClient/server/db"
	"github.com/gorilla/mux"
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
}

func (wh *WishListHandler) GetWishListHandler(w http.ResponseWriter, r *http.Request) {
	wh.Logger.Info("WishListHandler for get called")
	vars := mux.Vars(r)
	idStr := vars["wishListId"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to parse wish list id from request"))
		return
	}

	row, err := wh.Db.GetWishList(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to fetch row"))
		return
	}
	b, err := json.Marshal(row)
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

func (wh *WishListHandler) DeleteWishListHandler(w http.ResponseWriter, r *http.Request) {
	wh.Logger.Info("WishListHandler for delete called")
	vars := mux.Vars(r)
	idStr := vars["wishListId"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to parse wish list id from request"))
		return
	}

	if err := wh.Db.DeleteWishList(id); err != nil {
		wh.Logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to delete from wish list"))
		return
	}
}
