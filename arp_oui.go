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
)

func isValidFile(path string) bool {
	if stat, err := os.Stat(path); err != nil {
		return false
	} else {
		return stat.Size() > 0
	}
}

func ensureFile() error {
	if !isValidFile(ouiUri) {
		r, err := http.Get(ouiFilePath)
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
		fields := strings.Fields(line)

		if strings.HasPrefix(hwAddr, fields[manufacturerPrefixIdx]) {
			return fields[manufacturerPrefixIdx], nil
		}
	}

	return "", nil
}
