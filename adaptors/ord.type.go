package adaptors

type GetInscriptionIdsResponse struct {
	Ids       []string `json:"ids"`
	More      bool     `json:"more"`
	PageIndex int      `json:"page_index"`
}

type GetInscriptionResponse struct {
	Address     string `json:"address"`
	Fee         int64  `json:"fee"`
	Height      int64  `json:"height"`
	Id          string `json:"id"`
	Next        string `json:"next"`
	Previous    string `json:"previous"`
	ContentType string `json:"content_type"`
	Satpoint    string `json:"satpoint"`
	Timestamp   int64  `json:"timestamp"`
	Number      int64  `json:"number"`
	Value       int64  `json:"value"`
}

type GetOutputResponse struct {
	Address      string   `json:"address"`
	Indexed      bool     `json:"indexed"`
	Inscriptions []string `json:"inscriptions"`
	ScriptPubkey string   `json:"script_pubkey"`
	Spent        bool     `json:"spent"`
	Transaction  string   `json:"transaction"`
	Value        int64    `json:"value"`
}

type GetContentResponse struct {
	Protocol string `json:"p"`
	Action   string `json:"op"`
	Tick     string `json:"tick"`
	Amount   string `json:"amt"`
}
