package models

type Product struct {
	ID    int     `json:"id,omitempty"`
	UID   string  `json:"uid,omitempty"`
	Name  string  `json:"name,omitempty"`
	Type  string  `json:"type,omitempty"`
	Count int     `json:"count,omitempty"`
	Price float64 `json:"price,omitempty"`
}
