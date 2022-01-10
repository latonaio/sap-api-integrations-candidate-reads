package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-candidate-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetCandidate(candidateID string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(candidateID)
				wg.Done()
			}()
		case "Languages":
			func() {
				c.Languages(candidateID)
				wg.Done()
			}()
		case "Education":
			func() {
				c.Education(candidateID)
				wg.Done()
			}()
		case "Certificates":
			func() {
				c.Certificates(candidateID)
				wg.Done()
			}()
		case "OutsideWorkExperience":
			func() {
				c.OutsideWorkExperience(candidateID)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(candidateID string) {
	headerData, err := c.callCandidateSrvAPIRequirementHeader("Candidate", candidateID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(headerData)
	
	languagesData, err := c.callLanguages(headerData[0].Languages)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(languagesData)

	educationData, err := c.callEducation(headerData[0].Education)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(educationData)
	
	certificatesData, err := c.callCertificates(headerData[0].Certificates)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(certificatesData)
	
	outsideWorkExperienceData, err := c.callOutsideWorkExperience(headerData[0].OutsideWorkExperience)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outsideWorkExperienceData)

	jobsAppliedData, err := c.callJobsApplied(headerData[0].JobsApplied)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(jobsAppliedData)

	resumeData, err := c.callResume(headerData[0].Resume)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(resumeData)

	stateData, err := c.callState(headerData[0].State)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(stateData)
}

func (c *SAPAPICaller) callCandidateSrvAPIRequirementHeader(api, candidateID string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "odata/v2", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeader(req, candidateID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callLanguages(url string) ([]sap_api_output_formatter.Languages, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToLanguages(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callEducation(url string) ([]sap_api_output_formatter.Education, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToEducation(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callCertificates(url string) ([]sap_api_output_formatter.Certificates, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCertificates(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callOutsideWorkExperience(url string) ([]sap_api_output_formatter.OutsideWorkExperience, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToOutsideWorkExperience(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callJobsApplied(url string) ([]sap_api_output_formatter.JobsApplied, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToJobsApplied(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callResume(url string) (*sap_api_output_formatter.Resume, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToResume(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callState(url string) ([]sap_api_output_formatter.State, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToState(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Languages(candidateID string) {
	languagesData, err := c.callCandidateSrvAPIRequirementLanguages("CandidateBackground_Languages", candidateID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(languagesData)
}

func (c *SAPAPICaller) callCandidateSrvAPIRequirementLanguages(api, candidateID string) ([]sap_api_output_formatter.Languages, error) {
	url := strings.Join([]string{c.baseURL, "odata/v2", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithLanguages(req, candidateID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToLanguages(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Education(candidateID string) {
	educationData, err := c.callCandidateSrvAPIRequirementEducation("CandidateBackground_Education", candidateID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(educationData)
}

func (c *SAPAPICaller) callCandidateSrvAPIRequirementEducation(api, candidateID string) ([]sap_api_output_formatter.Education, error) {
	url := strings.Join([]string{c.baseURL, "odata/v2", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithEducation(req, candidateID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToEducation(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Certificates(candidateID string) {
	certificatesData, err := c.callCandidateSrvAPIRequirementCertificates("CandidateBackground_Certificates", candidateID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(certificatesData)
}

func (c *SAPAPICaller) callCandidateSrvAPIRequirementCertificates(api, candidateID string) ([]sap_api_output_formatter.Certificates, error) {
	url := strings.Join([]string{c.baseURL, "odata/v2", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithCertificates(req, candidateID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCertificates(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) OutsideWorkExperience(candidateID string) {
	outsideWorkExperienceData, err := c.callCandidateSrvAPIRequirementOutsideWorkExperience("CandidateBackground_OutsideWorkExperience", candidateID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outsideWorkExperienceData)
}

func (c *SAPAPICaller) callCandidateSrvAPIRequirementOutsideWorkExperience(api, candidateID string) ([]sap_api_output_formatter.OutsideWorkExperience, error) {
	url := strings.Join([]string{c.baseURL, "odata/v2", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithOutsideWorkExperience(req, candidateID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToOutsideWorkExperience(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithHeader(req *http.Request, candidateID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("candidateId eq '%s'", candidateID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithLanguages(req *http.Request, candidateID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("candidateId eq '%s'", candidateID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithEducation(req *http.Request, candidateID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("candidateId eq '%s'", candidateID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithCertificates(req *http.Request, candidateID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("candidateId eq '%s'", candidateID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithOutsideWorkExperience(req *http.Request, candidateID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("candidateId eq '%s'", candidateID))
	req.URL.RawQuery = params.Encode()
}
