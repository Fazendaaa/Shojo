package shojo

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dgraph-io/badger/v3"
	"gopkg.in/yaml.v3"
)

func fetchDataBase() (database *badger.DB, fail error) {
	dirname, fail := os.UserHomeDir()

	if fail != nil {
		return database, fmt.Errorf("%w;\nerror while reading user folder to init shojo's database directory", fail)
	}

	databasePath := filepath.Join(dirname, ".shojo")

	if _, fail = os.Stat(databasePath); nil != fail {
		fail = os.MkdirAll(databasePath, os.ModePerm)
	}

	if nil != fail {
		return database, fmt.Errorf("%w;\nerror while creating shojo's database directory", fail)
	}

	options := badger.DefaultOptions(databasePath)
	IsModeDebug := true

	if IsModeDebug {
		options.Logger = nil
	}

	database, fail = badger.Open(options)

	if fail != nil {
		return database, fmt.Errorf("%w;\nerror while opening shojo's database", fail)
	}

	return database, fail
}

func addPackageToDatabase(packageName string) (fail error) {
	database, fail := fetchDataBase()

	if nil != fail {
		return fmt.Errorf("%w;\nerror fecthing shojo's database to add '%s' package", fail, packageName)
	}

	defer database.Close()

	fail = database.Update(func(transaction *badger.Txn) (fail error) {
		fail = transaction.Set([]byte(packageName), []byte("42"))

		if nil != fail {
			return fmt.Errorf("%w;\nerror adding '%s' package entry to database", fail, packageName)
		}

		return fail
	})

	if nil != fail {
		return fmt.Errorf("%w;\nerror while finishing up shojo's database after adding '%s' package info to it", fail, packageName)
	}

	return fail
}

func readPackage(packageName string) (packageInfo Package, fail error) {
	database, fail := fetchDataBase()

	if nil != fail {
		return packageInfo, fmt.Errorf("%w;\nerror fecthing shojo's database to read '%s' package", fail, packageName)
	}

	var valueCopy []byte

	defer database.Close()

	fail = database.View(func(transaction *badger.Txn) (fail error) {
		item, fail := transaction.Get([]byte(packageName))

		if nil != fail {
			return fmt.Errorf("%w;\nerror fetching '%s' package entry from database", fail, packageName)
		}

		valueCopy, fail = item.ValueCopy(nil)

		if nil != fail {
			return fmt.Errorf("%w;\nerror fetching '%s' package data from database", fail, packageName)
		}

		return fail
	})

	if nil != fail {
		return packageInfo, fmt.Errorf("%w;\nerror updating database after fetching '%s' package data from it", fail, packageName)
	}

	fail = yaml.Unmarshal(valueCopy, &valueCopy)

	if nil != fail {
		return packageInfo, fmt.Errorf("%w;\nerror updating database after fetching '%s' package data from it", fail, packageName)
	}

	return packageInfo, fail
}

func updatePackage(packageName string) (packageInfo Package, fail error) {
	database, fail := fetchDataBase()

	if nil != fail {
		return packageInfo, fmt.Errorf("%w;\nerror fecthing shojo's database to update '%s' package", fail, packageName)
	}

	var valueCopy []byte

	defer database.Close()

	fail = database.Update(func(transaction *badger.Txn) (fail error) {
		fail = transaction.Set([]byte("answer"), []byte("24"))

		if nil != fail {
			return fail
		}

		return fail
	})

	if nil != fail {
		return packageInfo, fail
	}

	fmt.Println("rm package: %s", valueCopy)

	return packageInfo, fail
}

func rmPackage(packageName string) (fail error) {
	database, fail := fetchDataBase()

	if nil != fail {
		return fmt.Errorf("%w;\nerror fecthing shojo's database to remove '%s' package", fail, packageName)
	}

	defer database.Close()

	fail = database.Update(func(transaction *badger.Txn) error {
		return transaction.Delete([]byte("answer"))
	})

	if nil != fail {
		return fail
	}

	return fail
}
