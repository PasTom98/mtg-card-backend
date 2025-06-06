package models

type Card struct {
	Id         int64    `json:"id,omitempty"`
	ScryfallID string   `json:"scryfall_id,omitempty"`
	Name       string   `json:"name,omitempty"`
	OracleText string   `json:"oracle_text,omitempty"`
	ManaCost   string   `json:"mana_cost,omitempty"`
	Types      []string `json:"types,omitempty"`
	Set        string   `json:"set,omitempty"`
}
