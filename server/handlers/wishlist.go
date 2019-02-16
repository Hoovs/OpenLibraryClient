package handlers

import (
	"go.uber.org/zap"
	"net/http"
)

type WishListHandler struct {
	Logger *zap.Logger
}

func (wh *WishListHandler) GetWishListHandler(w http.ResponseWriter, r *http.Request) {
	wh.Logger.Info("WishListHandler for get called")
}

func (wh *WishListHandler) DeleteWishListHandler(w http.ResponseWriter, r *http.Request) {
	wh.Logger.Info("WishListHandler for delete called")
}
