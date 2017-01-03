package amslib

import (
	"log"

	//amslib "bitbucket.org/usautoparts/svc-ams-lib"
	"gopkg.in/mgo.v2/bson"
)

/*
* argument: SellerID, MethodName
* return int (summation of all count)
 */
func Counter(seller string, apiName string) (bool, error) {

	//assign AMSlib COnfig/ Environment variables
	cnf := NewConfig()
	if !cnf.Validate() {
		log.Println("[ERROR] Environment variables not set")
	}
	//connect to MONGODB

	mngSession, err := CreateSession(cnf.MongoConnectionURL)
	//check for error
	if err != nil {
		log.Println("[ERROR]", err.Error())
		return false, err
	}
	//close connection when done
	defer mngSession.Close()

	//cnf.MongoDBListing is competitor_items database
	//api_call_counter collection
	APICallCounter := mngSession.DB(cnf.MongoDBListing).C("api_call_counter")
	log.Println("Connected to mongo")

	_, err = APICallCounter.Upsert(bson.M{"sid": seller, "api": apiName}, bson.M{"$inc": bson.M{"ctr": 1}})

	if err != nil {
		log.Println("[ERROR]", err.Error())
		return false, err
	}

	return true, nil
}
