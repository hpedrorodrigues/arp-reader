package arpreader

import (
	"os/exec"
	"strings"
)

const (
	ipAddrIdx = 1
	hwAddrIdx = 3
)

func GetTable() (ArpTable, error) {
	output, err := exec.Command("arp", "-an").Output()
	if err != nil {
		return nil, err
	}

	var table ArpTable

	lines := strings.Split(string(output), "\n")

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		fields := strings.Fields(line)

		ipAddr := strings.TrimFunc(fields[ipAddrIdx], func(r rune) bool { return r == '(' || r == ')' })
		hwAddr := fields[hwAddrIdx]

		table = append(table, ARPEntry{
			IPAddr: ipAddr,
			HWAddr: hwAddr,
		})
	}

	return table, nil
}
