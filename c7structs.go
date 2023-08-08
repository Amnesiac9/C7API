package C7API

import "time"

// C7 autogenerated structs
type C7Orders struct {
	Orders []C7Order `json:"orders"`
	Total  int       `json:"total"`
}

type C7Order struct {
	//AdditionalData any `json:"additionalData"`
	//AppData        any `json:"appData"`
	//AppSync        any `json:"appSync"`
	BillTo struct {
		Address           string `json:"address"`
		Address2          string `json:"address2"`
		BirthDate         string `json:"birthDate"`
		City              string `json:"city"`
		Company           string `json:"company"`
		CountryCode       string `json:"countryCode"`
		CustomerAddressID string `json:"customerAddressId"`
		FirstName         string `json:"firstName"`
		Honorific         string `json:"honorific"`
		ID                string `json:"id"`
		LastName          string `json:"lastName"`
		Phone             string `json:"phone"`
		StateCode         string `json:"stateCode"`
		ZipCode           string `json:"zipCode"`
	} `json:"billTo"`
	//BottleDepositTotal    int    `json:"bottleDepositTotal"`
	//CarrierPickupLocation string `json:"carrierPickupLocation"`
	//CarryOut struct {
	//  ID string `json:"id"`
	// 	InventoryLocationID string `json:"inventoryLocationId"`
	// 	customerAddressId string `json:"customerAddressId"`
	// 	BirthDate string `json:"birthDate"`
	// 	Honorific string `json:"honorific"`
	// 	FirstName string `json:"firstName"`
	// 	LastName string `json:"lastName"`
	// 	Company string `json:"company"`
	// 	Phone string `json:"phone"`
	// 	Address string `json:"address"`
	// 	Address2 string `json:"address2"`
	// 	City string `json:"city"`
	// 	StateCode string `json:"stateCode"`
	// 	ZipCode string `json:"zipCode"`
	// 	CountryCode string `json:"countryCode"`
	//} `json:"carryOut"`
	//CartID                string `json:"cartId"`
	Channel string `json:"channel"`
	Club    *struct {
		//ClubID              string `json:"clubId"`
		//ClubPackageID       string `json:"clubPackageId"`
		ClubPackageTitle string `json:"clubPackageTitle"`
		ClubTitle        string `json:"clubTitle"`
		//ID                  string `json:"id"`
		ShipmentBuildStatus string `json:"shipmentBuildStatus"`
	} `json:"club"`
	//ComplianceStatus string `json:"complianceStatus"`
	// ConnectionInformation struct {
	// 	CustomerIpAddress string `json:"customerIpAddress"`
	// 	UserAgent         string `json:"userAgent"`
	// } `json:"connectionInformation"`
	// Coupons []struct {
	// 	Code          string `json:"code"`
	// 	CouponID      string `json:"couponId"`
	// 	ID            string `json:"id"`
	// 	InUse         bool   `json:"inUse"`
	// 	ProductValue  int    `json:"productValue"`
	// 	ShippingValue int    `json:"shippingValue"`
	// 	Title         string `json:"title"`
	// 	TotalValue    int    `json:"totalValue"`
	// } `json:"coupons"`
	//CreatedAt time.Time `json:"createdAt"`
	Customer struct {
		//AppData   string `json:"appData"`
		//AppSync   string `json:"appSync"`
		//Avatar    string `json:"avatar"`
		BirthDate string `json:"birthDate"`
		City      string `json:"city"`
		Clubs     []struct {
			CancelDate       time.Time `json:"cancelDate"`
			ClubID           string    `json:"clubId"`
			ClubMembershipID string    `json:"clubMembershipId"`
			ClubTitle        string    `json:"clubTitle"`
			SignupDate       time.Time `json:"signupDate"`
		} `json:"clubs"`
		CountryCode string    `json:"countryCode"`
		CreatedAt   time.Time `json:"createdAt"`
		//EmailMarketingStatus string    `json:"emailMarketingStatus"`
		Emails []struct {
			Email  string `json:"email"`
			ID     string `json:"id"`
			Status string `json:"status"`
		} `json:"emails"`
		Honorific string `json:"honorific"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		ID        string `json:"id"`
		Flags     []struct {
			//ID      string `json:"id"`
			Content string `json:"content"`
		} `json:"flags"`
		//HasAccount       bool      `json:"hasAccount"`
		// LastActivityDate time.Time `json:"lastActivityDate"`
		// LoginActivity    struct {
		// 	LastLoginAt  time.Time `json:"lastLoginAt"`
		// 	LastLogoutAt time.Time `json:"lastLogoutAt"`
		// 	LoginIp      string    `json:"loginIP"`
		// } `json:"loginActivity"`
		MetaData map[string]string `json:"metaData"`
		//Notifications    []any  `json:"notifications"`
		OrderInformation struct {
			// AcquisitionChannel string `json:"acquisitionChannel"`
			CurrentClubTitle string `json:"currentClubTitle"`
			// CurrentWebCartID        string    `json:"currentWebCartId"`
			// DaysInClub              int       `json:"daysInClub"`
			// DaysInCurrentClub       int       `json:"daysInCurrentClub"`
			// GrossProfit             int       `json:"grossProfit"`
			// LastOrderDate           time.Time `json:"lastOrderDate"`
			// LastOrderID             string    `json:"lastOrderId"`
			// LifetimeValue           int       `json:"lifetimeValue"`
			// LifetimeValueSeedAmount int       `json:"lifetimeValueSeedAmount"`
			// OrderCount              int       `json:"orderCount"`
			// Rank                    int       `json:"rank"`
			// RankTrend               string    `json:"rankTrend"`
			// YearlyValue             struct {
			// 	Y2023 int `json:"2023"`
			// } `json:"yearlyValue"`
		} `json:"orderInformation"`
		Phones []struct {
			ID    string `json:"id"`
			Phone string `json:"phone"`
		} `json:"phones"`
		Products []struct {
			Product struct {
				Image     string `json:"image"`
				Price     int    `json:"price"`
				ProductID string `json:"productId"`
				Quantity  int    `json:"quantity"`
				Sku       string `json:"sku"`
				Title     string `json:"title"`
				// WineProperties struct {
				// 	Appellation string `json:"appellation"`
				// 	CountryCode string `json:"countryCode"`
				// 	Region      string `json:"region"`
				// 	Type        string `json:"type"`
				// 	Varietal    string `json:"varietal"`
				// 	Vintage     int    `json:"vintage"`
				// } `json:"wineProperties,omitempty"`
			} `json:"product"`
			PurchaseDate time.Time `json:"purchaseDate"`
		} `json:"products"`
		StateCode string    `json:"stateCode"`
		UpdatedAt time.Time `json:"updatedAt"`
		ZipCode   string    `json:"zipCode"`
	} `json:"customer"`
	CustomerID string `json:"customerId"`
	//Duties     []any  `json:"duties"`
	//DutyTotal int `json:"dutyTotal"`
	Flags []struct {
		//ID      string `json:"id"`
		Content string `json:"content"`
	} `json:"flags"`
	// Fraud struct {
	// 	Comment      string `json:"comment"`
	// 	IsFraudulent bool   `json:"isFraudulent"`
	// } `json:"fraud"`
	//FraudCheckStatus  string `json:"fraudCheckStatus"`
	FulfillmentStatus string `json:"fulfillmentStatus"`
	//Fulfillments      []any  `json:"fulfillments"`
	GiftMessage string `json:"giftMessage"`
	ID          string `json:"id"`
	//IsNoDuty     bool   `json:"isNoDuty"`
	//IsNonTaxable bool   `json:"isNonTaxable"`
	Items []struct {
		AllocationID  string `json:"allocationId"`
		BottleDeposit int    `json:"bottleDeposit"`
		BundleItems   []struct {
			//AllocationID               string   `json:"allocationId"`
			//BottleDeposit              int      `json:"bottleDeposit"`
			//CollectionIds              []string `json:"collectionIds"`
			//ComparePrice               int      `json:"comparePrice"`
			//CostOfGood                 int      `json:"costOfGood"`
			//DepartmentCode             string   `json:"departmentCode"`
			//DepartmentID               string   `json:"departmentId"`
			//IsOverrideOperatingRegions bool     `json:"isOverrideOperatingRegions"`
			//IsPriceOverride     bool    `json:"isPriceOverride"`
			//OriginalPrice       float64 `json:"originalPrice"`
			//ProductSlug                string   `json:"productSlug"`
			//Vendor              string  `json:"vendor"`
			HasInventory        bool    `json:"hasInventory"`
			HasShipping         bool    `json:"hasShipping"`
			Image               string  `json:"image"`
			InventoryLocationID string  `json:"inventoryLocationId"`
			Price               float64 `json:"price"`
			ProductID           string  `json:"productId"`
			ProductTitle        string  `json:"productTitle"`
			ProductVariantID    string  `json:"productVariantId"`
			ProductVariantTitle string  `json:"productVariantTitle"`
			Quantity            int     `json:"quantity"`
			Sku                 string  `json:"sku"`
			Tax                 float64 `json:"tax"`
			TaxType             string  `json:"taxType"`
			Type                string  `json:"type"`
			VolumeInMl          int     `json:"volumeInML"`
			Weight              float64 `json:"weight"`
		} `json:"bundleItems"`
		//CollectionIds       string `json:"collectionIds"`
		//ComparePrice        int    `json:"comparePrice"`
		//CostOfGood 		  int `json:"costOfGood"`
		//DepartmentCode      string `json:"departmentCode"`
		//DepartmentID        string `json:"departmentId"`
		//IsPriceOverride     bool     `json:"isPriceOverride"`
		//ItemData            string   `json:"itemData"`
		//Modifiers           string   `json:"modifiers"`
		//Notes               []string `json:"notes"`
		//OriginalPrice       int     `json:"originalPrice"`
		//ProductSlug         string  `json:"productSlug"`
		//QuantitySentToPrep  int     `json:"quantitySentToPrep"`
		//TaxType             string  `json:"taxType"`
		//Type                string  `json:"type"`
		//Vendor              string  `json:"vendor"`
		HasShipping         bool    `json:"hasShipping"`
		ID                  string  `json:"id"`
		Image               string  `json:"image"`
		InventoryLocationID string  `json:"inventoryLocationId"`
		Price               float64 `json:"price"`
		ProductID           string  `json:"productId"`
		ProductTitle        string  `json:"productTitle"`
		ProductVariantID    string  `json:"productVariantId"`
		ProductVariantTitle string  `json:"productVariantTitle"`
		PurchaseType        string  `json:"purchaseType"`
		Quantity            int     `json:"quantity"`
		QuantityFulfilled   int     `json:"quantityFulfilled"`
		Sku                 string  `json:"sku"`
		Tax                 int     `json:"tax"`
		VolumeInMl          int     `json:"volumeInML"`
		Weight              float64 `json:"weight"`
	} `json:"items"`
	// LinkedOrders *[]struct {
	// 	OrderID      string `json:"orderId"`
	// 	OrderNumber  int    `json:"orderNumber"`
	// 	PurchaseType string `json:"purchaseType"`
	// } `json:"linkedOrders"`
	//LoyaltyPointsEarned int       `json:"loyaltyPointsEarned"`
	//OrderFulfilledDate  time.Time `json:"orderFulfilledDate"`
	//MetaData            string    `json:"metaData"`
	//OrderSource         string    `json:"orderSource"`
	OrderDeliveryMethod string    `json:"orderDeliveryMethod"`
	OrderNumber         int       `json:"orderNumber"`
	OrderSubmittedDate  time.Time `json:"orderSubmittedDate"`
	OrderPaidDate       time.Time `json:"orderPaidDate"`
	OrderTags           []struct {
		AppliesToCondition string    `json:"appliesToCondition"`
		CreatedAt          time.Time `json:"createdAt"`
		ID                 string    `json:"id"`
		ObjectType         string    `json:"objectType"`
		Title              string    `json:"title"`
		Type               string    `json:"type"`
		UpdatedAt          time.Time `json:"updatedAt"`
	} `json:"orderTags"`
	PaymentStatus string `json:"paymentStatus"`
	// PickupBy      struct {
	// 	Id                  string `json:"id"`
	// 	City                string `json:"city"`
	// 	Phone               string `json:"phone"`
	// 	Address             string `json:"address"`
	// 	Company             string `json:"company"`
	// 	ZipCode             string `json:"zipCode"`
	// 	Address2            string `json:"address2"`
	// 	FirstName           string `json:"firstName"`
	// 	LastName            string `json:"lastName"`
	// 	Honorific           string `json:"honorific"`
	// 	StateCode           string `json:"stateCode"`
	// 	CountryCode         string `json:"countryCode"`
	// 	CustomerAddressID   string `json:"customerAddressId"`
	// 	InventoryLocationID string `json:"inventoryLocationId"`
	// } `json:"pickupBy"`
	//PosProfile          string   `json:"posProfile"`
	//PosProfileID        string   `json:"posProfileId"`
	//PreviousOrderID     string   `json:"previousOrderId"`
	//PreviousOrderNumber int      `json:"previousOrderNumber"`
	//Promotions        []string `json:"promotions"`
	PurchaseType      string `json:"purchaseType"`
	RefundOrderID     string `json:"refundOrderId"`
	RefundOrderNumber int    `json:"refundOrderNumber"`
	//Reservation       string   `json:"reservation"`
	// SalesAssociate    struct {
	// 	AccountID string `json:"accountId"`
	// 	Name      string `json:"name"`
	// } `json:"salesAssociate"`
	SelectedShippingOptions struct {
		RequestedShipDate string `json:"requestedShipDate,omitempty"`
		//ShippingServiceID string `json:"shippingServiceId"`
	} `json:"selectedShippingOptions"`
	ShipTo struct {
		//CustomerAddressID string `json:"customerAddressId"`
		//BirthDate   string `json:"birthDate"`
		//ID          string `json:"id"`
		Address     string `json:"address"`
		Address2    string `json:"address2"`
		City        string `json:"city"`
		Company     string `json:"company"`
		CountryCode string `json:"countryCode"`
		FirstName   string `json:"firstName"`
		Honorific   string `json:"honorific"`
		LastName    string `json:"lastName"`
		Phone       string `json:"phone"`
		StateCode   string `json:"stateCode"`
		ZipCode     string `json:"zipCode"`
	} `json:"shipTo"`
	ShipTotal int `json:"shipTotal"`
	Shipping  []struct {
		Carrier           string `json:"carrier"`
		Code              string `json:"code"`
		Price             int    `json:"price"`
		Service           string `json:"service"`
		ShippingServiceID string `json:"shippingServiceId"`
		Tax               int    `json:"tax"`
		Title             string `json:"title"`
		//ID      string `json:"id"`
		//OriginalPrice     int    `json:"originalPrice"`
		//ProcessorResponse string `json:"processorResponse"`
		//Vendor            string `json:"vendor"`
	} `json:"shipping"`
	ShippingInstructions string `json:"shippingInstructions"`
	ShippingStatus       string `json:"shippingStatus"`
	SubTotal             int    `json:"subTotal"`
	Tags                 []struct {
		Title string `json:"title"`
		Type  string `json:"type"`
		//AppliesToCondition string `json:"appliesToCondition"`
		//ID                 string `json:"id"`
		//ObjectType         string    `json:"objectType"`
		//CreatedAt          time.Time `json:"createdAt"`
		//UpdatedAt          time.Time `json:"updatedAt"`
	} `json:"tags"`
	TaxSaleType string `json:"taxSaleType"`
	TaxTotal    int    `json:"taxTotal"`
	// Taxes       []struct {
	// 	Cannabis           int     `json:"cannabis"`
	// 	CountryCode        string  `json:"countryCode"`
	// 	Food               int     `json:"food"`
	// 	Freight            int     `json:"freight"`
	// 	GeneralMerchandise float64 `json:"generalMerchandise"`
	// 	ID                 string  `json:"id"`
	// 	IsFlatRate         bool    `json:"isFlatRate"`
	// 	IsIncludedInPrice  bool    `json:"isIncludedInPrice"`
	// 	IsNonTaxable       bool    `json:"isNonTaxable"`
	// 	Price              int     `json:"price"`
	// 	SortOrder          int     `json:"sortOrder"`
	// 	StateCode          string  `json:"stateCode"`
	// 	Title              string  `json:"title"`
	// 	Vendor             string  `json:"vendor"`
	// 	Wine               float64 `json:"wine"`
	// } `json:"taxes"`
	Tenders []struct {
		AmountTendered int       `json:"amountTendered"`
		ChargeStatus   string    `json:"chargeStatus"`
		ChargeType     string    `json:"chargeType"`
		CreatedAt      time.Time `json:"createdAt"`
		// CreditCard     struct {
		// 	AuthorizationID      string `json:"authorizationId"`
		// 	Bin                  string `json:"bin"`
		// 	CardBrand            string `json:"cardBrand"`
		// 	CardHolderName       string `json:"cardHolderName"`
		// 	CustomerCreditCardID string `json:"customerCreditCardId"`
		// 	ExpiryMo             int    `json:"expiryMo"`
		// 	ExpiryYr             int    `json:"expiryYr"`
		// 	Gateway              string `json:"gateway"`
		// 	MaskedCardNumber     string `json:"maskedCardNumber"`
		// 	OneTimeToken         string `json:"oneTimeToken"`
		// 	ProcessorResponse    string `json:"processorResponse"`
		// 	TokenOnFile          string `json:"tokenOnFile"`
		// } `json:"creditCard"`
		// ErrorCode string `json:"errorCode"`
		// GiftCard struct {
		// 	ID            string `json:"id"`
		// 	MaskedCode    string `json:"maskedCode"`
		// 	TransactionID string `json:"transactionId"`
		// } `json:"giftCard"`
		//ID string `json:"id"`
		// Loyalty struct {
		// 	Points        int    `json:"points"`
		// 	TransactionID string `json:"transactionId"`
		// } `json:"loyalty"`
		// OtherPaymentMethod string    `json:"otherPaymentMethod"`
		// PaymentDate        time.Time `json:"paymentDate"`
		// PreviousTenderID   string    `json:"previousTenderId"`
		TenderType string `json:"tenderType"`
		// Tip        int    `json:"tip"`
		// UpdatedAt  time.Time `json:"updatedAt"`
	} `json:"tenders"`
	TipTotal      int       `json:"tipTotal"`
	Total         int       `json:"total"`
	TotalAfterTip int       `json:"totalAfterTip"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type FulfillmentAllItems struct {
	SendTransactionEmail bool      `json:"sendTransactionEmail"`
	Type                 string    `json:"type"`
	FulfillmentDate      time.Time `json:"fulfillmentDate"`
	Shipped              struct {
		TrackingNumbers []string `json:"trackingNumbers"`
		Carrier         string   `json:"carrier"`
	} `json:"shipped"`
	PackageCount int `json:"packageCount"`
}
