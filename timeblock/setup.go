package timeblock

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	"github.com/infobloxopen/go-trees/iptree"
	"net"
	"strconv"
	"strings"
)

const pluginName = "timeblock"

func init() { plugin.Register(pluginName, setup) }


func setup(c *caddy.Controller) error {
	a, err := parse(c)
	if err != nil {
		return plugin.Error(pluginName, err)
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		a.Next = next
		return a
	})

	return nil
}

func parse(c *caddy.Controller) (TIME, error) {
	//Brute force config read, no checks!!
	a := TIME{}
	a.iptree = iptree.NewTree()
	c.Next()                  // Skip the plugin name, "timeblock" in this case.
	c.Next()
	value := c.Val()
	a.incRange = createRange(value)
	c.Next()
	cidr :=normalize(c.Val())
	_, source, _ := net.ParseCIDR(cidr)
	a.iptree.InplaceInsertNet(source, struct{}{})
	return a, nil
}

func createRange(s string)(IncRange)  {
	result := IncRange{}
	sp := strings.Split(s, "-")
	daySplit := strings.Split(sp[0],":")
	result.dayStart, _ = strconv.Atoi(daySplit[0])
	result.dayEnd, _ = strconv.Atoi(daySplit[1])

	lowerTimeSplit :=strings.Split(sp[1],":")
	lh,_ :=strconv.Atoi(lowerTimeSplit[0])
	lm, _:= strconv.Atoi(lowerTimeSplit[1])
	lower := lh*60 + lm
	higherTimeSplit :=strings.Split(sp[2],":")
	hh,_ :=strconv.Atoi(higherTimeSplit[0])
	hm, _:= strconv.Atoi(higherTimeSplit[1])
	higher := hh*60 + hm
	result.rangeStart = lower
	result.rangeEnd = higher
	return result
}
func normalize(rawNet string) string {
	if idx := strings.IndexAny(rawNet, "/"); idx >= 0 {
		return rawNet
	}

	if idx := strings.IndexAny(rawNet, ":"); idx >= 0 {
		return rawNet + "/128"
	}
	return rawNet + "/32"
}



