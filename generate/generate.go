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

//生成用户guid
//ts为用户创建时的时间戳,占guid的高33位
//order为用户顺序的数字id,占guid的中27位
//region为用户创建时的区域,占guid的低4位
func CreateUserGuid(ts, order, region uint64) uint64 {
	tmp0 := ts << 31
	tmp1 := order << 4
	guid := region | tmp1 | tmp0
	return guid
}

//根据用户guid解析出时间戳,顺序id,区域
func GetDataFromGuid(guid uint64) (uint64, uint64, uint64) {
	ts := guid >> 31
	var tmp2 uint64 = 0x7ffffff //27位都是1的16进制
	tmp3 := guid >> 4
	id := tmp3 & tmp2
	var tmp4 uint64 = 0xf //4位都是1的16进制
	region := guid & tmp4
	return ts, id, region
}
