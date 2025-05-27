package database

import (
	"database/sql"

	"github.com/Celio-Batalha/pos-go-listOrders/internal/entity"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) FindAll() ([]*entity.Order, error) {
	rows, err := r.db.Query("SELECT id, price, tax, final_price FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*entity.Order
	for rows.Next() {
		var order entity.Order
		if err := rows.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}
	return orders, nil
}

// func (r *OrderRepository) FindAll() ([]*entity.Order, error) {
// 	stmt, err := r.db.Prepare("SELECT id, price, tax, final_price FROM orders")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer stmt.Close()

// 	rows, err := stmt.Query()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var orders []*entity.Order
// 	for rows.Next() {
// 		var order entity.Order
// 		err := rows.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice)
// 		if err != nil {
// 			return nil, err
// 		}
// 		orders = append(orders, &order)
// 	}

// 	return orders, nil
// }
