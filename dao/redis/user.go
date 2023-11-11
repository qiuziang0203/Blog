package redis

import (
	"context"
	"strconv"
)

func DelUserInfo(userid int64) {
	ctx := context.Background()
	Rdb.Del(ctx, "user_info::user_id"+strconv.FormatInt(userid, 10))
}
