package criteria

// SdnMode Mode for criteria for SdnRepository
type SdnMode int

const (
	// SdnModeWeak use like
	SdnModeWeak SdnMode = iota
	// SdnModeStrong use equals
	SdnModeStrong
)

// SdnCriteria are params for getting Sdn from a repository
type SdnCriteria struct {
	MaybeFirstName string
	MaybeLastName  string
	Mode           SdnMode
	Limit          int
}
