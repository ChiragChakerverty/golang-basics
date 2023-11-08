package main

import (
	"fmt"
	"strings"
)

func greetUsers(ConferenceName string , numberOfTickets int , remainingTickets int) {
	fmt.Println("Welcome to", ConferenceName ,"by Chirag")
	fmt.Println("We have total of" , numberOfTickets ,"tickets and" , remainingTickets , "are remaining") 
	fmt.Println("To attend the talk , get your tickets from here")
}

func takingUserInput() (string , string , string , int) {
	var FirstName string
	var LastName string
	var UserTickets int
	var email string 
	
	fmt.Printf("Enter your first name : ")
	fmt.Scan(&FirstName)
	
	fmt.Printf("Enter your last name : ")
	fmt.Scan(&LastName)
	
	fmt.Printf("enter your email address : ")
	fmt.Scan(&email)
	
	fmt.Printf("Enter the number of tickets you want for the event : ")
	fmt.Scan(&UserTickets)

	return FirstName , LastName , email , UserTickets
}

func giveFirstNames(bookings []UserData) []string {
	var FirstNames[] string
		for _,booking := range bookings {
			FirstNames = append(FirstNames , booking.FirstName)
		} 
	return FirstNames
}

func isUserInputValid (FirstName string, LastName string, email string, UserTickets *int, remainingTickets *int) (bool,bool,bool,bool) {
		isNameValid := (len(FirstName)>=3) && (len(LastName)>=2)
		isEmailValid := strings.Contains(email,"@gmail.com")
		isUserTicketValid := (*UserTickets > 0) && (*UserTickets <= *remainingTickets)
		userInputValid := isNameValid && isEmailValid && isUserTicketValid 

		return isNameValid,isEmailValid,isUserTicketValid,userInputValid
}

func PrintBookingDetails(FirstName string, LastName string, email string, UserTickets *int, remainingTickets *int,bookings []UserData) []UserData {

	// UserData := make(map[string]string) // making a map to store user data 
	// UserData["FirstName"] = FirstName
	// UserData["LastName"] = LastName
	// UserData["email"] = email
	// UserData["NumberOfTickets"] = strconv.FormatInt(int64(*UserTickets),10)

	userData := UserData {
		FirstName: FirstName,
		LastName: LastName,
		email:email,
		numberOfTickets:*UserTickets,
	}

	fmt.Println("Thank you" , FirstName , LastName , "for booking" , *UserTickets , "tickets. You will receive a confirmation mail at" , email)
			*remainingTickets = *remainingTickets - *UserTickets
			fmt.Println("Now" , *remainingTickets , "tickets are left")
			bookings = append(bookings, userData)

		return bookings // in order to have the change in bookings
}

func PrintReasonOfInvalidation (isNameValid bool , isEmailValid bool , isUserTicketValid bool) {
	if !isNameValid {
		fmt.Println("The name you entered is too short")
	}

	if !isEmailValid {
		fmt.Println("The email you entered is invalid")
	}

	if !isUserTicketValid {
		fmt.Println("Number of tickets you are asking for is not available")
	}
	fmt.Println("Please try again")
}