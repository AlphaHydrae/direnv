package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"syscall"
)

// `direnv exec DIR <COMMAND> ...`
var CmdExec = &Cmd{
	Name: "exec",
	Desc: "Executes a command after loading the first .envrc found in DIR",
	Args: []string{"DIR", "COMMAND"},
	Fn: func(env Env, args []string) (err error) {
		var rcPath string
		var config *Config
		var oldEnv Env
		var newEnv Env

		if len(args) < 3 {
			err = fmt.Errorf("missing COMMAND argument")
			return
		}

		if len(args) > 1 {
			rcPath = args[1]
		} else {
			if rcPath, err = os.Getwd(); err != nil {
				return
			}
		}

		if config, err = LoadConfig(env); err != nil {
			return
		}

		rc := FindRC(rcPath, config.AllowDir())
		if rc == nil {
			return fmt.Errorf(".envrc not found")
		}

		// Restore pristine environment first
		oldEnv, err = config.EnvBackup()
		if err != nil {
			oldEnv = env
			err = nil
		}

		newEnv, err = rc.Load(config, oldEnv)

		exepath, _ := lookPath(args[2], newEnv["PATH"])
		log("exepath=%s", exepath)
		err = syscall.Exec(exepath, args[2:], newEnv.ToGoEnv())
		return
	},
}

// Similar to os/exec.LookPath except we pass in the PATH
func lookPath(file string, pathenv string) (string, error) {
	if strings.Contains(file, "/") {
		err := findExecutable(file)
		if err == nil {
			return file, nil
		}
		return "", err
	}
	if pathenv == "" {
		return "", errNotFound
	}
	for _, dir := range strings.Split(pathenv, ":") {
		if dir == "" {
			// Unix shell semantics: path element "" means "."
			dir = "."
		}
		path := dir + "/" + file
		if err := findExecutable(path); err == nil {
			return path, nil
		}
	}
	return "", errNotFound
}

// ErrNotFound is the error resulting if a path search failed to find an executable file.
var errNotFound = errors.New("executable file not found in $PATH")

func findExecutable(file string) error {
	d, err := os.Stat(file)
	if err != nil {
		return err
	}
	if m := d.Mode(); !m.IsDir() && m&0111 != 0 {
		return nil
	}
	return os.ErrPermission
}
