//全局配置
package conf

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

var Conf [][]string

func init() {
	dir, _ := os.Getwd()
	f, err := os.Open(dir + "/plays/conf/conf.csv")
	if err != nil {
		log.Println("conf", err)
	}
	r := csv.NewReader(f)
	str, err := r.ReadAll()
	if err != nil {
		log.Println("conf", err)
	}
	Conf = str
}

func GetConfAll() ([][]string) {
	return Conf
}

func GetConfInt(name string) (value int) {
	for _, v := range Conf {
		if v[1] == name {
			value, _ = strconv.Atoi(v[2])
			return
		}
	}
	return 0
}

func GetConfString(name string) (value string) {

	return ""
}
