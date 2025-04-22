package main

import (
	// "fmt"
	"bufio"
	"fmt"
	"os"
)

func main() {
  	// f,err:=os.Open("file.txt")
	// if err!=nil{
	// 	//log error
	// 	panic(err)
	// }
	// fileInfo,err := f.Stat()
	// f,err:= os.Open("file.txt")
	// if err!= nil{
	// 	panic (err)
	// }
	// defer f.Close()

	// buf := make([]byte, 10)
	// d,err:= f.Read(buf)
	// if err!=nil{ panic(err)
	// }
	// for i:=0 ;i<len(buf);i++{
	// fmt.Println("Data",d, string(buf[i]))
	// }
	// data,err:= os.ReadFile("file.txt")
	// if err!=nil {
	// 	panic(err)
	// }
	// fmt.Println(string (data))

	//read folders
	// dir,err:=os.Open("./")
	// if err!=nil{
	// 	panic(err)
	// }
	// defer dir.Close()
	// fileInfo,err:=dir.ReadDir(-1)
	// for _, fi :=range  fileInfo{
	// 	fmt.Println(fi.Name(),fi.IsDir())
	// }
	// f,err:=os.Create("sex.txt")
	// if err !=nil {
	// 	panic(err)
	// }
	// defer f.Close() 

	// bytes:= []byte("hello gokang")
	// f.Write(bytes)

	//tranfer file data to other file 
	//in streaming fashion

	sourceFile,err:= os.Open("sex.txt")
	if err != nil{
		panic(err)
	}
	defer sourceFile.Close()
	destFile,err:= os.Create("exm.txt")
	if err!=nil{
		panic(err)
	}
	defer destFile.Close()
	reader:= bufio.NewReader(sourceFile)
	writer:= bufio.NewWriter(destFile)

	for {
		b,err:=reader.ReadByte()
		if err !=nil{
			if err.Error()!= "EOF"{
				panic(err)
			}
			break
		}
		er:=writer.WriteByte(b)
		if er!=nil{
			panic(er)
		}
	}
	writer.Flush()
	fmt.Println("successfully return to new file")
}