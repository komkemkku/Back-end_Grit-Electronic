package migrations

import "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/model"

func Models() []any {
	return []any{

		(*model.Users)(nil),
		(*model.Roles)(nil),
		(*model.Permissions)(nil),
		(*model.RolePermissions)(nil),
		(*model.Admins)(nil),
		(*model.AdminLogs)(nil),
		(*model.AdminRoles)(nil),
		(*model.Products)(nil),
		(*model.Categories)(nil),
		(*model.SystemBanks)(nil),
		(*model.Payments)(nil),
		(*model.Reviews)(nil),
		(*model.Wishlists)(nil),
		(*model.Carts)(nil),
		(*model.Orders)(nil),
		(*model.OrderDetails)(nil),
		(*model.Notifications)(nil),
		(*model.Images)(nil),
		(*model.Shipments)(nil),

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
