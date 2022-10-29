package connectors

import "fmt"

type ConnectorInterface interface {
	Setup()
	ProductList()
	Product()
	Cart()
	Login()
	Register()
	Guest()
	AddProduct()
	RemoveProduct()
	AddQuantity()
	RemoveQuantity()
	AddDeliveryAddress()
	GetDeliveryOptions()
	AddDeliveryOptions()
	GetPaymentMethods()
	Category()
}

type Connector struct {
	Page string
}

func (c Connector) ProductList() {
	fmt.Println(c.Page)
}

func (c Connector) Product() {
	fmt.Println(c.Page)
}

func (c Connector) Cart() {
	fmt.Println(c.Page)
}

func (c Connector) Login() {
	fmt.Println(c.Page)
}

func (c Connector) Register() {
	fmt.Println(c.Page)
}

func (c Connector) Guest() {
	fmt.Println(c.Page)
}

func (c Connector) AddProduct() {
	fmt.Println(c.Page)
}

func (c Connector) RemoveProduct() {
	fmt.Println(c.Page)
}

func (c Connector) AddQuantity() {
	fmt.Println(c.Page)
}

func (c Connector) RemoveQuantity() {
	fmt.Println(c.Page)
}

func (c Connector) AddDeliveryAddress() {
	fmt.Println(c.Page)
}

func (c Connector) DeliveryOptions() {
	fmt.Println(c.Page)
}

func (c Connector) PaymentMethods() {
	fmt.Println(c.Page)
}

func (c Connector) Category() {
	fmt.Println(c.Page)
}

func (c Connector) AddDeliveryOptions() {
	fmt.Println(c.Page)
}

func (c Connector) GetDeliveryOptions() {
	fmt.Println(c.Page)
}

func (c Connector) GetPaymentMethods() {
	fmt.Println(c.Page)
}

func (c Connector) HomePage() {
	fmt.Println(c.Page)
}

func (c Connector) Blog() {
	fmt.Println(c.Page)
}
