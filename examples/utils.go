package examples

import (
	"encoding/json"
	"fmt"
	"github.com/cattail/databricks-sdk-go/databricks"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var secrets = LoadSecrets()
var client = GetClient()
var ctx context.Context = nil

func GetClient() *databricks.APIClient {
	cfg := databricks.NewConfiguration()
	cfg.AddDefaultHeader("Authorization", "Bearer " + secrets.Token)
	cfg.BasePath = secrets.Domain + "/api/2.0"
	return databricks.NewAPIClient(cfg)
}

type Secrets struct {
	Domain      string `json:"domain"`
	Token       string `json:"token"`
}

func LoadSecrets() *Secrets {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadFile(filepath.Join(dir, "../examples/secrets.json"))
	if err != nil {
		panic(err)
	}

	var sc Secrets
	err = json.Unmarshal(content, &sc)
	if err != nil {
		panic(err)
	}

	return &sc
}

func JsonPrint(response interface{}) {
	result, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))
}
