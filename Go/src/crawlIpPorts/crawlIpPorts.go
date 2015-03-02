package main

import(
	"fmt"
	"./myjob"
	"./myfile"
	"net/http" 
	"./myweb"
	"./myregex"
	"io/ioutil"
	"net/url"
	"./myrand"
	"./mytime"
	"time"
)

const patternForIpPort = `<td>(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})</td>\s*<td>(\d+)</td>`
const patternForIpPorts = `<table class="sortable"([\s\S]*?)</table>`


func main(){
	fmt.Printf("爬取开始\n")
	myjob.Run(3, crawl)
}


func ChangeProxy(p string) *http.Client{
	proxyUrl, err := url.Parse(p)
	if err != nil {
	    fmt.Println("Bad proxy URL", err)
	    return nil
	}

	tr := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl), 
	}

	client := &http.Client{
		Transport: tr,
	}
	return client;
}


/*
* 每次随机拿一个代理，然后创建http.Client去抓网页
* 成功了标记，不成功不管他（下次还会再爬）
*/
func crawl() string{
	resMsg := mytime.LogTime() + " crawl."

	proxies := myfile.GetFileLines("./proxy.txt")
	proxyUrl := "http://" + proxies[ myrand.RandInt(0, len(proxies)) ]
	resMsg += "via:" + proxyUrl + "."

	client := ChangeProxy( proxyUrl )

	IpPorts := crawlIpPort( client )

	if len(IpPorts) >= 1{
		outputFile := "./proxy0.txt"
		fmt.Printf("output to :%s\n", outputFile)

		c := ""
		for _, IpPort := range IpPorts{
			t0 := time.Now()
			proxyUrlTest := "http://" + IpPort

			clientTest := ChangeProxy( proxyUrlTest )
			resp, err := myweb.Crawl( "http://www.baidu.com/", clientTest )

			if err==nil && resp!=nil{
				body, _ := ioutil.ReadAll(resp.Body)
				t1 := time.Now()

				if t1.Sub(t0).Seconds() > 20{
					fmt.Printf("%s time error\n", IpPort)
					continue
				}

				if len(string(body)) < 100{
					fmt.Printf("%s nothing crawled\n", IpPort)
					continue
				}

				// fmt.Printf("output to :%s\n", string(body))
				fmt.Printf("%s OK\n", IpPort)
				c += IpPort + "\n"
			}
		}
		myfile.SaveFile(outputFile, []byte(c) )
		return "success\n"
	}

	return "error crwaling: nothing crwaled\n"
}


func crawlIpPort( client *http.Client ) []string{
	url := "http://cn-proxy.com/"
	fmt.Printf("crawling: %s\n", url)
	resp, err := myweb.Crawl( url, client )
	IpPorts := make([]string, 0, 100)
	
	if err==nil && resp!=nil{
		//把信息拿出来
		body, _ := ioutil.ReadAll(resp.Body)  
		table := myregex.Parse( string(body), patternForIpPorts )

		if len(table) >= 2{
			arr := myregex.Parse( table[1][1], patternForIpPort )

			for _, addr := range arr{
				if len(addr) >= 3{
					Ip := addr[1]
					Port := addr[2]
					IpPorts = append( IpPorts, Ip + ":" + Port )
				}
			}
		}else{
			fmt.Printf("error crwaling: nothing crwaled\n")
		}
	}else{
		fmt.Printf("error crwaling: %s, msg:%s\n", url, err)
	}
	return IpPorts
}