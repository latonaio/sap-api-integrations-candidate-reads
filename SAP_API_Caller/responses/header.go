package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			CandidateID             string      `json:"candidateId"`
			Country                 string      `json:"country"`
			PartnerMemberID         string      `json:"partnerMemberId"`
			LastLoginDateTime       string      `json:"lastLoginDateTime"`
			LastModifiedDateTime    string      `json:"lastModifiedDateTime"`
			AnonymizedDateTime      string      `json:"anonymizedDateTime"`
			Anonymized              string      `json:"anonymized"`
			PublicIntranet          bool        `json:"publicIntranet"`
			UsersSysID              string      `json:"usersSysId"`
			ExternalCandidate       bool        `json:"externalCandidate"`
			PrimaryEmail            string      `json:"primaryEmail"`
			CreationDateTime        string      `json:"creationDateTime"`
			Zip                     string      `json:"zip"`
			HomePhone               string      `json:"homePhone"`
			FirstName               string      `json:"firstName"`
			PrivacyAcceptDateTime   string      `json:"privacyAcceptDateTime"`
			CurrentTitle            string      `json:"currentTitle"`
			ConsentToMarketing      string      `json:"consentToMarketing"`
			AgreeToPrivacyStatement string      `json:"agreeToPrivacyStatement"`
			DateofAvailability      string      `json:"dateofAvailability"`
			LastName                string      `json:"lastName"`
			City                    string      `json:"city"`
			DataPrivacyID           string      `json:"dataPrivacyId"`
			VisibilityOption        bool        `json:"visibilityOption"`
			Address                 string      `json:"address"`
			CandidateLocale         string      `json:"candidateLocale"`
			Address2                string      `json:"address2"`
			ContactEmail            string      `json:"contactEmail"`
			ShareProfile            string      `json:"shareProfile"`
			PartnerSource           string      `json:"partnerSource"`
			MiddleName              string      `json:"middleName"`
			CellPhone               string      `json:"cellPhone"`
			Education               struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"education"`
			InsideWorkExperience struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"insideWorkExperience"`
			JobsApplied struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"jobsApplied"`
			CandidateProfileConversionInfo struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"candidateProfileConversionInfo"`
			CoverLetter struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"coverLetter"`
			State struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"state"`
			Attachment2 struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"attachment2"`
			Attachment1 struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"attachment1"`
			Tags struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"tags"`
			Certificates struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"certificates"`
			JobReqFwdCandidates struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"jobReqFwdCandidates"`
			OutsideWorkExperience struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"outsideWorkExperience"`
			Resume struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"resume"`
			Languages struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"languages"`
			CandidateProfileExtension struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"candidateProfileExtension"`
			User struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"user"`
		} `json:"results"`
	} `json:"d"`
}
