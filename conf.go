package conf

import (
	"flag"
	"fmt"
	"log"

	"github.com/go-ini/ini"
)

var cfg *ini.File
var env string

// Setup init dotenv package
func Setup(config map[string]interface{}) error {
	var err error
	// load command env params
	flag.StringVar(&env, "env", env, "set project env, if not seted, default production, can use: development, test, production")
	flag.Parse()

	if env == "" {
		env = "prod"
		log.Println("The project environment variable was not obtained，use 'prod' environment file")
	} else {
		log.Printf("use %s environment file", env)
	}

	// load config directory files
	file := fmt.Sprintf("%s/%s.ini", config["directory"], env)
	cfg, err = ini.Load(file)

	if err != nil {
		log.Fatalf("Fail to parse '%s/%s.ini'： %v", config["directory"], env, err)
		return err
	}

	err = cfg.Section("").MapTo(config["data"])
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
