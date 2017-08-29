package deserializers

import (
	pb "bitbucket.org/instinctools/gluten/shared/rpc/cli"
	"encoding/json"
	"fmt"
)

type Mapa struct {
	Name string
	Params map[string]string
}

func DeserializeJsonToProto(jsonProject string) *pb.Project {
	deserializedProject := pb.Project{}
	fmt.Printf("JSON : %s\n", jsonProject) 
	err := json.Unmarshal([]byte(jsonProject), &deserializedProject)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Printf("Project : %s\n", deserializedProject) 

	return &deserializedProject
}
