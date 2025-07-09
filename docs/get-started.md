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

> **IMPORTANT!** Open the the project folder in VSCode. A popup should appear in the bottom right, prompting you to install the recommended extensions for this workspace. **YES, DO IT**.
>
> Those extension will help maintain great code quality and consistent style (Biomejs, Golang extensions). Others will will make it easier to work with docker and go.
>
> **Golang** has more tools that need to be installed to maintain consistent code style. Once VSCode is open:
> 1. Open the command pallete (CTRL + SHIFT + P)
> 2. Type "go install/update"
> 3. Choose the first option
> 3. Wait for a popup that will ask you to select which tools to install. and select all of them.
> 4. The output tab will show up at the bottom logging the progress. Wait until you see  message that says all tools have been install successfully.

### 2. Install Dependencies

The project is a monorepo containing Go, and JavaScript/TypeScript code.

This command will install all dependencies for the entire project using `pnpm` and run modified `install` scripts for each service that isn't JS/TS.

```bash
# Install the build tool (Turbo)
npm install --global turbo@2
```

```bash
# Install project dependencies
pnpm install
```

## Running the Project

After installing all necassary dependencies, running this command from the root will run all dev servers of the available servers. Refer to each service's `README.md` for the opened ports and availabe services.

```bash
pnpm dev
```

*Running this command will run the [`docker-compose.yaml`](../docker-compose.yaml) files that which will build docker containers for all available services in this project and run them as though they are in a production envrionment* (Coming Soon)

```bash
pnpm start
```

Docker creates certain assets for the system to run all together while mimicing a aproduction environment, these include volumes and networks.

Should you want to stop the run "production" instance run this command:

```bash
# Run this to stop servers but keep all its assets
docker compose down
```

or

```bash
# Run this to stop servers and discard all its assets as well
docker compose docker -v
```

Once either command finishes, the application services will be running and accessible on their configured ports. You can view logs for all services directly in your terminal.

## What's Next?

You're all set up! Now that your environment is ready, here are some other documents you might find helpful:

*   [**Architecture Overview:**](./architecture.md) To understand how the different parts of the project fit together.
*   [**Contributing Guide:**](./contributing.md) For guidelines on how to contribute to the project effectively.

