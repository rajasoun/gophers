package model

//We need to mock "ReadFile(string)" function from "ioutil" and "Getwd()" function
//from "os packages" for unit tests.
//So, we created "readerI" interface with "readFile" and "get_wd" methods.
// This interface will have two implementations, one for main functionality and one for
// mock implementation which will return dummy values

type readerI interface {
	readFile(path string) ([]byte, error)
	get_wd() (string, error)
}
