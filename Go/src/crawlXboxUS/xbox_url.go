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

const  path = "./gamelink_html"
const  xbox360url = "http://marketplace.xbox.com/en-US/Games?page="
const  xbox1url = "http://store.xbox.com/en-US/Xbox-One?Page="

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

	index := 0
	final_index := 0

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

	index = 1
	final_index = 64

	for i := index; i <= final_index; i++{
		url := xbox360url + strconv.Itoa(i)
		filePath := path + "/360_" + strconv.Itoa(i) + ".html"
		j := myweb.CrawlJob{ url, filePath }

		linkChannel <- j
	}

	index = 1
	final_index = 11

	for i := index; i <= final_index; i++{
		url := xbox1url + strconv.Itoa(i)
		filePath := path + "/1_" + strconv.Itoa(i) + ".html"
		j := myweb.CrawlJob{ url, filePath }

		linkChannel <- j
	}

	wg.Wait()
}