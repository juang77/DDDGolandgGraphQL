package entity

import (
	"html"
	"strings"
	"time"
)

type Product struct {
	ID          int64     `gorm:"primary_key;auto_increment" json:"product_id"`
	Name        string    `gorm:"size:200;not null;unique;" json:"product_name"`
	Description string    `gorm:"text;not null;" json:"product_description"`
	Prise       float64   `gorm:"not null;" json:"product_prise"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (p *Product) BeforeSave() {
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.Description = html.EscapeString(strings.TrimSpace(p.Description))
}

func (p *Product) Prepare() {
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.Description = html.EscapeString(strings.TrimSpace(p.Description))
}

func (p *Product) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)

	switch strings.ToLower(action) {
	case "update":
		if p.Name == "" || p.Name == "null" {
			errorMessages["name_required"] = "name is required"
		}
		if p.Description == "" || p.Description == "null" {
			errorMessages["desc_required"] = "description is required"
		}
		if p.Prise <= 0 {
			errorMessages["prise_required"] = "prise must be greater than zero"
		}
	default:
		if p.Name == "" || p.Name == "null" {
			errorMessages["name_required"] = "name is required"
		}
		if p.Description == "" || p.Description == "null" {
			errorMessages["desc_required"] = "description is required"
		}
		if p.Prise <= 0 {
			errorMessages["prise_required"] = "prise must be greater than zero"
		}
	}
	return errorMessages
}
