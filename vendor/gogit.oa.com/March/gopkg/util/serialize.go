package util

import (
	"fmt"
	"strconv"
	"strings"
)

// money:44406411,msavecount:0,exp:160,sub:0,mbank:0,cas:214,tm:1570781757.4382,djmoney:0,tblmoney:0,bycoins:0,mid:223326
func UnSerialize(src string) map[string]string {
	sl := strings.Split(src, ",")
	res := make(map[string]string, 0)
	for _, v := range sl {
		vsl := strings.Split(v, ":")
		if len(vsl) == 2 {
			res[vsl[0]] = vsl[1]
		}
	}
	return res
}

func Serialize(data map[string]interface{}) string {
	var res []string
	for k, v := range data {
		switch v.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			s := fmt.Sprintf("%s:%d", k, v)
			res = append(res, s)
		case float32:
			s := fmt.Sprintf("%s:%s", k, strconv.FormatFloat(float64(v.(float32)), 'f', -1, 32))
			res = append(res, s)
		case float64:
			s := fmt.Sprintf("%s:%s", k, strconv.FormatFloat(v.(float64), 'f', -1, 64))
			res = append(res, s)
		case string:
			s := fmt.Sprintf("%s:%s", k, v)
			res = append(res, s)
		}
	}

	return strings.Join(res, ",")
}
