package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

// DbConfig defines all the configurations required to connect to db
type DbConfig struct {
	Address      string `yaml:"address"`
	Port         string `yaml:"port"`
	DatabaseName string `yaml:"databaseName"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
}

// DbCtxt defines the connections
type DbCtxt struct {
	client *gorm.DB
}

// GetNewDbCtxt gets new db ctxt object
func GetNewDbCtxt() *DbCtxt {
	db := &DbCtxt{}
	return db
}

// ConnectToDb creates connection to the db
func (db *DbCtxt) ConnectToDb(config DbConfig) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Username,
		config.Password,
		config.Address,
		config.Port,
		config.DatabaseName)
	log.Debugf("dsn [%s]", dsn)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Errorf("error in connecting to mysql - [%s]", err.Error())
		return err
	}
	db.client = conn
	return nil
}

// InitDatabase initiliases the database
// only needed when tables needs to be created
func (db *DbCtxt) InitDatabase() error {
	var models []interface{}
	models = append(models,
		&Hotel{},
		&Room{},
		&RatePlan{},
	)
	for _, model := range models {
		err := db.client.AutoMigrate(model)
		if err != nil {
			return err
		}
	}
	return nil
}
