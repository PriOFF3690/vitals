package main

import (
	"fmt"
	"os"

	"github.com/PriOFF3690/vitals/commands"
	"github.com/PriOFF3690/vitals/utils"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "--help" || os.Args[1] == "-h" {
		printHelp()
		return
	}

	switch os.Args[1] {
	case "system":
		commands.PrintSystemInfo()
	case "disk":
		commands.PrintDiskInfo()
	case "network":
		commands.PrintNetworkInfo()
	case "scan":
		commands.PrintPortScanner()
	case "monitor":
		commands.PrintSystemMonitor()
	case "--version":
		fmt.Println("vitals version 1.0.0")
	default:
		fmt.Println(utils.ColorRed + "Unknown command: " + os.Args[1] + utils.ColorReset)
		printHelp()
	}
}

func printHelp() {
	fmt.Println(utils.ColorBold + utils.ColorCyan + "Vitals - Cybersecurity and System Monitoring Tool" + utils.ColorReset)
	fmt.Println("-------------------------------------------------")
	fmt.Println()

	fmt.Println("A cross-platform CLI utility to scan, monitor, and analyze system vitals")
	fmt.Println("like CPU, memory, disk, network, and open ports — essential for security")
	fmt.Println("audits and system diagnostics.")
	fmt.Println()

	fmt.Println(utils.ColorCyan + "USAGE:" + utils.ColorReset)
	fmt.Println("  vitals <command> [options]")
	fmt.Println()

	fmt.Println(utils.ColorCyan + "COMMANDS:" + utils.ColorReset)
	fmt.Println("  system      Show detailed system information")
	fmt.Println("              └─ Options:")
	fmt.Println("                 --host-only     Show only host information")
	fmt.Println("                 --cpu-only      Show only CPU details")
	fmt.Println("                 --mem-only      Show only memory usage")
	fmt.Println()
	fmt.Println("  disk        Show disk usage statistics and mounted partitions")
	fmt.Println()
	fmt.Println("  network     Display network interface details and I/O statistics")
	fmt.Println()
	fmt.Println("  scan        Perform a TCP port scan with banner grabbing")
	fmt.Println("              └─ Options:")
	fmt.Println("                 --host <ip>         Target IP or domain (default: 127.0.0.1)")
	fmt.Println("                 --ports <range>     Port range to scan (e.g., 1-1024)")
	fmt.Println("                 --timeout <sec>     Timeout per port (default: 1)")
	fmt.Println("                 --threads <num>     Concurrent scan threads (default: 100)")
	fmt.Println()

	fmt.Println("  monitor     Live system monitor with top processes and usage stats")
	fmt.Println("              └─ Options:")
	fmt.Println("                 --interval <sec>    Refresh interval in seconds (default: 2)")
	fmt.Println("                 --duration <sec>    Total monitoring duration (0 = infinite)")
	fmt.Println("                 --sort <metric>     Sort by: cpu | mem | pid | name (default: cpu)")
	fmt.Println()

	fmt.Println(utils.ColorCyan + "GLOBAL OPTIONS:" + utils.ColorReset)
	fmt.Println("  --help, -h      Show this help message")
	fmt.Println("  --version       Show version information")
	fmt.Println()

	fmt.Println(utils.ColorCyan + "EXAMPLES:" + utils.ColorReset)
	fmt.Println("  vitals system")
	fmt.Println("  vitals system --cpu-only")
	fmt.Println("  vitals scan --host 192.168.1.100 --ports 20-1000 --threads 200")
	fmt.Println("  vitals monitor --interval 1 --duration 10 --sort mem")
	fmt.Println()

	fmt.Println(utils.ColorCyan + "DOCUMENTATION:" + utils.ColorReset)
	fmt.Println("  https://github.com/PriOFF3690/vitals")
	fmt.Println()
}
