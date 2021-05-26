package model

import (
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
)
type LogInfo struct {
	AppId int `db:"app_id"`
	AppName string `db:"app_name"`
	LogId int `db:"log_id"`
	LogPath string `db:"log_path"`
	CreateTime string `db:"create_time"`
	Status int `db:"status"`
	Topic string `db:"topic"`
}

func GetAllLogInfo()(logList []LogInfo, err error){
	err = Db.Select(&logList,
		"select a.app_id, b.app_name, a.create_time, a.topic, a.log_id, a.status, a.log_path from tbl_log_info a, tbl_app_info b where a.app_id=b.app_id")
	if err != nil {
		logs.Warn("Get all app info failed, err:%v", err)
		return
	}
	return
}

func CreateLog(info *LogInfo)(err error){
	conn, err := Db.Begin()
	if err != nil {
		logs.Warn("CreateLog failed, Db.Begin error:%v", err)
		return
	}

	defer func(){
		if err != nil {
			conn.Rollback()
			return
		}
		conn.Commit()
	}()
	var appId []int
	err = Db.Select(&appId, "select app_id from tbl_app_info where app_name = ?", info.AppName)
	if err != nil {
		logs.Warn("select app_id failed, Db.Exec error:%v", err)
	}
	info.AppId = appId[0]
	r , err := conn.Exec("insert into tbl_log_info(app_id, log_path, topic) values(?, ?, ?)",
		info.AppId, info.LogPath, info.Topic)

	if err != nil {
		logs.Warn("CreateLog failed, Db.Exec error:%v", err)
	}

	_, err = r.LastInsertId()
	if err != nil {
		logs.Warn("createLog failed, Db.LastInsertId error:%v", err)
		return
	}
	return
}