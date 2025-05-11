# CSV-DIFF

**CSV-DIFF** is a powerful and flexible CLI tool written in Go for comparing two CSV files. It highlights added,
deleted, updated, and reordered rows based on configurable composite keys. It’s useful for tracking changes in tabular
data like exports from databases, spreadsheets, or APIs.

---

## 📑 Table of Contents

- [Features](#-features)
- [Installation](#-installation)
- [Usage](#-usage)
- [Makefile](#-makefile)
- [License](#-license)
- [Contributing](#-contributing)
- [Code of Conduct](#-code-of-conduct)

---

## 📌 Features

- Compare two CSV files and identify:
    - Added, deleted, updated, or reordered rows
- Use composite keys to match rows across files
- Output results in JSON or CSV
- Filter output to focus on specific change types
- Optionally include unchanged rows
- Easy to use via CLI flags

---

## 🛠️ Installation

### Install via `go install`

Make sure you have [Go](https://golang.org/dl/) installed.

```bash
go install github.com/manishjalui11/csv-diff@latest
```

This will install csv-diff into your $GOPATH/bin.

---

## 🚀 Usage

```bash
csv-diff [flags] <original.csv> <new.csv>
```

### Example

```bash
csv-diff -k 0,1 -s json -o diff.json file_old.csv file_new.csv
```

### CLI Flags

| Flag                | Short | Description                                                                   |
|---------------------|-------|-------------------------------------------------------------------------------|
| --save              | -s    | Format to save output (json or csv)                                           |
| --output            | -o    | Filename to save changes                                                      |
| --key               | -k    | Comma-separated list of column indexes to use as a composite key (default: 0) |
| --show-changes      | -c    | Show detailed changes                                                         |
| --include-unchanged | -z    | Include unchanged rows in the output                                          |
| --filter-added      | -a    | Only show added rows                                                          |
| --filter-deleted    | -d    | Only show deleted rows                                                        |
| --filter-updated    | -u    | Only show updated rows                                                        |
| --filter-reordered  | -r    | Only show reordered rows                                                      |
| --help              | -h    | Show usage/help information                                                   |
| --version           | -v    | Show app version                                                              |

---

## 🧰 Makefile
The project includes a Makefile to simplify common development tasks.

Common Targets
```bash
make build        # Build the csv-diff binary
make test         # Run tests
make test-race    # Run tests with race detection and coverage
make run          # Build and run csv-diff with sample files
make clean        # Remove built binaries
make help         # Show available targets
```
You can always run make help to list all available targets.

---

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for more details.

---

## 🤝 Contributing

We welcome contributions from the community! Whether it's bug fixes, new features, or improvements, your help is
appreciated.

To get started:

1. **Fork the repository** and create your branch (`git checkout -b feature-branch`).
2. **Make your changes** and test them thoroughly.
3. **Commit your changes** with clear and concise commit messages.
4. **Push to your fork** (`git push origin feature-branch`).
5. **Open a Pull Request** describing your changes and why they're useful.

If you encounter any issues or have ideas for new features,
please [open an issue](https://github.com/manishjalui11/csv-diff/issues) or submit a feature request.

By contributing, you agree to follow the project's coding guidelines and adhere to our code of conduct.

Thank you for being awesome and helping improve the project! 🙌

---

## 📜 Code of Conduct

We are committed to creating a welcoming environment for all contributors. Please review and follow
our [Code of Conduct](./CODE_OF_CONDUCT.md).
