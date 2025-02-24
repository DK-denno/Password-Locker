package main

import (
	"fmt"
	"time"

	Users "github.com/DK-denno/passwordLocker/passwordLocker/users"
	Accountcredentials "github.com/DK-denno/passwordLocker/passwordLocker/accountcredentials"
)


func displayLoggedOutMenu() int {
	//defer recover()
	fmt.Println(" \n Here is your Menu :- ")
	fmt.Println("1). Create New Account")
	fmt.Println("2). Login")
	fmt.Println("3). About Us")
	fmt.Print(" Pick an Option >> ")

	var option int
	_, err := fmt.Scanf("%d", &option)

	if err != nil {
		panic(err)
	}

	fmt.Println(" Option Picked is :- ", fmt.Sprintf("%d", option))
	return option
}

func displayCreateUserAccount() {
	//defer recover()
	fmt.Println("Certainly, Kindly provide me with the following details and you will be good to go")
	var username, email, password string
	fmt.Print(" Username >> ")
	_, err := fmt.Scanf("%s", &username)

	fmt.Print(" Email >> ")
	_, err2 := fmt.Scanf("%s", &email)

	fmt.Print(" Password >> ")
	_, err3 := fmt.Scanf("%s", &password)

	if err != nil {
		panic(err)
	}

	if err2 != nil {
		panic(err2)
	}

	if err3 != nil {
		panic(err3)
	}

	_, err4 := Users.AddUserAccount(username, email, password)
	if err != nil {
		panic(err4)
	}
}

func displayLoginUserAccount(user Users.Users) {
	// fmt.Println(Users.UsersDb)
	fmt.Println("Certainly, Kindly provide me with the following details and you will be good to go")
	var username, password string
	fmt.Print(" Username >> ")
	_, err := fmt.Scanf("%s", &username)

	fmt.Print(" Password >> ")
	_, err3 := fmt.Scanf("%s", &password)

	if err != nil {
		panic(err)
	}

	if err3 != nil {
		panic(err3)
	}

	xusr, _err := Users.GetUser(username, password)
	// fmt.Println(xusr)
	if err != nil {
		displayLoginUserAccount(xusr)
		panic(_err)
	} else {
		fmt.Println(xusr)

		if xusr.Username != "" {
			displayLoggedInMainMenu(xusr , displayMainMenu)
		} else {
			fmt.Print(" Password >> ")
			displayLoggedOutMenu()
		}

	}
}

func displayLoggedInMainMenu(user Users.Users, operation func(user Users.Users)) {
	//defer recover()
	fmt.Println("[*], Login Successfull")
	fmt.Println(user)

	operation(user)
}

func displayAboutUs() int {
	//defer recover()
	fmt.Println("Hey there, fellow devs! 👋" +
		"\n [-] We know the struggle—too many passwords, not enough brain space." +
		"\n\n [-] That’s why we built Password Locker, a secure and hassle-free way" +
		"to store and manage all your credentials.\n\n [-] No more “forgot password" +
		"loops or sticky notes with login info (we’ve all been there)." +
		"\n\n [-]  With strong encryption, easy access, and cross-device sync," +
		"we make sure your passwords stay safe while you focus on coding" +
		"the next big thing. \n\n [-] Whether you're juggling API keys, SSH credentials," +
		"or just trying to keep your logins organized, we’ve got your back." +
		" \n\n [-]  Less password stress, more dev time. [-] \n\n That’s the goal. 🚀" +
		"[-]\n \n\n Let’s keep it secure and simple! 🔐💻")

	return 1
}

func displayMainMenu(user Users.Users) {
	// //defer recover()
	fmt.Println("\n Main Menu :- ")

	fmt.Println("1). View Profile")
	fmt.Println("2). Edit Profile Details")
	fmt.Println("3). Add New Account Credentials")
	fmt.Println("4). View Account Credentials")
	fmt.Println("5). Search For Account")
	fmt.Println("6). Delete Account")
	fmt.Println("0). Log Out")

	fmt.Print(" Pick an Option >> ")

	var option int
	_, err := fmt.Scanf("%d", &option)

	if err != nil {
		panic(err)
	}

	fmt.Println(" Option Picked is :- ", fmt.Sprintf("%d", option))

	switch option {
		case 1:
			displayViewProfile(user, displayMainMenu)
		case 2:
			displayEditProfile(user, displayMainMenu)
		case 3:
			displayAddNewAccountCredentials(user, displayMainMenu, true)
		case 4:
			displayAllCredentials(user, displayMainMenu)
		case 0:
			fmt.Print("\n\n\n [*[ Goodbye! ")
			displayLoggedOutMenu()
		default:
			displayMainMenu(user)
	}
}

// Function to display user profile
func displayAddNewAccountCredentials(user Users.Users,
	backButton func(user Users.Users), shouldDisplayText bool) {

		if shouldDisplayText {
			fmt.Println("[*] Hi " + user.Username +
				", Lets add a new Account :-  ")
			time.Sleep(1 * time.Second)
			fmt.Println("[*] Give me a minute ")
			time.Sleep(1 * time.Second)
			fmt.Println("[*] Almooooooooost :- ")
			time.Sleep(1 * time.Second)
			fmt.Println("[*] Okay Done! Sorry I kept you waiting you waiting, " +
				"my back hurts sometimes")
			fmt.Println("[*] ========================= [*]")
		}

		var accountName, accountPassword, accountUserName string
		fmt.Print("[*] Please Input PlatForm Name (facebook, instagram, 2wirra etc) >> ")
		_, err := fmt.Scanln(&accountName)

		fmt.Print("[*] Please Input Account User Name >> ")
		_, err2 := fmt.Scanln(&accountUserName)

		fmt.Print("[*] Please Input Account Password >> ")
		_, err3 := fmt.Scanln(&accountPassword)

		if err != nil {
			panic(err)
		}

		if err2 != nil {
			panic(err2)
		}

		if err3 != nil {
			panic(err3)
		}

		_, err4 := Accountcredentials.AddUserCredentials(
			accountName, accountUserName, accountPassword)

		if err4 != nil {
			panic(err4)
		}
		fmt.Println(Accountcredentials.AccountcredentialsDb)
		fmt.Println("New Credentials for " +
			accountName + " has been added. Would You like to add another one, " +
			" View All Credsentials or Go home")
		fmt.Println("1). Add Another One")
		fmt.Println("2). View All")
		fmt.Println("1). Go Home")
		var opt int
		fmt.Print(" Password >> ")
		_, err5 := fmt.Scanf("%d", &opt)
		if err5 != nil {
			panic(err4)
		}


		switch opt {
			case 1:
				displayAddNewAccountCredentials(user, backButton, false)
			case 2:
				displayAllCredentials(user, displayMainMenu)
		}

	displayAddNewAccountCredentials(user, backButton, shouldDisplayText)

}

// Function to display user profile
func displayAllCredentials(
	user Users.Users, backButton func(user Users.Users)) {

	Accountcredentials.ViewCredentials()

	backButton(user) // ✅ Now it runs only when user is ready
}



// Function to display user profile
func displayViewProfile(user Users.Users, backButton func(user Users.Users)) {
	fmt.Println("\n===== View Profile =====")
	fmt.Printf("Username: %s\n", user.Username)
	fmt.Printf("Email: %s\n", user.Email)

	// Ask user if they want to go back
	fmt.Println("\nPress Enter to go back...")
	fmt.Scanln() // Wait for user input before navigating back

	backButton(user) // ✅ Now it runs only when user is ready
}

// Function to edit user profile
func displayEditProfile(user Users.Users, backButton func(user Users.Users)) {
	fmt.Println("\n Edit Profile :- ")
	fmt.Println("Which of the following fields would you like to edit?")

	fmt.Println("1). Username")
	fmt.Println("2). Email")
	fmt.Println("3). Password")
	fmt.Println("0). Back") // ✅ Changed "00" to "0"

	fmt.Println("\n Username: " + user.Username)
	fmt.Println("Email: " + user.Email)

	var option int
	fmt.Print("\nInput option: ")
	_, err := fmt.Scanln(&option) // ✅ Using Scanln() to avoid errors

	if err != nil {
		fmt.Println("Invalid input! Please enter a number.")
		displayEditProfile(user, backButton) // Restart the menu
		return
	}

	fmt.Println("Option Picked:", option)

	switch option {
	case 1:
		editField("Username", user, func(updatedUser Users.Users) {
		    displayViewProfile(updatedUser, displayMainMenu)
		})
	case 2:
		editField("Email", user, func(updatedUser Users.Users) {
		    displayViewProfile(updatedUser, displayMainMenu)
		})
	case 3:
		editField("Password", user, func(updatedUser Users.Users) {
		    displayViewProfile(updatedUser, displayMainMenu)
		})
	case 0:
		backButton(user) // ✅ Corrected back option
	default:
		fmt.Println("Invalid option! Try again.")
		displayEditProfile(user, backButton) // Restart the menu
	}
}

func editField(field string, user Users.Users, nextButton func(user Users.Users)) {
	var newValue, oldValue string

	fmt.Print("Input Old Value >> ")
	_, err := fmt.Scanln(&oldValue) // ✅ Using Scanln() instead of Scanf()

	fmt.Print("Input New Value >> ")
	_, err2 := fmt.Scanln(&newValue) // ✅ Using Scanln() instead of Scanf()

	if err != nil && err2 != nil {
		fmt.Println("Invalid input! Please try again.")
		editField(field, user, nextButton)
	}

	if user.Username != oldValue || user.Email != oldValue || user.Password != oldValue {
		fmt.Println("Old Value is Wrong")
		displayEditProfile(user, nextButton)
	}

	// Update user details
	updatedUser, err := user.EditUserDetails(user.Username, field, newValue, oldValue)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// // ✅ Save updated user back into UsersDb
	// Users.UsersDb[user.Username] = updatedUser
	fmt.Println("Updated User:", updatedUser)

	// ✅ Call the next function to display updated profile
	nextButton(updatedUser)
}

func main()  {
	fmt.Println("Hellow and Welcome to Our password Locker done in Golang :-")
	fmt.Println(" A Pretty exciting peice of tech ")

	for {

		opt := displayLoggedOutMenu()
		switch opt {
			case 1:
				displayCreateUserAccount()
			case 2:
				displayLoginUserAccount(Users.Users{ })
			case 3:
				displayAboutUs()
		}
	}
}