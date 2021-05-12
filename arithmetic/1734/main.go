package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
)

func main() {
	s := `{
    "project_id": 39,
    "task_name": "test_100",
    "description": "注册当天截止今天最高等级角色统计",
    "node": {
        "duplicate_83075415": {
            "model": "duplicate",
            "model_name": "去重",
            "type": "processor",
            "type_name": "算法",
            "values": {
			"cols":{
					"key" : "cnt1",
                   "value":"string"
            	},
                "duplicate_cols": [
                    "vopenid",
                    "vgamesvrid"
                ]
            }
        },
        "duplicate_9c91e7cd": {
            "model": "duplicate",
            "model_name": "去重",
            "type": "processor",
            "type_name": "算法",
            "values": {
                "duplicate_cols": [
                    "vopenid",
                    "level",
                    "vgamesvrid"
                ]
            }
        },
        "max_2bfe204f": {
            "model": "max",
            "model_name": "取最大值",
            "type": "base_calculation",
            "type_name": "基础计算",
            "values": {
                "group": 1,
                "group_cols": [
                    "vopenid",
                    "vgamesvrid"
                ],
                "colname": "level",
                "alias_field": "level"
            }
        },
        "merge_78dd692d": {
            "model": "merge",
            "model_name": "合并数据",
            "type": "processor",
            "type_name": "算法",
            "values": {
                "how": "left",
                "on": [
                    "vopenid",
                    "vgamesvrid"
                ],
                "is_fillna": 0,
                "sort": [
                    "duplicate_83075415",
                    "merge_bd7dbae8"
                ]
            }
        },
        "merge_2af6b446": {
            "model": "merge",
            "model_name": "合并数据",
            "type": "processor",
            "type_name": "算法",
            "values": {
                "how": "inner",
                "on": [
                    "vopenid",
                    "level",
                    "vgamesvrid"
                ],
                "is_fillna": 0,
                "sort": []
            }
        },
        "merge_bd7dbae8": {
            "model": "merge",
            "model_name": "合并数据",
            "type": "processor",
            "type_name": "算法",
            "values": {
                "how": "inner",
                "on": [
                    "vroleid",
                    "level",
                    "vopenid",
                    "vgamesvrid"
                ],
                "is_fillna": 0,
                "sort": []
            }
        },
        "presto_58384822": {
            "model": "presto",
            "model_name": "Presto",
            "type": "origin",
            "type_name": "数据源",
            "values": {
                "catalog": "hive",
                "db": "{{.dbname}}",
                "query": "select vgamesvrid,vopenid from dwd_playerregister_log where date >= '{{.stat_day}}' and date <  '{{.stat_today}}'\r\n"
            }
        },
        "presto_06ee0c23": {
            "model": "presto",
            "model_name": "Presto",
            "type": "origin",
            "type_name": "数据源",
            "values": {
                "catalog": "hive",
                "db": "{{.dbname}}",
                "query": "select vgamesvrid,vopenid,vroleid,Max(level) as level from dwd_roleleave_log where  date < '{{.stat_over}}'  group by vopenid,vroleid,vgamesvrid"
            }
        },
        "presto_ea7ade44": {
            "model": "presto",
            "model_name": "Presto",
            "type": "origin",
            "type_name": "数据源",
            "values": {
                "catalog": "hive",
                "db": "{{.dbname}}",
                "query": "select vgamesvrid,vopenid,vroleid,level,jobid,ijoblevel from dwd_roleleave_log where date < '{{.stat_over}}'"
            }
        },
        "insert_d0cced32": {
            "model": "insert",
            "model_name": "插入列",
            "type": "processor",
            "type_name": "算法",
            "values": {
                "value": "{{.stat_day}}",
                "alias_field": "stat_day"
            }
        },
        "duplicate_8f486100": {
            "model": "duplicate",
            "model_name": "去重",
            "type": "processor",
            "type_name": "算法",
            "values": {
                "duplicate_cols": [
                    "stat_day",
                    "vopenid",
                    "vroleid",
                    "level",
                    "jobid",
                    "vgamesvrid"
                ]
            }
        },
        "selectmodel_4fa81cba": {
            "model": "selectmodel",
            "model_name": "自定义筛选",
            "type": "processor",
            "type_name": "算法",
            "values": {
                "select_col": "level",
                "select_symbol": ">",
                "select_num": "0"
            }
        },
        "IMongo_53280647": {
            "model": "IMongo",
            "model_name": "IMongo",
            "type": "destination",
            "type_name": "入库",
            "values": {
                "db": "{{.dbname}}",
                "table": "test_100",
                "method": "replaced"
            }
        }
    },
    "graph": {
        "duplicate_83075415": [
            "merge_78dd692d"
        ],
        "duplicate_9c91e7cd": [
            "max_2bfe204f",
            "merge_2af6b446"
        ],
        "max_2bfe204f": [
            "merge_2af6b446"
        ],
        "merge_78dd692d": [
            "insert_d0cced32"
        ],
        "merge_2af6b446": [
            "merge_bd7dbae8"
        ],
        "merge_bd7dbae8": [
            "merge_78dd692d"
        ],
        "presto_58384822": [
            "duplicate_83075415"
        ],
        "presto_06ee0c23": [
            "duplicate_9c91e7cd"
        ],
        "presto_ea7ade44": [
            "merge_bd7dbae8"
        ],
        "insert_d0cced32": [
            "duplicate_8f486100"
        ],
        "duplicate_8f486100": [
            "selectmodel_4fa81cba"
        ],
        "selectmodel_4fa81cba": [
            "IMongo_53280647"
        ],
        "IMongo_53280647": []
    },
    "page": "{\"nodes\":[{\"vA757be44\":\"\",\"type\":\"node\",\"size\":\"170*34\",\"conf\":\"[{\\\"desc\\\":\\\"不填则按行去重\\\",\\\"hidden\\\":0,\\\"type\\\":\\\"list\\\",\\\"required\\\":0,\\\"title\\\":\\\"去重字段\\\",\\\"defaultValue\\\":\\\"\\\",\\\"prop\\\":\\\"duplicate_cols\\\"}]\",\"key\":\"duplicate\",\"label\":\"去重\",\"worktype\":\"processor\",\"typename\":\"算法\",\"x\":336.5,\"y\":156,\"id\":\"83075415\",\"tooltip\":[[\"ID\",\"duplicate_83075415\"],[\"状态\",\"未运行\"]],\"workId\":\"duplicate_83075415\",\"shape\":\"duplicate_83075415CustomNode\",\"status\":\"未运行\",\"index\":0,\"run_time\":\"\",\"data\":\"\"},{\"vA757be44\":\"\",\"type\":\"node\",\"size\":\"170*34\",\"conf\":\"[{\\\"prop\\\":\\\"group\\\",\\\"hidden\\\":0,\\\"defaultValue\\\":0,\\\"type\\\":\\\"radio\\\",\\\"title\\\":\\\"分组\\\",\\\"options\\\":[{\\\"label\\\":\\\"是\\\",\\\"value\\\":1},{\\\"label\\\":\\\"否\\\",\\\"value\\\":0}],\\\"required\\\":1,\\\"desc\\\":\\\"\\\"},{\\\"prop\\\":\\\"group_cols\\\",\\\"hidden\\\":\\\"${group} === 0\\\",\\\"defaultValue\\\":\\\"\\\",\\\"type\\\":\\\"list\\\",\\\"title\\\":\\\"分组字段\\\",\\\"required\\\":1,\\\"desc\\\":\\\"\\\"},{\\\"prop\\\":\\\"colname\\\",\\\"hidden\\\":0,\\\"defaultValue\\\":\\\"\\\",\\\"type\\\":\\\"input\\\",\\\"title\\\":\\\"字段名\\\",\\\"required\\\":1,\\\"desc\\\":\\\"\\\"},{\\\"prop\\\":\\\"alias_field\\\",\\\"hidden\\\":0,\\\"defaultValue\\\":\\\"\\\",\\\"type\\\":\\\"input\\\",\\\"title\\\":\\\"别名\\\",\\\"required\\\":0,\\\"desc\\\":\\\"默认是该节点的ID名称\\\"}]\",\"key\":\"max\",\"label\":\"取最大值\",\"worktype\":\"base_calculation\",\"typename\":\"基础计算\",\"x\":657,\"y\":242,\"id\":\"2bfe204f\",\"tooltip\":[[\"ID\",\"max_2bfe204f\"],[\"状态\",\"未运行\"]],\"workId\":\"max_2bfe204f\",\"shape\":\"max_2bfe204fCustomNode\",\"status\":\"未运行\",\"index\":1,\"run_time\":\"\",\"data\":\"\"},{\"vA757be44\":\"\",\"type\":\"node\",\"size\":\"170*34\",\"conf\":\"[{\\\"desc\\\":\\\"不填则按行去重\\\",\\\"hidden\\\":0,\\\"type\\\":\\\"list\\\",\\\"required\\\":0,\\\"title\\\":\\\"去重字段\\\",\\\"defaultValue\\\":\\\"\\\",\\\"prop\\\":\\\"duplicate_cols\\\"}]\",\"key\":\"duplicate\",\"label\":\"去重\",\"worktype\":\"processor\",\"typename\":\"算法\",\"x\":657,\"y\":157.5,\"id\":\"9c91e7cd\",\"tooltip\":[[\"ID\",\"duplicate_9c91e7cd\"],[\"状态\",\"未运行\"]],\"workId\":\"duplicate_9c91e7cd\",\"shape\":\"duplicate_9c91e7cdCustomNode\",\"status\":\"未运行\",\"index\":4,\"run_time\":\"\",\"data\":\"\"},{\"vA757be44\":\"\",\"type\":\"node\",\"size\":\"170*34\",\"conf\":\"[{\\\"prop\\\":\\\"how\\\",\\\"title\\\":\\\"合并方式\\\",\\\"type\\\":\\\"select\\\",\\\"required\\\":1,\\\"options\\\":[{\\\"label\\\":\\\"mysql_inner_join\\\",\\\"value\\\":\\\"inner\\\"},{\\\"label\\\":\\\"mysql_outer_join\\\",\\\"value\\\":\\\"outer\\\"},{\\\"label\\\":\\\"mysql_left_join\\\",\\\"value\\\":\\\"left\\\"},{\\\"label\\\":\\\"mysql_right_join\\\",\\\"value\\\":\\\"right\\\"}],\\\"defaultValue\\\":\\\"\\\",\\\"hidden\\\":0,\\\"desc\\\":\\\"遵循mysql的join方式\\\"},{\\\"prop\\\":\\\"on\\\",\\\"title\\\":\\\"合并字段\\\",\\\"type\\\":\\\"list\\\",\\\"required\\\":1,\\\"defaultValue\\\":\\\"\\\",\\\"hidden\\\":0,\\\"desc\\\":\\\"\\\"},{\\\"prop\\\":\\\"is_fillna\\\",\\\"title\\\":\\\"是否填充NULL值\\\",\\\"type\\\":\\\"radio\\\",\\\"required\\\":1,\\\"options\\\":[{\\\"label\\\":\\\"是\\\",\\\"value\\\":1},{\\\"label\\\":\\\"否\\\",\\\"value\\\":0}],\\\"defaultValue\\\":0,\\\"hidden\\\":0,\\\"desc\\\":\\\"\\\"},{\\\"prop\\\":\\\"fillna_col\\\",\\\"title\\\":\\\"填充列\\\",\\\"type\\\":\\\"list\\\",\\\"required\\\":1,\\\"defaultValue\\\":\\\"\\\",\\\"hidden\\\":\\\"${is_fillna} === 0\\\",\\\"desc\\\":\\\"a=1 or b=\\\\\\\"c\\\\\\\"\\\"},{\\\"prop\\\":\\\"sort\\\",\\\"title\\\":\\\"合并队列\\\",\\\"type\\\":\\\"list\\\",\\\"required\\\":0,\\\"defaultValue\\\":\\\"\\\",\\\"hidden\\\":0,\\\"desc\\\":\\\"合并队列可以理解为每次从队列里取头部的两个数据进行合并然后，将合并后的数据放回队列头部，依次执行直到队列为空，队列的值为节点名称,例如节点:presto_59bcef72\\\"}]\",\"key\":\"merge\",\"label\":\"合并数据\",\"worktype\":\"processor\",\"typename\":\"算法\",\"x\":828.5,\"y\":323,\"id\":\"2af6b446\",\"tooltip\":[[\"ID\",\"merge_2af6b446\"],[\"状态\",\"未运行\"]],\"workId\":\"merge_2af6b446\",\"shape\":\"merge_2af6b446CustomNode\",\"status\":\"未运行\",\"index\":7,\"run_time\":\"\",\"data\":\"\"},{\"vA757be44\":\"\",\"type\":\"node\",\"size\":\"170*34\",\"conf\":\"[{\\\"prop\\\":\\\"how\\\",\\\"title\\\":\\\"合并方式\\\",\\\"type\\\":\\\"select\\\",\\\"required\\\":1,\\\"options\\\":[{\\\"label\\\":\\\"mysql_inner_join\\\",\\\"value\\\":\\\"inner\\\"},{\\\"label\\\":\\\"mysql_outer_join\\\",\\\"value\\\":\\\"outer\\\"},{\\\"label\\\":\\\"mysql_left_join\\\",\\\"value\\\":\\\"left\\\"},{\\\"label\\\":\\\"mysql_right_join\\\",\\\"value\\\":\\\"right\\\"}],\\\"defaultValue\\\":\\\"\\\",\\\"hidden\\\":0,\\\"desc\\\":\\\"遵循mysql的join方式\\\"},{\\\"prop\\\":\\\"on\\\",\\\"title\\\":\\\"合并字段\\\",\\\"type\\\":\\\"list\\\",\\\"required\\\":1,\\\"defaultValue\\\":\\\"\\\",\\\"hidden\\\":0,\\\"desc\\\":\\\"\\\"},{\\\"prop\\\":\\\"is_fillna\\\",\\\"title\\\":\\\"是否填充NULL值\\\",\\\"type\\\":\\\"radio\\\",\\\"required\\\":1,\\\"options\\\":[{\\\"label\\\":\\\"是\\\",\\\"value\\\":1},{\\\"label\\\":\\\"否\\\",\\\"value\\\":0}],\\\"defaultValue\\\":0,\\\"hidden\\\":0,\\\"desc\\\":\\\"\\\"},{\\\"prop\\\":\\\"fillna_col\\\",\\\"title\\\":\\\"填充列\\\",\\\"type\\\":\\\"list\\\",\\\"required\\\":1,\\\"defaultValue\\\":\\\"\\\",\\\"hidden\\\":\\\"${is_fillna} === 0\\\",\\\"desc\\\":\\\"a=1 or b=\\\\\\\"c\\\\\\\"\\\"},{\\\"prop\\\":\\\"sort\\\",\\\"title\\\":\\\"合并队列\\\",\\\"type\\\":\\\"list\\\",\\\"required\\\":0,\\\"defaultValue\\\":\\\"\\\",\\\"hidden\\\":0,\\\"desc\\\":\\\"合并队列可以理解为每次从队列里取头部的两个数据进行合并然后，将合并后的数据放回队列头部，依次执行直到队列为空，队列的值为节点名称,例如节点:presto_59bcef72\\\"}]\",\"key\":\"merge\",\"label\":\"合并数据\",\"worktype\":\"processor\",\"typename\":\"算法\",\"x\":933,\"y\":424,\"id\":\"bd7dbae8\",\"tooltip\":[[\"ID\",\"merge_bd7dbae8\"],[\"状态\",\"未运行\"]],\"workId\":\"merge_bd7dbae8\",\"shape\":\"merge_bd7dbae8CustomNode\",\"status\":\"未运行\",\"index\":9,\"run_time\":\"\",\"data\":\"\"},{\"vA757be44\":\"\",\"type\":\"node\",\"size\":\"170*34\",\"conf\":\"[{\\\"prop\\\":\\\"how\\\",\\\"title\\\":\\\"合并方式\\\",\\\"type\\\":\\\"select\\\",\\\"required\\\":1,\\\"options\\\":[{\\\"label\\\":\\\"mysql_inner_join\\\",\\\"value\\\":\\\"inner\\\"},{\\\"label\\\":\\\"mysql_outer_join\\\",\\\"value\\\":\\\"outer\\\"},{\\\"label\\\":\\\"mysql_left_join\\\",\\\"value\\\":\\\"left\\\"},{\\\"label\\\":\\\"mysql_right_join\\\",\\\"value\\\":\\\"right\\\"}],\\\"defaultValue\\\":\\\"\\\",\\\"hidden\\\":0,\\\"desc\\\":\\\"遵循mysql的join方式\\\"},{\\\"prop\\\":\\\"on\\\",\\\"title\\\":\\\"合并字段\\\",\\\"type\\\":\\\"list\\\",\\\"required\\\":1,\\\"defaultValue\\\":\\\"\\\",\\\"hidden\\\":0,\\\"desc\\\":\\\"\\\"},{\\\"prop\\\":\\\"is_fillna\\\",\\\"title\\\":\\\"是否填充NULL值\\\",\\\"type\\\":\\\"radio\\\",\\\"required\\\":1,\\\"options\\\":[{\\\"label\\\":\\\"是\\\",\\\"value\\\":1},{\\\"label\\\":\\\"否\\\",\\\"value\\\":0}],\\\"defaultValue\\\":0,\\\"hidden\\\":0,\\\"desc\\\":\\\"\\\"},{\\\"prop\\\":\\\"fillna_col\\\",\\\"title\\\":\\\"填充列\\\",\\\"type\\\":\\\"list\\\",\\\"required\\\":1,\\\"defaultValue\\\":\\\"\\\",\\\"hidden\\\":\\\"${is_fillna} === 0\\\",\\\"desc\\\":\\\"a=1 or b=\\\\\\\"c\\\\\\\"\\\"},{\\\"prop\\\":\\\"sort\\\",\\\"title\\\":\\\"合并队列\\\",\\\"type\\\":\\\"list\\\",\\\"required\\\":0,\\\"defaultValue\\\":\\\"\\\",\\\"hidden\\\":0,\\\"desc\\\":\\\"合并队列可以理解为每次从队列里取头部的两个数据进行合并然后，将合并后的数据放回队列头部，依次执行直到队列为空，队列的值为节点名称,例如节点:presto_59bcef72\\\"}]\",\"key\":\"merge\",\"label\":\"合并数据\",\"worktype\":\"processor\",\"typename\":\"算法\",\"x\":726,\"y\":487,\"id\":\"78dd692d\",\"tooltip\":[[\"ID\",\"merge_78dd692d\"],[\"状态\",\"未运行\"]],\"workId\":\"merge_78dd692d\",\"shape\":\"merge_78dd692dCustomNode\",\"status\":\"未运行\",\"index\":11,\"run_time\":\"\",\"data\":\"\"},{\"vA757be44\":\"\",\"type\":\"node\",\"size\":\"170*34\",\"conf\":\"[{\\\"hidden\\\":0,\\\"title\\\":\\\"值\\\",\\\"desc\\\":\\\"\\\",\\\"prop\\\":\\\"value\\\",\\\"type\\\":\\\"input\\\",\\\"required\\\":1,\\\"defaultValue\\\":\\\"\\\"},{\\\"hidden\\\":0,\\\"title\\\":\\\"别名\\\",\\\"desc\\\":\\\"默认是该节点的ID名称\\\",\\\"prop\\\":\\\"alias_field\\\",\\\"type\\\":\\\"input\\\",\\\"required\\\":0,\\\"defaultValue\\\":\\\"\\\"}]\",\"key\":\"insert\",\"label\":\"插入列\",\"worktype\":\"processor\",\"typename\":\"算法\",\"x\":725.5,\"y\":560,\"id\":\"d0cced32\",\"tooltip\":[[\"ID\",\"insert_d0cced32\"],[\"状态\",\"未运行\"]],\"workId\":\"insert_d0cced32\",\"shape\":\"insert_d0cced32CustomNode\",\"status\":\"未运行\",\"index\":13,\"run_time\":\"\",\"data\":\"\"},{\"v-5a6fd616\":\"\",\"type\":\"node\",\"size\":\"170*34\",\"conf\":\"[{\\\"prop\\\":\\\"catalog\\\",\\\"title\\\":\\\"查询引擎类型\\\",\\\"type\\\":\\\"select\\\",\\\"required\\\":1,\\\"defaultValue\\\":\\\"kudu\\\",\\\"options\\\":[{\\\"label\\\":\\\"kudu\\\",\\\"value\\\":\\\"kudu\\\"},{\\\"label\\\":\\\"mongodb\\\",\\\"value\\\":\\\"mongodb\\\"},{\\\"label\\\":\\\"mysql\\\",\\\"value\\\":\\\"mysql\\\"},{\\\"label\\\":\\\"hive\\\",\\\"value\\\":\\\"hive\\\"},{\\\"label\\\":\\\"mongodb2\\\",\\\"value\\\":\\\"mongodb2\\\"}],\\\"hidden\\\":0,\\\"desc\\\":\\\"\\\"},{\\\"prop\\\":\\\"db\\\",\\\"title\\\":\\\"查询数据库\\\",\\\"type\\\":\\\"input\\\",\\\"required\\\":0,\\\"defaultValue\\\":\\\"\\\",\\\"hidden\\\":0,\\\"desc\\\":\\\"不填则为业务默认的数据库\\\"},{\\\"prop\\\":\\\"query\\\",\\\"title\\\":\\\"查询语句\\\",\\\"type\\\":\\\"sql\\\",\\\"required\\\":1,\\\"defaultValue\\\":\\\"\\\",\\\"hidden\\\":0,\\\"desc\\\":\\\"\\\"}]\",\"key\":\"presto\",\"label\":\"Presto\",\"worktype\":\"origin\",\"typename\":\"数据源\",\"x\":337,\"y\":74,\"id\":\"58384822\",\"tooltip\":[[\"ID\",\"presto_58384822\"],[\"状态\",\"未运行\"]],\"workId\":\"presto_58384822\",\"shape\":\"presto_58384822CustomNode\",\"status\":\"未运行\",\"index\":14,\"run_time\":\"\",\"data\":\"\"},{\"v-5a6fd616\":\"\",\"type\":\"node\",\"size\":\"170*34\",\"conf\":\"[{\\\"prop\\\":\\\"catalog\\\",\\\"title\\\":\\\"查询引擎类型\\\",\\\"type\\\":\\\"select\\\",\\\"required\\\":1,\\\"defaultValue\\\":\\\"kudu\\\",\\\"options\\\":[{\\\"label\\\":\\\"kudu\\\",\\\"value\\\":\\\"kudu\\\"},{\\\"label\\\":\\\"mongodb\\\",\\\"value\\\":\\\"mongodb\\\"},{\\\"label\\\":\\\"mysql\\\",\\\"value\\\":\\\"mysql\\\"},{\\\"label\\\":\\\"hive\\\",\\\"value\\\":\\\"hive\\\"},{\\\"label\\\":\\\"mongodb2\\\",\\\"value\\\":\\\"mongodb2\\\"}],\\\"hidden\\\":0,\\\"desc\\\":\\\"\\\"},{\\\"prop\\\":\\\"db\\\",\\\"title\\\":\\\"查询数据库\\\",\\\"type\\\":\\\"input\\\",\\\"required\\\":0,\\\"defaultValue\\\":\\\"\\\",\\\"hidden\\\":0,\\\"desc\\\":\\\"不填则为业务默认的数据库\\\"},{\\\"prop\\\":\\\"query\\\",\\\"title\\\":\\\"查询语句\\\",\\\"type\\\":\\\"sql\\\",\\\"required\\\":1,\\\"defaultValue\\\":\\\"\\\",\\\"hidden\\\":0,\\\"desc\\\":\\\"\\\"}]\",\"key\":\"presto\",\"label\":\"Presto\",\"worktype\":\"origin\",\"typename\":\"数据源\",\"x\":659,\"y\":66,\"id\":\"06ee0c23\",\"tooltip\":[[\"ID\",\"presto_06ee0c23\"],[\"状态\",\"未运行\"]],\"workId\":\"presto_06ee0c23\",\"shape\":\"presto_06ee0c23CustomNode\",\"status\":\"未运行\",\"index\":16,\"run_time\":\"\",\"data\":\"\"},{\"v-5a6fd616\":\"\",\"type\":\"node\",\"size\":\"170*34\",\"conf\":\"[{\\\"prop\\\":\\\"catalog\\\",\\\"title\\\":\\\"查询引擎类型\\\",\\\"type\\\":\\\"select\\\",\\\"required\\\":1,\\\"defaultValue\\\":\\\"kudu\\\",\\\"options\\\":[{\\\"label\\\":\\\"kudu\\\",\\\"value\\\":\\\"kudu\\\"},{\\\"label\\\":\\\"mongodb\\\",\\\"value\\\":\\\"mongodb\\\"},{\\\"label\\\":\\\"mysql\\\",\\\"value\\\":\\\"mysql\\\"},{\\\"label\\\":\\\"hive\\\",\\\"value\\\":\\\"hive\\\"},{\\\"label\\\":\\\"mongodb2\\\",\\\"value\\\":\\\"mongodb2\\\"}],\\\"hidden\\\":0,\\\"desc\\\":\\\"\\\"},{\\\"prop\\\":\\\"db\\\",\\\"title\\\":\\\"查询数据库\\\",\\\"type\\\":\\\"input\\\",\\\"required\\\":0,\\\"defaultValue\\\":\\\"\\\",\\\"hidden\\\":0,\\\"desc\\\":\\\"不填则为业务默认的数据库\\\"},{\\\"prop\\\":\\\"query\\\",\\\"title\\\":\\\"查询语句\\\",\\\"type\\\":\\\"sql\\\",\\\"required\\\":1,\\\"defaultValue\\\":\\\"\\\",\\\"hidden\\\":0,\\\"desc\\\":\\\"\\\"}]\",\"key\":\"presto\",\"label\":\"Presto\",\"worktype\":\"origin\",\"typename\":\"数据源\",\"x\":1070,\"y\":320,\"id\":\"ea7ade44\",\"tooltip\":[[\"ID\",\"presto_ea7ade44\"],[\"状态\",\"未运行\"]],\"workId\":\"presto_ea7ade44\",\"shape\":\"presto_ea7ade44CustomNode\",\"status\":\"未运行\",\"index\":18,\"run_time\":\"\",\"data\":\"\"},{\"v-5a6fd616\":\"\",\"type\":\"node\",\"size\":\"170*34\",\"conf\":\"[{\\\"desc\\\":\\\"不填则按行去重\\\",\\\"hidden\\\":0,\\\"type\\\":\\\"list\\\",\\\"required\\\":0,\\\"title\\\":\\\"去重字段\\\",\\\"defaultValue\\\":\\\"\\\",\\\"prop\\\":\\\"duplicate_cols\\\"}]\",\"key\":\"duplicate\",\"label\":\"去重\",\"worktype\":\"processor\",\"typename\":\"算法\",\"x\":726,\"y\":648,\"id\":\"8f486100\",\"tooltip\":[[\"ID\",\"duplicate_8f486100\"],[\"状态\",\"未运行\"]],\"workId\":\"duplicate_8f486100\",\"shape\":\"duplicate_8f486100CustomNode\",\"status\":\"未运行\",\"index\":21,\"run_time\":\"\",\"data\":\"\"},{\"vA757be44\":\"\",\"type\":\"node\",\"size\":\"170*34\",\"conf\":\"[{\\\"prop\\\":\\\"db\\\",\\\"title\\\":\\\"数据库名\\\",\\\"type\\\":\\\"input\\\",\\\"required\\\":0,\\\"defaultValue\\\":\\\"\\\",\\\"hidden\\\":0,\\\"desc\\\":\\\"若不填写，默认为数据库名称默认是该业务名\\\"},{\\\"prop\\\":\\\"table\\\",\\\"title\\\":\\\"表名称\\\",\\\"type\\\":\\\"input\\\",\\\"required\\\":0,\\\"defaultValue\\\":\\\"\\\",\\\"hidden\\\":0,\\\"desc\\\":\\\"若不填写，默认为表名称默认是该作业名\\\"},{\\\"prop\\\":\\\"method\\\",\\\"title\\\":\\\"方法\\\",\\\"type\\\":\\\"select\\\",\\\"required\\\":0,\\\"options\\\":[{\\\"label\\\":\\\"更新\\\",\\\"value\\\":\\\"upsert\\\"},{\\\"label\\\":\\\"插入\\\",\\\"value\\\":\\\"append\\\"},{\\\"label\\\":\\\"替换\\\",\\\"value\\\":\\\"replaced\\\"}],\\\"defaultValue\\\":\\\"upsert\\\",\\\"hidden\\\":0,\\\"desc\\\":\\\"\\\"},{\\\"prop\\\":\\\"keys\\\",\\\"title\\\":\\\"主键\\\",\\\"type\\\":\\\"list\\\",\\\"required\\\":1,\\\"defaultValue\\\":\\\"\\\",\\\"hidden\\\":\\\"${method} != \\\\\\\"upsert\\\\\\\"\\\",\\\"desc\\\":\\\"以那些字段做唯一索引\\\"}]\",\"key\":\"IMongo\",\"label\":\"IMongo\",\"worktype\":\"destination\",\"typename\":\"入库\",\"x\":726.5,\"y\":818.5,\"id\":\"53280647\",\"tooltip\":[[\"ID\",\"IMongo_53280647\"],[\"状态\",\"未运行\"]],\"workId\":\"IMongo_53280647\",\"shape\":\"IMongo_53280647CustomNode\",\"status\":\"未运行\",\"index\":22,\"run_time\":\"\",\"data\":\"\"},{\"v-5a6fd616\":\"\",\"type\":\"node\",\"size\":\"170*34\",\"conf\":\"[{\\\"prop\\\":\\\"select_col\\\",\\\"title\\\":\\\"筛选字段\\\",\\\"type\\\":\\\"input\\\",\\\"required\\\":1,\\\"defaultValue\\\":\\\"\\\",\\\"hidden\\\":0,\\\"desc\\\":\\\"选择条件筛选的字段\\\"},{\\\"prop\\\":\\\"select_symbol\\\",\\\"title\\\":\\\"选择条件\\\",\\\"type\\\":\\\"input\\\",\\\"required\\\":1,\\\"defaultValue\\\":\\\"\\\",\\\"hidden\\\":0,\\\"desc\\\":\\\"输入筛选条件的符号，比如>,<,=\\\"},{\\\"prop\\\":\\\"select_num\\\",\\\"title\\\":\\\"输入具体数值\\\",\\\"type\\\":\\\"input\\\",\\\"required\\\":1,\\\"defaultValue\\\":\\\"\\\",\\\"hidden\\\":0,\\\"desc\\\":\\\"输入数字\\\"}]\",\"key\":\"selectmodel\",\"label\":\"自定义筛选\",\"worktype\":\"processor\",\"typename\":\"算法\",\"x\":727,\"y\":736,\"id\":\"4fa81cba\",\"tooltip\":[[\"ID\",\"selectmodel_4fa81cba\"],[\"状态\",\"未运行\"]],\"workId\":\"selectmodel_4fa81cba\",\"shape\":\"selectmodel_4fa81cbaCustomNode\",\"status\":\"未运行\",\"index\":25,\"run_time\":\"\",\"data\":\"\"}],\"edges\":[{\"source\":\"83075415\",\"sourceAnchor\":2,\"target\":\"78dd692d\",\"targetAnchor\":3,\"id\":\"06b184c1\",\"shape\":\"flow-polyline-round\",\"index\":2},{\"source\":\"9c91e7cd\",\"sourceAnchor\":2,\"target\":\"2bfe204f\",\"targetAnchor\":0,\"id\":\"1f339c42\",\"shape\":\"flow-polyline-round\",\"index\":3},{\"source\":\"9c91e7cd\",\"sourceAnchor\":1,\"target\":\"2af6b446\",\"targetAnchor\":0,\"id\":\"423cf038\",\"shape\":\"flow-polyline-round\",\"index\":5},{\"source\":\"2bfe204f\",\"sourceAnchor\":2,\"target\":\"2af6b446\",\"targetAnchor\":3,\"id\":\"bf7f4fbf\",\"shape\":\"flow-polyline-round\",\"index\":6},{\"source\":\"78dd692d\",\"sourceAnchor\":2,\"target\":\"d0cced32\",\"targetAnchor\":0,\"id\":\"2e36bc2c\",\"shape\":\"flow-polyline-round\",\"index\":8},{\"source\":\"2af6b446\",\"sourceAnchor\":2,\"target\":\"bd7dbae8\",\"targetAnchor\":0,\"id\":\"7b5a719d\",\"shape\":\"flow-polyline-round\",\"index\":10},{\"source\":\"bd7dbae8\",\"sourceAnchor\":2,\"target\":\"78dd692d\",\"targetAnchor\":1,\"id\":\"a0673199\",\"shape\":\"flow-polyline-round\",\"index\":12},{\"source\":\"58384822\",\"sourceAnchor\":2,\"target\":\"83075415\",\"targetAnchor\":0,\"id\":\"81f4b97a\",\"shape\":\"flow-polyline-round\",\"index\":15},{\"source\":\"06ee0c23\",\"sourceAnchor\":2,\"target\":\"9c91e7cd\",\"targetAnchor\":0,\"id\":\"94076c9f\",\"shape\":\"flow-polyline-round\",\"index\":17},{\"source\":\"ea7ade44\",\"sourceAnchor\":2,\"target\":\"bd7dbae8\",\"targetAnchor\":1,\"id\":\"c860055c\",\"shape\":\"flow-polyline-round\",\"index\":19},{\"source\":\"d0cced32\",\"sourceAnchor\":2,\"target\":\"8f486100\",\"targetAnchor\":0,\"id\":\"e711f995\",\"shape\":\"flow-polyline-round\",\"index\":20},{\"source\":\"8f486100\",\"sourceAnchor\":2,\"target\":\"4fa81cba\",\"targetAnchor\":0,\"id\":\"59914e26\",\"shape\":\"flow-polyline-round\",\"index\":23},{\"source\":\"4fa81cba\",\"sourceAnchor\":2,\"target\":\"53280647\",\"targetAnchor\":0,\"id\":\"46edb5ac\",\"shape\":\"flow-polyline-round\",\"index\":24}]}",
    "type": 1,
    "global_values": {
        "stat_day": "2020-06-23",
        "stat_today": "2020-06-24",
        "dbname": "ro_tlog_tx_test2",
        "stat_over": "2020-06-24"
    },
    "cron": "*/1 * * * *"
}`

	res, err := simplejson.NewJson([]byte(s))
	m, err := res.Get("node").Get("duplicate_83075415").Get("values").Get("cols").Map()
	res.Get("node").Get("duplicate_83075415").Get("values").Del("cols")
	res.Get("node").Get("duplicate_83075415").Get("values").Set("cols", gin.H{m["key"].(string): m["value"]})
	s1, err := json.Marshal(res)
	fmt.Println(string(s1), err)
}

func decode(encoded []int) []int {
	n := len(encoded)
	total := 0
	for i := 1; i <= n+1; i++ {
		total ^= i
	}
	odd := 0
	for i := 1; i < n; i += 2 {
		odd ^= encoded[i]
	}
	perm := make([]int, n+1)
	perm[0] = total ^ odd
	for i, v := range encoded {
		perm[i+1] = perm[i] ^ v
	}
	return perm
}
