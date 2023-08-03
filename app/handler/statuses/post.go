package statuses

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/auth"
)

type AddRequest struct {
	AccountID int64  `json:"account_id"`
	Content   string `json:"content"`
}

// responseの構造体
type StatusResponse struct {
	StatusID int64 `json:"status_id"`
}

func (h *handler) PostStatus(w http.ResponseWriter, r *http.Request) {

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account := auth.AccountOf(r)
	fmt.Println(`😎😎😎`)
	fmt.Println(req.Content)
	status := &object.Status{
		Account: account,
		Content: req.Content,
	}

	ctx := context.TODO()
	statusID, err := h.sr.AddStatus(ctx, account, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//新しく作成したstatusのIDを返す
	status.ID = statusID

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
