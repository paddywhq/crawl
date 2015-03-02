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
const patternForResult = `<a href="(.*?)"  data-ds-appid=".*?" onmouseover=".*?" onmouseout=".*?" class="search_result_row ds_collapse_flag" >`

func main(){
	files := myfile.GetFileList( readPath )

	c := ""

	for _, f := range files{
		fmt.Printf("\nfile: %s\n", f)

		l := myfile.GetFileContent( readPath + "/" + f )

		result := myregex.Parse( l, patternForResult )

		for _, game := range result{
			url := game[1]
			c += url + "\n"
		}
	}

	myfile.SaveFile( writePath, []byte(c) )
}