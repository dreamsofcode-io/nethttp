package invoice

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type Handler struct{}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Invoice created!"))
}

func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	monster, exists := loadInvoices()[r.PathValue("id")]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(monster)
}

func (h *Handler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Invoice updated!"))
}

func (h *Handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Invoice deleted!"))
}

func (h *Handler) PatchByID(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
}

type Invoice struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Amount    uint      `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

var invoices []Invoice = []Invoice{
	{ID: 1, Name: "Invoice #1", Amount: 100, Timestamp: time.Now().Add(-24 * time.Hour)},
	{ID: 2, Name: "Invoice #2", Amount: 150, Timestamp: time.Now().Add(-48 * time.Hour)},
	{ID: 3, Name: "Invoice #3", Amount: 200, Timestamp: time.Now().Add(-72 * time.Hour)},
	{ID: 4, Name: "Invoice #4", Amount: 75, Timestamp: time.Now().Add(-96 * time.Hour)},
	{ID: 5, Name: "Invoice #5", Amount: 300, Timestamp: time.Now().Add(-120 * time.Hour)},
}

func loadInvoices() map[string]Invoice {
	res := make(map[string]Invoice, len(invoices))

	for _, x := range invoices {
		res[strconv.Itoa(int(x.ID))] = x
	}

	return res
}
