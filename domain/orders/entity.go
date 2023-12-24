package orders

import (
	"time"

	"github.com/heriant0/pos-makanan/domain/products"
	paymentgateway "github.com/heriant0/pos-makanan/external/payment-gateway"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Id         primitive.ObjectID `bson:"_id"`
	InvoiceId  string             `bson:"invoice_id"`
	Product    products.Product   `bson:"product"`
	Quantity   int                `bson:"quantity"`
	Amount     float32            `bson:"amount"`
	InvoiceUrl string             `bson:"invoice_url"`
	Status     string             `bson:"status"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
	DeletedAt  *time.Time         `bson:"deleted_at"`
}

type OrderStatus string

func (o OrderStatus) value() string {
	return string(o)
}

const (
	Paid    OrderStatus = "Paid"
	Pending OrderStatus = "Pending"
	Expired OrderStatus = "Expired"
)

type Invoice struct {
	ExternalId string  `json:"external_id"`
	Amount     float32 `json:"amount"`
	Status     OrderStatus
	InvoiceUrl string
}

func (i Invoice) toInvoiceRequest() paymentgateway.InvoiceRequest {
	return paymentgateway.InvoiceRequest{
		Amount:     i.Amount,
		ExternalId: i.ExternalId,
	}
}
