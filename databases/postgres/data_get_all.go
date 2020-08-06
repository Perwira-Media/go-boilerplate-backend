package postgres

import (
	model "github.com/Perwira-Media/go-boilerplate-backend/models"
)

// GetAllData retrieves all data from postgresdb
func (p *Postgres) GetAllData() ([]model.Data, error) {
	result := []model.Data{}

	query := "SELECT * FROM data"

	data, err := p.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer data.Close()

	for data.Next() {
		datum := model.Data{}
		err = data.Scan(&datum.Name)
		if err != nil {
			return nil, err
		}

		result = append(result, datum)
	}

	return result, nil
}
