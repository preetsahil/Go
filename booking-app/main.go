package main

import (
	"fmt"
	"sync"
	"time"
	// "strconv"
	// "booking-app/validate"
	// --if we add validateuserinput.go into folder validate and package name(validate) is different from main then we do
	//this to use function of different package and capatilize the first letter of function and call like this validate.ValidateUserInput()
)

var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
// map[string]string-->in place of UserData only in next line
var bookings =make([]UserData,0)   
type UserData struct{
	firstName string
	lastName string
	email string
	numberoftickets uint
}

var wg=sync.WaitGroup{}

func main() {
	greetUsers()

	
	// for{
	// 	firstName, lastName, email, userTickets := getUserInput()
	// 	isvalidName, isvalidemail, isvalidusertickets := validate.ValidateUserInput(firstName, lastName, email, userTickets,remainingTickets)

	// 	if isvalidName && isvalidemail && isvalidusertickets {
	// 		bookTickets(userTickets, firstName, lastName, email)
	// 		wg.Add(1)
	// 	go	sendTicket(userTickets, firstName, lastName,email)
	// 		firstNames := getfirstNames()
	// 		fmt.Printf("The first name of all the bookings are: %v\n", firstNames)
	// 		if remainingTickets == 0 {
	// 			fmt.Printf("Our conference is fully booked. Come back next year\n")
	// 			break
	// 		}
	// 	} else {
	// 		if !isvalidName {
	// 			fmt.Println("firstname or lastname you entered is too short")
	// 		}
	// 		if !isvalidemail {
	// 			fmt.Println("email address you entered doesn't contain @ sign")
	// 		}
	// 		if !isvalidusertickets {
	// 			fmt.Println("number of tickets you entered is invalid")
	// 		}
	// 	}
		
	// }
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getfirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)  booking["firstName"]
		firstNames = append(firstNames,booking.firstName )
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter your userTickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	//create a map for user
	// var UserData =make(map[string]string)
	var userData =UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberoftickets: userTickets,
	}

	// UserData["firstName"]=firstName
	// UserData["lastName"]=lastName
	// UserData["email"]=email
	// UserData["numberoftickets"]=strconv.FormatUint(uint64(userTickets),10)
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings,userData)
	fmt.Printf("list of bookings is %v\n",bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive the confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v Tickets remaining for %v.\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint,firstName string,lastname string, email string){
     time.Sleep(20*time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v",userTickets,firstName,lastname)
	fmt.Println("#################")
	fmt.Printf("Sending ticket: \n %v \nto email address %v\n",ticket,email)
	fmt.Println("#################")
	wg.Done()
}