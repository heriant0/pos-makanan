package orders

import (
	"context"

	paymentgateway "github.com/heriant0/pos-makanan/external/payment-gateway"
)

type PaymentAdapter interface {
	CreateInvoice(ctx context.Context, req paymentgateway.InvoiceRequest) (paymentgateway.InvoiceResponse, error)
}

type paymentRepo struct {
	xendit PaymentAdapter
}

func newPaymentRepo(adapter PaymentAdapter) paymentRepository {
	return paymentRepo{adapter}
}

// createInvoice implements paymentRepository.
func (p paymentRepo) createInvoice(ctx context.Context, invoice Invoice) (res Invoice, err error) {
	req := invoice.toInvoiceRequest()

	result, err := p.xendit.CreateInvoice(ctx, req)

	res = Invoice{
		InvoiceUrl: result.InvoiceUrl,
	}

	return
}
