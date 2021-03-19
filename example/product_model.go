package main

// swagger:model ProductModel
type ProductModel struct {
	Name string
}

func NewProductModel(name string) *ProductModel {
	return &ProductModel{Name: name}
}
