package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-candidate-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			CandidateID:             data.CandidateID,
			Country:                 data.Country,
			PartnerMemberID:         data.PartnerMemberID,
			LastLoginDateTime:       data.LastLoginDateTime,
			LastModifiedDateTime:    data.LastModifiedDateTime,
			AnonymizedDateTime:      data.AnonymizedDateTime,
			Anonymized:              data.Anonymized,
			PublicIntranet:          data.PublicIntranet,
			UsersSysID:              data.UsersSysID,
			ExternalCandidate:       data.ExternalCandidate,
			PrimaryEmail:            data.PrimaryEmail,
			CreationDateTime:        data.CreationDateTime,
			Zip:                     data.Zip,
			HomePhone:               data.HomePhone,
			FirstName:               data.FirstName,
			PrivacyAcceptDateTime:   data.PrivacyAcceptDateTime,
			CurrentTitle:            data.CurrentTitle,
			ConsentToMarketing:      data.ConsentToMarketing,
			AgreeToPrivacyStatement: data.AgreeToPrivacyStatement,
			DateofAvailability:      data.DateofAvailability,
			LastName:                data.LastName,
			City:                    data.City,
			DataPrivacyID:           data.DataPrivacyID,
			VisibilityOption:        data.VisibilityOption,
			Address:                 data.Address,
			CandidateLocale:         data.CandidateLocale,
			Address2:                data.Address2,
			ContactEmail:            data.ContactEmail,
			ShareProfile:            data.ShareProfile,
			PartnerSource:           data.PartnerSource,
			MiddleName:              data.MiddleName,
			CellPhone:               data.CellPhone,
            Languages:               data.Languages.Deferred.URI,
            Education:               data.Education.Deferred.URI,
            Certificates:            data.Certificates.Deferred.URI,
            OutsideWorkExperience:   data.OutsideWorkExperience.Deferred.URI,
            JobsApplied:             data.JobsApplied.Deferred.URI,
            Resume:                  data.Resume.Deferred.URI,
			State:                   data.State.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToLanguages(raw []byte, l *logger.Logger) ([]Languages, error) {
	pm := &responses.Languages{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Languages. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	languages := make([]Languages, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		languages = append(languages, Languages{
	BackgroundElementID:  data.BackgroundElementID,
	LastModifiedDateTime: data.LastModifiedDateTime,
	WritingProf:          data.WritingProf,
	Language:             data.Language,
	ReadingProf:          data.ReadingProf,
	BgOrderPos:           data.BgOrderPos,
	CandidateID:          data.CandidateID,
	SpeakingProf:         data.SpeakingProf,
		})
	}

	return languages, nil
}

func ConvertToEducation(raw []byte, l *logger.Logger) ([]Education, error) {
	pm := &responses.Education{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Education. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	education := make([]Education, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		education = append(education, Education{
	BackgroundElementID:  data.BackgroundElementID,
	SchoolCity:           data.SchoolCity,
	DegreeDate:           data.DegreeDate,
	LastModifiedDateTime: data.LastModifiedDateTime,
	SchoolState:          data.SchoolState,
	EndDate:              data.EndDate,
	SchoolCountry:        data.SchoolCountry,
	SchoolType:           data.SchoolType,
	SchoolPhone:          data.SchoolPhone,
	Major:                data.Major,
	PresentStudent:       data.PresentStudent,
	School:               data.School,
	BgOrderPos:           data.BgOrderPos,
	SchoolZip:            data.SchoolZip,
	Degree:               data.Degree,
	SchoolAddress:        data.SchoolAddress,
	CandidateID:          data.CandidateID,
	StartDate:            data.StartDate,
		})
	}

	return education, nil
}

func ConvertToCertificates(raw []byte, l *logger.Logger) ([]Certificates, error) {
	pm := &responses.Certificates{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Certificates. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	certificates := make([]Certificates, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		certificates = append(certificates, Certificates{
	BackgroundElementID:  data.BackgroundElementID,
	LastModifiedDateTime: data.LastModifiedDateTime,
	LicenseState:         data.LicenseState,
	EndDate:              data.EndDate,
	Description:          data.Description,
	LicenseName:          data.LicenseName,
	Institution:          data.Institution,
	LicenseType:          data.LicenseType,
	Name:                 data.Name,
	BgOrderPos:           data.BgOrderPos,
	LicenseCountry:       data.LicenseCountry,
	LicenseNumber:        data.LicenseNumber,
	CandidateID:          data.CandidateID,
	StartDate:            data.StartDate,
		})
	}

	return certificates, nil
}

func ConvertToOutsideWorkExperience(raw []byte, l *logger.Logger) ([]OutsideWorkExperience, error) {
	pm := &responses.OutsideWorkExperience{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to OutsideWorkExperience. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	outsideWorkExperience := make([]OutsideWorkExperience, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		outsideWorkExperience = append(outsideWorkExperience, OutsideWorkExperience{
	BackgroundElementID:  data.BackgroundElementID,
	EmployerAddress:      data.EmployerAddress,
	LastModifiedDateTime: data.LastModifiedDateTime,
	EndDate:              data.EndDate,
	EmployerState:        data.EmployerState,
	StartTitle:           data.StartTitle,
	BgOrderPos:           data.BgOrderPos,
	Employer:             data.Employer,
	EmployerPhone:        data.EmployerPhone,
	PresentEmployer:      data.PresentEmployer,
	EmployerCountry:      data.EmployerCountry,
	EmployerZip:          data.EmployerZip,
	EmployerCity:         data.EmployerCity,
	EmployerEmail:        data.EmployerEmail,
	BusinessType:         data.BusinessType,
	CandidateID:          data.CandidateID,
	EmployerContact:      data.EmployerContact,
	StartDate:            data.StartDate,
		})
	}

	return outsideWorkExperience, nil
}

func ConvertToJobsApplied(raw []byte, l *logger.Logger) ([]JobsApplied, error) {
	pm := &responses.JobsApplied{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to JobsApplied. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	jobsApplied := make([]JobsApplied, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		jobsApplied = append(jobsApplied, JobsApplied{
	ApplicationID:                   data.ApplicationID,
	PosTitle:                        data.PosTitle,
	Yesanswer:                       data.Yesanswer,
	JobCode:                         data.JobCode,
	CandJust:                        data.CandJust,
	WotcStatus:                      data.WotcStatus,
	Ssn:                             data.Ssn,
	InstrEEOlable:                   data.InstrEEOlable,
	UsersSysID:                      data.UsersSysID,
	ConvictionDetails:               data.ConvictionDetails,
	Zip:                             data.Zip,
	PsychTestDte:                    data.PsychTestDte,
	FormerEmployee:                  data.FormerEmployee,
	PhoneScreenDate:                 data.PhoneScreenDate,
	Jobofferdate:                    data.Jobofferdate,
	CustRefCheckTitle1:              data.CustRefCheckTitle1,
	ExtTitle:                        data.ExtTitle,
	StatusComments:                  data.StatusComments,
	CustRefCheckLastName:            data.CustRefCheckLastName,
	InstrAppAdvice:                  data.InstrAppAdvice,
	Screendate:                      data.Screendate,
	StartDate:                       data.StartDate,
	Status:                          data.Status,
	LastName:                        data.LastName,
	InstrAppAdditionalInfo:          data.InstrAppAdditionalInfo,
	Gender:                          data.Gender,
	City:                            data.City,
	SourceLabel:                     data.SourceLabel,
	ProfileUpdated:                  data.ProfileUpdated,
	Criminalconvictionlabel:         data.Criminalconvictionlabel,
	VisaNum:                         data.VisaNum,
	DuplicateProfile:                data.DuplicateProfile,
	JobRateOfPay:                    data.JobRateOfPay,
	PsychTestTme:                    data.PsychTestTme,
	CountryCode:                     data.CountryCode,
	AverageRating:                   data.AverageRating,
	InductionAdd:                    data.InductionAdd,
	Haveyoubeenconvicted:            data.Haveyoubeenconvicted,
	JobReqID:                        data.JobReqID,
	Address:                         data.Address,
	DateAvail:                       data.DateAvail,
	InstrCandAppDetails:             data.InstrCandAppDetails,
	NonApplicantStatus:              data.NonApplicantStatus,
	ResumeUploadDate:                data.ResumeUploadDate,
	DateOfBirth:                     data.DateOfBirth,
	MedLoc:                          data.MedLoc,
	AppStatusSetItemID:              data.AppStatusSetItemID,
	CandConversionProcessed:         data.CandConversionProcessed,
	Pleaseselect:                    data.Pleaseselect,
	StateOther:                      data.StateOther,
	ReferenceComments:               data.ReferenceComments,
	AnonymizedFlag:                  data.AnonymizedFlag,
	ReferredBy:                      data.ReferredBy,
	CellPhone:                       data.CellPhone,
	ApplicationTemplateID:           data.ApplicationTemplateID,
	Country:                         data.Country,
	LastModifiedDateTime:            data.LastModifiedDateTime,
	Confnumber:                      data.Confnumber,
	Rating:                          data.Rating,
	WotcScreeningURL:                data.WotcScreeningURL,
	Source:                          data.Source,
	InstrAdditionalDocuments:        data.InstrAdditionalDocuments,
	AgencyInfo:                      data.AgencyInfo,
	Reference:                       data.Reference,
	CustRefCheckPhone1:              data.CustRefCheckPhone1,
	KnownAs:                         data.KnownAs,
	MedDate:                         data.MedDate,
	InstrubkgrndChkAttachment:       data.InstrubkgrndChkAttachment,
	InstrAppReturn:                  data.InstrAppReturn,
	CustRefCheckFirstName:           data.CustRefCheckFirstName,
	Reasonableaccomodation:          data.Reasonableaccomodation,
	TimeToHire:                      data.TimeToHire,
	WotcID:                          data.WotcID,
	Initials:                        data.Initials,
	JobStartDate:                    data.JobStartDate,
	HomePhone:                       data.HomePhone,
	OwnershpDate:                    data.OwnershpDate,
	MedTime:                         data.MedTime,
	FirstName:                       data.FirstName,
	InstrCandApproval:               data.InstrCandApproval,
	InstrEEOinstructions:            data.InstrEEOinstructions,
	CurrentTitle:                    data.CurrentTitle,
	LastModifiedByProxy:             data.LastModifiedByProxy,
	InstrCandScreen:                 data.InstrCandScreen,
	Whyask:                          data.Whyask,
	AnonymizedDate:                  data.AnonymizedDate,
	WotcFormsURL:                    data.WotcFormsURL,
	CandidateID:                     data.CandidateID,
	DataSource:                      data.DataSource,
	CandTypeWhenHired:               data.CandTypeWhenHired,
	Crimincalconvictioninstructions: data.Crimincalconvictioninstructions,
	HiredOn:                         data.HiredOn,
	InstrCandFields:                 data.InstrCandFields,
	PhoneScreenDetails:              data.PhoneScreenDetails,
	InstrCandPersInfo:               data.InstrCandPersInfo,
	BaseSalary:                      data.BaseSalary,
	Title:                           data.Title,
	VoluntaryselfID:                 data.VoluntaryselfID,
	CustRefCheckOrg1:                data.CustRefCheckOrg1,
	Owner:                           data.Owner,
	HireDate:                        data.HireDate,
	Address2:                        data.Address2,
	ContactEmail:                    data.ContactEmail,
	JobAppGUID:                      data.JobAppGUID,
	LastModifiedBy:                  data.LastModifiedBy,
	ExportedOn:                      data.ExportedOn,
	InstrCandAssess:                 data.InstrCandAssess,
	CurrentLocation:                 data.CurrentLocation,
	InductionDate:                   data.InductionDate,
	InstrWotc:                       data.InstrWotc,
	MiddleName:                      data.MiddleName,
	AppLocale:                       data.AppLocale,
	SnapShotDate:                    data.SnapShotDate,
		})
	}

	return jobsApplied, nil
}

func ConvertToResume(raw []byte, l *logger.Logger) (*Resume, error) {
	pm := &responses.Resume{}

	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Resume. unmarshal error: %w", err)
	}
	
	return &Resume{
	AttachmentID:           pm.D.AttachmentID,
	Country:                pm.D.Country,
	FileName:               pm.D.FileName,
	LastModifiedDateTime:   pm.D.LastModifiedDateTime,
	OwnerIDType:            pm.D.OwnerIDType,
	DocumentType:           pm.D.DocumentType,
	Deletable:              pm.D.Deletable,
	Description:            pm.D.Description,
	Language:               pm.D.Language,
	MimeType:               pm.D.MimeType,
	ModuleCategory:         pm.D.ModuleCategory,
	OwnerID:                pm.D.OwnerID,
	Deprecable:             pm.D.Deprecable,
	PiiFlag:                pm.D.PiiFlag,
	FileExtension:          pm.D.FileExtension,
	DocumentEntityID:       pm.D.DocumentEntityID,
	Module:                 pm.D.Module,
	DocumentCategory:       pm.D.DocumentCategory,
	ExternalID:             pm.D.ExternalID,
	UserID:                 pm.D.UserID,
	Searchable:             pm.D.Searchable,
	DocumentEntityType:     pm.D.DocumentEntityType,
	CreatedDate:            pm.D.CreatedDate,
	Viewable:               pm.D.Viewable,
	FileSize:               pm.D.FileSize,
	SoftDelete:             pm.D.SoftDelete,
	LastAccessed:           pm.D.LastAccessed,
	FileContent:            pm.D.FileContent,
	ImageConvertInProgress: pm.D.ImageConvertInProgress,
	}, nil
}

func ConvertToState(raw []byte, l *logger.Logger) ([]State, error) {
	pm := &responses.State{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to State. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	state := make([]State, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		state = append(state, State{
	ID:                   data.ID,
	MinValue:             data.MinValue,
	ExternalCode:         data.ExternalCode,
	MaxValue:             data.MaxValue,
	OptionValue:          data.OptionValue,
	SortOrder:            data.SortOrder,
	MdfExternalCode:      data.MdfExternalCode,
	Status:               data.Status,
		})
	}

	return state, nil
}
