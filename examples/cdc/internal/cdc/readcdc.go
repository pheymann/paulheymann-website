package cdc

import (
	"os"

	"gopkg.in/yaml.v2"
)

func ReadCDC(contractPath string) (*ClientDefinedContract, error) {
	contractFile, err := os.ReadFile(contractPath)
	if err != nil {
		return nil, err
	}

	var contract ClientDefinedContract
	if err := yaml.Unmarshal([]byte(contractFile), &contract); err != nil {
		return nil, err
	}

	return &contract, nil
}

func MustReadCDC(contractPath string) ClientDefinedContract {
	contract, err := ReadCDC(contractPath)
	if err != nil {
		panic(err)
	}

	return *contract
}

func MustLoadCDC(cdcPath string) ClientDefinedContract {
	rootPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return MustReadCDC(rootPath + "/../../cdc" + cdcPath)
}
