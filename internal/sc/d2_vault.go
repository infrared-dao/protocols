// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sc

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// EnforcedOptionParam is an auto generated low-level Go binding around an user-defined struct.
type EnforcedOptionParam struct {
	Eid     uint32
	MsgType uint16
	Options []byte
}

// InboundPacket is an auto generated low-level Go binding around an user-defined struct.
type InboundPacket struct {
	Origin    Origin
	DstEid    uint32
	Receiver  common.Address
	Guid      [32]byte
	Value     *big.Int
	Executor  common.Address
	Message   []byte
	ExtraData []byte
}

// MessagingFee is an auto generated low-level Go binding around an user-defined struct.
type MessagingFee struct {
	NativeFee  *big.Int
	LzTokenFee *big.Int
}

// MessagingReceipt is an auto generated low-level Go binding around an user-defined struct.
type MessagingReceipt struct {
	Guid  [32]byte
	Nonce uint64
	Fee   MessagingFee
}

// OFTFeeDetail is an auto generated low-level Go binding around an user-defined struct.
type OFTFeeDetail struct {
	FeeAmountLD *big.Int
	Description string
}

// OFTLimit is an auto generated low-level Go binding around an user-defined struct.
type OFTLimit struct {
	MinAmountLD *big.Int
	MaxAmountLD *big.Int
}

// OFTReceipt is an auto generated low-level Go binding around an user-defined struct.
type OFTReceipt struct {
	AmountSentLD     *big.Int
	AmountReceivedLD *big.Int
}

// Origin is an auto generated low-level Go binding around an user-defined struct.
type Origin struct {
	SrcEid uint32
	Sender [32]byte
	Nonce  uint64
}

// SendParam is an auto generated low-level Go binding around an user-defined struct.
type SendParam struct {
	DstEid       uint32
	To           [32]byte
	AmountLD     *big.Int
	MinAmountLD  *big.Int
	ExtraOptions []byte
	ComposeMsg   []byte
	OftCmd       []byte
}

// VaultV3ConstructorArgs is an auto generated low-level Go binding around an user-defined struct.
type VaultV3ConstructorArgs struct {
	Strategy         common.Address
	Owner            common.Address
	Trader           common.Address
	Depositor        common.Address
	Asset            common.Address
	Name             string
	Symbol           string
	MaxDeposits      *big.Int
	WhitelistAsset   common.Address
	WhitelistBalance *big.Int
	LzEndpoint       common.Address
	DateDeposits     *big.Int
	DateTrading      *big.Int
	DateEnd          *big.Int
}

// VaultV3Epoch is an auto generated low-level Go binding around an user-defined struct.
type VaultV3Epoch struct {
	FundingStart *big.Int
	EpochStart   *big.Int
	EpochEnd     *big.Int
}

// D2VaultMetaData contains all meta data concerning the D2Vault contract.
var D2VaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_maxDeposits\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_whitelistAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_whitelistBalance\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_lzEndpoint\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"dateDeposits\",\"type\":\"uint80\"},{\"internalType\":\"uint80\",\"name\":\"dateTrading\",\"type\":\"uint80\"},{\"internalType\":\"uint80\",\"name\":\"dateEnd\",\"type\":\"uint80\"}],\"internalType\":\"structVaultV3.ConstructorArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidDelegate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEndpointCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLocalDecimals\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"InvalidOptions\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LzTokenUnavailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"}],\"name\":\"NoPeer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"}],\"name\":\"NotEnoughNative\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"OnlyEndpoint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"}],\"name\":\"OnlyPeer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySelf\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"SimulationResult\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"}],\"name\":\"SlippageExceeded\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"msgType\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structEnforcedOptionParam[]\",\"name\":\"_enforcedOptions\",\"type\":\"tuple[]\"}],\"name\":\"EnforcedOptionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fundingStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochEnd\",\"type\":\"uint256\"}],\"name\":\"EpochStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsCustodied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsReturned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inspector\",\"type\":\"address\"}],\"name\":\"MsgInspectorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMax\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMax\",\"type\":\"uint256\"}],\"name\":\"NewMaxDeposits\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"NewWhitelistStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountReceivedLD\",\"type\":\"uint256\"}],\"name\":\"OFTReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSentLD\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountReceivedLD\",\"type\":\"uint256\"}],\"name\":\"OFTSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"peer\",\"type\":\"bytes32\"}],\"name\":\"PeerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"preCrimeAddress\",\"type\":\"address\"}],\"name\":\"PreCrimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_EPOCH_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_FUNDING_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEND\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEND_AND_CALL\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"}],\"name\":\"allowInitializePath\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"approvalRequired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"_msgType\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_extraOptions\",\"type\":\"bytes\"}],\"name\":\"combineOptions\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"custodied\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"custodiedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"custodyFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimalConversionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"emergencyEndEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyFreeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_fundingStart\",\"type\":\"uint80\"},{\"internalType\":\"uint80\",\"name\":\"_epochStart\",\"type\":\"uint80\"},{\"internalType\":\"uint80\",\"name\":\"_epochEnd\",\"type\":\"uint80\"}],\"name\":\"emergencyStartEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endpoint\",\"outputs\":[{\"internalType\":\"contractILayerZeroEndpointV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"msgType\",\"type\":\"uint16\"}],\"name\":\"enforcedOptions\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"enforcedOption\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochs\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"fundingStart\",\"type\":\"uint80\"},{\"internalType\":\"uint80\",\"name\":\"epochStart\",\"type\":\"uint80\"},{\"internalType\":\"uint80\",\"name\":\"epochEnd\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"frozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEpochInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint80\",\"name\":\"fundingStart\",\"type\":\"uint80\"},{\"internalType\":\"uint80\",\"name\":\"epochStart\",\"type\":\"uint80\"},{\"internalType\":\"uint80\",\"name\":\"epochEnd\",\"type\":\"uint80\"}],\"internalType\":\"structVaultV3.Epoch\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_maxDeposits\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_whitelistAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_whitelistBalance\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_lzEndpoint\",\"type\":\"address\"},{\"internalType\":\"uint80\",\"name\":\"dateDeposits\",\"type\":\"uint80\"},{\"internalType\":\"uint80\",\"name\":\"dateTrading\",\"type\":\"uint80\"},{\"internalType\":\"uint80\",\"name\":\"dateEnd\",\"type\":\"uint80\"}],\"internalType\":\"structVaultV3.ConstructorArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"isComposeMsgSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFunding\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_peer\",\"type\":\"bytes32\"}],\"name\":\"isPeer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"lzReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structInboundPacket[]\",\"name\":\"_packets\",\"type\":\"tuple[]\"}],\"name\":\"lzReceiveAndRevert\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"lzReceiveSimulate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"msgInspector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nextNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oApp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oAppVersion\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"senderVersion\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"receiverVersion\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oftVersion\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"}],\"name\":\"peers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"peer\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preCrime\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraOptions\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"composeMsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"oftCmd\",\"type\":\"bytes\"}],\"internalType\":\"structSendParam\",\"name\":\"_sendParam\",\"type\":\"tuple\"}],\"name\":\"quoteOFT\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountLD\",\"type\":\"uint256\"}],\"internalType\":\"structOFTLimit\",\"name\":\"oftLimit\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"feeAmountLD\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structOFTFeeDetail[]\",\"name\":\"oftFeeDetails\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountSentLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountReceivedLD\",\"type\":\"uint256\"}],\"internalType\":\"structOFTReceipt\",\"name\":\"oftReceipt\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraOptions\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"composeMsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"oftCmd\",\"type\":\"bytes\"}],\"internalType\":\"structSendParam\",\"name\":\"_sendParam\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"_payInLzToken\",\"type\":\"bool\"}],\"name\":\"quoteSend\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"msgFee\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"returnFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraOptions\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"composeMsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"oftCmd\",\"type\":\"bytes\"}],\"internalType\":\"structSendParam\",\"name\":\"_sendParam\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"_fee\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"send\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"fee\",\"type\":\"tuple\"}],\"internalType\":\"structMessagingReceipt\",\"name\":\"msgReceipt\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountSentLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountReceivedLD\",\"type\":\"uint256\"}],\"internalType\":\"structOFTReceipt\",\"name\":\"oftReceipt\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"name\":\"setDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"msgType\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structEnforcedOptionParam[]\",\"name\":\"_enforcedOptions\",\"type\":\"tuple[]\"}],\"name\":\"setEnforcedOptions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newMax\",\"type\":\"uint256\"}],\"name\":\"setMaxDeposits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_msgInspector\",\"type\":\"address\"}],\"name\":\"setMsgInspector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_peer\",\"type\":\"bytes32\"}],\"name\":\"setPeer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_preCrime\",\"type\":\"address\"}],\"name\":\"setPreCrime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"}],\"name\":\"setTrader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_whitelistAsset\",\"type\":\"address\"}],\"name\":\"setWhitelistAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_whitelistBalance\",\"type\":\"uint256\"}],\"name\":\"setWhitelistBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setWhitelistStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_users\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"_statuses\",\"type\":\"bool[]\"}],\"name\":\"setWhitelistStatuses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sharedDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_fundingStart\",\"type\":\"uint80\"},{\"internalType\":\"uint80\",\"name\":\"_epochStart\",\"type\":\"uint80\"},{\"internalType\":\"uint80\",\"name\":\"_epochEnd\",\"type\":\"uint80\"}],\"name\":\"startEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"started\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trader\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// D2VaultABI is the input ABI used to generate the binding from.
// Deprecated: Use D2VaultMetaData.ABI instead.
var D2VaultABI = D2VaultMetaData.ABI

// D2Vault is an auto generated Go binding around an Ethereum contract.
type D2Vault struct {
	D2VaultCaller     // Read-only binding to the contract
	D2VaultTransactor // Write-only binding to the contract
	D2VaultFilterer   // Log filterer for contract events
}

// D2VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type D2VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// D2VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type D2VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// D2VaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type D2VaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// D2VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type D2VaultSession struct {
	Contract     *D2Vault          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// D2VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type D2VaultCallerSession struct {
	Contract *D2VaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// D2VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type D2VaultTransactorSession struct {
	Contract     *D2VaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// D2VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type D2VaultRaw struct {
	Contract *D2Vault // Generic contract binding to access the raw methods on
}

// D2VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type D2VaultCallerRaw struct {
	Contract *D2VaultCaller // Generic read-only contract binding to access the raw methods on
}

// D2VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type D2VaultTransactorRaw struct {
	Contract *D2VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewD2Vault creates a new instance of D2Vault, bound to a specific deployed contract.
func NewD2Vault(address common.Address, backend bind.ContractBackend) (*D2Vault, error) {
	contract, err := bindD2Vault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &D2Vault{D2VaultCaller: D2VaultCaller{contract: contract}, D2VaultTransactor: D2VaultTransactor{contract: contract}, D2VaultFilterer: D2VaultFilterer{contract: contract}}, nil
}

// NewD2VaultCaller creates a new read-only instance of D2Vault, bound to a specific deployed contract.
func NewD2VaultCaller(address common.Address, caller bind.ContractCaller) (*D2VaultCaller, error) {
	contract, err := bindD2Vault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &D2VaultCaller{contract: contract}, nil
}

// NewD2VaultTransactor creates a new write-only instance of D2Vault, bound to a specific deployed contract.
func NewD2VaultTransactor(address common.Address, transactor bind.ContractTransactor) (*D2VaultTransactor, error) {
	contract, err := bindD2Vault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &D2VaultTransactor{contract: contract}, nil
}

// NewD2VaultFilterer creates a new log filterer instance of D2Vault, bound to a specific deployed contract.
func NewD2VaultFilterer(address common.Address, filterer bind.ContractFilterer) (*D2VaultFilterer, error) {
	contract, err := bindD2Vault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &D2VaultFilterer{contract: contract}, nil
}

// bindD2Vault binds a generic wrapper to an already deployed contract.
func bindD2Vault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := D2VaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_D2Vault *D2VaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _D2Vault.Contract.D2VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_D2Vault *D2VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _D2Vault.Contract.D2VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_D2Vault *D2VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _D2Vault.Contract.D2VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_D2Vault *D2VaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _D2Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_D2Vault *D2VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _D2Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_D2Vault *D2VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _D2Vault.Contract.contract.Transact(opts, method, params...)
}

// MAXEPOCHDURATION is a free data retrieval call binding the contract method 0x4f17a2fa.
//
// Solidity: function MAX_EPOCH_DURATION() view returns(uint256)
func (_D2Vault *D2VaultCaller) MAXEPOCHDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "MAX_EPOCH_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXEPOCHDURATION is a free data retrieval call binding the contract method 0x4f17a2fa.
//
// Solidity: function MAX_EPOCH_DURATION() view returns(uint256)
func (_D2Vault *D2VaultSession) MAXEPOCHDURATION() (*big.Int, error) {
	return _D2Vault.Contract.MAXEPOCHDURATION(&_D2Vault.CallOpts)
}

// MAXEPOCHDURATION is a free data retrieval call binding the contract method 0x4f17a2fa.
//
// Solidity: function MAX_EPOCH_DURATION() view returns(uint256)
func (_D2Vault *D2VaultCallerSession) MAXEPOCHDURATION() (*big.Int, error) {
	return _D2Vault.Contract.MAXEPOCHDURATION(&_D2Vault.CallOpts)
}

// MINFUNDINGDURATION is a free data retrieval call binding the contract method 0x86ee6b93.
//
// Solidity: function MIN_FUNDING_DURATION() view returns(uint256)
func (_D2Vault *D2VaultCaller) MINFUNDINGDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "MIN_FUNDING_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINFUNDINGDURATION is a free data retrieval call binding the contract method 0x86ee6b93.
//
// Solidity: function MIN_FUNDING_DURATION() view returns(uint256)
func (_D2Vault *D2VaultSession) MINFUNDINGDURATION() (*big.Int, error) {
	return _D2Vault.Contract.MINFUNDINGDURATION(&_D2Vault.CallOpts)
}

// MINFUNDINGDURATION is a free data retrieval call binding the contract method 0x86ee6b93.
//
// Solidity: function MIN_FUNDING_DURATION() view returns(uint256)
func (_D2Vault *D2VaultCallerSession) MINFUNDINGDURATION() (*big.Int, error) {
	return _D2Vault.Contract.MINFUNDINGDURATION(&_D2Vault.CallOpts)
}

// SEND is a free data retrieval call binding the contract method 0x1f5e1334.
//
// Solidity: function SEND() view returns(uint16)
func (_D2Vault *D2VaultCaller) SEND(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "SEND")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// SEND is a free data retrieval call binding the contract method 0x1f5e1334.
//
// Solidity: function SEND() view returns(uint16)
func (_D2Vault *D2VaultSession) SEND() (uint16, error) {
	return _D2Vault.Contract.SEND(&_D2Vault.CallOpts)
}

// SEND is a free data retrieval call binding the contract method 0x1f5e1334.
//
// Solidity: function SEND() view returns(uint16)
func (_D2Vault *D2VaultCallerSession) SEND() (uint16, error) {
	return _D2Vault.Contract.SEND(&_D2Vault.CallOpts)
}

// SENDANDCALL is a free data retrieval call binding the contract method 0x134d4f25.
//
// Solidity: function SEND_AND_CALL() view returns(uint16)
func (_D2Vault *D2VaultCaller) SENDANDCALL(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "SEND_AND_CALL")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// SENDANDCALL is a free data retrieval call binding the contract method 0x134d4f25.
//
// Solidity: function SEND_AND_CALL() view returns(uint16)
func (_D2Vault *D2VaultSession) SENDANDCALL() (uint16, error) {
	return _D2Vault.Contract.SENDANDCALL(&_D2Vault.CallOpts)
}

// SENDANDCALL is a free data retrieval call binding the contract method 0x134d4f25.
//
// Solidity: function SEND_AND_CALL() view returns(uint16)
func (_D2Vault *D2VaultCallerSession) SENDANDCALL() (uint16, error) {
	return _D2Vault.Contract.SENDANDCALL(&_D2Vault.CallOpts)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_D2Vault *D2VaultCaller) AllowInitializePath(opts *bind.CallOpts, origin Origin) (bool, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "allowInitializePath", origin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_D2Vault *D2VaultSession) AllowInitializePath(origin Origin) (bool, error) {
	return _D2Vault.Contract.AllowInitializePath(&_D2Vault.CallOpts, origin)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_D2Vault *D2VaultCallerSession) AllowInitializePath(origin Origin) (bool, error) {
	return _D2Vault.Contract.AllowInitializePath(&_D2Vault.CallOpts, origin)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_D2Vault *D2VaultCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_D2Vault *D2VaultSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _D2Vault.Contract.Allowance(&_D2Vault.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_D2Vault *D2VaultCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _D2Vault.Contract.Allowance(&_D2Vault.CallOpts, owner, spender)
}

// ApprovalRequired is a free data retrieval call binding the contract method 0x9f68b964.
//
// Solidity: function approvalRequired() pure returns(bool)
func (_D2Vault *D2VaultCaller) ApprovalRequired(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "approvalRequired")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovalRequired is a free data retrieval call binding the contract method 0x9f68b964.
//
// Solidity: function approvalRequired() pure returns(bool)
func (_D2Vault *D2VaultSession) ApprovalRequired() (bool, error) {
	return _D2Vault.Contract.ApprovalRequired(&_D2Vault.CallOpts)
}

// ApprovalRequired is a free data retrieval call binding the contract method 0x9f68b964.
//
// Solidity: function approvalRequired() pure returns(bool)
func (_D2Vault *D2VaultCallerSession) ApprovalRequired() (bool, error) {
	return _D2Vault.Contract.ApprovalRequired(&_D2Vault.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_D2Vault *D2VaultCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_D2Vault *D2VaultSession) Asset() (common.Address, error) {
	return _D2Vault.Contract.Asset(&_D2Vault.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_D2Vault *D2VaultCallerSession) Asset() (common.Address, error) {
	return _D2Vault.Contract.Asset(&_D2Vault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_D2Vault *D2VaultCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_D2Vault *D2VaultSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _D2Vault.Contract.BalanceOf(&_D2Vault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_D2Vault *D2VaultCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _D2Vault.Contract.BalanceOf(&_D2Vault.CallOpts, account)
}

// CombineOptions is a free data retrieval call binding the contract method 0xbc70b354.
//
// Solidity: function combineOptions(uint32 _eid, uint16 _msgType, bytes _extraOptions) view returns(bytes)
func (_D2Vault *D2VaultCaller) CombineOptions(opts *bind.CallOpts, _eid uint32, _msgType uint16, _extraOptions []byte) ([]byte, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "combineOptions", _eid, _msgType, _extraOptions)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CombineOptions is a free data retrieval call binding the contract method 0xbc70b354.
//
// Solidity: function combineOptions(uint32 _eid, uint16 _msgType, bytes _extraOptions) view returns(bytes)
func (_D2Vault *D2VaultSession) CombineOptions(_eid uint32, _msgType uint16, _extraOptions []byte) ([]byte, error) {
	return _D2Vault.Contract.CombineOptions(&_D2Vault.CallOpts, _eid, _msgType, _extraOptions)
}

// CombineOptions is a free data retrieval call binding the contract method 0xbc70b354.
//
// Solidity: function combineOptions(uint32 _eid, uint16 _msgType, bytes _extraOptions) view returns(bytes)
func (_D2Vault *D2VaultCallerSession) CombineOptions(_eid uint32, _msgType uint16, _extraOptions []byte) ([]byte, error) {
	return _D2Vault.Contract.CombineOptions(&_D2Vault.CallOpts, _eid, _msgType, _extraOptions)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_D2Vault *D2VaultCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_D2Vault *D2VaultSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _D2Vault.Contract.ConvertToAssets(&_D2Vault.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_D2Vault *D2VaultCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _D2Vault.Contract.ConvertToAssets(&_D2Vault.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_D2Vault *D2VaultCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_D2Vault *D2VaultSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _D2Vault.Contract.ConvertToShares(&_D2Vault.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_D2Vault *D2VaultCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _D2Vault.Contract.ConvertToShares(&_D2Vault.CallOpts, assets)
}

// Custodied is a free data retrieval call binding the contract method 0x103d113f.
//
// Solidity: function custodied() view returns(bool)
func (_D2Vault *D2VaultCaller) Custodied(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "custodied")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Custodied is a free data retrieval call binding the contract method 0x103d113f.
//
// Solidity: function custodied() view returns(bool)
func (_D2Vault *D2VaultSession) Custodied() (bool, error) {
	return _D2Vault.Contract.Custodied(&_D2Vault.CallOpts)
}

// Custodied is a free data retrieval call binding the contract method 0x103d113f.
//
// Solidity: function custodied() view returns(bool)
func (_D2Vault *D2VaultCallerSession) Custodied() (bool, error) {
	return _D2Vault.Contract.Custodied(&_D2Vault.CallOpts)
}

// CustodiedAmount is a free data retrieval call binding the contract method 0xd9cd2bcc.
//
// Solidity: function custodiedAmount() view returns(uint256)
func (_D2Vault *D2VaultCaller) CustodiedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "custodiedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CustodiedAmount is a free data retrieval call binding the contract method 0xd9cd2bcc.
//
// Solidity: function custodiedAmount() view returns(uint256)
func (_D2Vault *D2VaultSession) CustodiedAmount() (*big.Int, error) {
	return _D2Vault.Contract.CustodiedAmount(&_D2Vault.CallOpts)
}

// CustodiedAmount is a free data retrieval call binding the contract method 0xd9cd2bcc.
//
// Solidity: function custodiedAmount() view returns(uint256)
func (_D2Vault *D2VaultCallerSession) CustodiedAmount() (*big.Int, error) {
	return _D2Vault.Contract.CustodiedAmount(&_D2Vault.CallOpts)
}

// DecimalConversionRate is a free data retrieval call binding the contract method 0x963efcaa.
//
// Solidity: function decimalConversionRate() view returns(uint256)
func (_D2Vault *D2VaultCaller) DecimalConversionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "decimalConversionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecimalConversionRate is a free data retrieval call binding the contract method 0x963efcaa.
//
// Solidity: function decimalConversionRate() view returns(uint256)
func (_D2Vault *D2VaultSession) DecimalConversionRate() (*big.Int, error) {
	return _D2Vault.Contract.DecimalConversionRate(&_D2Vault.CallOpts)
}

// DecimalConversionRate is a free data retrieval call binding the contract method 0x963efcaa.
//
// Solidity: function decimalConversionRate() view returns(uint256)
func (_D2Vault *D2VaultCallerSession) DecimalConversionRate() (*big.Int, error) {
	return _D2Vault.Contract.DecimalConversionRate(&_D2Vault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_D2Vault *D2VaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_D2Vault *D2VaultSession) Decimals() (uint8, error) {
	return _D2Vault.Contract.Decimals(&_D2Vault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_D2Vault *D2VaultCallerSession) Decimals() (uint8, error) {
	return _D2Vault.Contract.Decimals(&_D2Vault.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_D2Vault *D2VaultCaller) Endpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "endpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_D2Vault *D2VaultSession) Endpoint() (common.Address, error) {
	return _D2Vault.Contract.Endpoint(&_D2Vault.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_D2Vault *D2VaultCallerSession) Endpoint() (common.Address, error) {
	return _D2Vault.Contract.Endpoint(&_D2Vault.CallOpts)
}

// EnforcedOptions is a free data retrieval call binding the contract method 0x5535d461.
//
// Solidity: function enforcedOptions(uint32 eid, uint16 msgType) view returns(bytes enforcedOption)
func (_D2Vault *D2VaultCaller) EnforcedOptions(opts *bind.CallOpts, eid uint32, msgType uint16) ([]byte, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "enforcedOptions", eid, msgType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EnforcedOptions is a free data retrieval call binding the contract method 0x5535d461.
//
// Solidity: function enforcedOptions(uint32 eid, uint16 msgType) view returns(bytes enforcedOption)
func (_D2Vault *D2VaultSession) EnforcedOptions(eid uint32, msgType uint16) ([]byte, error) {
	return _D2Vault.Contract.EnforcedOptions(&_D2Vault.CallOpts, eid, msgType)
}

// EnforcedOptions is a free data retrieval call binding the contract method 0x5535d461.
//
// Solidity: function enforcedOptions(uint32 eid, uint16 msgType) view returns(bytes enforcedOption)
func (_D2Vault *D2VaultCallerSession) EnforcedOptions(eid uint32, msgType uint16) ([]byte, error) {
	return _D2Vault.Contract.EnforcedOptions(&_D2Vault.CallOpts, eid, msgType)
}

// Epochs is a free data retrieval call binding the contract method 0xc6b61e4c.
//
// Solidity: function epochs(uint256 ) view returns(uint80 fundingStart, uint80 epochStart, uint80 epochEnd)
func (_D2Vault *D2VaultCaller) Epochs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	FundingStart *big.Int
	EpochStart   *big.Int
	EpochEnd     *big.Int
}, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "epochs", arg0)

	outstruct := new(struct {
		FundingStart *big.Int
		EpochStart   *big.Int
		EpochEnd     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FundingStart = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EpochStart = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EpochEnd = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Epochs is a free data retrieval call binding the contract method 0xc6b61e4c.
//
// Solidity: function epochs(uint256 ) view returns(uint80 fundingStart, uint80 epochStart, uint80 epochEnd)
func (_D2Vault *D2VaultSession) Epochs(arg0 *big.Int) (struct {
	FundingStart *big.Int
	EpochStart   *big.Int
	EpochEnd     *big.Int
}, error) {
	return _D2Vault.Contract.Epochs(&_D2Vault.CallOpts, arg0)
}

// Epochs is a free data retrieval call binding the contract method 0xc6b61e4c.
//
// Solidity: function epochs(uint256 ) view returns(uint80 fundingStart, uint80 epochStart, uint80 epochEnd)
func (_D2Vault *D2VaultCallerSession) Epochs(arg0 *big.Int) (struct {
	FundingStart *big.Int
	EpochStart   *big.Int
	EpochEnd     *big.Int
}, error) {
	return _D2Vault.Contract.Epochs(&_D2Vault.CallOpts, arg0)
}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() view returns(bool)
func (_D2Vault *D2VaultCaller) Frozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "frozen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() view returns(bool)
func (_D2Vault *D2VaultSession) Frozen() (bool, error) {
	return _D2Vault.Contract.Frozen(&_D2Vault.CallOpts)
}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() view returns(bool)
func (_D2Vault *D2VaultCallerSession) Frozen() (bool, error) {
	return _D2Vault.Contract.Frozen(&_D2Vault.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_D2Vault *D2VaultCaller) GetCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "getCurrentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_D2Vault *D2VaultSession) GetCurrentEpoch() (*big.Int, error) {
	return _D2Vault.Contract.GetCurrentEpoch(&_D2Vault.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_D2Vault *D2VaultCallerSession) GetCurrentEpoch() (*big.Int, error) {
	return _D2Vault.Contract.GetCurrentEpoch(&_D2Vault.CallOpts)
}

// GetCurrentEpochInfo is a free data retrieval call binding the contract method 0xbabc394f.
//
// Solidity: function getCurrentEpochInfo() view returns((uint80,uint80,uint80))
func (_D2Vault *D2VaultCaller) GetCurrentEpochInfo(opts *bind.CallOpts) (VaultV3Epoch, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "getCurrentEpochInfo")

	if err != nil {
		return *new(VaultV3Epoch), err
	}

	out0 := *abi.ConvertType(out[0], new(VaultV3Epoch)).(*VaultV3Epoch)

	return out0, err

}

// GetCurrentEpochInfo is a free data retrieval call binding the contract method 0xbabc394f.
//
// Solidity: function getCurrentEpochInfo() view returns((uint80,uint80,uint80))
func (_D2Vault *D2VaultSession) GetCurrentEpochInfo() (VaultV3Epoch, error) {
	return _D2Vault.Contract.GetCurrentEpochInfo(&_D2Vault.CallOpts)
}

// GetCurrentEpochInfo is a free data retrieval call binding the contract method 0xbabc394f.
//
// Solidity: function getCurrentEpochInfo() view returns((uint80,uint80,uint80))
func (_D2Vault *D2VaultCallerSession) GetCurrentEpochInfo() (VaultV3Epoch, error) {
	return _D2Vault.Contract.GetCurrentEpochInfo(&_D2Vault.CallOpts)
}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_D2Vault *D2VaultCaller) IsComposeMsgSender(opts *bind.CallOpts, arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "isComposeMsgSender", arg0, arg1, _sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_D2Vault *D2VaultSession) IsComposeMsgSender(arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	return _D2Vault.Contract.IsComposeMsgSender(&_D2Vault.CallOpts, arg0, arg1, _sender)
}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_D2Vault *D2VaultCallerSession) IsComposeMsgSender(arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	return _D2Vault.Contract.IsComposeMsgSender(&_D2Vault.CallOpts, arg0, arg1, _sender)
}

// IsFunding is a free data retrieval call binding the contract method 0x13b53153.
//
// Solidity: function isFunding() view returns(bool)
func (_D2Vault *D2VaultCaller) IsFunding(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "isFunding")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFunding is a free data retrieval call binding the contract method 0x13b53153.
//
// Solidity: function isFunding() view returns(bool)
func (_D2Vault *D2VaultSession) IsFunding() (bool, error) {
	return _D2Vault.Contract.IsFunding(&_D2Vault.CallOpts)
}

// IsFunding is a free data retrieval call binding the contract method 0x13b53153.
//
// Solidity: function isFunding() view returns(bool)
func (_D2Vault *D2VaultCallerSession) IsFunding() (bool, error) {
	return _D2Vault.Contract.IsFunding(&_D2Vault.CallOpts)
}

// IsInEpoch is a free data retrieval call binding the contract method 0x5aaab8ce.
//
// Solidity: function isInEpoch() view returns(bool)
func (_D2Vault *D2VaultCaller) IsInEpoch(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "isInEpoch")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInEpoch is a free data retrieval call binding the contract method 0x5aaab8ce.
//
// Solidity: function isInEpoch() view returns(bool)
func (_D2Vault *D2VaultSession) IsInEpoch() (bool, error) {
	return _D2Vault.Contract.IsInEpoch(&_D2Vault.CallOpts)
}

// IsInEpoch is a free data retrieval call binding the contract method 0x5aaab8ce.
//
// Solidity: function isInEpoch() view returns(bool)
func (_D2Vault *D2VaultCallerSession) IsInEpoch() (bool, error) {
	return _D2Vault.Contract.IsInEpoch(&_D2Vault.CallOpts)
}

// IsPeer is a free data retrieval call binding the contract method 0x5a0dfe4d.
//
// Solidity: function isPeer(uint32 _eid, bytes32 _peer) view returns(bool)
func (_D2Vault *D2VaultCaller) IsPeer(opts *bind.CallOpts, _eid uint32, _peer [32]byte) (bool, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "isPeer", _eid, _peer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPeer is a free data retrieval call binding the contract method 0x5a0dfe4d.
//
// Solidity: function isPeer(uint32 _eid, bytes32 _peer) view returns(bool)
func (_D2Vault *D2VaultSession) IsPeer(_eid uint32, _peer [32]byte) (bool, error) {
	return _D2Vault.Contract.IsPeer(&_D2Vault.CallOpts, _eid, _peer)
}

// IsPeer is a free data retrieval call binding the contract method 0x5a0dfe4d.
//
// Solidity: function isPeer(uint32 _eid, bytes32 _peer) view returns(bool)
func (_D2Vault *D2VaultCallerSession) IsPeer(_eid uint32, _peer [32]byte) (bool, error) {
	return _D2Vault.Contract.IsPeer(&_D2Vault.CallOpts, _eid, _peer)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_D2Vault *D2VaultCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_D2Vault *D2VaultSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _D2Vault.Contract.MaxDeposit(&_D2Vault.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_D2Vault *D2VaultCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _D2Vault.Contract.MaxDeposit(&_D2Vault.CallOpts, arg0)
}

// MaxDeposits is a free data retrieval call binding the contract method 0x354d594b.
//
// Solidity: function maxDeposits() view returns(uint256)
func (_D2Vault *D2VaultCaller) MaxDeposits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "maxDeposits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposits is a free data retrieval call binding the contract method 0x354d594b.
//
// Solidity: function maxDeposits() view returns(uint256)
func (_D2Vault *D2VaultSession) MaxDeposits() (*big.Int, error) {
	return _D2Vault.Contract.MaxDeposits(&_D2Vault.CallOpts)
}

// MaxDeposits is a free data retrieval call binding the contract method 0x354d594b.
//
// Solidity: function maxDeposits() view returns(uint256)
func (_D2Vault *D2VaultCallerSession) MaxDeposits() (*big.Int, error) {
	return _D2Vault.Contract.MaxDeposits(&_D2Vault.CallOpts)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_D2Vault *D2VaultCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_D2Vault *D2VaultSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _D2Vault.Contract.MaxMint(&_D2Vault.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_D2Vault *D2VaultCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _D2Vault.Contract.MaxMint(&_D2Vault.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_D2Vault *D2VaultCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_D2Vault *D2VaultSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _D2Vault.Contract.MaxRedeem(&_D2Vault.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_D2Vault *D2VaultCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _D2Vault.Contract.MaxRedeem(&_D2Vault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_D2Vault *D2VaultCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_D2Vault *D2VaultSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _D2Vault.Contract.MaxWithdraw(&_D2Vault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_D2Vault *D2VaultCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _D2Vault.Contract.MaxWithdraw(&_D2Vault.CallOpts, owner)
}

// MsgInspector is a free data retrieval call binding the contract method 0x111ecdad.
//
// Solidity: function msgInspector() view returns(address)
func (_D2Vault *D2VaultCaller) MsgInspector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "msgInspector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MsgInspector is a free data retrieval call binding the contract method 0x111ecdad.
//
// Solidity: function msgInspector() view returns(address)
func (_D2Vault *D2VaultSession) MsgInspector() (common.Address, error) {
	return _D2Vault.Contract.MsgInspector(&_D2Vault.CallOpts)
}

// MsgInspector is a free data retrieval call binding the contract method 0x111ecdad.
//
// Solidity: function msgInspector() view returns(address)
func (_D2Vault *D2VaultCallerSession) MsgInspector() (common.Address, error) {
	return _D2Vault.Contract.MsgInspector(&_D2Vault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_D2Vault *D2VaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_D2Vault *D2VaultSession) Name() (string, error) {
	return _D2Vault.Contract.Name(&_D2Vault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_D2Vault *D2VaultCallerSession) Name() (string, error) {
	return _D2Vault.Contract.Name(&_D2Vault.CallOpts)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_D2Vault *D2VaultCaller) NextNonce(opts *bind.CallOpts, arg0 uint32, arg1 [32]byte) (uint64, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "nextNonce", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_D2Vault *D2VaultSession) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _D2Vault.Contract.NextNonce(&_D2Vault.CallOpts, arg0, arg1)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_D2Vault *D2VaultCallerSession) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _D2Vault.Contract.NextNonce(&_D2Vault.CallOpts, arg0, arg1)
}

// OApp is a free data retrieval call binding the contract method 0x52ae2879.
//
// Solidity: function oApp() view returns(address)
func (_D2Vault *D2VaultCaller) OApp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "oApp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OApp is a free data retrieval call binding the contract method 0x52ae2879.
//
// Solidity: function oApp() view returns(address)
func (_D2Vault *D2VaultSession) OApp() (common.Address, error) {
	return _D2Vault.Contract.OApp(&_D2Vault.CallOpts)
}

// OApp is a free data retrieval call binding the contract method 0x52ae2879.
//
// Solidity: function oApp() view returns(address)
func (_D2Vault *D2VaultCallerSession) OApp() (common.Address, error) {
	return _D2Vault.Contract.OApp(&_D2Vault.CallOpts)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_D2Vault *D2VaultCaller) OAppVersion(opts *bind.CallOpts) (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "oAppVersion")

	outstruct := new(struct {
		SenderVersion   uint64
		ReceiverVersion uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SenderVersion = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ReceiverVersion = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_D2Vault *D2VaultSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _D2Vault.Contract.OAppVersion(&_D2Vault.CallOpts)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_D2Vault *D2VaultCallerSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _D2Vault.Contract.OAppVersion(&_D2Vault.CallOpts)
}

// OftVersion is a free data retrieval call binding the contract method 0x156a0d0f.
//
// Solidity: function oftVersion() pure returns(bytes4 interfaceId, uint64 version)
func (_D2Vault *D2VaultCaller) OftVersion(opts *bind.CallOpts) (struct {
	InterfaceId [4]byte
	Version     uint64
}, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "oftVersion")

	outstruct := new(struct {
		InterfaceId [4]byte
		Version     uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InterfaceId = *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	outstruct.Version = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// OftVersion is a free data retrieval call binding the contract method 0x156a0d0f.
//
// Solidity: function oftVersion() pure returns(bytes4 interfaceId, uint64 version)
func (_D2Vault *D2VaultSession) OftVersion() (struct {
	InterfaceId [4]byte
	Version     uint64
}, error) {
	return _D2Vault.Contract.OftVersion(&_D2Vault.CallOpts)
}

// OftVersion is a free data retrieval call binding the contract method 0x156a0d0f.
//
// Solidity: function oftVersion() pure returns(bytes4 interfaceId, uint64 version)
func (_D2Vault *D2VaultCallerSession) OftVersion() (struct {
	InterfaceId [4]byte
	Version     uint64
}, error) {
	return _D2Vault.Contract.OftVersion(&_D2Vault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_D2Vault *D2VaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_D2Vault *D2VaultSession) Owner() (common.Address, error) {
	return _D2Vault.Contract.Owner(&_D2Vault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_D2Vault *D2VaultCallerSession) Owner() (common.Address, error) {
	return _D2Vault.Contract.Owner(&_D2Vault.CallOpts)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_D2Vault *D2VaultCaller) Peers(opts *bind.CallOpts, eid uint32) ([32]byte, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "peers", eid)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_D2Vault *D2VaultSession) Peers(eid uint32) ([32]byte, error) {
	return _D2Vault.Contract.Peers(&_D2Vault.CallOpts, eid)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_D2Vault *D2VaultCallerSession) Peers(eid uint32) ([32]byte, error) {
	return _D2Vault.Contract.Peers(&_D2Vault.CallOpts, eid)
}

// PreCrime is a free data retrieval call binding the contract method 0xb731ea0a.
//
// Solidity: function preCrime() view returns(address)
func (_D2Vault *D2VaultCaller) PreCrime(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "preCrime")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PreCrime is a free data retrieval call binding the contract method 0xb731ea0a.
//
// Solidity: function preCrime() view returns(address)
func (_D2Vault *D2VaultSession) PreCrime() (common.Address, error) {
	return _D2Vault.Contract.PreCrime(&_D2Vault.CallOpts)
}

// PreCrime is a free data retrieval call binding the contract method 0xb731ea0a.
//
// Solidity: function preCrime() view returns(address)
func (_D2Vault *D2VaultCallerSession) PreCrime() (common.Address, error) {
	return _D2Vault.Contract.PreCrime(&_D2Vault.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_D2Vault *D2VaultCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_D2Vault *D2VaultSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _D2Vault.Contract.PreviewDeposit(&_D2Vault.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_D2Vault *D2VaultCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _D2Vault.Contract.PreviewDeposit(&_D2Vault.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_D2Vault *D2VaultCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_D2Vault *D2VaultSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _D2Vault.Contract.PreviewMint(&_D2Vault.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_D2Vault *D2VaultCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _D2Vault.Contract.PreviewMint(&_D2Vault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_D2Vault *D2VaultCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_D2Vault *D2VaultSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _D2Vault.Contract.PreviewRedeem(&_D2Vault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_D2Vault *D2VaultCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _D2Vault.Contract.PreviewRedeem(&_D2Vault.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_D2Vault *D2VaultCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_D2Vault *D2VaultSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _D2Vault.Contract.PreviewWithdraw(&_D2Vault.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_D2Vault *D2VaultCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _D2Vault.Contract.PreviewWithdraw(&_D2Vault.CallOpts, assets)
}

// QuoteOFT is a free data retrieval call binding the contract method 0x0d35b415.
//
// Solidity: function quoteOFT((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam) view returns((uint256,uint256) oftLimit, (int256,string)[] oftFeeDetails, (uint256,uint256) oftReceipt)
func (_D2Vault *D2VaultCaller) QuoteOFT(opts *bind.CallOpts, _sendParam SendParam) (struct {
	OftLimit      OFTLimit
	OftFeeDetails []OFTFeeDetail
	OftReceipt    OFTReceipt
}, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "quoteOFT", _sendParam)

	outstruct := new(struct {
		OftLimit      OFTLimit
		OftFeeDetails []OFTFeeDetail
		OftReceipt    OFTReceipt
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OftLimit = *abi.ConvertType(out[0], new(OFTLimit)).(*OFTLimit)
	outstruct.OftFeeDetails = *abi.ConvertType(out[1], new([]OFTFeeDetail)).(*[]OFTFeeDetail)
	outstruct.OftReceipt = *abi.ConvertType(out[2], new(OFTReceipt)).(*OFTReceipt)

	return *outstruct, err

}

// QuoteOFT is a free data retrieval call binding the contract method 0x0d35b415.
//
// Solidity: function quoteOFT((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam) view returns((uint256,uint256) oftLimit, (int256,string)[] oftFeeDetails, (uint256,uint256) oftReceipt)
func (_D2Vault *D2VaultSession) QuoteOFT(_sendParam SendParam) (struct {
	OftLimit      OFTLimit
	OftFeeDetails []OFTFeeDetail
	OftReceipt    OFTReceipt
}, error) {
	return _D2Vault.Contract.QuoteOFT(&_D2Vault.CallOpts, _sendParam)
}

// QuoteOFT is a free data retrieval call binding the contract method 0x0d35b415.
//
// Solidity: function quoteOFT((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam) view returns((uint256,uint256) oftLimit, (int256,string)[] oftFeeDetails, (uint256,uint256) oftReceipt)
func (_D2Vault *D2VaultCallerSession) QuoteOFT(_sendParam SendParam) (struct {
	OftLimit      OFTLimit
	OftFeeDetails []OFTFeeDetail
	OftReceipt    OFTReceipt
}, error) {
	return _D2Vault.Contract.QuoteOFT(&_D2Vault.CallOpts, _sendParam)
}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, bool _payInLzToken) view returns((uint256,uint256) msgFee)
func (_D2Vault *D2VaultCaller) QuoteSend(opts *bind.CallOpts, _sendParam SendParam, _payInLzToken bool) (MessagingFee, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "quoteSend", _sendParam, _payInLzToken)

	if err != nil {
		return *new(MessagingFee), err
	}

	out0 := *abi.ConvertType(out[0], new(MessagingFee)).(*MessagingFee)

	return out0, err

}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, bool _payInLzToken) view returns((uint256,uint256) msgFee)
func (_D2Vault *D2VaultSession) QuoteSend(_sendParam SendParam, _payInLzToken bool) (MessagingFee, error) {
	return _D2Vault.Contract.QuoteSend(&_D2Vault.CallOpts, _sendParam, _payInLzToken)
}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, bool _payInLzToken) view returns((uint256,uint256) msgFee)
func (_D2Vault *D2VaultCallerSession) QuoteSend(_sendParam SendParam, _payInLzToken bool) (MessagingFee, error) {
	return _D2Vault.Contract.QuoteSend(&_D2Vault.CallOpts, _sendParam, _payInLzToken)
}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() view returns(uint8)
func (_D2Vault *D2VaultCaller) SharedDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "sharedDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() view returns(uint8)
func (_D2Vault *D2VaultSession) SharedDecimals() (uint8, error) {
	return _D2Vault.Contract.SharedDecimals(&_D2Vault.CallOpts)
}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() view returns(uint8)
func (_D2Vault *D2VaultCallerSession) SharedDecimals() (uint8, error) {
	return _D2Vault.Contract.SharedDecimals(&_D2Vault.CallOpts)
}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(bool)
func (_D2Vault *D2VaultCaller) Started(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "started")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(bool)
func (_D2Vault *D2VaultSession) Started() (bool, error) {
	return _D2Vault.Contract.Started(&_D2Vault.CallOpts)
}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(bool)
func (_D2Vault *D2VaultCallerSession) Started() (bool, error) {
	return _D2Vault.Contract.Started(&_D2Vault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_D2Vault *D2VaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_D2Vault *D2VaultSession) Symbol() (string, error) {
	return _D2Vault.Contract.Symbol(&_D2Vault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_D2Vault *D2VaultCallerSession) Symbol() (string, error) {
	return _D2Vault.Contract.Symbol(&_D2Vault.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_D2Vault *D2VaultCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_D2Vault *D2VaultSession) Token() (common.Address, error) {
	return _D2Vault.Contract.Token(&_D2Vault.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_D2Vault *D2VaultCallerSession) Token() (common.Address, error) {
	return _D2Vault.Contract.Token(&_D2Vault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_D2Vault *D2VaultCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_D2Vault *D2VaultSession) TotalAssets() (*big.Int, error) {
	return _D2Vault.Contract.TotalAssets(&_D2Vault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_D2Vault *D2VaultCallerSession) TotalAssets() (*big.Int, error) {
	return _D2Vault.Contract.TotalAssets(&_D2Vault.CallOpts)
}

// TotalDeposits is a free data retrieval call binding the contract method 0x7d882097.
//
// Solidity: function totalDeposits() view returns(uint256)
func (_D2Vault *D2VaultCaller) TotalDeposits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "totalDeposits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDeposits is a free data retrieval call binding the contract method 0x7d882097.
//
// Solidity: function totalDeposits() view returns(uint256)
func (_D2Vault *D2VaultSession) TotalDeposits() (*big.Int, error) {
	return _D2Vault.Contract.TotalDeposits(&_D2Vault.CallOpts)
}

// TotalDeposits is a free data retrieval call binding the contract method 0x7d882097.
//
// Solidity: function totalDeposits() view returns(uint256)
func (_D2Vault *D2VaultCallerSession) TotalDeposits() (*big.Int, error) {
	return _D2Vault.Contract.TotalDeposits(&_D2Vault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_D2Vault *D2VaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_D2Vault *D2VaultSession) TotalSupply() (*big.Int, error) {
	return _D2Vault.Contract.TotalSupply(&_D2Vault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_D2Vault *D2VaultCallerSession) TotalSupply() (*big.Int, error) {
	return _D2Vault.Contract.TotalSupply(&_D2Vault.CallOpts)
}

// Trader is a free data retrieval call binding the contract method 0x1758078b.
//
// Solidity: function trader() view returns(address)
func (_D2Vault *D2VaultCaller) Trader(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "trader")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Trader is a free data retrieval call binding the contract method 0x1758078b.
//
// Solidity: function trader() view returns(address)
func (_D2Vault *D2VaultSession) Trader() (common.Address, error) {
	return _D2Vault.Contract.Trader(&_D2Vault.CallOpts)
}

// Trader is a free data retrieval call binding the contract method 0x1758078b.
//
// Solidity: function trader() view returns(address)
func (_D2Vault *D2VaultCallerSession) Trader() (common.Address, error) {
	return _D2Vault.Contract.Trader(&_D2Vault.CallOpts)
}

// WhitelistAsset is a free data retrieval call binding the contract method 0x0d89ddec.
//
// Solidity: function whitelistAsset() view returns(address)
func (_D2Vault *D2VaultCaller) WhitelistAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "whitelistAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WhitelistAsset is a free data retrieval call binding the contract method 0x0d89ddec.
//
// Solidity: function whitelistAsset() view returns(address)
func (_D2Vault *D2VaultSession) WhitelistAsset() (common.Address, error) {
	return _D2Vault.Contract.WhitelistAsset(&_D2Vault.CallOpts)
}

// WhitelistAsset is a free data retrieval call binding the contract method 0x0d89ddec.
//
// Solidity: function whitelistAsset() view returns(address)
func (_D2Vault *D2VaultCallerSession) WhitelistAsset() (common.Address, error) {
	return _D2Vault.Contract.WhitelistAsset(&_D2Vault.CallOpts)
}

// WhitelistBalance is a free data retrieval call binding the contract method 0x8c60f820.
//
// Solidity: function whitelistBalance() view returns(uint256)
func (_D2Vault *D2VaultCaller) WhitelistBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "whitelistBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WhitelistBalance is a free data retrieval call binding the contract method 0x8c60f820.
//
// Solidity: function whitelistBalance() view returns(uint256)
func (_D2Vault *D2VaultSession) WhitelistBalance() (*big.Int, error) {
	return _D2Vault.Contract.WhitelistBalance(&_D2Vault.CallOpts)
}

// WhitelistBalance is a free data retrieval call binding the contract method 0x8c60f820.
//
// Solidity: function whitelistBalance() view returns(uint256)
func (_D2Vault *D2VaultCallerSession) WhitelistBalance() (*big.Int, error) {
	return _D2Vault.Contract.WhitelistBalance(&_D2Vault.CallOpts)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_D2Vault *D2VaultCaller) Whitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _D2Vault.contract.Call(opts, &out, "whitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_D2Vault *D2VaultSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _D2Vault.Contract.Whitelisted(&_D2Vault.CallOpts, arg0)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_D2Vault *D2VaultCallerSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _D2Vault.Contract.Whitelisted(&_D2Vault.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_D2Vault *D2VaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_D2Vault *D2VaultSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _D2Vault.Contract.Approve(&_D2Vault.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_D2Vault *D2VaultTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _D2Vault.Contract.Approve(&_D2Vault.TransactOpts, spender, amount)
}

// CustodyFunds is a paid mutator transaction binding the contract method 0xfde7a941.
//
// Solidity: function custodyFunds() returns(uint256)
func (_D2Vault *D2VaultTransactor) CustodyFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "custodyFunds")
}

// CustodyFunds is a paid mutator transaction binding the contract method 0xfde7a941.
//
// Solidity: function custodyFunds() returns(uint256)
func (_D2Vault *D2VaultSession) CustodyFunds() (*types.Transaction, error) {
	return _D2Vault.Contract.CustodyFunds(&_D2Vault.TransactOpts)
}

// CustodyFunds is a paid mutator transaction binding the contract method 0xfde7a941.
//
// Solidity: function custodyFunds() returns(uint256)
func (_D2Vault *D2VaultTransactorSession) CustodyFunds() (*types.Transaction, error) {
	return _D2Vault.Contract.CustodyFunds(&_D2Vault.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_D2Vault *D2VaultTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_D2Vault *D2VaultSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _D2Vault.Contract.DecreaseAllowance(&_D2Vault.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_D2Vault *D2VaultTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _D2Vault.Contract.DecreaseAllowance(&_D2Vault.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_D2Vault *D2VaultTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_D2Vault *D2VaultSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _D2Vault.Contract.Deposit(&_D2Vault.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_D2Vault *D2VaultTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _D2Vault.Contract.Deposit(&_D2Vault.TransactOpts, assets, receiver)
}

// EmergencyEndEpoch is a paid mutator transaction binding the contract method 0x9bf05b18.
//
// Solidity: function emergencyEndEpoch(uint256 _amount) returns()
func (_D2Vault *D2VaultTransactor) EmergencyEndEpoch(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "emergencyEndEpoch", _amount)
}

// EmergencyEndEpoch is a paid mutator transaction binding the contract method 0x9bf05b18.
//
// Solidity: function emergencyEndEpoch(uint256 _amount) returns()
func (_D2Vault *D2VaultSession) EmergencyEndEpoch(_amount *big.Int) (*types.Transaction, error) {
	return _D2Vault.Contract.EmergencyEndEpoch(&_D2Vault.TransactOpts, _amount)
}

// EmergencyEndEpoch is a paid mutator transaction binding the contract method 0x9bf05b18.
//
// Solidity: function emergencyEndEpoch(uint256 _amount) returns()
func (_D2Vault *D2VaultTransactorSession) EmergencyEndEpoch(_amount *big.Int) (*types.Transaction, error) {
	return _D2Vault.Contract.EmergencyEndEpoch(&_D2Vault.TransactOpts, _amount)
}

// EmergencyFreeze is a paid mutator transaction binding the contract method 0xf3d4b942.
//
// Solidity: function emergencyFreeze() returns()
func (_D2Vault *D2VaultTransactor) EmergencyFreeze(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "emergencyFreeze")
}

// EmergencyFreeze is a paid mutator transaction binding the contract method 0xf3d4b942.
//
// Solidity: function emergencyFreeze() returns()
func (_D2Vault *D2VaultSession) EmergencyFreeze() (*types.Transaction, error) {
	return _D2Vault.Contract.EmergencyFreeze(&_D2Vault.TransactOpts)
}

// EmergencyFreeze is a paid mutator transaction binding the contract method 0xf3d4b942.
//
// Solidity: function emergencyFreeze() returns()
func (_D2Vault *D2VaultTransactorSession) EmergencyFreeze() (*types.Transaction, error) {
	return _D2Vault.Contract.EmergencyFreeze(&_D2Vault.TransactOpts)
}

// EmergencyStartEpoch is a paid mutator transaction binding the contract method 0x8ea0675a.
//
// Solidity: function emergencyStartEpoch(uint80 _fundingStart, uint80 _epochStart, uint80 _epochEnd) returns()
func (_D2Vault *D2VaultTransactor) EmergencyStartEpoch(opts *bind.TransactOpts, _fundingStart *big.Int, _epochStart *big.Int, _epochEnd *big.Int) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "emergencyStartEpoch", _fundingStart, _epochStart, _epochEnd)
}

// EmergencyStartEpoch is a paid mutator transaction binding the contract method 0x8ea0675a.
//
// Solidity: function emergencyStartEpoch(uint80 _fundingStart, uint80 _epochStart, uint80 _epochEnd) returns()
func (_D2Vault *D2VaultSession) EmergencyStartEpoch(_fundingStart *big.Int, _epochStart *big.Int, _epochEnd *big.Int) (*types.Transaction, error) {
	return _D2Vault.Contract.EmergencyStartEpoch(&_D2Vault.TransactOpts, _fundingStart, _epochStart, _epochEnd)
}

// EmergencyStartEpoch is a paid mutator transaction binding the contract method 0x8ea0675a.
//
// Solidity: function emergencyStartEpoch(uint80 _fundingStart, uint80 _epochStart, uint80 _epochEnd) returns()
func (_D2Vault *D2VaultTransactorSession) EmergencyStartEpoch(_fundingStart *big.Int, _epochStart *big.Int, _epochEnd *big.Int) (*types.Transaction, error) {
	return _D2Vault.Contract.EmergencyStartEpoch(&_D2Vault.TransactOpts, _fundingStart, _epochStart, _epochEnd)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_D2Vault *D2VaultTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_D2Vault *D2VaultSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _D2Vault.Contract.IncreaseAllowance(&_D2Vault.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_D2Vault *D2VaultTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _D2Vault.Contract.IncreaseAllowance(&_D2Vault.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x25c86a12.
//
// Solidity: function initialize((address,address,address,address,address,string,string,uint256,address,uint256,address,uint80,uint80,uint80) args) returns()
func (_D2Vault *D2VaultTransactor) Initialize(opts *bind.TransactOpts, args VaultV3ConstructorArgs) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "initialize", args)
}

// Initialize is a paid mutator transaction binding the contract method 0x25c86a12.
//
// Solidity: function initialize((address,address,address,address,address,string,string,uint256,address,uint256,address,uint80,uint80,uint80) args) returns()
func (_D2Vault *D2VaultSession) Initialize(args VaultV3ConstructorArgs) (*types.Transaction, error) {
	return _D2Vault.Contract.Initialize(&_D2Vault.TransactOpts, args)
}

// Initialize is a paid mutator transaction binding the contract method 0x25c86a12.
//
// Solidity: function initialize((address,address,address,address,address,string,string,uint256,address,uint256,address,uint80,uint80,uint80) args) returns()
func (_D2Vault *D2VaultTransactorSession) Initialize(args VaultV3ConstructorArgs) (*types.Transaction, error) {
	return _D2Vault.Contract.Initialize(&_D2Vault.TransactOpts, args)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_D2Vault *D2VaultTransactor) LzReceive(opts *bind.TransactOpts, _origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "lzReceive", _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_D2Vault *D2VaultSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _D2Vault.Contract.LzReceive(&_D2Vault.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_D2Vault *D2VaultTransactorSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _D2Vault.Contract.LzReceive(&_D2Vault.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceiveAndRevert is a paid mutator transaction binding the contract method 0xbd815db0.
//
// Solidity: function lzReceiveAndRevert(((uint32,bytes32,uint64),uint32,address,bytes32,uint256,address,bytes,bytes)[] _packets) payable returns()
func (_D2Vault *D2VaultTransactor) LzReceiveAndRevert(opts *bind.TransactOpts, _packets []InboundPacket) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "lzReceiveAndRevert", _packets)
}

// LzReceiveAndRevert is a paid mutator transaction binding the contract method 0xbd815db0.
//
// Solidity: function lzReceiveAndRevert(((uint32,bytes32,uint64),uint32,address,bytes32,uint256,address,bytes,bytes)[] _packets) payable returns()
func (_D2Vault *D2VaultSession) LzReceiveAndRevert(_packets []InboundPacket) (*types.Transaction, error) {
	return _D2Vault.Contract.LzReceiveAndRevert(&_D2Vault.TransactOpts, _packets)
}

// LzReceiveAndRevert is a paid mutator transaction binding the contract method 0xbd815db0.
//
// Solidity: function lzReceiveAndRevert(((uint32,bytes32,uint64),uint32,address,bytes32,uint256,address,bytes,bytes)[] _packets) payable returns()
func (_D2Vault *D2VaultTransactorSession) LzReceiveAndRevert(_packets []InboundPacket) (*types.Transaction, error) {
	return _D2Vault.Contract.LzReceiveAndRevert(&_D2Vault.TransactOpts, _packets)
}

// LzReceiveSimulate is a paid mutator transaction binding the contract method 0xd045a0dc.
//
// Solidity: function lzReceiveSimulate((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_D2Vault *D2VaultTransactor) LzReceiveSimulate(opts *bind.TransactOpts, _origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "lzReceiveSimulate", _origin, _guid, _message, _executor, _extraData)
}

// LzReceiveSimulate is a paid mutator transaction binding the contract method 0xd045a0dc.
//
// Solidity: function lzReceiveSimulate((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_D2Vault *D2VaultSession) LzReceiveSimulate(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _D2Vault.Contract.LzReceiveSimulate(&_D2Vault.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceiveSimulate is a paid mutator transaction binding the contract method 0xd045a0dc.
//
// Solidity: function lzReceiveSimulate((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_D2Vault *D2VaultTransactorSession) LzReceiveSimulate(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _D2Vault.Contract.LzReceiveSimulate(&_D2Vault.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_D2Vault *D2VaultTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_D2Vault *D2VaultSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _D2Vault.Contract.Mint(&_D2Vault.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_D2Vault *D2VaultTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _D2Vault.Contract.Mint(&_D2Vault.TransactOpts, shares, receiver)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address _owner) returns(uint256)
func (_D2Vault *D2VaultTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "redeem", shares, receiver, _owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address _owner) returns(uint256)
func (_D2Vault *D2VaultSession) Redeem(shares *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _D2Vault.Contract.Redeem(&_D2Vault.TransactOpts, shares, receiver, _owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address _owner) returns(uint256)
func (_D2Vault *D2VaultTransactorSession) Redeem(shares *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _D2Vault.Contract.Redeem(&_D2Vault.TransactOpts, shares, receiver, _owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_D2Vault *D2VaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_D2Vault *D2VaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _D2Vault.Contract.RenounceOwnership(&_D2Vault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_D2Vault *D2VaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _D2Vault.Contract.RenounceOwnership(&_D2Vault.TransactOpts)
}

// ReturnFunds is a paid mutator transaction binding the contract method 0x45755dd6.
//
// Solidity: function returnFunds(uint256 _amount) returns()
func (_D2Vault *D2VaultTransactor) ReturnFunds(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "returnFunds", _amount)
}

// ReturnFunds is a paid mutator transaction binding the contract method 0x45755dd6.
//
// Solidity: function returnFunds(uint256 _amount) returns()
func (_D2Vault *D2VaultSession) ReturnFunds(_amount *big.Int) (*types.Transaction, error) {
	return _D2Vault.Contract.ReturnFunds(&_D2Vault.TransactOpts, _amount)
}

// ReturnFunds is a paid mutator transaction binding the contract method 0x45755dd6.
//
// Solidity: function returnFunds(uint256 _amount) returns()
func (_D2Vault *D2VaultTransactorSession) ReturnFunds(_amount *big.Int) (*types.Transaction, error) {
	return _D2Vault.Contract.ReturnFunds(&_D2Vault.TransactOpts, _amount)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)) msgReceipt, (uint256,uint256) oftReceipt)
func (_D2Vault *D2VaultTransactor) Send(opts *bind.TransactOpts, _sendParam SendParam, _fee MessagingFee, _refundAddress common.Address) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "send", _sendParam, _fee, _refundAddress)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)) msgReceipt, (uint256,uint256) oftReceipt)
func (_D2Vault *D2VaultSession) Send(_sendParam SendParam, _fee MessagingFee, _refundAddress common.Address) (*types.Transaction, error) {
	return _D2Vault.Contract.Send(&_D2Vault.TransactOpts, _sendParam, _fee, _refundAddress)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)) msgReceipt, (uint256,uint256) oftReceipt)
func (_D2Vault *D2VaultTransactorSession) Send(_sendParam SendParam, _fee MessagingFee, _refundAddress common.Address) (*types.Transaction, error) {
	return _D2Vault.Contract.Send(&_D2Vault.TransactOpts, _sendParam, _fee, _refundAddress)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_D2Vault *D2VaultTransactor) SetDelegate(opts *bind.TransactOpts, _delegate common.Address) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "setDelegate", _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_D2Vault *D2VaultSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _D2Vault.Contract.SetDelegate(&_D2Vault.TransactOpts, _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_D2Vault *D2VaultTransactorSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _D2Vault.Contract.SetDelegate(&_D2Vault.TransactOpts, _delegate)
}

// SetEnforcedOptions is a paid mutator transaction binding the contract method 0xb98bd070.
//
// Solidity: function setEnforcedOptions((uint32,uint16,bytes)[] _enforcedOptions) returns()
func (_D2Vault *D2VaultTransactor) SetEnforcedOptions(opts *bind.TransactOpts, _enforcedOptions []EnforcedOptionParam) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "setEnforcedOptions", _enforcedOptions)
}

// SetEnforcedOptions is a paid mutator transaction binding the contract method 0xb98bd070.
//
// Solidity: function setEnforcedOptions((uint32,uint16,bytes)[] _enforcedOptions) returns()
func (_D2Vault *D2VaultSession) SetEnforcedOptions(_enforcedOptions []EnforcedOptionParam) (*types.Transaction, error) {
	return _D2Vault.Contract.SetEnforcedOptions(&_D2Vault.TransactOpts, _enforcedOptions)
}

// SetEnforcedOptions is a paid mutator transaction binding the contract method 0xb98bd070.
//
// Solidity: function setEnforcedOptions((uint32,uint16,bytes)[] _enforcedOptions) returns()
func (_D2Vault *D2VaultTransactorSession) SetEnforcedOptions(_enforcedOptions []EnforcedOptionParam) (*types.Transaction, error) {
	return _D2Vault.Contract.SetEnforcedOptions(&_D2Vault.TransactOpts, _enforcedOptions)
}

// SetMaxDeposits is a paid mutator transaction binding the contract method 0x00fed902.
//
// Solidity: function setMaxDeposits(uint256 _newMax) returns()
func (_D2Vault *D2VaultTransactor) SetMaxDeposits(opts *bind.TransactOpts, _newMax *big.Int) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "setMaxDeposits", _newMax)
}

// SetMaxDeposits is a paid mutator transaction binding the contract method 0x00fed902.
//
// Solidity: function setMaxDeposits(uint256 _newMax) returns()
func (_D2Vault *D2VaultSession) SetMaxDeposits(_newMax *big.Int) (*types.Transaction, error) {
	return _D2Vault.Contract.SetMaxDeposits(&_D2Vault.TransactOpts, _newMax)
}

// SetMaxDeposits is a paid mutator transaction binding the contract method 0x00fed902.
//
// Solidity: function setMaxDeposits(uint256 _newMax) returns()
func (_D2Vault *D2VaultTransactorSession) SetMaxDeposits(_newMax *big.Int) (*types.Transaction, error) {
	return _D2Vault.Contract.SetMaxDeposits(&_D2Vault.TransactOpts, _newMax)
}

// SetMsgInspector is a paid mutator transaction binding the contract method 0x6fc1b31e.
//
// Solidity: function setMsgInspector(address _msgInspector) returns()
func (_D2Vault *D2VaultTransactor) SetMsgInspector(opts *bind.TransactOpts, _msgInspector common.Address) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "setMsgInspector", _msgInspector)
}

// SetMsgInspector is a paid mutator transaction binding the contract method 0x6fc1b31e.
//
// Solidity: function setMsgInspector(address _msgInspector) returns()
func (_D2Vault *D2VaultSession) SetMsgInspector(_msgInspector common.Address) (*types.Transaction, error) {
	return _D2Vault.Contract.SetMsgInspector(&_D2Vault.TransactOpts, _msgInspector)
}

// SetMsgInspector is a paid mutator transaction binding the contract method 0x6fc1b31e.
//
// Solidity: function setMsgInspector(address _msgInspector) returns()
func (_D2Vault *D2VaultTransactorSession) SetMsgInspector(_msgInspector common.Address) (*types.Transaction, error) {
	return _D2Vault.Contract.SetMsgInspector(&_D2Vault.TransactOpts, _msgInspector)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_D2Vault *D2VaultTransactor) SetPeer(opts *bind.TransactOpts, _eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "setPeer", _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_D2Vault *D2VaultSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _D2Vault.Contract.SetPeer(&_D2Vault.TransactOpts, _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_D2Vault *D2VaultTransactorSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _D2Vault.Contract.SetPeer(&_D2Vault.TransactOpts, _eid, _peer)
}

// SetPreCrime is a paid mutator transaction binding the contract method 0xd4243885.
//
// Solidity: function setPreCrime(address _preCrime) returns()
func (_D2Vault *D2VaultTransactor) SetPreCrime(opts *bind.TransactOpts, _preCrime common.Address) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "setPreCrime", _preCrime)
}

// SetPreCrime is a paid mutator transaction binding the contract method 0xd4243885.
//
// Solidity: function setPreCrime(address _preCrime) returns()
func (_D2Vault *D2VaultSession) SetPreCrime(_preCrime common.Address) (*types.Transaction, error) {
	return _D2Vault.Contract.SetPreCrime(&_D2Vault.TransactOpts, _preCrime)
}

// SetPreCrime is a paid mutator transaction binding the contract method 0xd4243885.
//
// Solidity: function setPreCrime(address _preCrime) returns()
func (_D2Vault *D2VaultTransactorSession) SetPreCrime(_preCrime common.Address) (*types.Transaction, error) {
	return _D2Vault.Contract.SetPreCrime(&_D2Vault.TransactOpts, _preCrime)
}

// SetTrader is a paid mutator transaction binding the contract method 0xa6bc18f9.
//
// Solidity: function setTrader(address _trader) returns()
func (_D2Vault *D2VaultTransactor) SetTrader(opts *bind.TransactOpts, _trader common.Address) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "setTrader", _trader)
}

// SetTrader is a paid mutator transaction binding the contract method 0xa6bc18f9.
//
// Solidity: function setTrader(address _trader) returns()
func (_D2Vault *D2VaultSession) SetTrader(_trader common.Address) (*types.Transaction, error) {
	return _D2Vault.Contract.SetTrader(&_D2Vault.TransactOpts, _trader)
}

// SetTrader is a paid mutator transaction binding the contract method 0xa6bc18f9.
//
// Solidity: function setTrader(address _trader) returns()
func (_D2Vault *D2VaultTransactorSession) SetTrader(_trader common.Address) (*types.Transaction, error) {
	return _D2Vault.Contract.SetTrader(&_D2Vault.TransactOpts, _trader)
}

// SetWhitelistAsset is a paid mutator transaction binding the contract method 0x8ca67a8c.
//
// Solidity: function setWhitelistAsset(address _whitelistAsset) returns()
func (_D2Vault *D2VaultTransactor) SetWhitelistAsset(opts *bind.TransactOpts, _whitelistAsset common.Address) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "setWhitelistAsset", _whitelistAsset)
}

// SetWhitelistAsset is a paid mutator transaction binding the contract method 0x8ca67a8c.
//
// Solidity: function setWhitelistAsset(address _whitelistAsset) returns()
func (_D2Vault *D2VaultSession) SetWhitelistAsset(_whitelistAsset common.Address) (*types.Transaction, error) {
	return _D2Vault.Contract.SetWhitelistAsset(&_D2Vault.TransactOpts, _whitelistAsset)
}

// SetWhitelistAsset is a paid mutator transaction binding the contract method 0x8ca67a8c.
//
// Solidity: function setWhitelistAsset(address _whitelistAsset) returns()
func (_D2Vault *D2VaultTransactorSession) SetWhitelistAsset(_whitelistAsset common.Address) (*types.Transaction, error) {
	return _D2Vault.Contract.SetWhitelistAsset(&_D2Vault.TransactOpts, _whitelistAsset)
}

// SetWhitelistBalance is a paid mutator transaction binding the contract method 0x2d463560.
//
// Solidity: function setWhitelistBalance(uint256 _whitelistBalance) returns()
func (_D2Vault *D2VaultTransactor) SetWhitelistBalance(opts *bind.TransactOpts, _whitelistBalance *big.Int) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "setWhitelistBalance", _whitelistBalance)
}

// SetWhitelistBalance is a paid mutator transaction binding the contract method 0x2d463560.
//
// Solidity: function setWhitelistBalance(uint256 _whitelistBalance) returns()
func (_D2Vault *D2VaultSession) SetWhitelistBalance(_whitelistBalance *big.Int) (*types.Transaction, error) {
	return _D2Vault.Contract.SetWhitelistBalance(&_D2Vault.TransactOpts, _whitelistBalance)
}

// SetWhitelistBalance is a paid mutator transaction binding the contract method 0x2d463560.
//
// Solidity: function setWhitelistBalance(uint256 _whitelistBalance) returns()
func (_D2Vault *D2VaultTransactorSession) SetWhitelistBalance(_whitelistBalance *big.Int) (*types.Transaction, error) {
	return _D2Vault.Contract.SetWhitelistBalance(&_D2Vault.TransactOpts, _whitelistBalance)
}

// SetWhitelistStatus is a paid mutator transaction binding the contract method 0x0c424284.
//
// Solidity: function setWhitelistStatus(address _user, bool _status) returns()
func (_D2Vault *D2VaultTransactor) SetWhitelistStatus(opts *bind.TransactOpts, _user common.Address, _status bool) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "setWhitelistStatus", _user, _status)
}

// SetWhitelistStatus is a paid mutator transaction binding the contract method 0x0c424284.
//
// Solidity: function setWhitelistStatus(address _user, bool _status) returns()
func (_D2Vault *D2VaultSession) SetWhitelistStatus(_user common.Address, _status bool) (*types.Transaction, error) {
	return _D2Vault.Contract.SetWhitelistStatus(&_D2Vault.TransactOpts, _user, _status)
}

// SetWhitelistStatus is a paid mutator transaction binding the contract method 0x0c424284.
//
// Solidity: function setWhitelistStatus(address _user, bool _status) returns()
func (_D2Vault *D2VaultTransactorSession) SetWhitelistStatus(_user common.Address, _status bool) (*types.Transaction, error) {
	return _D2Vault.Contract.SetWhitelistStatus(&_D2Vault.TransactOpts, _user, _status)
}

// SetWhitelistStatuses is a paid mutator transaction binding the contract method 0x7e046634.
//
// Solidity: function setWhitelistStatuses(address[] _users, bool[] _statuses) returns()
func (_D2Vault *D2VaultTransactor) SetWhitelistStatuses(opts *bind.TransactOpts, _users []common.Address, _statuses []bool) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "setWhitelistStatuses", _users, _statuses)
}

// SetWhitelistStatuses is a paid mutator transaction binding the contract method 0x7e046634.
//
// Solidity: function setWhitelistStatuses(address[] _users, bool[] _statuses) returns()
func (_D2Vault *D2VaultSession) SetWhitelistStatuses(_users []common.Address, _statuses []bool) (*types.Transaction, error) {
	return _D2Vault.Contract.SetWhitelistStatuses(&_D2Vault.TransactOpts, _users, _statuses)
}

// SetWhitelistStatuses is a paid mutator transaction binding the contract method 0x7e046634.
//
// Solidity: function setWhitelistStatuses(address[] _users, bool[] _statuses) returns()
func (_D2Vault *D2VaultTransactorSession) SetWhitelistStatuses(_users []common.Address, _statuses []bool) (*types.Transaction, error) {
	return _D2Vault.Contract.SetWhitelistStatuses(&_D2Vault.TransactOpts, _users, _statuses)
}

// StartEpoch is a paid mutator transaction binding the contract method 0x06e67d30.
//
// Solidity: function startEpoch(uint80 _fundingStart, uint80 _epochStart, uint80 _epochEnd) returns()
func (_D2Vault *D2VaultTransactor) StartEpoch(opts *bind.TransactOpts, _fundingStart *big.Int, _epochStart *big.Int, _epochEnd *big.Int) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "startEpoch", _fundingStart, _epochStart, _epochEnd)
}

// StartEpoch is a paid mutator transaction binding the contract method 0x06e67d30.
//
// Solidity: function startEpoch(uint80 _fundingStart, uint80 _epochStart, uint80 _epochEnd) returns()
func (_D2Vault *D2VaultSession) StartEpoch(_fundingStart *big.Int, _epochStart *big.Int, _epochEnd *big.Int) (*types.Transaction, error) {
	return _D2Vault.Contract.StartEpoch(&_D2Vault.TransactOpts, _fundingStart, _epochStart, _epochEnd)
}

// StartEpoch is a paid mutator transaction binding the contract method 0x06e67d30.
//
// Solidity: function startEpoch(uint80 _fundingStart, uint80 _epochStart, uint80 _epochEnd) returns()
func (_D2Vault *D2VaultTransactorSession) StartEpoch(_fundingStart *big.Int, _epochStart *big.Int, _epochEnd *big.Int) (*types.Transaction, error) {
	return _D2Vault.Contract.StartEpoch(&_D2Vault.TransactOpts, _fundingStart, _epochStart, _epochEnd)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_D2Vault *D2VaultTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_D2Vault *D2VaultSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _D2Vault.Contract.Transfer(&_D2Vault.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_D2Vault *D2VaultTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _D2Vault.Contract.Transfer(&_D2Vault.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_D2Vault *D2VaultTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_D2Vault *D2VaultSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _D2Vault.Contract.TransferFrom(&_D2Vault.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_D2Vault *D2VaultTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _D2Vault.Contract.TransferFrom(&_D2Vault.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_D2Vault *D2VaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_D2Vault *D2VaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _D2Vault.Contract.TransferOwnership(&_D2Vault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_D2Vault *D2VaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _D2Vault.Contract.TransferOwnership(&_D2Vault.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address _owner) returns(uint256)
func (_D2Vault *D2VaultTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _D2Vault.contract.Transact(opts, "withdraw", assets, receiver, _owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address _owner) returns(uint256)
func (_D2Vault *D2VaultSession) Withdraw(assets *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _D2Vault.Contract.Withdraw(&_D2Vault.TransactOpts, assets, receiver, _owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address _owner) returns(uint256)
func (_D2Vault *D2VaultTransactorSession) Withdraw(assets *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _D2Vault.Contract.Withdraw(&_D2Vault.TransactOpts, assets, receiver, _owner)
}

// D2VaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the D2Vault contract.
type D2VaultApprovalIterator struct {
	Event *D2VaultApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *D2VaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D2VaultApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(D2VaultApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *D2VaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D2VaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D2VaultApproval represents a Approval event raised by the D2Vault contract.
type D2VaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_D2Vault *D2VaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*D2VaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _D2Vault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &D2VaultApprovalIterator{contract: _D2Vault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_D2Vault *D2VaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *D2VaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _D2Vault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D2VaultApproval)
				if err := _D2Vault.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_D2Vault *D2VaultFilterer) ParseApproval(log types.Log) (*D2VaultApproval, error) {
	event := new(D2VaultApproval)
	if err := _D2Vault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// D2VaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the D2Vault contract.
type D2VaultDepositIterator struct {
	Event *D2VaultDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *D2VaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D2VaultDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(D2VaultDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *D2VaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D2VaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D2VaultDeposit represents a Deposit event raised by the D2Vault contract.
type D2VaultDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_D2Vault *D2VaultFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*D2VaultDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _D2Vault.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &D2VaultDepositIterator{contract: _D2Vault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_D2Vault *D2VaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *D2VaultDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _D2Vault.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D2VaultDeposit)
				if err := _D2Vault.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_D2Vault *D2VaultFilterer) ParseDeposit(log types.Log) (*D2VaultDeposit, error) {
	event := new(D2VaultDeposit)
	if err := _D2Vault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// D2VaultEnforcedOptionSetIterator is returned from FilterEnforcedOptionSet and is used to iterate over the raw logs and unpacked data for EnforcedOptionSet events raised by the D2Vault contract.
type D2VaultEnforcedOptionSetIterator struct {
	Event *D2VaultEnforcedOptionSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *D2VaultEnforcedOptionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D2VaultEnforcedOptionSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(D2VaultEnforcedOptionSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *D2VaultEnforcedOptionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D2VaultEnforcedOptionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D2VaultEnforcedOptionSet represents a EnforcedOptionSet event raised by the D2Vault contract.
type D2VaultEnforcedOptionSet struct {
	EnforcedOptions []EnforcedOptionParam
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEnforcedOptionSet is a free log retrieval operation binding the contract event 0xbe4864a8e820971c0247f5992e2da559595f7bf076a21cb5928d443d2a13b674.
//
// Solidity: event EnforcedOptionSet((uint32,uint16,bytes)[] _enforcedOptions)
func (_D2Vault *D2VaultFilterer) FilterEnforcedOptionSet(opts *bind.FilterOpts) (*D2VaultEnforcedOptionSetIterator, error) {

	logs, sub, err := _D2Vault.contract.FilterLogs(opts, "EnforcedOptionSet")
	if err != nil {
		return nil, err
	}
	return &D2VaultEnforcedOptionSetIterator{contract: _D2Vault.contract, event: "EnforcedOptionSet", logs: logs, sub: sub}, nil
}

// WatchEnforcedOptionSet is a free log subscription operation binding the contract event 0xbe4864a8e820971c0247f5992e2da559595f7bf076a21cb5928d443d2a13b674.
//
// Solidity: event EnforcedOptionSet((uint32,uint16,bytes)[] _enforcedOptions)
func (_D2Vault *D2VaultFilterer) WatchEnforcedOptionSet(opts *bind.WatchOpts, sink chan<- *D2VaultEnforcedOptionSet) (event.Subscription, error) {

	logs, sub, err := _D2Vault.contract.WatchLogs(opts, "EnforcedOptionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D2VaultEnforcedOptionSet)
				if err := _D2Vault.contract.UnpackLog(event, "EnforcedOptionSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEnforcedOptionSet is a log parse operation binding the contract event 0xbe4864a8e820971c0247f5992e2da559595f7bf076a21cb5928d443d2a13b674.
//
// Solidity: event EnforcedOptionSet((uint32,uint16,bytes)[] _enforcedOptions)
func (_D2Vault *D2VaultFilterer) ParseEnforcedOptionSet(log types.Log) (*D2VaultEnforcedOptionSet, error) {
	event := new(D2VaultEnforcedOptionSet)
	if err := _D2Vault.contract.UnpackLog(event, "EnforcedOptionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// D2VaultEpochStartedIterator is returned from FilterEpochStarted and is used to iterate over the raw logs and unpacked data for EpochStarted events raised by the D2Vault contract.
type D2VaultEpochStartedIterator struct {
	Event *D2VaultEpochStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *D2VaultEpochStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D2VaultEpochStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(D2VaultEpochStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *D2VaultEpochStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D2VaultEpochStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D2VaultEpochStarted represents a EpochStarted event raised by the D2Vault contract.
type D2VaultEpochStarted struct {
	Epoch        *big.Int
	FundingStart *big.Int
	EpochStart   *big.Int
	EpochEnd     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEpochStarted is a free log retrieval operation binding the contract event 0xfc6c725a951b30501cae57e21f8ee2c863bc9eb91a4fc9bc85ef4f4e07b861e1.
//
// Solidity: event EpochStarted(uint256 indexed epoch, uint256 fundingStart, uint256 epochStart, uint256 epochEnd)
func (_D2Vault *D2VaultFilterer) FilterEpochStarted(opts *bind.FilterOpts, epoch []*big.Int) (*D2VaultEpochStartedIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _D2Vault.contract.FilterLogs(opts, "EpochStarted", epochRule)
	if err != nil {
		return nil, err
	}
	return &D2VaultEpochStartedIterator{contract: _D2Vault.contract, event: "EpochStarted", logs: logs, sub: sub}, nil
}

// WatchEpochStarted is a free log subscription operation binding the contract event 0xfc6c725a951b30501cae57e21f8ee2c863bc9eb91a4fc9bc85ef4f4e07b861e1.
//
// Solidity: event EpochStarted(uint256 indexed epoch, uint256 fundingStart, uint256 epochStart, uint256 epochEnd)
func (_D2Vault *D2VaultFilterer) WatchEpochStarted(opts *bind.WatchOpts, sink chan<- *D2VaultEpochStarted, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _D2Vault.contract.WatchLogs(opts, "EpochStarted", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D2VaultEpochStarted)
				if err := _D2Vault.contract.UnpackLog(event, "EpochStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEpochStarted is a log parse operation binding the contract event 0xfc6c725a951b30501cae57e21f8ee2c863bc9eb91a4fc9bc85ef4f4e07b861e1.
//
// Solidity: event EpochStarted(uint256 indexed epoch, uint256 fundingStart, uint256 epochStart, uint256 epochEnd)
func (_D2Vault *D2VaultFilterer) ParseEpochStarted(log types.Log) (*D2VaultEpochStarted, error) {
	event := new(D2VaultEpochStarted)
	if err := _D2Vault.contract.UnpackLog(event, "EpochStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// D2VaultFundsCustodiedIterator is returned from FilterFundsCustodied and is used to iterate over the raw logs and unpacked data for FundsCustodied events raised by the D2Vault contract.
type D2VaultFundsCustodiedIterator struct {
	Event *D2VaultFundsCustodied // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *D2VaultFundsCustodiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D2VaultFundsCustodied)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(D2VaultFundsCustodied)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *D2VaultFundsCustodiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D2VaultFundsCustodiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D2VaultFundsCustodied represents a FundsCustodied event raised by the D2Vault contract.
type D2VaultFundsCustodied struct {
	Epoch  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsCustodied is a free log retrieval operation binding the contract event 0x69e193dd4c77613d0e599740c9e2cd88fb7b4a9d11ef9b1f6226d392c941f471.
//
// Solidity: event FundsCustodied(uint256 indexed epoch, uint256 amount)
func (_D2Vault *D2VaultFilterer) FilterFundsCustodied(opts *bind.FilterOpts, epoch []*big.Int) (*D2VaultFundsCustodiedIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _D2Vault.contract.FilterLogs(opts, "FundsCustodied", epochRule)
	if err != nil {
		return nil, err
	}
	return &D2VaultFundsCustodiedIterator{contract: _D2Vault.contract, event: "FundsCustodied", logs: logs, sub: sub}, nil
}

// WatchFundsCustodied is a free log subscription operation binding the contract event 0x69e193dd4c77613d0e599740c9e2cd88fb7b4a9d11ef9b1f6226d392c941f471.
//
// Solidity: event FundsCustodied(uint256 indexed epoch, uint256 amount)
func (_D2Vault *D2VaultFilterer) WatchFundsCustodied(opts *bind.WatchOpts, sink chan<- *D2VaultFundsCustodied, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _D2Vault.contract.WatchLogs(opts, "FundsCustodied", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D2VaultFundsCustodied)
				if err := _D2Vault.contract.UnpackLog(event, "FundsCustodied", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFundsCustodied is a log parse operation binding the contract event 0x69e193dd4c77613d0e599740c9e2cd88fb7b4a9d11ef9b1f6226d392c941f471.
//
// Solidity: event FundsCustodied(uint256 indexed epoch, uint256 amount)
func (_D2Vault *D2VaultFilterer) ParseFundsCustodied(log types.Log) (*D2VaultFundsCustodied, error) {
	event := new(D2VaultFundsCustodied)
	if err := _D2Vault.contract.UnpackLog(event, "FundsCustodied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// D2VaultFundsReturnedIterator is returned from FilterFundsReturned and is used to iterate over the raw logs and unpacked data for FundsReturned events raised by the D2Vault contract.
type D2VaultFundsReturnedIterator struct {
	Event *D2VaultFundsReturned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *D2VaultFundsReturnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D2VaultFundsReturned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(D2VaultFundsReturned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *D2VaultFundsReturnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D2VaultFundsReturnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D2VaultFundsReturned represents a FundsReturned event raised by the D2Vault contract.
type D2VaultFundsReturned struct {
	Epoch  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsReturned is a free log retrieval operation binding the contract event 0xe5e9cfeede9ff1fc77b415bf8346e29706d16794b3bdeca347ac54a7fd3e0c3c.
//
// Solidity: event FundsReturned(uint256 indexed epoch, uint256 amount)
func (_D2Vault *D2VaultFilterer) FilterFundsReturned(opts *bind.FilterOpts, epoch []*big.Int) (*D2VaultFundsReturnedIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _D2Vault.contract.FilterLogs(opts, "FundsReturned", epochRule)
	if err != nil {
		return nil, err
	}
	return &D2VaultFundsReturnedIterator{contract: _D2Vault.contract, event: "FundsReturned", logs: logs, sub: sub}, nil
}

// WatchFundsReturned is a free log subscription operation binding the contract event 0xe5e9cfeede9ff1fc77b415bf8346e29706d16794b3bdeca347ac54a7fd3e0c3c.
//
// Solidity: event FundsReturned(uint256 indexed epoch, uint256 amount)
func (_D2Vault *D2VaultFilterer) WatchFundsReturned(opts *bind.WatchOpts, sink chan<- *D2VaultFundsReturned, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _D2Vault.contract.WatchLogs(opts, "FundsReturned", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D2VaultFundsReturned)
				if err := _D2Vault.contract.UnpackLog(event, "FundsReturned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFundsReturned is a log parse operation binding the contract event 0xe5e9cfeede9ff1fc77b415bf8346e29706d16794b3bdeca347ac54a7fd3e0c3c.
//
// Solidity: event FundsReturned(uint256 indexed epoch, uint256 amount)
func (_D2Vault *D2VaultFilterer) ParseFundsReturned(log types.Log) (*D2VaultFundsReturned, error) {
	event := new(D2VaultFundsReturned)
	if err := _D2Vault.contract.UnpackLog(event, "FundsReturned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// D2VaultMsgInspectorSetIterator is returned from FilterMsgInspectorSet and is used to iterate over the raw logs and unpacked data for MsgInspectorSet events raised by the D2Vault contract.
type D2VaultMsgInspectorSetIterator struct {
	Event *D2VaultMsgInspectorSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *D2VaultMsgInspectorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D2VaultMsgInspectorSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(D2VaultMsgInspectorSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *D2VaultMsgInspectorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D2VaultMsgInspectorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D2VaultMsgInspectorSet represents a MsgInspectorSet event raised by the D2Vault contract.
type D2VaultMsgInspectorSet struct {
	Inspector common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMsgInspectorSet is a free log retrieval operation binding the contract event 0xf0be4f1e87349231d80c36b33f9e8639658eeaf474014dee15a3e6a4d4414197.
//
// Solidity: event MsgInspectorSet(address inspector)
func (_D2Vault *D2VaultFilterer) FilterMsgInspectorSet(opts *bind.FilterOpts) (*D2VaultMsgInspectorSetIterator, error) {

	logs, sub, err := _D2Vault.contract.FilterLogs(opts, "MsgInspectorSet")
	if err != nil {
		return nil, err
	}
	return &D2VaultMsgInspectorSetIterator{contract: _D2Vault.contract, event: "MsgInspectorSet", logs: logs, sub: sub}, nil
}

// WatchMsgInspectorSet is a free log subscription operation binding the contract event 0xf0be4f1e87349231d80c36b33f9e8639658eeaf474014dee15a3e6a4d4414197.
//
// Solidity: event MsgInspectorSet(address inspector)
func (_D2Vault *D2VaultFilterer) WatchMsgInspectorSet(opts *bind.WatchOpts, sink chan<- *D2VaultMsgInspectorSet) (event.Subscription, error) {

	logs, sub, err := _D2Vault.contract.WatchLogs(opts, "MsgInspectorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D2VaultMsgInspectorSet)
				if err := _D2Vault.contract.UnpackLog(event, "MsgInspectorSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMsgInspectorSet is a log parse operation binding the contract event 0xf0be4f1e87349231d80c36b33f9e8639658eeaf474014dee15a3e6a4d4414197.
//
// Solidity: event MsgInspectorSet(address inspector)
func (_D2Vault *D2VaultFilterer) ParseMsgInspectorSet(log types.Log) (*D2VaultMsgInspectorSet, error) {
	event := new(D2VaultMsgInspectorSet)
	if err := _D2Vault.contract.UnpackLog(event, "MsgInspectorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// D2VaultNewMaxDepositsIterator is returned from FilterNewMaxDeposits and is used to iterate over the raw logs and unpacked data for NewMaxDeposits events raised by the D2Vault contract.
type D2VaultNewMaxDepositsIterator struct {
	Event *D2VaultNewMaxDeposits // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *D2VaultNewMaxDepositsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D2VaultNewMaxDeposits)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(D2VaultNewMaxDeposits)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *D2VaultNewMaxDepositsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D2VaultNewMaxDepositsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D2VaultNewMaxDeposits represents a NewMaxDeposits event raised by the D2Vault contract.
type D2VaultNewMaxDeposits struct {
	OldMax *big.Int
	NewMax *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewMaxDeposits is a free log retrieval operation binding the contract event 0xd026c0acfe6cb0f0c80c55b5d91a908394572c2a440fe841386448dc9c16ea77.
//
// Solidity: event NewMaxDeposits(uint256 oldMax, uint256 newMax)
func (_D2Vault *D2VaultFilterer) FilterNewMaxDeposits(opts *bind.FilterOpts) (*D2VaultNewMaxDepositsIterator, error) {

	logs, sub, err := _D2Vault.contract.FilterLogs(opts, "NewMaxDeposits")
	if err != nil {
		return nil, err
	}
	return &D2VaultNewMaxDepositsIterator{contract: _D2Vault.contract, event: "NewMaxDeposits", logs: logs, sub: sub}, nil
}

// WatchNewMaxDeposits is a free log subscription operation binding the contract event 0xd026c0acfe6cb0f0c80c55b5d91a908394572c2a440fe841386448dc9c16ea77.
//
// Solidity: event NewMaxDeposits(uint256 oldMax, uint256 newMax)
func (_D2Vault *D2VaultFilterer) WatchNewMaxDeposits(opts *bind.WatchOpts, sink chan<- *D2VaultNewMaxDeposits) (event.Subscription, error) {

	logs, sub, err := _D2Vault.contract.WatchLogs(opts, "NewMaxDeposits")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D2VaultNewMaxDeposits)
				if err := _D2Vault.contract.UnpackLog(event, "NewMaxDeposits", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewMaxDeposits is a log parse operation binding the contract event 0xd026c0acfe6cb0f0c80c55b5d91a908394572c2a440fe841386448dc9c16ea77.
//
// Solidity: event NewMaxDeposits(uint256 oldMax, uint256 newMax)
func (_D2Vault *D2VaultFilterer) ParseNewMaxDeposits(log types.Log) (*D2VaultNewMaxDeposits, error) {
	event := new(D2VaultNewMaxDeposits)
	if err := _D2Vault.contract.UnpackLog(event, "NewMaxDeposits", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// D2VaultNewWhitelistStatusIterator is returned from FilterNewWhitelistStatus and is used to iterate over the raw logs and unpacked data for NewWhitelistStatus events raised by the D2Vault contract.
type D2VaultNewWhitelistStatusIterator struct {
	Event *D2VaultNewWhitelistStatus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *D2VaultNewWhitelistStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D2VaultNewWhitelistStatus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(D2VaultNewWhitelistStatus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *D2VaultNewWhitelistStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D2VaultNewWhitelistStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D2VaultNewWhitelistStatus represents a NewWhitelistStatus event raised by the D2Vault contract.
type D2VaultNewWhitelistStatus struct {
	User   common.Address
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewWhitelistStatus is a free log retrieval operation binding the contract event 0xf8ccc4383c96e7994507db23f575d09dee79c80d522b953dbb53c3c9a3f970d3.
//
// Solidity: event NewWhitelistStatus(address indexed user, bool status)
func (_D2Vault *D2VaultFilterer) FilterNewWhitelistStatus(opts *bind.FilterOpts, user []common.Address) (*D2VaultNewWhitelistStatusIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _D2Vault.contract.FilterLogs(opts, "NewWhitelistStatus", userRule)
	if err != nil {
		return nil, err
	}
	return &D2VaultNewWhitelistStatusIterator{contract: _D2Vault.contract, event: "NewWhitelistStatus", logs: logs, sub: sub}, nil
}

// WatchNewWhitelistStatus is a free log subscription operation binding the contract event 0xf8ccc4383c96e7994507db23f575d09dee79c80d522b953dbb53c3c9a3f970d3.
//
// Solidity: event NewWhitelistStatus(address indexed user, bool status)
func (_D2Vault *D2VaultFilterer) WatchNewWhitelistStatus(opts *bind.WatchOpts, sink chan<- *D2VaultNewWhitelistStatus, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _D2Vault.contract.WatchLogs(opts, "NewWhitelistStatus", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D2VaultNewWhitelistStatus)
				if err := _D2Vault.contract.UnpackLog(event, "NewWhitelistStatus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewWhitelistStatus is a log parse operation binding the contract event 0xf8ccc4383c96e7994507db23f575d09dee79c80d522b953dbb53c3c9a3f970d3.
//
// Solidity: event NewWhitelistStatus(address indexed user, bool status)
func (_D2Vault *D2VaultFilterer) ParseNewWhitelistStatus(log types.Log) (*D2VaultNewWhitelistStatus, error) {
	event := new(D2VaultNewWhitelistStatus)
	if err := _D2Vault.contract.UnpackLog(event, "NewWhitelistStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// D2VaultOFTReceivedIterator is returned from FilterOFTReceived and is used to iterate over the raw logs and unpacked data for OFTReceived events raised by the D2Vault contract.
type D2VaultOFTReceivedIterator struct {
	Event *D2VaultOFTReceived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *D2VaultOFTReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D2VaultOFTReceived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(D2VaultOFTReceived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *D2VaultOFTReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D2VaultOFTReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D2VaultOFTReceived represents a OFTReceived event raised by the D2Vault contract.
type D2VaultOFTReceived struct {
	Guid             [32]byte
	SrcEid           uint32
	ToAddress        common.Address
	AmountReceivedLD *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOFTReceived is a free log retrieval operation binding the contract event 0xefed6d3500546b29533b128a29e3a94d70788727f0507505ac12eaf2e578fd9c.
//
// Solidity: event OFTReceived(bytes32 indexed guid, uint32 srcEid, address indexed toAddress, uint256 amountReceivedLD)
func (_D2Vault *D2VaultFilterer) FilterOFTReceived(opts *bind.FilterOpts, guid [][32]byte, toAddress []common.Address) (*D2VaultOFTReceivedIterator, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _D2Vault.contract.FilterLogs(opts, "OFTReceived", guidRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return &D2VaultOFTReceivedIterator{contract: _D2Vault.contract, event: "OFTReceived", logs: logs, sub: sub}, nil
}

// WatchOFTReceived is a free log subscription operation binding the contract event 0xefed6d3500546b29533b128a29e3a94d70788727f0507505ac12eaf2e578fd9c.
//
// Solidity: event OFTReceived(bytes32 indexed guid, uint32 srcEid, address indexed toAddress, uint256 amountReceivedLD)
func (_D2Vault *D2VaultFilterer) WatchOFTReceived(opts *bind.WatchOpts, sink chan<- *D2VaultOFTReceived, guid [][32]byte, toAddress []common.Address) (event.Subscription, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _D2Vault.contract.WatchLogs(opts, "OFTReceived", guidRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D2VaultOFTReceived)
				if err := _D2Vault.contract.UnpackLog(event, "OFTReceived", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOFTReceived is a log parse operation binding the contract event 0xefed6d3500546b29533b128a29e3a94d70788727f0507505ac12eaf2e578fd9c.
//
// Solidity: event OFTReceived(bytes32 indexed guid, uint32 srcEid, address indexed toAddress, uint256 amountReceivedLD)
func (_D2Vault *D2VaultFilterer) ParseOFTReceived(log types.Log) (*D2VaultOFTReceived, error) {
	event := new(D2VaultOFTReceived)
	if err := _D2Vault.contract.UnpackLog(event, "OFTReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// D2VaultOFTSentIterator is returned from FilterOFTSent and is used to iterate over the raw logs and unpacked data for OFTSent events raised by the D2Vault contract.
type D2VaultOFTSentIterator struct {
	Event *D2VaultOFTSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *D2VaultOFTSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D2VaultOFTSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(D2VaultOFTSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *D2VaultOFTSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D2VaultOFTSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D2VaultOFTSent represents a OFTSent event raised by the D2Vault contract.
type D2VaultOFTSent struct {
	Guid             [32]byte
	DstEid           uint32
	FromAddress      common.Address
	AmountSentLD     *big.Int
	AmountReceivedLD *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOFTSent is a free log retrieval operation binding the contract event 0x85496b760a4b7f8d66384b9df21b381f5d1b1e79f229a47aaf4c232edc2fe59a.
//
// Solidity: event OFTSent(bytes32 indexed guid, uint32 dstEid, address indexed fromAddress, uint256 amountSentLD, uint256 amountReceivedLD)
func (_D2Vault *D2VaultFilterer) FilterOFTSent(opts *bind.FilterOpts, guid [][32]byte, fromAddress []common.Address) (*D2VaultOFTSentIterator, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _D2Vault.contract.FilterLogs(opts, "OFTSent", guidRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return &D2VaultOFTSentIterator{contract: _D2Vault.contract, event: "OFTSent", logs: logs, sub: sub}, nil
}

// WatchOFTSent is a free log subscription operation binding the contract event 0x85496b760a4b7f8d66384b9df21b381f5d1b1e79f229a47aaf4c232edc2fe59a.
//
// Solidity: event OFTSent(bytes32 indexed guid, uint32 dstEid, address indexed fromAddress, uint256 amountSentLD, uint256 amountReceivedLD)
func (_D2Vault *D2VaultFilterer) WatchOFTSent(opts *bind.WatchOpts, sink chan<- *D2VaultOFTSent, guid [][32]byte, fromAddress []common.Address) (event.Subscription, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _D2Vault.contract.WatchLogs(opts, "OFTSent", guidRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D2VaultOFTSent)
				if err := _D2Vault.contract.UnpackLog(event, "OFTSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOFTSent is a log parse operation binding the contract event 0x85496b760a4b7f8d66384b9df21b381f5d1b1e79f229a47aaf4c232edc2fe59a.
//
// Solidity: event OFTSent(bytes32 indexed guid, uint32 dstEid, address indexed fromAddress, uint256 amountSentLD, uint256 amountReceivedLD)
func (_D2Vault *D2VaultFilterer) ParseOFTSent(log types.Log) (*D2VaultOFTSent, error) {
	event := new(D2VaultOFTSent)
	if err := _D2Vault.contract.UnpackLog(event, "OFTSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// D2VaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the D2Vault contract.
type D2VaultOwnershipTransferredIterator struct {
	Event *D2VaultOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *D2VaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D2VaultOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(D2VaultOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *D2VaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D2VaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D2VaultOwnershipTransferred represents a OwnershipTransferred event raised by the D2Vault contract.
type D2VaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_D2Vault *D2VaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*D2VaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _D2Vault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &D2VaultOwnershipTransferredIterator{contract: _D2Vault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_D2Vault *D2VaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *D2VaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _D2Vault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D2VaultOwnershipTransferred)
				if err := _D2Vault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_D2Vault *D2VaultFilterer) ParseOwnershipTransferred(log types.Log) (*D2VaultOwnershipTransferred, error) {
	event := new(D2VaultOwnershipTransferred)
	if err := _D2Vault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// D2VaultPeerSetIterator is returned from FilterPeerSet and is used to iterate over the raw logs and unpacked data for PeerSet events raised by the D2Vault contract.
type D2VaultPeerSetIterator struct {
	Event *D2VaultPeerSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *D2VaultPeerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D2VaultPeerSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(D2VaultPeerSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *D2VaultPeerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D2VaultPeerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D2VaultPeerSet represents a PeerSet event raised by the D2Vault contract.
type D2VaultPeerSet struct {
	Eid  uint32
	Peer [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPeerSet is a free log retrieval operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_D2Vault *D2VaultFilterer) FilterPeerSet(opts *bind.FilterOpts) (*D2VaultPeerSetIterator, error) {

	logs, sub, err := _D2Vault.contract.FilterLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return &D2VaultPeerSetIterator{contract: _D2Vault.contract, event: "PeerSet", logs: logs, sub: sub}, nil
}

// WatchPeerSet is a free log subscription operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_D2Vault *D2VaultFilterer) WatchPeerSet(opts *bind.WatchOpts, sink chan<- *D2VaultPeerSet) (event.Subscription, error) {

	logs, sub, err := _D2Vault.contract.WatchLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D2VaultPeerSet)
				if err := _D2Vault.contract.UnpackLog(event, "PeerSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePeerSet is a log parse operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_D2Vault *D2VaultFilterer) ParsePeerSet(log types.Log) (*D2VaultPeerSet, error) {
	event := new(D2VaultPeerSet)
	if err := _D2Vault.contract.UnpackLog(event, "PeerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// D2VaultPreCrimeSetIterator is returned from FilterPreCrimeSet and is used to iterate over the raw logs and unpacked data for PreCrimeSet events raised by the D2Vault contract.
type D2VaultPreCrimeSetIterator struct {
	Event *D2VaultPreCrimeSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *D2VaultPreCrimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D2VaultPreCrimeSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(D2VaultPreCrimeSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *D2VaultPreCrimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D2VaultPreCrimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D2VaultPreCrimeSet represents a PreCrimeSet event raised by the D2Vault contract.
type D2VaultPreCrimeSet struct {
	PreCrimeAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPreCrimeSet is a free log retrieval operation binding the contract event 0xd48d879cef83a1c0bdda516f27b13ddb1b3f8bbac1c9e1511bb2a659c2427760.
//
// Solidity: event PreCrimeSet(address preCrimeAddress)
func (_D2Vault *D2VaultFilterer) FilterPreCrimeSet(opts *bind.FilterOpts) (*D2VaultPreCrimeSetIterator, error) {

	logs, sub, err := _D2Vault.contract.FilterLogs(opts, "PreCrimeSet")
	if err != nil {
		return nil, err
	}
	return &D2VaultPreCrimeSetIterator{contract: _D2Vault.contract, event: "PreCrimeSet", logs: logs, sub: sub}, nil
}

// WatchPreCrimeSet is a free log subscription operation binding the contract event 0xd48d879cef83a1c0bdda516f27b13ddb1b3f8bbac1c9e1511bb2a659c2427760.
//
// Solidity: event PreCrimeSet(address preCrimeAddress)
func (_D2Vault *D2VaultFilterer) WatchPreCrimeSet(opts *bind.WatchOpts, sink chan<- *D2VaultPreCrimeSet) (event.Subscription, error) {

	logs, sub, err := _D2Vault.contract.WatchLogs(opts, "PreCrimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D2VaultPreCrimeSet)
				if err := _D2Vault.contract.UnpackLog(event, "PreCrimeSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePreCrimeSet is a log parse operation binding the contract event 0xd48d879cef83a1c0bdda516f27b13ddb1b3f8bbac1c9e1511bb2a659c2427760.
//
// Solidity: event PreCrimeSet(address preCrimeAddress)
func (_D2Vault *D2VaultFilterer) ParsePreCrimeSet(log types.Log) (*D2VaultPreCrimeSet, error) {
	event := new(D2VaultPreCrimeSet)
	if err := _D2Vault.contract.UnpackLog(event, "PreCrimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// D2VaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the D2Vault contract.
type D2VaultTransferIterator struct {
	Event *D2VaultTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *D2VaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D2VaultTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(D2VaultTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *D2VaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D2VaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D2VaultTransfer represents a Transfer event raised by the D2Vault contract.
type D2VaultTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_D2Vault *D2VaultFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*D2VaultTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _D2Vault.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &D2VaultTransferIterator{contract: _D2Vault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_D2Vault *D2VaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *D2VaultTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _D2Vault.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D2VaultTransfer)
				if err := _D2Vault.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_D2Vault *D2VaultFilterer) ParseTransfer(log types.Log) (*D2VaultTransfer, error) {
	event := new(D2VaultTransfer)
	if err := _D2Vault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// D2VaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the D2Vault contract.
type D2VaultWithdrawIterator struct {
	Event *D2VaultWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *D2VaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D2VaultWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(D2VaultWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *D2VaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D2VaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D2VaultWithdraw represents a Withdraw event raised by the D2Vault contract.
type D2VaultWithdraw struct {
	Sender   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_D2Vault *D2VaultFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*D2VaultWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _D2Vault.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &D2VaultWithdrawIterator{contract: _D2Vault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_D2Vault *D2VaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *D2VaultWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _D2Vault.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D2VaultWithdraw)
				if err := _D2Vault.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_D2Vault *D2VaultFilterer) ParseWithdraw(log types.Log) (*D2VaultWithdraw, error) {
	event := new(D2VaultWithdraw)
	if err := _D2Vault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
