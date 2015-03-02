package main

import(
	"encoding/json"
	"./myfile"
	"fmt"
	"strings"
	// "strconv"
	// "regexp"
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

const path = "./gameinfo_json"

func main(){

	files := myfile.GetFileList( path )

	csv := "AchievementsEarned,PlayersTracked,TotalAchievements,Gamerscore,Challenges,AllClub,Name,Image,Platform,Icon\r\n"
	
	for _, f := range files{
		if !strings.HasSuffix( f, ".json" ){
			continue
		}
		fmt.Printf("file: %s\n", f)

		j := myfile.GetFileContent( path + "/" + f )

		//fmt.Printf( "json: %s\n", j )

		var game GameInfo
		err := json.Unmarshal( []byte(j), &game )
		if err != nil{
			fmt.Printf( "error unmarshal: %v\n", err )
		}

		line := fmt.Sprintf( 
			"\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\"\r\n",
			game.AchievementsEarned, game.PlayersTracked, game.TotalAchievements, game.Gamerscore, game.Challenges, game.AllClub, game.Name, game.Image, game.Platform, game.Icon)

		csv += line
		// fmt.Printf( "line: %s", line )
	}

	csv = strings.Replace( csv, "\\\"", "\"\"", -1 )
	
	csv = strings.Replace( csv, "&#8230;", "…", -1 )
	csv = strings.Replace( csv, "&#8217;", "'", -1 )
	csv = strings.Replace( csv, "&#8211;", "-", -1 )
	csv = strings.Replace( csv, "&#038;", "&", -1 )
	csv = strings.Replace( csv, "&amp;", "&", -1 )
	csv = strings.Replace( csv, "&nbsp;", " ", -1 )
	csv = strings.Replace( csv, "&lt;", "<", -1 )
	csv = strings.Replace( csv, "&gt;", ">", -1 )
	csv = strings.Replace( csv, "&quot;", "\"", -1 )
	csv = strings.Replace( csv, "&apos;", "'", -1 )
	csv = strings.Replace( csv, "&cent;", "￠", -1 )
	csv = strings.Replace( csv, "&pound;", "£", -1 )
	csv = strings.Replace( csv, "&yen;", "¥", -1 )
	csv = strings.Replace( csv, "&euro;", "€", -1 )
	csv = strings.Replace( csv, "&sect;", "§", -1 )
	csv = strings.Replace( csv, "&copy;", "©", -1 )
	csv = strings.Replace( csv, "&reg;", "®", -1 )
	csv = strings.Replace( csv, "&trade;", "™", -1 )
	csv = strings.Replace( csv, "&times;", "×", -1 )
	csv = strings.Replace( csv, "&divide;", "÷", -1 )
	myfile.SaveFile( "./games.csv", []byte(csv) )
}