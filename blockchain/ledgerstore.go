package blockchain

import (
	"github.com/elastos/Elastos.ELA.SideChain/core"

	. "github.com/elastos/Elastos.ELA.Utility/common"
)

// IChainStore provides func with store package.
type IChainStore interface {
	SaveBlock(b *core.Block) error
	GetBlock(hash Uint256) (*core.Block, error)
	GetBlockHash(height uint32) (Uint256, error)
	IsDoubleSpend(tx *core.Transaction) bool

	GetHeader(hash Uint256) (*core.Header, error)

	RollbackBlock(hash Uint256) error

	GetTransaction(txId Uint256) (*core.Transaction, uint32, error)
	GetTxReference(tx *core.Transaction) (map[*core.Input]*core.Output, error)

	PersistAsset(assetid Uint256, asset core.Asset) error
	GetAsset(hash Uint256) (*core.Asset, error)

	PersistMainchainTx(mainchainTxHash Uint256)
	GetMainchainTx(mainchainTxHash Uint256) (byte, error)

	GetCurrentBlockHash() Uint256
	GetHeight() uint32

	RemoveHeaderListElement(hash Uint256)

	GetUnspent(txid Uint256, index uint16) (*core.Output, error)
	ContainsUnspent(txid Uint256, index uint16) (bool, error)
	GetUnspentFromProgramHash(programHash Uint168, assetid Uint256) ([]*UTXO, error)
	GetUnspentsFromProgramHash(programHash Uint168) (map[Uint256][]*UTXO, error)
	GetAssets() map[Uint256]*core.Asset

	PersistTrimmedBlock(b *core.Block) error
	RollbackTrimmedBlock(b *core.Block) error
	PersistBlockHash(b *core.Block) error
	RollbackBlockHash(b *core.Block) error
	PersistCurrentBlock(b *core.Block) error
	RollbackCurrentBlock(b *core.Block) error
	PersistUnspendUTXOs(b *core.Block) error
	RollbackUnspendUTXOs(b *core.Block) error
	PersistTransactions(b *core.Block) error
	RollbackTransactions(b *core.Block) error
	RollbackTransaction(txn *core.Transaction) error
	RollbackAsset(assetId Uint256) error
	RollbackMainchainTx(mainchainTxHash Uint256) error
	PersistUnspend(b *core.Block) error
	RollbackUnspend(b *core.Block) error

	IsTxHashDuplicate(txhash Uint256) bool
	IsMainchainTxHashDuplicate(mainchainTxHash Uint256) bool
	IsBlockInStore(hash Uint256) bool
	Close()
}
