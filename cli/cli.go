package cli

func Run() int {
	var a = app{
		filters: initFilters(),
		config:  initConfig(),
	}
	return a.op()
}
