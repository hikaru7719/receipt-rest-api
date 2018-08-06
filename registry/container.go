package registry

import "github.com/hikaru7719/receipt-rest-api/interface/server/handler"

type Registry interface {
	NewAppHandler() handler.GetReceiptHandler
}

type Container struct{
}

func (c *Container) NewAppHandler() handler.GetReceiptHandler{
	return handler.NewGetReceiptHandler()
}