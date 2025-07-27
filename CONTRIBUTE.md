# 🧠 Contributing to Smartik

Welcome! We're glad you're here. This document outlines the conventions and expectations for contributing code to the SmarTik project. Please read this carefully before submitting a pull request.

---

## 🚀 Project Stack

- **Backend**: Go
- **Frontend**: JavaScript, React
- **Styling**: Tailwind CSS 
- **Testing**: Go test, Jest, React Testing Library

---

## ✅ Code Review Checklist

### 1. Design

- 🔹 Changes must fit into the **existing architecture** and follow established patterns.
- 🔹 Avoid mixing concerns (e.g., no business logic in React components, no UI logic in Go services).
- 🔹 Place code in the appropriate module, folder, or layer.
- 🔹 Don’t create new concepts unnecessarily.

### 2. Functionality

- 🔹 Code must work as intended. Understand the feature's goal and ensure it’s met.
- 🔹 Handle all edge cases and validate inputs.
- 🔹 Pay attention to concurrency, resource leaks, and error propagation in Go.
- 🔹 Test UI/UX changes manually or with screen recordings/demos.

### 3. Complexity

- 🔹 Functions should be **small**, focused, and follow **single responsibility**.
- 🔹 Avoid early abstraction and unnecessary generalization.
- 🔹 Refactor or extract code if a method or component grows beyond 40–50 lines.
- 🔹 Don’t sacrifice clarity for cleverness.

### 4. Test Quality

- 🔹 All new logic must have **tests** in the same pull request.
- 🔹 Tests must be deterministic and meaningful.
- 🔹 Avoid testing implementation details—test observable behavior.
- 🔹 Keep each test focused and name them descriptively.

### 5. Naming and Documentation

- 🔹 Use clear, consistent, and descriptive names.
  - ✅ `fetchUserData()`, ❌ `getInfo()`
- 🔹 Follow project casing conventions:
  - Go: `camelCase` for variables, `PascalCase` for exported
  - JS: `camelCase` for variables/functions, `PascalCase` for components
- 🔹 Comments must explain *why*, not *what*.
- 🔹 Update external documentation if builds, tests, or configs are affected.

### 6. Style and Consistency

- 🔹 Use formatters:
  - Go: `go fmt`, `golangci-lint`
  - JS: Prettier, ESLint
- 🔹 Match existing patterns and file structures.
- 🔹 Use React hooks for logic reuse; keep components lean.
- 🔹 Handle errors consistently across the codebase.
- 🔹 Mark style-only comments as `// nit:` or `// optional:` in PRs.

---

## 🧪 Testing & Linting

### Go

```sh
go fmt ./...
golangci-lint run
go test ./...
```
### React/JS

npm run lint
npm run test
npm run format

### 🗂 Project Structure Overview

### Backend (Go)

```bash
/cmd/smartik        # Application entry point
/internal/
  /handler           # HTTP handlers
  /service           # Business logic
  /repository        # Data access
  /model             # Types and structs
```
### Frontend (React)

```bash
/src
  /components        # UI components
  /pages             # Route views
  /hooks             # Custom hooks
  /services          # API interactions
  /store             # Global state
```
### 📥 Pull Request Process

Fork the repo and create a branch: feature/<name> or fix/<name>.

- Follow all style and testing guidelines.

- Include relevant tests.

- Request a review with a clear description of the changes.

- Be open to feedback!

# 💬 Communication
Please use descriptive PR titles, meaningful commit messages, and respectful review comments. Collaboration builds better software.