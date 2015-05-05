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

const readPathJson = "./gamelink_json"
const readPathHtml = "./gameinfo_html"
const writePath = "./gameinfo_json"
const patternForResult = `"index"\s*:\s*"(\d*)",\s*"title"\s*:\s*"(.*)",\s*"description"\s*:\s*"(.*)",\s*"context"\s*:\s*"(.*)",\s*"content_type"\s*:\s*"(.*)",\s*"feature"\s*:\s*"(.*)",\s*"keywords"\s*:\s*"(.*)",\s*"platform"\s*:\s*"(.*)",\s*"genre"\s*:\s*"(.*)",\s*"rating"\s*:\s*"(.*)",\s*"image"\s*:\s*"(.*)",\s*"video_url"\s*:\s*"(.*)",\s*"video_desc"\s*:\s*"(.*)",\s*"image_url"\s*:\s*"(.*)",\s*"image_desc"\s*:\s*"(.*)",\s*"release_date"\s*:\s*"(.*)",\s*"release_date_display"\s*:\s*"(.*)",\s*"url"\s*:\s*"(.*)",\s*"psnId"\s*:\s*"(.*)",\s*"buttonTxt"\s*:\s*"(.*)",\s*"buttonURL"\s*:\s*"(.*)",\s*"buttonColor"\s*:\s*"(.*)",\s*"ageRatingImage"\s*:\s*"(.*)",\s*"ageRatingImageAlt"\s*:\s*"(.*)",`
const patternForDetail1 = `<div class="richtext copy\s*?">([\s\S]*?)</div>`
const patternForDetail2 = `<!-- Text block - CM168 -->([\s\S]*?)<!-- End Text block - CM168 -->`
const patternForPublisher = `<li>\s*?<span>\s*?Publisher:\s*?</span>\s*(.*?)\s*?</li>`
const patternForDeveloper = `<li>\s*?<span>\s*?Developer:\s*?</span>\s*(.*?)\s*?</li>`

func main(){

	os.Mkdir( writePath, os.ModeSticky | 0755)

	files := myfile.GetFileList( readPathJson )
	for _, f := range files{
		fmt.Printf("\nfile: %s\n", f)

		jsonfile := myfile.GetFileContent( readPathJson + "/" + f )

		//有些简介中有br
		// jsonfile = strings.Replace( jsonfile, "<br />", "\n", -1)

		result := myregex.Parse( jsonfile, patternForResult )
		
		for _, info := range result{
			game := GameInfo{}

			game.Index = info[1]
			game.Title = info[2]
			game.Description = info[3]
			game.Context = info[4]
			game.Content_type = info[5]
			game.Feature = info[6]
			game.Keywords = info[7]
			game.Platform = info[8]
			game.Genre = info[9]
			// game.Rating = info[10]
			game.Image = info[11]
			game.Video_url = info[12]
			// game.Video_desc = info[13]
			game.Image_url = info[14]
			// game.Image_desc = info[15]
			game.Release_date = info[16]
			game.Release_date_display = info[17]
			game.Url = info[18]
			game.PsnId = info[19]
			// game.ButtonTxt = info[20]
			// game.ButtonURL = info[21]
			// game.ButtonColor = info[22]
			game.AgeRatingImage = info[23]
			game.AgeRatingImageAlt = info[24]

			urlSplit := strings.Split( game.Url, "/" )
			filePath := readPathHtml + "/" + urlSplit[ len(urlSplit)-2 ] + ".html"

			htmlfile := myfile.GetFileContent( filePath )

			//有些简介中有br
			// htmlfile = strings.Replace( htmlfile, "<br />", "\n", -1)

			details1 := myregex.Parse( htmlfile, patternForDetail1 )
			details2 := myregex.Parse( htmlfile, patternForDetail2 )
			publisher := myregex.Parse( htmlfile, patternForPublisher )
			developer := myregex.Parse( htmlfile, patternForDeveloper )

			c := ""
			for _, detail := range details1{
				c += detail[1] + "\n"
			}
			for _, detail := range details2{
				c += detail[1] + "\n"
			}
			
			game.Detail = c
			game.Publisher = publisher[0][1]
			game.Developer = developer[0][1]

			a := regexp.MustCompile("<a(.*?)>")
			game.Detail = a.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</a>", "", -1 )
			p := regexp.MustCompile("<p(.*?)>")
			game.Detail = p.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</p>", "", -1 )
			ul := regexp.MustCompile("<ul(.*?)>")
			game.Detail = ul.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</ul>", "", -1 )
			li := regexp.MustCompile("<li(.*?)>")
			game.Detail = li.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</li>", "", -1 )
			div := regexp.MustCompile("<div(.*?)>")
			game.Detail = div.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</div>", "", -1 )
			game.Detail = strings.Replace( game.Detail, "</div >", "", -1 )
			span := regexp.MustCompile("<span(.*?)>")
			game.Detail = span.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</span>", "", -1 )
			game.Detail = strings.Replace( game.Detail, "</span >", "", -1 )
			font := regexp.MustCompile("<font(.*?)>")
			game.Detail = font.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</font>", "", -1 )
			strong := regexp.MustCompile("<strong(.*?)>")
			game.Detail = strong.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</strong>", "", -1 )
			em := regexp.MustCompile("<em(.*?)>")
			game.Detail = em.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</em>", "", -1 )
			h1 := regexp.MustCompile("<h1(.*?)>")
			game.Detail = h1.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</h1>", "", -1 )
			h2 := regexp.MustCompile("<h2(.*?)>")
			game.Detail = h2.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</h2>", "", -1 )
			h3 := regexp.MustCompile("<h3(.*?)>")
			game.Detail = h3.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</h3>", "", -1 )
			h4 := regexp.MustCompile("<h4(.*?)>")
			game.Detail = h4.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</h4>", "", -1 )
			h5 := regexp.MustCompile("<h5(.*?)>")
			game.Detail = h5.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</h5>", "", -1 )
			h6 := regexp.MustCompile("<h6(.*?)>")
			game.Detail = h6.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</h6>", "", -1 )
			i := regexp.MustCompile("<i(.*?)>")
			game.Detail = i.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</i>", "", -1 )
			u := regexp.MustCompile("<u(.*?)>")
			game.Detail = u.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</u>", "", -1 )
			b := regexp.MustCompile("<b(.*?)>")
			game.Detail = b.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</b>", "", -1 )
			sup := regexp.MustCompile("<sup(.*?)>")
			game.Detail = sup.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</sup>", "", -1 )
			small := regexp.MustCompile("<small(.*?)>")
			game.Detail = small.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</small>", "", -1 )
			crosslinking := regexp.MustCompile("<cross-linking(.*?)>")
			game.Detail = crosslinking.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</cross-linking>", "", -1 )
			center := regexp.MustCompile("<center(.*?)>")
			game.Detail = center.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</center>", "", -1 )
			style := regexp.MustCompile(`<style[\s\S]*</style>`)
			game.Detail = style.ReplaceAllString( game.Detail, "" )
			script := regexp.MustCompile(`<script[\s\S]*</script>`)
			game.Detail = script.ReplaceAllString( game.Detail, "" )
			table := regexp.MustCompile(`<table[\s\S]*</table>`)
			game.Detail = table.ReplaceAllString( game.Detail, "" )
			TABLE := regexp.MustCompile(`<TABLE[\s\S]*</TABLE>`)
			game.Detail = TABLE.ReplaceAllString( game.Detail, "" )
			xml := regexp.MustCompile(`<\?xml[\s\S]*?>`)
			game.Detail = xml.ReplaceAllString( game.Detail, "" )
			tag := regexp.MustCompile(`<!-- Start of Double\s*Click Spotlight Tag[\s\S]*Please do not remove-->`)
			game.Detail = tag.ReplaceAllString( game.Detail, "" )
			A := regexp.MustCompile("<A(.*?)>")
			game.Detail = A.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</A>", "", -1 )
			P := regexp.MustCompile("<P(.*?)>")
			game.Detail = P.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</P>", "", -1 )
			UL := regexp.MustCompile("<UL(.*?)>")
			game.Detail = UL.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</UL>", "", -1 )
			LI := regexp.MustCompile("<LI(.*?)>")
			game.Detail = LI.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</LI>", "", -1 )
			DIV := regexp.MustCompile("<DIV(.*?)>")
			game.Detail = DIV.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</DIV>", "", -1 )
			SPAN := regexp.MustCompile("<SPAN(.*?)>")
			game.Detail = SPAN.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</SPAN>", "", -1 )
			game.Detail = strings.Replace( game.Detail, "</SPAN >", "", -1 )
			FONT := regexp.MustCompile("<FONT(.*?)>")
			game.Detail = FONT.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</FONT>", "", -1 )
			STRONG := regexp.MustCompile("<STRONG(.*?)>")
			game.Detail = STRONG.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</STRONG>", "", -1 )
			EM := regexp.MustCompile("<EM(.*?)>")
			game.Detail = EM.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</EM>", "", -1 )
			H1 := regexp.MustCompile("<H1(.*?)>")
			game.Detail = H1.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</H1>", "", -1 )
			H2 := regexp.MustCompile("<H2(.*?)>")
			game.Detail = H2.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</H2>", "", -1 )
			H3 := regexp.MustCompile("<H3(.*?)>")
			game.Detail = H3.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</H3>", "", -1 )
			H4 := regexp.MustCompile("<H4(.*?)>")
			game.Detail = H4.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</H4>", "", -1 )
			H5 := regexp.MustCompile("<H5(.*?)>")
			game.Detail = H5.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</H5>", "", -1 )
			H6 := regexp.MustCompile("<H6(.*?)>")
			game.Detail = H6.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</H6>", "", -1 )
			I := regexp.MustCompile("<I(.*?)>")
			game.Detail = I.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</I>", "", -1 )
			U := regexp.MustCompile("<U(.*?)>")
			game.Detail = U.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</U>", "", -1 )
			B := regexp.MustCompile("<B(.*?)>")
			game.Detail = B.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</B>", "", -1 )
			SUP := regexp.MustCompile("<SUP(.*?)>")
			game.Detail = SUP.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</SUP>", "", -1 )
			SMALL := regexp.MustCompile("<SMALL(.*?)>")
			game.Detail = SMALL.ReplaceAllString( game.Detail, "" )
			game.Detail = strings.Replace( game.Detail, "</SMALL>", "", -1 )
			game.Detail = strings.Replace( game.Detail, "<BR>", " ", -1 )
			game.Detail = strings.Replace( game.Detail, "<BR/>", " ", -1 )
			game.Detail = strings.Replace( game.Detail, "</BR>", " ", -1 )
			game.Detail = strings.Replace( game.Detail, "<br>", " ", -1 )
			game.Detail = strings.Replace( game.Detail, "</br>", " ", -1 )
			game.Detail = strings.Replace( game.Detail, "</divmain_image>", "", -1 )
			game.Detail = strings.Replace( game.Detail, "</div main_image>", "", -1 )
			game.Detail = strings.Replace( game.Detail, "<!--StartFragment -->", "", -1 )
			game.Detail = strings.Replace( game.Detail, "<o:p></o:p>", "", -1 )
			game.Detail = strings.Replace( game.Detail, "<!--[if !IE]>-->]]><![endif]-->", "", -1 )
			game.Detail = strings.Replace( game.Detail, "<!--[if !IE]>--><![CDATA[<![endif]-->", "", -1 )
			game.Detail = strings.Replace( game.Detail, "<!--[if !IE]>--><![endif]-->", "", -1 )
			game.Detail = strings.Replace( game.Detail, "<br/>", " ", -1 )
			game.Detail = strings.Replace( game.Detail, "<heading>", " ", -1 )
			game.Detail = strings.Replace( game.Detail, "\n", " ", -1 )
			game.Detail = strings.Replace( game.Detail, "\r", " ", -1 )
			game.Detail = strings.Replace( game.Detail, "\t", " ", -1 )
			game.Detail = strings.Replace( game.Detail, "\f", "", -1 )
			game.Detail = strings.Replace( game.Detail, "\v", "", -1 )
			game.Detail = strings.Replace( game.Detail, "\\\"", "\"", -1 )
			game.Detail = strings.Replace( game.Detail, "\"", "\"\"", -1 )
			space := regexp.MustCompile(`\s+`)
			game.Detail = space.ReplaceAllString( game.Detail, ` ` )
			game.Description = strings.Replace( game.Description, "\n", " ", -1 )
			game.Description = strings.Replace( game.Description, "\r", " ", -1 )
			game.Description = strings.Replace( game.Description, "\\n", " ", -1 )
			game.Description = strings.Replace( game.Description, "\\r", " ", -1 )
			game.Description = strings.Replace( game.Description, "\\\"", "\"", -1 )
			game.Description = strings.Replace( game.Description, "\"", "\"\"", -1 )
			game.Description = space.ReplaceAllString( game.Description, ` ` )

			j,err := json.Marshal( &game )
			if err != nil {
				fmt.Println(err)
			}
			//fmt.Printf( "json: %s\n", string(j) )

			myfile.SaveFile( writePath + "/" + urlSplit[ len(urlSplit)-2 ] + ".json", []byte(j) )
		}
	}
}