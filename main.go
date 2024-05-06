package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	Yellow = "\033[1;33m"
	White  = "\033[0m"
)

type User struct {
	ID   int
	Name string
}

type Message struct {
	SenderID   int
	ReceiverID int
	Content    string
	Timestamp  time.Time
}

var users = []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
	{ID: 3, Name: "Charlie"},
}

func sendMessage(senderID, receiverID int, content string) {

	fmt.Printf("User %d sent message to User %d: %s\n", senderID, receiverID, content)
}

func broadcastMessage(senderID int, content string) {

	for _, user := range users {
		if user.ID != senderID {
			fmt.Printf("User %d received broadcast from User %d: %s\n", user.ID, senderID, content)
		}
	}
}

func viewMessageLog(userID int) {

	fmt.Printf("Message log for User %d:\n", userID)

}

func main() {

	var chatterGoSymbol = fmt.Sprintf(`%v
	 ___   ____                         
        /  __| / ___|  ___   
       | |    | |  _  / _ \ 
       | |__  | |_| || (_) | 
        \____|\_____| \___/ 
                             
	   
		%v `, Yellow, White)
	fmt.Println(chatterGoSymbol)
	fmt.Println("Welcome to the chatterGo!")
	for {
		fmt.Println("1. Send Message between two users")
		fmt.Println("2. Broadcast Message to all users")
		fmt.Println("3. View Message Log of a User")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			var senderID, receiverID int
			var content string
			fmt.Print("Enter sender ID: ")
			fmt.Scanln(&senderID)
			fmt.Print("Enter receiver ID: ")
			fmt.Scanln(&receiverID)
			fmt.Print("Enter message content: ")
			content, _ = bufio.NewReader(os.Stdin).ReadString('\n')
			content = strings.TrimSpace(content)
			sendMessage(senderID, receiverID, content)
		case "2":
			var senderID int
			var content string
			fmt.Print("Enter sender ID: ")
			fmt.Scanln(&senderID)
			fmt.Print("Enter broadcast message content: ")
			content, _ = bufio.NewReader(os.Stdin).ReadString('\n')
			content = strings.TrimSpace(content)
			broadcastMessage(senderID, content)
		case "3":
			var userID int
			fmt.Print("Enter user ID: ")
			fmt.Scanln(&userID)
			viewMessageLog(userID)
		case "4":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a valid option.")
		}
	}
}
