package orders

import (
	"fmt"
	"time"

	cla "github.com/guesswho311/godip/variants/classical/common"
	dip "github.com/guesswho311/godip/common"
)

func init() {
	generators = append(generators, func() dip.Order { return &hold{} })
}

func Hold(source dip.Province) *hold {
	return &hold{
		targets: []dip.Province{source},
	}
}

type hold struct {
	targets []dip.Province
}

func (self *hold) String() string {
	return fmt.Sprintf("%v %v", self.targets[0], cla.Hold)
}

func (self *hold) Flags() map[dip.Flag]bool {
	return nil
}

func (self *hold) Type() dip.OrderType {
	return cla.Hold
}

func (self *hold) DisplayType() dip.OrderType {
	return cla.Hold
}

func (self *hold) Targets() []dip.Province {
	return self.targets
}

func (self *hold) At() time.Time {
	return time.Now()
}

func (self *hold) Adjudicate(r dip.Resolver) error {
	return nil
}

func (self *hold) Options(v dip.Validator, nation dip.Nation, src dip.Province) (result dip.Options) {
	if src.Super() != src {
		return
	}
	if v.Phase().Type() == cla.Movement {
		if v.Graph().Has(src) {
			if unit, actualSrc, ok := v.Unit(src); ok {
				if unit.Nation == nation {
					result = dip.Options{
						dip.SrcProvince(actualSrc): nil,
					}
				}
			}
		}
	}
	return
}

func (self *hold) Validate(v dip.Validator) (dip.Nation, error) {
	if v.Phase().Type() != cla.Movement {
		return "", cla.ErrInvalidPhase
	}
	if !v.Graph().Has(self.targets[0]) {
		return "", cla.ErrInvalidTarget
	}
	var ok bool
	var unit dip.Unit
	unit, self.targets[0], ok = v.Unit(self.targets[0])
	if !ok {
		return "", cla.ErrMissingUnit
	}
	return unit.Nation, nil
}

func (self *hold) Execute(state dip.State) {
}
