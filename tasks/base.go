package tasks

import (
	"Helloc/models/utils"
	"fmt"
	"github.com/jasonlvhit/gocron"
	"log"
	"os"
	"time"
)

var scheduler = gocron.NewScheduler()
var sc chan bool

func StartTasks() {
	scheduler.Every(3600).Second().Do(defaultPrintTask)
	scheduler.Every(3600).Second().Do(func (){
		f, err := os.OpenFile("task.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer f.Close()
		n, err := f.WriteString("今天天气真好!\n")
		fmt.Println(n, err)

	})

	sc = scheduler.Start()
	<-sc
	fmt.Println("sc close")
}

func Close() {
	fmt.Println("sc close")
	close(sc)
}

func LoopTask() {
	s := gocron.NewScheduler()
	s.Every(2).Second().Do(defaultPrintTask)
	sc := s.Start()
	go func() {
		for {
			if value, _ := utils.RGetString("close"); value!=""{
				close(sc)
			}
			time.Sleep(time.Second)
		}
	}()
	fmt.Println(233333)
	<-sc // 等待子go通知结束
}

func writeTask() {
	f, err := os.OpenFile("/home/youcan/Desktop/Go/works/Helloc/ab.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	fmt.Println(err)
	defer f.Close()
	fmt.Println("今天天气真好!")
	n, err := f.WriteString("今天天气真好!")
	fmt.Println(n, err)

}

func PrintTask() {
	s := gocron.NewScheduler()
	s.Every(1).Seconds().Do(printTask, "printTask")
	s.Every(4).Seconds().Do(defaultPrintTask)
	fmt.Println("test")
	sc := s.Start() // keep the channel
	time.Sleep(8 * time.Second)
	s.Remove(printTask)
	fmt.Println("remove printTask")
	time.Sleep(8 * time.Second)
	s.Clear()
	fmt.Println("clear")
	close(sc)
}

func printTask(s string) {
	for i:=0; i< 2;i++ {
		fmt.Println(s)
	}
}

func defaultPrintTask() {
	for i:=0; i< 2;i++ {
		fmt.Println("defaultPrintTask")
	}
}




