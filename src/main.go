package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

const (
	PORT = 8080
)

func getIP() string {
	var ips strings.Builder

	interfaces, err := net.Interfaces()
	if err != nil {
		log.Println("Could not get network interfaces", err.Error())
		return ""
	}

	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Println("Could not get interface's address", err.Error())
			continue
		}

		for _, addr := range addrs {
			ip := net.IPv4(127, 0, 0, 1)
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			ips.WriteString(ip.String())
			ips.WriteRune('\n')
		}
	}

	return ips.String()
}

func getHostname() string {
	str, err := os.Hostname()
	if err != nil {
		log.Println("Could not get the hostname", err.Error())
		return ""
	}

	return str
}

func getVersion() string {
	version := os.Getenv("VERSION")
	return version
}

func main() {
	addr := fmt.Sprintf("127.0.0.1:%v", PORT)
	fmt.Printf("Server running on http://%s", addr)
	http.HandleFunc("/", HandleIndex)
	http.ListenAndServe(addr, nil)
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	type ViewData struct {
		IP       string
		Hostname string
		Version  string
	}
	data := ViewData{
		getIP(),
		getHostname(),
		getVersion(),
	}

	tmpl, err := template.ParseFiles("index.templ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
