package main

import(
	"fmt"
	// "./myjob"
	"./myfile"
	"strings"
	"net/http" 
	"./myweb"
	// "./myregex"
	// "io/ioutil"
	"os"
	"net/url"
	"./myrand"
	"./mytime"
	// "strconv"
	"sync"
	//"regexp"
)

const urlFilePath = `./gameUrl.txt`
const path = "./gameinfo_html"
const urlXbox360Base = "http://marketplace.xbox.com"
const urlXbox1Base = "http://store.xbox.com"

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

func main(){
	os.Mkdir( path, os.ModeSticky | 0755)

	gameInfos := myfile.GetFileLines( urlFilePath )
	
	wg := new(sync.WaitGroup)
	linkChannel := make(chan myweb.CrawlJob, 5)

	const  worker_count = 10
	for i:=0; i<worker_count; i++{
		proxies := myfile.GetFileLines("./proxy.txt")
		resMsg := mytime.LogTime() + " crawl."

		proxyUrl := "http://" + proxies[ myrand.RandInt(0, len(proxies)) ]
		resMsg += "via:" + proxyUrl + "."

		client := ChangeProxy( proxyUrl )

		wg.Add(1)
		go myweb.CrawlAndSaveWorker( linkChannel, client, path, wg )
	}

	for _, info := range gameInfos{
		url := info
		urlSplit := strings.Split( url, "/" )
		var filePath string
		if strings.Contains(url, urlXbox360Base) {
			filePath = path + "/360_" + urlSplit[ len(urlSplit)-2 ] + "_" + strings.Replace(urlSplit[ len(urlSplit)-1 ], "?nosplash=1", "", -1) + ".html"
			filePath = strings.Replace(filePath, "*", "", -1)
		} else {
			filePath = path + "/1_" + urlSplit[ len(urlSplit)-2 ] + "_" + strings.Replace(urlSplit[ len(urlSplit)-1 ], "?nosplash=1", "", -1) + ".html"
			filePath = strings.Replace(filePath, "*", "", -1)
		}

		j := myweb.CrawlJob{ url, filePath }

		linkChannel <- j
	}

	wg.Wait()
}