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

const readPathLink = "./gamelink_html"
const readPathInfo = "./gameinfo_html"
const writePath = "./gameinfo_json"

const patternForResult = `<a href="(.*?)"  data-ds-appid=".*?" onmouseover=".*?" onmouseout=".*?" class="search_result_row ds_collapse_flag" >\s*<div class="col search_capsule"><img src="(.*?)" alt=".*?" width="120" height="45"></div>[\s\S]*?</p>`
const patternForName = `<b>标题:</b>\s*(.*?)<br>`
const patternForIconBig = `<img class="game_header_image_full" src="(.*?)">`
const patternForDetail = `<div id="game_area_description" class="game_area_description">\s*<h2>关于这款游戏</h2>\s*([\s\S]*?)\s*</div>`
const patternForDate = `<b>发行日期:</b>\s*(\d+年\d+月\d+日)<br>`
const patternForDeveloper = `<a href="http://store.steampowered.com/search/\?developer=.*?">(.*?)</a>`
const patternForPublisher = `<a href="http://store.steampowered.com/search/\?publisher=.*?">(.*?)</a>`
const patternForGenre = `<a href="http://store.steampowered.com/genre/.*?">(.*?)</a>`
const patternForPlatform = `<span class="platform_img (.*?)"></span>`
const patternForImage = `<a class="highlight_screenshot_link" data-screenshotid=".*?" target="_blank" href="(.*?)"?">`
const patternForVideo = `FILENAME: "(.*?)",`

func main(){

	os.Mkdir( writePath, os.ModeSticky | 0755)

	files := myfile.GetFileList( readPathLink )
	for _, f := range files{
		fmt.Printf("\nfile: %s\n", f)

		linkFile := myfile.GetFileContent( readPathLink + "/" + f )

		result := myregex.Parse( linkFile, patternForResult )
		for _, game := range result{
			info := GameInfo{}

			url := game[1]

			urlSplit := strings.Split( url, "/" )
			filePath := readPathInfo + "/" + urlSplit[ len(urlSplit)-2 ] + ".html"
			filePath = strings.Replace(filePath, "*", "", -1)

			htmlFile := myfile.GetFileContent( filePath )

			name := myregex.Parse( htmlFile, patternForName )
			if len(name) >= 1 {
				info.Name = name[0][1]
			}

			info.Icon_small = game[2]

			iconBig := myregex.Parse( htmlFile, patternForIconBig )
			if len(iconBig) >= 1 {
				info.Icon_big = iconBig[0][1]
			}

			detail := myregex.Parse( htmlFile, patternForDetail )
			if len(detail) >= 1 {
				info.Detail = detail[0][1]

				a := regexp.MustCompile("<a(.*?)>")
				info.Detail = a.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</a>", "", -1 )
				p := regexp.MustCompile("<p(.*?)>")
				info.Detail = p.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</p>", "", -1 )
				ul := regexp.MustCompile("<ul(.*?)>")
				info.Detail = ul.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</ul>", "", -1 )
				li := regexp.MustCompile("<li(.*?)>")
				info.Detail = li.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</li>", "", -1 )
				div := regexp.MustCompile("<div(.*?)>")
				info.Detail = div.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</div>", "", -1 )
				info.Detail = strings.Replace( info.Detail, "</div >", "", -1 )
				span := regexp.MustCompile("<span(.*?)>")
				info.Detail = span.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</span>", "", -1 )
				info.Detail = strings.Replace( info.Detail, "</span >", "", -1 )
				font := regexp.MustCompile("<font(.*?)>")
				info.Detail = font.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</font>", "", -1 )
				strong := regexp.MustCompile("<strong(.*?)>")
				info.Detail = strong.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</strong>", "", -1 )
				em := regexp.MustCompile("<em(.*?)>")
				info.Detail = em.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</em>", "", -1 )
				h1 := regexp.MustCompile("<h1(.*?)>")
				info.Detail = h1.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</h1>", "", -1 )
				h2 := regexp.MustCompile("<h2(.*?)>")
				info.Detail = h2.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</h2>", "", -1 )
				h3 := regexp.MustCompile("<h3(.*?)>")
				info.Detail = h3.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</h3>", "", -1 )
				h4 := regexp.MustCompile("<h4(.*?)>")
				info.Detail = h4.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</h4>", "", -1 )
				h5 := regexp.MustCompile("<h5(.*?)>")
				info.Detail = h5.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</h5>", "", -1 )
				h6 := regexp.MustCompile("<h6(.*?)>")
				info.Detail = h6.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</h6>", "", -1 )
				i := regexp.MustCompile("<i(.*?)>")
				info.Detail = i.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</i>", "", -1 )
				u := regexp.MustCompile("<u(.*?)>")
				info.Detail = u.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</u>", "", -1 )
				b := regexp.MustCompile("<b(.*?)>")
				info.Detail = b.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</b>", "", -1 )
				sup := regexp.MustCompile("<sup(.*?)>")
				info.Detail = sup.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</sup>", "", -1 )
				small := regexp.MustCompile("<small(.*?)>")
				info.Detail = small.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</small>", "", -1 )
				crosslinking := regexp.MustCompile("<cross-linking(.*?)>")
				info.Detail = crosslinking.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</cross-linking>", "", -1 )
				center := regexp.MustCompile("<center(.*?)>")
				info.Detail = center.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</center>", "", -1 )
				style := regexp.MustCompile(`<style[\s\S]*</style>`)
				info.Detail = style.ReplaceAllString( info.Detail, "" )
				script := regexp.MustCompile(`<script[\s\S]*</script>`)
				info.Detail = script.ReplaceAllString( info.Detail, "" )
				table := regexp.MustCompile(`<table[\s\S]*</table>`)
				info.Detail = table.ReplaceAllString( info.Detail, "" )
				TABLE := regexp.MustCompile(`<TABLE[\s\S]*</TABLE>`)
				info.Detail = TABLE.ReplaceAllString( info.Detail, "" )
				xml := regexp.MustCompile(`<\?xml[\s\S]*?>`)
				info.Detail = xml.ReplaceAllString( info.Detail, "" )
				tag := regexp.MustCompile(`<!-- Start of Double\s*Click Spotlight Tag[\s\S]*Please do not remove-->`)
				info.Detail = tag.ReplaceAllString( info.Detail, "" )
				A := regexp.MustCompile("<A(.*?)>")
				info.Detail = A.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</A>", "", -1 )
				P := regexp.MustCompile("<P(.*?)>")
				info.Detail = P.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</P>", "", -1 )
				UL := regexp.MustCompile("<UL(.*?)>")
				info.Detail = UL.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</UL>", "", -1 )
				LI := regexp.MustCompile("<LI(.*?)>")
				info.Detail = LI.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</LI>", "", -1 )
				DIV := regexp.MustCompile("<DIV(.*?)>")
				info.Detail = DIV.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</DIV>", "", -1 )
				SPAN := regexp.MustCompile("<SPAN(.*?)>")
				info.Detail = SPAN.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</SPAN>", "", -1 )
				info.Detail = strings.Replace( info.Detail, "</SPAN >", "", -1 )
				FONT := regexp.MustCompile("<FONT(.*?)>")
				info.Detail = FONT.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</FONT>", "", -1 )
				STRONG := regexp.MustCompile("<STRONG(.*?)>")
				info.Detail = STRONG.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</STRONG>", "", -1 )
				EM := regexp.MustCompile("<EM(.*?)>")
				info.Detail = EM.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</EM>", "", -1 )
				H1 := regexp.MustCompile("<H1(.*?)>")
				info.Detail = H1.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</H1>", "", -1 )
				H2 := regexp.MustCompile("<H2(.*?)>")
				info.Detail = H2.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</H2>", "", -1 )
				H3 := regexp.MustCompile("<H3(.*?)>")
				info.Detail = H3.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</H3>", "", -1 )
				H4 := regexp.MustCompile("<H4(.*?)>")
				info.Detail = H4.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</H4>", "", -1 )
				H5 := regexp.MustCompile("<H5(.*?)>")
				info.Detail = H5.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</H5>", "", -1 )
				H6 := regexp.MustCompile("<H6(.*?)>")
				info.Detail = H6.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</H6>", "", -1 )
				I := regexp.MustCompile("<I(.*?)>")
				info.Detail = I.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</I>", "", -1 )
				U := regexp.MustCompile("<U(.*?)>")
				info.Detail = U.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</U>", "", -1 )
				B := regexp.MustCompile("<B(.*?)>")
				info.Detail = B.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</B>", "", -1 )
				SUP := regexp.MustCompile("<SUP(.*?)>")
				info.Detail = SUP.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</SUP>", "", -1 )
				SMALL := regexp.MustCompile("<SMALL(.*?)>")
				info.Detail = SMALL.ReplaceAllString( info.Detail, "" )
				info.Detail = strings.Replace( info.Detail, "</SMALL>", "", -1 )
				info.Detail = strings.Replace( info.Detail, "<BR>", " ", -1 )
				info.Detail = strings.Replace( info.Detail, "<BR/>", " ", -1 )
				info.Detail = strings.Replace( info.Detail, "</BR>", " ", -1 )
				info.Detail = strings.Replace( info.Detail, "<br>", " ", -1 )
				info.Detail = strings.Replace( info.Detail, "</br>", " ", -1 )
				info.Detail = strings.Replace( info.Detail, "</divmain_image>", "", -1 )
				info.Detail = strings.Replace( info.Detail, "</div main_image>", "", -1 )
				info.Detail = strings.Replace( info.Detail, "<!--StartFragment -->", "", -1 )
				info.Detail = strings.Replace( info.Detail, "<o:p></o:p>", "", -1 )
				info.Detail = strings.Replace( info.Detail, "<!--[if !IE]>-->]]><![endif]-->", "", -1 )
				info.Detail = strings.Replace( info.Detail, "<!--[if !IE]>--><![CDATA[<![endif]-->", "", -1 )
				info.Detail = strings.Replace( info.Detail, "<!--[if !IE]>--><![endif]-->", "", -1 )
				info.Detail = strings.Replace( info.Detail, "<br/>", " ", -1 )
				info.Detail = strings.Replace( info.Detail, "<heading>", " ", -1 )
				info.Detail = strings.Replace( info.Detail, "</blockquote>", " ", -1 )
				info.Detail = strings.Replace( info.Detail, "</iframe>", " ", -1 )
				info.Detail = strings.Replace( info.Detail, "<Li>", " ", -1 )
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

			date := myregex.Parse( htmlFile, patternForDate )
			if len(date) >= 1 {
				info.Date = date[0][1]
			}

			developer := myregex.Parse( htmlFile, patternForDeveloper )
			if len(developer) >= 1 {
				for _, content := range developer{
					info.Developer = append(info.Developer, content[1])
				}
			}

			publisher := myregex.Parse( htmlFile, patternForPublisher )
			if len(publisher) >= 1 {
				for _, content := range publisher{
					info.Publisher = append(info.Publisher, content[1])
				}
			}

			genre := myregex.Parse( htmlFile, patternForGenre )
			if len(genre) >= 1 {
				for _, content := range genre{
					info.Genre = append(info.Genre, content[1])
				}
			}

			image := myregex.Parse( htmlFile, patternForImage )
			if len(image) >= 1 {
				for _, content := range image{
					info.Image = append(info.Image, content[1])
				}
			}

			video := myregex.Parse( htmlFile, patternForVideo )
			if len(video) >= 1 {
				for _, content := range video{
					info.Video = append(info.Video, content[1])
				}
			}

			info.Id = urlSplit[ len(urlSplit)-2 ]

			platform := myregex.Parse( game[0], patternForPlatform )
			if len(platform) >= 1 {
				for _, content := range platform{
					info.Platform = append(info.Platform, content[1])
				}
			}

			j,err := json.Marshal( &info )
			if err != nil {
				fmt.Println(err)
			}
			//fmt.Printf( "json: %s\n", string(j) )

			myfile.SaveFile( writePath + "/" + urlSplit[ len(urlSplit)-2 ] + ".json", []byte(j) )
		}
	}
}