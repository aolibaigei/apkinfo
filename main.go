package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type PackageInfo struct {
	Package  string
	Version  string
	Homepage string
}

var infopath = "/var/lib/dpkg/status"

func (v *PackageInfo) getInfo() (string, string, string) {
	return v.Package, v.Version, v.Homepage
}

func parsinfo(info string) string {
	splitno := strings.Index(info, ":")

	v := info[splitno:]

	return strings.TrimSpace(strings.TrimPrefix(v, ":"))
}

func main() {
	fi, err := os.Open(infopath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	pkginfo := &PackageInfo{}
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if len(a) != 0 {
			//
			if strings.Contains(string(a), "Package:") {
				pkginfo.Package = parsinfo(string(a))
			}
			///
			if strings.Contains(string(a), "Version:") {
				pkginfo.Version = parsinfo(string(a))
			}
			if strings.Contains(string(a), "Homepage:") {
				pkginfo.Homepage = parsinfo(string(a))
			}

		}
		if len(a) == 0 {
			fmt.Println(pkginfo)
			// info := fmt.Sprintf("%s %s %s %s %s \n", bug.Package, bug.Version, bug.Bugs, bug.Homepage, bug.Origin)
			// _, err1 := io.WriteString(f, info)
			// if err1 != nil {
			// 	panic(err1)
			// }
		}
		if c == io.EOF {
			break
		}

	}
}
