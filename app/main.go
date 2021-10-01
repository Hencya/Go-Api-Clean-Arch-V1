package main

import (
	_userUsecase "book_crud_ca/bussinesses/users"
	_userController "book_crud_ca/controllers/users"
	_driverFactory "book_crud_ca/drivers"
	_usersRepo "book_crud_ca/drivers/databases/users"
	_dbDriver "book_crud_ca/drivers/mysql"

	_routes "book_crud_ca/app/routes"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"time"
)

func init() {
	viper.SetConfigFile(`app/config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_usersRepo.Users{},
	)
}

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configDB.InitDB()
	dbMigrate(db)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()

	userRepo := _driverFactory.NewUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, timeoutContext)
	userCtrl := _userController.NewUserController(userUsecase)

	routesInit := _routes.ControllerList{
		UserController: *userCtrl,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
