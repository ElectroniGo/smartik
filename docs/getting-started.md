# Getting Started

This guide will walk you through the process of setting up your development environment for Smartik.

## Prerequisites

Before you begin, ensure you have the following installed:

### [Node.js](https://nodejs.org/)

We're using Node.js as the tooling that the entire solution will be built on top of.

- **version:** 22.15.x (See the [.nvmrc](../.nvmrc) version file)
   > **NOTE:** Any version above 20.18 should work just fine though. We'll accept that.

### [PNPM](https://pnpm.io/)

We're using the PNPM package manager to govern our project's dependencies. This is just like the regular npm. they operate the same.
PNPM just approaches dependency management for multi-package repos a lot more cleanly than npm, I'd say. _(creates a dedicated `node_modules` folder for every package that has its own `package.json`)_

To enable/install it, run:

```bash
corepack enable pnpm
```

```bash
corepack prepare pnpm@10.9.0 # version must be exactly this or higher to avoid compatibilty with Turborepo
```

If your get an error setting up using `corepack`, you can alternatively just use npm:

```bash
npm install -g pnpm@10.9.0
```

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/algoblue/smartik.git
   ```

   Or, if you've setup your SSH keys (preferrable)

   ```bash
   git clone git@github.com:algoblue/smartik.git
   ```

   ```bash
   cd smartik
   ```

1. **Install dependencies:**

   ```bash
   pnpm install
   ```

## Running the Project

To run the development server, use the following command:

```bash
pnpm dev
```

> _Refer to the `README.md`'s in each of the available apps/packages to see what ports are open when the dev server is running._
