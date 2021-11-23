package handler

import (
	"encoding/json"
	"github.com/teten-nugraha/go-social-network/internal/service"
	"net/http"
)

type loginInput struct {
	Email string
}

func (h *handler) login (w http.ResponseWriter, r *http.Request) {

	var in loginInput
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out, err := h.service.Login(r.Context(), in.Email)
	if err == service.ErrInvalidEmail {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	if err == service.ErrUserNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		RespondError(w, err)
		return
	}

	Respond(w, out, http.StatusOK)
}