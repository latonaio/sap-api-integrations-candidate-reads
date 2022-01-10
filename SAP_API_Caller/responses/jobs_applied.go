package responses

type JobsApplied struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ApplicationID                   string      `json:"applicationId"`
			PosTitle                        string      `json:"posTitle"`
			Yesanswer                       string      `json:"yesanswer"`
			JobCode                         string      `json:"jobCode"`
			CandJust                        string      `json:"candJust"`
			WotcStatus                      string      `json:"wotcStatus"`
			Ssn                             string      `json:"ssn"`
			InstrEEOlable                   string      `json:"instr_EEOlable"`
			UsersSysID                      string      `json:"usersSysId"`
			ConvictionDetails               string      `json:"conviction_details"`
			Zip                             string      `json:"zip"`
			PsychTestDte                    string      `json:"psychTestDte"`
			FormerEmployee                  bool        `json:"formerEmployee"`
			PhoneScreenDate                 string      `json:"phoneScreenDate"`
			Jobofferdate                    string      `json:"jobofferdate"`
			CustRefCheckTitle1              string      `json:"cust_refCheckTitle1"`
			ExtTitle                        string      `json:"extTitle"`
			StatusComments                  string      `json:"statusComments"`
			CustRefCheckLastName            string      `json:"cust_refCheckLastName"`
			InstrAppAdvice                  string      `json:"instr_AppAdvice"`
			Screendate                      string      `json:"screendate"`
			StartDate                       string      `json:"startDate"`
			Status                          string      `json:"status"`
			LastName                        string      `json:"lastName"`
			InstrAppAdditionalInfo          string      `json:"instr_AppAdditionalInfo"`
			Gender                          string      `json:"gender"`
			City                            string      `json:"city"`
			SourceLabel                     string      `json:"sourceLabel"`
			ProfileUpdated                  string      `json:"profileUpdated"`
			Criminalconvictionlabel         string      `json:"criminalconvictionlabel"`
			VisaNum                         string      `json:"visaNum"`
			DuplicateProfile                string      `json:"duplicateProfile"`
			JobRateOfPay                    string      `json:"jobRateOfPay"`
			PsychTestTme                    string      `json:"psychTestTme"`
			CountryCode                     string      `json:"countryCode"`
			AverageRating                   string      `json:"averageRating"`
			InductionAdd                    string      `json:"inductionAdd"`
			Haveyoubeenconvicted            string      `json:"haveyoubeenconvicted"`
			JobReqID                        string      `json:"jobReqId"`
			Address                         string      `json:"address"`
			DateAvail                       string      `json:"dateAvail"`
			InstrCandAppDetails             string      `json:"instr_CandAppDetails"`
			NonApplicantStatus              string      `json:"nonApplicantStatus"`
			ResumeUploadDate                string      `json:"resumeUploadDate"`
			DateOfBirth                     string      `json:"dateOfBirth"`
			MedLoc                          string      `json:"medLoc"`
			AppStatusSetItemID              string      `json:"appStatusSetItemId"`
			CandConversionProcessed         string      `json:"candConversionProcessed"`
			Pleaseselect                    string      `json:"pleaseselect"`
			StateOther                      string      `json:"stateOther"`
			ReferenceComments               string      `json:"referenceComments"`
			AnonymizedFlag                  string      `json:"anonymizedFlag"`
			ReferredBy                      string      `json:"referredBy"`
			CellPhone                       string      `json:"cellPhone"`
			ApplicationTemplateID           string      `json:"applicationTemplateId"`
			Country                         string      `json:"country"`
			LastModifiedDateTime            string      `json:"lastModifiedDateTime"`
			Confnumber                      string      `json:"confnumber"`
			Rating                          string      `json:"rating"`
			WotcScreeningURL                string      `json:"wotcScreeningURL"`
			Source                          string      `json:"source"`
			InstrAdditionalDocuments        string      `json:"instr_AdditionalDocuments"`
			AgencyInfo                      string      `json:"agencyInfo"`
			Reference                       string      `json:"reference"`
			CustRefCheckPhone1              string      `json:"cust_refCheckPhone1"`
			KnownAs                         string      `json:"knownAs"`
			MedDate                         string      `json:"medDate"`
			InstrubkgrndChkAttachment       string      `json:"instrubkgrndChkAttachment"`
			InstrAppReturn                  string      `json:"instrAppReturn"`
			CustRefCheckFirstName           string      `json:"cust_refCheckFirstName"`
			Reasonableaccomodation          string      `json:"reasonableaccomodation"`
			TimeToHire                      string      `json:"timeToHire"`
			WotcID                          string      `json:"wotcId"`
			Initials                        string      `json:"initials"`
			JobStartDate                    string      `json:"jobStartDate"`
			HomePhone                       string      `json:"homePhone"`
			OwnershpDate                    string      `json:"ownershpDate"`
			MedTime                         string      `json:"medTime"`
			FirstName                       string      `json:"firstName"`
			InstrCandApproval               string      `json:"instrCandApproval"`
			InstrEEOinstructions            string      `json:"instr_EEOinstructions"`
			CurrentTitle                    string      `json:"currentTitle"`
			LastModifiedByProxy             string      `json:"lastModifiedByProxy"`
			InstrCandScreen                 string      `json:"instrCandScreen"`
			Whyask                          string      `json:"whyask"`
			AnonymizedDate                  string      `json:"anonymizedDate"`
			WotcFormsURL                    string      `json:"wotcFormsURL"`
			CandidateID                     string      `json:"candidateId"`
			DataSource                      string      `json:"dataSource"`
			CandTypeWhenHired               string      `json:"candTypeWhenHired"`
			Crimincalconvictioninstructions string      `json:"crimincalconvictioninstructions"`
			HiredOn                         string      `json:"hiredOn"`
			InstrCandFields                 string      `json:"instrCandFields"`
			PhoneScreenDetails              string      `json:"phoneScreenDetails"`
			InstrCandPersInfo               string      `json:"instr_CandPersInfo"`
			BaseSalary                      string      `json:"baseSalary"`
			Title                           string      `json:"title"`
			VoluntaryselfID                 string      `json:"voluntaryselfID"`
			CustRefCheckOrg1                string      `json:"cust_refCheckOrg1"`
			Owner                           string      `json:"owner"`
			HireDate                        string      `json:"hireDate"`
			Address2                        string      `json:"address2"`
			ContactEmail                    string      `json:"contactEmail"`
			JobAppGUID                      string      `json:"jobAppGuid"`
			LastModifiedBy                  string      `json:"lastModifiedBy"`
			ExportedOn                      string      `json:"exportedOn"`
			InstrCandAssess                 string      `json:"instrCandAssess"`
			CurrentLocation                 string      `json:"currentLocation"`
			InductionDate                   string      `json:"inductionDate"`
			InstrWotc                       string      `json:"instrWotc"`
			MiddleName                      string      `json:"middleName"`
			AppLocale                       string      `json:"appLocale"`
			SnapShotDate                    string      `json:"snapShotDate"`
			Education                       struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"education"`
			Disabilityselection struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"disabilityselection"`
			ManagerApproval struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"managerApproval"`
			CoverLetter struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"coverLetter"`
			JobApplicationComments struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"jobApplicationComments"`
			State struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"state"`
			JobApplicationAssessmentOrder struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"jobApplicationAssessmentOrder"`
			ReferredByNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"referredByNav"`
			OfferLetter struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"offerLetter"`
			RefCheckComp struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"refCheckComp"`
			JobApplicationInterview struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"jobApplicationInterview"`
			ReasonReject struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"reasonReject"`
			Certificates struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"certificates"`
			StatusID struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"statusId"`
			JobRequisition struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"jobRequisition"`
			Prescreen struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"prescreen"`
			CandType struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"candType"`
			Ethnicity struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"ethnicity"`
			JobApplicationQuestionResponse struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"jobApplicationQuestionResponse"`
			JobOffer struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"jobOffer"`
			ConvictionResponse struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"conviction_response"`
			CountryPicklist struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"countryPicklist"`
			VisaCustom struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"visaCustom"`
			Resume struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"resume"`
			UserNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"userNav"`
			Languages struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"languages"`
			RankHireMan struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"rankHireMan"`
			CurrentEmpQ struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"currentEmpQ"`
			CustRefCheckState struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"cust_refCheckState"`
			CustRefCheckCountry struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"cust_refCheckCountry"`
			Attachment3 struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"attachment3"`
			InsideWorkExperience struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"insideWorkExperience"`
			VeteranStatus struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"veteranStatus"`
			JobAppStatus struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"jobAppStatus"`
			Attachment1 struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"attachment1"`
			BkgrndChkAttachment struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"bkgrndChkAttachment"`
			Citizen struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"citizen"`
			CustRefCheckType struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"cust_refCheckType"`
			JobApplicationOnboardingData struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"jobApplicationOnboardingData"`
			CandidateSource struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"candidateSource"`
			JobApplicationAudit struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"jobApplicationAudit"`
			OutsideWorkExperience struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"outsideWorkExperience"`
			Salutation struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"salutation"`
			DisabilityStatus struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"disabilityStatus"`
			JobApplicationStatusAuditTrail struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"jobApplicationStatusAuditTrail"`
			Candidate struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"candidate"`
			Atsi struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"atsi"`
			BkgrndChkStatus struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"bkgrndChkStatus"`
		} `json:"results"`
	} `json:"d"`
}
