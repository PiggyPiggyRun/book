package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string
	Age  uint
}

func main() {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/users?charset=utf8&parseTime=True")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create
	db.Create(&User{Name: "James Bond", Age: 40})

	// Read
	var user User
	db.First(&user, 1) // find user with id 1
	fmt.Println(user)

	db.First(&user, "Name = ?", "James Bond") // find James Bond
	fmt.Println(user)

	// Update - update Bond's age
	db.Model(&user).Update("Age", 41)
	fmt.Println(user)

	// Delete - delete user
	db.Delete(&user)

	createTwoUsers(db)
}

func createTwoUsers(db *gorm.DB) {
	userA := User{Name: "UserA", Age: 20}
	userB := User{Name: "UserB", Age: 20}

	tx := db.Begin()
	if err := tx.Create(&userA).Error; err != nil {
		tx.Rollback()
	}
	if err := tx.Create(&userB).Error; err != nil {
		tx.Rollback()
	}

	//commit!
	tx.Commit()
}
