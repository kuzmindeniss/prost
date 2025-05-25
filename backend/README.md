## Backend for frontend run
```bash
air -d -c cmd/bff/.air.toml
```

## Tg-bot run
```bash
air -d -c cmd/tg/.air.toml
```

## SQLC generation
```bash
sqlc generate
```

## Migration

### Up
```bash
goose --dir internal/db/sql/schema postgres postgres://user:password@localhost:5432/prost up
```

### Down
```bash
goose --dir internal/db/sql/schema postgres postgres://user:password@localhost:5432/prost down
```
