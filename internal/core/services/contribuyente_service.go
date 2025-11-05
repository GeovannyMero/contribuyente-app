// Contiene la lógica de aplicación. Recibe la interfaz del repositorio como dependencia.
package services

import (
	"github.com/geovannymero/contribuyente-app/internal/core/domain"
	"github.com/geovannymero/contribuyente-app/internal/core/ports"
)

// UserService implementa la lógica de negocio para los usuarios.
type ContribuyenteService struct {
	repo ports.ContribuyenteRepository
}

func NewContribuyenteService(r ports.ContribuyenteRepository) *ContribuyenteService {
	return &ContribuyenteService{
		repo: r,
	}
}

func (s *ContribuyenteService) Get(province string) ([]domain.Contribuyente, error) {
	return s.repo.Get(province)
}
