package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/msadg/userAgent"
)

const (
	UA   = "ua"
	DT   = "dt"
	RAND = "rand"
	ALL  = "all"
	NULL = ""
)

func main() {
	// add := ":9580"
	var port string
	flag.StringVar(&port, "p", ":9580", "HTTP Server Port")
	// /api?ua={<rand>,<browserName>,<[all]>}[&dt=<rand/count/[all]>]
	// ua : rand / browserName / all
	// browserName: {chrome, opera, firefox, ie, safari, edge}
	// if ua == browserName and dt(datatype)={<rand>,<count>,[all]}
	// dt : rand / count / [all]default
	// rand 随机
	// count 多少条
	// all 默认的模式，就是不写
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		if r.Method == http.MethodGet {

			ua := r.FormValue(UA)
			switch ua {

			case RAND: // 随机单条信息
				data := userAgent.Rand()
				w.Write(toJson(data))

			case ALL: // 所有 User-Agent
				data := userAgent.All()
				w.Write(toJson(data))

			default: // 判断是不是浏览器
				if ua == "ie" {
					ua = "internet explorer"
				} else if ua == "edge" {
					ua = "microsoft edge"
				}

				browserNames := userAgent.ListBrowsers()

				for _, name := range browserNames {
					if name == ua {
						// 判断需要怎么获取信息
						dt := r.FormValue(DT)
						switch dt {

						case RAND: // 随机一条
							d := userAgent.RandBs(name)
							w.Write(toJson(d))

						case NULL: // 没有写
							fallthrough

						case ALL: // 全部获取
							data := userAgent.BrowserAll(name) // 获取浏览器所有 User-Agent
							w.Write(toJson(data))

						default: // 其他参数
							w.Write(toJson([]byte("参数错误")))
						}
					}
				}
				w.Write(nil)
			}
		}
	})
	log.Fatal(http.ListenAndServe(port, nil))
}

// 转换成JSON
func toJson(v any) []byte {
	bs, err := json.Marshal(v)
	checkError(err)
	return bs
}

// 错误检查
func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
