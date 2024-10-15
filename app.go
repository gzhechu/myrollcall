package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/exp/rand"
)

// App struct
type App struct {
	ctx   context.Context
	Names []Person
}

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func ShuffleSlice[T any](slice []T) []T {
	n := len(slice)
	rand.Seed(uint64(time.Now().UnixNano())) // 使用当前时间作为随机种子

	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)                   // 生成 [0, i] 范围内的随机索引
		slice[i], slice[j] = slice[j], slice[i] // 交换元素
	}
	return slice
}

func (a *App) GetNameList(name string) []Person {
	fmt.Println("get name list.")
	//	a.LoadNames()
	return ShuffleSlice(a.Names)
}

func (a *App) LoadFile() {
	opt := runtime.OpenDialogOptions{
		Title: "选择名单文件",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Text File (*.txt;*.csv)",
				Pattern:     "*.txt;*.csv",
			},
		},
	}
	filePath, err := runtime.OpenFileDialog(a.ctx, opt)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	a.loadNamesFronFile(filePath)
}

func (a *App) loadNamesFronFile(f string) {
	// 打开文件
	// file, err := os.Open("22软工1班.txt")
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	a.Names = make([]Person, 0)

	// 创建扫描器逐行读取文件
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// 解析每一行
		parts := strings.Split(line, "\t")
		if len(parts) != 2 {
			log.Println("Invalid line format:", line)
			continue
		}

		name := parts[1]
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Println("Invalid age format:", parts[1])
			continue
		}

		// 创建 Person 对象并添加到列表
		person := Person{Name: name, Id: id}
		a.Names = append(a.Names, person)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// 输出结果
	fmt.Println(a.Names)
	fmt.Println(len(a.Names))
}
