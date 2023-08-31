package db

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zapadapter"
	"go.uber.org/zap"
)

func Setup(dsn string) *sqlx.DB {
	zapCfg := zap.NewDevelopmentConfig()
	zapCfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel) // whatever minimum level
	zapCfg.DisableCaller = true
	logger, _ := zapCfg.Build()

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}

	loggerAdapter := zapadapter.New(logger)
	db = sqldblogger.OpenDriver(dsn, db.Driver(), loggerAdapter)
	return sqlx.NewDb(db, "pgx")
}
