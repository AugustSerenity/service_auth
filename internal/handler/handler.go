package handler

import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/AugustSerenity/service_auth/internal/handler/model"
	"github.com/beevik/guid"
)

type Handler struct {
	service Service
}

func New(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) Route() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("POST /auth-token", h.CreateToken)
	router.HandleFunc("POST /auth-refresh", h.RefreshToken)

	return router
}

func (h *Handler) CreateToken(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "Empty data", http.StatusBadRequest)
		return
	}

	if !guid.IsGuid(userID) {
		http.Error(w, "invalid guid", http.StatusBadRequest)
		return
	}

	userIP := getClientIP(r)

	accessToken, refreshToken, err := h.service.CreateToken(r.Context(), userID, userIP)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	resp := model.CreateTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}

func getClientIP(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req model.RefreshTokenRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.AccessToken == "" || req.RefreshToken == "" {
		http.Error(w, "Missing tokens", http.StatusBadRequest)
		return
	}

	clientIP := getClientIP(r)

	newAccess, newRefresh, err := h.service.RefreshToken(r.Context(), req.AccessToken, req.RefreshToken, clientIP)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	resp := model.CreateTokenResponse{
		AccessToken:  newAccess,
		RefreshToken: newRefresh,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
