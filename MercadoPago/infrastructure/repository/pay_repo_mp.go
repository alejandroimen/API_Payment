package repository

import (
    "bytes"
    "encoding/json"
    "errors"
    "net/http"
    "log"
    "os"

    "github.com/alejandroimen/API_Payment/MercadoPago/domain/entities"
    "github.com/mercadopago/sdk-go/pkg/config"
)

type PayRepo interface {
    CreatePayment(payment entities.Payment) (interface{}, error)
}

type PayRepoMP struct {
    cfg *config.Config
}

func NewPayRepoMP(cfg *config.Config) *PayRepoMP {
    return &PayRepoMP{cfg: cfg}
}

func (r *PayRepoMP) CreatePayment(payment entities.Payment) (interface{}, error) {
    // Crear payload
    payload := map[string]interface{}{
        "preapproval_plan_id": os.Getenv("PLAN_ID"),
        "reason":              payment.PlanID,
        "payer_email":         payment.PayerEmail,
        "card_token_id":       payment.Token,
        "auto_recurring": map[string]interface{}{
            "frequency":          1,
            "frequency_type":     "months",
            "transaction_amount": payment.Amount,
            "currency_id":        payment.Currency,
        },
        "back_url": "https://chat.openai.com/",
        "status":   "authorized",
    }

    payloadBytes, _ := json.Marshal(payload)
    req, err := http.NewRequest("POST", "https://api.mercadopago.com/preapproval", bytes.NewBuffer(payloadBytes))
    if err != nil {
        log.Println("Error creando la solicitud:", err)
        return nil, err
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+os.Getenv("APP_USER"))
    log.Printf("Bearer: %s", "Bearer "+os.Getenv("APP_USER"))

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error enviando la solicitud:", err)
        return nil, err
    }
    defer resp.Body.Close()
 
    if resp.StatusCode != http.StatusCreated {
        log.Printf("Error en la respuesta de MercadoPago: StatusCode=%d", resp.StatusCode)
        var errorBody map[string]interface{}
        json.NewDecoder(resp.Body).Decode(&errorBody)
        log.Printf("Detalles del error: %v", errorBody)
        return nil, errors.New("Error en la respuesta de MercadoPago")
    }

    var respBody map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
        log.Println("Error decodificando la respuesta:", err)
        return nil, err
    }

    return respBody, nil
}

