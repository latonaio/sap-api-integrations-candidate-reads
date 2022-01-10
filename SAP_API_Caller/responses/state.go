package responses

type State struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ID                   string `json:"id"`
			MinValue             string `json:"minValue"`
			ExternalCode         string `json:"externalCode"`
			MaxValue             string `json:"maxValue"`
			OptionValue          string `json:"optionValue"`
			SortOrder            int    `json:"sortOrder"`
			MdfExternalCode      string `json:"mdfExternalCode"`
			Status               string `json:"status"`
			ParentPicklistOption struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"parentPicklistOption"`
			PicklistLabels struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"picklistLabels"`
			Picklist struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"picklist"`
			ChildPicklistOptions struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"childPicklistOptions"`
		} `json:"results"`
	} `json:"d"`
}