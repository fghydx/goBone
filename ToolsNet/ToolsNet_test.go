package ToolsNet

import "testing"

func TestGetLocalIP(t *testing.T) {
	ips := GetLocalIP()
	t.Logf("%v", ips)
}
