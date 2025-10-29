package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contribuyente struct {
	Id                      primitive.ObjectID `json:"_id" bson:"_id"`
	Ruc                     string             `json:"ruc" bson:"ruc"`
	RazonSocial             string             `json:"razon_social" bson:"razon_social"`
	CodigoJuridiccion       string             `json:"codigo_juridiccion" bson:"codigo_juridiccion"`
	EstadoContribuyente     string             `json:"estado_contribuyente" bson:"estado_contribuyente"`
	ClaseContribuyente      string             `json:"clase_contribuyente" bson:"clase_contribuyente"`
	Obligado                string             `json:"obligado" bson:"obligado"`
	TipoContribuyente       string             `json:"tipo_contribuyente" bson:"tipo_contribuyente"`
	NumeroEstablecimiento   int32              `json:"numero_establecimiento" bson:"numero_establecimiento"`
	NombreFantasiaComercial string             `json:"nombre_fantasia_comercial" bson:"nombre_fantasia_comercial"`
	EstadoEstablecimiento   string             `json:"estado_establecimiento" bson:"estado_establecimiento"`
	Provincia               string             `json:"descripcion_provincia_est" bson:"descripcion_provincia_est"`
	Canton                  string             `json:"descripcion_canton_est" bson:"descripcion_canton_est"`
	Parroquia               string             `json:"descripcion_parroquia_est" bson:"descripcion_parroquia_est"`
	CodigoCiiu              string             `json:"codigo_ciiu" bson:"codigo_ciiu"`
	ActividadEconomica      string             `json:"actividad_economica" bson:"actividad_economica"`
	AgenteRetencion         string             `json:"agente_retencion" bson:"agente_retencion"`
	Especial                string             `json:"especial" bson:"especial"`
}
