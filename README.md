# Live Chat Support System

A real-time chat support system built with Go, featuring WebSocket communication, JWT authentication, and message persistence using SQLite and Redis.

## Tech Stack

- **Backend Framework**: Go with Gin Web Framework
- **Database**: 
  - SQLite (Primary database)
  - Redis (Message caching and real-time features)
- **Authentication**: JWT (JSON Web Tokens)
- **WebSocket**: Gorilla WebSocket for real-time communication
- **Testing**: Go testing framework with Testify suite
- **Logging**: Uber's Zap logger

## Features

- User authentication (Register/Login)
- Real-time messaging using WebSockets
- Message persistence in SQLite database
- Recent messages caching in Redis
- Secure password hashing using bcrypt
- JWT-based authentication
- Comprehensive test suite
- Production-grade logging

## Project Structure

```
├── config/
│   └── config.go           # Configuration management
├── controllers/
│   ├── auth_controller.go  # Authentication handlers
│   └── chat_controller.go  # Chat-related handlers
├── database/
│   └── database.go         # Database connection and migrations
├── middleware/
│   └── jwt_middleware.go   # JWT authentication middleware
├── models/
│   ├── message.go         # Message model
│   └── user.go            # User model
├── routes/
│   └── routes.go          # Route definitions
├── tests/
│   └── endpoints_test.go  # API endpoint tests
├── utils/
│   ├── logger.go          # Logging utility
│   └── redis_client.go    # Redis connection
├── websocket/
│   └── websocket.go       # WebSocket handler
└── main.go                # Application entry point
```

## Prerequisites

- Go 1.19 or higher
- Redis server
- SQLite

## Installation & Setup

1. Clone the repository:
```bash
git clone <repository-url>
cd livechat-support
```

2. Install dependencies:
```bash
go mod download
```

3. Start Redis server:
```bash
redis-server
```

4. Set up environment variables (optional):
```bash
export JWT_SECRET=your_custom_secret
```

5. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

- `POST /register` - Register a new user
- `POST /login` - User login
- `GET /ws` - WebSocket connection endpoint
- `POST /save-message` - Save a new message
- `GET /recent-messages` - Retrieve recent messages

## Running Tests

Execute the test suite:
```bash
go test ./tests -v
```

## Development

1. The application uses SQLite for development. The database file will be created automatically.
2. Redis is used for caching recent messages and managing real-time features.
3. WebSocket connections are handled at the `/ws` endpoint.

## Security Features

- Password hashing using bcrypt
- JWT-based authentication
- Secure WebSocket implementation
- Environment-based configuration

## Monitoring and Logging

The application uses Uber's Zap logger for structured logging:
- Production-grade logging in production mode
- Development-friendly logging in development mode
- Test-specific logging configuration for tests

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.