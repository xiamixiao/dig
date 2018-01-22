package main

import (
	"os"
	"io/ioutil"
	"dig/dnsutil"
	"fmt"
	"regexp"
)

func main() {
	resolv, err := ioutil.ReadFile("/etc/resolv.conf")
	if err != nil {
		fmt.Println(err)
	}
	reg := regexp.MustCompile(`(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)`)
	arr := reg.FindAllString(string(resolv), -1)

	domain := "www.baidu.com"
	if len(os.Args) == 2 {
		domain = os.Args[1]
	}
	dig := new(dnsutil.Dig)
	dig.SetDNS(arr[0])
	a, _ := dig.CNAME(domain)
	fmt.Println(a)
	if a != nil {
		for _, v := range a {
			fmt.Println(v.Target)
		}
	}
}