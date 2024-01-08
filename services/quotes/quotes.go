package services

import (
	"fmt"

	"github.com/pedrohscramos/teste-frete-rapido/database/mysql"
	"github.com/pedrohscramos/teste-frete-rapido/models"
)

var (
	repository *GormRepository
)

type GormRepository struct {
	db *mysql.DB
}

func NewGormRepository(db *mysql.DB) *GormRepository {
	repository = &GormRepository{
		db: db,
	}

	return repository
}

func GetRepository() GormRepository {
	return *repository
}

type QuoteService interface {
	InsertQuote(c *models.Quote) (*models.Quote, error)
	GetLastQuotes(limit uint64) (interface{}, error)
}

type ResponseQuotes struct {
	ID       uint64  `json:"id"`
	Name     string  `json:"name"`
	Service  string  `json:"service"`
	Deadline string  `json:"deadline"`
	Price    float64 `json:"price"`
}

func (r *GormRepository) InsertQuote(c *models.Quote) (*models.Quote, error) {
	result := r.db.Database.Create(c)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *GormRepository) GetLastQuotes(limit uint64) (interface{}, error) {
	var quotes = []ResponseQuotes{}

	query := fmt.Sprintf(`SELECT q.id, q.name, q.service, q.deadline, q.price FROM quotes AS q ORDER BY q.price DESC LIMIT 0,%v`, limit)

	err := r.db.Database.Raw(query).Scan(&quotes).Error

	var responseQuotes []ResponseQuotes

	for _, q := range quotes {
		responseQuotes = append(responseQuotes, ResponseQuotes{
			ID:       q.ID,
			Name:     q.Name,
			Service:  q.Service,
			Deadline: q.Deadline,
			Price:    q.Price,
		})
	}

	var data = map[string]interface{}{
		"quotes": &responseQuotes,
	}

	return data, err
}
