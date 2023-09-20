package db_models

import (
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type DynamicModel struct {
	ID   uint                   `gorm:"primaryKey"`
	Data map[string]interface{} `gorm:"type:jsonb"`
}

func InitModels(db *gorm.DB) {
	if !db.Migrator().HasTable(&DynamicModel{}) {
		db.AutoMigrate(&DynamicModel{})
		log.Printf("Model created %v", DynamicModel{})
	}
}

func CreateRecord(db *gorm.DB, jsonData []byte) {
	entry := DynamicModel{
		Data: map[string]interface{}{},
	}
	err := json.Unmarshal(jsonData, &entry.Data)
	if err != nil {
		log.Println(err)
	}
	db.Create(&entry)
	fmt.Println("Record created")
}
