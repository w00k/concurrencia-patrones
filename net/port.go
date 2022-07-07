package main

import (
	"fmt"
	"net"
)

func main() {
	//revisar si los puertos estan libres y hacer una conexión tcp,
	//si funciona imprimimos que el puerto esta abierto
	for i := 0; i < 100; i++ {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "scanme.nmap.org", i))
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Printf("Port %d is open\n", i)
	}
}
