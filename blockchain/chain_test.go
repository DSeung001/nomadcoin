package blockchain

import (
	"github.com/nomadcoders/nomadcoin/utils"
	"reflect"
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

func TestBlocks(t *testing.T) {
	fakeBlock := 0
	dbStorage = fakeDB{
		fakeFindBlock: func() []byte {
			var b *Block
			if fakeBlock == 0 {
				b = &Block{
					Height:   2,
					PrevHash: "x",
				}
			} else if fakeBlock == 1 {
				b = &Block{
					Height: 1,
				}
			}
			fakeBlock++

			return utils.ToBytes(b)
		},
	}

	bc := &blockchain{}
	blocks := Blocks(bc)
	if reflect.TypeOf(blocks) != reflect.TypeOf([]*Block{}) {
		t.Error("Blocks() should return a slice of return")
	}
}

func TestFindTx(t *testing.T) {
	t.Run("Tx not found", func(t *testing.T) {
		dbStorage = fakeDB{
			fakeFindBlock: func() []byte {
				b := &Block{
					Height:       2,
					Transactions: []*Tx{},
				}

				return utils.ToBytes(b)
			},
		}

		tx := FindTx(&blockchain{NewestHash: "x"}, "test")
		if tx != nil {
			t.Error("Tx Should be not found")

		}
	})

	t.Run("Tx should be found", func(t *testing.T) {
		dbStorage = fakeDB{
			fakeFindBlock: func() []byte {
				b := &Block{
					Height: 2,
					Transactions: []*Tx{
						{ID: "test"},
					},
				}

				return utils.ToBytes(b)
			},
		}

		tx := FindTx(&blockchain{NewestHash: "x"}, "test")
		if tx == nil {
			t.Error("Tx Should be found")

		}
	})
}