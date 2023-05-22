CREATE TABLE IF NOT EXISTS telegram (
  id       serial       primary key ,
  chat_id integer not null unique,
  command varchar(255) not null,
  action varchar(255)
)