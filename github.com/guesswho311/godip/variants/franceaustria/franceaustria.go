package franceaustria

import (
	"github.com/guesswho311/godip/state"
	"github.com/guesswho311/godip/variants/classical"
	"github.com/guesswho311/godip/variants/classical/orders"
	"github.com/guesswho311/godip/variants/classical/start"
	"github.com/guesswho311/godip/variants/common"

	cla "github.com/guesswho311/godip/variants/classical/common"
	dip "github.com/guesswho311/godip/common"
)

var FranceAustriaVariant = common.Variant{
	Name: "France vs Austria",
	Graph: func() dip.Graph {
		okNations := map[dip.Nation]bool{
			cla.France:  true,
			cla.Austria: true,
			cla.Neutral: true,
		}
		neutral := cla.Neutral
		result := start.Graph()
		for _, node := range result.Nodes {
			if node.SC != nil && !okNations[*node.SC] {
				node.SC = &neutral
			}
		}
		return result
	},
	Start: func() (result *state.State, err error) {
		if result, err = classical.Start(); err != nil {
			return
		}
		if err = result.SetUnits(map[dip.Province]dip.Unit{
			"bre": dip.Unit{cla.Fleet, cla.France},
			"par": dip.Unit{cla.Army, cla.France},
			"mar": dip.Unit{cla.Army, cla.France},
			"tri": dip.Unit{cla.Fleet, cla.Austria},
			"vie": dip.Unit{cla.Army, cla.Austria},
			"bud": dip.Unit{cla.Army, cla.Austria},
		}); err != nil {
			return
		}
		result.SetSupplyCenters(map[dip.Province]dip.Nation{
			"bre": cla.France,
			"par": cla.France,
			"mar": cla.France,
			"tri": cla.Austria,
			"vie": cla.Austria,
			"bud": cla.Austria,
		})
		return
	},
	Blank:             classical.Blank,
	Phase:             classical.Phase,
	ParseOrders:       orders.ParseAll,
	ParseOrder:        orders.Parse,
	OrderTypes:        orders.OrderTypes(),
	Nations:           []dip.Nation{cla.Austria, cla.France},
	PhaseTypes:        cla.PhaseTypes,
	Seasons:           cla.Seasons,
	UnitTypes:         cla.UnitTypes,
	SoloSupplyCenters: 18,
	SVGMap: func() ([]byte, error) {
		return classical.Asset("svg/map.svg")
	},
	SVGVersion: "1482957154",
	SVGUnits: map[dip.UnitType]func() ([]byte, error){
		cla.Army: func() ([]byte, error) {
			return classical.Asset("svg/army.svg")
		},
		cla.Fleet: func() ([]byte, error) {
			return classical.Asset("svg/fleet.svg")
		},
	},
	CreatedBy: "",
	Version: "",
	Description: "A two player variant on the classical map.",
	Rules: "The first to 18 supply centers is the winner. The rules are as per classical Diplomacy, but with only France and Austria.",
}
