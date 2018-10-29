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
)
