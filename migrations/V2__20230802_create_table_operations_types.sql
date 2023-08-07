create table operations_types (
	id serial primary key,
	operation_type_id int,
	description varchar(25),
	type varchar(10)
);

insert into operations_types(operation_type_id,"description", "type") values (1, 'COMPRA A VISTA', 'DEBIT');
insert into operations_types(operation_type_id,"description", "type") values (2, 'COMPRA PARCELADA', 'DEBIT');
insert into operations_types(operation_type_id,"description", "type") values (3, 'SAQUE', 'DEBIT');
insert into operations_types(operation_type_id,"description", "type") values (4, 'PAGAMENTO', 'CREDIT');
