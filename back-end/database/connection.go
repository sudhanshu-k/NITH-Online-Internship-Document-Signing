package database

import (
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/middleware"
	"os"

	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error // define error here to prevent overshadowing the global DB

	env := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(env), &gorm.Config{})

	middleware.LogIfError(err, "Failed to connect to DB")
	println("Connection succesful to DB")

	err = DB.AutoMigrate(&model.Student{}, &model.Faculty{}, &model.Profile{}, &model.Session{})
	middleware.FatalError(err)
	println("Migrated DB")
}