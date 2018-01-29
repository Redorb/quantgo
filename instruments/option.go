package quantgo

import . "github.com/redorb/quantgo"

type OptionResults struct {
	delta              float64
	gamma              float64
	theta              float64
	vega               float64
	rho                float64
	dividendRho        float64
	deltaForward       float64
	elasticity         float64
	thetaPerDay        float64
	strikeSensitivity  float64
	itmCashProbability float64
	value              float64
}

type VanillaOption struct {
	payoff        StrikedTypePayoff
	exercise      Exercise
	pricingEngine PricingEngine
	results       OptionResults
}

func (o *OptionResults) reset() {
	o.delta = 0
	o.gamma = 0
	o.theta = 0
	o.vega = 0
	o.rho = 0
	o.dividendRho = 0
	o.deltaForward = 0
	o.elasticity = 0
	o.thetaPerDay = 0
	o.strikeSensitivity = 0
	o.itmCashProbability = 0
	o.value = 0
}
