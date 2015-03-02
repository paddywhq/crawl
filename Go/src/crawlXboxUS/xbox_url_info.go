package main

import(
	"fmt"
	"./myfile"
	"./myregex"
	// "strconv"
	// "os"
	// "encoding/json"
	"strings"
)

const readPath = "./gamelink_html"
const writePath = "./gameUrl.txt"
const urlXbox360Base = "http://marketplace.xbox.com"
const urlXbox1Base = "http://store.xbox.com"
const patternForXbox360Result = `<h2><a href="(.*?)" title=".*?">.*?</a></h2>`
const patternForXbox1Result = `<a Title=".*?" data-alt=".*?" data-contentid=".*?" data-slot=".*?" href="(.*?)">.*?</a>`

func main(){
	files := myfile.GetFileList( readPath )

	c := ""

	for _, f := range files{
		fmt.Printf("\nfile: %s\n", f)

		l := myfile.GetFileContent( readPath + "/" + f )

		if strings.Contains(f, "360_") {
			result := myregex.Parse( l, patternForXbox360Result )
			for _, game := range result{
				url := game[1]
				c += urlXbox360Base + url + "?nosplash=1\n"
			}
		} else {
			result := myregex.Parse( l, patternForXbox1Result )
			for _, game := range result{
				url := game[1]
				c += urlXbox1Base + url + "?nosplash=1\n"
			}
		}
	}

	myfile.SaveFile( writePath, []byte(c) )
}