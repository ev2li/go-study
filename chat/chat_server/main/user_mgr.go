package main

import (
	"go-study/chat/chat_server/model"
)

var (
	mgr *model.UserMgr
)

func initUserMgr() {
	mgr = model.NewUserMgr(pool)
}
