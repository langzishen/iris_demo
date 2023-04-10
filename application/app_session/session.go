package app_session

import (
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/sessions/sessiondb/boltdb"
	"os"
)

//定义home通用的session
var Sess_Home = sessions.New(sessions.Config{Cookie: "sessionHomeKey"})
var Sessdb_Home, _ = boltdb.New("./sessions_db.db", os.FileMode(0750))
