package main

import (
	"Assignment1/Hotel"
	"Assignment1/Employee"
	"Assignment1/Gym"
	"Assignment1/Wallet"
)

func main() {
	Hotel.RunHotelMenu()
	Employee.RunEmployeeDemo()
	Gym.RunGymDemo()
	Wallet.RunWalletMenu()
}