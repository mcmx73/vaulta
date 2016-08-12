package tapeencript

import (
	"github.com/mcmx73/vaulta/components/store"
	"github.com/DeepForestTeam/mobiussign/components/config"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"crypto/rand"
	"errors"
	"fmt"
)

const (
	TapeStore = "Blocks"
)

var (
	ErrDuplicateId = errors.New("duplicate block id. can not save")
	ErrTooLongBlock = errors.New("can not encode data: block too long")
)

type TapeBlock struct {
	BlockId   string
	BlockData []byte
}

func (this *TapeBlock)Save() (err error) {
	dumb := TapeBlock{}
	err = store.Get(TapeStore, this.BlockId, &dumb)
	if err != store.ErrKeyNotFound && err != store.ErrSectionNotFound {
		return ErrDuplicateId
	}
	_, err = store.Set(TapeStore, this.BlockId, this)
	return
}

func (this *TapeBlock) Load() (err error) {
	err = store.Get(TapeStore, this.BlockId, this)
	return
}

func (this *TapeBlock)Encode() (new_key string, err error) {
	max_block_size, err := config.GetInt64("ENC_BLOCK_SIZE")
	if err != nil {
		max_block_size = 16384
	}
	if len(this.BlockData) > int(max_block_size) {
		return new_key, ErrTooLongBlock
	}

	return
}

func (this *TapeBlock)FillRandom(count int) {
	for {
		this.BlockId = generateRandKey()
		generateRandBlock(&this.BlockData)
		err := this.Save()
		if err != nil {
			log.Error(err)
		}
		count = count - 1
		if count == 0 {
			return
		}
	}
}

func (this *TapeBlock)TotalCoutunt() (count int64) {
	count, _ = store.Count(TapeStore)
	return
}

func generateRandKey() (pepper string) {
	pepper_bytes := make([]byte, 8)
	rand.Read(pepper_bytes)
	pepper = fmt.Sprintf("%x", pepper_bytes)
	return
}
func generateRandBlock(block *[]byte) {
	max_block_size, err := config.GetInt64("ENC_BLOCK_SIZE")
	if err != nil {
		max_block_size = 16384
	}
	*block = make([]byte, max_block_size)
	rand.Read(*block)

}