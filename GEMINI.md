# Gemini CLI Usage Guide

This document outlines how to effectively use the Gemini CLI within the MikroNet project.

## 1. General Interaction

The Gemini CLI is an interactive assistant designed to help you with various software engineering tasks. You can ask it questions, request code modifications, or seek assistance with project-related queries.

### Getting Help

To get general help or information about the Gemini CLI's features, you can use the `/help` command.

## 2. Code Analysis and Modification

Gemini CLI can assist with understanding, refactoring, and debugging code.

-   **Explaining Code**: Ask "Explain this function: `function_name` in `path/to/file.go`"
-   **Refactoring**: Request refactoring tasks, e.g., "Refactor `services/authentication/internal/controller/auth_controller.go` to improve error handling."
-   **Bug Fixing**: Describe a bug and ask for a fix, e.g., "There's a bug in `services/user/internal/service/user_service.go` where users are not being properly authenticated. Can you help fix it?"
-   **Adding Features**: Request new features, e.g., "Add a new endpoint to the authentication service for password reset."

## 3. Testing

Gemini CLI can help with running and understanding tests.

-   **Running Tests**: While Gemini CLI doesn't directly execute arbitrary commands without your confirmation, you can ask it to identify the correct test commands for a specific service or the entire project. For example, "How do I run tests for the `authentication` service?" or "Run all tests."
-   **Creating Tests**: "Create unit tests for the `auth_service.go`."

## 4. Development Workflow

Gemini CLI can integrate into your development workflow for various tasks:

-   **File Operations**: Request file creation, deletion, or modification.
-   **Dependency Management**: Inquire about or request updates to `go.mod` or `go.sum` files.
-   **Documentation**: Ask Gemini to generate or update documentation files.

## Important Notes

-   Always provide clear and specific instructions to the Gemini CLI.
-   Be prepared to confirm critical commands that modify the file system or system state.
-   The Gemini CLI will adhere to existing project conventions and coding styles.