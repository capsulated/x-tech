create table if not exists crypto (
    time timestamp not null,
    ticker_source varchar(10) not null,
    ticker_target varchar(10) not null,
    rate bigint not null
);

create unique index if not exists crypto_timestamp_index on crypto(time);

comment on table crypto is 'crypto history data';


create table if not exists fiat (
    time timestamp not null,
    ticker_source varchar(10) not null,
    ticker_target varchar(10) not null,
    rate bigint not null,
    constraint fiat_timestamp_constraint unique (time, ticker_source)
);

create unique index if not exists fiat_timestamp_index on fiat(time);

comment on table fiat is 'fiat history data';