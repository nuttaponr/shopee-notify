package shopee

type Shopee struct {
	Error    interface{} `json:"error"`
	ErrorMsg interface{} `json:"error_msg"`
	Data     struct {
		Name       string `json:"name"`
		ItemStatus string `json:"item_status"`
		Models     []struct {
			Name        string `json:"name"`
			Stock       int    `json:"stock"`
			NormalStock int    `json:"normal_stock"`
			ModelId     int64  `json:"modelid"`
		} `json:"models"`
	} `json:"data"`
}
