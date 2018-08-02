package pincode

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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
	jsonFile, err :=  http.Get("https://github.com/krs92/go-indian-pincodes/blob/master/pincodes.json?raw=true")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Body.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile.Body)

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
