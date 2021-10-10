package main

import (
	arpreader "github.com/hpedrorodrigues/arp-reader"
	"log"
)

func main() {
	table, err := arpreader.GetTable(&arpreader.TableConfig{IgnoreManufacturer: false})
	if err != nil {
		log.Fatalln(err)
	}

	for _, entry := range table {
		log.Printf("Manufacturer: %s, IpAddr: %s, HWAddr: %s\n", entry.Manufacturer, entry.IPAddr, entry.HWAddr)
	}
}
