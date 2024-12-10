---
hide:
  - toc
---
forgex provides a convenient CLI tool to effortlessly set up your Go projects. Follow the steps below to install the tool on your system.

## Binary Installation

To install the forgex CLI tool as a binary, Run the following command:

```sh
go install github.com/DEVRhylme-Foundation/forgex@latest
```

This command installs the forgex binary, automatically binding it to your `$GOPATH`.


## Building and Installing from Source

If you prefer to build and install forgex directly from the source code, you can follow these steps:

Clone the forgex repository from GitHub:

```sh
git clone https://github.com/DEVRhylme-Foundation/forgex
```
   
Build the forgex binary:

```sh
go build
```
Install in your `$PATH` to make it accessible system-wide:

```sh
go install
```

Verify the installation by running:

```sh
forgex version
```

This should display the version information of the installed forgex.

Now you have successfully built and installed forgex from the source code.
