package cli

func Run() int {
	var a = app{
		filters: initFilters(),
		config:  getConfig(),
	}
	return a.op()
}
