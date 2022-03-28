package handler

import (
	"net/http"

	"github.com/booogaga/url-cropper/internal/storage"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func Redirect(stor storage.CropURLStorage, slogger *zap.SugaredLogger) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// define shortID from users query
		shortID := chi.URLParam(r, "shortID")
		// defint corresponding long URL from database
		longURL, err := stor.Resolve(shortID)
		if err != nil {
			slogger.Debugf("resolving error %w", err)
			JSONError(rw, err, http.StatusBadRequest)
			return
		}
		slogger.Debugf("successful redirect %s -> %s", shortID, longURL)
		// implement redirect
		http.Redirect(rw, r, longURL, http.StatusPermanentRedirect)
	}
}
