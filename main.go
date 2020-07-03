package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	// "net"
	"os"
	// "regexp"
	// "strconv"
	// "strings"
	"sync"
)
var  ipAddr string
var startPort,endPort,timeout int 
var showDetailInfo bool
var logfile *os.File
func main(){
	
    flag.StringVar(&ipAddr,"a","","ip地址")
	flag.IntVar(&startPort,"s",0,"起始端口")
	flag.IntVar(&endPort,"e",65535,"结束端口")

	flag.IntVar(&timeout,"t",1000,"超时时间毫秒")
	flag.BoolVar(&showDetailInfo,"i",false,"显示详细扫描信息")
	flag.Parse()

  startTime:=time.Now().UnixNano()
	logfile, _ := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	log.SetOutput(logfile)

	var wg sync.WaitGroup
	fmt.Printf("正在扫描IP地址%s  端口号%d~%d ;超时时间%d毫秒\r\n",ipAddr,startPort,endPort,timeout)
	for{
	  for i := 0; i < 5; i++ {		  
		   if startPort>endPort{
			ts:= time.Now().UnixNano()	
			seconds:= float64((ts - startTime) / 1e9)
			fmt.Printf("扫描完成-耗时：%f 秒 ，扫描结果查看日志内容",seconds)
			logfile.Close()
			return
		   }
		   fullAddr:=fmt.Sprintf("%s:%d",ipAddr,startPort)
		   startPort++;
		   wg.Add(1)
		   go scanPort(fullAddr,&wg)
	  }
	  wg.Wait()
	
	}

	

}

func scanPort( tcpAddress string,wg *sync.WaitGroup){
	
	_,err:= net.DialTimeout("tcp",tcpAddress,time.Millisecond*time.Duration(timeout))
	
	if err!=nil{
		if showDetailInfo{
			fmt.Printf("%s 未开放\r\n",tcpAddress)
		}
		
	}else{	
		log.Printf("扫描到一个开放端口：%s \r\n",tcpAddress)
	}
	wg.Done()
	
}