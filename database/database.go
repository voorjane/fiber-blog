package database

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"html/template"
	"log"
	"os"
)

type PgConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"dbname"`
}

type Post struct {
	Id       uint16
	Title    string
	Announce string
	Text     template.HTML
}

func ConnectToDB() (*gorm.DB, error) {
	var pg PgConfig
	pg.GetConf()

	cfg := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", pg.Host, pg.Username, pg.Password, pg.Database, pg.Port)
	db, err := gorm.Open(postgres.Open(cfg), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Post{})
	return db, nil
}

func (p *PgConfig) GetConf() *PgConfig {
	conf, err := os.ReadFile("database/pg.yaml")
	if err != nil {
		log.Fatalf("file not found: %v", err)
	}
	err = yaml.Unmarshal(conf, p)
	if err != nil {
		log.Fatalf("error unmarshalling yaml file: %v", err)
	}
	return p
}
