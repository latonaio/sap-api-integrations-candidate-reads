package responses

type OutsideWorkExperience struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			BackgroundElementID  string      `json:"backgroundElementId"`
			EmployerAddress      string      `json:"employerAddress"`
			LastModifiedDateTime string      `json:"lastModifiedDateTime"`
			EndDate              string      `json:"endDate"`
			EmployerState        string      `json:"employerState"`
			StartTitle           string      `json:"startTitle"`
			BgOrderPos           string      `json:"bgOrderPos"`
			Employer             string      `json:"employer"`
			EmployerPhone        string      `json:"employerPhone"`
			PresentEmployer      string      `json:"presentEmployer"`
			EmployerCountry      string      `json:"employerCountry"`
			EmployerZip          string      `json:"employerZip"`
			EmployerCity         string      `json:"employerCity"`
			EmployerEmail        string      `json:"employerEmail"`
			BusinessType         string      `json:"businessType"`
			CandidateID          string      `json:"candidateId"`
			EmployerContact      string      `json:"employerContact"`
			StartDate            string      `json:"startDate"`
			Candidate            struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"candidate"`
			PresentEmployerNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"presentEmployerNav"`
			EmployerStateNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"employerStateNav"`
			BusinessTypeNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"businessTypeNav"`
			EmployerCountryNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"employerCountryNav"`
		} `json:"results"`
	} `json:"d"`
}
