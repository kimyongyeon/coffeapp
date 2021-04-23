package main

import (
	"fmt"
)

type Customer struct {
	name string
	phone int
	addr string
	postCode string
}

type CustomerRepo interface {
	getName() string
	getPhone() int
	getAddr() string
	getPostCode() string
}

func (customer Customer) getName() string {
	return customer.name
}
func (customer Customer) getPhone() int {
	return customer.phone
}
func (customer Customer) getAddr() string {
	return customer.addr
}
func (customer Customer) getPostCode() string {
	return customer.postCode
}

func customerPrint(customerRepo CustomerRepo) {
	fmt.Println(customerRepo.getName())
	fmt.Println(customerRepo.getPhone())
	fmt.Println(customerRepo.getAddr())
	fmt.Println(customerRepo.getPostCode())
}


func main() {
	//fmt.Println("hello go")
	//var customer = Customer{name: "김용연", phone:123123, addr:"서울특별시", postCode: "123123"}
	//customerPrint(customer)
	//product.Main()


}