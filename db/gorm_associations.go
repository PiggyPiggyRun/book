package main

import (
	_ "fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type HotelChain struct {
	gorm.Model
	Name string
}

type Hotel struct {
	gorm.Model
	Name    string
	NoRooms uint
	Chain   HotelChain `gorm:"foreignkey:ChainId"` // use ChainId as foreign key
	ChainId uint
}

func main() {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/users?charset=utf8&parseTime=True")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&HotelChain{})
	db.AutoMigrate(&Hotel{})

	db.Model(&Hotel{}).AddForeignKey("chain_id", "hotel_chains(id)", "CASCADE", "CASCADE")

	// Create some entities and save
	taj := HotelChain{Name: "Taj"}
	db.Save(&taj)
	vivanta := Hotel{Name: "Vivanta by the sea", NoRooms: 400, Chain: taj}
	db.Save(&vivanta)

}
