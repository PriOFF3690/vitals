package commands

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"sort"
	"strings"
	"syscall"
	"text/tabwriter"
	"time"

	"github.com/PriOFF3690/vitals/utils"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
)

func PrintSystemMonitor() {
	fs := flag.NewFlagSet("monitor", flag.ExitOnError)
	interval := fs.Int("interval", 2, "Refresh interval in seconds")
	sortBy := fs.String("sort", "cpu", "Sort process list by 'cpu', 'mem', 'pid' or 'name'")
	duration := fs.Int("duration", 0, "Monitoring duration in seconds (0 for infinite)")

	args := os.Args[2:]
	fs.Parse(args)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sig
		fmt.Println(utils.ColorRed + "\nMonitoring stopped..." + utils.ColorReset)
		os.Exit(0)
	}()

	clearTerminal()
	start := time.Now()

	for {
		clearTerminal()
		fmt.Println(utils.ColorCyan + "Vitals System Monitor (Press Ctrl+C to stop)" + utils.ColorReset)
		fmt.Println(utils.ColorCyan + "--------------------------------------------" + utils.ColorReset)
		printVitals(*sortBy, start)
		time.Sleep(time.Duration(*interval) * time.Second)

		if *duration > 0 && time.Since(start).Seconds() >= float64(*duration) {
			fmt.Printf("\n%s Monitoring completed in %d seconds.%s\n", utils.ColorGreen, *duration, utils.ColorReset)
			break
		}
	}
}

func printVitals(sortBy string, start time.Time) {
	cpuPercent, _ := cpu.Percent(0, false)
	vm, _ := mem.VirtualMemory()
	diskUsage, _ := disk.Usage("/")
	netIO, _ := net.IOCounters(false)

	fmt.Printf("%sCPU Usage:%s\t%.2f%%\n", utils.ColorGreen, utils.ColorReset, cpuPercent[0])
	fmt.Printf("%sMemory Usage:%s\t%.2f%% of %.2f GB\n", utils.ColorGreen, utils.ColorReset, vm.UsedPercent, toGB(vm.Total))
	fmt.Printf("%sDisk Usage:%s\t%.2f%% of %.2f GB\n", utils.ColorGreen, utils.ColorReset, diskUsage.UsedPercent, toGB(diskUsage.Total))

	if len(netIO) > 0 {
		rx := toMB(netIO[0].BytesRecv)
		tx := toMB(netIO[0].BytesSent)
		fmt.Printf("%sNetwork I/O:%s\tDownload: %.2f MB\tUpload: %.2f MB\n", utils.ColorGreen, utils.ColorReset, rx, tx)
	}

	fmt.Printf("%sMonitoring time:%s %.2f sec\n", utils.ColorCyan, utils.ColorReset, time.Since(start).Seconds())

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Printf("\n%sTop Processes (By %s):%s\n", utils.ColorCyan, strings.ToUpper(sortBy), utils.ColorReset)
	fmt.Println(utils.ColorCyan + "--------------------------------------------" + utils.ColorReset)
	fmt.Fprintln(w, "PID\tName\tCPU %\tMemory %")

	procs, _ := process.Processes()
	type procInfo struct {
		PID    int32
		Name   string
		CPU    float64
		Memory float32
	}
	var top []procInfo

	for _, p := range procs {
		name, err1 := p.Name()
		cpu, err2 := p.CPUPercent()
		mem, err3 := p.MemoryPercent()
		if err1 != nil || err2 != nil || err3 != nil {
			continue
		}
		top = append(top, procInfo{PID: p.Pid, Name: name, CPU: cpu, Memory: mem})
	}

	switch sortBy {
	case "mem":
		sort.Slice(top, func(i, j int) bool { return top[i].Memory > top[j].Memory })
	case "pid":
		sort.Slice(top, func(i, j int) bool { return top[i].PID > top[j].PID })
	case "name":
		sort.Slice(top, func(i, j int) bool {
			return strings.ToLower(top[i].Name) < strings.ToLower(top[j].Name)
		})
	default:
		sort.Slice(top, func(i, j int) bool { return top[i].CPU > top[j].CPU })
	}

	for i := 0; i < len(top) && i < 10; i++ {
		fmt.Fprintf(w, "%d\t%-15s\t%.1f\t%.1f\n", top[i].PID, top[i].Name, top[i].CPU, top[i].Memory)
	}
	w.Flush()
}

func toGB(bytes uint64) float64 {
	return float64(bytes) / (1024 * 1024 * 1024)
}

func toMB(bytes uint64) float64 {
	return float64(bytes) / (1024 * 1024)
}

func clearTerminal() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Print("\033[H\033[2J") // Moves cursor to top and clears screen
	}
}
