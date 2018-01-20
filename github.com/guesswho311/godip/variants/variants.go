package variants

import (
	"github.com/guesswho311/godip/variants/common"
	"github.com/guesswho311/godip/variants/classical"
	"github.com/guesswho311/godip/variants/coldwar"
	"github.com/guesswho311/godip/variants/fleetrome"
	"github.com/guesswho311/godip/variants/franceaustria"
	"github.com/guesswho311/godip/variants/pure"
	"github.com/guesswho311/godip/variants/ancientmediterranean"
)

func init() {
	for _, variant := range OrderedVariants {
		Variants[variant.Name] = variant
	}
}

var Variants = map[string]common.Variant{}

var OrderedVariants = []common.Variant{
	classical.ClassicalVariant,
	coldwar.ColdWarVariant,
	fleetrome.FleetRomeVariant,
	franceaustria.FranceAustriaVariant,
	pure.PureVariant,
	ancientmediterranean.AncientMediterraneanVariant,
}
