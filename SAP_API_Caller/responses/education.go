package responses

type Education struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			BackgroundElementID  string      `json:"backgroundElementId"`
			SchoolCity           string      `json:"schoolCity"`
			DegreeDate           string      `json:"degreeDate"`
			LastModifiedDateTime string      `json:"lastModifiedDateTime"`
			SchoolState          string      `json:"schoolState"`
			EndDate              string      `json:"endDate"`
			SchoolCountry        string      `json:"schoolCountry"`
			SchoolType           string      `json:"schoolType"`
			SchoolPhone          string      `json:"schoolPhone"`
			Major                string      `json:"major"`
			PresentStudent       int         `json:"presentStudent"`
			School               string      `json:"school"`
			BgOrderPos           string      `json:"bgOrderPos"`
			SchoolZip            string      `json:"schoolZip"`
			Degree               string      `json:"degree"`
			SchoolAddress        string      `json:"schoolAddress"`
			CandidateID          string      `json:"candidateId"`
			StartDate            string      `json:"startDate"`
			DegreeNav            struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"degreeNav"`
			Candidate struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"candidate"`
			SchoolStateNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"schoolStateNav"`
			MajorNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"majorNav"`
			PresentStudentNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"presentStudentNav"`
			SchoolCountryNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"schoolCountryNav"`
		} `json:"results"`
	} `json:"d"`
}
