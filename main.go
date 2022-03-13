package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func main() {
	raw, err := ioutil.ReadFile(os.Getenv("CONFIG_FILE"))
	if err != nil {
		panic(err)
	}
	var conf config
	if err := yaml.Unmarshal(raw, &conf); err != nil {
		panic(err)
	}

	output := os.Getenv("UPSTREAM_OUTPUT_FILE")
	f, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for i := 0; i < len(conf.Proxies); i++ {
		if conf.Proxies[i].Type == "" {
			conf.Proxies[i].Type = "TCP"
		}
		if conf.Proxies[i].Port <= 0 || conf.Proxies[i].Port > 65535 {
			panic(fmt.Errorf("invalid port value [%v] in [%v]", conf.Proxies[i].Port, conf.Proxies[i].Name))
		}
		if _, err := f.WriteString(GenerateUpstream(conf.Proxies[i])); err != nil {
			panic(err)
		}
	}
}
