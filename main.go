package main

import (
	"fmt"
	"strings"
)

// Package level variables
// Initialization
var courseName = "GO Langauage basics"
var author = "Prakasa Rao Lankalapalli"

const totNumberOfSeats uint = 100

// var bookedSeats = 0
var remaingSeats uint = totNumberOfSeats
var courseFee uint = 1000
var courseDiscount = 10.4

// Slice dynamic size
var userNameList []string

// Array fixed size
var userEmailList [totNumberOfSeats]string
var userCount uint

func main() {

	printWelcomeMessage()
	printCourseDetails()

	userCount = 0
	for remaingSeats > 0 {

		var userFirstName string
		var userLastName string
		var userEmailId string
		var numberOfSeats uint
		var confirmation bool

		var userCity string

		//Taking input from users

		userFirstName, userLastName, userEmailId, userCity, numberOfSeats, confirmation = getUserInput(userFirstName, userLastName, userEmailId, userCity, numberOfSeats, confirmation, remaingSeats, courseFee)

		isUserInputValid := validateUserInput(userFirstName, userLastName, userEmailId, userCity, numberOfSeats, remaingSeats)

		if !isUserInputValid {
			continue
		}

		remaingSeats = remaingSeats - numberOfSeats
		fmt.Println("Number of seats left ", remaingSeats)

		userEmailList[userCount] = userEmailId
		userNameList = append(userNameList, userFirstName+" "+userLastName)
		userCount = userCount + 1

		fmt.Printf("Discount got is %v for City %v\n", getDiscount(userCity), userCity)

		fmt.Printf("User Name list %v\n", userNameList)
		fmt.Printf("User Name Type %T\n", userNameList)
		fmt.Printf("User Name size %v\n", len(userNameList))

		fmt.Printf("User Email list %v\n", userEmailList)
		fmt.Printf("User Email Type %T\n", userEmailList)
		fmt.Printf("User Email size %v\n", len(userEmailList))

		fmt.Printf("Printing Users List **************************\n")
		//Ususal syntax for loop is for index, list to iterate. But
		//We don't need first variable so marked as _
		for _, booking := range userNameList {
			var names = strings.Fields(booking)
			fmt.Printf("First name is %v and Last name is %v\n", names[0], names[1])
		}
	}
}

func getDiscount(userCity string) float32 {
	var discount float32

	switch userCity {
	case "B":
		discount = 10.5
	case "C":
		discount = 15.5
	case "H":
		discount = 1.5
	default:
		discount = 0
	}
	return discount
}

func printWelcomeMessage() {
	fmt.Println("Welcome to ", courseName, "Mr. Prakash. Author````", author)
	fmt.Printf("This is a short course with basics of GO. This course has %v seats available\n", totNumberOfSeats)
}

func printCourseDetails() {
	fmt.Printf("Course Fee is %v and discount can be availed upto %v\n", courseFee, courseDiscount)
}

func getUserInput(userFirstName string, userLastName string, userEmailId string, userCity string, numberOfSeats uint, confirmation bool, remaingSeats uint, courseFee uint) (string, string, string, string, uint, bool) {
	fmt.Print("Please enter user first name:\n")
	fmt.Scan(&userFirstName)
	fmt.Print("Please enter user last name:\n")
	fmt.Scan(&userLastName)
	fmt.Print("Please enter Email ID:\n")
	fmt.Scan(&userEmailId)
	fmt.Print("Please enter City [C/B/H]\n")
	fmt.Scan(&userCity)
	fmt.Printf("Please enter number of seats to book. Available at present are %v:\n", remaingSeats)
	fmt.Scan(&numberOfSeats)
	fmt.Printf("\nCost of the course for %v is %v", numberOfSeats, (numberOfSeats * uint(courseFee)))
	fmt.Println("Please confirm (Y/N)")
	fmt.Scan(&confirmation)

	return userFirstName, userLastName, userEmailId, userCity, numberOfSeats, confirmation

}
