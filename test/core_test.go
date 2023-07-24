/**
 @author : ikkk
 @date   : 2023/7/19
 @text   : nil
**/

package test

import (
	"car_pooling_admini/tools/moudle"
	"fmt"
	"github.com/goccy/go-json"
	"testing"
	"time"
)

func TestCore(T *testing.T) {
	var sysUsers []moudle.SysUser
	re := moudle.DB.Find(&sysUsers)
	fmt.Println(re.Error)
	fmt.Println(sysUsers)
}

func TestTime(t *testing.T) {
	re := time.Now()
	fmt.Println(re.Format("2006-01-02 15:04:05"))
}

type Student struct {
	Name string `json:"name"`
}

func (s *Student) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func TestLink(T *testing.T) {
	//cli.Set(ctx, "demo", "key1", 0)
	res, err := moudle.Cli.Get(moudle.Ctx, "a").Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
