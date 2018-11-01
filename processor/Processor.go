package processor

import (
	"database/sql"
)

func UpdateOrder(db *sql.DB) error{

	update, err := db.Query("UPDATE orders as o SET o.status='D' where o.created_at < NOW() and status='A'")

	if err != nil {
		return err
	}

	defer update.Close()

	return nil

}

func UpdateOrderProcessor(db *sql.DB) error{

	update, err := db.Query("UPDATE order_products as op SET op.status='D' where op.created_at < NOW() and op.status='A'")

	if err != nil {
		return err
	}

	defer update.Close()

	return nil

}