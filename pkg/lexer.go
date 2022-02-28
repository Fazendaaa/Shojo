package shojo

import (
	"fmt"
	"reflect"
	"strconv"
)

func texToProject(origin map[interface{}]interface{}) (tex Tex, fail error) {
	if _, ok := origin["tex"]; ok {
		result, converted := origin["tex"].(map[interface{}]interface{})

		if !converted {
			return tex, fmt.Errorf("Tex definition presented and it's also malformed")
		}

		tex.Version, converted = result["version"].(string)

		if !converted {
			return tex, fmt.Errorf(`Tex version definition presented and it's also
malformed, expected 'string' and got '%s'`, reflect.TypeOf(tex.Version))
		}
	}

	return tex, fail
}

func tlmgrToProject(origin map[interface{}]interface{}) (tlmgr TLMGR, fail error) {
	if _, ok := origin["tlmgr"]; ok {
		result, converted := origin["tlmgr"].(map[interface{}]interface{})

		if !converted {
			return tlmgr, fmt.Errorf("TLMGR definition presented and it's also malformed")
		}

		version, converted := result["version"].(int)

		if !converted {
			return tlmgr, fmt.Errorf(`TLMGR version definition presented and it's
also malformed, expected 'string' and got '%s'`, reflect.TypeOf(version))
		}

		tlmgr.Version = strconv.Itoa(version)
	}

	return tlmgr, fail
}

func repositoryToProject(origin map[interface{}]interface{}) (repository Repository, fail error) {
	var converted bool = true

	if _, ok := origin["repository"]; ok {
		repository.URL, converted = origin["repository"].(string)
	}
	if !converted {
		value, converted := origin["repository"].(map[interface{}]interface{})

		if !converted {
			return repository, fmt.Errorf("repository definition presented and it's also malformed")
		}

		if _, ok := value["url"]; ok {
			repository.URL, converted = value["url"].(string)

			if !converted {
				return repository, fmt.Errorf("repository url definition presented and it's also malformed")
			}
		}
	}

	return repository, fail
}

func interfaceToPackage(origin interface{}) (result Package, fail error) {
	read, ok := origin.(string)

	if !ok {
		return result, fmt.Errorf("Package definition is not a string")
	}

	result.Name = read

	return result, fail
}

func packagesToProject(origin map[interface{}]interface{}) (packages []Package, fail error) {
	read, ok := origin["packages"].([]interface{})

	if !ok {
		return packages, fmt.Errorf("packages definition is not presented")
	}

	packages = make([]Package, len(read))

	for index, data := range read {
		packages[index], fail = interfaceToPackage(data)

		if nil != fail {
			return packages, fmt.Errorf("%w;\nmalformed packages definition", fail)
		}
	}

	return packages, fail
}

// projectFunc handles the pre-processing stage of the config file, just
// checking whether or not the whole file structure seems right -- it doesn't
// check whether or not it's right for execution, many values could be set as
// wrong; like a package name with some typo
func projectFunc(origin map[interface{}]interface{}) (_ interface{}, fail error) {
	result := Project{}
	result.Tex, fail = texToProject(origin)

	if nil != fail {
		return result, fmt.Errorf("%w;\nmalformed tex definition", fail)
	}

	result.TLMGR, fail = tlmgrToProject(origin)

	if nil != fail {
		return result, fmt.Errorf("%w;\nmalformed tlmgr definition", fail)
	}

	result.Repository, fail = repositoryToProject(origin)

	if nil != fail {
		return result, fmt.Errorf("%w;\nmalformed repository definition", fail)
	}

	result.Packages, fail = packagesToProject(origin)

	if nil != fail {
		return result, fmt.Errorf("%w;\nmalformed packages definition", fail)
	}

	return result, fail
}
