package processor

import (
	"fmt"
	"os"
)

const (
	usage = `SUMA tools 
Usage should be "uyuni-cli [command]"

Available Command:`
)

type ToolsCommandManager struct {
	commandsOrdered [] string
	toolsCommands map[string]toolCmd
}

func (manager *ToolsCommandManager) Execute(toolsCommandId string)  {
	if value, ok:= manager.toolsCommands[toolsCommandId]; ok {
		value.Execute()
	}else{
		manager.UsagePrint()
		os.Exit(1)
	}
}

func (manager *ToolsCommandManager) UsagePrint() {
	fmt.Println(usage)
	for _, toolID := range manager.commandsOrdered{
		cmd := manager.toolsCommands[toolID]
		fmt.Println("  - ", cmd.Info())
	}
}

func (manager *ToolsCommandManager) registerTool(tool toolCmd) {
	manager.commandsOrdered = append(manager.commandsOrdered, tool.Id)
	manager.toolsCommands[tool.Id] = tool
}

func GetToolsCommandManager() ToolsCommandManager {
	manager := ToolsCommandManager{make([]string,0), make(map[string]toolCmd)}
	manager.registerTool(externalToolCommand("spacewalk-sql", "/usr/bin/spacewalk-sql", "execute sql command directly on database"))
	manager.registerTool(externalToolCommand("spacewalk-repo-sync", "/usr/bin/spacewalk-repo-sync", "start repository synchronization"))
	manager.registerTool(externalToolCommand("spacewalk-debug", "/usr/bin/spacewalk-debug", "export debug information"))
	manager.registerTool(externalToolCommand("satpasswd", "/usr/bin/satpasswd", "reset user password"))
	manager.registerTool(externalToolCommand("spacecmd", "/usr/bin/spacecmd", "command line tool to perform actions"))
	//manager.registerTool(externalToolCommand("satwho", "/usr/bin/satwho", "show which user is being used for authentication"))
	return manager
}