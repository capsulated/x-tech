create table if not exists crypto (
    time timestamp not null,
    base varchar(5) not null,
    currency varchar(5) not null,
    rate float4 not null
);

create unique index if not exists crypto_timestamp_index on crypto(time);

comment on table crypto is 'crypto history data';


create table if not exists fiat (
    time timestamp not null,
    base varchar(3) not null,
    currency varchar(3) not null,
    rate float4 not null,
    constraint fiat_timestamp_constraint unique (time, base)
);

create index if not exists fiat_timestamp_index on fiat(time);

comment on table fiat is 'fiat history data';