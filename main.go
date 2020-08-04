package main

import (
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	// get working dic
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Current Dir is:", dir)
	ips, err := getIps()
	if err != nil {
		log.Fatal(err)
	}
	for _,v := range ips{
		log.Printf("start file server at http://%s:8080",v)
	}

	fileserver := http.FileServer(http.Dir(dir))
	err = http.ListenAndServe(":8080", fileserver)
	if err != nil {
		log.Fatal(err)
	}
}

func getIps() ([]string, error) {
	var result []string
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			default:
				continue
			}
			if ip.To4() != nil{
				result = append(result, ip.String())
			}

		}
	}
	return result, nil
}
