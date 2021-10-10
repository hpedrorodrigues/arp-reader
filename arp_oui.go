package arpreader

import (
	"bufio"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	ouiUri      = "http://standards-oui.ieee.org/oui/oui.txt"
	ouiFilePath = "/tmp/oui.txt"
	ouiFileMode = 0644

	manufacturerPrefixIdx = 0
	manufacturerNameIdx   = 2
)

func isValidFile(path string) bool {
	if stat, err := os.Stat(path); err != nil {
		return false
	} else {
		return stat.Size() > 0
	}
}

func ensureFile() error {
	if !isValidFile(ouiFilePath) {
		r, err := http.Get(ouiUri)
		if err != nil {
			return err
		}
		defer r.Body.Close()

		if r.StatusCode != 200 {
			return errors.New("cannot retrieve manufacturer list")
		}

		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}

		if err = ioutil.WriteFile(ouiFilePath, content, ouiFileMode); err != nil {
			return err
		}
	}

	return nil
}

func FindManufacturer(hwAddr string) (string, error) {
	if err := ensureFile(); err != nil {
		return "", err
	}

	f, err := os.Open(ouiFilePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Scan()

	for sc.Scan() {
		line := sc.Text()
		if len(line) == 0 {
			continue
		}

		if !strings.Contains(line, "(hex)") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}

		prefix := strings.ToLower(strings.ReplaceAll(fields[manufacturerPrefixIdx], "-", ":"))
		shortHwAddr := strings.Join(strings.Split(hwAddr, ":")[:3], ":")
		if strings.Contains(prefix, shortHwAddr) {
			return strings.Join(fields[manufacturerNameIdx:], " "), nil
		}
	}

	return "", nil
}
