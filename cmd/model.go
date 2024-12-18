package cmd

import (
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/database/migrations"
	"context"
	"log"

	"github.com/uptrace/bun"
)

func modelUp(db *bun.DB) error {
	log.Printf("Executing model up...")
	if len(migrations.Models()) == 0 {
		log.Printf("No models to migrate")
		return nil
	}
	for _, mod := range migrations.Models() {
		if _, err := db.NewCreateTable().Model(mod).Exec(context.Background()); err != nil {
			return err
		}
	}
	return nil
}

func modelDown(db *bun.DB) error {
	log.Printf("Executing model down...")
	if len(migrations.Models()) == 0 {
		log.Printf("No models to migrate")
		return nil
	}
	for _, mod := range migrations.Models() {
		if _, err := db.NewDropTable().Model(mod).Exec(context.Background()); err != nil {
			return err
		}
	}
	return nil
}
