# Term Git Pulse

![GitHub Workflow Status](https://img.shields.io/badge/build-passing-brightgreen)
![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go Version](https://img.shields.io/badge/go-%3E%3D1.20-cyan)

**Term Git Pulse** is a blazingly fast, standalone terminal dashboard and activity visualizer for local Git repositories. Built in Go, it reads directly from your `.git` logs to render beautiful contribution frequency charts, branch status summaries, recent commit velocity, and author contribution breakdowns right in your CLI.

## ✨ Key Features

- **Zero External Dependencies**: Operates purely on raw Git logs and local filesystem parsing.
- **Activity Pulse Matrix**: ASCII/ANSI heat map showing commit frequency over time.
- **Branch Overview**: Quick status of active, local, and remote-tracking branches.
- **Contributor Breakdown**: Percentage share and commit counts per author.
- **Lightning Fast**: Concurrent log parsing optimized for large Git histories.

---

## 📂 Project Structure

```text
term-git-pulse/
├── go.mod
├── main.go
└── README.md
```

---

## 🚀 Getting Started

### Prerequisites

- [Go](https://golang.org/) (version 1.20 or higher)
- Git installed in your system PATH

### Installation

Clone the repository and build the binary:

```bash
git clone https://github.com/elite-dev/term-git-pulse.git
cd term-git-pulse
go build -o term-git-pulse main.go
```

---

## 💡 Usage

Run the tool inside any valid Git repository:

```bash
./term-git-pulse
```

Or specify a custom path to a repository:

```bash
./term-git-pulse -path /path/to/your/repo
```

### Sample Output

```text
======================================================================
  🚀 TERM GIT PULSE - Repository Dashboard
======================================================================
📁 Repository: /Users/developer/projects/awesome-app
🌿 Active Branch: main
📊 Total Commits: 412
👥 Total Authors: 6

----------------------------------------------------------------------
📈 COMMIT ACTIVITY PULSE (Last 30 Days)
----------------------------------------------------------------------
[████████░░] 2023-10-14 (12 commits)
[██████████] 2023-10-15 (18 commits)
[████░░░░░░] 2023-10-16 (5 commits)
[██████░░░░] 2023-10-17 (9 commits)
[████████░░] 2023-10-18 (14 commits)
[░░░░░░░░░░] 2023-10-19 (0 commits)
[░░░░░░░░░░] 2023-10-20 (0 commits)

----------------------------------------------------------------------
🏆 TOP CONTRIBUTORS
----------------------------------------------------------------------
  1. Jane Doe           - 215 commits ( 52.2%)
  2. John Smith         - 120 commits ( 29.1%)
  3. Alice Developer    -  50 commits ( 12.1%)
  4. Bob Engineer       -  27 commits (  6.6%)

----------------------------------------------------------------------
🌿 RECENT BRANCHES
----------------------------------------------------------------------
  * main
    feature/auth-refresh
    bugfix/memory-leak
    experiment/wasm-runtime
======================================================================
```

---

## 📜 License

This project is open-source under the terms of the [MIT License](LICENSE).
