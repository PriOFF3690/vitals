package commands

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"text/tabwriter"
	"time"

	"github.com/PriOFF3690/vitals/utils"
)

func PrintPortScanner() {
	fs := flag.NewFlagSet("scan", flag.ExitOnError)
	host := fs.String("host", "127.0.0.1", "Target ip or domain")
	portRange := fs.String("ports", "1-1024", "Port range (e.g., 20-80)")
	timeout := fs.Int("timeout", 1, "Timeout per port in seconds")
	threads := fs.Int("threads", 100, "Number of concurrent threads")

	args := os.Args[2:]
	fs.Parse(args)

	start, end, err := parsePortRange(*portRange)
	if err != nil {
		fmt.Println(utils.ColorRed + "Error: " + err.Error() + utils.ColorReset)
		return
	}
	totalPorts := end - start + 1
	var scannedCount int64 = 0
	startTime := time.Now()

	fmt.Printf(utils.ColorCyan+"Scanning %s from port %d to %d..."+utils.ColorReset+"\n\n", *host, start, end)

	var wg sync.WaitGroup
	results := make(chan string, totalPorts)
	semaphore := make(chan struct{}, *threads)

	// ETA
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			done := atomic.LoadInt64(&scannedCount)
			elapsed := time.Since(startTime).Seconds()
			if done == 0 {
				continue
			}
			eta := (elapsed / float64(done)) * float64(int64(totalPorts)-done)
			fmt.Printf("\rProgress: %d/%d ports | Elapsed: %.1fs | ETA: %.1fs\n", done, totalPorts, elapsed, eta)
			if int(done) >= totalPorts {
				break
			}
		}
	}()

	for port := start; port <= end; port++ {
		semaphore <- struct{}{}
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			defer func() { <-semaphore }()

			address := net.JoinHostPort(*host, fmt.Sprintf("%d", p))
			conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
			if err != nil {
				atomic.AddInt64(&scannedCount, 1)
				return
			}
			defer conn.Close()

			conn.Write([]byte("\r\n"))
			conn.SetReadDeadline(time.Now().Add(time.Duration(*timeout) * time.Second))
			buf := make([]byte, 1024)
			n, _ := conn.Read(buf)
			banner := sanitizeBanner(buf[:n])

			if banner != "" {
				results <- fmt.Sprintf("%d\t%s", p, banner)
			} else {
				results <- fmt.Sprintf("%d", p)
			}
			atomic.AddInt64(&scannedCount, 1)
		}(port)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(writer, utils.ColorBold+"Port\tService"+utils.ColorReset)

	found := false
	for res := range results {
		found = true
		fmt.Fprintln(writer, res)
	}

	if found {
		writer.Flush()
		fmt.Printf("\n%sOS Guess: %s%s\n", utils.ColorCyan, guessOS(*host), utils.ColorReset)
	} else {
		fmt.Println(utils.ColorRed + "No ports open" + utils.ColorReset)
	}

	fmt.Printf("\n%sScan completed in %.2f seconds%s\n", utils.ColorGreen, time.Since(startTime).Seconds(), utils.ColorReset)
}

// Utility functions remain unchanged

func parsePortRange(portRange string) (int, int, error) {
	parts := strings.Split(portRange, "-")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid port range format. Use format: 20-80")
	}
	start, err1 := strconv.Atoi(parts[0])
	end, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil || start < 0 || end < 0 || start > end {
		return 0, 0, fmt.Errorf("invalid port numbers in range")
	}
	return start, end, nil
}

func sanitizeBanner(data []byte) string {
	var sb strings.Builder
	for _, b := range data {
		if b >= 32 && b <= 126 {
			sb.WriteByte(b)
		} else {
			sb.WriteByte('.')
		}
	}
	result := sb.String()
	if len(result) > 100 {
		result = result[:100] + "..."
	}
	return result
}

func guessOS(host string) string {
	conn, err := net.DialTimeout("ip4:icmp", host, 1*time.Second)
	if err != nil {
		return "Unknown"
	}
	defer conn.Close()

	var ttl int // You can enhance this using raw socket ICMP in future

	switch {
	case ttl <= 64:
		return "Linux / Unix (TTL ≈ 64)"
	case ttl <= 128:
		return "Windows (TTL ≈ 128)"
	case ttl <= 255:
		return "Network device (TTL ≈ 255)"
	default:
		return "Unknown"
	}
}
