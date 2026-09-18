package parser

import (
	"encoding/json"
	"fmt"
	"k8s-learning/factory"
	"k8s-learning/k8sinterface"
	"os"
	"strings"

	"k8s.io/apimachinery/pkg/util/yaml"
)

func Parser(filename string, registry *factory.Registry) []k8sinterface.Resource {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	decoder := yaml.NewYAMLOrJSONDecoder(strings.NewReader(string(data)), 1024)
	var resources []k8sinterface.Resource
	for {
		var raw map[string]interface{}
		err := decoder.Decode(&raw)
		if err != nil {
			break
		}
		kind, ok := raw["kind"].(string)
		if !ok {
			continue
		}
		obj := registry.Create(kind)
		if obj == nil {
			fmt.Printf("不支持的资源类型:%s\n", kind)
			continue
		}
		jsonData, err := json.Marshal(raw)
		if err != nil {
			continue
		}

		err = json.Unmarshal(jsonData, obj)
		if err != nil {
			continue
		}
		resources = append(resources, obj)
	}
	return resources
}
