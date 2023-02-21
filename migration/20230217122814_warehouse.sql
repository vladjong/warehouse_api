-- +goose Up
-- +goose StatementBegin
CREATE TABLE warehouse (
  id bigserial primary key,
  name varchar(50) unique not null,
  is_available boolean not null
);

CREATE TABLE product (
  id uuid primary key,
  name varchar(50) unique not null,
  size int not null,
  qty int not null
);

CREATE TABLE warehouse_product (
  id bigserial primary key,
  warehouse_id bigserial references warehouse(id) on delete CASCADE,
  product_id uuid references product(id) on delete CASCADE,
  qty int not null,
  constraint unique_warehouse_product unique(warehouse_id, product_id)
);

CREATE TABLE product_reservation (
  id bigserial primary key,
  id_warehouse_product bigserial unique references warehouse_product(id) on delete CASCADE,
  qty int not null
);

CREATE INDEX product_id_in_warehouse_product ON warehouse_product(product_id);

CREATE INDEX warehouse_id_in_warehouse_product ON warehouse_product(warehouse_id);

CREATE INDEX warehouse_product_in_product_reservation ON product_reservation(id_warehouse_product);


CREATE OR REPLACE FUNCTION get_cnt_qty_by_id(idd uuid)
    RETURNS int AS
$$
    SELECT SUM(product_reservation.qty)
    FROM product_reservation
             JOIN warehouse_product wp on wp.id = product_reservation.id_warehouse_product
    WHERE wp.qty = 0 AND wp.product_id = idd
    GROUP BY wp.product_id;
$$
    LANGUAGE 'sql';

CREATE OR REPLACE FUNCTION check_product_in_zero()
    RETURNS trigger AS
$$
BEGIN
    DELETE FROM product WHERE get_cnt_qty_by_id(product.id) = 0;
    RETURN NULL;
END;
$$
    LANGUAGE 'plpgsql';

CREATE TRIGGER check_update_product
    AFTER UPDATE
    ON product_reservation
    FOR EACH ROW
EXECUTE FUNCTION check_product_in_zero();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX warehouse_id_in_warehouse_product CASCADE;
DROP INDEX product_id_in_warehouse_product CASCADE;
DROP INDEX warehouse_product_in_product_reservation CASCADE;
DROP TABLE warehouse CASCADE;
DROP TABLE product CASCADE;
DROP TABLE product_reservation CASCADE;
DROP TABLE warehouse_product CASCADE;
DROP FUNCTION get_cnt_qty_by_id(uuid);
DROP FUNCTION check_product_in_zero();
-- +goose StatementEnd
