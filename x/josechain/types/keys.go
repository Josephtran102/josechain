package types

const (
	// ModuleName defines the module name
	ModuleName = "josechain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_josechain"
)

var (
	ParamsKey = []byte("p_josechain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
