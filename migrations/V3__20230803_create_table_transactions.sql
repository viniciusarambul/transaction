create table transactions (
	id serial primary key,
    idempotency_key varchar(40) unique,
    account_id int,
	operation_type_id int,
	amount numeric(10,2),
	event_date timestamp,
	created_at timestamp,
	updated_at timestamp,
    CONSTRAINT fk_accounts
      FOREIGN KEY(account_id) 
	  REFERENCES accounts(id)
);