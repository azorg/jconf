// File: "jconf.go"

package jconf

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/azorg/xlog"
)

// Read JSON config file
func Read(conf any, fileName string) error {
	if fileName == "" {
		return nil
	}

	file, err := os.Open(fileName)
	if err != nil {
		xlog.Error("can't read config file; use defaults",
			"err", err, "fileName", fileName)
		return err
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	if err = dec.Decode(conf); err != nil {
		xlog.Fatal("can't decode JSON config", "err", err, "fileName", fileName)
	}
	return nil
}

// Write JSON config file
func Write(conf any, fileName string) error {
	if fileName == "" {
		err := errors.New("empty file name")
		xlog.Error("can't write config file", "err", err)
		return err
	}

	file, err := os.Create(fileName)
	if err != nil {
		xlog.Error("can't create config file",
			"err", err, "fileName", fileName)
		return err
	}

	err = file.Chmod(FILE_MODE)
	if err != nil {
		xlog.Error("can't set config file mode",
			"err", err, "fileName", fileName)
		return err
	}

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	if err = enc.Encode(conf); err != nil {
		xlog.Error("can't encode and write config file",
			"err", err, "fileName", fileName)
		return err
	}

	if err = file.Close(); err != nil {
		xlog.Error("can't close config file",
			"err", err, "fileName", fileName)
		return err
	}
	return nil
}

// Show JSON config structure to stdout
func Show(conf any) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(conf); err != nil {
		xlog.Error("can't encode JSON and write to stdout", "err", err)
		return err
	}
	return nil
}

// EOF: "jconf.go"
