package generate

import (
	"fmt"
	"github.com/sony/sonyflake"
	"github.com/zsly3n3/go1.18_utils/common"
	"time"
)

var sf = sonyflake.NewSonyflake(sonyflake.Settings{
	StartTime: time.Date(2020, 12, 16, 0, 0, 0, 0, time.Local)})

func Unique() (string, error) {
	id, err := sf.NextID()
	sid := fmt.Sprint(id)
	return sid, err
}

//集群中生成所属单例ID
//传入进程ID即可
func ClusterUniqueID(pid int) (string, error) {
	id, err := Unique()
	if err != nil {
		return "", err
	}
	rs := common.Md5(fmt.Sprintf(`%s%d`, id, pid))
	return rs, nil
}
