package main

import(
	"encoding/json"
	"myfile"
	"fmt"
	"strings"
	// "strconv"
	"regexp"
)

type GameInfoPSN struct{
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

const uspath = "./gameinfo_json_us"
const hkpath = "./gameinfo_json_hk"
const jppath = "./gameinfo_json_jp"
const eupath = "./gameinfo_json_eu"
const psnpath = "./gameinfo_json_psn"

func main(){
	mall := make(map[string]bool)
	mus := make(map[string]bool)
	mgenrepsnprofile := make(map[string]bool)
	mgenreplaystation := make(map[string]bool)
	iall := 0
	iwithus := 0

	files := myfile.GetFileList( psnpath )

	for _, f := range files{
		if !strings.HasSuffix( f, ".json" ){
			continue
		}
		//fmt.Printf("file: %s\n", f)

		j := myfile.GetFileContent( psnpath + "/" + f )

		//fmt.Printf( "json: %s\n", j )

		var game GameInfoPSN
		err := json.Unmarshal( []byte(j), &game )
		if err != nil{
			fmt.Printf( "error unmarshal: %v\n", err )
		}
		space := regexp.MustCompile(`[\s 　]+`)

		game.Name = strings.Replace( game.Name, "&amp;", "&", -1 )
		game.Name = strings.Replace( game.Name, "&nbsp;", " ", -1 )
		game.Name = strings.Replace( game.Name, "&lt;", "<", -1 )
		game.Name = strings.Replace( game.Name, "&gt;", ">", -1 )
		game.Name = strings.Replace( game.Name, "&quot;", "\"", -1 )
		game.Name = strings.Replace( game.Name, "&apos;", "'", -1 )
		game.Name = strings.Replace( game.Name, "&cent;", "￠", -1 )
		game.Name = strings.Replace( game.Name, "&pound;", "£", -1 )
		game.Name = strings.Replace( game.Name, "&yen;", "¥", -1 )
		game.Name = strings.Replace( game.Name, "&euro;", "€", -1 )
		game.Name = strings.Replace( game.Name, "&sect;", "§", -1 )
		game.Name = strings.Replace( game.Name, "&copy;", "©", -1 )
		game.Name = strings.Replace( game.Name, "&reg;", "®", -1 )
		game.Name = strings.Replace( game.Name, "&trade;", "™", -1 )
		game.Name = strings.Replace( game.Name, "&times;", "×", -1 )
		game.Name = strings.Replace( game.Name, "&divide;", "÷", -1 )

		game.Name = strings.Replace( game.Name, "®", "", -1 )
		game.Name = strings.Replace( game.Name, "™", "", -1 )
		game.Name = strings.Replace( game.Name, "PlayStation®Vita", "", -1 )
		game.Name = strings.Replace( game.Name, "PlayStation®3", "", -1 )
		game.Name = strings.Replace( game.Name, "PlayStation®4", "", -1 )
		game.Name = strings.Replace( game.Name, "PSP®", "", -1 )
		game.Name = strings.Replace( game.Name, "PS3™", "", -1 )
		game.Name = strings.Replace( game.Name, "PS4™", "", -1 )
		game.Name = strings.Replace( game.Name, "PS Vita", "", -1 )
		game.Name = strings.Replace( game.Name, "PSVita", "", -1 )
		game.Name = strings.Replace( game.Name, "PS3", "", -1 )
		game.Name = strings.Replace( game.Name, "PS4", "", -1 )
		game.Name = strings.Replace( game.Name, "PSP", "", -1 )
		game.Name = strings.Replace( game.Name, "PlayStationVita", "", -1 )
		game.Name = strings.Replace( game.Name, "PlayStation3", "", -1 )
		game.Name = strings.Replace( game.Name, "PlayStation4", "", -1 )

		Name := space.ReplaceAllString( game.Name, "" )
		plat := regexp.MustCompile(`[\(（].*?[\)）]`)
		Name = plat.ReplaceAllString( Name, `` )
		Name = strings.Replace( Name, "&", "", -1 )
		Name = strings.Replace( Name, "•", "", -1 )
		Name = strings.Replace( Name, ":", "", -1 )
		Name = strings.Replace( Name, ".", "", -1 )
		Name = strings.Replace( Name, "!", "", -1 )
		Name = strings.Replace( Name, "'", "", -1 )
		Name = strings.Replace( Name, "\"", "", -1 )
		Name = strings.Replace( Name, "\\", "", -1 )
		Name = strings.Replace( Name, "’", "", -1 )
		Name = strings.Replace( Name, "-", "", -1 )
		Name = strings.Replace( Name, ",", "", -1 )
		Name = strings.ToUpper( Name )
		// fmt.Printf( "line: %s", line )

		_, exist := mall[Name]
		if exist{
		}else{
			mall[Name] = true
		}

		Genre := strings.Split( game.Genre, ", " )
		for _, genre := range Genre {
			mgenrepsnprofile[genre] = true
		}
	}

	ius := 0
	files = myfile.GetFileList( uspath )

	for _, f := range files{
		if !strings.HasSuffix( f, ".json" ){
			continue
		}
		//fmt.Printf("file: %s\n", f)

		j := myfile.GetFileContent( uspath + "/" + f )

		//fmt.Printf( "json: %s\n", j )

		var game GameInfo
		err := json.Unmarshal( []byte(j), &game )
		if err != nil{
			fmt.Printf( "error unmarshal: %v\n", err )
		}
		space := regexp.MustCompile(`[\s 　]+`)

		game.Name = strings.Replace( game.Name, "&amp;", "&", -1 )
		game.Name = strings.Replace( game.Name, "&nbsp;", " ", -1 )
		game.Name = strings.Replace( game.Name, "&lt;", "<", -1 )
		game.Name = strings.Replace( game.Name, "&gt;", ">", -1 )
		game.Name = strings.Replace( game.Name, "&quot;", "\"", -1 )
		game.Name = strings.Replace( game.Name, "&apos;", "'", -1 )
		game.Name = strings.Replace( game.Name, "&cent;", "￠", -1 )
		game.Name = strings.Replace( game.Name, "&pound;", "£", -1 )
		game.Name = strings.Replace( game.Name, "&yen;", "¥", -1 )
		game.Name = strings.Replace( game.Name, "&euro;", "€", -1 )
		game.Name = strings.Replace( game.Name, "&sect;", "§", -1 )
		game.Name = strings.Replace( game.Name, "&copy;", "©", -1 )
		game.Name = strings.Replace( game.Name, "&reg;", "®", -1 )
		game.Name = strings.Replace( game.Name, "&trade;", "™", -1 )
		game.Name = strings.Replace( game.Name, "&times;", "×", -1 )
		game.Name = strings.Replace( game.Name, "&divide;", "÷", -1 )

		game.Name = strings.Replace( game.Name, "®", "", -1 )
		game.Name = strings.Replace( game.Name, "™", "", -1 )
		game.Name = strings.Replace( game.Name, "PlayStation®Vita", "", -1 )
		game.Name = strings.Replace( game.Name, "PlayStation®3", "", -1 )
		game.Name = strings.Replace( game.Name, "PlayStation®4", "", -1 )
		game.Name = strings.Replace( game.Name, "PSP®", "", -1 )
		game.Name = strings.Replace( game.Name, "PS3™", "", -1 )
		game.Name = strings.Replace( game.Name, "PS4™", "", -1 )
		game.Name = strings.Replace( game.Name, "PS Vita", "", -1 )
		game.Name = strings.Replace( game.Name, "PSVita", "", -1 )
		game.Name = strings.Replace( game.Name, "PS3", "", -1 )
		game.Name = strings.Replace( game.Name, "PS4", "", -1 )
		game.Name = strings.Replace( game.Name, "PSP", "", -1 )
		game.Name = strings.Replace( game.Name, "PlayStationVita", "", -1 )
		game.Name = strings.Replace( game.Name, "PlayStation3", "", -1 )
		game.Name = strings.Replace( game.Name, "PlayStation4", "", -1 )
		game.Name = space.ReplaceAllString( game.Name, ` ` )

		Name := space.ReplaceAllString( game.Name, "" )
		Name = strings.Replace( Name, "&", "", -1 )
		Name = strings.Replace( Name, "•", "", -1 )
		Name = strings.Replace( Name, ":", "", -1 )
		Name = strings.Replace( Name, ".", "", -1 )
		Name = strings.Replace( Name, "!", "", -1 )
		Name = strings.Replace( Name, "'", "", -1 )
		Name = strings.Replace( Name, "\"", "", -1 )
		Name = strings.Replace( Name, "\\", "", -1 )
		Name = strings.Replace( Name, "’", "", -1 )
		Name = strings.Replace( Name, "-", "", -1 )
		Name = strings.Replace( Name, ",", "", -1 )

		plat := regexp.MustCompile(`\(.*?\)`)
		Name = plat.ReplaceAllString( Name, `` )

		Name = strings.ToUpper( Name )
		// fmt.Printf( "line: %s", line )

		_, exist := mall[Name]
		if exist{
			if mall[Name]{
				iall++
				ius++
				mall[Name] = false
			}else{
				ius++
			}
		}else{
			// fmt.Printf("not found: %s\n", Name)
		}

		_, existus := mus[Name]
		if existus{
		}else{
			mus[Name] = true
		}

		for _, genre := range game.Metadata.Game_genre.Values {
			mgenreplaystation[genre] = true
		}
	}

	ieu := 0
	ieuwithus := 0
	files = myfile.GetFileList( eupath )

	for _, f := range files{
		if !strings.HasSuffix( f, ".json" ){
			continue
		}
		//fmt.Printf("file: %s\n", f)

		j := myfile.GetFileContent( eupath + "/" + f )

		//fmt.Printf( "json: %s\n", j )

		var game GameInfo
		err := json.Unmarshal( []byte(j), &game )
		if err != nil{
			fmt.Printf( "error unmarshal: %v\n", err )
		}
		space := regexp.MustCompile(`[\s 　]+`)

		game.Name = strings.Replace( game.Name, "&amp;", "&", -1 )
		game.Name = strings.Replace( game.Name, "&nbsp;", " ", -1 )
		game.Name = strings.Replace( game.Name, "&lt;", "<", -1 )
		game.Name = strings.Replace( game.Name, "&gt;", ">", -1 )
		game.Name = strings.Replace( game.Name, "&quot;", "\"", -1 )
		game.Name = strings.Replace( game.Name, "&apos;", "'", -1 )
		game.Name = strings.Replace( game.Name, "&cent;", "￠", -1 )
		game.Name = strings.Replace( game.Name, "&pound;", "£", -1 )
		game.Name = strings.Replace( game.Name, "&yen;", "¥", -1 )
		game.Name = strings.Replace( game.Name, "&euro;", "€", -1 )
		game.Name = strings.Replace( game.Name, "&sect;", "§", -1 )
		game.Name = strings.Replace( game.Name, "&copy;", "©", -1 )
		game.Name = strings.Replace( game.Name, "&reg;", "®", -1 )
		game.Name = strings.Replace( game.Name, "&trade;", "™", -1 )
		game.Name = strings.Replace( game.Name, "&times;", "×", -1 )
		game.Name = strings.Replace( game.Name, "&divide;", "÷", -1 )

		game.Name = strings.Replace( game.Name, "®", "", -1 )
		game.Name = strings.Replace( game.Name, "™", "", -1 )
		game.Name = strings.Replace( game.Name, "PlayStation®Vita", "", -1 )
		game.Name = strings.Replace( game.Name, "PlayStation®3", "", -1 )
		game.Name = strings.Replace( game.Name, "PlayStation®4", "", -1 )
		game.Name = strings.Replace( game.Name, "PSP®", "", -1 )
		game.Name = strings.Replace( game.Name, "PS3™", "", -1 )
		game.Name = strings.Replace( game.Name, "PS4™", "", -1 )
		game.Name = strings.Replace( game.Name, "PS Vita", "", -1 )
		game.Name = strings.Replace( game.Name, "PSVita", "", -1 )
		game.Name = strings.Replace( game.Name, "PSP", "", -1 )
		game.Name = strings.Replace( game.Name, "PS3", "", -1 )
		game.Name = strings.Replace( game.Name, "PS4", "", -1 )
		game.Name = strings.Replace( game.Name, "PSN", "", -1 )
		game.Name = strings.Replace( game.Name, "PlayStationVita", "", -1 )
		game.Name = strings.Replace( game.Name, "PlayStation3", "", -1 )
		game.Name = strings.Replace( game.Name, "PlayStation4", "", -1 )
		game.Name = strings.Replace( game.Name, "[]", "", -1 )
		game.Name = space.ReplaceAllString( game.Name, ` ` )

		Name := space.ReplaceAllString( game.Name, "" )
		Name = strings.Replace( Name, "&", "", -1 )
		Name = strings.Replace( Name, "•", "", -1 )
		Name = strings.Replace( Name, ":", "", -1 )
		Name = strings.Replace( Name, ".", "", -1 )
		Name = strings.Replace( Name, "!", "", -1 )
		Name = strings.Replace( Name, "'", "", -1 )
		Name = strings.Replace( Name, "\"", "", -1 )
		Name = strings.Replace( Name, "\\", "", -1 )
		Name = strings.Replace( Name, "’", "", -1 )
		Name = strings.Replace( Name, "-", "", -1 )
		Name = strings.Replace( Name, ",", "", -1 )

		plat := regexp.MustCompile(`\(.*?\)`)
		Name = plat.ReplaceAllString( Name, `` )

		Name = strings.ToUpper( Name )
		// fmt.Printf( "line: %s", line )

		_, exist := mall[Name]
		if exist{
			if mall[Name]{
				iall++
				ieu++
				mall[Name] = false
			}else{
				ieu++
			}
		}else{
			// fmt.Printf("not found: %s\n", Name)
		}

		_, existus := mus[Name]
		if existus{
			if mus[Name]{
				iwithus++
				ieuwithus++
				mus[Name] = false
			}else{
				ieuwithus++
			}
		}else{
		}

		for _, genre := range game.Metadata.Game_genre.Values {
			mgenreplaystation[genre] = true
		}
	}

	ijp := 0
	files = myfile.GetFileList( jppath )

	for _, f := range files{
		if !strings.HasSuffix( f, ".json" ){
			continue
		}
		//fmt.Printf("file: %s\n", f)

		j := myfile.GetFileContent( jppath + "/" + f )

		//fmt.Printf( "json: %s\n", j )

		var game GameInfo
		err := json.Unmarshal( []byte(j), &game )
		if err != nil{
			fmt.Printf( "error unmarshal: %v\n", err )
		}
		space := regexp.MustCompile(`[\s 　]+`)

		game.Name = strings.Replace( game.Name, "&amp;", "&", -1 )
		game.Name = strings.Replace( game.Name, "&nbsp;", " ", -1 )
		game.Name = strings.Replace( game.Name, "&lt;", "<", -1 )
		game.Name = strings.Replace( game.Name, "&gt;", ">", -1 )
		game.Name = strings.Replace( game.Name, "&quot;", "\"", -1 )
		game.Name = strings.Replace( game.Name, "&apos;", "'", -1 )
		game.Name = strings.Replace( game.Name, "&cent;", "￠", -1 )
		game.Name = strings.Replace( game.Name, "&pound;", "£", -1 )
		game.Name = strings.Replace( game.Name, "&yen;", "¥", -1 )
		game.Name = strings.Replace( game.Name, "&euro;", "€", -1 )
		game.Name = strings.Replace( game.Name, "&sect;", "§", -1 )
		game.Name = strings.Replace( game.Name, "&copy;", "©", -1 )
		game.Name = strings.Replace( game.Name, "&reg;", "®", -1 )
		game.Name = strings.Replace( game.Name, "&trade;", "™", -1 )
		game.Name = strings.Replace( game.Name, "&times;", "×", -1 )
		game.Name = strings.Replace( game.Name, "&divide;", "÷", -1 )

		game.Name = strings.Replace( game.Name, "®", "", -1 )
		game.Name = strings.Replace( game.Name, "™", "", -1 )
		game.Name = strings.Replace( game.Name, "PlayStation®Vita", "", -1 )
		game.Name = strings.Replace( game.Name, "PlayStation®3", "", -1 )
		game.Name = strings.Replace( game.Name, "PlayStation®4", "", -1 )
		game.Name = strings.Replace( game.Name, "PSP®", "", -1 )
		game.Name = strings.Replace( game.Name, "PS3™", "", -1 )
		game.Name = strings.Replace( game.Name, "PS4™", "", -1 )
		game.Name = strings.Replace( game.Name, "PS Vita", "", -1 )
		game.Name = strings.Replace( game.Name, "PSVita", "", -1 )
		game.Name = strings.Replace( game.Name, "PS3", "", -1 )
		game.Name = strings.Replace( game.Name, "PS4", "", -1 )
		game.Name = strings.Replace( game.Name, "PSP", "", -1 )
		game.Name = strings.Replace( game.Name, "PlayStationVita", "", -1 )
		game.Name = strings.Replace( game.Name, "PlayStation3", "", -1 )
		game.Name = strings.Replace( game.Name, "PlayStation4", "", -1 )
		game.Name = space.ReplaceAllString( game.Name, ` ` )

		Name := space.ReplaceAllString( game.Name, "" )
		Name = strings.Replace( Name, "&", "", -1 )
		Name = strings.Replace( Name, "•", "", -1 )
		Name = strings.Replace( Name, ":", "", -1 )
		Name = strings.Replace( Name, ".", "", -1 )
		Name = strings.Replace( Name, "!", "", -1 )
		Name = strings.Replace( Name, "'", "", -1 )
		Name = strings.Replace( Name, "\"", "", -1 )
		Name = strings.Replace( Name, "\\", "", -1 )
		Name = strings.Replace( Name, "’", "", -1 )
		Name = strings.Replace( Name, "-", "", -1 )
		Name = strings.Replace( Name, ",", "", -1 )

		plat := regexp.MustCompile(`[\(（].*?[\)）]`)
		Name = plat.ReplaceAllString( Name, `` )

		Name = strings.ToUpper( Name )
		// fmt.Printf( "line: %s", line )

		_, exist := mall[Name]
		if exist{
			if mall[Name]{
				iall++
				ijp++
				mall[Name] = false
			}else{
				ijp++
			}
		}else{
			// fmt.Printf("not found: %s\n", Name)
		}

		for _, genre := range game.Metadata.Game_genre.Values {
			mgenreplaystation[genre] = true
		}
	}

	ihk := 0
	ihkwithus := 0
	files = myfile.GetFileList( hkpath )

	for _, f := range files{
		if !strings.HasSuffix( f, ".json" ){
			continue
		}
		//fmt.Printf("file: %s\n", f)

		j := myfile.GetFileContent( hkpath + "/" + f )

		//fmt.Printf( "json: %s\n", j )

		var game GameInfo
		err := json.Unmarshal( []byte(j), &game )
		if err != nil{
			fmt.Printf( "error unmarshal: %v\n", err )
		}
		space := regexp.MustCompile(`[\s 　]+`)

		Name := ""
		if len(game.Metadata.Hiragana.Values) > 0 {
			Name = game.Metadata.Hiragana.Values[0]
		}

		Name = strings.Replace( Name, "&amp;", "&", -1 )
		Name = strings.Replace( Name, "&nbsp;", " ", -1 )
		Name = strings.Replace( Name, "&lt;", "<", -1 )
		Name = strings.Replace( Name, "&gt;", ">", -1 )
		Name = strings.Replace( Name, "&quot;", "\"", -1 )
		Name = strings.Replace( Name, "&apos;", "'", -1 )
		Name = strings.Replace( Name, "&cent;", "￠", -1 )
		Name = strings.Replace( Name, "&pound;", "£", -1 )
		Name = strings.Replace( Name, "&yen;", "¥", -1 )
		Name = strings.Replace( Name, "&euro;", "€", -1 )
		Name = strings.Replace( Name, "&sect;", "§", -1 )
		Name = strings.Replace( Name, "&copy;", "©", -1 )
		Name = strings.Replace( Name, "&reg;", "®", -1 )
		Name = strings.Replace( Name, "&trade;", "™", -1 )
		Name = strings.Replace( Name, "&times;", "×", -1 )
		Name = strings.Replace( Name, "&divide;", "÷", -1 )

		Name = strings.Replace( Name, "®", "", -1 )
		Name = strings.Replace( Name, "™", "", -1 )
		Name = strings.Replace( Name, "PlayStation®Vita", "", -1 )
		Name = strings.Replace( Name, "PlayStation®3", "", -1 )
		Name = strings.Replace( Name, "PlayStation®4", "", -1 )
		Name = strings.Replace( Name, "PSP®", "", -1 )
		Name = strings.Replace( Name, "PS3™", "", -1 )
		Name = strings.Replace( Name, "PS4™", "", -1 )
		Name = strings.Replace( Name, "PS Vita", "", -1 )
		Name = strings.Replace( Name, "PSVita", "", -1 )
		Name = strings.Replace( Name, "PS3", "", -1 )
		Name = strings.Replace( Name, "PS4", "", -1 )
		Name = strings.Replace( Name, "PSP", "", -1 )
		Name = strings.Replace( Name, "PlayStationVita", "", -1 )
		Name = strings.Replace( Name, "PlayStation3", "", -1 )
		Name = strings.Replace( Name, "PlayStation4", "", -1 )

		Name = space.ReplaceAllString( Name, "" )
		Name = strings.Replace( Name, "&", "", -1 )
		Name = strings.Replace( Name, "•", "", -1 )
		Name = strings.Replace( Name, ":", "", -1 )
		Name = strings.Replace( Name, ".", "", -1 )
		Name = strings.Replace( Name, "!", "", -1 )
		Name = strings.Replace( Name, "'", "", -1 )
		Name = strings.Replace( Name, "\"", "", -1 )
		Name = strings.Replace( Name, "\\", "", -1 )
		Name = strings.Replace( Name, "’", "", -1 )
		Name = strings.Replace( Name, "-", "", -1 )
		Name = strings.Replace( Name, ",", "", -1 )

		plat := regexp.MustCompile(`[\(（].*?[\)）]`)
		Name = plat.ReplaceAllString( Name, `` )

		Name = strings.ToUpper( Name )
		Name = strings.Replace( Name, "FULLGAME", "", -1 )
		// fmt.Printf( "line: %s", line )

		_, exist := mall[Name]
		if exist{
			if mall[Name]{
				iall++
				ihk++
				mall[Name] = false
			}else{
				ihk++
			}
		}else{
			// fmt.Printf("not found: %s\n", Name)
		}

		_, existus := mus[Name]
		if existus{
			if mus[Name]{
				iwithus++
				ihkwithus++
				mus[Name] = false
			}else{
				ihkwithus++
			}
		}else{
		}

		for _, genre := range game.Metadata.Game_genre.Values {
			mgenreplaystation[genre] = true
		}
	}

	for name, value := range mall{
		if value {
			fmt.Printf("not found: %s\n", name)
		}
	}

	c := ""
	for name, _ := range mgenrepsnprofile{
		c += name + "\n"
	}
	myfile.SaveFile( "./genre_psnprofile.txt", []byte(c) )

	c = ""
	for name, _ := range mgenreplaystation{
		c += name + "\n"
	}
	myfile.SaveFile( "./genre_playstation.txt", []byte(c) )

	fmt.Printf("%d\nus:%d\neu:%d\njp:%d\nhk:%d\n", iall, ius, ieu, ijp, ihk)
	fmt.Printf("%d\neu:%d\nhk:%d\n", iwithus, ieuwithus, ihkwithus)

}