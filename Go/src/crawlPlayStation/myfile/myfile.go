package myfile

import (
	"io/ioutil" 
	"bufio"
	"os"
	"fmt"
)

func GetFileLines( path string ) []string{
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("open file rror:%v\n", err)
	}
	defer file.Close()

	lines := make([]string, 0, 10000)
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}
	return lines
}

func GetFileContent( path string ) string{
	data, _ := ioutil.ReadFile(path)
	return string(data)
}

func SaveFile( path string, content []byte ) error {
	 err := ioutil.WriteFile( path, content, 0644)
	 return err
}

func FileExists( path string ) bool{
	if _, err:= os.Stat(path); os.IsNotExist(err){
		return false
	}
	return true
}

// func GetFileSize( path string ) int {
// 	f, err := file.Open( path )
// 	if err != nil {
// 		fmt.Printf( "error open file: %v", err )
// 	}

// 	fi, err := f.Stat()
// 	if err != nil{
// 		fmt.Printf("error read status of file: %v", err)
// 	}

// 	return fi.Size()
// }

func GetFileList( path string) []string{
	fileInfo, _ := ioutil.ReadDir( path )

	size := len(fileInfo)
	paths := make( []string, size, size )
	i := 0
    for _, f := range fileInfo {
        paths[i] = f.Name()
        i++
    }
    return paths
}
