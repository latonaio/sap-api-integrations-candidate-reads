package responses

type Certificates struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			BackgroundElementID  string      `json:"backgroundElementId"`
			LastModifiedDateTime string      `json:"lastModifiedDateTime"`
			LicenseState         string      `json:"licenseState"`
			EndDate              string      `json:"endDate"`
			Description          string      `json:"description"`
			LicenseName          string      `json:"licenseName"`
			Institution          string      `json:"institution"`
			LicenseType          string      `json:"licenseType"`
			Name                 string      `json:"name"`
			BgOrderPos           string      `json:"bgOrderPos"`
			LicenseCountry       string      `json:"licenseCountry"`
			LicenseNumber        string      `json:"licenseNumber"`
			CandidateID          string      `json:"candidateId"`
			StartDate            string      `json:"startDate"`
			LicenseStateNav      struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"licenseStateNav"`
			Candidate struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"candidate"`
			LicenseTypeNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"licenseTypeNav"`
			LicenseCountryNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"licenseCountryNav"`
		} `json:"results"`
	} `json:"d"`
}
