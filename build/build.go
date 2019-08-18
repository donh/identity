package build

import (
	"encoding/json"
	"io/ioutil"
)

// Build generates two files: KeyManager.abi and KeyManager.bin
func Build() {
	dat, err := ioutil.ReadFile("./contracts/KeyManager.json")
	if err != nil {
		panic(err.Error())
	}
	output := make(map[string]interface{})
	_ = json.Unmarshal(dat, &output)
	abi, _ := json.Marshal(output["abi"])
	bytecode, _ := json.Marshal(output["bytecode"])

	_ = ioutil.WriteFile("../contracts/KeyManager.abi", abi, 0644)
	_ = ioutil.WriteFile("../contracts/KeyManager.bin", bytecode, 0644)
}
