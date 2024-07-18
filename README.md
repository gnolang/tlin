# tlin: Lint for gno

Advance Linter for go-like grammar languages.

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/gnoswap-labs/tlin/CI?label=build)
![License](https://img.shields.io/badge/License-MIT-blue.svg)

## Introduction

tlin is an linter designed for both [Go](https://go.dev/) and [gno](https://gno.land/) programming languages. It leverages the functionality of [golangci-lint](https://github.com/golangci/golangci-lint) as its main linting engine, providing powerful code analysis for go-like grammar languages.

Inspired by Rust's [clippy](https://github.com/rust-lang/rust-clippy), tlin aims to provide additional code improvement suggestions beyond the default golangci-lint rules.

## Features

- Support for Go (.go) and Gno (.gno) files
- Ability to add custom lint rules
- Additional code improvement suggestion, such as detecting unnecessary code (🚧 WIP)

## Installation

- Requirements:
    - Go: 1.22 or higher
    - latest version of gno

To install tlin CLI, run:

```bash
go install ./cmd/tlin
```

## Usage

```bash
tlin <path>
```

Replace `<path>` with the file or directory path you want to analyze.

To check the current directory, run:

```bash
tlin .
```

## Adding Custom Lint Rules

tlin allows addition of custom lint rules beyond the default golangci-lint rules. To add a new lint rule, follow these steps:

1. Add a function defining the new rule in the `internal/rule_set.go` file.

Example:

```go
func (e *Engine) detectNewRule(filename string) ([]Issue, error) {
    // rule implementation
}
```

2. Add the new rule to the `Run` method in the `internal/lint.go` file.

```go
newRuleIssues, err := e.detectNewRule(tempFile)
if err != nil {
    return nil, fmt.Errorf("error detecting new rule: %w", err)
}
filtered = append(filtered, newRuleIssues...)
```

3. (Optional) Create a new formatter for the new rule in the `formatter` pacakge.
  a. Create a new file named after your lint rule (e.g., `new_rule.go`) in the `formatter` package.

  b. Implement the `IssueFormatter` interface for your new rule:

  ```go
  type NewRuleFormatter struct{}

  func (f *NewRuleFormatter) Format(
    issue internal.Issue,
    snippet *internal.SourceCode,
    ) string {
        // Implementation of the formatting logic for the new rule
    }
  ```

  c. Add the new formatter to the `GetFormatter` function in `formatter/fmt.go`:

  ```go
    // rule set
    const (
        // ...
        NewRule = "new_rule" // <- define the new rule as constant
    )

    func GetFormatter(rule string) IssueFormatter {
        switch rule {
        // ...
        case NewRule:
            return &NewRuleFormatter{}
        default:
            return &DefaultFormatter{}
        }
    }
    ```

4. If necessary, update the `FormatIssueWithArrow` function in `formatter/fmt.go` to handle any special formatting requirements for your new rule.

By following these steps, you can add new lint rules and ensure they are properly formatted when displayed in the CLI.


## Contributing

We welcome all forms of contributions, including bug reports, feature requests, and pull requests. Please feel free to open an issue or submit a pull request.

## License

This project is distributed under the MIT License. See `LICENSE` for more information.
