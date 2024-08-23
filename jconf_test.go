// File: "jconf_test.go"

package jconf

import (
	"fmt"
	//"os"
	"testing"

	"github.com/azorg/xlog"
	"github.com/stretchr/testify/require"
)

func init() {
	// Setup logger
	conf := xlog.Conf{
		Level: "flood", Src: true, //SrcLong: true,
		Tint: true, TimeTint: "15:04:05.999",
	}
	xlog.Env(&conf)
	xlog.Setup(conf)
	xlog.Current().SetDefaultLogs()
}

// Simple config structute
type Conf struct {
	URL  string    `json:"url"`  // some URL
	Port int       `json:"port"` // listen port
	Log  xlog.Conf `json:"log"`  // logger settings
}

const JSON_CONF_FILE_NAME = "simple.jconf"
const YAML_CONF_FILE_NAME = "simple.yaml"

// Show JSON/YAML config
func Test_Show(t *testing.T) {
	conf := &Conf{
		URL:  "prot://host.name:PORT",
		Port: 7777,
		Log:  xlog.NewConf(),
	}

	fmt.Println("JSON:")
	err := Show(conf)
	require.NoError(t, err, "Show() fail")

	fmt.Println("YAML:")
	err = ShowYAML(conf)
	require.NoError(t, err, "ShowYAML() fail")
}

// Write config (JSON+YAML)
func Test_Write(t *testing.T) {
	conf := &Conf{
		URL:  "https://host.name:8443",
		Port: 6666,
		Log:  xlog.NewConf(),
	}

	err := Write(conf, JSON_CONF_FILE_NAME)
	require.NoError(t, err, "Write() JSON fail")

	err = Write(conf, YAML_CONF_FILE_NAME)
	require.NoError(t, err, "Write() YAML fail")
}

// Read and show config (JSON+YAML)
func Test_Read(t *testing.T) {
	conf := &Conf{}

	err := Read(conf, JSON_CONF_FILE_NAME)
	require.NoError(t, err, "Read() JSON fail")

	err = Show(conf)
	require.NoError(t, err, "Show() JSON fail")

	err = os.Remove(JSON_CONF_FILE_NAME)
	require.NoError(t, err, "os.Remove() JSON fail")

	err = Read(conf, YAML_CONF_FILE_NAME)
	require.NoError(t, err, "Read() YAML fail")

	err = ShowYAML(conf)
	require.NoError(t, err, "Show() YAML fail")

	err = os.Remove(YAML_CONF_FILE_NAME)
	require.NoError(t, err, "os.Remove() YAML fail")
}

// EOF: "jconf_test.go"
