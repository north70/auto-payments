CREATE TABLE IF NOT EXISTS auto_payments (
    id       serial       primary key ,
    chat_id integer not null,
    name     varchar(255) not null,
    period_day integer not null,
    payment_day integer not null,
    amount integer not null,
    count_pay integer not null,
    next_pay_date timestamp not null,
    created_at timestamp
);
ALTER TABLE auto_payments DROP CONSTRAINT IF EXISTS chat_id_name_unique;

ALTER TABLE auto_payments ADD CONSTRAINT chat_id_name_unique UNIQUE (chat_id, name);

