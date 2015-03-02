package main

import(
	"fmt"
	"./myfile"
	"./myregex"
	// "strconv"
	// "os"
	// "encoding/json"
	// "strings"
)

const readPath = "./gamelink_html"
const writePath = "./gameUrl.txt"
const urlBase = "http://psnprofiles.com"
const patternForResult = `<td\s*style\s*=\s*"width:\s*100%;"\s*>\s*<a\s*class="bold"\s*href="(.*?)"\s*>`

func main(){
	files := myfile.GetFileList( readPath )

	c := ""

	for _, f := range files{
		fmt.Printf("\nfile: %s\n", f)

		l := myfile.GetFileContent( readPath + "/" + f )

		result := myregex.Parse( l, patternForResult )

		for _, game := range result{
			url := game[1]
			c += urlBase + url + "\n"
		}
	}

	myfile.SaveFile( writePath, []byte(c) )
}