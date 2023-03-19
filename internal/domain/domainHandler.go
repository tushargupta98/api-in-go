package domain

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type DomainHandler struct {
	repo DomainRepository
}

func NewDomainHandler(repo DomainRepository) *DomainHandler {
	return &DomainHandler{repo}
}

// List godoc
// @Summary List domains
// @Description Get a list of all domains
// @Tags Domain
// @Security ApiKeyAuth
// @ID list-domain
// @Success 200 {array} Domain
// @Failure 500 {string} string "Internal server error"
// @Router /domain [get]
func (h *DomainHandler) List(w http.ResponseWriter, r *http.Request) {
	domain, err := h.repo.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(domain)
}

// Create godoc
// @Summary Create a domain
// @Description Create a new domain
// @Tags Domain
// @Security ApiKeyAuth
// @ID create-domain
// @Accept json
// @Produce json
// @Param date_range body Domain true "Domain object"
// @Success 201 {string} string "Created"
// @Header 201 {string} Location "/domain/{id}"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /domain [post]
func (h *DomainHandler) Create(w http.ResponseWriter, r *http.Request) {
	var domain Domain
	err := json.NewDecoder(r.Body).Decode(&domain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := h.repo.Create(domain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	location := r.URL.Path + "/" + strconv.Itoa(id)
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
}

// Get godoc
// @Summary Get a domain
// @Description Get a domain by ID
// @Tags Domain
// @Security ApiKeyAuth
// @ID get-domain
// @Param id path int true "Domain ID"
// @Success 200 {object} Domain
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Domain not found"
// @Router /domain/{id} [get]
func (h *DomainHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	domain, err := h.repo.Get(id)
	if err != nil {
		http.Error(w, "Domain not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(domain)
}

// Update godoc
// @Summary Update a domain
// @Description Update a domain by ID
// @Tags Domain
// @Security ApiKeyAuth
// @ID update-domain
// @Accept json
// @Produce json
// @Param id path int true "Domain ID"
// @Param date_range body Domain true "Domain object"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Invalid ID or bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /domain/{id} [put]
func (h *DomainHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var domain Domain
	err = json.NewDecoder(r.Body).Decode(&domain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.repo.Update(id, domain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Delete godoc
// @Summary Delete a domain
// @Description Delete a domain by ID
// @Tags Domain
// @Security ApiKeyAuth
// @ID delete-domain
// @Param id path int true "Domain ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Invalid ID"
// @Failure 500 {string} string "Internal server error"
// @Router /domain/{id} [delete]
func (h *DomainHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	err = h.repo.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func parseIDParam(r *http.Request) (int, error) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}
	return id, nil
}
