package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

//Deposit() -> escribir (condición de carrera)
func Deposit(amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done() // decremento
	lock.Lock()     // bloquea el programa
	balance += amount
	lock.Unlock() // desbloquea el programa
}

//Balance() -> leer balance
func Balance(lock *sync.RWMutex) int {
	lock.RLock() // bloquea las escrituras, pero no las lecturas
	b := balance
	lock.RUnlock() // desbloquea el programa
	return b
}

func main() {
	var wg sync.WaitGroup // "contador"
	var lock sync.RWMutex // para eliminar la condición de carrera
	for i := 1; i <= 5; i++ {
		wg.Add(1) // agrego 1
		go Deposit(i*100, &wg, &lock)
	}
	wg.Wait() // bloquedo el programa hasta que finalice
	fmt.Println(Balance(&lock))
}
