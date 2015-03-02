package myweb

import (
	"fmt"
	"net/http" 
	"time"
	//"strings"
	"sync"
	"io/ioutil" 
	"myfile"
)

type CrawlJob struct{
	Url string
	File string
}

func Crawl( url string, client *http.Client ) (*http.Response, error) {
	time.Sleep(1000 * time.Millisecond)
	
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122 Safari/537.36")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Add("Cookie", "GeoIP=US:Absecon:39.4899:-74.4773:v4; uls-previous-languages=%5B%22zh-cn%22%5D; mediaWiki.user.sessionId=uWZBQIGC5ZjlmClgOl39ACHVCTk8Jw8Q; TBLkisOn=0")
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