package migrations

import "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"

func Models() []any {
	return []any{

		// (*model.Users)(nil),
		// (*model.UserRole)(nil),
		// (*model.Products)(nil),
		// (*model.Category)(nil),
		// (*model.Shipments)(nil),
		// (*model.SystemBank)(nil),
		// (*model.ProductsImage)(nil),
		(*model.Orders)(nil),
		// (*model.Order_details)(nil),
	}
}

func RawBeforeQueryMigrate() []string {
	return []string{
		`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`,
	}
}

func RawAfterQueryMigrate() []string {
	return []string{}
}
