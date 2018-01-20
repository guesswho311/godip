package common

import (
	"github.com/guesswho311/godip/state"
	"github.com/guesswho311/godip/variants/classical"
	"github.com/guesswho311/godip/variants/classical/orders"
	"github.com/guesswho311/godip/variants/common"

	dip "github.com/guesswho311/godip/common"
)

type Phase struct {
	Season        dip.Season
	Year          int
	Type          dip.PhaseType
	Units         map[dip.Province]dip.Unit
	Orders        map[dip.Nation]map[dip.Province][]string
	SupplyCenters map[dip.Province]dip.Nation
	Dislodgeds    map[dip.Province]dip.Unit
	Dislodgers    map[dip.Province]dip.Province
	Bounces       map[dip.Province]map[dip.Province]bool
	Resolutions   map[dip.Province]string
}

func NewPhase(state *state.State) *Phase {
	currentPhase := state.Phase()
	p := &Phase{
		Orders:      map[dip.Nation]map[dip.Province][]string{},
		Resolutions: map[dip.Province]string{},
		Season:      currentPhase.Season(),
		Year:        currentPhase.Year(),
		Type:        currentPhase.Type(),
	}
	var resolutions map[dip.Province]error
	p.Units, p.SupplyCenters, p.Dislodgeds, p.Dislodgers, p.Bounces, resolutions = state.Dump()
	for prov, err := range resolutions {
		if err == nil {
			p.Resolutions[prov] = "OK"
		} else {
			p.Resolutions[prov] = err.Error()
		}
	}
	return p
}

func (self *Phase) State(variant common.Variant) (*state.State, error) {
	parsedOrders, err := orders.ParseAll(self.Orders)
	if err != nil {
		return nil, err
	}
	return classical.Blank(variant.Phase(
		self.Year,
		self.Season,
		self.Type,
	)).Load(
		self.Units,
		self.SupplyCenters,
		self.Dislodgeds,
		self.Dislodgers,
		self.Bounces,
		parsedOrders,
	), nil
}
