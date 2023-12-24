package orders

import (
	"context"

	"github.com/heriant0/pos-makanan/utility"
)

type orderRepository interface {
	insertOrder(ctx context.Context, order Order) error
	findLastestOrderInvoiceId(ctx context.Context) (string, error)
	updateOrderStatus(ctx context.Context, invoice Invoice) error
	findOneByInvoiceId(ctx context.Context, invoiceId string) (Order, error)
}

type paymentRepository interface {
	createInvoice(ctx context.Context, invoice Invoice) (Invoice, error)
}

type productRepository interface {
	findProductById(ctx context.Context, id int) (Product, error)
	updateProductStockById(ctx context.Context, id, stock int) error
}

type service struct {
	orderRepo   orderRepository
	paymentRepo paymentRepository
	productRepo productRepository
}

func newService(orderRepo orderRepository, paymentRepo paymentRepository, productRepo productRepository) service {
	return service{
		orderRepo:   orderRepo,
		paymentRepo: paymentRepo,
		productRepo: productRepo,
	}
}

func (svc service) createOrder(ctx context.Context, req createOrderRequest) (res createOrderResponse, err error) {
	latestInvoiceId, err := svc.orderRepo.findLastestOrderInvoiceId(ctx)
	if err != nil {
		return
	}

	invoiceId := utility.GenerateInvoiceId(&latestInvoiceId)

	product, err := svc.productRepo.findProductById(ctx, req.ProductId)
	if err != nil {
		return
	}

	amount := float32(req.Quantity) * float32(product.Price)
	invoice := Invoice{
		ExternalId: invoiceId,
		Amount:     amount,
	}

	invoiceResult, err := svc.paymentRepo.createInvoice(ctx, invoice)
	if err != nil {
		return
	}

	orderInsert := req.toOrderEntity()
	orderInsert.InvoiceId = invoiceId
	orderInsert.InvoiceUrl = invoiceResult.InvoiceUrl
	orderInsert.Amount = amount
	orderInsert.Status = Pending.value()
	orderInsert.Product = ProductRecord{
		Id:    req.ProductId,
		Price: float32(product.Price),
		Name:  product.Name,
	}

	err = svc.orderRepo.insertOrder(ctx, orderInsert)
	if err != nil {
		return
	}

	res = createOrderResponse{
		InvoiceUrl: invoiceResult.InvoiceUrl,
	}
	return
}

func (svc service) changeOrderStatus(ctx context.Context, req updateOrderRequest) (err error) {
	invoiceRequest := req.toInvoiceEntity()

	order, err := svc.orderRepo.findOneByInvoiceId(ctx, req.ExternalID)
	if err != nil {
		return
	}

	err = svc.productRepo.updateProductStockById(ctx, order.Product.Id, order.Quantity)
	if err != nil {
		return
	}

	err = svc.orderRepo.updateOrderStatus(ctx, invoiceRequest)
	if err != nil {
		return
	}

	return
}
