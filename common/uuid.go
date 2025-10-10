package common

import (
	"fmt"

	"github.com/speps/go-hashids/v2"
)

var hd *hashids.HashIDData

func init() {
	// salt nên được cấu hình trong env
	hd = hashids.NewData()
	hd.Salt = "your-secret-salt-key"
	hd.MinLength = 8
}

func MaskID(id int64) (string, error) {
	if id <= 0 {
		return "", nil
	}
	h, _ := hashids.NewWithData(hd)
	return h.EncodeInt64([]int64{id})
}

func UnmaskID(code string) (int64, error) {
	h, _ := hashids.NewWithData(hd)
	nums, err := h.DecodeInt64WithError(code)
	if err != nil || len(nums) == 0 {
		return 0, fmt.Errorf("invalid masked id")
	}
	return nums[0], nil
}
