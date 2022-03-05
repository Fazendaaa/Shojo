package shojo

import (
	"os"
	"path/filepath"

	"github.com/dgraph-io/badger/v3"
)

func initDataBase() (database *badger.DB, fail error) {
	dirname, fail := os.UserHomeDir()

	if fail != nil {
		return database, fail
	}

	databasePath := filepath.Join(dirname, ".shojo")

	if _, fail = os.Stat(databasePath); nil != fail {
		fail = os.MkdirAll(databasePath, os.ModePerm)

		if nil != fail {
			return database, fail
		}
	}

	database, fail = badger.Open(badger.DefaultOptions(dirname))

	if fail != nil {
		return database, fail
	}

	return database, fail
}

func addPackageToDatabase(packageName string) (fail error) {
	database, fail := initDataBase()

	if nil != fail {
		return fail
	}

	defer database.Close()

	return fail
}

func readPackage(packageName string) (packageInfo Package, fail error) {
	database, fail := initDataBase()

	if nil != fail {
		return packageInfo, fail
	}

	defer database.Close()

	return packageInfo, fail
}

func updatePackage(packageName string) (packageInfo Package, fail error) {
	database, fail := initDataBase()

	if nil != fail {
		return packageInfo, fail
	}

	defer database.Close()

	return packageInfo, fail
}

func rmPackage(packageName string) (packageInfo Package, fail error) {
	database, fail := initDataBase()

	if nil != fail {
		return packageInfo, fail
	}

	defer database.Close()

	return packageInfo, fail
}
