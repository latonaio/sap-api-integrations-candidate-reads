package main

import (
	sap_api_caller "sap-api-integrations-candidate-reads/SAP_API_Caller"
	"sap-api-integrations-candidate-reads/SAP_API_Input_Reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Candidate_Candidate_By_Name_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/successfactors/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Header", "Languages", "Education", "Certificates", "OutsideWorkExperience",
			"CandidateByName",
		}
	}

	caller.AsyncGetCandidate(
		inoutSDC.Candidate.CandidateID,
		inoutSDC.Candidate.FirstName,
		inoutSDC.Candidate.LastName,
		accepter,
	)
}
