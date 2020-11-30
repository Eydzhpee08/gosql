package app

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/Eydzhpee08/crud/pkg/customers"

)

type Server struct {
	mux          *http.ServeMux
	customersSvc *customers.Service
	items		 *customers.Customer
}

func NewServer(mux *http.ServeMux, customersSvc *customers.Service) *Server {
	return &Server{mux: mux, customersSvc: customersSvc}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) Init() {
	s.mux.HandleFunc("/customers.getById", s.handleGetCustomerByID)
	s.mux.HandleFunc("/customers.getAll", s.handleGetCustomerAll)
	s.mux.HandleFunc("/customers.getAllActive", s.handleGetCustomerAllActive)
	s.mux.HandleFunc("/customers.save", s.handleGetCustomerSave)
	//s.mux.HandleFunc("/customers.getById", s.handleGetCustomerByID)
}

func (s *Server) handleGetCustomerByID(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return
	}

	item, err := s.customersSvc.ByID(r.Context(), id)
	if errors.Is(err, customers.ErrNotFound) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(item)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		return
	}

}

func (s *Server) handleGetCustomerAll(w http.ResponseWriter, r *http.Request) {

	item, err := s.customersSvc.All(r.Context())
	if errors.Is(err, customers.ErrNotFound) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(item)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		return
	}

}
func (s *Server) handleGetCustomerAllActive(w http.ResponseWriter, r *http.Request) {

	item, err := s.customersSvc.AllActive(r.Context())
	if errors.Is(err, customers.ErrNotFound) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(item)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		return
	}

}

func (s *Server) handleGetCustomerSave(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	phone := r.URL.Query().Get("phone")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return
	}
	itemR := &customers.Customer{
		ID: id,
		Name: name,
		Phone: phone,
		Active: strconv.FormatBool(true),
		Created: time.Date(time.Now().Year(),time.December,time.Now().Day(),time.Now().Hour(),time.Now().Minute(),time.Now().Second(),time.Now().Second(),time.UTC),
	}

	item, err := s.customersSvc.Save(r.Context(), itemR)
	if errors.Is(err, customers.ErrNotFound) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(item)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		return
	}

}