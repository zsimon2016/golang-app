package main

import (
	"flag"
	"fmt"
	"log"
	car "service/simon/api"
	"time"

	// "gitee.com/simon_git_code/godb/taos"
	// util "gitee.com/simon_git_code/golang-utils"

	"github.com/rcrowley/go-metrics"
	"github.com/rpcxio/rpcx-etcd/serverplugin"
	"github.com/smallnest/rpcx/server"

	// metrics "github.com/rcrowley/go-metrics"
	// "time"
	_ "log"
	// "gitee.com/simon_git_code/godb/mysql"
	// "gitee.com/simon_git_code/godb/taos"
	// "gitee.com/simon_git_code/smallnest/rpcx/server"
	// "simon/mongo"
	// "context"
	// util1 "car/herx/system/util"
	// "net/http"
	// _ "net/http/pprof"
)

var (
	addr     = flag.String("addr", "0.0.0.0:8975", "server address")
	papi     = flag.String("papi", "0.0.0.0:9090", "papi address")
	etcdAddr = flag.String("etcdAddr", "localhost:2379", "nacos address")
	basePath = flag.String("basePath", "/services", "basePath")
)

// var logDB *taos.Db

func main() {
	flag.Parse()

	s := server.NewServer()
	// port := uint64(0)
	// nacosGroup := "MALL_GROUP"
	addrss := ""
	fmt.Println(*addr)
	if *addr != "0.0.0.0:8975" {
		addrss = *addr
	}
	addRegistryPlugin(s)
	// nacosUtil, err := util.NewNacosUtil(nacosUser, nacosPassword, nacosAddr, nacosPort, nacosDsGroup, nacosGroup, &addrss, &port, nacosDataId)
	// if err != nil {
	// 	panic("注册中心连接失败")
	// }

	// nacosUtil.AddRegistryPlugin(s)
	// confMap,b:= util.LoadConfig("./config.json")
	// if b == false {
	// 	fmt.Println("配置文件载入出错")
	// }
	// confMap := make(map[string]interface{})
	// nacosUtil.GetConfig(&confMap)
	// fmt.Println("confMap", confMap)
	// cnf := &confMap
	// dbconf := confMap["tbox_db"].(map[string]interface{})

	//注册时链接数据库
	// _ = util.InitDB(dbconf)

	// Db := new(mysql.Db)
	// Db.DbConn = new(mysql.DbConnection)
	// Db.DbConn.Connt(&dbconf)

	// //redis
	// rConf := confMap["redis"].(map[string]interface{})
	// rhost := rConf["host"].(string)
	// rpwd := rConf["password"].(string)
	// pool := util.GetRedisPool(&rhost, &rpwd)
	// redis := util.Redis{pool}
	// util1.CreateRDS(confMap)
	// tdbconf := confMap["td_db"].(map[string]interface{})
	// _ = taos.InitDB(tdbconf)

	// tdbconf["database"] = "tbox_log"
	// logDB = new(taos.Db)
	// logDB.DbConn = new(taos.DbConnection)
	// logDB.DbConn.Connt(&tdbconf)
	// es.Builder(conf)
	car := car.Car{}

	s.RegisterName("Bi", &car, "")
	// go car.SyncMil()
	serr := s.Serve("tcp", addrss)
	if serr != nil {
		panic(serr)
	}

}
func addRegistryPlugin(s *server.Server) {
	r := &serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		EtcdServers:    []string{*etcdAddr},
		BasePath:       *basePath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}
