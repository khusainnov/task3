package entity

type CSVData struct {
	State                 interface{} `csv:"State"`
	ZipCode               interface{} `csv:"ZipCode"`
	TaxRegionName         interface{} `csv:"TaxRegionName"`
	StateRate             interface{} `csv:"StateRate"`
	EstimatedCombinedRate interface{} `csv:"EstimatedCombinedRate"`
	EstimatedCountyRate   interface{} `csv:"EstimatedCountyRate"`
	EstimatedCityRate     interface{} `csv:"EstimatedCityRate"`
	EstimatedSpecialRate  interface{} `csv:"EstimatedSpecialRate"`
	RiskLevel             interface{} `csv:"RiskLevel"`
}
