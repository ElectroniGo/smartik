# Project Documentation

Welcome to the **Smartik** documentation! This directory contains comprehensive guides and references to help you understand, set up, and contribute to the Smartik educational platform.

## 📋 Table of Contents

### Getting Started
- [**Getting Started Guide**](./get-started.md) - Complete setup instructions for local development, including technology stack overview and troubleshooting

### Development & Contributing  
- [**Contributing Guidelines**](./contributing.md) - Code standards, workflow, and contribution process

### Service Documentation
- [**API Service Documentation**](../services/api/README.md) - Comprehensive Go API service documentation with endpoints, models, and examples
- [**Desktop Application**](../apps/desktop/README.md) - Electron desktop application documentation

## 🏗️ Architecture Overview

Smartik is built as a **monorepo** with multiple interconnected services:

### Backend Services
- **Go API Service** - RESTful API using Echo framework, PostgreSQL, and GORM
- **Text Extraction Service** - Python service for PDF processing and language detection
- **Database** - PostgreSQL with automatic migrations
- **Object Storage** - MinIO for file management
- **Message Queue** - RabbitMQ for asynchronous processing

### Frontend Applications  
- **Desktop Application** - Cross-platform Electron app
- **Web Interface** - (Future: React/Next.js web application)

### Infrastructure
- **Docker Compose** - Local development and production deployment
- **Turborepo** - Monorepo management and build orchestration
- **CI/CD Pipeline** - GitHub Actions with automated testing and deployment

## 🚀 Quick Start

1. **Prerequisites**: Go v1.24, Node.js v22.15+, Docker v28.x
2. **Clone & Setup**:
   ```bash
   git clone https://github.com/ElectroniGo/smartik.git
   cd smartik
   pnpm install
   chmod +x ./scripts/init_dotenv.sh && ./scripts/init_dotenv.sh
   ```
3. **Start Development**:
   ```bash
   docker compose up -d  # Start infrastructure
   pnpm dev             # Start all development servers
   ```

For detailed instructions, see the [Getting Started Guide](./get-started.md).

## 📚 Key Features

- **Student Management** - Comprehensive student data management
- **Exam System** - Digital exam creation and management  
- **Answer Script Processing** - Automated PDF processing with text extraction
- **Memorandum Management** - Educational content organization
- **Multi-language Support** - Automatic language detection for documents
- **Cross-platform Desktop App** - Native desktop experience

## 🔧 Development Tools

- **Live Reloading** - Air for Go services, built-in for Node.js
- **Code Quality** - Automated linting, formatting, and pre-commit hooks  
- **Testing** - Comprehensive test suites with coverage reporting
- **Security Analysis** - SonarCloud and CodeQL integration
- **Documentation** - Automated documentation generation

## 🛠️ Service URLs (Development)

| Service | URL | Purpose |
|---------|-----|---------|
| API Service | http://localhost:1323 | REST API endpoints |
| MinIO Console | http://localhost:9001 | File storage management |
| RabbitMQ Management | http://localhost:15672 | Message queue monitoring |
| Text Extraction API | http://localhost:8001 | PDF processing service |
| PostgreSQL | localhost:5432 | Database access |

## 📖 Additional Resources

- **GitHub Repository**: [ElectroniGo/smartik](https://github.com/ElectroniGo/smartik)
- **Issues & Bug Reports**: Use GitHub Issues for tracking
- **Service-specific READMEs**: Each service has detailed documentation in its directory

---

💡 **Tip**: Start with the [Getting Started Guide](./get-started.md) for your first setup, then refer to service-specific documentation for detailed development information.