create schema sc_transaction

create table accounts (
	id serial primary key,
	document varchar(30),
	limit_max decimal(10,2)
)
