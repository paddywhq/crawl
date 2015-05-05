package main

import (
	"fmt"
	// "strconv"
	"os"
	"net/http"
	"io/ioutil"
	// "sync"
	"./myweb"
	"./myfile"
	"./mytime"
	"./myrand"
//	"./myregex"
	"net/url"
)

const  path = "./userInfo"
const  patternForUserId = `"id":"(\d*)"`
const  patternForGameId = `"productId":"(.*?)"`
const  patternForNotFound = `The server found no data for the requested entity`

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

func MyProxy() *http.Client{
	client := &http.Client{
	}
	return client;
}

func CrawlUserAchievementsPlayStation(userName string) {
	proxies := myfile.GetFileLines("./proxy.txt")
	
	for ;; {
		resMsg := mytime.LogTime() + " crawl."

		proxyUrl := "http://" + proxies[ myrand.RandInt(0, len(proxies)) ]
		resMsg += "via:" + proxyUrl + "."

		//client := ChangeProxy( proxyUrl )
		client := MyProxy( )

		url := "https://my.playstation.com/playstation/psn/public/trophies/?onlineId="+userName//&_=1429175993318
		fmt.Printf("crawling: %s\n", url)
		resp, err := myweb.Crawl( url, client )

		if err==nil && resp!=nil{
			//把信息拿出来
			body, _ := ioutil.ReadAll(resp.Body)  
			// fmt.Printf("%s", string(body))

			myfile.SaveFile( path + "/x_" + userName + ".txt", []byte(string(body)) )
/*
			achievements := myregex.Parse( string(body), patternForGameId )

			c := ""

			if len(achievements) >= 1{
				m := make(map[string]bool)

				fmt.Printf("Success.\n")
				for _, achievement := range achievements {
					_, exist := m[achievement[1]]
					if exist{
					}else{
						c += achievement[1] + "\n"
						m[achievement[1]] = true
					}
				}

				fmt.Printf("Saving...\n")
				myfile.SaveFile( path + "/" + userName + ".txt", []byte(c) )
				break;
				// myfile.SaveFile( path + "/" + userName + ".txt", []byte(c) )
			} else {
				fmt.Printf("Not found.\n")
				fmt.Printf("%s\n", string(body))
			}
*/
			break
		} else {
			fmt.Printf("error crwaling: %s, msg:%s\n", url, err)
		}
	}

	// myfile.SaveFile( path + "/" + userName + ".txt", []byte(c) )
}

/*
func CrawlUserPlayStation(userName string)  bool {
	proxies := myfile.GetFileLines("./proxy.txt")
	
	for ;; {
		resMsg := mytime.LogTime() + " crawl."

		proxyUrl := "http://" + proxies[ myrand.RandInt(0, len(proxies)) ]
		resMsg += "via:" + proxyUrl + "."

		//client := ChangeProxy( proxyUrl )
		client := MyProxy( )

		url := "https://profile.xboxlive.com/users/gt(" + userName + ")/profile/settings"
		fmt.Printf("crawling: %s\n", url)
		resp, err := myweb.Crawl( url, client )

		if err==nil && resp!=nil{
			//把信息拿出来
			body, _ := ioutil.ReadAll(resp.Body)  
			// fmt.Printf("%s", string(body))
			idJson := myregex.Parse( string(body), patternForUserId )
			notfound := myregex.Parse( string(body), patternForNotFound )

			if len(notfound) >= 1{
				fmt.Printf("Id not found.\n")
				return false
			} 

			if len(idJson) >= 1{
				id := idJson[0][1]
				CrawlUserAchievementsPlayStation(userName, id)
				break;
			} else {
				fmt.Printf("Not found.\n")
				fmt.Printf("%s\n", string(body))
				//myfile.SaveFile( "./0.txt", []byte(string(body)) )
			}
		} else {
			fmt.Printf("error crwaling: %s, msg:%s\n", url, err)
		}
	}

	// myfile.SaveFile( path + "/" + userName + ".txt", []byte(c) )
	return true
}
*/

func CrawlUserGamesPlayStation(userName string) {
	CrawlUserAchievementsPlayStation(userName)
	//if CrawlUserPlayStation(userName) {
	//	fmt.Printf("Success.\n")
	//} else {
	//	fmt.Printf("Wrong username.\n")
	//}
}

func main() {
	os.Mkdir( path, os.ModeSticky | 0755)

	CrawlUserGamesPlayStation( "jdfblakscbclebalivwiuhflkj" )
	CrawlUserGamesPlayStation( "guishou928" )
	//CrawlUserAchievementsXbox("kasaan","2533274807950917")
}