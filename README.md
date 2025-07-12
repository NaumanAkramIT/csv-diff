# CSV Diff Tool 🛠️

![GitHub Release](https://img.shields.io/github/v/release/NaumanAkramIT/csv-diff?style=flat-square) ![License](https://img.shields.io/badge/license-MIT-blue.svg)

Welcome to the **CSV Diff Tool** repository! This project offers a fast and flexible command-line tool for comparing CSV files using composite keys. It generates clean, structured diffs, making it easier for you to identify changes between datasets.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Command-Line Options](#command-line-options)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Features

- **Fast Comparison**: Quickly compare large CSV files.
- **Composite Key Support**: Use multiple columns as keys for accurate comparisons.
- **Structured Output**: Get clean diffs that are easy to read.
- **Cross-Platform**: Works on Linux, macOS, and Windows.
- **Open Source**: Free to use and modify under the MIT License.

## Installation

To get started, download the latest release of the CSV Diff Tool from our [Releases page](https://github.com/NaumanAkramIT/csv-diff/releases). Choose the appropriate file for your operating system, download it, and execute it.

### For Linux and macOS

1. Open your terminal.
2. Download the binary file using `curl` or `wget`:
   ```bash
   curl -LO https://github.com/NaumanAkramIT/csv-diff/releases/download/vX.X.X/csv-diff-linux-amd64
   chmod +x csv-diff-linux-amd64
   sudo mv csv-diff-linux-amd64 /usr/local/bin/csv-diff
   ```

### For Windows

1. Download the `.exe` file from the [Releases page](https://github.com/NaumanAkramIT/csv-diff/releases).
2. Place the executable in a directory included in your system's PATH.

## Usage

Once installed, you can start using the CSV Diff Tool from your command line. The basic syntax is:

```bash
csv-diff [options] <file1.csv> <file2.csv>
```

This command will compare `file1.csv` and `file2.csv` and output the differences.

## Command-Line Options

| Option         | Description                              |
|----------------|------------------------------------------|
| `-k, --keys`   | Specify composite keys for comparison.   |
| `-o, --output` | Specify output format (e.g., JSON, CSV).|
| `-h, --help`   | Display help information.                |
| `-v, --version`| Show the version of the tool.           |

## Examples

### Basic Comparison

To compare two CSV files without any options:

```bash
csv-diff file1.csv file2.csv
```

### Using Composite Keys

If you want to compare based on specific columns, use the `-k` option:

```bash
csv-diff -k column1,column2 file1.csv file2.csv
```

### Outputting in JSON Format

To get the differences in JSON format, use the `-o` option:

```bash
csv-diff -o json file1.csv file2.csv
```

## Contributing

We welcome contributions! If you have suggestions for improvements or new features, please open an issue or submit a pull request. Make sure to follow our code of conduct and contribution guidelines.

1. Fork the repository.
2. Create your feature branch: `git checkout -b feature/MyFeature`
3. Commit your changes: `git commit -m 'Add some feature'`
4. Push to the branch: `git push origin feature/MyFeature`
5. Open a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For questions or feedback, please reach out to the maintainer:

- **Nauman Akram**: [GitHub Profile](https://github.com/NaumanAkramIT)

---

Thank you for checking out the CSV Diff Tool! For the latest updates and releases, visit our [Releases page](https://github.com/NaumanAkramIT/csv-diff/releases). Happy comparing!