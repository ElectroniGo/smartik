# Contributing to the Project

This document provides guidelines for contributing to this project. Please read it carefully to ensure a smooth and effective collaboration process.

## Contribution Workflow

1.  **Clone the repo** locally:
    ```bash
    git clone https://github.com/algoblue/smartik.git
    ```
    or 
    ```bash
    # RECOMMENDED
    git clone git@github.com:algoblue/smartik.git
    ```
3.  **Create a new branch** for your changes. Please use a descriptive name:
    ```bash
    git switch -c feature/your-new-feature
    ```
    or
    ```bash
    git switch -c fix/your-bug-fix
    ```
    or
    ```bash
    git switch -c chore/your-maintainance-change
    ```
4.  **Set up the development environment** by following the instructions in the [getting started](./get-started.md).
5.  **Make your changes** in your local repository.
6.  **Ensure** that everything is working correctly.
7.  **Commit your changes** using a descriptive commit message that follows our Commit Message Convention.
    > **type(scope): descriptive message**
    > ---
    > type: whether your change is a chore (no direct to the functionality of the project), feature (new functionality), fix (bug fix), etc.
    > scope: the folder name of the package your changes are meant for.
             (e.g. you added a button in `client/desktop` the scope of the commit would be `desktop`). Leave out the scope if your changes are for the global scope (in the repo's root)
    > descriptive message: you actual commit message
8.  **Push your branch** to GitHub: 
    ```bash
    git push origin feature/your-amazing-feature
    ```
9.  **Open a Pull Request** to the `main` branch of the repository.
10. **Request a review** from copilot. If you would like a fresh pair of eyes on your changes from other teammates, you can request a reveiw from them as well. Make sure you use Copilot to review your code though.

## Development Environment Setup

### Installation

1.  After cloning the repository, install the JavaScript/TypeScript dependencies:
    ```bash
    pnpm install
    ```
2.  Install the Go module dependencies:
    ```bash
    go mod tidy
    ```

### Running the Project

The easiest way to run the entire application stack is by using Docker Compose. This will build and run the Go backend, the frontend services, and any other required services.

```bash
pnpm dev
```

## Coding Style and Conventions

### Commit Message Convention

We follow the Conventional Commits specification. This makes the commit history easier to read. Each commit message should be in the format `type(scope): descriptive message`.

**Example:** `feat(api): add user authentication endpoint`