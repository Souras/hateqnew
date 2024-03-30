package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/sourabh/models"
)

var db *sql.DB

func InitDb() error {
	connectionString := "user=casaos dbname=todos password=casaos host=192.168.1.12 port=5432 sslmode=disable" // Replace with your credentials
	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}
	return nil
}

func GetProducts() ([]models.Product, error) {
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func GetProduct(id int64) (models.Product, error) {
	row := db.QueryRow("SELECT * FROM products WHERE id = $1", id)
	var p models.Product
	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Product{}, err
		}
		return models.Product{}, err
	}
	return p, nil
}

func CreateProduct(p models.Product) (models.Product, error) {
	stmt, err := db.Prepare("INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return models.Product{}, err
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(p.Name, p.Price).Scan(&id)
	if err != nil {
		return models.Product{}, err
	}

	p.ID = id
	return p, nil
}

func UpdateProduct(id int64, p models.Product) (models.Product, error) {
	stmt, err := db.Prepare("UPDATE products SET name = $1, price = $2 WHERE id = $3 RETURNING id")
	if err != nil {
		return models.Product{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(p.Name, p.Price, id)
	if err != nil {
		return models.Product{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Product{}, err
	}

	if rowsAffected == 0 {
		return models.Product{}, sql.ErrNoRows
	}

	p.ID = id
	return p, nil
}

func DeleteProduct(id int64) error {
	stmt, err := db.Prepare("DELETE FROM products WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
