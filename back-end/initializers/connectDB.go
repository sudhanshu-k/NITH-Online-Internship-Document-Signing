package initializers

import (
	"context"

	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/database"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/middleware"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(config *Config) {
	var err error // define error here to prevent overshadowing the global DB

	database.DB, err = pgxpool.New(context.Background(), config.DATABASE_URL)

	middleware.LogIfError(err, "Failed to connect to DB")
	println("Connection succesful to DB")
}
