package initializers

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/config"
	"github.com/sudhanshu-k/NITH-Online-Internship-Document-Signing/tree/main/back-end/utils"
)

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigType("dotenv")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		utils.Logger.Fatal("Unable to read .env.", zap.Error(err))
	}

	err = viper.Unmarshal(config.Config)
	if err != nil {
		utils.Logger.Fatal("Unable to unmarshal.", zap.Error(err))
	}

	utils.Logger.Info("Succesfully loaded config.")
}
