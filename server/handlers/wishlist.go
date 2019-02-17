package handlers

import (
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

func (wh *WishListHandler) GetWishListHandler(w http.ResponseWriter, r *http.Request) {
	wh.Logger.Info("WishListHandler for get called")
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
