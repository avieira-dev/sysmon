 <div align="center">
    <h1>sysmon</h1>
    <p>A Linux system monitoring tool written in Go, designed to collect and analyze system resource information.</p>
    <p>
        <img src="https://img.shields.io/badge/status-in%20development-yellow?style=flat-square" />
        <img src="https://img.shields.io/badge/platform-Linux-FCC624?style=flat-square&logo=linux&logoColor=black" />
        <img src="https://img.shields.io/badge/language-Go-00ADD8?style=flat-square&logo=go&logoColor=white" />
        <img src="https://img.shields.io/badge/interface-CLI-1f6feb?style=flat-square" />
    </p>
</div>

---

## Overview

**Sysmon** is a Linux system monitoring tool developed in Go. Its goal is to collect, process, and present system resource information through a lightweight command-line interface.

By combining low-level system information with Go's capabilities, Sysmon aims to provide a foundation for monitoring system resources and analyzing system activity.

> [!WARNING]  
> **Sysmon is currently under active development.**

The project is in its early stages. Features, architecture, and interfaces may change as development progresses. Only the functionality explicitly documented as implemented should be considered available.

---

## Table of Contents

- [Features](#features) 
- [Installation](#installation)
    - [Requirements](#requirements)
    - [Steps](#steps)
        - [Clone](#1-clone-the-repository)
        - [Build](#2-build-the-project)
- [Usage](#usage)
- [CPU Statistics](#cpu-statistics)
- [Running Tests](#running-tests)
- [Project Structure](#project-structure)
- [Roadmap](#roadmap)
- [Author](#author)
- [License](#license)

---

## Features

- **CPU Statistics Collection:** Reads aggregate CPU time counters from the Linux /proc/stat interface.
- **CPU Data Parsing:** Extracts and converts the ten CPU counters into a structured Go representation.
- **Input Validation:** Checks CPU identifiers, field quantities, numeric values, and integer overflow.
- **Error Handling:** Provides contextual errors when reading or parsing CPU statistics fails.
- **Unit Testing:** Includes automated tests for valid input and invalid parsing scenarios.
- **Lightweight Architecture:** Separates system information reading, data parsing, and application entry points.

---

## Installation

### Requirements

- Linux-based operating system.
- Go 1.26.8 or later.
- Git.

### Steps

#### 1. Clone the repository

```bash
git clone https://github.com/avieira-dev/sysmon.git
cd sysmon
```

#### 2. Build the project

```bash
go build -o sysmon ./cmd/sysmon
```

> [!NOTE]  
> The executable will be generated in the project root directory.

---

## Usage

During the current development stage, Sysmon can be executed directly using Go:

```bash
go run ./cmd/sysmon
```

Alternatively, build the project and run the generated executable:

```bash
go build -o sysmon ./cmd/sysmon
./sysmon
```

The current implementation reads and displays the aggregate CPU time counters available in `/proc/stat`.

> [!NOTE]  
> Starting with the first release, installation instructions will be provided to make Sysmon available as a system-wide command, allowing it to be executed from any directory in the terminal.

---

## CPU Statistics

| Field       | Description                                                                                        |
|:------------|:---------------------------------------------------------------------------------------------------|
| `User`      | Time spent executing user-space processes.                                                         |
| `Nice`      | Time spent executing user-space processes with adjusted nice priority.                             |
| `System`    | Time spent executing kernel code.                                                                  |
| `Idle`      | Time the CPU spent idle.                                                                           |
| `IOWait`    | Time the CPU spent idle while I/O was outstanding.                                                 |
| `IRQ`       | Time spent processing hardware interrupts.                                                         |
| `SoftIRQ`   | Time spent processing software interrupts.                                                         |
| `Steal`     | Time a virtual CPU was unable to run because the hypervisor was executing another virtual machine. |
| `Guest`     | Time spent running virtual CPUs for guest operating systems.                                       |
| `GuestNice` | Time spent running niced guest virtual CPUs.                                                       |

> [!IMPORTANT]  
> These values represent accumulated CPU time counters, not CPU usage percentages.

---

## Running Tests

Run all project tests with:

```bash
go test ./...
```

To display detailed test execution:

```bash
go test ./internal/procfs -v
```

The current test suite covers:

- Valid CPU statistics input.
- Invalid CPU identifiers.
- Insufficient input fields.
- Non-numeric CPU values.
- Values exceeding the `uint64` limit.

---

## Project Structure

```text
sysmon/
├── cmd/
│   └── sysmon/
│       └── main.go
├── internal/
│   └── procfs/
│       ├── cpu.go
│       ├── parser_test.go
│       ├── parser.go
│       └── reader.go
├── .gitignore
├── go.mod
├── LICENSE
└── README.md
```

---

## Roadmap

The roadmap represents the current development vision for Sysmon. Planned features and priorities may change as the project evolves, and new features may be added over time.

| Feature                                              | Status                      |
|------------------------------------------------------|-----------------------------|
| **CPU Monitoring**                                   |                             |
| Linux `/proc/stat` CPU data collection               | ████████████████████ `100%` |
| CPU statistics parsing and validation                | ████████████████████ `100%` |
| Automated parser unit tests                          | ████████████████████ `100%` |
| CPU utilization calculation                          | ░░░░░░░░░░░░░░░░░░░░ `0%`   |
| Per-core CPU monitoring                              | ░░░░░░░░░░░░░░░░░░░░ `0%`   |
| **Memory Monitoring**                                |                             |
| RAM usage monitoring                                 | ░░░░░░░░░░░░░░░░░░░░ `0%`   |
| Swap usage monitoring                                | ░░░░░░░░░░░░░░░░░░░░ `0%`   |
| Memory usage statistics                              | ░░░░░░░░░░░░░░░░░░░░ `0%`   |
| **Storage Monitoring**                               |                             |
| Disk usage monitoring                                | ░░░░░░░░░░░░░░░░░░░░ `0%`   |
| Filesystem information                               | ░░░░░░░░░░░░░░░░░░░░ `0%`   |
| Disk I/O statistics                                  | ░░░░░░░░░░░░░░░░░░░░ `0%`   |
| **Network Monitoring**                               |                             |
| Network interface information                        | ░░░░░░░░░░░░░░░░░░░░ `0%`   |
| Network traffic monitoring                           | ░░░░░░░░░░░░░░░░░░░░ `0%`   |
| Network throughput statistics                        | ░░░░░░░░░░░░░░░░░░░░ `0%`   |
| **Process Monitoring**                               |                             |
| Process listing                                      | ░░░░░░░░░░░░░░░░░░░░ `0%`   |
| Per-process CPU usage                                | ░░░░░░░░░░░░░░░░░░░░ `0%`   |
| Per-process memory usage                             | ░░░░░░░░░░░░░░░░░░░░░ `0%`  |
| Process information and details                      | ░░░░░░░░░░░░░░░░░░░░ `0%`   |
| **Terminal Interface**                               |                             |
| Interactive terminal interface                       | ░░░░░░░░░░░░░░░░░░░░ `0%`   |
| Real-time monitoring and automatic refresh           | ░░░░░░░░░░░░░░░░░░░░ `0%`   |
| Customizable display options                         | ░░░░░░░░░░░░░░░░░░░░ `0%`   |
| **Distribution and Configuration**                   |                             |
| Global command-line installation                     | ░░░░░░░░░░░░░░░░░░░░ `0%`   |
| Configuration file support                           | ░░░░░░░░░░░░░░░░░░░░ `0%`   |
| Command-line arguments and flags                     | ░░░░░░░░░░░░░░░░░░░░ `0%`   |
| Cross-distribution compatibility                     | ░░░░░░░░░░░░░░░░░░░░ `0%`   |

> [!NOTE]  
> The roadmap is subject to change as the project evolves. Features and topics may be added, modified, or removed at any time.

---

## Author

**Alexandre Vieira**  
GitHub: [@avieira-dev](https://github.com/avieira-dev)

---

## License

Distributed under the [MIT License](LICENSE). See `LICENSE` for details.