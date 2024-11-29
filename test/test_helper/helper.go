package testhelper

import (
	"data-spaces-backend/presentation/http/echo/handler"
	mocks "data-spaces-backend/test/mock"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewMockHandler(host string) handler.OuranosHandler {
	cfpUsecase := new(mocks.ICfpUsecase)
	cfpHandler := handler.NewCfpHandler(cfpUsecase)
	cfpCertificationUsecase := new(mocks.ICfpCertificationUsecase)
	cfpCertificationHandler := handler.NewCfpCertificationHandler(cfpCertificationUsecase)
	partsUsecase := new(mocks.IPartsUsecase)
	partsStructureUsecase := new(mocks.IPartsStructureUsecase)
	partsHandler := handler.NewPartsHandler(partsUsecase, partsStructureUsecase, host)
	partsStructureHandler := handler.NewPartsStructureHandler(partsStructureUsecase)
	tradeUsecase := new(mocks.ITradeUsecase)
	tradeHandler := handler.NewTradeHandler(tradeUsecase, host)
	statusUsecase := new(mocks.IStatusUsecase)
	statusHandler := handler.NewStatusHandler(statusUsecase, host)
	h := handler.NewOuranosHandler(cfpHandler, cfpCertificationHandler, partsHandler, partsStructureHandler, tradeHandler, statusHandler)

	return h
}

func NewMockDB() (*gorm.DB, error) {
	// DBのコネクションを初期化
	uuid, _ := uuid.NewRandom()
	dbStr := "file:memdb_" + uuid.String() + "?mode=memory&cache=shared"
	db, err := gorm.Open(sqlite.Open(dbStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = initTable(db)
	if err != nil {
		return nil, err
	}
	err = insertSeeders(db)
	if err != nil {
		return nil, err
	}
	return db, err
}

func initTable(db *gorm.DB) error {
	// init DB connection
	var err error

	// get DDL files
	setupDir := "../../../setup/migrations_sqlite"
	files, err := os.ReadDir(setupDir)
	if err != nil {
		return err
	}

	// read DDL
	var upQueries []string
	var downQueries []string
	for _, file := range files {
		b, err := os.ReadFile(fmt.Sprintf("%s/%s", setupDir, file.Name()))
		if err != nil {
			return err
		}
		ls := strings.Split(string(b), "\n")
		l := strings.Join(ls, "")
		qs := strings.Split(l, ";")
		if file.Name()[len(file.Name())-6:] == "up.sql" {
			upQueries = append(upQueries, qs[:len(qs)-1]...)
		} else if file.Name()[len(file.Name())-8:] == "down.sql" {
			downQueries = append(downQueries, qs[:len(qs)-1]...)
		} else {
			return fmt.Errorf("unknown file(%s) is founded", file.Name())
		}
	}

	// execute down DDL by reverse order
	for idx := 0; idx < len(downQueries); idx++ {
		err := db.Exec(fmt.Sprintf("%s;", downQueries[len(downQueries)-idx-1])).Error
		if err != nil {
			return err
		}
	}

	// execute up DDL by normal order
	for idx := 0; idx < len(upQueries); idx++ {
		err := db.Exec(fmt.Sprintf("%s;", upQueries[idx])).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func insertSeeders(db *gorm.DB) error {
	var err error

	// get DML files
	setupDir := "../../../setup/seeders_sqlite"
	files, err := os.ReadDir(setupDir)
	if err != nil {
		return err
	}

	// read DML
	var queries []string
	for _, file := range files {
		b, err := os.ReadFile(fmt.Sprintf("%s/%s", setupDir, file.Name()))
		if err != nil {
			return err
		}
		ls := strings.Split(string(b), "\n")
		l := strings.Join(ls, "")
		qs := strings.Split(l, ";")
		queries = append(queries, qs[:len(qs)-1]...)
	}

	// execute DML
	for idx := 0; idx < len(queries); idx++ {
		err := db.Exec(fmt.Sprintf("%s;", queries[idx])).Error
		if err != nil {
			return err
		}
	}
	return nil
}
