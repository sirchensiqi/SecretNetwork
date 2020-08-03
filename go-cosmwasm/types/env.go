package types

import "github.com/enigmampc/cosmos-sdk/x/auth"

//---------- Env ---------

// Env defines the state of the blockchain environment this contract is
// running in. This must contain only trusted data - nothing from the Tx itself
// that has not been verfied (like Signer).
//
// Env are json encoded to a byte slice before passing to the wasm contract.
type Env struct {
	Block             BlockInfo           `json:"block"`
	Message           MessageInfo         `json:"message"`
	Contract          ContractInfo        `json:"contract"`
	Key               ContractKey         `json:"contract_key"`
	Bytes             [][]byte            `json:"sign_bytes"`
	Signatures        []auth.StdSignature `json:"signatures"`
	CallbackSignature []byte              `json:"cb_signature" yaml:"signature"` // Optional

}

type ContractKey string

type BlockInfo struct {
	// block height this transaction is executed
	Height uint64 `json:"height"`
	// time in seconds since unix epoch - since cosmwasm 0.3
	Time    uint64 `json:"time"`
	ChainID string `json:"chain_id"`
}

type MessageInfo struct {
	// binary encoding of sdk.AccAddress executing the contract
	Sender CanonicalAddress `json:"sender"`
	// amount of funds send to the contract along with this message
	SentFunds Coins `json:"sent_funds"`
}

type ContractInfo struct {
	// binary encoding of sdk.AccAddress of the contract, to be used when sending messages
	Address CanonicalAddress `json:"address"`
}
