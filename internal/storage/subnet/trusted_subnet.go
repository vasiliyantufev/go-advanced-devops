package subnet

var TrustedSubnet = []string{
	"192.168.1.63",
	"192.168.1.0",
	"192.168.1.1",
}

func GetTrustedSubnet() []string {
	return TrustedSubnet
}
