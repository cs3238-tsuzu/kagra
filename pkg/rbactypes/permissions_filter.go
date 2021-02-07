package rbactypes

// PermissionsFilter should return if the Permission will be in the filtered Permissions
type PermissionsFilter func(p *Permission) bool

var (
	// FilterDenyingAll filters Permission denying all(not matter)
	FilterDenyingAll = func(p *Permission) bool {
		return !p.DeniesAll()
	}
)
