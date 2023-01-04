package internal

type KeyParam struct {
	Nonce string `json:"nonce"`
	Tag   string `json:"tag"`
}

type Slot struct {
	Type      int      `json:"type"`
	UUID      string   `json:"uuid"`
	Key       string   `json:"key"`
	KeyParams KeyParam `json:"key_params"`
	N         int      `json:"n"`
	R         int      `json:"r"`
	P         int      `json:"p"`
	Salt      string   `json:"salt"`
	Repaired  bool     `json:"repaired"`
}

type VaultHeader struct {
	Slots  []Slot   `json:"slots"`
	Params KeyParam `json:"params"`
}

type Vault struct {
	Version int         `json:"version"`
	Header  VaultHeader `json:"header"`
	DB      string      `json:"db"`
}

type DBEntryInfo struct {
	Secret string `json:"secret"`
	Algo   string `json:"algo"`
	Digits int    `json:"digits"`
	Period int    `json:"period"`
}

type DBEntry struct {
	Type   string      `json:"type"`
	Name   string      `json:"name"`
	Issuer string      `json:"issuer"`
	Group  string      `json:"group"`
	Info   DBEntryInfo `json:"info"`
}

type DB struct {
	Version int       `json:"version"`
	Entries []DBEntry `json:"entries"`
}
