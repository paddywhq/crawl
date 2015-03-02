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
	GameOwners string
	PlatinumAchievers string
	AverageCompletion string
	TrophiesEarned string
	AllCompleted string
	Name string
	Image string
	Platform string
	Developer string
	Publisher string
	Genre string
	Theme string
}

const readPath = "./gameinfo_html"
const writePath = "./gameinfo_json"
const patternForGameOwners = `<span class="stat">(.*?)<span>Game Owners</span></span>`
const patternForPlatinumAchievers = `<span class="stat">(.*?)<span>Platinum Achievers</span></span>`
const patternForAverageCompletion = `<span class="stat">(.*?)<span>Average Completion</span></span>`
const patternForTrophiesEarned = `<span class="stat">(.*?)<span>Trophies Earned</span></span>`
const patternForAllCompleted = `<span class="stat">(.*?)<span>100% Completed</span></span>`
const patternForName = `<h3 class="ellipsis" style="height: 19px;width: \d*?px;">\s*([\s\S]*?)\s*</h3>`
const patternForImage = `<img class="game-image fancy"[\s\S]*?src="(.*?)"`
const patternForPlatform = `<img[\s\S]*?alt="(.*?)"\s*/?>`
const patternForDeveloper = `<tr>\s*<td>Developers?</td>\s*<td>(.*?)</td>\s*</tr>`
const patternForPublisher = `<tr>\s*<td>Publishers?</td>\s*<td>(.*?)</td>\s*</tr>`
const patternForGenre = `<tr>\s*<td>Genres?</td>\s*<td>(.*?)</td>\s*</tr>`
const patternForTheme = `<tr>\s*<td>Themes?</td>\s*<td>(.*?)</td>\s*</tr>`

func main(){

	os.Mkdir( writePath, os.ModeSticky | 0755)

	files := myfile.GetFileList( readPath )
	for _, f := range files{
		fmt.Printf("\nfile: %s\n", f)

		l := myfile.GetFileContent( readPath + "/" + f )

		info := GameInfo{}

		gameOwners := myregex.Parse( l, patternForGameOwners )
		if len(gameOwners) == 1 && len(gameOwners[0]) == 2{
			info.GameOwners = gameOwners[0][1]
		}

		platinumAchievers := myregex.Parse( l, patternForPlatinumAchievers )
		if len(platinumAchievers) == 1 && len(platinumAchievers[0]) == 2{
			info.PlatinumAchievers = platinumAchievers[0][1]
		}

		averageCompletion := myregex.Parse( l, patternForAverageCompletion )
		if len(averageCompletion) == 1 && len(averageCompletion[0]) == 2{
			info.AverageCompletion = averageCompletion[0][1]
		}

		trophiesEarned := myregex.Parse( l, patternForTrophiesEarned )
		if len(trophiesEarned) == 1 && len(trophiesEarned[0]) == 2{
			info.TrophiesEarned = trophiesEarned[0][1]
		}

		allCompleted := myregex.Parse( l, patternForAllCompleted )
		if len(allCompleted) == 1 && len(allCompleted[0]) == 2{
			info.AllCompleted = allCompleted[0][1]
		}

		name := myregex.Parse( l, patternForName )
		if len(name) == 1 && len(name[0]) == 2{
			info.Name = name[0][1]
			info.Name = strings.Replace( info.Name, " Trophies", "", -1 )
			space := regexp.MustCompile(`\s+`)
			info.Name = space.ReplaceAllString( info.Name, ` ` )
		}

		image := myregex.Parse( l, patternForImage )
		if len(image) == 1 && len(image[0]) == 2{
			info.Image = image[0][1]
		}

		platform := myregex.Parse( l, patternForPlatform )
		info.Platform = ""
		if len(platform) >= 1 {
			for i, content := range platform{
				if i != 0 {
					if i == 1{
						info.Platform += content[1]
					}else{
						info.Platform += ", " + content[1]
					}
				}
			}
		}

		developer := myregex.Parse( l, patternForDeveloper )
		if len(developer) == 1 && len(developer[0]) == 2{
			info.Developer = developer[0][1]
			info.Developer = strings.Replace( info.Developer, "<nobr>", "", -1 )
			info.Developer = strings.Replace( info.Developer, "</nobr>", "", -1 )
		}

		publisher := myregex.Parse( l, patternForPublisher )
		if len(publisher) == 1 && len(publisher[0]) == 2{
			info.Publisher = publisher[0][1]
			info.Publisher = strings.Replace( info.Publisher, "<nobr>", "", -1 )
			info.Publisher = strings.Replace( info.Publisher, "</nobr>", "", -1 )
		}

		genre := myregex.Parse( l, patternForGenre )
		if len(genre) == 1 && len(genre[0]) == 2{
			info.Genre = genre[0][1]
			info.Genre = strings.Replace( info.Genre, "<nobr>", "", -1 )
			info.Genre = strings.Replace( info.Genre, "</nobr>", "", -1 )
		}

		theme := myregex.Parse( l, patternForTheme )
		if len(theme) == 1 && len(theme[0]) == 2{
			info.Theme = theme[0][1]
			info.Theme = strings.Replace( info.Theme, "<nobr>", "", -1 )
			info.Theme = strings.Replace( info.Theme, "</nobr>", "", -1 )
		}

		j,err := json.Marshal( &info )
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Printf( "json: %s\n", string(j) )

		myfile.SaveFile( writePath + "/" + strings.Replace( f, ".html", "", -1 ) + ".json", []byte(j) )
	}
}