package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/wesmota/go-jobsity-chat-server/handlers/presenter"
	"github.com/wesmota/go-jobsity-chat-server/models"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Info().Msgf("Error decoding: %v", err)
		ErrResponse(ErrInRequestMarshaling, w)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Info().Interface("user", user).Msg("SignUp")
	login, err := h.AuthService.SignUp(context.Background(), user)
	if err != nil {
		log.Info().Msgf("Error signing up: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(login)
	w.WriteHeader(http.StatusOK)

}
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	payload := presenter.Login{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Info().Msgf("Error decoding: %v", err)
		ErrResponse(ErrInRequestMarshaling, w)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	loginData, err := h.AuthService.Login(context.Background(), payload.Email, payload.Password)
	if err != nil {
		log.Info().Msgf("Error logging in: %v", err)
		ErrResponse(errors.New("invalid credentials"), w)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(loginData)
	w.WriteHeader(http.StatusOK)
	return
}
