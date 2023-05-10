package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação Linha de Comando"
	app.Usage = "Busca Ips e Nomes de Servidores"

	app.Commands = []cli.Command {
		{
			Name: "ip",
			Usage: "Busca Ips de domínios na internet",
			Flags: []cli.Flag {
				cli.StringFlag{
					Name: "domain",
					Value: "google.com.br",
				},
			},
			Action: BuscarIps,
		},
	}

	return app
}

func BuscarIps(c *cli.Context) {
	host := c.String("domain")

	ips, erro := net.LookupIP(host); 
	if erro != nil {
		log.Fatal(erro)
	}

	for _ , ip := range ips {
		fmt.Println(ip)
	}
}