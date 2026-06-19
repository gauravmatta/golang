package servemux

import (
	"customerManager/domain"
	"encoding/json"
	"net/http"
)

type CustomerHandler struct {
	Repository domain.CustomerRepository // interface for persistence
}

func (h *CustomerHandler) Post(w http.ResponseWriter, r *http.Request) {
	var customer domain.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, err = h.Repository.FindById(customer.ID); err == nil {
		http.Error(w, "Customer already exists", http.StatusConflict)
		return
	}
	if err = h.Repository.Create(customer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *CustomerHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	customers, err := h.Repository.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		marshal, err := json.Marshal(customers)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(marshal)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

}
