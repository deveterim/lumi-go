# Lumi Go SDK (`lumi-go`)

> SDK oficial em Go para a infraestrutura de e-mails transacionais **Lumi**, com interface drop-in replacement compatível com o Resend Go SDK.

---

## 📦 Instalação

```bash
go get -u github.com/deveterim/lumi-go
```

---

## 🚀 Como Usar

### 1. Inicializar Cliente e Enviar E-mail

```go
package main

import (
	"fmt"
	"log"

	"github.com/deveterim/lumi-go"
)

func main() {
	client := lumi.NewClient("ma_live_123456789")
	// Ou configure a variável LUMI_API_KEY e use lumi.NewClient("")

	params := &lumi.SendEmailRequest{
		From:    "onboarding@suaempresa.ao",
		To:      []string{"cliente@gmail.com"},
		Subject: "Boas-vindas!",
		Html:    "<h1>Olá!</h1><p>Seu e-mail transacional foi entregue.</p>",
	}

	sent, err := client.Emails.Send(params)
	if err != nil {
		log.Fatalf("Erro ao enviar: %v", err)
	}

	fmt.Printf("E-mail enviado! ID: %s (Status: %s)\n", sent.Id, sent.Status)
}
```

### 2. Envio em Lote (Batch)

```go
emails := []*lumi.SendEmailRequest{
	{
		From:    "news@suaempresa.ao",
		To:      []string{"user1@gmail.com"},
		Subject: "Atualização #1",
		Html:    "<p>Texto 1</p>",
	},
	{
		From:    "news@suaempresa.ao",
		To:      []string{"user2@gmail.com"},
		Subject: "Atualização #2",
		Html:    "<p>Texto 2</p>",
	},
}

resp, err := client.Batch.Send(emails)
```

### 3. Gestão de Domínios & API Keys

```go
// Criar domínio
domain, err := client.Domains.Create("suaempresa.ao")

// Verificar DNS
verified, err := client.Domains.Verify(domain.Id)

// Listar API Keys
keys, err := client.ApiKeys.List()
```

### 4. Validação de Webhooks HMAC

```go
isValid := client.Webhooks.VerifySignature(payloadBytes, signatureHeader, webhookSecret)
```

---

## 📄 Licença

MIT © [Lumi](https://lumi.email)
