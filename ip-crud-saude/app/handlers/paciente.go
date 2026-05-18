package handlers

import "time"

// Paciente representa a estrutura dos dados que serão salvos no banco e lidos no JSON
type Paciente struct {
	ID          int       `json:"id"`
	Nome        string    `json:"nome"`
	Idade       int       `json:"idade"`
	Sintomas    string    `json:"sintomas"`
	DataCriacao time.Time `json:"data_criacao,omitempty"` // omitempty esconde o campo se ele estiver vazio
}
