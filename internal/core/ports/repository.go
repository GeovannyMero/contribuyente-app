//Define el contrato de cómo interactuar con los datos. Esta interfaz no sabe que se usará MongoDB.

package ports

import "github.com/geovannymero/contribuyente-app/internal/core/domain"

type ContribuyenteRepository interface {
	Get(province string, page, limit int) ([]domain.Contribuyente, int64, error)
}
