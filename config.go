package main

type proxy struct {
	Name string `yaml:"name"`
	Addr string `yaml:"addr"`
	Type string `yaml:"type"`
	Port int    `yaml:"port"`
}

type config struct {
	Proxies []proxy `yaml:"proxies"`
}
