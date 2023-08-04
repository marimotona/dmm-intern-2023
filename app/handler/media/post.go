package media

import (
	"context"
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/domain/object"
)

type AddMwdiaRequest struct {
	MediaID  int64  `json:"media_id"`
	MediaURL string `json:"media_url"`
}

//// 質問 ////
// swagger UI で画像は `Media to be uploaded (encoded using multipart/form-data)` となっているが、どうやってURLとして扱うのか？

func (h *handler) PostMedia(w http.ResponseWriter, r *http.Request) {
	var req AddMwdiaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	media := &object.Media{
		MediaID:  req.MediaID,
		MediaURL: req.MediaURL,
	}

	ctx := context.TODO()
	mediaID, err := h.mr.AddMedia(ctx, media)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	media.MediaID = mediaID

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(media); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
