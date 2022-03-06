package shojo

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dgraph-io/badger/v3"
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

// encodeToBytes https://gist.github.com/SteveBate/042960baa7a4795c3565#file-struct_to_bytes-go-L29
func encodeToBytes(p interface{}) []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("uncompressed size (bytes): ", len(buf.Bytes()))
	return buf.Bytes()
}

// bytesToShow https://gist.github.com/SteveBate/042960baa7a4795c3565#file-struct_to_bytes-go-L29
func bytesToShow(input []byte) (show Show, fail error) {
	dec := gob.NewDecoder(bytes.NewReader(input))
	fail = dec.Decode(&show)

	if nil != fail {
		return show, fmt.Errorf("%w;\nerror while converting information stored in database to shojo", fail)
	}

	return show, fail
}

func addPackageToDatabase(show Show) (fail error) {
	database, fail := fetchDataBase()

	if nil != fail {
		return fmt.Errorf("%w;\nerror fecthing shojo's database to add '%s' package", fail, show.Package)
	}

	defer database.Close()

	fail = database.Update(func(transaction *badger.Txn) (fail error) {
		fail = transaction.Set([]byte(show.Package), encodeToBytes(show))

		if nil != fail {
			return fmt.Errorf("%w;\nerror adding '%s' package entry to database", fail, show.Package)
		}

		return fail
	})

	if nil != fail {
		return fmt.Errorf("%w;\nerror while finishing up shojo's database after adding '%s' package info to it", fail, show.Package)
	}

	return fail
}

func readPackage(packageName string) (show Show, fail error) {
	database, fail := fetchDataBase()

	if nil != fail {
		return show, fmt.Errorf("%w;\nerror fecthing shojo's database to read '%s' package", fail, packageName)
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
		return show, fmt.Errorf("%w;\nerror updating database after fetching '%s' package data from it", fail, show.Package)
	}

	return bytesToShow(valueCopy)
}

func updatePackage(show Show) (fail error) {
	database, fail := fetchDataBase()

	if nil != fail {
		return fmt.Errorf("%w;\nerror fecthing shojo's database to update '%s' package", fail, show.Package)
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
		return fail
	}

	fmt.Println("rm package: %s", valueCopy)

	return fail
}

func rmPackage(show Show) (fail error) {
	database, fail := fetchDataBase()

	if nil != fail {
		return fmt.Errorf("%w;\nerror fecthing shojo's database to remove '%s' package", fail, show.Package)
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
