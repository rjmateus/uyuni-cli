package processor

type ToolCmd interface {
	Execute()
	Info() string
	getId() string
}

func externalToolCommand(id string, execPath string, desc string, providePackage string) ToolCmd {
	return externalToolCmd{id: id, description: desc,
		execPath: execPath, providePackage: providePackage}
}

func internalToolCommand(id string, ft func(), desc string) ToolCmd {
	return internalToolCmd{id: id, description: desc, run: ft}
}
