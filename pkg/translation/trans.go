package translation

import (
	"common_service/global"
	"fmt"
	"strconv"
)

func TransT(key interface{}, params ...string) string {
	s, err := global.Trans.T(key, params...)
	if err != nil {
		global.Logger.Err(fmt.Errorf("global.Trans.T '%s': %w", key, err)).Send()
		switch key.(type) {
		case int:
			s = strconv.Itoa(key.(int))
		case string:
			s = key.(string)
		}
	}
	return s
}
