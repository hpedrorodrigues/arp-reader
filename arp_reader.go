package arpreader

type ARPEntry struct {
	IPAddr string
	HWAddr string
}

type ArpTable []ARPEntry
