package response

type ProductResponses struct {
    ID        int64              `json:"id"`
    Name      string             `json:"name"`
    Price     float64            `json:"price"`
    Detail    string             `json:"detail"`
    Stock     int64              `json:"stock"`
    Image     string             `json:"image"`
    Category  CategoryResponses  `json:"category"`
    Created_at int64              `json:"created_at"`
    Updated_at int64              `json:"updated_at"`
}

type CategoryResponses struct {
    ID   int64  `json:"id"`
    Name string `json:"name"`
    
}
