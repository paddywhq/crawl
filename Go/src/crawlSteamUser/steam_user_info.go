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

const  allGamesPath = "./allGames"
const  recentPlayedPath = "./recentPlayed"
const  wishlistPath = "./wishlist"
const  patternForUserId = `"id":"(\d*)"`
const  patternForGameId = `"appid":(\d*),`
const  patternForWishlist = `<a href="http://steamcommunity.com/app/(\d*)">`
const  patternForNoWishlist = `<title>Steam Community :: .*? :: Games</title>`
const  patternForNotFound = `<div class="error_ctn">`

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

func CrawlUserGamesSteamAllGames(userName string) bool {
	os.Mkdir( allGamesPath, os.ModeSticky | 0755)

	proxies := myfile.GetFileLines("./proxy.txt")
	
	for ;; {
		resMsg := mytime.LogTime() + " crawl."

		proxyUrl := "http://" + proxies[ myrand.RandInt(0, len(proxies)) ]
		resMsg += "via:" + proxyUrl + "."

		//client := ChangeProxy( proxyUrl )
		client := MyProxy( )

		url := "http://steamcommunity.com/id/" + userName + "/games/?tab=all"
		fmt.Printf("crawling: %s\n", url)
		resp, err := myweb.Crawl( url, client )

		if err==nil && resp!=nil{
			//把信息拿出来
			body, _ := ioutil.ReadAll(resp.Body)  
			// fmt.Printf("%s", string(body))

			myfile.SaveFile( allGamesPath + "/x_" + userName + ".txt", []byte(string(body)) )
			games := myregex.Parse( string(body), patternForGameId )
			notfound := myregex.Parse( string(body), patternForNotFound )

			c := ""

			if len(notfound) >= 1{
				fmt.Printf("User not found.\n")
				return false
			} 

			if len(games) >= 0{
				m := make(map[string]bool)

				fmt.Printf("Success.\n")
				for _, game := range games {
					_, exist := m[game[1]]
					if exist{
					}else{
						c += game[1] + "\n"
						m[game[1]] = true
					}
				}

				fmt.Printf("Saving...\n")
				myfile.SaveFile( allGamesPath + "/" + userName + ".txt", []byte(c) )
				break
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
	return true
}

func CrawlUserGamesSteamRecentPlayed(userName string) {
	os.Mkdir( recentPlayedPath, os.ModeSticky | 0755)

	proxies := myfile.GetFileLines("./proxy.txt")
	
	for ;; {
		resMsg := mytime.LogTime() + " crawl."

		proxyUrl := "http://" + proxies[ myrand.RandInt(0, len(proxies)) ]
		resMsg += "via:" + proxyUrl + "."

		//client := ChangeProxy( proxyUrl )
		client := MyProxy( )

		url := "http://steamcommunity.com/id/" + userName + "/games/?tab=recent"
		fmt.Printf("crawling: %s\n", url)
		resp, err := myweb.Crawl( url, client )

		if err==nil && resp!=nil{
			//把信息拿出来
			body, _ := ioutil.ReadAll(resp.Body)  
			// fmt.Printf("%s", string(body))

			myfile.SaveFile( recentPlayedPath + "/x_" + userName + ".txt", []byte(string(body)) )
			games := myregex.Parse( string(body), patternForGameId )
			notfound := myregex.Parse( string(body), patternForNotFound )

			c := ""

			if len(notfound) >= 1{
				fmt.Printf("User not found.\n")
				break
			} 

			if len(games) >= 0{
				m := make(map[string]bool)

				fmt.Printf("Success.\n")
				for _, game := range games {
					_, exist := m[game[1]]
					if exist{
					}else{
						c += game[1] + "\n"
						m[game[1]] = true
					}
				}

				fmt.Printf("Saving...\n")
				myfile.SaveFile( recentPlayedPath + "/" + userName + ".txt", []byte(c) )
				break
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

func CrawlUserGamesSteamWishlist(userName string) {
	os.Mkdir( wishlistPath, os.ModeSticky | 0755)

	proxies := myfile.GetFileLines("./proxy.txt")
	
	for ;; {
		resMsg := mytime.LogTime() + " crawl."

		proxyUrl := "http://" + proxies[ myrand.RandInt(0, len(proxies)) ]
		resMsg += "via:" + proxyUrl + "."

		//client := ChangeProxy( proxyUrl )
		client := MyProxy( )

		url := "http://steamcommunity.com/id/" + userName + "/wishlist"
		fmt.Printf("crawling: %s\n", url)
		resp, err := myweb.Crawl( url, client )

		if err==nil && resp!=nil{
			//把信息拿出来
			body, _ := ioutil.ReadAll(resp.Body)  
			// fmt.Printf("%s", string(body))

			myfile.SaveFile( wishlistPath + "/x_" + userName + ".txt", []byte(string(body)) )
			games := myregex.Parse( string(body), patternForWishlist )
			notfound := myregex.Parse( string(body), patternForNotFound )
			noWishlist := myregex.Parse( string(body), patternForNoWishlist )

			c := ""

			if len(notfound) >= 1{
				fmt.Printf("User not found.\n")
				break
			} 

			if len(noWishlist) == 0{
				fmt.Printf("Saving...\n")
				myfile.SaveFile( wishlistPath + "/" + userName + ".txt", []byte(c) )
				break
			} 

			if len(games) >= 0{
				m := make(map[string]bool)

				fmt.Printf("Success.\n")
				for _, game := range games {
					_, exist := m[game[1]]
					if exist{
					}else{
						c += game[1] + "\n"
						m[game[1]] = true
					}
				}

				fmt.Printf("Saving...\n")
				myfile.SaveFile( wishlistPath + "/" + userName + ".txt", []byte(c) )
				break
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

func CrawlUserGamesSteam(userName string) {
	if CrawlUserGamesSteamAllGames(userName) {
		CrawlUserGamesSteamRecentPlayed(userName)
		CrawlUserGamesSteamWishlist(userName)
	} else {
		fmt.Printf("Wrong username.\n")
	}
}

func main() {
	CrawlUserGamesSteam("sdvlksfnvkclkashkalkdjf")
	CrawlUserGamesSteam("mdudu")
	CrawlUserGamesSteam("cascascas")
	CrawlUserGamesSteam("cas")
	// http://steamcommunity.com/profiles/76561198086715847
}