package main

import (
	//"main/models"
	//"fmt"
	"fmt"
	"encoding/json"
	"github.com/aliveyun/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/bson"
	//"context"
	//"go.mongodb.org/mongo-driver/mongo/options"
)
type pull struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	DeviceID  string    `bson:"device"`
	RTSP     string             `bson:"rtsp"`
	RTMP    string             `bson:"rtmp"`

}
var store *mongo.MgoStore
var defaultCollection *pull


func (c *pull) Name() string {
	return "pull"
}

func (c *pull) GetId() primitive.ObjectID {
	return c.Id
}

func (c *pull) SetId(id primitive.ObjectID) {
	c.Id = id
}


func main() {

	conf := &mongo.MgoConf{
		User:        "admin",
		Password:    "123456",
		DataSource:  []string{"10.55.16.241:27017"},
		DB:          "aliveyun",
		AuthDB:      "?connect=direct",
		ReplicaSet:  "",
		MaxPoolSize: 20,
	}
	dbCli, err :=NewDB(conf)
	if err != nil {
		panic(err)
	}
	b := []byte(`{
		"message": "getdevicelist",
		"pushdevice": "robot",
	"sn": "1",
		"body": {
	
			"robot": [{
					  "rtsp":"rtsp://admin:Aa123456@10.55.133.11:554/Streaming/Channels/1",
					"rtmp":"rtmp://10.55.23.36:1935/live/robot"
				},
				{
					  "rtsp":"rtsp://admin:abc12345@10.55.17.52:554/h264/ch1/main/av_stream",
					"rtmp":"rtmp://10.55.23.36:1935/live/robot2"
	
				}
			]
		}
	}`)
	dbCli.AddDevice("aliveyun",b)
	m,_:=dbCli.GetDevice("aliveyun")
	var msg MsgHead
	if err := json.Unmarshal(m, &msg); err != nil {
		//return errors.Wrapf(err, "Unmarshal %s", m)
		fmt.Println("4444",err)
	}
	fmt.Println("555",msg)
	/*dbCli, err := mongo.NewClient(conf)
	if err != nil {
		panic(err)
	}
	dbStore := mongo.NewStore(dbCli, conf.DB)
	store = dbStore*/
	/*def:= &pull{
		DeviceID:     "alive",
		RTSP:    "RTSP//",
		RTMP:   "rtmp//",
	}

	db,_:=mongo.NewDB(conf)
	err := db.Store().InsertOne(context.TODO(), def)
	if err != nil {
		//t.Fatal(err)
		fmt.Println("store.InsertOne",err)
	}

	err = db.Store().InsertOne(context.TODO(), def)
	if err != nil {
		//t.Fatal(err)
		fmt.Println("store.InsertOne",err)
	}


	col := &pull{}
	filter := bson.D{
		{"device", "alive"},
	}
	opt := options.FindOne().SetSort(bson.D{{"_id", -1}})
	finder := mongo.NewOneFinder(col).Where(filter).Options(opt)

	b, err := db.Store().FindOne(context.TODO(), finder)
	if err != nil {
		fmt.Println("store.FindOne",err)
	}
	fmt.Println("FindOne",b)
	fmt.Println("FindOne ",col)

	//
	col = &pull{}
	filter = bson.D{
		{"device", "alive"},
	}
	opt1 := options.Find().SetSort(bson.D{{"_id", -1}})
	records := make([]*pull, 0)
	finder1 := mongo.NewFinder(col).Where(filter).Options(opt1).Records(&records)

	err = db.Store().FindMany(context.TODO(), finder1)
	if err != nil {
		fmt.Println("store.FindMany",err)
	}
	for _, item := range records {

		fmt.Println("FindMany ",item)
	}*/
	//
}