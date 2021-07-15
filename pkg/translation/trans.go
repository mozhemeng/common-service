package translation

import (
	"common_service/global"
	"github.com/pkg/errors"
	"strconv"
)

func TransT(key interface{}, params ...string) string {
	s, err := global.Trans.T(key, params...)
	if err != nil {
		global.Logger.Error(errors.Wrapf(err, "global.Trans.T: '%s'", key))
		switch key.(type) {
		case int:
			s = strconv.Itoa(key.(int))
		case string:
			s = key.(string)
		}
	}
	return s
}
