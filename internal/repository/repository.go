package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"wb-tech/internal/model"
)

type Repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) Save(ctx context.Context, orderData model.OrderData) (err error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("SQL: AddOrderData: begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	_, err = tx.NamedExecContext(ctx,
		`INSERT INTO order_info (order_uid, track_number, entry, customer_id, delivery_service, date_created, 
                        shardkey, sm_id, oof_shard) 
			 VALUES (:order_uid, :track_number, :entry, :customer_id, :delivery_service, :date_created, 
			         :shardkey, :sm_id, :oof_shard)`,
		orderData)
	if err != nil {
		return fmt.Errorf("SQL: AddOrderData: insert order_info:%w", err)
	}

	_, err = tx.NamedExecContext(ctx,
		`INSERT INTO delivery_info (order_uid, name, phone, zip, city, address, region, email)
			 VALUES (:order_uid, :name, :phone, :zip, :city, :address, :region, :email)`, orderData)
	if err != nil {
		return fmt.Errorf("SQL: AddOrderData: insert delivery_info:%w", err)
	}

	_, err = tx.NamedExecContext(ctx,
		`INSERT INTO payment_info (order_uid, transaction, request_id, currency, provider, amount, 
                          payment_dt, bank, delivery_cost, goods_total, custom_fee) 
			 VALUES (:order_uid, :transaction, :request_id, :currency, :provider, :amount, :payment_dt, :bank, 
			         :delivery_cost, :goods_total, :custom_fee)`,
		orderData)
	if err != nil {
		return fmt.Errorf("SQL: AddOrderData: insert payment_info:%w", err)
	}

	stmt, err := tx.PrepareNamedContext(ctx,
		`INSERT INTO order_items (chrt_id, order_uid, track_number, price, rid, name, sale, size, total_price,
	                    nm_id, brand, status)
			   VALUES (:chrt_id, :order_uid, :track_number, :price, :rid, :name, :sale, :size, :total_price,
			           :nm_id, :brand, :status)`)
	if err != nil {
		return fmt.Errorf("SQL: Save: prepare order_items statement: %w", err)
	}
	defer stmt.Close()

	for _, item := range orderData.Items {
		_, err = stmt.ExecContext(ctx, item)
		if err != nil {
			return fmt.Errorf("SQL: Save: insert order_items: %w", err)
		}
	}

	return nil
}
func (r *Repo) LoadCache(ctx context.Context) (map[string]model.OrderData, error) {
	// Выбираем все данные из всех таблиц
	query := `
		SELECT
				oi.*,
				di.name, di.phone, di.zip,
				di.city, di.address, di.region, di.email,
				pi.transaction, pi.request_id, pi.currency, pi.provider, pi.amount, pi.payment_dt, pi.bank,
				pi.delivery_cost, pi.goods_total, pi.custom_fee,
				(
					SELECT json_agg(
						json_build_object(
							'chrt_id', oit.chrt_id,
							'order_uid', oit.order_uid,
							'track_number', oit.track_number,
							'price', oit.price,
							'rid', oit.rid,
							'name', oit.name,
							'sale', oit.sale,
							'size', oit.size,
							'total_price', oit.total_price,
							'nm_id', oit.nm_id,
							'brand', oit.brand,
							'status', oit.status
						)
					)
					FROM order_items oit
					WHERE oit.order_uid = oi.order_uid
				) AS items
			FROM order_info oi
			JOIN delivery_info di ON oi.order_uid = di.order_uid
			JOIN payment_info pi ON oi.order_uid = pi.order_uid
	`

	rows, err := r.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("SQL: LoadCache: %w", err)
	}
	defer rows.Close()

	cache := make(map[string]model.OrderData)

	for rows.Next() {
		var order model.OrderData
		if err := rows.StructScan(&order); err != nil {
			return nil, fmt.Errorf("SQL: LoadCache: structScan: %w", err)
		}
		cache[order.OrderUid] = order
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("SQL: LoadCache: %w", err)
	}

	return cache, nil
}
