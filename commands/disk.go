package commands

import (
	"fmt"

	"github.com/PriOFF3690/vitals/utils"

	"github.com/shirou/gopsutil/disk"
)

func PrintDiskInfo() {
	partitions, err := disk.Partitions(false)
	if err != nil {
		fmt.Println(utils.ColorRed + "Error reading partitions: " + err.Error() + utils.ColorReset)
		return
	}

	for _, p := range partitions {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			continue
		}
		fmt.Printf(utils.ColorCyan+"Mountpoint: %s"+utils.ColorReset+"\n", usage.Path)
		fmt.Printf("  %sTotal:%s %.2f GB\n", utils.ColorGreen, utils.ColorReset, float64(usage.Total)/1e9)
		fmt.Printf("  %sUsed:%s  %.2f GB (%.2f%%)\n", utils.ColorGreen, utils.ColorReset, float64(usage.Used)/1e9, usage.UsedPercent)
		fmt.Printf("  %sFree:%s  %.2f GB\n\n", utils.ColorGreen, utils.ColorReset, float64(usage.Free)/1e9)
	}
}
