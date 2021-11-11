package main

type proxy struct {
	Name string
	Addr string
	Type string
	Port int
}

type config struct {
	Proxies []proxy
}
