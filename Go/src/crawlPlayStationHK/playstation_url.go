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

const path =  "./gamelink_json"
const baseUrl = `https://store.sonyentertainmentnetwork.com/chihiroview/viewfinder?https%3A%2F%2Fstore.sonyentertainmentnetwork.com%2Fstore%2Fapi%2Fchihiro%2F00_09_000%2Fcontainer%2FHK%2Fzh%2F999%2FSTORE-MSF86012-GAMESALL%2F0%3Fsize%3D30%26start%3D`

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

	index := 0
	final_index := 1860

	for i := index; i <= final_index; i += 30 {
		url := baseUrl + strconv.Itoa(i)
		filePath := path + "/" + strconv.Itoa(i) + ".json"
		j := myweb.CrawlJob{ url, filePath }

		linkChannel <- j
	}

	// wg.Wait()
}