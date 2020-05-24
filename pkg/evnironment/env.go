// !build js,wasm

package environment

import "os"

var (
	defaultDataPath   = os.Getenv("HOME") + "/.local/share"
	defaultConfigPath = os.Getenv("HOME") + "/.config"
	name              = "/pspk"
	mode              = os.FileMode(0766)
)

func LoadDataPath() (path string) {
	defer func() {
		_ = os.Mkdir(path, mode)
	}()
	env, ok := os.LookupEnv("XDG_DATA_HOME")
	if !ok {
		return defaultDataPath + name
	}
	return env + name
}

func LoadConfigPath() (path string) {
	defer func() {
		_ = os.Mkdir(path, mode)
	}()
	env, ok := os.LookupEnv("XDG_CONFIG_HOME")
	if !ok {
		return defaultConfigPath + name
	}
	return env + name
}
