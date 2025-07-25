# Getting Started

Welcome! This guide will walk you through the process of setting up the project for local development. Following these steps will ensure you have a working environment to build, test, and contribute.

## Prerequisites

Before you begin, please ensure you have the following software installed on your system.

| Software | Version Requirement | Notes |
| :--- | :--- | :--- |
| [**Go**](https://go.dev/install) | `v1.24.x` | Required for the core & API services. (If your version is not of `1.24` see the [Go managed version installation guide](https://go.dev/mange-install)) |
| [**Node.js**](https://nodejs.org/download) | `v22.15 +` | The project uses `pnpm` which will automatically manage and use this version. (Preferred installation is through [nvm](https://github.com/nvm-sh/nvm?tab=readme-ov-file#installing-and-updating)) |
| [**Docker**](https://docs.docker.com/engine/install/) | `v28.x` | Required for running the application stack via Docker Compose. (Preferred installation is through install [Docker Desktop](https://docs.docker.com/desktop/)) |
| [**Air**](https://github.com/air-verse/air?tab=readme-ov-file#installation) | `v1.62.x` | (Optional) Used for live reloading in the package & serives written in Go |

> **Note on Node.js:** This project uses `pnpm` and is configured to enforce a specific Node.js version. When you run `pnpm install`, it will automatically download and use the correct version for you.

## Setup

Follow these steps to get your development environment up and running.

### 1. Clone the repository

```bash
# Using HTTPS
git clone https://github.com/algoblue/smartik.git
```

```bash
# Or using SSH
git clone git@github.com:algoblue/smartik.git
```

```bash
# Navigate into the project directory
cd smartik
```

### 2. Install Dependencies

The project is a monorepo containing Go and JavaScript code.
 command will install all dependencies for the project using `pnpm` and run modified `install` scripts for each service.

```bash
pnpm install
```

## Running the Project

The most straightforward way to run the entire application stack is with Docker Compose. This handles building the containers for each service and running them together in an orchestrated way.

### For Development

Look at the `README.md` of the specific service/app you would like to work on. There you will find the environment variables you need to set & the services the dev server relies on (e.g., a running postgres or redis database server)

```bash
# Run all development servers
pnpm dev
```

```bash
# Run a specific app or services' development server
pnpm dev --filter=PACKAGE
```

> Replace `PACKAGE` with the name of the specific app or service, usually the name of the folder it's in. Otherwise, use the `name` field in that app or service's `package.json`

### For Production (Testing out all app and services working together)

```bash
# Run all services using docker compose
pnpm start
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
