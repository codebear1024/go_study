package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

// 写一个程序，解析配置文件xxx.ini
type configMysql struct {
	Username string `config:"username"`
	Password string `config:"password"`
	Addr     string `config:"addr"`
	Port     uint32 `config:"port"`
}

type configRedis struct {
	Host string `config:"host"`
	Port uint32 `config:"port"`
}

const (
	configTypeUnknown = iota
	configTypeMysql
	configTypeRedis
)

var mysql configMysql
var redis configRedis

func main() {
	var configType int
	// 打开文件
	confile, err := os.Open("./xxx.ini")
	if err != nil {
		fmt.Println("open", err)
		return
	}
	defer confile.Close()
	// 一行一行的读文件
	reader := bufio.NewReader(confile)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if line == "\r\n" {
			continue
		}
		// 用strings.TrimSpace()删除开头和结尾的空格
		//line = strings.TrimSpace(line)
		// 用string。ReplaceAll()删除中间的空格
		line = strings.ReplaceAll(line, " ", "")
		// 解析每一行读出来的文件
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if line == "[mysql]" {
			configType = configTypeMysql
			continue
		} else if line == "[redis]" {
			configType = configTypeRedis
			continue
		}
		if !unicode.IsLetter(rune(line[0])) {
			continue
		}
		data := strings.Split(line, "=")
		if len(data) != 2 {
			continue
		}
		parseData(data[0], data[1], configType)
	}
	fmt.Println(mysql)
	fmt.Println(redis)
}

func parseData(configKey, configValue string, configType int) (err error) {
	var v reflect.Value
	if configType == configTypeMysql {
		v = reflect.ValueOf(&mysql).Elem() // 下面会修改结构体的值，所以要使用取地址和Eleme()
	} else if configType == configTypeRedis {
		v = reflect.ValueOf(&redis).Elem()
	} else {
		err = fmt.Errorf("unknow config type")
		return
	}
	for i := 0; i < v.NumField(); i++ { // 这里也可以使用t.NumField()进行遍历，t = reflect.TypeOf()
		filed := v.Field(i) //v.Field(i)表示结构体第个字段的值
		if !filed.IsValid() || !filed.CanSet() {
			continue
		}
		if v.Type().Field(i).Tag.Get("config") == configKey { //v.Type()等同于reflect.Type()获取的类型，是两种不同的获取类型的方法
			switch filed.Kind() {
			case reflect.String:
				filed.SetString(configValue)
			case reflect.Uint32:
				num, err := strconv.Atoi(configValue)
				if err == nil {
					filed.SetUint(uint64(num))
				}
			}
		}
	}
	return
}
