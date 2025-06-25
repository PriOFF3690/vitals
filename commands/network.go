package commands

import (
	"fmt"

	"github.com/PriOFF3690/vitals/utils"

	net "github.com/shirou/gopsutil/v3/net"
)

func PrintNetworkInfo() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(utils.ColorRed + "Error getting network interface: " + err.Error() + utils.ColorReset)
		return
	}

	for _, iface := range interfaces {
		if len(iface.Addrs) == 0 {
			continue
		}

		fmt.Println(utils.ColorCyan + "Interface: " + iface.Name + utils.ColorReset)
		fmt.Println("  MAC Address: " + iface.HardwareAddr)

		for _, addr := range iface.Addrs {
			fmt.Println("  Address: " + addr.Addr)
		}
		fmt.Println()
	}

	// Network I/O Stats
	stats, err := net.IOCounters(false)
	if err == nil && len(stats) > 0 {
		stat := stats[0]
		fmt.Println(utils.ColorCyan + "Total Network I/O" + utils.ColorReset)
		fmt.Printf("  Bytes Sent: %.2f MB\n", float64(stat.BytesSent)/1e6)
		fmt.Printf("  Bytes Recv: %.2f MB\n\n", float64(stat.BytesRecv)/1e6)
	}
}
