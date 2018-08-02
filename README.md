# go-indian-pincodes
Simple program to get District, Taluk and State name 

## Installing
```bash
go get github.com/krs92/go-indian-pincodes
```

## Getting Started

```go
  data, err := pincode.Lookup(504293)
  
  and Response will be 
  
{
"officeName":"Ada B.O",
"pincode":504293,
"taluk":"Asifabad",
"districtName":"Adilabad",
"stateName":"ANDHRA PRADESH"
}

```

## The pincode data used in the module has been provided by data.gov.in.



## TO-DO
* Get Pincode from district 
* Get picode from Statename
