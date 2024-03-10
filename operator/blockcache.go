package operator

type BlockCache interface {
	Get(height uint64) (BlockDef, bool)
	Set(height uint64, blockDef BlockDef)
}
type BlockDef interface{}
