package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
)

type OrderData struct {
	OrderUid          string `json:"order_uid" db:"order_uid" validate:"required"`
	TrackNumber       string `json:"track_number" db:"track_number" validate:"required"`
	Entry             string `json:"entry" db:"entry"`
	Delivery          `json:"delivery" validate:"required"`
	Payment           `json:"payment" validate:"required"`
	Items             Items     `json:"items" db:"items" validate:"required"`
	Locale            string    `json:"locale" db:"locale"`
	InternalSignature string    `json:"internal_signature" db:"internal_signature"`
	CustomerId        string    `json:"customer_id" db:"customer_id" validate:"required"`
	DeliveryService   string    `json:"delivery_service" db:"delivery_service" validate:"required"`
	Shardkey          string    `json:"shardkey" db:"shardkey"`
	SmId              int       `json:"sm_id" db:"sm_id"`
	DateCreated       time.Time `json:"date_created" db:"date_created" validate:"required"`
	OofShard          string    `json:"oof_shard" db:"oof_shard"`
}

type Delivery struct {
	Name    string `json:"name" db:"name" validate:"required"`
	Phone   string `json:"phone" db:"phone" validate:"required"`
	Zip     string `json:"zip" db:"zip"`
	City    string `json:"city" db:"city" validate:"required"`
	Address string `json:"address" db:"address"  validate:"required"`
	Region  string `json:"region" db:"region" `
	Email   string `json:"email" db:"email"  validate:"required,email"`
}

type Payment struct {
	Transaction  string `json:"transaction" db:"transaction" validate:"required"`
	RequestId    string `json:"request_id" db:"request_id"`
	Currency     string `json:"currency" db:"currency" validate:"required"`
	Provider     string `json:"provider" db:"provider"`
	Amount       int    `json:"amount" db:"amount" validate:"required"`
	PaymentDt    int    `json:"payment_dt" db:"payment_dt"`
	Bank         string `json:"bank" db:"bank"`
	DeliveryCost int    `json:"delivery_cost" db:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total" db:"goods_total"`
	CustomFee    int    `json:"custom_fee" db:"custom_fee"`
}

type Items []struct {
	ChrtId      int64  `json:"chrt_id" db:"chrt_id"`
	OrderUid    string `json:"-" db:"order_uid"`
	TrackNumber string `json:"track_number" db:"track_number"`
	Price       int    `json:"price" db:"price" validate:"required"`
	Rid         string `json:"rid" db:"rid"`
	Name        string `json:"name" db:"name" validate:"required"`
	Sale        int    `json:"sale" db:"sale"`
	Size        string `json:"size" db:"size"`
	TotalPrice  int    `json:"total_price" db:"total_price" validate:"required"`
	NmId        int    `json:"nm_id" db:"nm_id"`
	Brand       string `json:"brand" db:"brand"`
	Status      int    `json:"status" db:"status"`
}

func (i *Items) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSON value")
	}
	return json.Unmarshal(bytes, i)
}

func (i *Items) Value() (driver.Value, error) {
	return json.Marshal(i)
}

func (o *OrderData) Validate() error {
	validate := validator.New()

	if err := validate.Struct(o); err != nil {
		return err
	}

	// Валидация вложенных структур
	if err := validate.Struct(o.Delivery); err != nil {
		return err
	}
	if err := validate.Struct(o.Payment); err != nil {
		return err
	}
	for _, item := range o.Items {
		if err := validate.Struct(item); err != nil {
			return err
		}
	}

	return nil
}
