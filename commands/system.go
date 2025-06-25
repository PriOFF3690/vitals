package commands

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/PriOFF3690/vitals/utils"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

func PrintSystemInfo() {
	// Define flags
	hostOnly := flag.Bool("host-only", false, "Show only Host info")
	cpuOnly := flag.Bool("cpu-only", false, "Show only CPU info")
	memOnly := flag.Bool("mem-only", false, "Show only Memory info")

	args := os.Args[2:]
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flag.CommandLine.BoolVar(cpuOnly, "cpu-only", false, "Show only CPU info")
	flag.CommandLine.BoolVar(memOnly, "mem-only", false, "Show only Memory info")
	flag.CommandLine.BoolVar(hostOnly, "host-only", false, "Show only Host info")
	flag.CommandLine.Parse(args)

	if *hostOnly {
		printHost()
		return
	} else if *cpuOnly {
		printCPU()
		return
	} else if *memOnly {
		printMem()
		return
	}

	// Default: all
	printHost()
	printCPU()
	printMem()
}

// Prints Host Info
func printHost() {
	fmt.Println(utils.ColorCyan + "\n[+] Host Information" + utils.ColorReset)
	hostInfo, _ := host.Info()

	fmt.Println(utils.ColorBold + "Hostname:      " + utils.ColorReset + hostInfo.Hostname)
	fmt.Println(utils.ColorBold + "OS:            " + utils.ColorReset + hostInfo.Platform + " " + hostInfo.PlatformVersion)
	fmt.Println(utils.ColorBold + "Kernel:        " + utils.ColorReset + hostInfo.KernelVersion)
	fmt.Println(utils.ColorBold + "Architecture:  " + utils.ColorReset + hostInfo.KernelArch)
	fmt.Println(utils.ColorBold + "Uptime:        " + utils.ColorReset + (time.Duration(hostInfo.Uptime) * time.Second).String())
	boot := time.Unix(int64(hostInfo.BootTime), 0)
	fmt.Println(utils.ColorBold + "Boot Time:     " + utils.ColorReset + boot.Format("2006-01-02 15:04:05"))

	users, _ := host.Users()
	fmt.Printf(utils.ColorBold+"Users Logged In:"+utils.ColorReset+" %d\n", len(users))
	for _, u := range users {
		fmt.Printf(" - %s on %s\n", u.User, u.Terminal)
	}
}

// Prints CPU Info
func printCPU() {
	fmt.Println(utils.ColorCyan + "\n[+] CPU Information" + utils.ColorReset)

	info, _ := cpu.Info()
	if len(info) > 0 {
		fmt.Println(utils.ColorBold + "Model:         " + utils.ColorReset + info[0].ModelName)
	}

	physical, _ := cpu.Counts(false)
	logical, _ := cpu.Counts(true)
	fmt.Printf(utils.ColorBold+"Physical Cores:"+utils.ColorReset+" %d\n", physical)
	fmt.Printf(utils.ColorBold+"Logical Cores: "+utils.ColorReset+" %d\n", logical)
}

// Prints Memory Info
func printMem() {
	fmt.Println(utils.ColorCyan + "\n[+] Memory Information" + utils.ColorReset)

	vm, _ := mem.VirtualMemory()
	fmt.Printf(utils.ColorBold+"Memory:        "+utils.ColorReset+"%.2f%% used (%.2f GB / %.2f GB)\n",
		vm.UsedPercent, float64(vm.Used)/1e9, float64(vm.Total)/1e9)

	swap, _ := mem.SwapMemory()
	fmt.Printf(utils.ColorBold+"Swap:          "+utils.ColorReset+"%.2f%% used (%.2f GB / %.2f GB)\n",
		swap.UsedPercent, float64(swap.Used)/1e9, float64(swap.Total)/1e9)
}
