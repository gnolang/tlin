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
- Additional code improvement suggestion, such as detecting unnecessary code
- Auto-fixing for some lint rules (🚧 WIP)
- Cyclomatic complexity analysis

## Installation

- Requirements:
  - Go: 1.22 or higher
  - latest version of gno

To install tlin CLI, follow these steps:

1. Clone the repository

```bash
git clone https://github.com/gnoswap-labs/tlin
```

2. Move to the cloned directory

```bash
cd tlin
```

3. Install the CLI

```bash
go install ./cmd/tlin
```

that's it! You can now use the `tlin` command in your terminal.

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

> ⚠️ Must update relevant tests if you have added a new rule or formatter.

1. Implement the `LintRule` interface for your new rule:

```go
type NewRule struct{}

func (r *NewRule) Check(filename string, node *ast.File) ([]types.Issue, error) {
    // Implement your lint rule logic here
    // return a slice of Issues and any error encountered
}
```

2. Register your new rule in the `registerDefaultRules` method of the `Engine` struct in `internal/engine.go`:

```go
func (e *Engine) registerDefaultRules() {
    e.rules = append(e.rules,
        &GolangciLintRule{},
        // ...
        &NewRule{}, // Add your new rule here
    )
}
```

3. (Optional) if your rule requires special formatting, create a new formatter in the `formatter` package:

   a. Create a new file (e.g., `formatter/new_rule.go`).

   b. Implement the `IssueFormatter` interface for your new rule:

   ```go
   type NewRuleFormatter struct{}

   func (f *NewRuleFormatter) Format(
       issue types.Issue,
       snippet *internal.SourceCode,
   ) string {
       // Implement formatting logic for new rule here.
   }
   ```

   c. Add the new formatter to the `GetFormatter` function in `formatter/fmt.go`.

   ```go
   func GetFormatter(rule string) IssueFormatter {
       switch rule {
       // ...
       case "new_rule": // Add your new rule here
           return &NewRuleFormatter{}
       default:
           return &DefaultFormatter{}
       }
   }
   ```

By following these steps, you can add new lint rules and ensure they are properly formatted when displayed in the CLI.

## Cyclomatic Complexity Analysis

tlin includes a feature to analyze the [cyclomatic complexity](https://en.wikipedia.org/wiki/Cyclomatic_complexity) of your code. This feature is a quantitative measure of the number of linearly independent paths through a program's source code. It's often used to indicate the complexity of a program and can be used to estimate the effort required to test or maintain a function or method.

High cyclomatic complexity can make code harder to understand and maintain. Generally, lower cyclomatic complexity is better.

### Running Cyclomatic Complexity Analysis

To run this analysis, use the `-cyclo` flag along with an optional `-threshold` flag to set the [complexity threshold](https://learn.microsoft.com/en-us/visualstudio/code-quality/code-metrics-cyclomatic-complexity?view=vs-2022). The default threshold is 10.

```bash
tlin -cyclo -threshold <threshold> <path>
```

For example:

```bash
tlin -cyclo -threshold 15 .
```

This command will analyze all `.go` and `.gno` files in the current directory and its subdirectories, reporting functions with a complexity higher than 15.

The output will show:

- The function name and its location
- The cyclomatic complexity score

Use this feature to identify complex parts of your codebase that might benefit from refactoring.

## Contributing

We welcome all forms of contributions, including bug reports, feature requests, and pull requests. Please feel free to open an issue or submit a pull request.

## Credits

- [@GodDrinkTeJAVA](https://github.com/GodDrinkTeJAVA) - Project name (`tlin`) suggestion

## License

This project is distributed under the MIT License. See `LICENSE` for more information.
