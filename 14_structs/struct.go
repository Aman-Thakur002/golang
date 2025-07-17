package main

//user defined data-structure
// - use to group mutiple fields

import (
	"fmt"
	"time"
)

type order struct {
	id        int
	amount    float32
	status    string
	createdAt time.Time // nano second precision
}
// contructor like approach
func newOrder(id int, amount float32, status string) order{
	 myOrder := order {
		id : id,
		amount : amount,
		status : status,
	 }

	 return myOrder
}

// manipulation of objects of structure through functions 
// -> use * when you need to update the field 

func (o *order) changeStatus(status string) { // (intial letter of struct and name of struct),  (o order) is receiver type
	o.status = status   // no need to derefernece use * with "o" as struct do it on its own
}

func (o order) getAmount() float32 {
	return o.amount
}

func main() {

	myOrder := order{
		id:     1,
		amount: 100.0,
		status: "pending",
	}

	fmt.Println(myOrder)

	myOrder.changeStatus("Confirmed") // used the fcuntion to manipulate the object myOrder
	myOrder.createdAt = time.Now()

	fmt.Println(myOrder)
	fmt.Println(myOrder.getAmount())

	o1 := newOrder(1,200,"Approved")
	fmt.Println(o1)
	o2 := newOrder(2,500,"Cancelled")
	fmt.Println(o2)

	//one time use struct 
	language := struct {
		name string
		isGood bool
	} {"Golang", true}

	fmt.Println(language)

}
