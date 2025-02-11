package response

type WishlistResponses struct {
    ID               int              `json:"id"`
    User             UserRespWistlist `json:"user"`    // เปลี่ยนชื่อ field จาก UserID เป็น User
    Product          ProductRespCart  `json:"product"` // เปลี่ยนชื่อ field จาก ProductID เป็น Product
    // PricePerProduct  float64          `json:"price_per_product"`
    // AmountPerProduct int              `json:"amount_per_product"`
    CreatedAt        int64            `json:"created_at"`
    UpdatedAt        int64            `json:"updated_at"`
}
