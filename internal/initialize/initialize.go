package initialize

var allowedinittypes = []string{"node", "server", "standalone"}

func Initialize(inittype string) {
	if !checkarrayvalue(allowedinittypes, inittype) {
		panic("wrong initialization type")
	}

	switch inittype {
	case "node":
		InitializeNode()
	case "server":
		InitializeServer()
	default:
		InitializeStandalone()
	}
}

func checkarrayvalue(list []string, val string) bool {
	for _, v := range list {
		if val == v {
			return true
		}
	}

	return false
}
