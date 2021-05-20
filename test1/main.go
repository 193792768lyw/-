package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	s := `{
            "cast1":{
                "from":[
                    "sql4"
                ],
                "name":"cast",
                "values":{
                    "cols":[
                        {
                            "name":"_eventname",
                            "newValue":"boolean"
                        },
                        {
                            "name":"_accountid",
                            "newValue":"boolean"
                        },
                        {
                            "name":"app_plat",
                            "newValue":"string"
                        }
                    ]
                }
            },
            "join5":{
                "from":[
                    "rename2",
                    "sql6"
                ],
                "name":"join",
                "values":{
                    "joinkey":[
                        "_eventname"
                    ],
                    "jointype":"inner",
                    "left_field_cols":[
                        "_eventname"
                    ]
                }
            },
            "rename2":{
                "from":[
                    "cast1"
                ],
                "name":"rename",
                "values":{
                    "recols":[
                        {
                            "name":"_eventname",
                            "newValue":"事件名",
                            "oldValue":"事件名"
                        },
                        {
                            "name":"_accountid",
                            "newValue":"账号id",
                            "oldValue":"账号id"
                        },
                        {
                            "name":"app_plat",
                            "newValue":"APP平台",
                            "oldValue":"APP平台"
                        }
                    ]
                }
            },
            "sql4":{
                "from":[

                ],
                "name":"sql",
                "values":{
                    "query":{
                        "cols":[
                            "_eventname",
                            "_accountid",
                            "_time"
                        ],
                        "filter_info":{
                            "filters":[

                            ],
                            "relative":1
                        },
                        "table_name":"dwd_register_log",
                        "type":"event"
                    }
                }
            },
            "sql6":{
                "from":[

                ],
                "name":"sql",
                "values":{
                    "query":{
                        "cols":[
                            "_eventname"
                        ],
                        "filter_info":{
                            "filters":[

                            ],
                            "relative":1
                        },
                        "table_name":"dwd_register_log",
                        "type":"event"
                    }
                }
            }
        }`

	js, err := simplejson.NewJson([]byte(s))
	if err != nil {
		fmt.Println(err)
	}
	daMap, err := js.Map()
	for _, v := range daMap {
		m1 := v.(map[string]interface{})
		if m1["name"] == "rename" || m1["name"] == "cast" {
			m2 := m1["values"].(map[string]interface{})
			res := map[string]interface{}{}
			for _, v2 := range m2 {
				m3 := v2.([]interface{})
				for _, v3 := range m3 {
					m4 := v3.(map[string]interface{})
					res[m4["name"].(string)] = m4["newValue"]
				}
			}
			m1["values"] = res
		}
	}
	ttt, _ := json.Marshal(daMap)
	fmt.Println(ttt)
}

var qq int

func main1() {
	var id cron.EntryID
	var err error
	c := cron.New()
	id, err = c.AddFunc("* * * * *", func() {
		fmt.Println("Every hour on the half hour", time.Now())
		c.Remove(id)
	})
	fmt.Println(id, err)
	//c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	//c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.Start()

	time.Sleep(999999 * time.Second)
}
