-- +goose Up
-- +goose StatementBegin
insert into stock(sku, available) 
values (1076963, 1000),
       (1148162, 1000),
       (1625903, 1000),
       (2618151, 1000),
       (2956315, 1000),
       (2958025, 1000),
       (3596599, 1000),
       (3618852, 1000),
       (4288068, 1000),
       (4465995, 1000);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
truncate table stock;
-- +goose StatementEnd
