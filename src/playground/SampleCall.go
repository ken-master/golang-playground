package main

import (
	"encoding/json"
	"log"

	ebayapi "svc-ebay-api"
	//"io/ioutil"
	//"getapiaccess" //local structs
	//"getitem"
	//"github.com/franela/goreq"
)

func main() {
	c := ebayapi.NewConfig()

	id := []string{"110183992526", "110183992527"}
	//log.Printf("%v", id)
	//110183992526
	handle := ebayapi.GetMultipleItemsRequest{
		Xmlns: "urn:ebay:apis:eBLBaseComponents",
		//EBayAuthToken: c.Token,
		//ErrorLanguage: "en_US",
		//WarningLevel:  "High",
		//Version:       "967",
		ItemID: id,
	}

	//log.Printf("%v", handle)
	response, num, err := handle.Fetch(*c)
	if err != nil {
		panic(err)
	}

	log.Printf("%v", num)

	//var res interface{}

	res, err := json.Marshal(response)

	log.Printf("%v\n", string(res))
}
