package arpreader

import (
	"bufio"
	"os"
	"strings"
)

const (
	ipAddr int = iota
	hwType
	flags
	hwAddr
	mask
	device
)

const arpFile = "/proc/net/arp"

func GetTable() (ArpTable, error) {
	f, err := os.Open(arpFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Scan()

	var table ArpTable
	for sc.Scan() {
		line := sc.Text()
		fields := strings.Fields(line)

		table = append(table, ARPEntry{
			IPAddr: fields[ipAddr],
			HWAddr: fields[hwAddr],
		})
	}

	return table, nil
}
