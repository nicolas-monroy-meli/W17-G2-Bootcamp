package handler

import (
	"net/http"

	internal "github.com/nicolas-monroy-meli/W17-G2-Bootcamp/internal/interfaces"
)

// NewSectionHandler creates a new instance of the section handler
func NewSectionHandler(sv internal.SectionService) *SectionHandler {
	return &SectionHandler{
		sv: sv,
	}
}

// SectionHandler is the default implementation of the section handler
type SectionHandler struct {
	// sv is the service used by the handler
	sv internal.SectionService
}

// GetAll returns all sections
func (h *SectionHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// GetByID returns a section
func (h *SectionHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Create creates a new section
func (h *SectionHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Update updates a section
func (h *SectionHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Delete deletes a section
func (h *SectionHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
