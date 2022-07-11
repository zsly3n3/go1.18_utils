package generate

import (
	"fmt"
	"github.com/sony/sonyflake"
	"github.com/zsly3n3/go1.18_utils/common"
	"time"
)

var sf = sonyflake.NewSonyflake(sonyflake.Settings{
	StartTime: time.Date(2020, 12, 16, 0, 0, 0, 0, time.Local)})

//非分布式唯一
//func Unique() (string, error) {
//	id, err := sf.NextID()
//	sid := fmt.Sprint(id)
//	return sid, err
//}

//分布式唯一ID:md5(雪花ID和进程ID)
//传入进程ID
func ClusterUniqueID(pid int) (string, error) {
	id, err := sf.NextID()
	if err != nil {
		return "", err
	}
	rs := common.Md5(fmt.Sprintf(`%d%d`, id, pid))
	return rs, nil
}
