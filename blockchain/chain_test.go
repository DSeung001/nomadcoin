package blockchain

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
