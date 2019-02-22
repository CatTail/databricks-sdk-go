package examples

import (
	"encoding/json"
	sw "github.com/cattail/databricks-sdk-go/client"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var secrets = LoadSecrets()

func GetClient() *sw.APIClient {
	cfg := sw.NewConfiguration()
	cfg.BasePath = secrets.Domain + "/api/2.0"
	return sw.NewAPIClient(cfg)
}

func GetAuth() context.Context {
	return context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
		Key: secrets.Token,
		Prefix: "Bearer",
	})
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

	content, err := ioutil.ReadFile(filepath.Join(dir, "../secrets.json"))
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
