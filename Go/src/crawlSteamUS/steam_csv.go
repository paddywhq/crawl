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
	Name string
	Icon_small string
	Icon_big string
	Detail string
	Date string
	Developer []string
	Publisher []string
	Genre []string
	Image []string
	Video []string
	Id string
	Platform []string
}

const path = "./gameinfo_json"

func main(){

	files := myfile.GetFileList( path )

	csv := "Id,Name,Icon_small,Icon_big,Detail,Date,Developer,Publisher,Genre,Image,Video,Platform\r\n"

	m := make(map[string]bool)
	
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
			"\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%+v\",\"%+v\",\"%+v\",\"%+v\",\"%+v\",\"%+v\"\r\n",
			game.Id, game.Name, game.Icon_small, game.Icon_big, game.Detail, game.Date, game.Developer, game.Publisher, game.Genre, game.Image, game.Video, game.Platform)

		csv += line

		for _, genre := range game.Genre {
			m[genre] = true
		}
		// fmt.Printf( "line: %s", line )
	}

	csv = strings.Replace( csv, "&#246;", "ö", -1 )
	csv = strings.Replace( csv, "&#187;", "»", -1 )
	csv = strings.Replace( csv, "&#171;", "«", -1 )
	csv = strings.Replace( csv, "&#243;", "ó", -1 )
	csv = strings.Replace( csv, "&#239;", "ï", -1 )
	csv = strings.Replace( csv, "&#172;", "", -1 )
	csv = strings.Replace( csv, "&#235;", "ë", -1 )
	csv = strings.Replace( csv, "&#215;", "×", -1 )
	csv = strings.Replace( csv, "&#173;", " ", -1 )
	csv = strings.Replace( csv, "&#179;", "³", -1 )
	csv = strings.Replace( csv, "&#178;", "²", -1 )
	csv = strings.Replace( csv, "&#225;", "á", -1 )
	csv = strings.Replace( csv, "&#233;", "é", -1 )
	csv = strings.Replace( csv, "&#237;", "í", -1 )
	csv = strings.Replace( csv, "&#227;", "ã", -1 )
	csv = strings.Replace( csv, "&#231;", "ç", -1 )
	csv = strings.Replace( csv, "&#241;", "ñ", -1 )
	csv = strings.Replace( csv, "&#169;", "©", -1 )
	csv = strings.Replace( csv, "&#160;", " ", -1 )
	csv = strings.Replace( csv, "&#252;", "ü", -1 )
	csv = strings.Replace( csv, "&#183;", "·", -1 )
	csv = strings.Replace( csv, "&amp;", "&", -1 )
	csv = strings.Replace( csv, "&#39;", "'", -1 )
	csv = strings.Replace( csv, "&nbsp;", " ", -1 )
	csv = strings.Replace( csv, "&lt;", "<", -1 )
	csv = strings.Replace( csv, "&gt;", ">", -1 )
	csv = strings.Replace( csv, "&quot;", "\"\"", -1 )
	csv = strings.Replace( csv, "&apos;", "'", -1 )
	csv = strings.Replace( csv, "&cent;", "￠", -1 )
	csv = strings.Replace( csv, "&pound;", "£", -1 )
	csv = strings.Replace( csv, "&yen;", "¥", -1 )
	csv = strings.Replace( csv, "&euro;", "€", -1 )
	csv = strings.Replace( csv, "&sect;", "§", -1 )
	csv = strings.Replace( csv, "&copy;", "©", -1 )
	csv = strings.Replace( csv, "&reg;", "®", -1 )
	csv = strings.Replace( csv, "&#174;", "®", -1 )
	csv = strings.Replace( csv, "&trade;", "™", -1 )
	csv = strings.Replace( csv, "&times;", "×", -1 )
	csv = strings.Replace( csv, "&divide;", "÷", -1 )
	myfile.SaveFile( "./games.csv", []byte(csv) )

	c := ""
	for name, _ := range m{
		c += name + "\n"
	}
	myfile.SaveFile( "./genre_steam_us.txt", []byte(c) )
}