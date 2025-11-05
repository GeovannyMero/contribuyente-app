//Implementa la interfaz definida en ports/. Aquí reside el código de MongoDB.

package mongodb

import (
	"context"

	"github.com/geovannymero/contribuyente-app/internal/core/domain"
	"github.com/geovannymero/contribuyente-app/internal/core/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UserRepository implementa la interfaz ports.UserRepository para MongoDB.
type ContribuyenteRepository struct {
	collection *mongo.Collection
}

var _ ports.ContribuyenteRepository = (*ContribuyenteRepository)(nil)

func NewMongoRepository(client *mongo.Client, dbName, collName string) ports.ContribuyenteRepository {
	return &ContribuyenteRepository{
		collection: client.Database(dbName).Collection(collName),
	}
}

func (r *ContribuyenteRepository) Get(province string) ([]domain.Contribuyente, error) {
	var contribuyentes []domain.Contribuyente

	opts := options.Find().SetLimit(20)
	res, err := r.collection.Find(context.TODO(), bson.D{}, opts)

	for res.Next(context.TODO()) {
		var con domain.Contribuyente

		res.Decode(&con)
		contribuyentes = append(contribuyentes, con)

	}
	return contribuyentes, err
}
