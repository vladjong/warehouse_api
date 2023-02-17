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
  warehouse_id bigserial references warehouse(id),
  product_id uuid references product(id),
  qty int not null,
  constraint unique_warehouse_product unique(warehouse_id, product_id)
);

CREATE TABLE product_reservation (
  id bigserial primary key,
  id_warehouse_product bigserial references warehouse_product(id),
  qty int not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE warehouse CASCADE;
DROP TABLE product CASCADE;
DROP TABLE product_reservation CASCADE;
DROP TABLE warehouse_product CASCADE;
-- +goose StatementEnd
