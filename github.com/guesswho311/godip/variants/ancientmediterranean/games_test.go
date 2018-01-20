package ancientmediterranean

import (
	"testing"

	dip "github.com/guesswho311/godip/common"
	tst "github.com/guesswho311/godip/variants/testing"
)

func init() {
	dip.Debug = true
}

func TestGames(t *testing.T) {
	tst.TestGames(t, AncientMediterraneanVariant)
}
