package db

import (
	"log"
)

// InsertOrUpdateConfig : To insert or update config in database
func InsertOrUpdateConfig(workflowId, workflowName, version, extension string) error {
	var err error
	config := getSingleConfig(workflowId, version)
	if config.WorkflowID != "" {
		err = updateConfig(workflowId, workflowName, version, extension)
	} else {
		err = insertConfig(workflowId, workflowName, version, extension)
	}

	if err != nil {
		log.Println("Error in updating config: ", err)
		return err
	}
	return nil
}

// GetAllConfig : to get all configs
func GetAllConfig() []WorkflowConfig {
	configs := getAllConfigs()
	result := make([]WorkflowConfig, 0)

	for _, c := range configs {
		result = append(result, WorkflowConfig{WorkflowName: c.WorkflowName, WorkflowID: c.WorkflowID, Version: c.Version, CreatedAt: c.CreatedAt.Local().String(), UpdatedAt: c.UpdatedAt.Local().String()})
	}
	return result
}

// GetActiveConfigs : to get all configs
func GetActiveConfigs() []WorkflowConfig {
	configs := getActiveConfigs()
	result := make([]WorkflowConfig, 0)

	for _, c := range configs {
		result = append(result, WorkflowConfig{WorkflowName: c.WorkflowName, WorkflowID: c.WorkflowID, Version: c.Version, CreatedAt: c.CreatedAt.Local().String(), UpdatedAt: c.UpdatedAt.Local().String()})
	}
	return result
}

// GetSingelConfig : to get all configs
func GetSingelConfig(workflowID, version string) WorkflowConfig {
	c := getSingleConfig(workflowID, version)

	result := WorkflowConfig{WorkflowName: c.WorkflowName, WorkflowID: c.WorkflowID, Version: c.Version, CreatedAt: c.CreatedAt.Local().String(), UpdatedAt: c.UpdatedAt.Local().String()}

	return result
}

// GetActiveConfig : to get active config of the workflow
func GetActiveConfig(workflowID string) WorkflowConfig {
	c := getActiveConfig(workflowID)
	result := WorkflowConfig{WorkflowID: c.WorkflowID, Version: c.Version, Active: c.Active, FileExt: c.FileExt}

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
func DeactivateConfig(workflowId, version string) error {
	err := updateActiveStatus(workflowId, version, false)
	if err != nil {
		log.Println("Error in updating active status: ", err)
		return err
	}
	return nil
}

// DeleteConfig
func DeleteConfig(workflowId, version string) error {
	err := deleteConfig(workflowId, version)
	if err != nil {
		log.Println("Error in updating active status: ", err)
		return err
	}
	return nil
}
