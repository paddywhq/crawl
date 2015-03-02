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

const  path =  "./gamelink_html"

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
	final_index := 1000

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
		url := "http://psnprofiles.com/games/?order=a-z&letter=&platform=&p=" + strconv.Itoa(i)
		filePath := path + "/" + strconv.Itoa(i) + ".html"
		j := myweb.CrawlJob{ url, filePath }

		linkChannel <- j
	}

	// wg.Wait()
}