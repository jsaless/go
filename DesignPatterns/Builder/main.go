package main

import "fmt"

type IBuilder interface {
	setName(string) IBuilder
	setProfession(string) IBuilder
	setAge(int) IBuilder
	build() Person
}

type Person struct {
	name       string
	age        int
	profession string
}

type PersonBuilder struct {
	person Person
}

func (b *PersonBuilder) setName(name string) *PersonBuilder {
	b.person.name = name
	return b
}

func (b *PersonBuilder) setProfession(profession string) *PersonBuilder {
	b.person.profession = profession
	return b
}

func (b *PersonBuilder) setAge(age int) *PersonBuilder {
	b.person.age = age
	return b
}

func (b *PersonBuilder) build() *Person {
	return &Person{
		name:       b.person.name,
		age:        b.person.age,
		profession: b.person.profession,
	}
}

func main() {
	personBuilder := &PersonBuilder{}

	personBuilder.setName("João").setAge(20).setProfession("DevOps")

	person := personBuilder.build()

	fmt.Println(person)
}
