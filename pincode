package pincode

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)


type Pincode struct {
	OfficeName   string `json:"officeName"`
	Pincode      int64  `json:"pincode"`
	Taluk        string `json:"taluk"`
	DistrictName string `json:"districtName"`
	StateName    string `json:"stateName"`
}

type Pincodes []Pincode

func Lookup(data int64) (interface{}, error) {
	// Open our jsonFile
	jsonFile, err := os.Open("./pincodes.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var pincode Pincodes
	json.Unmarshal(byteValue, &pincode)
	//fmt.Print(pincode, data)
	for i := 0; i < len(pincode); i++ {
		if pincode[i].Pincode == data {
			return pincode[i], nil
		}
		return nil, errors.New("pincode is not available")
	}
	return nil, nil
}
