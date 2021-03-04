package processor

import (
	"errors"
	"fmt"
)

const (
	usage = `uyuni tools 
Usage should be "uyuni-cli [command]"

Available Command:`
)

type ToolsCommandManager struct {
	commandsOrdered []string
	toolsCommands   map[string]ToolCmd
}

func (manager *ToolsCommandManager) Execute(toolsCommandId string) error {
	if value, ok := manager.toolsCommands[toolsCommandId]; ok {
		return value.Execute()
	} else {
		manager.UsagePrint()
		return errors.New("command not found")
	}
}

func (manager *ToolsCommandManager) UsagePrint() {
	fmt.Println(usage)
	for _, toolID := range manager.commandsOrdered {
		cmd := manager.toolsCommands[toolID]
		fmt.Println("  - ", cmd.Info())
	}
}

func (manager *ToolsCommandManager) registerTool(tool ToolCmd) {
	manager.commandsOrdered = append(manager.commandsOrdered, tool.getId())
	manager.toolsCommands[tool.getId()] = tool
}

func GetToolsCommandManager() ToolsCommandManager {
	manager := ToolsCommandManager{make([]string, 0), make(map[string]ToolCmd)}
	manager.registerTool(externalToolCommand("spacewalk-sql", "/usr/bin/spacewalk-sql", "execute sql command directly on database", "susemanager-schema"))
	manager.registerTool(externalToolCommand("spacewalk-repo-sync", "/usr/bin/spacewalk-repo-sync", "start repository synchronization", "spacewalk-backend-tools"))
	manager.registerTool(externalToolCommand("spacewalk-debug", "/usr/bin/spacewalk-debug", "export debug information", "spacewalk-backend-tools"))
	manager.registerTool(externalToolCommand("satpasswd", "/usr/bin/satpasswd", "reset user password", "spacewalk-backend-tools"))
	manager.registerTool(externalToolCommand("spacecmd", "/usr/bin/spacecmd", "command line tool to perform actions", "spacecmd"))
	manager.registerTool(externalToolCommand("satwho", "/usr/bin/satwho", "show which user is being used for authentication", "spacewalk-backend-tools"))
	return manager
}
