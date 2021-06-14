package db

import (
	"log"
	"path/filepath"
	"sync"
	"time"

	"github.com/coredevelopment/workflow/internal/constants"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Here are the internal methods listed to perform db operations
// Workflow Config schema to store in db
type workflowConfig struct {
	ID           uint
	WorkflowID   string `gorm:"index"`
	WorkflowName string
	Version      string
	Deleted      bool
	Active       bool
	FileExt      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

var (
	dbInstance *gorm.DB
	once       sync.Once
)

// func main() {
// 	InsertOrUpdateConfig("wf1", "workflow_1", "v2")
// 	InsertOrUpdateConfig("wf1", "workflow_1", "v1")
// 	InsertOrUpdateConfig("wf2", "workflow_2", "v1")
// 	InsertOrUpdateConfig("wf1", "workflow_updated", "v2")
// 	fmt.Println("----------------------------------------------------------------")
// 	fmt.Println(getSingleConfig("wf1", "v2"))
// 	fmt.Println("----------------------------------------------------------------")
// 	fmt.Println(getAllConfigs())
// 	fmt.Println(updateActiveStatus("wf1", "v1", false))
// 	fmt.Println("----------------------------------------------------------------")
// 	fmt.Println(getActiveConfigs())
// 	// deleteConfig("wf1", "v2")
// }

// getInstance : to create single instance of the database
func getInstance() *gorm.DB {
	dbPath := filepath.FromSlash(constants.DB_PATH)
	// dbPath := "../../configs/engine-configs/workflow.db"
	if dbInstance == nil {
		once.Do(
			func() {
				var err error
				// log.Println("Creating Single Instance Now")
				dbInstance, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

				// Creating DB schemas
				dbInstance.AutoMigrate(&workflowConfig{})

				if err != nil {
					log.Println("Error connecting to Database")
				}
			})
	}
	return dbInstance
}

// getSingleConfig : to get single workflow config
func getSingleConfig(workflowId, version string) workflowConfig {
	db := getInstance()
	var config workflowConfig
	result := db.Where("workflow_id = ? AND version = ? AND deleted = ?", workflowId, version, false).Find(&config)
	if result.Error != nil {
		log.Println("Error getting data :", result.Error)
	}
	return config
}

// getActiveConfig : to get active workflow config
func getActiveConfig(workflowId string) workflowConfig {
	db := getInstance()
	var config workflowConfig
	result := db.Where("workflow_id = ? AND deleted = ?", workflowId, false).Find(&config)
	if result.Error != nil {
		log.Println("Error getting data :", result.Error)
	}
	return config
}

// getAllConfigs - to get all the configs present in database
func getAllConfigs() []workflowConfig {
	db := getInstance()
	var config []workflowConfig
	result := db.Where("deleted = ?", false).Find(&config)
	if result.Error != nil {
		log.Println("Error getting data :", result.Error)
	}
	return config
}

// getActiveConfigs - to get all the configs present in database
func getActiveConfigs() []workflowConfig {
	db := getInstance()
	var config []workflowConfig
	result := db.Where("deleted = ? AND active = ?", false, true).Find(&config)
	if result.Error != nil {
		log.Println("Error getting data :", result.Error)
	}
	return config
}

// insertConfig : to insert a new config in database
func insertConfig(workflowId, workflowName, version, extension string) error {
	db := getInstance()
	newConfig := workflowConfig{WorkflowID: workflowId, WorkflowName: workflowName, Version: version, CreatedAt: time.Now(), Deleted: false, Active: true, FileExt: extension}
	result := db.Create(&newConfig)
	return result.Error
}

// updateConfig : to update the config
func updateConfig(workflowId, workflowName, version, extension string) error {
	db := getInstance()

	result := db.Where("workflow_id = ? AND version = ? AND deleted = ?", workflowId, version, false).Updates(workflowConfig{WorkflowName: workflowName})
	return result.Error
}

// updateActiveStatus : to update active status of config
func updateActiveStatus(workflowId, version string, active bool) error {
	db := getInstance()

	result := db.Table("workflow_configs").Where("( workflow_id = ? OR workflow_name = ? ) AND version = ? AND deleted = ?", workflowId, workflowId, version, false).Update("active", active)
	return result.Error
}

// deleteConfig : to update delete flag
func deleteConfig(workflowId, version string) error {
	db := getInstance()
	result := db.Where("( workflow_id = ? OR workflow_name = ? ) AND version = ?", workflowId, workflowId, version).Updates(workflowConfig{Deleted: true})
	return result.Error
}
