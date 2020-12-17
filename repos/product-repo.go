package repos

import (
	"github.com/vSterlin/sw-store/database"
)

type ProductRepo struct {
}

type productModel struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

func NewProductRepo() *ProductRepo {
	return &ProductRepo{}
}

func (pr *ProductRepo) FindAll() []*productModel {

	rows, err := database.DB.QueryMany("SELECT * FROM products;")
	if err != nil {
		panic(err)
	}
	pmArr := []*productModel{}
	for rows.Next() {
		pm := &productModel{}
		rows.Scan(&pm.ID, &pm.Name, &pm.Price, &pm.Description, &pm.CreatedAt, &pm.UpdatedAt)
		pmArr = append(pmArr, pm)
	}
	return pmArr
}

func (pr *ProductRepo) FindOne(id int) (*productModel, error) {
	row := database.DB.QueryOne("SELECT * FROM products WHERE id=$1;", id)
	pm := &productModel{}

	err := row.Scan(&pm.ID, &pm.Name, &pm.Price, &pm.Description, &pm.CreatedAt, &pm.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return pm, nil
}
