package orders

import (
	"context"

	"github.com/heriant0/pos-makanan/utility"
)

type orderRepository interface {
	insertOrder(ctx context.Context, order Order) error
	findLastestOrderInvoiceId(ctx context.Context) (string, error)
	updateOrderStatus(ctx context.Context, invoice Invoice) error
}

type paymentRepository interface {
	createInvoice(ctx context.Context, invoice Invoice) (Invoice, error)
}

type service struct {
	orderRepo   orderRepository
	paymentRepo paymentRepository
}

func newService(orderRepo orderRepository, paymentRepo paymentRepository) service {
	return service{orderRepo, paymentRepo}
}

func (svc service) createOrder(ctx context.Context, req createOrderRequest) (res createOrderResponse, err error) {
	latestInvoiceId, err := svc.orderRepo.findLastestOrderInvoiceId(ctx)
	if err != nil {
		return
	}

	invoiceId := utility.GenerateInvoiceId(&latestInvoiceId)

	price := float32(50000)
	amount := float32(req.Quantity) * price
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

	err = svc.orderRepo.updateOrderStatus(ctx, invoiceRequest)
	if err != nil {
		return
	}

	return
}
