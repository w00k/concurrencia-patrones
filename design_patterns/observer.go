package main

import "fmt"

//patron que permite que diferentes objetos se suscriban a otro objeto cuando ciertos eventos se ejecutan
//es de tipo pasivo, porque espera un evento para enviar la notificación

type Topic interface {
	register(observer Observer)
	broadcast()
}

type Observer interface {
	getId() string
	updateValue(string)
}

// Item -> No disponible
// Item -> Avise -> Disponible
type Item struct {
	observers []Observer
	name      string
	available bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

func (i *Item) broadcast() {
	//iterar los observadores
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

//cambia el estado de available a true y el broadcast notifica a los observadores
func (i *Item) UpdateAvailable() {
	fmt.Printf("Item %s is available \n", i.name)
	i.available = true
	i.broadcast()
}

type EmailClient struct {
	id string
}

func (eC *EmailClient) getId() string {
	return eC.id
}

func (eC *EmailClient) updateValue(value string) {
	fmt.Printf("Sending Email - %s available for client %s\n", value, eC.id)
}

func main() {
	nvidiaItem := NewItem("RTX 3080")
	firstObserver := &EmailClient{
		id: "12ab",
	}
	secondObserver := &EmailClient{
		id: "34bc",
	}

	nvidiaItem.register(firstObserver)
	nvidiaItem.register(secondObserver)
	nvidiaItem.UpdateAvailable()
}
