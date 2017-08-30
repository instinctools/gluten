package client

import (
	"bitbucket.org/instinctools/gluten/shared/logging"
	"encoding/json"
	"io/ioutil"
)

func ReadJSONFile(pathToFile string) string {
	plan, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("File reading failed")
	}
	data := Project{}
	err = json.Unmarshal(plan, &data)
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("Unmarshal failed")
	}
	str, err := json.Marshal(data)
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("Marshal failed")
	}
	return string(str)
}

func AutoGenerateConfig(filename string) {
	//TODO - generateJSON should be a string, don't need to marshal it to string
	json := "{\"Name\": \"Project1\","
//				+ "\"Scenarios\": [{"
//					+"\"Name\": \"Scenario1\","
//						+"\"Cases\": [{"
//							+"\"Name\": \"Case1\","
//								+"\"Steps\": [{"
//									+"\"Name\": \"G1\","
//									+"\"Type\": \"GetRequestStep\","
//									+"\"Parameters\": {\"URL\": \"https://google.com\"},"
//									+"\"SubSteps\" : [{"
//										+"\"Name\": \"P1\","
//										+"\"Type\": \"GetRequestStep\","
//										+"\"Parameters\": {\"URL\": \"https://google.com\"}"
//									+"}]},"
//									+"{\"Name\": \"G2\","
//									+"\"Type\": \"GetRequestStep\","
//									+"\"Parameters\": {\"URL\": \"https://google.com\"}"
//								+"}]"
//							+"}]"
//						+"}]"
//					+"}"
	err := ioutil.WriteFile(filename, []byte(json), 0644)
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("File writing failed")
	}
}
