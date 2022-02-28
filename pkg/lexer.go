package shojo

import (
	"fmt"
	"reflect"
	"strconv"
)

func texToProject(origin map[interface{}]interface{}) (tex Tex, fail error) {
	var converted bool = true
	var value map[string]interface{}

	if _, ok := origin["tex"]; ok {
		tex.Version, converted = origin["tex"].(string)
	}
	if !converted {
		switch origin["tex"].(type) {
		default:
			value, converted = origin["tex"].(map[string]interface{})
		case float64:
			tex.Version = strconv.FormatFloat(origin["tex"].(float64), 'f', -1, 64)

			return tex, fail
		}

		if !converted {
			return tex, fmt.Errorf("Tex definition presented and it's also malformed")
		}

		switch value["version"].(type) {
		default:
			return tex, fmt.Errorf(`TLMGR version definition presented and it's also
	malformed, expected 'string' and got '%s'`, reflect.TypeOf(value["version"]))
		case string:
			tex.Version, converted = value["version"].(string)
		case float64:
			tex.Version = strconv.FormatFloat(value["version"].(float64), 'f', -1, 64)
		}

		if !converted {
			return tex, fmt.Errorf(`TLMGR version definition presented and it's also
malformed type during conversion, got unexpected '%s' type`,
				reflect.TypeOf(value["version"]))
		}
	}

	return tex, fail
}

func tlmgrToProject(origin map[interface{}]interface{}) (tlmgr TLMGR, fail error) {
	var value int = 0
	var converted bool = true

	if _, ok := origin["tlmgr"]; ok {
		value, converted = origin["tlmgr"].(int)

		if converted {
			tlmgr.Version = strconv.Itoa(value)
		}
	}
	if !converted {
		result, converted := origin["tlmgr"].(map[string]interface{})

		if !converted {
			return tlmgr, fmt.Errorf("TLMGR definition presented and it's also malformed")
		}

		switch result["version"].(type) {
		default:
			return tlmgr, fmt.Errorf(`TLMGR version definition presented and it's also
	malformed, expected 'string' and got '%s'`, reflect.TypeOf(result["version"]))
		case string:
			tlmgr.Version, _ = result["version"].(string)
		case int:
			tlmgr.Version = strconv.Itoa(result["version"].(int))
		}
	}

	return tlmgr, fail
}

func repositoryToProject(origin map[interface{}]interface{}) (repository Repository, fail error) {
	var converted bool = true

	if _, ok := origin["repository"]; ok {
		repository.URL, converted = origin["repository"].(string)
	}
	if !converted {
		value, converted := origin["repository"].(map[string]interface{})

		if !converted {
			return repository, fmt.Errorf("repository definition presented and it's also malformed")
		}

		_, ok := value["url"]

		if ok {
			repository.URL, converted = value["url"].(string)
		}
		if !converted || !ok {
			return repository, fmt.Errorf("repository url definition presented and it's also malformed")
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
