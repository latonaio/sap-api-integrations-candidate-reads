package responses

type Resume struct {
	D struct {
		Metadata struct {
			URI  string `json:"uri"`
			Type string `json:"type"`
		} `json:"__metadata"`
		AttachmentID           string      `json:"attachmentId"`
		Country                string      `json:"country"`
		FileName               string      `json:"fileName"`
		LastModifiedDateTime   string      `json:"lastModifiedDateTime"`
		OwnerIDType            string      `json:"ownerIdType"`
		DocumentType           string      `json:"documentType"`
		Deletable              bool        `json:"deletable"`
		Description            string      `json:"description"`
		Language               string      `json:"language"`
		MimeType               string      `json:"mimeType"`
		ModuleCategory         string      `json:"moduleCategory"`
		OwnerID                string      `json:"ownerId"`
		Deprecable             bool        `json:"deprecable"`
		PiiFlag                int         `json:"piiFlag"`
		FileExtension          string      `json:"fileExtension"`
		DocumentEntityID       string      `json:"documentEntityId"`
		Module                 string      `json:"module"`
		DocumentCategory       string      `json:"documentCategory"`
		ExternalID             string      `json:"externalId"`
		UserID                 string      `json:"userId"`
		Searchable             bool        `json:"searchable"`
		DocumentEntityType     string      `json:"documentEntityType"`
		CreatedDate            string      `json:"createdDate"`
		Viewable               bool        `json:"viewable"`
		FileSize               int         `json:"fileSize"`
		SoftDelete             bool        `json:"softDelete"`
		LastAccessed           string      `json:"lastAccessed"`
		FileContent            string      `json:"fileContent"`
		ImageConvertInProgress bool        `json:"imageConvertInProgress"`
		UserNav                struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"userNav"`
		CandidateNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"candidateNav"`
	} `json:"d"`
}
