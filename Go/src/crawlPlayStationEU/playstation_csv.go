package main

import(
	"encoding/json"
	"./myfile"
	"fmt"
	"strings"
	// "strconv"
	"regexp"
)

type ImageInfo struct{
	Url string
}

type DetailInfo struct{
	Name string
	Values []string
}

type MetadataInfo struct{
	Genre DetailInfo
	Hiragana DetailInfo
	Game_genre DetailInfo
	Playable_platform DetailInfo
}

type StarInfo struct{
	Star int
	Count int
}

type StarRatingInfo struct{
	Total string
	Score string
	Count []StarInfo
}

type MaterialsInfo struct{
	Urls []ImageInfo
}

type PromomediaInfo struct{
	Materials []MaterialsInfo
}

type SkusInfo struct{
	Display_price string
	Name string
}

type GameInfo struct{
	Id string
	Images []ImageInfo
	Long_desc string
	Metadata MetadataInfo
	Provider_name string
	Release_date string
	Name string
	Star_rating StarRatingInfo
	Promomedia []PromomediaInfo
	Skus []SkusInfo
}

const path = "./gameinfo_json"

func main(){

	files := myfile.GetFileList( path )

	csv := "Images,Long_desc,Genre,Hiragana,Game_genre,Playable_platform,Provider_name,Release_date,Name,Promomedia,Version\r\n"
	
	m := make(map[string]GameInfo)
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

		b := regexp.MustCompile(`<b[\s\S]*?>`)
		game.Long_desc = b.ReplaceAllString( game.Long_desc, "" )
		game.Long_desc = strings.Replace( game.Long_desc, "</b>", "", -1 )
		font := regexp.MustCompile(`<font[\s\S]*?>`)
		game.Long_desc = font.ReplaceAllString( game.Long_desc, "" )
		game.Long_desc = strings.Replace( game.Long_desc, "</font>", "", -1 )
		center := regexp.MustCompile(`<center[\s\S]*?>`)
		game.Long_desc = center.ReplaceAllString( game.Long_desc, "" )
		game.Long_desc = strings.Replace( game.Long_desc, "</center>", "", -1 )
		span := regexp.MustCompile(`<span[\s\S]*?>`)
		game.Long_desc = span.ReplaceAllString( game.Long_desc, "" )
		game.Long_desc = strings.Replace( game.Long_desc, "</span>", "", -1 )
		img := regexp.MustCompile(`<img[\s\S]*?>`)
		game.Long_desc = img.ReplaceAllString( game.Long_desc, "" )
		game.Long_desc = strings.Replace( game.Long_desc, "<br/>", " ", -1 )
		game.Long_desc = strings.Replace( game.Long_desc, "<br>", " ", -1 )
		game.Long_desc = strings.Replace( game.Long_desc, "</br>", " ", -1 )
		game.Long_desc = strings.Replace( game.Long_desc, "<BR/>", " ", -1 )
		game.Long_desc = strings.Replace( game.Long_desc, "<BR>", " ", -1 )
		game.Long_desc = strings.Replace( game.Long_desc, "</BR>", " ", -1 )
		game.Long_desc = strings.Replace( game.Long_desc, "<Br/>", " ", -1 )
		game.Long_desc = strings.Replace( game.Long_desc, "<Br>", " ", -1 )
		game.Long_desc = strings.Replace( game.Long_desc, "</Br>", " ", -1 )
		//game.Long_desc = strings.Replace( game.Long_desc, "\n", " ", -1 )
		game.Long_desc = strings.Replace( game.Long_desc, "\r", "\n", -1 )
		game.Long_desc = strings.Replace( game.Long_desc, "\t", " ", -1 )
		game.Long_desc = strings.Replace( game.Long_desc, "\f", "", -1 )
		game.Long_desc = strings.Replace( game.Long_desc, "\v", "", -1 )
		game.Long_desc = strings.Replace( game.Long_desc, "\\\"", "\"", -1 )
		game.Long_desc = strings.Replace( game.Long_desc, "\"", "\"\"", -1 )
		//space := regexp.MustCompile(`[\s 　]+`)
		//game.Long_desc = space.ReplaceAllString( game.Long_desc, ` ` )
		// plat := regexp.MustCompile(`\(.*?\)`)
		// game.Name = plat.ReplaceAllString( game.Name, `` )

		Name := game.Name
		Name = strings.Replace( Name, "®", "", -1 )
		Name = strings.Replace( Name, "™", "", -1 )
		Name = strings.Replace( Name, "PlayStationVita", "", -1 )
		Name = strings.Replace( Name, "PlayStation3", "", -1 )
		Name = strings.Replace( Name, "PlayStation4", "", -1 )
		Name = strings.Replace( Name, "PSP", "", -1 )
		Name = strings.Replace( Name, "PS3", "", -1 )
		Name = strings.Replace( Name, "PS4", "", -1 )
		Name = strings.Replace( Name, "PS Vita", "", -1 )
		Name = strings.Replace( Name, "PSN", "", -1 )
		Name = strings.Replace( Name, "()", "", -1 )
		Name = strings.Replace( Name, "[]", "", -1 )
		Name = space.ReplaceAllString( Name, ` ` )
		Name = space.ReplaceAllString( Name, "" )
		Name = strings.Replace( Name, ":", "", -1 )
		Name = strings.Replace( Name, ".", "", -1 )
		Name = strings.Replace( Name, "!", "", -1 )
		Name = strings.Replace( Name, "'", "", -1 )
		Name = strings.Replace( Name, "’", "", -1 )
		Name = strings.Replace( Name, "-", "", -1 )
		Name = strings.Replace( Name, ",", "", -1 )
		Name = strings.ToUpper( Name )

		// Version := ""
		
		// if len(game.Skus) >= 1{
		// 	Version = game.Skus[0].Name
		// }

		value, exist := m[Name]
		if exist{
			value.Metadata.Playable_platform.Values = append(value.Metadata.Playable_platform.Values, game.Metadata.Playable_platform.Values...)
			value.Metadata.Genre.Values = append(value.Metadata.Genre.Values, game.Metadata.Genre.Values...)
			value.Metadata.Game_genre.Values = append(value.Metadata.Game_genre.Values, game.Metadata.Game_genre.Values...)
			value.Promomedia = append(value.Promomedia, game.Promomedia...)
			if value.Release_date > game.Release_date && game.Release_date != ""{
				value.Release_date = game.Release_date
			}
			m[Name] = value
		}else{
			m[Name] = game
		}
	}

	for _, game := range m {
		Version := ""
		// Display_price := ""

		if len(game.Skus) >= 1{
			Version = game.Skus[0].Name
			// Display_price = game.Skus[0].Display_price
		}

		platforms := make(map[string]bool)
		for _, platform := range game.Metadata.Playable_platform.Values{
			platforms[platform] = true
		}
		Playable_platform := ""
		for platform, _ := range platforms{
			Playable_platform += platform + " "
		}

		genres := make(map[string]bool)
		for _, genre := range game.Metadata.Genre.Values{
			genres[genre] = true
		}
		Genre := ""
		for genre, _ := range genres{
			Genre += genre + " "
		}

		game_genres := make(map[string]bool)
		for _, game_genre := range game.Metadata.Game_genre.Values{
			game_genres[game_genre] = true
		}
		Game_genre := ""
		for game_genre, _ := range game_genres{
			Game_genre += game_genre + " "
		}

		promomedias := make(map[string]bool)
		for _, promomedia := range game.Promomedia{
			for _, material := range promomedia.Materials{
				for _, url := range material.Urls{
					promomedias[url.Url] = true
				}
			}
		}
		Promomedia := ""
		for url, _ := range promomedias{
			Promomedia += url + " "
		}

		line := fmt.Sprintf( 
			"\"%+v\",\"%s\",\"%s\",\"%+v\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\"\r\n",
			game.Images, game.Long_desc, Genre, game.Metadata.Hiragana.Values, Game_genre, Playable_platform, game.Provider_name, game.Release_date, game.Name, Promomedia, Version)

		csv += line
		// fmt.Printf( "line: %s", line )
	}

	csv = strings.Replace( csv, "&amp;", "&", -1 )
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
	csv = strings.Replace( csv, "&trade;", "™", -1 )
	csv = strings.Replace( csv, "&times;", "×", -1 )
	csv = strings.Replace( csv, "&divide;", "÷", -1 )
	myfile.SaveFile( "./games.csv", []byte(csv) )
}