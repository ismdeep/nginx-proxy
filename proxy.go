package main

import (
	"fmt"
	"github.com/google/uuid"
)

func GenerateUpstream(p proxy) string {
	udpReuse := ""
	if p.Type == "UDP" {
		udpReuse = " udp reuseport"
	}
	s := fmt.Sprintf("srv-%v", uuid.New().String())
	return fmt.Sprintf(`# %v
upstream %v {
    server %v;
}
server {
    listen %v%v;
    proxy_connect_timeout 8s;
    proxy_timeout 876000h;
    proxy_pass %v;
}

`, p.Name, s, p.Addr, p.Port, udpReuse, s)
}
