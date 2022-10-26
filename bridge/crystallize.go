package bridge

import (
	"fmt"
	"gostorefront/pkg/ui"
)

type Crystallize struct {
	name string
}

func (c Crystallize) Setup() {
	fmt.Println(c.name)
}

func (c Crystallize) ProductList() {
	fmt.Println(c.name)
}

func (c Crystallize) Product() {
	fmt.Println(c.name)
}

func (c Crystallize) Cart() {
	fmt.Println(c.name)
}

func (c Crystallize) Login() {
	fmt.Println(c.name)
}

func (c Crystallize) Register() {
	fmt.Println(c.name)
}

func (c Crystallize) Guest() {
	fmt.Println(c.name)
}

func (c Crystallize) AddProduct() {
	fmt.Println(c.name)
}

func (c Crystallize) RemoveProduct() {
	fmt.Println(c.name)
}

func (c Crystallize) AddQuantity() {
	fmt.Println(c.name)
}

func (c Crystallize) RemoveQuantity() {
	fmt.Println(c.name)
}

func (c Crystallize) AddDeliveryAddress() {
	fmt.Println(c.name)
}

func (c Crystallize) DeliveryOptions() {
	fmt.Println(c.name)
}

func (c Crystallize) PaymentMethods() {
	fmt.Println(c.name)
}

func (c Crystallize) Index() IndexData {
	return IndexData{
		Title: "Product Overviews",
		Head: ui.Head{
			Lang:            "en",
			PageTitle:       "Product Overviews",
			MetaDescription: "Product Overviews Meta",
		},
	}
}

func (c Crystallize) Category() CategoryData {
	return CategoryData{Title: "Category"}
}
