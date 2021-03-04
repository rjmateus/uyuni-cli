package processor

type internalToolCmd struct {
	id          string
	description string
	run         func() error
}

func (tool internalToolCmd) getId() string {
	return tool.id
}

func (tool internalToolCmd) Execute() error {
	return tool.run()
}

func (tool internalToolCmd) Info() string {
	return tool.id + ": " + tool.description
}
