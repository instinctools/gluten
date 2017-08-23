package steps

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (ce *TestCase) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	var rawMessageForCaseName *json.RawMessage
	err = json.Unmarshal(*objMap["Name"], &rawMessageForCaseName)
	if err != nil {
		return err
	}
	var str string
	err = json.Unmarshal(*rawMessageForCaseName, &str)
	if err != nil {
		return err
	}
	ce.Name = str
	var rawMessagesForSteps []*json.RawMessage
	err = json.Unmarshal(*objMap["Steps"], &rawMessagesForSteps)
	if err != nil {
		return err
	}
	ce.Steps = make([]Step, len(rawMessagesForSteps))
	var m map[string]interface{}
	for index, rawMessage := range rawMessagesForSteps {
		fmt.Println(string(*rawMessage))

		err = json.Unmarshal(*rawMessage, &m)
		if err != nil {
			return err
		}
		// TODO replace If with Switch/Case
		if m["type"] == "GetRequest" {
			var p GetRequestStep
			err := json.Unmarshal(*rawMessage, &p)
			if err != nil {
				return err
			}
			ce.Steps[index] = &p
		} else if m["type"] == "Composite" {
			var a CompositeStep
			err := json.Unmarshal(*rawMessage, &a)
			if err != nil {
				return err
			}
			ce.Steps[index] = &a
		} else {
			return errors.New("Unsupported type found!")
		}
	}
	return nil
}
