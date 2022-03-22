package main

import (
	"artificial-catalog/internal/plant"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	fmt.Println("Поехало!")

	logger := initZapLogger()

	initViperConfigger(logger)

	db := initDbConn(logger)

	router := mux.NewRouter()

	plantRepository := plant.NewPlantRepository(db, logger)
	plantService := plant.NewPlantService(plantRepository, logger)
	plant.RegisterPlantHandlers(router, plantService, logger)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Привет Мир") }).Methods("GET")

	configServer(router).ListenAndServe()
}

func configServer(router *mux.Router) *http.Server {
	return &http.Server{
		Handler:        router,
		Addr:           ":" + viper.GetString("DB_PORT"),
		WriteTimeout:   5 * time.Second,
		ReadTimeout:    5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func initViperConfigger(logger *zap.Logger) {
	viper.SetConfigName("app")
	viper.AddConfigPath("config/.")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Print(err)
		logger.Error("failed read in config", zap.Error(err))
		return
	}
}

func initDbConn(logger *zap.Logger) *sql.DB {
	db, err := sql.Open(
		viper.GetString("DB_DRIVER"),
		viper.GetString("DB_USER") + viper.GetString("DB_PASSWORD") + "@" + viper.GetString("DB_SOURSE") + "/" + "")
	if err != nil {
		logger.Error("failed connect to db", zap.Error(err))
		fmt.Println(err)
		os.Exit(1)
	}
	return db
}

func initZapLogger() *zap.Logger {
	logger := zap.NewExample()
	zap.ReplaceGlobals(logger)
	return logger
}
