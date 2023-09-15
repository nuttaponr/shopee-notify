package shopee

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Client struct {
	ShopID    string
	ItemID    string
	ModelID   int64
	client    *http.Client
	userAgent string
}

func New(shopID, itemID string, modelID int64, opts ...Option) *Client {
	c := &Client{
		ShopID:  shopID,
		ItemID:  itemID,
		ModelID: modelID,
		client:  http.DefaultClient,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

type Option func(*Client)

// WithUserAgent when need to change the user agent
func WithUserAgent(userAgent string) Option {
	return func(c *Client) {
		c.userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36"
		if userAgent != "" {
			c.userAgent = userAgent
		}

	}
}

func (c *Client) Call() (string, error) {
	url := fmt.Sprintf("https://shopee.co.th/api/v4/item/get?itemid=%s&shopid=%s", c.ItemID, c.ShopID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("User-Agent", c.userAgent)
	if err != nil {
		return "", err
	}

	res, err := c.client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return "", err
	}

	var result Shopee
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return "", err
	}

	for _, m := range result.Data.Models {
		msg := fmt.Sprintf("%s: %s มีของแล้ว!! จำนวน %v ชิ้น ", result.Data.Name, m.Name, m.NormalStock)
		log.Println(msg)
		if m.ModelId == c.ModelID && m.NormalStock > 0 {
			return msg + fmt.Sprintf("https://shopee.co.th/product/%s/%s", c.ShopID, c.ItemID), nil

		}
	}
	return "", nil
}
