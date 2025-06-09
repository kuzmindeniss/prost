## Backend for frontend run
```bash
air -d -c cmd/bff/.air.toml
```

## Tg-bot run
```bash
air -d -c cmd/tg/.air.toml
```

## Notification bot run
```bash
air -d -c cmd/tg_notifications/.air.toml
```

## Environment variables

For the application to work properly, you need to set the following environment variables:

1. `TELEGRAM_API_TOKEN` - Token for the main application bot
2. `NOTIFICATION_BOT_TOKEN` - Token for the notification bot
3. `ADMIN_CHAT_IDS` - Comma-separated list of chat IDs that should receive notifications
4. `RABBITMQ_URI` - URI for connecting to RabbitMQ (default: amqp://user:password@localhost:5672/)

You can set these in a `.env` file in the project root.

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
goose --dir internal/db/sql/schema postgres postgres://user:password@localhost:5432/prost down-to 0
```

## Bot Architecture

The project uses two Telegram bots connected via RabbitMQ message broker:

1. **Main application bot** - Collects applications from users
2. **RabbitMQ** - Message broker that decouples the bots and provides reliable message delivery
3. **Notification bot** - Receives messages via RabbitMQ and sends notifications to admin users

### Message Flow:

1. User submits an application through the main bot
2. Main bot saves the application to the database
3. Main bot publishes an "application.created" message to RabbitMQ
4. Notification bot consumes the message from RabbitMQ
5. Notification bot sends notifications to all subscribed admin users
