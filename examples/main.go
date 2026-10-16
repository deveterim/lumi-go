package main

import (
	"fmt"
	"log"

	"github.com/deveterim/lumi-go"
)

func main() {
	client := lumi.NewClient("lu_live_123456789")

	params := &lumi.SendEmailRequest{
		From:    "onboarding@suaempresa.com",
		To:      []string{"destinatario@exemplo.com"},
		Subject: "Sua fatura foi emitida",
		Html:    "<h1>Fatura #9482</h1><p>Obrigado pela preferência.</p>",
	}

	sent, err := client.Emails.Send(params)
	if err != nil {
		log.Fatalf("Erro ao enviar e-mail: %v", err)
	}

	fmt.Printf("E-mail submetido com sucesso! ID: %s (Status: %s)\n", sent.Id, sent.Status)
}
