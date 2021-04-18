package balance

import "fmt"

type BalanceMgr struct {
	allBalance map[string]Balance
}

var (
	mgr = BalanceMgr{
		allBalance: make(map[string]Balance),
	}
)

func (p *BalanceMgr) registerBalance(name string, b Balance) {
	p.allBalance[name] = b
}

func RegisterBalance(name string, b Balance) {
	mgr.registerBalance(name, b)
}

func DoBalance(name string, insts []*Instance) (inst *Instance, err error) {
	balancer, ok := mgr.allBalance[name]

	if !ok {
		err = fmt.Errorf("Not found %s balance", name)
		return
	}

	fmt.Printf("use %s balance!\n", name)

	inst, err = balancer.DoBalance(insts)
	return
}
