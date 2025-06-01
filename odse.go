package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

const version = "v1.0.0"

func printVersion() {
	fmt.Println("Version:", version)
}

func checkServiceStatus(service string) {
	cmd := exec.Command("systemctl", "is-active", service)
	output, err := cmd.Output()
	if err != nil {
		log.Printf("서비스 상태 확인 실패: %v\n", err)
		return
	}
	status := strings.TrimSpace(string(output))
	log.Printf("서비스 [%s] 상태: %s\n", service, status)
}

func sendWallMessage() {
	numbers := make([]int, 6)
	for i := 0; i < 6; i++ {
		numbers[i] = rand.Intn(100)
	}

	message := fmt.Sprintf("Hello~ 랜덤 숫자: %v", numbers)
	cmd := exec.Command("wall", message)
	err := cmd.Run()
	if err != nil {
		log.Printf("wall 메시지 전송 실패: %v\n", err)
	}
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--version" {
		printVersion()
		return
	}

	log.Println("프로그램 실행 시작")
	rand.Seed(time.Now().UnixNano())

	for {
		checkServiceStatus("testpg")
		sendWallMessage()
		time.Sleep(60 * time.Second) // 1분마다 실행
	}
}
