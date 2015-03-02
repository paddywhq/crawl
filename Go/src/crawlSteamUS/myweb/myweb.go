package myweb

import (
	"fmt"
	"net/http" 
	"time"
	//"strings"
	"sync"
	"io/ioutil" 
	"../myfile"
)

type CrawlJob struct{
	Url string
	File string
}

func Crawl( url string, client *http.Client ) (*http.Response, error) {
	time.Sleep(1000 * time.Millisecond)
	
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Cookie", "birthtime=476092801; fakeCC=US;")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8")
	resp, err := client.Do(req)

	defer func() {
		if err:=recover(); err!=nil{
			fmt.Printf( "got err: %v", err )
		}
	}()

	return resp, err
}


func CrawlAndSaveWorker( linkChannel chan CrawlJob, client *http.Client, rootPath string, wg *sync.WaitGroup){
	defer wg.Done()

	for job := range linkChannel{
		if myfile.FileExists( job.File ){
			continue
		}

		resp, err := Crawl( job.Url, client )
		if err != nil {
			fmt.Printf( "GET error %v\n", err )
		}
		if resp == nil{
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)  
		if err != nil {
		 	fmt.Println("Read error is: %v\n", err)
		}
		
		myfile.SaveFile( job.File, body )
		fmt.Printf( "Url:%s, Size:%d \n", job.Url, len(body) )
		resp.Body.Close()

	}
}