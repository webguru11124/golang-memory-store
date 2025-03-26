package api

import (
	"encoding/json"
	"golang-memory-store/internal/core"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	store *core.Store
}

func NewHandler(store *core.Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) Set(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
		TTL   int         `json:"ttl"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	h.store.Set(req.Key, req.Value, req.TTL)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	value, found := h.store.Get(key)
	if !found {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(value)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	h.store.Delete(key)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Push(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	list := h.store.GetList(req.Key)
	list.Push(req.Value)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Pop(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	list := h.store.GetList(key)
	value, found := list.Pop()
	if !found {
		http.Error(w, "No items in list", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(value)
}
