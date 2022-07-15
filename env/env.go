package env

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/alexflint/go-arg"
	"github.com/joho/godotenv"
)

func init() {
	changeDir()
	err := godotenv.Load(getEnvFileName())
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

func getEnvFileName() string {
	if envFile := os.Getenv("_env_file"); envFile != "" {
		return envFile
	}
	return ".env"
}

func Getenv(key string, fallback ...string) string {
	value := os.Getenv(key)
	if len(value) == 0 && len(fallback) > 0 {
		return fallback[0]
	}
	return value
}
func GetenvMust(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		panic(key + " should be setted")
	}
	return value
}

func GetenvInt(key string, fallback ...int) int {
	value := os.Getenv(key)
	if len(value) == 0 && len(fallback) > 0 {
		return fallback[0]
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return v
}

func GetenvFloat64(key string, fallback ...float64) float64 {
	value := os.Getenv(key)
	if len(value) == 0 && len(fallback) > 0 {
		return fallback[0]
	}
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}
	return v
}

func GetenvBool(key string, fallback bool) bool {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	v, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}
	return v
}
func FileAsBytes(p string) []byte {
	bs, err := os.ReadFile(p)
	if err != nil {
		panic(err)
	}
	return bs
}

func Exists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

func Setenv(key, value string) error {
	return os.Setenv(key, value)
}

func Getwd() string {
	wd, _ := os.Getwd()
	return wd
}

func changeDir() {
	wd, _ := os.Getwd()
	envFile := getEnvFileName()
	for filepath.Dir(wd) != wd {
		if Exists(filepath.Join(wd, envFile)) {
			break
		}
		wd = filepath.Dir(wd)
	}
	if !Exists(filepath.Join(wd, envFile)) {
		panic("Can not find " + envFile + " file")
	}
	fmt.Println("Find " + envFile + " file in " + wd)
	os.Chdir(wd)

}

//ParseArgs see https://github.com/alexflint/go-arg
func ParseArgs(argsRef interface{}) {
	arg.MustParse(argsRef)
}
