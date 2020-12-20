package store

import (
	"github.com/vSterlin/sw-store/internal/app/model"
)

type ProductRepository struct {
	store *Store
}

func (pr *ProductRepository) FindAll() ([]*model.Product, error) {

	rows, err := pr.store.db.Query("SELECT * FROM products;")
	if err != nil {
		return nil, err
	}

	pmArr := []*model.Product{}
	for rows.Next() {
		pm := &model.Product{}
		rows.Scan(&pm.ID, &pm.Name, &pm.Price, &pm.Description, &pm.CreatedAt, &pm.UpdatedAt)
		pmArr = append(pmArr, pm)
	}
	return pmArr, nil

}
func (pr *ProductRepository) FindById(id int) (*model.Product, error) {
	row := pr.store.db.QueryRow("SELECT * FROM products WHERE id=$1;", id)
	pm := &model.Product{}

	err := row.Scan(&pm.ID, &pm.Name, &pm.Price, &pm.Description, &pm.CreatedAt, &pm.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return pm, nil
}
