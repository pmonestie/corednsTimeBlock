package timeblock

import (
	"context"
	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/plugin/metrics"
	"github.com/coredns/coredns/request"
	"github.com/infobloxopen/go-trees/iptree"
	"github.com/miekg/dns"
	"net"
	"time"
)

// ACL enforces access control policies on DNS queries.
type TIME struct {
	Next       plugin.Handler
	incRange IncRange
	iptree        *iptree.Tree
}
type IncRange struct {
	dayStart int
	dayEnd int
	rangeStart int
	rangeEnd   int
}

// ServeDNS implements the plugin.Handler interface.
func (a TIME) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	state := request.Request{W: w, Req: r}
	ip := net.ParseIP(state.IP())
	hours, minutes, _ := time.Now().Clock()
	currentTimeMin := hours*60 + minutes
	currentDay := (int)(time.Now().Weekday())
	inDayRange := currentDay >= a.incRange.dayStart && currentDay <= a.incRange.dayEnd
	inTimeBlockRange := inDayRange && currentTimeMin >= a.incRange.rangeStart && currentTimeMin <= a.incRange.rangeEnd
	_, contained := a.iptree.GetByIP(ip)

	if contained && inTimeBlockRange {
		RequestBlockCount.WithLabelValues(metrics.WithServer(ctx), ip.String()).Inc()
		m := new(dns.Msg)
		m.SetRcode(r, dns.RcodeRefused)
		w.WriteMsg(m)
		return dns.RcodeSuccess, nil
	}
	RequestAllowCount.WithLabelValues(metrics.WithServer(ctx)).Inc()
	return plugin.NextOrFailure(state.Name(), a.Next, ctx, w, r)
}


// Name implements the plugin.Handler interface.
func (a TIME) Name() string {
	return "timeblock"
}
