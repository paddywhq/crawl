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
	Index string
	Title string
	Description string
	Context string
	Content_type string
	Feature string
	Keywords string
	Platform string
	Genre string
	// Rating string
	Image string
	Video_url string
	// Video_desc string
	Image_url string
	// image_desc string
	Release_date string
	Release_date_display string
	Url string
	PsnId string
	// ButtonTxt string
	// ButtonURL string
	// ButtonColor string
	AgeRatingImage string
	AgeRatingImageAlt string
	Detail string
	Publisher string
	Developer string
}

const path = "./gameinfo_json"

func main(){

	files := myfile.GetFileList( path )

	csv := "index,title,description,context,content_type,feature,keywords,platform,genre,image,video_url,image_url,release_date,release_date_display,url,psnId,ageRatingImage,ageRatingImageAlt,publisher,developer,detail\r\n"
	
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
			"\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\"\r\n",
			game.Index, game.Title, game.Description, game.Context, game.Content_type, game.Feature, game.Keywords, game.Platform, game.Genre, game.Image, game.Video_url, game.Image_url, game.Release_date, game.Release_date_display, game.Url, game.PsnId, game.AgeRatingImage, game.AgeRatingImageAlt, game.Publisher, game.Developer, game.Detail)

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