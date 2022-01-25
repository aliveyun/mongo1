package main

import (
	//"main/models"
	"errors"
	"github.com/aliveyun/mongo"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gogf/gf/os/glog"

)

type Device struct {

	c *mongo.Client

}

type MsgHead struct {
	Message    string   `json:"message"`
	Pushdevice string   `json:"pushdevice"`
	Sn         string   `json:"sn"`
	Body       *MsgBody `json:"body"`
}

type MsgBody struct {
	Robot []*RobotInfo
}

type RobotInfo struct {
	Rtsp string `json:"rtsp"`
	Rtmp string `json:"rtmp"`
}

const (
	HEAD_GETDEVICELIST = "getdevicelist"
	HEAD_ADDDEVICE     = "adddevice"
	HEAD_DELDEVICE     = "deldevice"
	HEAD_HEARTBEAT     = "heartbeat"
	HEAD_DISCONNECT     = "disconnect"
)


func NewDB(cfg *mongo.MgoConf) (*Device, error) {

	c, err := mongo.NewDB(cfg)
	if err != nil {
		return nil, err
	}
	return &Device{
		c:c,
	}, nil
}
func (c *Device) AddDevice(message []byte) error {
	var msg MsgHead
	if err := json.Unmarshal(message, &msg); err != nil {
		glog.Warningf("Unmarshal  Signal  err:%v\n", err)
		return err
	}
	// glog.Debugf("AddDevice:  %s",  message)	
	if msg.Body!=nil {	
		cols := make([]*pull, 0)	
		for _, v := range msg.Body.Robot {
			// glog.Debugf("recv: v=%v",  v)	
			// glog.Debugf("recv: %s %s",  v.Rtsp,v.Rtmp)	
			item:= &pull{
				DeviceID:     msg.Pushdevice,
				RTSP:   v.Rtsp,
				RTMP:   v.Rtmp,
			}
			cols = append(cols, item)
			id,_:=c.FindOneDevice(v.Rtsp ,v.Rtmp,msg.Pushdevice)
			glog.Debugf("FindOneDevice=%s",id)
			c.DelOneDevice(id)
		}
		docs := make([]interface{}, 0)
		for i := 0; i < len(cols); i++ {
			docs = append(docs, cols[i])
		}

		err := c.c.Store().InsertMany(context.TODO(), docs)
		if err != nil {
			glog.Debugf("InsertMany fail %s err=%s",docs,err)
		}

		return err
	}
  return errors.New("msg.Body is invalid")
}
func (c *Device) GetDevice(pushdevice string) ( []byte ,error) {


	col := &pull{}
	filter := bson.D{
		{"device", pushdevice},
	}
	opt := options.Find().SetSort(bson.D{{"_id", -1}})
	records := make([]*pull, 0)
	finder := mongo.NewFinder(col).Where(filter).Options(opt).Records(&records)

	err :=c.c.Store().FindMany(context.TODO(), finder)
	if err != nil {
		glog.Debugf("store.FindMany fail err=%s",err)
	}
	var msgBody MsgBody
    var	robot []*RobotInfo

	for _, item := range records {

		info:= &RobotInfo{
			Rtsp:   item.RTSP,
			Rtmp:  item.RTMP,
		}
		robot = append(robot,info)
	}
	msgBody.Robot = robot
	head := MsgHead{
		Message:    "getdevicelist",
		Pushdevice: pushdevice,
		Sn:         "1",
		Body: &msgBody,
	}
	m, err := json.Marshal(head)
	//glog.Debugf("store.FindMany fail err=%s",m)
	return m ,err
}



func (c *Device) FindOneDevice(rtsp,rtmp ,pushdevice string) ( string ,error) {
	col := &pull{}
	filter := bson.D{
		{"device", pushdevice},
		{"rtsp", rtsp},
		{"rtmp", rtmp},
	}
	opt := options.FindOne().SetSort(bson.D{{"_id", -1}})
	finder := mongo.NewOneFinder(col).Where(filter).Options(opt)
	_, err := c.c.Store().FindOne(context.TODO(), finder)
	return  col.GetId().Hex(), err
}

func (c *Device) DelOneDevice(ID string)  error {


	id, _ := primitive.ObjectIDFromHex(ID)
	col := &pull{
		Id: id,
	}
	_, err := c.c.Store().DeleteOne(context.TODO(), col)
	if err != nil {
		glog.Debugf("DelOneDevice id=%s fail=%s",ID,err)
	}
	glog.Debugf("DelOneDevice=%s",ID)
	//t.Log("cnt", cnt)
	return err
}