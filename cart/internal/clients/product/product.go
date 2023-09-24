package product

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"route256/cart/internal/models"
	"route256/cart/internal/services/cart"
)

type contextKeyType string

const tokenContextKey = contextKeyType("token")

const (
	getProductInfoPath = "/get_product"
)

type ProductService struct {
	name       string
	baseURL    string
	httpClient *http.Client
}

var _ cart.ProductProvider = (*ProductService)(nil)

func NewProductService(baseURL string) *ProductService {
	return &ProductService{
		name:       "product service",
		httpClient: &http.Client{},
		baseURL:    baseURL,
	}
}

func WithToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenContextKey, token)
}

func getTokenFromContext(ctx context.Context) string {
	val, ok := ctx.Value(tokenContextKey).(string)
	if !ok || len(val) < 1 {
		log.Println("bad token")
	}
	return val
}

func (c *ProductService) GetProductInfo(ctx context.Context, sku models.SKU) (name string, price uint32, err error) {
	reqBody := GetProductRequest{
		Token: getTokenFromContext(ctx),
		SKU:   uint32(sku),
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		return "", 0, fmt.Errorf("%s: failed to encode request %w", c.name, err)
	}

	reqURL, err := url.JoinPath(c.baseURL, getProductInfoPath)
	if err != nil {
		return "", 0, fmt.Errorf("%s: join path %w, base '%s' path '%s'", c.name, err, c.baseURL, getProductInfoPath)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewBuffer(data))
	if err != nil {
		return "", 0, fmt.Errorf("%s: failed to create HTTP request: %w", c.name, err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", 0, fmt.Errorf("%s: failed to execute HTTP request: %w", c.name, err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		var response GetProductErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			return "", 0, fmt.Errorf("%s: failed to decode error response: %w", c.name, err)
		}
		return "", 0, fmt.Errorf("%s: HTTP request responded with: %d , message: %s", c.name, resp.StatusCode, response.Message)
	}

	var response GetProductResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", 0, fmt.Errorf("%s: failed to decode error response: %w", c.name, err)
	}

	return response.Name, response.Price, nil
}