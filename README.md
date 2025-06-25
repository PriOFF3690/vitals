
# 🩺  Vitals

**Vitals** is a lightweight, cross-platform command-line tool built in Go for system monitoring and basic cybersecurity diagnostics. Inspired by how vitals are used to assess health, this tool helps analyze the digital "health" of your machine.

---


## 🔍 Features

- View system info (hostname, OS, CPU, memory)
- Display disk usage and partitions
- Show network interfaces and open connections
- Perform basic security scans
- Monitor system vitals continuously
- Color-coded terminal output for better visibility

---

## 📦 Releases

Download pre-built binaries from the [Releases](https://github.com/PriOFF3690/vitals/releases) page:

- 🪟 [Windows (64-bit)](https://github.com/PriOFF3690/vitals/releases/download/v1.0.0/vitals-windows-amd64.zip)
- 🐧 [Linux (64-bit)](https://github.com/PriOFF3690/vitals/releases/download/v1.0.0/vitals-linux-amd64.zip)
- 🍎 [macOS (Intel)](https://github.com/PriOFF3690/vitals/releases/download/v1.0.0/vitals-darwin-amd64.zip)

## 🛠 Installation

You have several options to get started with Vitals:

### 1. Install via go install (requires Go installed)

```bash
  go install github.com/PriOFF3690/vitals@latest
```
- Installs the vitals binary to your Go bin directory (e.g. $HOME/go/bin)
- Add that directory to your PATH to run vitals globally

### 2. Build from Source

```bash
# Clone the repository
git clone https://github.com/PriOFF3690/vitals.git
cd vitals

# Build the binary (Go must be installed)
go build -o vitals

# (Optional) Install it globally
sudo mv vitals /usr/local/bin/
```
    
### 3. Download Prebuilt Binary

Visit: [vitals Releases](https://github.com/PriOFF3690/vitals/releases)

Download the correct binary zip for your system and extract it. Example for Linux:
```bash
chmod +x vitals
./vitals --help
```
For Windows users (PowerShell):
```powershell
.\vitals.exe --help
```
## 🚀 Usage
Run the tool using:

```bash
vitals [command] [flags]
```


## 📘 Documentation

Available Commands
| Command   | Description                                   |
| --------- | --------------------------------------------- |
| `system`  | Display host information (OS, CPU, RAM, etc.) |
| `disk`    | Show disk usage and partitions                |
| `network` | List network interfaces and open connections  |
| `scan`    | Perform basic security checks (expandable)    |
| `monitor` | Monitor system vitals (CPU & memory live)     |

Use --help with any command for more info:
```bash
vitals system --help
vitals scan --help
```
## 📜 License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.



## 👨‍💻 Author
### Prince Gokhale
Cybersecurity Enthusiast & passionate techie

- [@LinkedIn](https://www.linkedin.com/in/prince-g-7262b123a)
- [@TryHackMe](https://tryhackme.com/p/shodan2109)
