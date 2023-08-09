# Transactions

Transaction is a simple service for create transactions.

## Start!

Clone the repository 

```bash
git clone https://github.com/viniciusarambul/transaction.git
```

## Utilities commands

```bash
# up postgres container
make infra.up

# make migrate
make flyway

# up API on default port 8080
make run-api

# run coverage test
make test.cover
```

## Routes

```bash
# Create account
curl --location --request POST 'localhost:8080/accounts' \
--header 'Content-Type: application/json' \
--data-raw '{
    "document": "123456789",
    "limit_max": 200.00
}'

# Find Account
curl --location --request GET 'localhost:8080/accounts/1' \
--data-raw ''

# Create Transaction
curl --location --request POST 'localhost:8080/transactions' \
--header 'Content-Type: application/json' \
--data-raw '{
    "idempotency_key": "1234561711212123",
    "account_id": 1,
    "operation_type_id": 2,
    "amount": 0.01
}'

```