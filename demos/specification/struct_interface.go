package main

import "fmt"

type (
	Event interface {
		notify()
	}
	ptrEvent struct {
		name string
	}

	entityEvent struct {
		name string
	}
)

func (e *ptrEvent) notify() {
	fmt.Printf("Implements by type: %T value: %v \n", e, e)
}

func (e entityEvent) notify() {
	fmt.Printf("Implements by type: %T value: %v \n", e, e)
}

func newPtrEvent() Event {
	return &ptrEvent{"aaaa"}
}

func newEntityEvent() Event {
	return entityEvent{"bbbb"}
}

func main() {
	e := newPtrEvent()
	e.notify()
	//ee := e.(ptrEvent) // ptrEvent does not implement Event (notify method has pointer receiver)
	ee := e.(*ptrEvent)
	fmt.Printf("Type: %T \n", ee)

	e2 := newEntityEvent()
	e2.notify()
	//ee2 := e2.(*entityEvent) // interface conversion: main.Event is main.entityEvent, not *main.entityEvent
	ee2 := e2.(entityEvent) // interface conversion: main.Event is main.entityEvent, not *main.entityEvent
	fmt.Printf("Type: %T \n", ee2)
}
