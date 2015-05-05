package main

import (
	"fmt"
	"strconv"
	"os"
	"net/http" 
	"sync"
	"./myweb"
	"./myfile"
	"./mytime"
	"./myrand"
	"net/url"
)

const  path =  "./gamelink_json/"

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

func main() {
	os.Mkdir( path, os.ModeSticky | 0755)

	index := 1
	final_index := 59

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

	for i := index; i <= final_index; i++{
		url := "http://www.playstation.com/en-us/explore/gamefinder/_jcr_content/content_par/gamefinder.query.json?_=&callback=jQuery18109199425186961889_1422498579295&count=100&i=1&page=" + strconv.Itoa(i) + "&q1=Game&q2=out_now&sort=release_date_desc&x1=content_type&x2=release_cat&callback=jQuery18109199425186961889_1422498579295"
		filePath := path + "/" + strconv.Itoa(i) + ".json"
		j := myweb.CrawlJob{ url, filePath }

		linkChannel <- j
	}

	// wg.Wait()
}