package vault

type Vault struct {
	Filename string
	Home     string
}

func NewVault() *Vault {
	return &Vault{
		Filename: "vault.db",
		Home:     ".ev",
	}
}
