package db

import (
	"log"

	"github.com/coredevelopment/workflow/internal/models"
)

// InsertOrUpdateConfig : To insert or update config in database
func InsertOrUpdateConfig(workflowId, workflowName, version string) {
	var err error
	config := getSingleConfig(workflowId, version)
	if config.WorkflowID != "" {
		err = updateConfig(workflowId, workflowName, version)
	} else {
		err = insertConfig(workflowId, workflowName, version)
	}

	if err != nil {
		log.Println("Error in updating config: ", err)
	}
}

// GetAllConfig : to get all configs
func GetAllConfig() []models.WorkflowConfig {
	configs := getAllConfigs()
	result := make([]models.WorkflowConfig, 0)

	for _, c := range configs {
		result = append(result, models.WorkflowConfig{WorkflowName: c.WorkflowName, WorkflowID: c.WorkflowID, Version: c.Version, CreatedAt: c.CreatedAt.Local().String(), UpdatedAt: c.UpdatedAt.Local().String()})
	}
	return result
}

// GetActiveConfigs : to get all configs
func GetActiveConfigs() []models.WorkflowConfig {
	configs := getActiveConfigs()
	result := make([]models.WorkflowConfig, 0)

	for _, c := range configs {
		result = append(result, models.WorkflowConfig{WorkflowName: c.WorkflowName, WorkflowID: c.WorkflowID, Version: c.Version, CreatedAt: c.CreatedAt.Local().String(), UpdatedAt: c.UpdatedAt.Local().String()})
	}
	return result
}

// GetSingelConfig : to get all configs
func GetSingelConfig(workflowID, version string) models.WorkflowConfig {
	c := getSingleConfig(workflowID, version)

	result := models.WorkflowConfig{WorkflowName: c.WorkflowName, WorkflowID: c.WorkflowID, Version: c.Version, CreatedAt: c.CreatedAt.Local().String(), UpdatedAt: c.UpdatedAt.Local().String()}

	return result
}

// ActivateConfig : to activate config
func ActivateConfig(workflowId, version string) {
	err := updateActiveStatus(workflowId, version, true)
	if err != nil {
		log.Println("Error in updating active status: ", err)
	}
}

// DeactivateConfig : to activate config
func DeactivateConfig(workflowId, version string) {
	err := updateActiveStatus(workflowId, version, false)
	if err != nil {
		log.Println("Error in updating active status: ", err)
	}
}

// DeleteConfig
func DeleteConfig(workflowId, version string) {
	err := deleteConfig(workflowId, version)
	if err != nil {
		log.Println("Error in updating active status: ", err)
	}
}
