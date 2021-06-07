package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	e, err := casbin.NewEnforcer("D:\\goworkstation\\Study\\Casbin\\keymatch_model.conf", "D:\\goworkstation\\Study\\Casbin\\keymatch_policy.csv")

	fmt.Printf("RBAC TENANTS test start\n") // output for debug

	// superAdmin
	if falg, _ := e.Enforce("superAdmin", "gy", "project", "read"); falg {
		log.Println("superAdmin can read project in gy")
	} else {
		log.Fatal("ERROR: superAdmin can not read project in gy")
	}

	if falg, _ := e.Enforce("superAdmin", "gy", "project", "write"); falg {
		log.Println("superAdmin can write project in gy")
	} else {
		log.Fatal("ERROR: superAdmin can not write project in gy")
	}

	if falg, _ := e.Enforce("superAdmin", "jn", "project", "read"); falg {
		log.Println("superAdmin can read project in jn")
	} else {
		log.Fatal("ERROR: superAdmin can not read project in jn")
	}

	if falg, _ := e.Enforce("superAdmin", "jn", "project", "write"); falg {
		log.Println("superAdmin can write project in jn")
	} else {
		log.Fatal("ERROR: superAdmin can not write project in jn")
	}

	// admin
	if falg, _ := e.Enforce("quyuan", "gy", "project", "read"); falg {
		log.Println("quyuan can read project in gy")
	} else {
		log.Fatal("ERROR: quyuan can not read project in gy")
	}

	if falg, _ := e.Enforce("quyuan", "gy", "project", "write"); falg {
		log.Println("quyuan can write project in gy")
	} else {
		log.Fatal("ERROR: quyuan can not write project in gy")
	}

	if falg, _ := e.Enforce("quyuan", "jn", "project", "read"); falg {
		log.Fatal("ERROR: quyuan can read project in jn")
	} else {
		log.Println("quyuan can not read project in jn")
	}

	if falg, _ := e.Enforce("quyuan", "jn", "project", "write"); falg {
		log.Fatal("ERROR: quyuan can write project in jn")
	} else {
		log.Println("quyuan can not write project in jn")
	}

	if falg, _ := e.Enforce("quyuan", "gy", "asse", "read"); falg {
		log.Fatal("ERROR: quyuan can read asse in gy")
	} else {
		log.Println("quyuan can not read asse in gy")
	}

	if falg, _ := e.Enforce("quyuan", "gy", "asse", "write"); falg {
		log.Fatal("ERROR: quyuan can write asse in gy")
	} else {
		log.Println("quyuan can not write asse in gy")
	}

	if falg, _ := e.Enforce("quyuan", "jn", "asse", "read"); falg {
		log.Println("quyuan can read asse in jn")
	} else {
		log.Fatal("ERROR: quyuan can not read asse in jn")
	}

	if falg, _ := e.Enforce("quyuan", "jn", "asse", "write"); falg {
		log.Println("quyuan can write asse in jn")
	} else {
		log.Fatal("ERROR: quyuan can not write asse in jn")
	}

	// wenyin
	if falg, _ := e.Enforce("wenyin", "gy", "asse", "write"); falg {
		log.Println("wenyin can write asse in gy")
	} else {
		log.Fatal("ERROR: wenyin can not write asse in gy")
	}

	if falg, _ := e.Enforce("wenyin", "jn", "asse", "write"); falg {
		log.Fatal("ERROR: wenyin can write asse in jn")
	} else {
		log.Println("wenyin can not write asse in jn")
	}

	// shangshang
	if falg, _ := e.Enforce("shangshang", "jn", "project", "write"); falg {
		log.Println("shangshang can write project in jn")
	} else {
		log.Fatal("ERROR: shangshang can not write project in jn")
	}

	if falg, _ := e.Enforce("shangshang", "gy", "project", "write"); falg {
		log.Fatal("ERROR: shangshang can write project in gy")
	} else {
		log.Println("shangshang can not write project in gy")
	}
	fmt.Println(err)
}

func main1() {

	e, err := casbin.NewEnforcer("D:\\goworkstation\\Study\\Casbin\\keymatch_model.conf", "D:\\goworkstation\\Study\\Casbin\\keymatch_policy.csv")

	fmt.Printf("RBAC test start\n") // output for debug

	// superAdmin
	if falg, _ := e.Enforce("superAdmin", "project", "read"); falg {
		log.Println("superAdmin can read project")
	} else {
		log.Fatal("ERROR: superAdmin can not read project")
	}

	if falg, _ := e.Enforce("superAdmin", "project", "write"); falg {
		log.Println("superAdmin can write project")
	} else {
		log.Fatal("ERROR: superAdmin can not write project")
	}

	// admin
	if falg, _ := e.Enforce("quyuan", "project", "read"); falg {
		log.Println("quyuan can read project")
	} else {
		log.Fatal("ERROR: quyuan can not read project")
	}

	if falg, _ := e.Enforce("quyuan", "project", "write"); falg {
		log.Println("quyuan can write project")
	} else {
		log.Fatal("ERROR: quyuan can not write project")
	}

	if falg, _ := e.Enforce("quyuan", "asse", "read"); falg {
		log.Println("quyuan can read asse")
	} else {
		log.Fatal("ERROR: quyuan can not read asse")
	}

	if falg, _ := e.Enforce("quyuan", "asse", "write"); falg {
		log.Println("quyuan can write asse")
	} else {
		log.Fatal("ERROR: quyuan can not write asse")
	}

	// zhuangjia
	if falg, _ := e.Enforce("wenyin", "project", "read"); falg {
		log.Fatal("ERROR: wenyin can read project")
	} else {
		log.Println("wenyin can not read project")
	}

	if falg, _ := e.Enforce("wenyin", "project", "write"); falg {
		log.Println("wenyin can write project")
	} else {
		log.Fatal("ERROR: wenyin can not write project")
	}

	if falg, _ := e.Enforce("wenyin", "asse", "read"); falg {
		log.Fatal("ERROR: wenyin can read asse")
	} else {
		log.Println("wenyin can not read asse")
	}

	if falg, _ := e.Enforce("wenyin", "asse", "write"); falg {
		log.Println("wenyin can write asse")
	} else {
		log.Fatal("ERROR: wenyin can not write asse")
	}

	// shangshang
	if falg, _ := e.Enforce("shangshang", "project", "read"); falg {
		log.Println("shangshang can read project")
	} else {
		log.Fatal("ERROR: shangshang can not read project")
	}

	if falg, _ := e.Enforce("shangshang", "project", "write"); falg {
		log.Fatal("ERROR: shangshang can write project")
	} else {
		log.Println("shangshang can not write project")
	}

	if falg, _ := e.Enforce("shangshang", "asse", "read"); falg {
		log.Println("shangshang can read asse")
	} else {
		log.Fatal("ERROR: shangshang can not read asse")
	}

	if falg, _ := e.Enforce("shangshang", "asse", "write"); falg {
		log.Fatal("ERROR: shangshang can write asse")
	} else {
		log.Println("shangshang can not write asse")
	}

	fmt.Println(err)
	//	// 使用MySQL数据库初始化一个Xorm适配器
	//	engine, err := xorm.NewEngine("mysql", "root:193792@tcp(127.0.0.1:3306)/test")
	//	if err != nil {
	//		os.Exit(1)
	//	}
	//	a, err := xormadapter.NewAdapterByEngine(engine)
	//	if err != nil {
	//		log.Fatalf("error: adapter: %s", err)
	//	}
	//
	//	m, err := model.NewModelFromString(`
	//[request_definition]
	//r = sub, obj, act
	//
	//[policy_definition]
	//p = sub, obj, act
	//
	//[policy_effect]
	//e = some(where (p.eft == allow))
	//
	//[matchers]
	//m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
	//`)
	//	if err != nil {
	//		log.Fatalf("error: model: %s", err)
	//	}
	//
	//	e, err := casbin.NewEnforcer(m, a)
	//	if err != nil {
	//		log.Fatalf("error: enforcer: %s", err)
	//	}
	//	fmt.Println(e)
}
