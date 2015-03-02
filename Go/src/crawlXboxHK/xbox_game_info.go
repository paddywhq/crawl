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
	Name string
	Icon_vertical string
	Icon_horizontal string
	Detail string
	Date string
	Developer string
	Publisher string
	Genre string
	Image []string
	Id string
	Platform string
}

const readPathLink = "./gamelink_html"
const readPathInfo = "./gameinfo_html"
const writePath = "./gameinfo_json"

const patternForXbox360Result = `<h2><a href="(.*?)" title=".*?">.*?</a></h2>`
const patternForXbox360Name = `<h1>(.*?)</h1>`
const patternForXbox360IconVertical = `<img alt=".*?" class="boxart" src="(.*?)" title=".*?" />`
const patternForXbox360IconHorizontal = `banner.png`
const patternForXbox360Detail = `<div class="Text">\s*<p>([\s\S]*?)</p>\s*</div>`
const patternForXbox360Date = `<li><label>原?來?的?發行日期\s*：</label>\s*(\d+/\d+/\d+)\s*</li>`
const patternForXbox360Developer = `<li><label>開發商\s*：</label>\s*(.*?)\s*</li>`
const patternForXbox360Publisher = `<li><label>發行者\s*：</label>\s*(.*?)\s*</li>`
const patternForXbox360Genre = `<li><label>遊戲類型\s*：</label>\s*(.*?)\s*</li>`
const patternForXbox360Image = `<div id="image\d+" class="TabPage image">\s*<img src="(.*?)" alt=".*?" />\s*</div>`

const patternForXbox1Result = `<a data-alt=".*?" data-contentid=".*?" data-slot="\d*" href="(.*?)"><img alt=".*?" src=".*?&amp;format=png&amp;h=294&amp;w=215" title=".*?" /></a>`
const patternForXbox1Name = `<h1 class="title">(.*?)</h1>`
const patternForXbox1IconVertical = `<img alt=".*?" src="(.*?)&amp;format=png&amp;h=294&amp;w=215" title=".*?" />`
const patternForXbox1IconHorizontal = `<img alt="" class="superHeroImage" id="superHeroImage" src="(.*?)&amp;format=jpg" />`
const patternForXbox1Detail = `<label>.*?說明:</label><span>([\s\S]*?)</span>`
const patternForXbox1Date = `<label>發行日期:</label>\s*<span>(\d+/\d+/\d+)</span>`
const patternForXbox1Developer = `<label>開發商:</label><span>([\s\S]*?)</span>`
const patternForXbox1Publisher = `<label>發行者:</label><span>([\s\S]*?)</span>`
const patternForXbox1Genre = `<label>類型:</label><span>([\s\S]*?)</span>`
const patternForXbox1Image = `url: "(.*?)\\u0026format=jpg\\u0026h=640\\u0026w=1138",`

func main(){

	os.Mkdir( writePath, os.ModeSticky | 0755)

	files := myfile.GetFileList( readPathLink )
	for _, f := range files{
		fmt.Printf("\nfile: %s\n", f)

		linkFile := myfile.GetFileContent( readPathLink + "/" + f )

		if strings.Contains(f, "360_") {
			result := myregex.Parse( linkFile, patternForXbox360Result )
			for _, game := range result{
				info := GameInfo{}

				url := game[1]
				urlSplit := strings.Split( url, "/" )

				if len(urlSplit) < 2 {
					continue
				}

				filePath := readPathInfo + "/360_" + urlSplit[ len(urlSplit)-2 ] + "_" + urlSplit[ len(urlSplit)-1 ] + ".html"
				filePath = strings.Replace(filePath, "*", "", -1)

				htmlFile := myfile.GetFileContent( filePath )

				name := myregex.Parse( htmlFile, patternForXbox360Name )
				if len(name) >= 1 {
					info.Name = name[0][1]
				}

				iconVertical := myregex.Parse( htmlFile, patternForXbox360IconVertical )
				if len(iconVertical) >= 1 {
					info.Icon_vertical = iconVertical[0][1]
					info.Icon_horizontal = strings.Replace( info.Icon_vertical, "boxartlg.jpg", patternForXbox360IconHorizontal, -1 )
				}

				detail := myregex.Parse( htmlFile, patternForXbox360Detail )
				if len(detail) >= 1 {
					info.Detail = detail[0][1]
				}

				date := myregex.Parse( htmlFile, patternForXbox360Date )
				if len(date) >= 1 {
					info.Date = date[0][1]
				}

				developer := myregex.Parse( htmlFile, patternForXbox360Developer )
				if len(developer) >= 1 {
					info.Developer = developer[0][1]
				}

				publisher := myregex.Parse( htmlFile, patternForXbox360Publisher )
				if len(publisher) >= 1 {
					info.Publisher = publisher[0][1]
				}

				genre := myregex.Parse( htmlFile, patternForXbox360Genre )
				if len(genre) >= 1 {
					info.Genre = genre[0][1]
				}

				image := myregex.Parse( htmlFile, patternForXbox360Image )
				if len(image) >= 1 {
					for _, content := range image{
						info.Image = append(info.Image, content[1])
					}
				}

				info.Id = urlSplit[ len(urlSplit)-1 ]
				info.Platform = "Xbox 360"

				j,err := json.Marshal( &info )
				if err != nil {
					fmt.Println(err)
				}
				//fmt.Printf( "json: %s\n", string(j) )

				myfile.SaveFile( writePath + "/360_" + urlSplit[ len(urlSplit)-2 ] + "_" + urlSplit[ len(urlSplit)-1 ] + ".json", []byte(j) )
			}
		} else {
			result := myregex.Parse( linkFile, patternForXbox1Result )
			for _, game := range result{
				info := GameInfo{}

				url := game[1]
				if strings.Contains(url, "/zh-HK/Xbox-One/Bundle/") {
					continue
				}

				urlSplit := strings.Split( url, "/" )
				filePath := readPathInfo + "/1_" + urlSplit[ len(urlSplit)-2 ] + "_" + urlSplit[ len(urlSplit)-1 ] + ".html"
				filePath = strings.Replace(filePath, "*", "", -1)

				htmlFile := myfile.GetFileContent( filePath )

				name := myregex.Parse( htmlFile, patternForXbox1Name )
				if len(name) >= 1 {
					info.Name = name[0][1]
				}

				iconVertical := myregex.Parse( game[0], patternForXbox1IconVertical )
				if len(iconVertical) >= 1 {
					info.Icon_vertical = iconVertical[0][1]
				}

				iconHorizontal := myregex.Parse( htmlFile, patternForXbox1IconHorizontal )
				if len(iconHorizontal) >= 1 {
					info.Icon_horizontal = iconHorizontal[0][1]
				}

				detail := myregex.Parse( htmlFile, patternForXbox1Detail )
				if len(detail) >= 1 {
					info.Detail = detail[0][1]
					info.Detail = strings.Replace( info.Detail, "\n", " ", -1 )
					info.Detail = strings.Replace( info.Detail, "\r", " ", -1 )
					info.Detail = strings.Replace( info.Detail, "\t", " ", -1 )
					info.Detail = strings.Replace( info.Detail, "\f", "", -1 )
					info.Detail = strings.Replace( info.Detail, "\v", "", -1 )
					info.Detail = strings.Replace( info.Detail, "\\\"", "\"", -1 )
					info.Detail = strings.Replace( info.Detail, "\"", "\"\"", -1 )
					space := regexp.MustCompile(`[\s 　]+`)
					info.Detail = space.ReplaceAllString( info.Detail, ` ` )
				}

				date := myregex.Parse( htmlFile, patternForXbox1Date )
				if len(date) >= 1 {
					info.Date = date[0][1]
				}

				developer := myregex.Parse( htmlFile, patternForXbox1Developer )
				if len(developer) >= 1 {
					info.Developer = developer[0][1]
					space := regexp.MustCompile(`[\s 　]+`)
					info.Developer = space.ReplaceAllString( info.Developer, ` ` )
				}

				publisher := myregex.Parse( htmlFile, patternForXbox1Publisher )
				if len(publisher) >= 1 {
					info.Publisher = publisher[0][1]
					space := regexp.MustCompile(`[\s 　]+`)
					info.Publisher = space.ReplaceAllString( info.Publisher, ` ` )
				}

				genre := myregex.Parse( htmlFile, patternForXbox1Genre )
				if len(genre) >= 1 {
					info.Genre = genre[0][1]
					space := regexp.MustCompile(`[\s 　]+`)
					info.Genre = space.ReplaceAllString( info.Genre, ` ` )
				}

				image := myregex.Parse( htmlFile, patternForXbox1Image )
				if len(image) >= 1 {
					for _, content := range image{
						info.Image = append(info.Image, content[1])
					}
				}

				info.Id = urlSplit[ len(urlSplit)-1 ]
				info.Platform = "Xbox One"

				j,err := json.Marshal( &info )
				if err != nil {
					fmt.Println(err)
				}
				//fmt.Printf( "json: %s\n", string(j) )

				myfile.SaveFile( writePath + "/1_" + urlSplit[ len(urlSplit)-2 ] + "_" + urlSplit[ len(urlSplit)-1 ] + ".json", []byte(j) )
			}
		}
	}
}