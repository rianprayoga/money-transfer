# TL;DR

There are 2 fields in payload, those are amount and success.

```javascript
{
  "amount": 123,
  "success": true
}
```

The **amount** is to specify how much we want to deduct the merchant balance, and the **success** is a flag to determine whether the money transfer is failed or success.

If the **success** is false, the transaction is failed and the money is returned to the merchant's balance. By default merchant's balance is 100.

For the pub-sub, we use redis for the simplicity.

# Prerequisite

go version go1.25.5

Client: Docker Engine Version:29.1.3

Server: Docker Engine Version:29.1.3

# Run database and redis

Inside Dockerfile folder run:

```
docker compose up
```

# Run the service

Inside root folder **money-transfer/** run:

```
go run ./cmd/
```

# Burst the API

There is test folder(**money-transfer/test**) to test the logic of the service.

The **burst.sh**

It will send 10 request at the same time with same payload in each request, which is reducing the balance by 30.
So there will be 3 request success and 7 will be error **insufficient balance**. It is because the default balance is 100 for the merchant.

# To Improve

```
- Error handling
- More test case to cover the transaction logic
```
