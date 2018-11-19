package flagfile

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var configFile = defaultConfig()

// Usage prints the usage information for the application including all flags
// and their values after parsing the configuration file
func Usage() {
	Parse()
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Fprintf(os.Stderr, "  -%s=%s: %s\n", f.Name, f.Value.String(), f.Usage)
	})
}

func checkFlagfile(dir string, name string) string {
	path := filepath.Join(dir, name)
	_, err := os.Open(path)
	if err == nil {
		return path
	}
	return ""
}

func defaultConfig() string {
	name := "flagfile"
	path := checkFlagfile("./", name)
	if path != "" {
		return path
	}

	path = checkFlagfile("./conf/", name)
	if path != "" {
		return path
	}

	path = checkFlagfile("../conf/", name)
	if path != "" {
		return path
	}

	return ""
}

func contains(list []*flag.Flag, f *flag.Flag) bool {
	for _, i := range list {
		if i == f {
			return true
		}
	}
	return false
}

func readConfig(filename string) map[string]string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read config file %s: %s", configFile, err)
	}
	lines := strings.Split(string(bytes), "\n")
	result := make(map[string]string, len(lines))
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || trimmed[0] == '#' {
			continue
		}
		parts := strings.Split(line, "=")
		if len(parts) < 2 {
			log.Fatalf("Invalid config line: %s", line)
		}
		result[strings.TrimSpace(parts[0])] = strings.TrimSpace(strings.Join(parts[1:], "="))
	}
	return result
}

// Parse parses the default configuration file and populates the global flags
// based on the contents of the file.
func Parse() {
	ParseSet(flag.CommandLine)
}

// ParseSet parses the default configuraiton file and populates the flags in
// the flag.FlagSet based on the contents of the file.
func ParseSet(flags *flag.FlagSet) {
	ParseFile(flags, configFile)
}

// ParseFile parses the specified configuration file and populates the flags
// in the flag.FlagSet based on the contents of the file.
func ParseFile(flags *flag.FlagSet, filename string) {
	if filename == "" {
		return
	}

	var (
		explicit []*flag.Flag
		all      []*flag.Flag
	)

	config := readConfig(filename)

	flags.Visit(func(f *flag.Flag) {
		explicit = append(explicit, f)
	})

	flags.VisitAll(func(f *flag.Flag) {
		all = append(all, f)
		if !contains(explicit, f) {
			// env
			val, exist := os.LookupEnv("nl_" + f.Name)
			if !exist {
				// flagfile
				val = config[f.Name]
			}
			if val != "" {
				err := f.Value.Set(val)
				if err != nil {
					log.Fatalf("Failed to set flag %s with value %s", f.Name, val)
				}
			}
		}
	})
Outer:
	for name, val := range config {
		for _, f := range all {
			if f.Name == name {
				continue Outer
			}
		}
		log.Fatalf("Unknown flag %s=%s in config file.", name, val)
	}
}
