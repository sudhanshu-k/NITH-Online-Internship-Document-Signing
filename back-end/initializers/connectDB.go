package initializers

import (
	"context"
	// "fmt"

	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/config"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/database"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/utils"

	// "github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/utils"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func ConnectDB() {
	var err error

	// connect to db
	database.DB, err = pgxpool.New(context.Background(), config.Config.DATABASE_URL)
	if err != nil {
		utils.Logger.Fatal("Unable to connect to database.", zap.Error(err))
	}

	_, err = database.DB.Exec(context.Background(), "select 1")
	if err != nil {
		utils.Logger.Fatal("Error occured while executing test query.", zap.Error(err))
	}

	utils.Logger.Info("Connected to database")
}
