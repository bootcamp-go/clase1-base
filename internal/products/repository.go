package internal

import (
	"database/sql"
	"log"

	"github.com/bootcamp-go/clase1-base/internal/models"
)

// Se podria hacer la inteface en clase y dejar que los bootcampers te guien para ver si enteiendieron el concepto
type Repository interface {
	GetOne(id int) (models.Product, error)
	Store(product models.Product) (models.Product, error)
	Update(id int, product models.Product) (models.Product, error)
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) GetOne(id int) (models.Product, error) {
	var product models.Product
	rows, err := r.db.Query("SELECT * FROM products where id = ?", id)
	if err != nil {
		log.Println(err)
		return product, err
	}
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			log.Println(err.Error())
			return product, err
		}
	}
	return product, nil
}

func (r *repository) Store(product models.Product) (models.Product, error) {
	stmt, err := r.db.Prepare("INSERT INTO products(name, type, count, price) VALUES( ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria

	var result sql.Result
	result, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price)
	if err != nil {
		return models.Product{}, err
	}
	insertedId, _ := result.LastInsertId() // del sql.Result devuelto en la ejecución obtenemos el Id insertado
	product.ID = int(insertedId)

	return product, nil
}

func (r *repository) Update(id int, product models.Product) (models.Product, error) {
	stmt, err := r.db.Prepare("UPDATE products SET name = ?, type = ?, count = ?, price = ? WHERE id = ?") // se prepara la sentencia SQL a ejecutar
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price, id)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil

}

func (r *repository) Delete(id int) error {
	stmt := "DELETE FROM products WHERE id = ?" // Explicar que esta sentencia es un hard delete (no se puede revertir)
	// y explicar que se va a eliminar el producto la difetencia con un soft y como es la sentencia
	_, err := r.db.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil

}
