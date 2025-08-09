# Getting Started

Welcome! This guide will walk you through the process of setting up the **Smartik** project for local development. Smartik is a comprehensive educational platform built as a monorepo with multiple services and applications.

## Technology Stack

Smartik leverages modern technologies across different domains:

### Backend Services
- **Go v1.24** - Main API service using Echo framework
- **PostgreSQL** - Primary database with GORM for ORM
- **MinIO** - Object storage for file management
- **RabbitMQ** - Message queue for asynchronous processing

### Frontend Applications
- **Electron** - Desktop application
- **Node.js v22.15+** - JavaScript runtime and build tools

### Development & Deployment
- **Turborepo** - Monorepo management and build orchestration
- **Docker & Docker Compose** - Containerization and local development
- **PNPM** - Package manager for efficient dependency management

### Code Quality & CI/CD
- **GitHub Actions** - Automated testing and deployment
- **SonarCloud** - Code quality and security analysis
- **CodeQL** - Security vulnerability scanning

## Prerequisites

Before you begin, please ensure you have the following software installed on your system.

| Software | Version Requirement | Notes |
| :--- | :--- | :--- |
| [**Go**](https://go.dev/doc/install) | `v1.24.x` | Required for the core & API services. (If your version is not of `1.24` see the [Go managed version installation guide](https://go.dev/doc/manage-install)) |
| [**Node.js**](https://nodejs.org/download) | `v22.15 +` | The project uses `pnpm` which will automatically manage and use this version. (Preferred installation is through [nvm](https://github.com/nvm-sh/nvm?tab=readme-ov-file#installing-and-updating)) |
| [**Docker**](https://docs.docker.com/engine/install/) | `v28.x` | Required for running the application stack via Docker Compose. (Preferred installation is through install [Docker Desktop](https://docs.docker.com/desktop/)) |
| [**Air**](https://github.com/air-verse/air?tab=readme-ov-file#installation) | `v1.62.x` | (Optional) Used for live reloading in the package & serives written in Go |

## Setup

Follow these steps to get your development environment up and running.

### 1. Clone the repository

```bash
# Using HTTPS
git clone https://github.com/ElectroniGo/smartik.git
```

```bash
# Or using SSH
git clone git@github.com:ElectroniGo/smartik.git
```

```bash
# Navigate into the project directory
cd smartik
```

### 2. Install Dependencies

The project is a monorepo containing Go and JavaScript code. The following command will install all dependencies for the project using `pnpm` and run modified `install` scripts for each service.

```bash
npm install
```

This command will:
- Install Node.js dependencies across all workspace packages
- Download Go modules for the API service
- Set up Turbo for build orchestration
- Configure development tools and linting

### 3. Environment Setup

Initialize the environment variables for all services:

```bash
# Make the script executable and run it
chmod +x ./scripts/init_dotenv.sh
./scripts/init_dotenv.sh
```

This script creates `.env` files with default development settings for:
- **API Service**: Database connections, MinIO settings, server configuration
- **Desktop App**: Application-specific environment variables
- **Development Tools**: Build and testing configurations

### 4. Database Setup

Start the PostgreSQL database using Docker Compose:

```bash
# Start only the database service
docker compose up postgresdb -d

# Or start the full stack (PostgreSQL + MinIO + RabbitMQ)
docker compose up -d
```

The database will be automatically initialized with the required tables through GORM's auto-migration feature.

## Running the Project

### Quick Start with Docker Compose

The easiest way to run the entire application stack is with Docker Compose:

```bash
# Start all services (API, Database, MinIO, RabbitMQ, Text Extraction)
docker compose up

# Or run in background
docker compose up -d

# View logs
docker compose logs -f

# Stop all services
docker compose down
```

This will start:
- **PostgreSQL Database** (port 5432)
- **MinIO Object Storage** (port 9000, console: 9001)
- **RabbitMQ Message Queue** (port 5672, management: 15672)
- **Go API Service** (port 1323)
- **Text Extraction Service** (port 8001)

### Development Mode

For active development, you'll want to run services individually with hot reloading:

#### 1. Start Infrastructure Services

```bash
# Start database, MinIO, and RabbitMQ
docker compose up postgresdb minio rabbitmq -d
```

#### 2. Run API Service with Live Reload

```bash
# In the API service directory
cd services/api

# Using Air for live reloading (recommended)
air

# Or run directly with Go
go run ./cmd/api
```

#### 3. Run Desktop Application

```bash
# In the desktop app directory
cd apps/desktop

# Start development server
pnpm dev
```

#### 4. Run Text Extraction Service

```bash
# In the text extraction service directory
cd services/text-extraction

# Install Python dependencies
pip install -r requirements.txt

# Run the service
python main.py
```

### Service URLs

When all services are running, you can access:

| Service | URL | Description |
|---------|-----|-------------|
| API Service | http://localhost:1323 | REST API endpoints |
| MinIO Console | http://localhost:9001 | Object storage management (minioadmin/minioadmin) |
| RabbitMQ Management | http://localhost:15672 | Message queue monitoring (admin/password) |
| Text Extraction | http://localhost:8001 | PDF processing service |
| PostgreSQL | localhost:5432 | Database (root/password) |

### Using Turborepo Commands

The project uses Turborepo for efficient build orchestration:

```bash
# Build all services
pnpm build

# Build specific service
pnpm build --filter=api

# Run tests across all services
pnpm test

# Lint all code
pnpm lint

# Run development servers
pnpm dev
```

## Environment Configuration

> **Important**: Make sure you run all scripts from the root of the repository (where `pnpm-lock.yaml` is located).

Initialize environment files for all services:

```bash
```bash
# Makes the script executable & runs it
chmod +x ./scripts/init_dotenv.sh && ./scripts/init_dotenv.sh
```

The script will create `.env` files with sensible defaults for:
- Database connections
- MinIO object storage settings
- RabbitMQ configuration
- Service ports and URLs

## Troubleshooting

### Common Issues

#### Port Conflicts
If you encounter port conflicts, check which services are using the required ports:

```bash
# Check if ports are in use
sudo lsof -i :1323  # API service
sudo lsof -i :5432  # PostgreSQL
sudo lsof -i :9000  # MinIO
sudo lsof -i :5672  # RabbitMQ
```

#### Docker Issues
```bash
# Clean up Docker resources
docker system prune -f

# Rebuild containers
docker compose build --no-cache

# Reset volumes (will lose data)
docker compose down -v
```

#### Database Connection Issues
```bash
# Check if PostgreSQL is running
docker compose ps postgresdb

# View database logs
docker compose logs postgresdb

# Connect to database manually
docker compose exec postgresdb psql -U root -d postgres
```

#### Go Module Issues
```bash
# Clean Go module cache
go clean -modcache

# Re-download dependencies
cd services/api
go mod download
```

### Development Tips

1. **Use Air for Go Development**: Install and use Air for automatic reloading of Go services
2. **Check Service Health**: All services include health check endpoints
3. **Monitor Logs**: Use `docker compose logs -f <service-name>` to monitor specific services
4. **Database Migrations**: The API service automatically handles database migrations on startup

## Testing

Run tests across the project:

```bash
# Run all tests
pnpm test

# Run tests for specific service
pnpm test --filter=api

# Run Go tests with coverage
cd services/api
go test -v -race -coverprofile=coverage.out ./...

# Run Python tests
cd services/text-extraction
python -m pytest test_main.py -v
```

## Production Deployment

For production deployment:

```bash
# Build production images
docker compose -f compose.yaml build

# Start production services
docker compose -f compose.yaml up -d

# Monitor service health
docker compose ps
```

## What's Next?

You're all set up! Now that your environment is ready, here are some resources to help you get started:

### Essential Documentation
- [**Contributing Guide**](./contributing.md) - Guidelines for contributing to the project
- [**API Documentation**](../services/api/README.md) - Comprehensive API service documentation

### Architecture & Services
- **API Service**: RESTful backend built with Go and Echo framework
- **Desktop App**: Electron-based cross-platform application
- **Text Extraction**: Python service for PDF processing and language detection

### Development Workflow
1. **Feature Development**: Create feature branches following the naming convention
2. **Code Quality**: Pre-commit hooks ensure code quality and formatting
3. **Testing**: Automated tests run on every pull request
4. **Documentation**: Update relevant documentation with your changes

### Getting Help
- Check existing issues on GitHub
- Review service-specific README files
- Use the development Discord/Slack channels (if available)

Happy coding! 🚀
```

```bash
# Run all development servers
npm dev
```

```bash
# Run a specific app or services' development server
npm dev --filter=PACKAGE
```

> Replace `PACKAGE` with the name of the specific app or service, usually the name of the folder it's in. Otherwise, use the `name` field in that app or service's `package.json`

### For Production (Testing out all apps and services working together)

```bash
# Run all services using docker compose
npm start
```

To stop all the running services run:

```bash
# Stops all running services and keeps all their data (virtual networks, named volumes, etc.)
docker compose down
```

```bash
# Stops all running services and cleans up after them, removing all their data (virtual networks, named volumes, etc.)
docker compose down -v
```

## What's Next?

You're all set up! Now that your environment is ready, here are some other documents you might find helpful:

*   [**Architecture Overview:**](./architecture.md) To understand how the different parts of the project fit together.
*   [**Contributing Guide:**](./contributing.md) For guidelines on how to contribute to the project effectively.
