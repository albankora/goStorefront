package bridge

import "gostorefront/pkg/ui"

type CategoryData struct {
	Title string
}

type IndexData struct {
	Head  ui.Head
	Title string
}

type BridgeInterface interface {
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
	DeliveryOptions()
	PaymentMethods()
	Index() IndexData
	Category() CategoryData
}

func Setup() BridgeInterface {
	return Crystallize{name: "crystallize"}
}
