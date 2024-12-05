-- +goose Up
-- +goose StatementBegin

create table rates
(
    id         integer generated always as identity,
    code       text,
    nominal    bigint,
    course     text,
    kopecks    bigint,
    ts         timestamp,
    created_at timestamp,
    updated_at timestamp
);

create index rates_code_ts_idx on rates(code, ts);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop index rates_code_ts_idx;
drop table rates;

-- +goose StatementEnd
