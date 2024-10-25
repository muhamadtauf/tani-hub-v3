package repository

import (
	"database/sql"
	"tani-hub-v3/structs"
)

func InsertOrder(db *sql.DB, order structs.Order) (err error) {
	sql := "INSERT INTO orders (code, status, total, user_id, address)" +
		" VALUES ($1, $2, $3, $4, $5)"

	errs := db.QueryRow(sql, order.Code, order.Status, order.Total, order.UserId, order.Address)

	for _, orderDetail := range order.OrderDetail {
		sql := "UPDATE products SET stock = stock - $1 WHERE id = $2"

		db.QueryRow(sql, orderDetail.Quantity, orderDetail.ProductId)

		//return errs.Err()
		//if errs != nil {
		//	panic(errs)
		//}
	}

	for _, orderDetail := range order.OrderDetail {
		sql := "INSERT INTO order_details (price, quantity, total, product_id, order_code)" +
			" VALUES ($1, $2, $3, $4, $5)"

		db.QueryRow(sql, orderDetail.Price, orderDetail.Quantity, orderDetail.Total, orderDetail.ProductId, orderDetail.OrderCode)

		//return errs.Err()
		//if errs != nil {
		//	panic(errs)
		//}
	}

	return errs.Err()
}

func GetAllOrder(db *sql.DB) (err error, results []structs.Order) {
	//sql := "SELECT * FROM orders"
	sql := "select orders.id, orders.code, orders.status,orders.total, orders.address, orders.created_at, orders.updated_at, orders.user_id, users.name from orders INNER join users on orders.user_id = users.id"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var order = structs.Order{}

		err = rows.Scan(&order.Id, &order.Code, &order.Status, &order.Total, &order.Address, &order.CreatedAt, &order.UpdatedAt, &order.UserId, &order.CustomerName)
		if err != nil {
			panic(err)
		}
		results = append(results, order)
	}
	return
}

func GetOrderByCode(db *sql.DB, order structs.Order) (err error, results []structs.Order) {
	sql := "SELECT * FROM orders WHERE code = $1"
	//sql := "SELECT * FROM orders INNER JOIN order_details ON orders.uuid = order_details.order_uuid WHERE uuid = $1"
	rows, err := db.Query(sql, order.Code)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var order = structs.Order{}

		err = rows.Scan(&order.Id, &order.Code, &order.Status, &order.Total, &order.Address, &order.CreatedAt, &order.UpdatedAt, &order.UserId)
		if err != nil {
			panic(err)
		}
		results = append(results, order)

		//detail
		sql := "SELECT * FROM order_details WHERE order_code = $1"
		rows, err := db.Query(sql, order.Code)
		if err != nil {
			panic(err)
		}

		defer rows.Close()
		for rows.Next() {
			var orderDetail = structs.OrderDetail{}

			err = rows.Scan(&orderDetail.Id, &orderDetail.Price, &orderDetail.Quantity, &orderDetail.Total, &orderDetail.ProductId, &orderDetail.OrderCode, &orderDetail.CreatedAt, &orderDetail.UpdatedAt)
			if err != nil {
				panic(err)
			}
			results[0].OrderDetail = append(results[0].OrderDetail, orderDetail)

		}
	}
	return
}

func UpdateOrderStatus(db *sql.DB, order structs.Order) (err error) {
	sql := "UPDATE orders SET status = $1 WHERE id = $2"

	errs := db.QueryRow(sql, order.Status, order.Id)

	return errs.Err()
}

func GetOrderByUserId(db *sql.DB, order structs.Order) (err error, results []structs.Order) {
	sql := "SELECT * FROM orders WHERE user_id = $1"
	rows, err := db.Query(sql, order.UserId)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var order = structs.Order{}

		err = rows.Scan(&order.Id, &order.Code, &order.Status, &order.Total, &order.Address, &order.CreatedAt, &order.UpdatedAt, &order.UserId)
		if err != nil {
			panic(err)
		}
		results = append(results, order)
	}
	return
}
