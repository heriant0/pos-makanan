package orders

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repository struct {
	dbMongo *mongo.Client
}

func newRespository(dbMongo *mongo.Client) repository {
	return repository{dbMongo}
}

func (r repository) insertOrder(ctx context.Context, order Order) (err error) {
	collection := r.dbMongo.Database(mongoDatabase).Collection(mongoOrderCollection)

	order.Id = primitive.NewObjectID()
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	order.DeletedAt = nil

	_, err = collection.InsertOne(ctx, order)
	if err != nil {
		return
	}

	return
}

func (r repository) findLastestOrderInvoiceId(ctx context.Context) (invoiceId string, err error) {
	collection := r.dbMongo.Database(mongoDatabase).Collection(mongoOrderCollection)
	options := options.FindOne().SetSort(bson.M{"invoice_id": -1})

	lastOrder := Order{}
	err = collection.FindOne(ctx, bson.M{}, options).Decode(&lastOrder)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return
		}
		err = nil
	}

	invoiceId = lastOrder.InvoiceId
	return
}

func (r repository) updateOrderStatus(ctx context.Context, invoice Invoice) (err error) {
	collection := r.dbMongo.Database(mongoDatabase).Collection(mongoOrderCollection)

	update := bson.M{
		"status": invoice.Status.value(),
	}
	_, err = collection.
		UpdateOne(
			ctx,
			bson.M{"invoice_id": invoice.ExternalId},
			bson.M{"$set": update},
		)

	return
}

func (r repository) findOneByInvoiceId(ctx context.Context, invoiceId string) (o Order, err error) {
	collection := r.dbMongo.Database(mongoDatabase).Collection(mongoOrderCollection)

	o = Order{}
	err = collection.FindOne(ctx, bson.M{"invoice_id": invoiceId}).Decode(&o)
	if err != nil {
		return
	}

	return
}
