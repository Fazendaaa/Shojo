package shojo

func tlmgr(subcommand string, paramters []string) (fail error) {
	_, fail = shell("tlmgr", append([]string{subcommand}, paramters...))

	return fail
}

func install(packageName string) (fail error) {
	return tlmgr("install", []string{packageName})
}
