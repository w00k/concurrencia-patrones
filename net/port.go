package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

//flag para recibirlo como parámetro site
var site = flag.String("site", "scanme.nmap.org", "URl to scan")

func main() {
	//revisar si los puertos estan libres y hacer una conexión tcp,
	//si funciona imprimimos que el puerto esta abierto

	flag.Parse() //permite usar *site como variable
	var wg sync.WaitGroup

	for i := 0; i < 65535; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("Port %d is open\n", port)
		}(i)
	}
	wg.Wait() //fuera del for, sino se va a bloquear uno por uno
}
