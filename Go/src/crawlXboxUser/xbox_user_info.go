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
	"./myregex"
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

func CrawlUserAchievementsXbox(userName string, userId string) {
	proxies := myfile.GetFileLines("./proxy.txt")
	
	for ;; {
		resMsg := mytime.LogTime() + " crawl."

		proxyUrl := "http://" + proxies[ myrand.RandInt(0, len(proxies)) ]
		resMsg += "via:" + proxyUrl + "."

		//client := ChangeProxy( proxyUrl )
		client := MyProxy( )

		url := "https://achievements.xboxlive.com/users/xuid(" + userId + ")/achievements?orderBy=UnlockTime&maxItems=50000"
		fmt.Printf("crawling: %s\n", url)
		url = `https://account.xbox.com/passport/setCookies.ashx?rru=https%3a%2f%2faccount.xbox.com%2fzh-CN%2fAccount%2fSignin&wa=wsignin1.0`
		//url = `https://account.xbox.com/Account/Signin`
		resp, err := myweb.GetCookieToken( url, client )
		url = "https://achievements.xboxlive.com/users/xuid(" + userId + ")/achievements?orderBy=UnlockTime&maxItems=50000"
		resp, err = myweb.Crawl( url, client )

		if err==nil && resp!=nil{
			//把信息拿出来
			body, _ := ioutil.ReadAll(resp.Body)  
			// fmt.Printf("%s", string(body))

			myfile.SaveFile( path + "/x_" + userName + ".txt", []byte(string(body)) )
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
		} else {
			fmt.Printf("error crwaling: %s, msg:%s\n", url, err)
		}
	}

	// myfile.SaveFile( path + "/" + userName + ".txt", []byte(c) )
}

func CrawlUserXbox(userName string)  bool {
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
				CrawlUserAchievementsXbox(userName, id)
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

func CrawlUserGamesXbox(userName string) {
	if CrawlUserXbox(userName) {
		fmt.Printf("Success.\n")
	} else {
		fmt.Printf("Wrong username.\n")
	}
}

func main() {
	os.Mkdir( path, os.ModeSticky | 0755)

	CrawlUserGamesXbox( "jdfblakscbclebalivwiuhflkj" )
	CrawlUserGamesXbox( "kasaan" )
	CrawlUserGamesXbox( "A Royal Hobo" )
	CrawlUserGamesXbox( "CI BO VICE21" )
	CrawlUserGamesXbox( "briixxloveee" )
	//CrawlUserAchievementsXbox("kasaan","2533274807950917")
}