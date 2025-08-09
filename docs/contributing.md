# Contributing to Smartik

Welcome! We appreciate your interest in contributing to **Smartik**. This document provides comprehensive guidelines to ensure a smooth and effective collaboration process.

## 🚀 Getting Started

Before contributing, ensure you have:

1. **Read the [Getting Started Guide](./get-started.md)** - Set up your local development environment
2. **Reviewed the [Architecture Overview](./README.md#-architecture-overview)** - Understand the project structure
3. **Installed all [prerequisites](./get-started.md#prerequisites)** - Go, Node.js, Docker, etc.

## 🔄 Contribution Workflow

### 1. Fork and Clone

```bash
# Fork the repository on GitHub, then clone your fork
git clone https://github.com/YOUR_USERNAME/smartik.git
cd smartik

# Add upstream remote
git remote add upstream https://github.com/ElectroniGo/smartik.git
```

### 2. Create a Feature Branch

```bash
# Create and switch to a feature branch
git checkout -b feat/your-feature-name

# Or for bug fixes
git checkout -b fix/issue-description

# Or for documentation
git checkout -b docs/update-readme
```

### 3. Make Your Changes

- Write clean, readable code
- Include relevant tests
- Update documentation as needed
- Follow the coding standards below

### 4. Test Your Changes

```bash
# Run all tests
pnpm test

# Run linting
pnpm lint

# Test specific service
pnpm test --filter=api

# Start services to test integration
docker compose up -d
```

### 5. Commit and Push

```bash
# Stage your changes
git add .

# Commit with conventional commit format
git commit -m "feat: add PDF text extraction service"

# Push to your fork
git push origin feat/your-feature-name
```

### 6. Create Pull Request

- Create a pull request from your fork to the main repository
- Use a clear, descriptive title
- Fill out the pull request template completely
- Reference any related issues

## 📝 Coding Standards

### Code Quality

**Write readable code** - Your code should be self-documenting, but include comments where necessary to explain:
- Complex business logic
- API integrations
- Configuration decisions
- Workarounds or temporary solutions

**Example:**
```go
// ProcessAnswerScript handles PDF upload and triggers async text extraction
// It validates file format, stores in MinIO, and queues for processing
func (h *AnswerScriptHandler) ProcessAnswerScript(c echo.Context) error {
    // Implementation with clear variable names and logic flow
}
```

### Language-Specific Guidelines

#### Go Services
- Follow [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use `gofmt` for formatting
- Use meaningful variable and function names
- Handle errors appropriately, don't ignore them

```go
// Good: Clear error handling
result, err := someOperation()
if err != nil {
    return fmt.Errorf("failed to perform operation: %w", err)
}

// Bad: Ignoring errors
result, _ := someOperation()
```

#### JavaScript/TypeScript
- Use ESLint and Prettier configurations provided
- Prefer `const` and `let` over `var`
- Use async/await over Promise chains
- Write unit tests for components and utilities

#### Python Services
- Follow PEP 8 style guide
- Use type hints where appropriate
- Write docstrings for functions and classes
- Use meaningful variable names

### Database Guidelines

- Always use migrations for schema changes
- Include both `up` and `down` migration paths
- Test migrations on a copy of production data
- Document any breaking schema changes

## 🎯 Commit Message Convention

We follow [Conventional Commits](https://www.conventionalcommits.org/) specification:

### Format
```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

### Types
- **feat**: New features
- **fix**: Bug fixes
- **docs**: Documentation updates
- **style**: Code formatting (no logic changes)
- **refactor**: Code refactoring
- **test**: Adding or updating tests
- **chore**: Maintenance tasks
- **ci**: CI/CD changes

### Examples
```bash
# New feature
git commit -m "feat(api): add student registration endpoint"

# Bug fix
git commit -m "fix(desktop): resolve memory leak in file upload"

# Documentation
git commit -m "docs: update API authentication examples"

# Breaking change
git commit -m "feat(api)!: migrate to new authentication system

BREAKING CHANGE: The authentication endpoint has changed from /auth/login to /api/v2/auth/login"
```

## 📚 Documentation Guidelines

### When to Update Documentation

Update documentation when you:
- Add new features or APIs
- Change existing functionality
- Modify configuration or environment variables
- Update deployment processes
- Fix bugs that weren't documented

### Documentation Standards

1. **README Files**: Each service should have a comprehensive README
2. **API Documentation**: Document all endpoints with examples
3. **Code Comments**: Explain complex business logic
4. **Architecture Decisions**: Document significant architectural choices

### Documentation Structure
```
docs/
├── README.md              # Overview and quick start
├── get-started.md         # Setup instructions  
├── contributing.md        # This file
└── service-docs/
    ├── api-reference.md   # API endpoints
    └── deployment.md      # Production deployment
```

## 🧪 Testing Guidelines

### Test Coverage Requirements
- **Go services**: Minimum 80% coverage for new code
- **JavaScript/TypeScript**: Unit tests for all utilities and components
- **Python services**: Test critical functions and API endpoints

### Testing Best Practices
```go
// Table-driven tests for Go
func TestProcessPDF(t *testing.T) {
    tests := []struct {
        name     string
        input    []byte
        expected string
        wantErr  bool
    }{
        {"valid PDF", validPDFBytes, "extracted text", false},
        {"invalid PDF", invalidBytes, "", true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := ProcessPDF(tt.input)
            // Assertions...
        })
    }
}
```

## 🔍 Code Review Process

### What We Look For
- **Functionality**: Does the code work as intended?
- **Performance**: Are there any performance implications?
- **Security**: Does the code introduce security vulnerabilities?
- **Maintainability**: Is the code easy to understand and modify?
- **Testing**: Are there adequate tests?
- **Documentation**: Is the change properly documented?

### Review Checklist
- [ ] Code follows project conventions
- [ ] Tests pass and provide adequate coverage
- [ ] Documentation is updated
- [ ] No security vulnerabilities introduced
- [ ] Performance impact is acceptable
- [ ] Breaking changes are properly documented

## 🐛 Reporting Issues

### Bug Reports
Include:
- Steps to reproduce
- Expected vs actual behavior
- Environment details (OS, Docker version, etc.)
- Relevant logs and error messages

### Feature Requests
Include:
- Clear description of the proposed feature
- Use cases and benefits
- Proposed implementation approach (if any)

## 🚀 Release Process

### Version Numbering
We use [Semantic Versioning](https://semver.org/):
- **MAJOR**: Breaking changes
- **MINOR**: New features (backward compatible)
- **PATCH**: Bug fixes (backward compatible)

### Release Schedule
- **Patch releases**: As needed for critical bugs
- **Minor releases**: Monthly or when significant features are ready
- **Major releases**: Planned with advance notice for breaking changes

## 🤝 Community Guidelines

### Code of Conduct
- Be respectful and inclusive
- Focus on constructive feedback
- Help newcomers get started
- Celebrate contributions of all sizes

### Getting Help
- Check existing documentation first
- Search existing issues and discussions
- Ask questions in pull requests or issues
- Join community discussions

---

<div align="center">
  
**Happy Contributing! 🎉**

*Every contribution, no matter how small, makes Smartik better for everyone.*

</div>