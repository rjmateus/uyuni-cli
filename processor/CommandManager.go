package processor

import (
	"fmt"
	"github.com/rmateus/sumatools/newTool"
	"os"
)

const (
	usage = `SUMA tools 
Usage should be "sumatools [command]"

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
	manager.registerTool(remoteToolCommand("spacewalk-sql", "/usr/bin/spacewalk-sql"))
	manager.registerTool(remoteToolCommand("spacewalk-repo-sync", "/usr/bin/spacewalk-repo-sync"))
	manager.registerTool(remoteToolCommand("satpasswd", "/usr/bin/satpasswd"))
	manager.registerTool(remoteToolCommand("spacecmd", "/usr/bin/spacecmd"))
	manager.registerTool(localToolCommand("newTool", newTool.ProcessNewTool))
	return manager
}