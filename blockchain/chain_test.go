package blockchain

import (
	"github.com/nomadcoders/nomadcoin/utils"
	"sync"
	"testing"
)

type fakeDB struct {
	fakeFindBlock func() []byte
	fakeLoadChain func() []byte
}

func (f fakeDB) FindBlock(hash string) []byte {
	return f.fakeFindBlock()
}
func (f fakeDB) LoadChain() []byte {
	return f.fakeLoadChain()
}
func (fakeDB) SaveBlock(hash string, data []byte) {

}
func (fakeDB) SaveChain(data []byte) {

}
func (fakeDB) DeleteAllBlocks() {

}

func TestBlockchain(t *testing.T) {
	t.Run("Should create blockchain", func(t *testing.T) {
		dbStorage = fakeDB{
			fakeLoadChain: func() []byte {
				return nil
			},
		}
		bc := Blockchain()
		if bc.Height != 1 {
			// b.AddBlock Test
			t.Error("Blockchain() should create a blockchain")

		}
	})

	t.Run("Should restore blockchain", func(t *testing.T) {
		once = *new(sync.Once)
		dbStorage = fakeDB{
			fakeLoadChain: func() []byte {
				bc := &blockchain{Height: 1, NewestHash: "xxx", CurrentDifficulty: 1}
				return utils.ToBytes(bc)
			},
		}
		bc := Blockchain()
		if bc.Height != 2 {
			// b.AddBlock Test
			t.Errorf("Blockchain() should restore a blockchain with a height of %d, got %d", 2, bc.Height)

		}
	})
}
