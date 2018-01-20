package droidippy

import (
	"testing"

	"github.com/guesswho311/godip/variants/classical"

	dip "github.com/guesswho311/godip/common"
	tst "github.com/guesswho311/godip/variants/testing"
)

func init() {
	dip.Debug = true
}

func TestDroidippyGames(t *testing.T) {
	tst.TestGames(t, classical.ClassicalVariant)
}
