# `@smartik/api`

This directory contains the self-hosted backend infrastructure for the Smartik application, powered by Convex. It uses Docker to run the core Convex backend, a PostgreSQL database, and the Convex dashboard.

## Prerequisites

Before you begin, ensure you have the following installed on your system:

- Docker
- Docker Compose (Included with Docker Desktop)

## Quick Start

1.  **Set up Environment Variables:**
    This project requires a `.env` file for configuration. You can create one by copying the example file:
    ```bash
    cp .env.example .env
    ```
    > **Note:** The default values in `.env.example` are suitable for local development. You do not need to change them to get started.

2.  **Start the Services:**
    Run the following command from this directory (`/api`) to build and start all the services in the background.
    ```bash
    pnpm run start
    ```

3.  **Access the Services:**
    Once the containers are running, you can access the services at these default URLs:
    -   **Convex Backend:** `http://127.0.0.1:3210` (Your client application should connect to this URL)
    -   **Convex Dashboard:** `http://127.0.0.1:6791` (Use this to view your data, logs, and functions)
    
    > To access the convex dashboard you need to have an admin key.
    > 
    > To get one, first make sure you've run step 2. Then, run the following command:
    > 
    > ```bash
    > docker exec api-backend-1 bash ./generate_admin_key.sh
    > ```
    > Copy the admin key & paste it when prompted in the dashboard's login form.
    

## Configuration

All configuration is managed through the `.env` file.

| Variable | Default | Description |
| :--- | :--- | :--- |
| `POSTGRES_USER` | `convex` | The username for the PostgreSQL database. |
| `POSTGRES_PASSWORD` | `localconvexpassword` | The password for the PostgreSQL database. |
| `POSTGRES_DB` | `convex_self_hosted` | The database name. **Do not change this value.** |
| `POSTGRES_URL` | `postgresql://...` | The full connection string for the backend to connect to the database. |
| `PORT` | `3210` | The host port for the main Convex backend service. |
| `SITE_PROXY_PORT` | `3211` | The host port for the Convex site proxy. |
| `DASHBOARD_PORT` | `6791` | The host port for the Convex dashboard web interface. |

If you change the `POSTGRES_USER` or `POSTGRES_PASSWORD`, you **must** also update the `POSTGRES_URL` to match.

## Services

The `docker-compose.yaml` file defines three main services:

-   **`db`**: A PostgreSQL database instance used for all data persistence.
-   **`backend`**: The core Convex backend that runs your functions, manages state, and serves the API.
-   **`dashboard`**: A web interface for viewing your project's data, logs, schema, and functions.

### Data Persistence

The database and Convex data are persisted using Docker named volumes (`db_data` and `convex_data`). This means your data will be saved even if you stop and restart the containers.

To completely reset your local environment and delete all data, run:
```bash
pnpm run down:v
```

## Useful Commands

These commands should be run from the `/api` directory.

```bash
# Start all services in detached mode
pnpm run dev

# Stop all services
pnpm run down

# View logs for all running services
pnpm run logs

# View logs for a specific service (e.g., backend)
docker-compose logs -f backend

# Stop services and remove all data volumes
pnpm run down:v
```
