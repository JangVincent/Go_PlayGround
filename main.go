package main

import (
	"context"
	"fmt"

	"github.com/Phantola/Go_PlayGround/database"
	"github.com/Phantola/Go_PlayGround/greetings"
	"github.com/Phantola/Go_PlayGround/models"
)



func main() {
	fmt.Println("Hello, world.")

	ctx := context.Background()
	db, err := database.GetDB()

	if err != nil {
		fmt.Println("Error getting database connection")
		return
	}

	defer database.CloseDB(db)

	user, err := models.FindUser(ctx, db, "1")

	if err != nil {
		fmt.Println("Error getting user")

		grettings, err := greetings.Hellos([]string{"Gladys", "Samantha", "Darrin"})
		if err != nil {
			fmt.Println("Error getting greetings")
		}

		fmt.Println(grettings)
		return
	}

	fmt.Println(user)
}
