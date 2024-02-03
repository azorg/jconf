// File: "jconf_test.go"

package jconf

import (
	"os"
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

const CONF_FILE_NAME = "simple.jconf"

// Show config
func Test_Show(t *testing.T) {
	conf := &Conf{
		URL:  "prot://host.name:PORT",
		Port: 7777,
		Log:  xlog.NewConf(),
	}

	err := Show(conf)
	require.NoError(t, err, "Show() fail")
}

// Write config
func Test_Write(t *testing.T) {
	conf := &Conf{
		URL:  "https://host.name:8443",
		Port: 6666,
		Log:  xlog.NewConf(),
	}

	err := Write(conf, CONF_FILE_NAME)
	require.NoError(t, err, "Write() fail")
}

// Read and show config
func Test_Read(t *testing.T) {
	conf := &Conf{}

	err := Read(conf, CONF_FILE_NAME)
	require.NoError(t, err, "Read() fail")

	err = Show(conf)
	require.NoError(t, err, "Show() fail")

	err = os.Remove(CONF_FILE_NAME)
	require.NoError(t, err, "os.Remove() fail")
}

// EOF: "jconf_test.go"
