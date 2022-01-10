package responses

type Languages struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			BackgroundElementID  string `json:"backgroundElementId"`
			LastModifiedDateTime string `json:"lastModifiedDateTime"`
			WritingProf          string `json:"writingProf"`
			Language             string `json:"language"`
			ReadingProf          string `json:"readingProf"`
			BgOrderPos           string `json:"bgOrderPos"`
			CandidateID          string `json:"candidateId"`
			SpeakingProf         string `json:"speakingProf"`
			LanguageNav          struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"languageNav"`
			SpeakingProfNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"speakingProfNav"`
			Candidate struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"candidate"`
			WritingProfNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"writingProfNav"`
			ReadingProfNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"readingProfNav"`
		} `json:"results"`
	} `json:"d"`
}
