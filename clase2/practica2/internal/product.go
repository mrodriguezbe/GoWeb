package internal

type Product struct {
	Name       string  `json:"name"`
	Id         int     `json:"id"`
	Quantity   int     `json:"quantity"`
	CodeValue  string  `json:"code_value"`
	Published  bool    `json:"is_published"`
	Expiration string  `json:"expiration"`
	Price      float64 `json:"price"`
}
