-- Таблица для хранения общей информации о заказе
CREATE TABLE order_info (
                            order_uid TEXT PRIMARY KEY,
                            track_number TEXT,
                            entry TEXT,
                            customer_id TEXT,
                            delivery_service TEXT,
                            date_created TIMESTAMP,
                            shardkey TEXT,
                            sm_id INT,
                            oof_shard TEXT
);

-- Таблица для хранения информации о доставке
CREATE TABLE delivery_info (
                               order_uid TEXT REFERENCES order_info(order_uid),
                               name TEXT,
                               phone TEXT,
                               zip TEXT,
                               city TEXT,
                               address TEXT,
                               region TEXT,
                               email TEXT
);

-- Таблица для хранения информации о платеже
CREATE TABLE payment_info (
                              order_uid TEXT REFERENCES order_info(order_uid),
                              transaction TEXT,
                              request_id TEXT,
                              currency TEXT,
                              provider TEXT,
                              amount INT,
                              payment_dt BIGINT,
                              bank TEXT,
                              delivery_cost INT,
                              goods_total INT,
                              custom_fee INT
);

-- Таблица для хранения информации о товарах в заказе
CREATE TABLE order_items (
                             chrt_id BIGSERIAL PRIMARY KEY,
                             order_uid TEXT REFERENCES order_info(order_uid),
                             track_number TEXT,
                             price INT,
                             rid TEXT,
                             name TEXT,
                             sale INT,
                             size TEXT,
                             total_price INT,
                             nm_id INT,
                             brand TEXT,
                             status INT
);