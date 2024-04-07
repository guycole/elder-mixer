package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const banner = "elder-mixer 0.0"

func main() {
	flag.Parse()

	rand.New(rand.NewSource(time.Now().UnixNano()))

	log.Println(banner)

	app := AppType{SugarLog: zapSetup(false)}
	app.SugarLog.Info(banner)

	envVars := [...]string{"FEATURE_FLAGS", "GRPC_PORT"}

	for index, element := range envVars {
		temp, err := os.LookupEnv(element)
		if err {
			app.SugarLog.Infof("%d:%s:%s", index, element, temp)
		} else {
			app.SugarLog.Fatal("missing:", element)
		}

		switch element {
		case "FEATURE_FLAGS":
			temp, err := strconv.Atoi(temp)
			if err == nil {
				app.FeatureFlags = uint32(temp)
			} else {
				app.SugarLog.Fatal("bad featureFlags")
			}
		case "GRPC_PORT":
			temp, err := strconv.Atoi(temp)
			if err == nil {
				app.GrpcPort = temp
			} else {
				app.SugarLog.Fatal("bad grpcPort")
			}
		default:
			app.SugarLog.Fatal("unknown environment var:", element)
		}
	}

	app.initialize()
	app.run()
}
