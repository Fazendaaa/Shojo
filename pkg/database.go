package shojo

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dgraph-io/badger/v3"
)

// encodeToBytes https://gist.github.com/SteveBate/042960baa7a4795c3565#file-struct_to_bytes-go-L29
func encodeToBytes(p interface{}) (tmp []byte, fail error) {
	buffer := bytes.Buffer{}
	enc := gob.NewEncoder(&buffer)
	fail = enc.Encode(p)

	if nil != fail {
		return tmp, fmt.Errorf("%w\n;error while encoded data to database", fail)
	}

	return buffer.Bytes(), fail
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

func (project *Project) fetchDataBase() (fail error) {
	dirname, fail := os.UserHomeDir()

	if nil != fail {
		return fmt.Errorf("%w;\nerror while reading user folder to init shojo's database directory", fail)
	}

	databasePath := filepath.Join(dirname, ".shojo")

	if _, fail = os.Stat(databasePath); nil != fail {
		fail = os.MkdirAll(databasePath, os.ModePerm)
	}

	if nil != fail {
		return fmt.Errorf("%w;\nerror while creating shojo's database directory", fail)
	}

	options := badger.DefaultOptions(databasePath)
	IsModeDebug := true

	if IsModeDebug {
		options.Logger = nil
	}

	project.database, fail = badger.Open(options)

	if nil != fail {
		return fmt.Errorf("%w;\nerror while opening shojo's database", fail)
	}

	return fail
}

func (project *Project) addPackageToDatabase(show Show) (fail error) {
	data, fail := encodeToBytes(show)

	if nil != fail {
		return fmt.Errorf("%w;\nerror while generating shojo's database to add '%s' package", fail, show.Package)
	}

	transaction := project.database.NewTransaction(true)

	defer transaction.Discard()

	fail = transaction.Set([]byte(show.Package), data)

	if nil != fail {
		return fmt.Errorf("%w;\nerror adding '%s' package entry to database", fail, show.Package)
	}

	fail = transaction.Commit()

	if nil != fail {
		return fmt.Errorf("%w;\nerror after adding '%s' package entry to database", fail, show.Package)
	}

	return fail
}

func (project *Project) readPackage(packageName string) (show Show, fail error) {
	var valueCopy []byte
	transaction := project.database.NewTransaction(false)

	defer transaction.Discard()

	item, fail := transaction.Get([]byte(packageName))

	if nil != fail {
		return show, fmt.Errorf("%w;\nerror fetching '%s' package entry from database", fail, packageName)
	}

	valueCopy, fail = item.ValueCopy(nil)

	if nil != fail {
		return show, fmt.Errorf("%w;\nerror fetching '%s' package data from database", fail, packageName)
	}

	return bytesToShow(valueCopy)
}

func (project *Project) updatePackage(show Show) (fail error) {
	data, fail := encodeToBytes(show)

	if nil != fail {
		return fmt.Errorf("%w;\nerror while generating shojo's database to update '%s' package", fail, show.Package)
	}

	transaction := project.database.NewTransaction(true)

	defer transaction.Discard()

	fail = transaction.Set([]byte(show.Package), data)

	if nil != fail {
		return fmt.Errorf("%w;\nerror updating '%s' package entry to database", fail, show.Package)
	}

	fail = transaction.Commit()

	if nil != fail {
		return fmt.Errorf("%w;\nerror after updating '%s' package entry to database", fail, show.Package)
	}

	return fail
}

func (project *Project) rmPackage(show Show) (fail error) {
	defer project.database.Close()

	transaction := project.database.NewTransaction(true)

	defer transaction.Discard()

	fail = transaction.Delete([]byte(show.Package))

	if nil != fail {
		return fmt.Errorf("%w;\nerror removing '%s' package entry to database", fail, show.Package)
	}

	fail = transaction.Commit()

	if nil != fail {
		return fmt.Errorf("%w;\nerror after removing '%s' package entry to database", fail, show.Package)
	}

	return fail
}
