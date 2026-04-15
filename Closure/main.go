package main

import "fmt"

func main() {
	userGiftcard := activateGiftcard()
	userGiftcard(10)
	userGiftcard2 := activateGiftcard()
	userGiftcard2(10)

	fmt.Println(userGiftcard(10))
	fmt.Println(userGiftcard2(10))
	//closure

}
func activateGiftcard() func(int) int {

	amount := 100

	debitFunction := func(debitamount int) int {
		amount -= debitamount
		return amount
	}

	return debitFunction
}
