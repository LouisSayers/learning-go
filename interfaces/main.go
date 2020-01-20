package main

import "fmt"

type Person struct {
	name string
}

type dbService interface {
	Get(id int) *Person
	Put(p *Person, id int)
}

type mongod struct {
	store map[int]*Person
}

func NewMongod() *mongod {
	return &mongod{
		store: make(map[int]*Person),
	}
}

func (m mongod) Get(id int) *Person {
	return m.store[id]
}

func (m mongod) Put(p *Person, id int) {
	m.store[id] = p
}

func main() {
	p1 := Person{name: "Joe Banks"}
	p2 := Person{name: "Angie Jacks"}

	var mdb dbService = NewMongod()
	mdb.Put(&p1, 1)
	mdb.Put(&p2, 2)

	fmt.Printf("Person1: %v, Person2: %v\n", *mdb.Get(1), *mdb.Get(2))
}
