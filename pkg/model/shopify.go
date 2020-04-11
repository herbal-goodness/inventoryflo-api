package model

// Product Related structs
type ShopifyProducts struct {
	Products []ShopifyProduct `json:"products"`
}

type ShopifyProduct struct {
	ID             int64            `json:"id,omitempty"`
	Title          string           `json:"title,omitempty"`
	Handle         string           `json:"handle,omitempty"`
	ProductType    string           `json:"product_type,omitempty"`
	Vendor         string           `json:"vendor,omitempty"`
	Variants       []ShopifyVariant `json:"variants,omitempty"`
	Options        []ShopifyOption  `json:"options,omitempty"`
	CreatedAt      string           `json:"created_at,omitempty"`
	UpdatedAt      string           `json:"updated_at,omitempty"`
	PublishedAt    string           `json:"published_at,omitempty"`
	PublishedScope string           `json:"published_scope,omitempty"`
	Image          ShopifyImage     `json:"image,omitempty"`
	Images         []ShopifyImage   `json:"images,omitempty"`
	BodyHTML       string           `json:"body_html,omitempty"`
	Tags           string           `json:"tags,omitempty"`
	TemplateSuffix string           `json:"template_suffix,omitempty"`
}

type ShopifyImage struct {
	ID         int64   `json:"id,omitempty"`
	ProductID  int64   `json:"product_id,omitempty"`
	Src        string  `json:"src,omitempty"`
	VariantIds []int64 `json:"variant_ids,omitempty"`
	Height     int64   `json:"height,omitempty"`
	Width      int64   `json:"width,omitempty"`
	Position   int64   `json:"position,omitempty"`
	Alt        string  `json:"alt,omitempty"`
	CreatedAt  string  `json:"created_at,omitempty"`
	UpdatedAt  string  `json:"updated_at,omitempty"`
}

type ShopifyOption struct {
	ID        int64    `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	ProductID int64    `json:"product_id,omitempty"`
	Values    []string `json:"values,omitempty"`
	Position  int64    `json:"position,omitempty"`
}

type ShopifyVariant struct {
	ID                   int64   `json:"id,omitempty"`
	Sku                  string  `json:"sku,omitempty"`
	ProductID            int64   `json:"product_id,omitempty"`
	Title                string  `json:"title,omitempty"`
	Price                string  `json:"price,omitempty"`
	Taxable              bool    `json:"taxable,omitempty"`
	Barcode              string  `json:"barcode,omitempty"`
	CompareAtPrice       string  `json:"compare_at_price,omitempty"`
	RequiresShipping     bool    `json:"requires_shipping,omitempty"`
	Grams                int64   `json:"grams,omitempty"`
	Weight               float64 `json:"weight,omitempty"`
	WeightUnit           string  `json:"weight_unit,omitempty"`
	FulfillmentService   string  `json:"fulfillment_service,omitempty"`
	InventoryItemID      int64   `json:"inventory_item_id,omitempty"`
	InventoryManagement  string  `json:"inventory_management,omitempty"`
	InventoryPolicy      string  `json:"inventory_policy,omitempty"`
	InventoryQuantity    int64   `json:"inventory_quantity,omitempty"`
	OldInventoryQuantity int64   `json:"old_inventory_quantity,omitempty"`
	Option1              string  `json:"option1,omitempty"`
	Option2              string  `json:"option2,omitempty"`
	Option3              string  `json:"option3,omitempty"`
	ImageID              int64   `json:"image_id,omitempty"`
	Position             int64   `json:"position,omitempty"`
	CreatedAt            string  `json:"created_at,omitempty"`
	UpdatedAt            string  `json:"updated_at,omitempty"`
}

// Order related structs
type ShopifyOrders struct {
	Orders []ShopifyOrder `json:"orders"`
}

type ShopifyOrder struct {
	ID                  int64             `json:"id,omitempty"`
	Name                string            `json:"name,omitempty"`
	Number              int64             `json:"number,omitempty"`
	OrderNumber         int64             `json:"order_number,omitempty"`
	Confirmed           bool              `json:"confirmed,omitempty"`
	OrderStatusURL      string            `json:"order_status_url,omitempty"`
	FinancialStatus     string            `json:"financial_status,omitempty"`
	Gateway             string            `json:"gateway,omitempty"`
	LineItems           []ShopifyLineItem `json:"line_items,omitempty"`
	Currency            string            `json:"currency,omitempty"`
	SubtotalPrice       string            `json:"subtotal_price,omitempty"`
	TotalDiscounts      string            `json:"total_discounts,omitempty"`
	TotalLineItemsPrice string            `json:"total_line_items_price,omitempty"`
	TotalPrice          string            `json:"total_price,omitempty"`
	TotalPriceUsd       string            `json:"total_price_usd,omitempty"`
	TotalTax            string            `json:"total_tax,omitempty"`
	TotalWeight         int64             `json:"total_weight,omitempty"`
	Customer            ShopifyCustomer   `json:"customer,omitempty"`
	Email               string            `json:"email,omitempty"`
	BillingAddress      ShopifyAddress    `json:"billing_address,omitempty"`
	ShippingAddress     ShopifyAddress    `json:"shipping_address,omitempty"`
	CreatedAt           string            `json:"created_at,omitempty"`
	UpdatedAt           string            `json:"updated_at,omitempty"`
	ProcessedAt         string            `json:"processed_at,omitempty"`
	CancelledAt         string            `json:"cancelled_at,omitempty"`
	Test                bool              `json:"test,omitempty"`
}

type ShopifyLineItem struct {
	ID                  int64  `json:"id,omitempty"`
	Sku                 string `json:"sku,omitempty"`
	ProductID           int64  `json:"product_id,omitempty"`
	Title               string `json:"title,omitempty"`
	Name                string `json:"name,omitempty"`
	ProductExists       bool   `json:"product_exists,omitempty"`
	Vendor              string `json:"vendor,omitempty"`
	FulfillableQuantity int64  `json:"fulfillable_quantity,omitempty"`
	FulfillmentService  string `json:"fulfillment_service,omitempty"`
	FulfillmentStatus   string `json:"fulfillment_status,omitempty"`
	RequiresShipping    bool   `json:"requires_shipping,omitempty"`
	Grams               int64  `json:"grams,omitempty"`
	Quantity            int64  `json:"quantity,omitempty"`
	Properties          []struct {
		Name  string `json:"name,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"properties,omitempty"`
	Price                      string `json:"price,omitempty"`
	TotalDiscount              string `json:"total_discount,omitempty"`
	Taxable                    bool   `json:"taxable,omitempty"`
	GiftCard                   bool   `json:"gift_card,omitempty"`
	VariantID                  int64  `json:"variant_id,omitempty"`
	VariantTitle               string `json:"variant_title,omitempty"`
	VariantInventoryManagement string `json:"variant_inventory_management,omitempty"`
}

type ShopifyCustomer struct {
	ID                        int64          `json:"id,omitempty"`
	FirstName                 string         `json:"first_name,omitempty"`
	LastName                  string         `json:"last_name,omitempty"`
	Phone                     string         `json:"phone,omitempty"`
	Email                     string         `json:"email,omitempty"`
	VerifiedEmail             bool           `json:"verified_email,omitempty"`
	DefaultAddress            ShopifyAddress `json:"default_address,omitempty"`
	State                     string         `json:"state,omitempty"`
	MarketingOptInLevel       string         `json:"marketing_opt_in_level,omitempty"`
	AcceptsMarketing          bool           `json:"accepts_marketing,omitempty"`
	AcceptsMarketingUpdatedAt string         `json:"accepts_marketing_updated_at,omitempty"`
	OrdersCount               int64          `json:"orders_count,omitempty"`
	LastOrderID               int64          `json:"last_order_id,omitempty"`
	LastOrderName             string         `json:"last_order_name,omitempty"`
	TaxExempt                 bool           `json:"tax_exempt,omitempty"`
	TotalSpent                string         `json:"total_spent,omitempty"`
	Currency                  string         `json:"currency,omitempty"`
	Note                      string         `json:"note,omitempty"`
	Tags                      string         `json:"tags,omitempty"`
	CreatedAt                 string         `json:"created_at,omitempty"`
	UpdatedAt                 string         `json:"updated_at,omitempty"`
}

type ShopifyAddress struct {
	ID           int64   `json:"id,omitempty"`
	CustomerID   int64   `json:"customer_id,omitempty"`
	Name         string  `json:"name,omitempty"`
	FirstName    string  `json:"first_name,omitempty"`
	LastName     string  `json:"last_name,omitempty"`
	Phone        string  `json:"phone,omitempty"`
	Company      string  `json:"company,omitempty"`
	Address1     string  `json:"address1,omitempty"`
	Address2     string  `json:"address2,omitempty"`
	City         string  `json:"city,omitempty"`
	Province     string  `json:"province,omitempty"`
	ProvinceCode string  `json:"province_code,omitempty"`
	Zip          string  `json:"zip,omitempty"`
	Country      string  `json:"country,omitempty"`
	CountryCode  string  `json:"country_code,omitempty"`
	CountryName  string  `json:"country_name,omitempty"`
	Latitude     float64 `json:"latitude,omitempty"`
	Longitude    float64 `json:"longitude,omitempty"`
	Default      bool    `json:"default,omitempty"`
}
