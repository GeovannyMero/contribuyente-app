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

func (r *ContribuyenteRepository) Get(province string, page, limit int) ([]domain.Contribuyente, int64, error) {
	var contribuyentes []domain.Contribuyente

	// 1. Calcular Skip (Offset)
	skip := int64((page - 1) * limit)

	// 2. Definir Opciones de Búsqueda (Skip, Limit y Sort)
	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(skip)
	findOptions.SetSort(bson.D{{"_id", 1}}) // Ordenar es crucial para paginación consistente

	//Filtros
	filter := bson.D{}

	if province != "" {
		filter = append(filter, bson.E{Key: "codigo_juridiccion", Value: province})
	}

	//opts := options.Find().SetLimit(20)
	res, err := r.collection.Find(context.TODO(), filter, findOptions)

	for res.Next(context.TODO()) {
		var con domain.Contribuyente

		res.Decode(&con)
		contribuyentes = append(contribuyentes, con)

	}
	return contribuyentes, int64(len(contribuyentes)), err
}

func (r *ContribuyenteRepository) GetTotal() ([]domain.ContrinuyenteTotal, error) {
	var todalContribuyente []domain.ContrinuyenteTotal

	//Pipeline
	pipeline := mongo.Pipeline{
		// Etapa de agrupacion
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$codigo_juridiccion"}, //Agrupa por el campo de codigo de juridiccion
			{Key: "total", Value: bson.D{
				{Key: "$sum", Value: 1}, //Cuenta los documentos
			}},
		},
		}},
	}

	res, err := r.collection.Aggregate(context.TODO(), pipeline)

	for res.Next(context.TODO()) {
		var con domain.ContrinuyenteTotal

		res.Decode(&con)
		todalContribuyente = append(todalContribuyente, con)

	}

	return todalContribuyente, err
}
