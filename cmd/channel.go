package cmd

import (
	"fmt"

	controllers "github.com/Fazendaaa/Shojo/controllers"
	"github.com/theckman/yacspin"
)

func consumeChannel(params []string,
	process string,
	spinner *yacspin.Spinner,
	resultChannel chan controllers.PackageResponse) (fail error) {
	for i := 0; i < len(params); i++ {
		result := <-resultChannel

		spinner.Message(fmt.Sprintf("'%s'", result.PackageName))

		if nil != result.Response {
			return fmt.Errorf(`\n%v;
error while %s package '%s'`, result.Response, process, result.PackageName)
		}
	}

	return fail
}
