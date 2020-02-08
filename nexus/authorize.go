package main

import (
	"fmt"
	"strings"

	"github.com/gammazero/nexus/v3/wamp"
)

type authorizer struct{}

func (a *authorizer) Authorize(sess *wamp.Session, msg wamp.Message) (bool, error) {
	m, ok := msg.(*wamp.Call)
	if !ok {
		return true, nil
	}
	if sess.Details["authrole"] == "admin" || sess.Details["authrole"] == "backend" {
		return true, nil
	}
	if strings.HasPrefix(fmt.Sprint(m.Procedure), "pi.") && sess.Details["authrole"] == "pi" {
		return true, nil
	}
	if strings.HasPrefix(fmt.Sprint(m.Procedure), "public.") && sess.Details["authrole"] == "authenticated" {
		return true, nil
	}
	return false, nil
}
