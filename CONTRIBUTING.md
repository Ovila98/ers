# Contributing to ers

Thank you for considering contributing to `ers`! I welcome all types of contributions, including bug fixes, new features, documentation improvements, and more. This guide will help you get started with contributing to the project.

## Table of Contents

1. [How to Contribute](#how-to-contribute)
   - [Reporting Bugs](#reporting-bugs)
   - [Feature Requests](#feature-requests)
   - [Code Contributions](#code-contributions)
2. [Development Workflow](#development-workflow)
3. [Coding Standards](#coding-standards)
4. [Pull Request Guidelines](#pull-request-guidelines)
5. [License](#license)

---

## How to Contribute

There are several ways you can contribute to `ers`:

### Reporting Bugs

If you find a bug, please create an issue on the [GitHub Issues page](https://github.com/ovila98/ers/issues) with the following details:

- **Description**: A clear and concise description of the bug.
- **Steps to Reproduce**: Detailed steps on how to reproduce the issue.
- **Expected Behavior**: What you expected to happen.
- **Actual Behavior**: What actually happened.
- **Environment**: The version of Go and any other environment details (operating system, etc.) that might be relevant.

### Feature Requests

If you have an idea for a new feature or improvement, I encourage you to:

1. **Search existing issues** to see if there’s already a similar request. If so, feel free to contribute to the discussion there.
2. **Open a new issue** if it hasn’t been requested yet. Describe the feature in detail and why you think it would be useful.

### Code Contributions

I welcome all code contributions, whether it’s fixing bugs, adding new features, improving documentation, or optimizing performance.

Before starting on a code contribution, it’s a good idea to first open an issue to discuss the changes you’re proposing. This will prevent duplicated effort and ensure that your changes align with the project's goals.

## Development Workflow

1. **Fork the repository**: Start by forking the `ers` repository on GitHub to your account.
2. **Clone your fork**:

   ```bash
   git clone https://github.com/ovila98/ers.git
   cd ers
   ```

3. **Create a new branch** for your feature or bugfix:

   ```bash
   git checkout -b feature/my-new-feature
   ```

4. **Make your changes**: Write clean and well-documented code. Be sure to include comments where necessary.

5. **Test your changes**: Ensure that your changes work correctly by running the tests. You can write additional tests if your changes introduce new functionality.

6. **Commit your changes**:

   ```bash
   git commit -am "Add feature/fix description"
   ```

7. **Push your branch** to GitHub:

   ```bash
   git push origin feature/my-new-feature
   ```

8. **Create a Pull Request**: Go to the original repository and submit a pull request with a description of your changes.

## Coding Standards

- Follow Go's standard code style and best practices.
- Ensure that functions, variables, and packages have meaningful names.
- Write [Go doc](https://blog.golang.org/godoc-documenting-go-code) comments for all public methods and functions.
- Keep your code clean and modular. Avoid long, complex functions.
- All contributions should be compatible with Go's latest stable version.

## Pull Request Guidelines

- Ensure that your changes do not break any existing functionality.
- All pull requests should be associated with an open issue (if applicable). Mention the issue in your pull request.
- Provide a clear description of what the pull request does and why it's needed.
- Keep your pull requests focused; avoid unrelated changes or large feature sets bundled into a single PR.
- Ensure all code passes linting and testing before submitting.

## License

By contributing, you agree that your contributions will be licensed under the same license as the project: **Apache License 2.0**.
