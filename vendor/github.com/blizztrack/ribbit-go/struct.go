package ribbit

type (
	SummaryItem struct {
		Product string
		Seqn    string
		Flags   string
	}

	RegionItem struct {
		Buildconfig   string
		Buildid       string
		Cdnconfig     string
		Keyring       string
		Region        string
		Versionsname  string
		Productconfig string
	}

	CdnItem struct {
		Name        string `json:"-,omitempty"`
		Region      string
		Path        string
		Hosts       string `json:"-,omitempty"`
		HostsList   []string
		Servers     string `json:"-,omitempty"`
		ServersList []string
		ConfigPath  string
	}
)
