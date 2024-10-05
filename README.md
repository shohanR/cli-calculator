# cli-calculator

This project implements basic funtionalities of a calculator. It can be run in terminal.

This calculator supports the following operations :

1. Addition
2. Subtraction
3. Multiplication
4. Division


## Features

- The user can perform basic arithmetic operations.
- The calculator supports multiple operations during the session.
- At the end of the session, all valid operations performed will be displayed to the user.


## How to run

<!-- TODO: add steps to build the software and run it without Go. -->

### Prerequisites

- If you have Go installed, you can run the project directly using 'go run'.
- If you don't have Go installed, follow the steps to build the binary on a machine with Go and run it on a machine without Go installed.

### Running with Go Installed

***1. Clone the repository***

    git clone https://github.com/shohanR/cli-calculator.git
    cd your-repository

***2. Run the cli-calculator***

    go run main.go

***3. Interact with Menu and Repeat or, Exit***

- Follow the on-screen instructions to perform operations or exit the calculator.

### Running without Go Installed

If you want to run the `cli-calculator` on a machine that doesn’t have Go installed, you need to build the executable binary first.

***1. Clone the repository***

On a machine that has Go installed, clone the repository:

    git clone https://github.com/shohanR/cli-calculator.git
    cd cli-calculator

***2. Build the Project***

Run the following command to generate the binary executable:

    go build -o cli-calculator

This command will create an executable file named `cli-calculator` (or `cli-calculator.exe` on Windows).

***3. Transfer the Binary***

Once the binary is built, transfer it to the target machine where Go is not installed. You can use any file transfer method (USB, SCP, email etc.).

***4. Run the Binary***

- On **Linux** or **macOS**:

        ./cli-calculator

- On **Windows**

        cli-calculator.exe

***5. Cross-Compiling for Other Platforms***

If you need to build the binary for a different platform, you can cross-compile it using Go:

- To build for **Linux**:

        GOOS=linux GOARCH=amd64 go build -o cli-calculator-linux

- To build for **Windows**:

        GOOS=windows GOARCH=amd64 go build -o cli-calculator.exe

- To build for **macOS**:

        GOOS=darwin GOARCH=amd64 go build -o cli-calculator-mac

Transfer the corresponding binary to the target platform and run it as described in Step 4.


## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/shohanR/cli-calculator/blob/main/LICENSE) file for details.


## New Features

The final user would like to see all the operations performed during the session. When the user closes the session you have to present every operation done (one per line). You can save them as strings in a `slice` structure. For now, we're only interested in saving the valid operations.
