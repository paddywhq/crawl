package main

import(
	"fmt"
	"./myfile"
	"./myregex"
	// "strconv"
	"os"
	"encoding/json"
	"strings"
	"regexp"
)

type GameInfo struct{
	AchievementsEarned string
	PlayersTracked string
	TotalAchievements string
	Gamerscore string
	Challenges string
	AllClub string
	Name string
	Image string
	Platform string
	Icon string
}

const readPath = "./gameinfo_html"
const writePath = "./gameinfo_json"
const patternForAchievementsEarned = `<strong>(.*?)</strong>\s*?Achievements earned`
const patternForPlayersTracked = `<strong>(.*?)</strong>\s*?Players Tracked`
const patternForTotalAchievements = `<strong>(.*?)</strong>\s*?Total Achievements`
const patternForGamerscore = `<strong>(.*?)</strong>\s*?Gamerscore`
const patternForChallenges = `<strong>(.*?)</strong>\s*?Challenges`
const patternForAllClub = `<strong>(.*?)</strong>\s*?100% Club`
const patternForName = `<h2><a href=".*?" rel="bookmark" title=".*?">\s*(.*?)\s*</a></h2>`
const patternForImage = `url\((.*?)\)`
const patternForPlatform = `<div style="margin-left: 6px" class="inline-pf generic">\s*(.*?)\s*</div>`
const patternForIcon = `<a class="image" href=".*?"><img src="(.*?)" /></a>`

func main(){

	os.Mkdir( writePath, os.ModeSticky | 0755)

	files := myfile.GetFileList( readPath )
	for _, f := range files{
		fmt.Printf("\nfile: %s\n", f)

		l := myfile.GetFileContent( readPath + "/" + f )

		info := GameInfo{}

		achievementsEarned := myregex.Parse( l, patternForAchievementsEarned )
		if len(achievementsEarned) == 1 && len(achievementsEarned[0]) == 2{
			info.AchievementsEarned = achievementsEarned[0][1]
		}

		playersTracked := myregex.Parse( l, patternForPlayersTracked )
		if len(playersTracked) == 1 && len(playersTracked[0]) == 2{
			info.PlayersTracked = playersTracked[0][1]
		}

		totalAchievements := myregex.Parse( l, patternForTotalAchievements )
		if len(totalAchievements) == 1 && len(totalAchievements[0]) == 2{
			info.TotalAchievements = totalAchievements[0][1]
		}

		gamerscore := myregex.Parse( l, patternForGamerscore )
		if len(gamerscore) == 1 && len(gamerscore[0]) == 2{
			info.Gamerscore = gamerscore[0][1]
		}

		challenges := myregex.Parse( l, patternForChallenges )
		if len(challenges) == 1 && len(challenges[0]) == 2{
			info.Challenges = challenges[0][1]
		}

		allClub := myregex.Parse( l, patternForAllClub )
		if len(allClub) == 1 && len(allClub[0]) == 2{
			info.AllClub = allClub[0][1]
		}

		name := myregex.Parse( l, patternForName )
		if len(name) == 1 && len(name[0]) == 2{
			info.Name = name[0][1]
			space := regexp.MustCompile(`\s+`)
			info.Name = space.ReplaceAllString( info.Name, ` ` )
		}

		image := myregex.Parse( l, patternForImage )
		if len(image) == 1 && len(image[0]) == 2{
			info.Image = image[0][1]
		}

		platform := myregex.Parse( l, patternForPlatform )
		if len(platform) >= 1 && len(platform[0]) == 2{
			info.Platform = platform[0][1]
		}

		icon := myregex.Parse( l, patternForIcon )
		if len(icon) == 1 && len(icon[0]) == 2{
			info.Icon = icon[0][1]
		}

		j,err := json.Marshal( &info )
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Printf( "json: %s\n", string(j) )

		myfile.SaveFile( writePath + "/" + strings.Replace( f, ".html", ".json", -1 ), []byte(j) )
	}
}