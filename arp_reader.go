package arpreader

type ARPEntry struct {
	IPAddr       string
	HWAddr       string
	Manufacturer string
}

type ArpTable []ARPEntry

type TableConfig struct {
	IgnoreManufacturer bool
}
