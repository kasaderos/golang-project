package controller_http

import (
	"encoding/json"
	"net/http"
	"route256/cart/internal/models"
	"route256/cart/internal/services/product"
)

type ItemAddRequest struct {
	User  int64  `json:"user,omitempty"`
	SKU   uint32 `json:"sku,omitempty"`
	Count uint16 `json:"count,omitempty"`
}

func (c *Controller) ItemAddHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	productServiceToken := r.Header.Get("X-Product-Service-Token")
	if len(productServiceToken) < 1 {
		http.Error(w, "empty product service token", http.StatusBadRequest)
		return
	}

	var req ItemAddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := req.validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.Usecases.AddItem(
		product.WithToken(ctx, productServiceToken),
		models.UserID(req.User),
		models.SKU(req.SKU),
		req.Count,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}
}

func (req *ItemAddRequest) validate() error {
	return nil
}
