package utils

import (
	"encoding/json"
	"fmt"
	"testing"

	core_utils "github.com/fluffy-bunny/grpcdotnetgo/pkg/utils"
)

func TestRedact(t *testing.T) {
	type Sensitive struct {
		Name     string `json:"name"`
		Password string `json:"password" redact:"true"`
	}
	obj := &Sensitive{
		Name:     "John",
		Password: "secret",
	}
	fmt.Println(core_utils.PrettyJSON(obj))
	json, _ := json.Marshal(obj)
	fmt.Println(string(json))

	dst := &Sensitive{}
	PrettyPrintRedacted(obj, dst)

}
