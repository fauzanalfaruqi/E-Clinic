package config

import (
	"avengers-clinic/model/dto"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

func ConnectDB(in dto.ConfigData, logger zerolog.Logger) (*sql.DB, error) {
	logger.Info().Msg("Trying Connect to DB...")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", in.DbConfig.Host, in.DbConfig.Port, in.DbConfig.User, in.DbConfig.Pass, in.DbConfig.Database)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to open database connection")
		return nil, err
	}

	if err = db.Ping(); err != nil {
		logger.Fatal().Err(err).Msg("Failed to ping database")
		return nil, err
	}

	logger.Info().Msg("Successfully connected to the database")
	return db, nil
}