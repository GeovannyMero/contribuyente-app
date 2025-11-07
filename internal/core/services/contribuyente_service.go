// Contiene la lógica de aplicación. Recibe la interfaz del repositorio como dependencia.
package services

import (
	"github.com/geovannymero/contribuyente-app/internal/core/domain"
	"github.com/geovannymero/contribuyente-app/internal/core/ports"
	response "github.com/geovannymero/contribuyente-app/pkg/utils"
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

func (s *ContribuyenteService) Get(province string, page, limit int) (response.Response[[]domain.Contribuyente], error) {
	data, totalCount, err := s.repo.Get(province, page, limit)
	println(totalCount)
	return response.Ok(data), err
}
