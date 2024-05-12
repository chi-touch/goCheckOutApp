package main

import (
	"fmt"
	"strings"
	"time"
)

type Item struct {
	Name     string
	Price    float64
	Quantity int
}

type Cart struct {
	Items []Item
}

func main() {
	var cart Cart
	var customerName string

	fmt.Println("WELCOME TO CHICHI CHECKOUT")
	fmt.Print("What is your name:")
	fmt.Scanln(&customerName)

	addItemsToCart(&cart)

	var cashierName string
	var discount float64

	fmt.Println("cashier enter your name")
	fmt.Scanln(&cashierName)

	fmt.Print("what is the given discount")
	fmt.Scanln(&discount)

	printInvoice(customerName, cashierName, cart, discount)

}

func addItemsToCart(cart *Cart) {
	for {
		var itemName string
		var price float64
		var quantity int

		fmt.Print("Enter what you bought: ")
		fmt.Scanln(&itemName)

		fmt.Println("Enter price of goods: ")
		fmt.Scanln(&price)

		fmt.Println("Enter the quantity of goods you want to buy: ")
		fmt.Scanln(&quantity)

		item := Item{Name: itemName, Price: price, Quantity: quantity}

		cart.Items = append(cart.Items, item)

		var addMoreItems string
		fmt.Print("Do you want to add more items? Enter 'Yes' or 'No': ")
		fmt.Scanln(&addMoreItems)

		if strings.ToUpper(addMoreItems) != "Yes" {
			break

		}
	}

}

func calculateSubtotal(cart Cart) float64 {
	var subtotal float64

	for _, item := range cart.Items {
		subtotal += item.Price * float64(item.Quantity)

	}
	return subtotal

}

func calculateDiscount(subtotal float64, discount float64) float64 {
	return (discount / 100) * subtotal
}

func calculateTotal(subtotal float64, discount float64) float64 {
	totalAfterDiscount := subtotal - calculateDiscount(subtotal, discount)
	vat := 0.175 * subtotal
	return totalAfterDiscount + vat
}

func printInvoice(customerName string, cashierName string, cart Cart, discount float64) {
	fmt.Println("==========================================================================")
	fmt.Println("								CHICHI STORE								")
	fmt.Println("								MAIN BRANCH									")
	fmt.Println("		LOCATION: 321, HERBERT MACAULAY WAY, SABO, YABA, LAGOS.				")
	fmt.Println("								PHONE: 081-0372-2570						")
	fmt.Println("==========================================================================")
	fmt.Printf("Date: %s\n", time.Now().Format("24-05-12"))
	fmt.Printf("Cashier: %s\n", cashierName)
	fmt.Printf("Customer: %s\n", customerName)
	fmt.Println("==========================================================================")
	fmt.Println(" ITEM \t\tPRICE \t\tTOTAL(NGN)")
	fmt.Println("==========================================================================")

	subtotal := calculateSubtotal(cart)
	discountAmount := calculateDiscount(subtotal, discount)
	total := calculateTotal(subtotal, discount)

	for _, item := range cart.Items {
		fmt.Printf("%s\t\t%.2f\t\t%d\t\t%.2f\n", item.Name, item.Price, item.Quantity, item.Price*float64(item.Quantity))
	}

	fmt.Println("--------------------------------------------------------------------------")
	fmt.Printf("\t\tSub Total:		%.2f\n", subtotal)
	fmt.Printf("\t\tDiscount:		%.2f\n", discountAmount)
	fmt.Printf("\t\tVAT @ 17.5%:		%.2f\n", 0.175*subtotal)
	fmt.Println("==========================================================================")
	fmt.Printf("\t\tTotal Bill:	 	%.2f\n", total)
	fmt.Println("==========================================================================")
	fmt.Println("		THIS IS NOT A RECEIPT- KINDLY PAY									")
	fmt.Println("==========================================================================")

}
