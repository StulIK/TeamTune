package main

import (
	"net/http"
	"../../../core"
	"../../projects"
	"../../../core/structures"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type ServiceDatabase struct {
	Dao *core.MongoDatabase
}
var Database = ServiceDatabase{&core.Dao}

//Check if correct username and password
func (r ServiceDatabase) getList(userID bson.ObjectId) bool {
	r.Dao.C("projects")

	var results []projects.Project
	fmt.Println(bson.M{"users": bson.M{"_id":userID}})
	err := Database.Dao.Collection.Find(bson.M{"users": bson.M{"$elemMatch":bson.M{"_id":userID}}}).Select(bson.M{"_id": 1, "name":1}).All(&results)
	if err != nil {
		core.SetResponse("database_error")
		return false
	}

	core.SetResponse("list_retrieved")
	core.SetData(results)
	return true
}


//Connects to database and listens to port
func main() {
	Database.Dao.Connect(core.Config.DatabaseHost + ":" + core.Config.DatabasePort, core.Config.DatabaseName)
	http.HandleFunc("/", do)
	http.ListenAndServe(core.Config.Host + ":1337", nil)
}

func do(w http.ResponseWriter, r *http.Request) {
	core.CORS(w)

	//Parses request data to
	var data structures.Session
	if core.DecodeRequest(&data, r){

		success,UserID := Database.Dao.CheckSession(data)
		if success {
			Database.getList(UserID) //Adds project to database
		}
	}

	//Prints R
	core.PrintReponse(w)
}
