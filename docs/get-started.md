# Getting Started

Welcome! This guide will walk you through the process of setting up the project for local development. Following these steps will ensure you have a working environment to build, test, and contribute.

## Prerequisites

Before you begin, please ensure you have the following software installed on your system.

| Software | Version Requirement | Notes |
| :--- | :--- | :--- |
| **Go** | `v1.x` | `1.21` or newer is recommended. Required for the core service. |
| **Python** | `v3.8` or newer | Required for the API |
| **Node.js** | `v22.15.1` | The project uses `pnpm` which will automatically manage and use this version. |
| **Git** | `v2.48.x` | Any recent version of Git should work fine. |
| **Docker** | `v28.x` | Required for running the application stack via Docker Compose. |

> **Note on Node.js:** This project uses `pnpm` and is configured to enforce a specific Node.js version. When you run `pnpm install`, it will automatically download and use the correct version for you.

## Setup

Follow these steps to get your development environment up and running.

### 1. Clone the Repository

First, clone the project to your local machine. If you plan to contribute, it's best to fork the repository first and clone your fork.

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

The project is a monorepo containing Go, Python, and JavaScript code.

This command will install all dependencies for the entire project services using `pnpm` and run modified `install` scripts for each service.

```bash
pnpm install
```

## Running the Project

The most straightforward way to run the entire application stack is with Docker Compose. This handles building the containers for each service and running them together in an orchestrated way.

```bash
# This command will all development servers
pnpm dev
```

```bash
# This command will build the images if they don't exist and start all services in production mode locally
pnpm start
```

Once either command finishes, the application services will be running and accessible on their configured ports. You can view logs for all services directly in your terminal.

To stop all the running services run:

```bash
# This command will stop all running services but still keep all their data, virtual networks, etc.
docker compose down
```

```bash
# This command will stop all running services and clean up after them, removing all their data, virtual networks, etc.
docker compose down -v
```

## What's Next?

You're all set up! Now that your environment is ready, here are some other documents you might find helpful:

*   [**Architecture Overview:**](./architecture.md) To understand how the different parts of the project fit together.
*   [**Contributing Guide:**](./contributing.md) For guidelines on how to contribute to the project effectively.
*   [**Deployment Guide:**](./deployment.md) For guidelines on the deployment of the project using docker compose.

