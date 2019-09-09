package context

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type dbContext struct {
	dbContext *gorm.DB
}

func newDbContext() dbContext {
	return dbContext{}
}

func (context *dbContext) open(connectString string) {
	db, err := gorm.Open("postgres", connectString)
	if err != nil {
		log.Print("Can't open db : %v\n ", err.Error())
		return
	}
	context.dbContext = db
}

func (context *dbContext) close() {
	err := context.dbContext.Close()
	if err != nil {
		log.Print("Error when close db : %v\n", err.Error())
	}
}

