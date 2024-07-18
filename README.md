# gitgen

`gitgen` is a command-line tool developed in Go that generates commit messages and code reviews based on code changes in your project by utilizing OpenAI's ChatGPT API.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Introduction

`gitgen` is designed to assist developers in creating detailed commit messages and/or performing code reviews automatically depending on their codebase changes. By leveraging the power of ChatGPT, `gitgen` analyzes the changes made to the code and generates meaningful output.

## Features

- Generate commit messages based on code changes.
- Performs detailed code reviews.

## Installation

To get started with `gitgen`, you need to have Go installed on your machine. You can download and install Go from [here](https://golang.org/dl/).

Once Go is installed, you can clone the `gitgen` repository and build the tool:

```sh
git clone https://github.com/seymahandekli/git-gen
cd git-gen
go build ./cmd/gitgen
```

## Usage

After building `gitgen`, you can use it from the command line. Below are some example commands to help you get started:

```sh
# Generate commit message
./gitgen

# Generate code review
./gitgen --prompt review
```

For more detailed usage instructions, refer to the `--help` option:

```sh
./gitgen --help
```

## Contributing

We welcome contributions from the community! If you'd like to contribute to `gitgen`, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Make your changes and commit them with clear messages.
4. Push your changes to your fork.
5. Submit a pull request to the `main` branch of this repository.

For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

---

We hope you find `gitgen` useful! If you have any questions or feedback, please feel free to open an issue on GitHub.

Happy coding!