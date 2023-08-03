package statuses

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *handler) GetStatusByID(w http.ResponseWriter, r *http.Request) {
	ids := chi.URLParam(r, "id")
	if ids == "" {
		http.Error(w, "no status id", http.StatusInternalServerError)
		return
	}

	id, err := strconv.ParseInt(ids, 10, 64)

	if err != nil {
		http.Error(w, "invalid status ID", http.StatusBadRequest)
		return
	}

	ctx := context.TODO()
	status, err := h.sr.FindByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if status == nil {
		http.Error(w, "status not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
