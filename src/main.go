package main

import (
	"errors"
	"rsoi/lab1/controllers"
	"rsoi/lab1/objects"
	"rsoi/lab1/utils"

	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func initDBConnection(cnf utils.DBConfiguration) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cnf.Host, cnf.User, cnf.Password, cnf.Name, cnf.Port)
	db, e := gorm.Open(cnf.Type, dsn)

	if e != nil {
		utils.Logger.Print("DB Connection failed")
		utils.Logger.Fatal(e)
	} else {
		utils.Logger.Print("DB Connection Established")
	}

	db.SingularTable(true)
	db.AutoMigrate(&objects.Person{})

	return db
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) < 2 {
		panic(errors.New("no config path"))
	}

	utils.InitConfig(os.Args[1])
	utils.InitLogger()
	defer utils.CloseLogger()

	db := initDBConnection(utils.Config.DB)
	defer db.Close()

	r := controllers.InitRouter(db)
	utils.Logger.Print("Server started")
	fmt.Printf("Server is running on http://localhost:%d\n", utils.Config.Port)
	code := controllers.RunRouter(r, utils.Config.Port)

	utils.Logger.Printf("Server ended with code %s", code)
}
