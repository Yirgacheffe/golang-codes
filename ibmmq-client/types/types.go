package types

// ----------------------------------------------------
type Player struct {
	Name    string `json:"name"`
	TaxNbr  string `json:"tax_nbr"`
	Tel     string `json:"tel"`
	Address string `json:"address"`
	Bank    Bank   `json:"bank"`
}

type Bank struct {
	Name    string `json:"name"`
	Account string `json:"account"`
}

type Seller struct {
	Player `json:"seller"`
}

type Buyer struct {
	Player `json:"buyer"`
}

type Item struct {
	Name        string `json:"name"`
	CatalogCode string `json:"catalog_code"`
	Spec        string `json:"spec"`
	Unit        string `json:"unit"`
	Quantity    string `json:"quantity"`
	Price       string `json:"price"`
	Amount      string `json:"amount"`
	TaxRate     string `json:"tax_rate"`
	TaxAmt      string `json:"tax_amt"`
}

type Dealer struct {
	Code       string `json:"code"`
	BranchCode string `json:"branch_code"`
}

type Request struct {
	ReqId            string `json:"req_id"`
	TimeStamp        string `json:"timestamp"`
	Kind             string `json:"kind"`
	SpecialOilMark   string `json:"special_oil_mark"`
	ListMark         string `json:"list_mark"`
	Dealer           Dealer `json:"dealer"`
	Buyer            Buyer  `json:"buyer"`
	Items            []Item `json:"items"`
	Seller           Seller `json:"seller"`
	TotalAmt         string `json:"total_amt"`
	TotalTax         string `json:"total_tax"`
	AmountWithoutTax string `json:"amount_without_tax"`
	Payee            string `json:"payee"`
	Reviewer         string `json:"reviewer"`
	Drawer           string `json:"drawer"`
	Remark           string `json:"remark"`
}

// ----------------------------------------------------
