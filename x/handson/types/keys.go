package types

const (
	// ModuleName defines the module name
	ModuleName = "handson"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_handson"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	PoolKey      = "Pool/value/"
	PoolCountKey = "Pool/count/"
)

const (
	PoolEventId     = "id"
	PoolEventAmount = "amount"
	PoolEventDenom  = "denom"

	PoolCreatedEventType    = "pool-added"
	PoolDepositedEventType  = "pool-deposited"
	PoolWithdrawedEventType = "pool-withdrawed"
)
