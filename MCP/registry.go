package main

import (
	"github.com/input-api/mcp-server/config"
	"github.com/input-api/mcp-server/models"
	tools_group "github.com/input-api/mcp-server/tools/group"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_group.CreateGet_groupTool(cfg),
		tools_group.CreateGet_group_yamlTool(cfg),
	}
}
