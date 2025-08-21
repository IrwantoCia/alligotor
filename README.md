# Alligotor 🐊

A powerful Domain-Driven Design (DDD) folder structure generator for Go applications. Alligotor helps developers quickly scaffold well-organized, scalable Go projects following clean architecture principles.

## 🚀 Features

- **DDD Architecture**: Generates complete Domain-Driven Design folder structures
- **Clean Architecture**: Implements layered architecture (Domain, Application, Delivery, Infrastructure)
- **Multiple Contexts**: Support for different business contexts (User, Payment, etc.)
- **Web Templates**: Includes ready-to-use web application templates with modern UI
- **Database Support**: PostgreSQL and SQLite support with GORM
- **Authentication**: JWT-based authentication system
- **Modern Stack**: Uses Gin framework, Zap logger, and other popular Go libraries

## 📁 Project Structure

```
alligotor/
├── cmd/gotor/           # Main CLI application
│   ├── libs/           # Core generation logic
│   └── templates/      # Template files for different contexts
├── templates/           # Template definitions
└── go.mod              # Go module file
```

## 🛠️ Installation

### Prerequisites
- Go 1.24.5 or higher

### Install from source
```bash
# Clone the repository
git clone https://github.com/IrwantoCia/alligotor.git
cd alligotor

# Install the CLI tool
make install
# or manually
go install ./cmd/gotor
```

## 📖 Usage

### Initialize a new project
```bash
gotor init <project_name>
```

This command generates a complete Go project with:
- Clean architecture folder structure
- Web application setup with Gin framework
- Database configuration (PostgreSQL/SQLite)
- Authentication middleware
- Logging with Zap
- Environment configuration with Viper
- Modern web UI templates

### Create a new business context
```bash
gotor context <context_name>
```

Generates a new business context with:
- Domain entities and repositories
- Application services
- HTTP handlers
- Database schemas and implementations

### Generate payment context
```bash
gotor payment
```

Creates a specialized payment context with:
- Payment entities and DTOs
- Payment gateway interfaces
- Repository implementations
- Payment-specific exceptions

## 🏗️ Generated Project Structure

When you run `gotor init <project_name>`, it creates:

```
<project_name>/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── auth_context/     # Authentication domain
│   ├── user_context/     # User management domain
│   ├── shared/           # Shared utilities
│   │   ├── config/       # Configuration management
│   │   ├── db/           # Database connection
│   │   ├── logger/       # Logging utilities
│   │   └── middleware/   # HTTP middleware
│   └── payment_context/  # Payment processing domain
├── web/                  # Web assets
│   ├── static/
│   │   ├── css/
│   │   └── js/
│   └── views/
├── go.mod
├── go.sum
├── Makefile
└── package.json
```

## 🔧 Available Commands

| Command | Description |
|---------|-------------|
| `gotor init <name>` | Initialize a new Go project with DDD structure |
| `gotor context <name>` | Create a new business context |
| `gotor payment` | Generate payment-specific context |

## 🧪 Testing

```bash
# Run all tests
make test

# Clean test cache
make clean
```

## 🏗️ Build

```bash
# Build the CLI tool
make build

# Install globally
make install
```

## 📚 Dependencies

The generated projects include:

- **Web Framework**: [Gin](https://github.com/gin-gonic/gin)
- **Database ORM**: [GORM](https://gorm.io/)
- **Authentication**: [JWT](https://github.com/golang-jwt/jwt)
- **Configuration**: [Viper](https://github.com/spf13/viper)
- **Logging**: [Zap](https://go.uber.org/zap)
- **Template Engine**: [Goview](https://github.com/foolin/goview)
- **Database Drivers**: PostgreSQL, SQLite

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Inspired by clean architecture principles
- Built with modern Go best practices
- Templates designed for scalability and maintainability

## 📞 Support

If you have any questions or need help, please open an issue on GitHub.

---

**Happy coding with Alligotor! 🐊✨**
