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

	flags := []cli.Flag {
		cli.StringFlag{
			Name: "domain",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command {
		{
			Name: "ip",
			Usage: "Busca Ips de domínios na internet",
			Action: BuscarIps,
			Flags: flags,
		},
		{
			Name: "server",
			Usage: "Busca o nomes de servidores na internet",
			Flags: flags,
			Action: BuscarServidores,
		},
	}

	return app
}

func BuscarIps(c *cli.Context) {
	domain := c.String("domain")

	ips, erro := net.LookupIP(domain); 
	if erro != nil {
		log.Fatal(erro)
	}

	for _ , ip := range ips {
		fmt.Println(ip)
	}
}

func BuscarServidores (c *cli.Context) {
	domain := c.String("domain")

	servidores, erro := net.LookupNS(domain);
	if erro != nil {
		log.Fatal(erro)
	}

	for _ , servidor := range servidores {
		fmt.Println(servidor)
	}
}