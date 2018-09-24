package core

import (
	"github.com/elastos/Elastos.ELA.SideChain/vm/interfaces"
	"github.com/elastos/Elastos.ELA.Utility/common"
)

var TransactionHelper ITransactionHelper

type ITransactionHelper interface {
	Name(txType TransactionType) string
	GetDataContainer(programHash *common.Uint168, tx *Transaction) interfaces.IDataContainer
}

type TransactionHelperBase struct {
}

func InitTransactionHelper() {
	TransactionHelper = &TransactionHelperBase{}
}

func (t *TransactionHelperBase) Name(txType TransactionType) string {
	switch txType {
	case CoinBase:
		return "CoinBase"
	case RegisterAsset:
		return "RegisterAsset"
	case TransferAsset:
		return "TransferAsset"
	case Record:
		return "Record"
	case Deploy:
		return "Deploy"
	case SideChainPow:
		return "SideChainPow"
	case RechargeToSideChain:
		return "RechargeToSideChain"
	case WithdrawFromSideChain:
		return "WithdrawFromSideChain"
	case TransferCrossChainAsset:
		return "TransferCrossChainAsset"
	default:
		return "Unknown"
	}
}

func (t *TransactionHelperBase) GetDataContainer(programHash *common.Uint168, tx *Transaction) interfaces.IDataContainer {
	return tx
}
