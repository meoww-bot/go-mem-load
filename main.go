package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// 获取内存信息的结构体
type MemInfo struct {
	Total     uint64
	Available uint64
}

// 读取并解析 /proc/meminfo 文件
func getMemInfo() (MemInfo, error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return MemInfo{}, err
	}
	defer file.Close()

	var memInfo MemInfo
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		key := fields[0][:len(fields[0])-1]
		value, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			continue
		}

		switch key {
		case "MemTotal":
			memInfo.Total = value
		case "MemAvailable":
			memInfo.Available = value
		}
	}

	if err := scanner.Err(); err != nil {
		return MemInfo{}, err
	}

	return memInfo, nil
}

func allocateAndUseMemory(sizeInMB int) []byte {
	size := sizeInMB * 1024 * 1024
	mem := make([]byte, size)
	for i := 0; i < size; i += 4096 {
		mem[i] = 0
	}
	return mem
}

func main() {

	memInfo, err := getMemInfo()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Total Memory: %d kB\n", memInfo.Total)

	percent := 0.5
	if len(os.Args) > 1 {
		if s, err := strconv.ParseFloat(os.Args[1], 32); err == nil {
			percent = s
		}
	}
	sizeInMB := int(percent * float64(memInfo.Total) / 1024)

	fmt.Printf("Allocating and using %d MB of memory\n", sizeInMB)
	mem := allocateAndUseMemory(sizeInMB)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	fmt.Println("Memory allocated and in use. Press Ctrl+C to exit.")
	for {
		select {
		case <-ticker.C:
			// 定期访问内存以确保内存保持在 used 状态
			for i := 0; i < len(mem); i += 4096 {
				mem[i] = byte(i % 256)
			}
		}
	}
}
