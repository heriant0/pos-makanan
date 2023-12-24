package paymentgateway

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/heriant0/pos-makanan/internal/config"
	"github.com/xendit/xendit-go/v3"
	"github.com/xendit/xendit-go/v3/invoice"
)

type Xendit struct {
	client    *xendit.APIClient
	secretKey string
	redirect  redirect
}

type redirect struct {
	Success string `json:"success_redirect_url"`
	Failure string `json:"failure_redirect_url"`
}

func NewXendit(secretKey string) Xendit {
	client := xendit.NewClient(secretKey)

	return Xendit{client: client, secretKey: secretKey}
}

func (x *Xendit) SetConfig(cfg config.Payment) *Xendit {
	x.redirect.Success = cfg.RedirectUrl.Success
	x.redirect.Failure = cfg.RedirectUrl.Failure

	return x
}

type InvoiceRequest struct {
	ExternalId string
	Amount     float32
}

func (i InvoiceRequest) toXenditRequest(successRedirectURI, failureRedirectURI string) invoice.CreateInvoiceRequest {
	req := invoice.NewCreateInvoiceRequestWithDefaults()
	req.ExternalId = i.ExternalId
	req.Amount = i.Amount

	return *req
}

type InvoiceResponse struct {
	InvoiceUrl string
}

func newInvoiceResponseFromXenditResponse(invoice invoice.Invoice) (res InvoiceResponse) {
	res = InvoiceResponse{
		InvoiceUrl: invoice.InvoiceUrl,
	}
	return
}

func (x Xendit) CreateInvoice(ctx context.Context, req InvoiceRequest) (res InvoiceResponse, err error) {
	invoiceReq := req.toXenditRequest(x.redirect.Success, x.redirect.Failure)
	invoice, httpResponse, errXendit := x.client.InvoiceApi.CreateInvoice(ctx).CreateInvoiceRequest(invoiceReq).Execute()

	if errXendit != nil {
		b, _ := json.Marshal(errXendit.FullError())
		fmt.Printf("Error when try to get balance with error detail : %v\n", string(b))
		fmt.Printf("Full HTTP response: %v\n", httpResponse)
		err = errXendit
		return
	}

	// jika invoice nya nil
	if invoice == nil {
		err = errors.New("invoice xendit is nil")
		fmt.Printf("Full HTTP response: %v\n", httpResponse)
		fmt.Printf("Error when try to get balance with error detail : %v\n", err.Error())
		return
	}

	res = newInvoiceResponseFromXenditResponse(*invoice)
	return
}
