package start

import (
	. "github.com/guesswho311/godip/variants/classical/common"
	"github.com/guesswho311/godip/common"
)

func SupplyCenters() map[common.Province]common.Nation {
	return map[common.Province]common.Nation{
		"edi": England,
		"lvp": England,
		"lon": England,
		"bre": France,
		"par": France,
		"mar": France,
		"kie": Germany,
		"ber": Germany,
		"mun": Germany,
		"ven": Italy,
		"rom": Italy,
		"nap": Italy,
		"tri": Austria,
		"vie": Austria,
		"bud": Austria,
		"con": Turkey,
		"ank": Turkey,
		"smy": Turkey,
		"sev": Russia,
		"mos": Russia,
		"stp": Russia,
		"war": Russia,
	}
}
