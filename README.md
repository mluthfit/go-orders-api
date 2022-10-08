# Objective

Build RESTful API of order and items including CRUD operations.

## Criteria

- [ ] No error
- [x] Using `gin-gonic/gin`
- [x] Using `gorm`
- [x] Using git
- [ ] Using proper HTTP methods
- [ ] Error handling
- [ ] Dependency injection implementation

## Endpoint

| Route            | HTTP     | Description               |
| ---------------- | -------- | ------------------------- |
| /orders          | `GET`    | Get all orders            |
| /orders          | `POST`   | Create order and items    |
| /orders/:orderId | `GET`    | Get a order               |
| /orders/:orderId | `PUT`    | Update a orders and items |
| /orders/:orderId | `DELETE` | Delete a order            |

Request body for `POST`

```json
{
  "orderedAt": "2019-11-09T21:21:46+07:00",
  "customerName": "John Doe",
  "items": [
    {
      "itemCode": "IP-001",
      "description": "IPhone 10X",
      "quantity": 1
    },
    {
      "itemCode": "B-001",
      "description": "Bag",
      "quantity": 10
    }
  ]
}
```

Request body for `PUT`

```json
{
  "customerName": "Spike Tyke",
  "orderedAt": "2020-11-09T21:21:46+07:00",
  "items": [
    {
      "lineItemId": 1,
      "itemCode": "IP-002",
      "description": "IPhone 10X",
      "quantity": 10
    }
  ]
}
```
