package entities

type Payment struct {
    PayerEmail string  `json:"payer_email"`
    Token      string  `json:"token"`
    PlanID     string  `json:"plan_id"`
    Amount     float64 `json:"amount"`
    Currency   string  `json:"currency"`
}