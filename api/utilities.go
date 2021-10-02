package api

import "io/ioutil"

func ReadFile(path string) ([]byte, error) { return ioutil.ReadFile(path) }

func GetData(item string) (string, error) {
	dat, err := ReadFile("./data/" + item + ".json")

	if err != nil { return "", err }

	return string(dat), nil
}
