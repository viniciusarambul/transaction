create schema if not exists sc_transaction;

create table accounts (
	id serial primary key,
	document varchar(30),
	limit_max decimal(10,2),
	created_at timestamp,
	updated_at timestamp
);
