package storage

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mjeffers-998/interview-task/structs"
)

type DB struct {
	db *sql.DB
}

func NewDB() (*DB, error) {
	db, err := sql.Open("sqlite3", "orders.db")
	if err != nil {
		return nil, err
	}
	err = RunMigration(db)
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}

func RunMigration(db *sql.DB) error {
	//obviously not a proper implementation, but for a short task it's enough
	_, err := db.Exec(`CREATE TABLE orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		order_date TIMESTAMP,
		order_value REAL,
		order_status TEXT CHECK (order_status IN ('new', 'cancelled', 'shipped')),
		store_id INTEGER
		);`)
	return err
}

func (d *DB) CreateOrder(order *structs.Order) (int64, error) {
	tx, err := d.db.Begin()
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()
	result, err := tx.Exec(`INSERT INTO orders (
		order_date, order_value, order_status, store_id
	) VALUES (
		$1, $2, $3, $4
	)`, order.OrderDate, order.OrderValue, order.Status, order.StoreID)
	if err != nil {
		return -1, err
	}
	lastID, err := result.LastInsertId()
	tx.Commit()
	return lastID, err
}

func (d *DB) CancelOrder(id int64) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec("UPDATE orders SET order_status = $1 WHERE order_id = $2", structs.CancelledOrder, id)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (d *DB) ListAllOrders(status string) ([]*structs.Order, error) {
	rows, err := d.db.Query("SELECT * FROM orders WHERE order_status = $1", status)
	if err != nil {
		return nil, err
	}
	orders := []*structs.Order{}
	for rows.Next() {
		tmp := structs.Order{}
		err := rows.Scan(&tmp.OrderID, &tmp.OrderDate, &tmp.OrderValue, &tmp.Status, &tmp.StoreID)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &tmp)
	}
	return orders, nil
}

func (d *DB) GetOrderByID(id int64) (*structs.Order, error) {
	var order structs.Order
	err := d.db.QueryRow("SELECT * FROM orders WHERE order_id = $1", id).Scan(&order.OrderID, &order.OrderDate, &order.OrderValue, &order.Status, &order.StoreID)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (d *DB) GetRevenue(startDate time.Time, endTime time.Time) (float32, error) {
	var revenue float32
	err := d.db.QueryRow("SELECT count(order_status) FROM order WHERE order_date BETWEEN $1 AND $2").Scan(&revenue)
	if err != nil {
		return -1, err
	}
	return revenue, nil
}
