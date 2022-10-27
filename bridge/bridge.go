package bridge

type CommerceBridgeInterface interface {
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

type ContentBridgeInterface interface {
	Setup()
	HomePage()
	Blob()
}

func CommerceSetup() CommerceBridgeInterface {
	return Crystallize{name: "crystallize"}
}

func ContentSetup() ContentBridgeInterface {
	return Contentful{name: "contentful"}
}
