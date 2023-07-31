// Package config is for reading the config file necessary for the server interface to run
// And exports the config file to the config model declared in its own package
package config

// Reading the configuration json file and transferring it to the config model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Config model for server hosting
type Config struct {
	// Host value
	Host string `json:"host"`

	// Port value
	Port string `json:"port"`

	// Front-end port
	ClientPort string `json:"client-port"`

	// Sqlite source path
	SqlitePath string `json:"sqlite"`
}

// Read from a provided config file
// Into the Config Structure from models
// Providing the host and port
func LoadServerConfig(path string) (*Config, error) {

	// Initialing the logger
	logger := log.New(os.Stdout, "CONFIG: ", log.Ldate|log.Ltime)

	// Open the json config file
	jsonFile, err := os.Open(path)

	// Handling the error
	// Returning the error and an nil as the model
	if nil != err {
		return nil, err
	}

	// Differ closing the file
	defer jsonFile.Close()

	// Logging there was no error opening the file
	logger.Println("Successfully opened " + path)

	// Read the open files as a bute array
	byteValue, err := ioutil.ReadAll(jsonFile)

	// Handling the error
	// Returning the error and an nil as the model
	if nil != err {
		return nil, err
	}

	// Initialize the mode
	var config Config

	// Unmarshal the byte value
	err = json.Unmarshal(byteValue, &config)

	// Handling the error
	// Returning the error and an nil as the model
	if nil != err {
		return nil, err
	}

	// Logging there was no error unmarshalling the json
	logger.Println("Successfully unmarshalled " + path)

	// Returning the config model
	return &config, nil
}
