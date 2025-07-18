# Contributing to Simple IDM Manager

Thank you for your interest in contributing to Simple IDM Manager! This document provides guidelines and information for contributors.

## Getting Started

### Prerequisites

- Go 1.21 or later
- Git

### Setting Up Development Environment

1. Fork the repository on GitHub
2. Clone your fork locally:
   ```bash
   git clone https://github.com/sw3do/simple-idm-manager.git
   cd simple-idm-manager
   ```
3. Add the original repository as upstream:
   ```bash
   git remote add upstream https://github.com/sw3do/simple-idm-manager.git
   ```
4. Install dependencies:
   ```bash
   go mod download
   ```

## Development Workflow

### Making Changes

1. Create a new branch for your feature or bugfix:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. Make your changes and ensure they follow the coding standards:
   ```bash
   go fmt ./...
   go vet ./...
   ```

3. Add tests for your changes:
   ```bash
   go test ./...
   ```

4. Build and test the application:
   ```bash
   go build -o simple-idm
   ./simple-idm -version
   ```

### Commit Guidelines

- Use clear and descriptive commit messages
- Follow the conventional commit format:
  - `feat:` for new features
  - `fix:` for bug fixes
  - `docs:` for documentation changes
  - `test:` for adding tests
  - `refactor:` for code refactoring
  - `ci:` for CI/CD changes

Example:
```
feat: add support for custom user agents

Adds a new -user-agent flag that allows users to specify
a custom user agent string for HTTP requests.
```

### Submitting Changes

1. Push your changes to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```

2. Create a Pull Request on GitHub with:
   - Clear title and description
   - Reference to any related issues
   - Screenshots or examples if applicable

## Code Standards

### Go Code Style

- Follow standard Go formatting (`go fmt`)
- Use meaningful variable and function names
- Add comments for exported functions and complex logic
- Keep functions small and focused
- Handle errors appropriately

### Testing

- Write unit tests for new functionality
- Ensure all tests pass before submitting
- Aim for good test coverage
- Use table-driven tests where appropriate

### Documentation

- Update README.md if adding new features
- Update CHANGELOG.md following the format
- Add inline comments for complex code

## Reporting Issues

### Bug Reports

When reporting bugs, please include:

- Go version (`go version`)
- Operating system and version
- Steps to reproduce the issue
- Expected vs actual behavior
- Error messages or logs
- Example URLs or files (if applicable)

### Feature Requests

For feature requests, please:

- Describe the feature and its use case
- Explain why it would be beneficial
- Provide examples of how it would work
- Consider backward compatibility

## Release Process

Releases are automated through GitHub Actions:

1. Update CHANGELOG.md with new version
2. Create and push a new tag:
   ```bash
   git tag v1.0.0
   git push upstream v1.0.0
   ```
3. GitHub Actions will automatically:
   - Build binaries for multiple platforms
   - Create a GitHub release
   - Upload release assets

## Getting Help

If you need help or have questions:

- Check existing issues and discussions
- Create a new issue with the "question" label
- Reach out to maintainers

## Code of Conduct

Please be respectful and constructive in all interactions. We want to maintain a welcoming environment for all contributors.

## License

By contributing to Simple IDM Manager, you agree that your contributions will be licensed under the MIT License.