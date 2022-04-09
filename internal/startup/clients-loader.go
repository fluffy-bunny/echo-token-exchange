package startup

import (
	"echo-starter/internal"
	"echo-starter/internal/models"
	"echo-starter/internal/utils"
	"errors"
	"fmt"
	"path"

	"github.com/spf13/viper"
	"github.com/xeipuuv/gojsonschema"
)

func (s *Startup) loadTestClients() (err error) {
	schemaPath := utils.ToCanonical(path.Join(internal.RootFolder, "static/clients/clients.schema.json"))
	documentPath := utils.ToCanonical(path.Join(internal.RootFolder, "static/clients/clients.json"))

	schemaLoader := gojsonschema.NewReferenceLoader(schemaPath)
	documentLoader := gojsonschema.NewReferenceLoader(documentPath)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return
	}
	if !result.Valid() {
		sError := ""
		for _, desc := range result.Errors() {
			sError += fmt.Sprintf("- %s\n", desc)
		}
		fmt.Println(sError)
		err = errors.New(fmt.Sprintf("document:[%s] did not pass schema:[%s] validation errors:[%s]",
			documentPath, schemaPath, sError))
		return
	}
	clientsConfig := viper.New()
	clientsConfig.SetConfigFile("static/clients/clients.json")
	err = clientsConfig.ReadInConfig()
	if err != nil {
		return
	}
	var clients []models.Client
	err = clientsConfig.UnmarshalKey("clients", &clients)
	if err != nil {
		return
	}

	// need to set the secret has as the original has them in the open
	for ic := range clients {
		for idx := range clients[ic].ClientSecrets {
			hash, _ := utils.GeneratePasswordHash(clients[ic].ClientSecrets[idx].Value)
			clients[ic].ClientSecrets[idx].Value = hash
		}
	}
	s.clients = clients
	return
}
