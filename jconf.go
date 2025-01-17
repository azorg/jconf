// File: "jconf.go"

package jconf

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/azorg/xlog"
	//"github.com/ghodss/yaml"
	"github.com/itchyny/json2yaml"
	"sigs.k8s.io/yaml"
)

// Check file name is YAML
func IsYAML(fileName string) bool {
	for _, ext := range YAML_EXTS {
		if len(fileName) >= len(ext) &&
			fileName[len(fileName)-len(ext):] == ext {
			return true
		}
	}
	return false
}

// Marshal YAML
func ToYAML(conf any) (string, error) {
	data, err := json.Marshal(conf)
	if err != nil {
		return "", err
	}

	// Convert JSON to YAML
	// (preserves the order of mapping keys!)
	input := strings.NewReader(string(data))
	var output strings.Builder
	err = json2yaml.Convert(&output, input)
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

// Write JSON/YAML config file
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

	defer func() {
		if err = file.Close(); err != nil {
			xlog.Crit("can't close config file",
				"err", err, "fileName", fileName)
		}
	}()

	err = file.Chmod(FILE_MODE)
	if err != nil {
		xlog.Error("can't set config file mode",
			"err", err, "fileName", fileName,
			"mode", fmt.Sprintf("%04o", FILE_MODE))
		return err
	}

	if IsYAML(fileName) {
		// Save as YAML
		str, err := ToYAML(conf)
		if err != nil {
			xlog.Error("can't marshal YAML", "err", err)
			return err
		}

		_, err = file.Write([]byte(str))
		if err != nil {
			xlog.Error("can't write to YAML file",
				"err", err, "fileName", fileName)
			return err
		}
		return nil
	}

	// Save as JSON
	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	if err = enc.Encode(conf); err != nil {
		xlog.Error("can't encode and write JSON config file",
			"err", err, "fileName", fileName)
		return err
	}

	return nil
}

// Read JSON config file
func Read(conf any, fileName string) error {
	if fileName == "" {
		return nil // do nothing
	}

	if IsYAML(fileName) {
		// Read as YAML
		data, err := os.ReadFile(fileName)
		if err != nil {
			xlog.Error("can't read YAML config file; use defaults",
				"err", err, "fileName", fileName)
			return err
		}

		err = yaml.Unmarshal(data, conf)
		if err != nil {
			xlog.Fatal("can't unmarshal YAML config",
				"err", err, "fileName", fileName)
			return err
		}
		return nil
	}

	// Read as JSON
	file, err := os.Open(fileName)
	if err != nil {
		xlog.Error("can't read JSON config file; use defaults",
			"err", err, "fileName", fileName)
		return err
	}

	defer func() {
		if err = file.Close(); err != nil {
			xlog.Crit("can't close config file",
				"err", err, "fileName", fileName)
		}
	}()

	dec := json.NewDecoder(file)
	if err = dec.Decode(conf); err != nil {
		xlog.Fatal("can't decode JSON config",
			"err", err, "fileName", fileName)
		return err
	}
	return nil
}

// Show structure to stdout as JSON
func Show(conf any) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(conf); err != nil {
		xlog.Error("can't encode JSON and write to stdout", "err", err)
		return err
	}
	return nil
}

// Show structure to stdout as YAML
func ShowYAML(conf any) error {
	str, err := ToYAML(conf)
	if err != nil {
		xlog.Crit("can't marshal JSON", "err", err)
		return err
	}
	fmt.Print(str)
	return nil
}

// EOF: "jconf.go"
