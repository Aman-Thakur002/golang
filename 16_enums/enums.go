package main

import "fmt"

// enums -> enumarated variables

// type orderStatus int
// const (
//  Received orderStatus = iota // iota is a predeclared identifier representing the untyped integer ordinal number of the current const specification in a (usually parenthesized) const declaration. It is zero-indexed.
//  Confirmed
//  Prepared
//  Delivered
// )

type orderStatus string
const (
	Received orderStatus = "Received"
	Confirmed orderStatus = "Confirmed"
	Prepared orderStatus = "Prepared"
	Delivered orderStatus = "Delivered"
)


func changeOrderStatus (status orderStatus){
	fmt.Println("Changed Order Status : ", status)
}

func main(){
    changeOrderStatus(Confirmed)
}