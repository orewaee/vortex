## Vortex

Vortex is an open-source Telegram-based ticketing system with a web-based user interface.


## Routes

#### `/v1/`

###### `POST /ticket` - open ticket

```json
{
    "chat_id": "int64",
    "topic": "string"
}
```

- `201` - created
- `409` - ticket already exists

