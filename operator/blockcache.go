package operator

type BlockCache interface {
	Get(chain string, height int64) ([]byte, bool)
	Set(chain string, height int64, blockDef []byte)
}
