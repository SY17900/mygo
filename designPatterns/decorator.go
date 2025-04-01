package designpatterns

import "fmt"

type Component interface {
	execute()
}

type ConcreteComponent struct{}

func (*ConcreteComponent) execute() {
	fmt.Println("Lowest Execution")
}

type Decorator0 struct {
	wrapee ConcreteComponent
}

func (d *Decorator0) execute() {
	d.wrapee.execute()
	fmt.Println("Decorator0 Execution")
}
