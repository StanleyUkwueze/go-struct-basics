package main

import "fmt"

type contactInfo struct{
	email string
	zipCode int
}
type Person struct {
	firstName string
	lastName  string
	contact contactInfo

}

func main() {
	jim := Person{
		firstName:  "Jim",
		lastName: "Party",
		contact: contactInfo{
			email:"jim@mail.com",
			zipCode:98765,
		},
	}

	jimPointer := &jim
	jimPointer.updateName("Jimmmmy")
	
	//fmt.Println(jimPointer)
	fmt.Println(&jimPointer)
	//jim.print()
}

func(pointerToPerson *Person) updateName(newFirstName string){
(*pointerToPerson).firstName = newFirstName
fmt.Println(&pointerToPerson)
}

func(p Person) print(){
	
	fmt.Printf("%+v", p)
}