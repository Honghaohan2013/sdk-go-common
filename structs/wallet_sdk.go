/*
Copyright ArxanFintech Technology Ltd. 2017-2018 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package structs

import (
	"net/http"

	"github.com/arxanchain/sdk-go-common/crypto/sign/ed25519"
	"github.com/arxanchain/sdk-go-common/protos/wallet"
)

////////////////////////////////////////////////////////////////////////////////
// Wallet Client Structs

// Register wallet request structure
type RegisterWalletBody struct {
	Id        Identifier         `json:"id"`         //Optional, if empty, generated by wallet service
	Type      DidType            `json:"type"`       //Wallet Type: 1. "Organization"; 2. "Dependent"; 3. "Independent"; 4. "Asset"
	Access    string             `json:"access"`     //Register user name
	Phone     string             `json:"phone"`      //Register user phone
	Email     string             `json:"email"`      //Register user email
	Secret    string             `json:"secret"`     //Register user passwd
	Metadata  interface{}        `json:"meta_data"`  //Register user metadata
	PublicKey *ed25519.PublicKey `json:"public_key"` //User public key. Optional, if empty, keypaire auto generated by wallet service
}

// Register subwallet request structure
type RegisterSubWalletBody struct {
	Id        Identifier         `json:"id"`         //main wallet id
	Type      DidType            `json:"type"`       //Wallet Type: 1."cash"; 2."fee"; 3."loan"; 4."interest"
	PublicKey *ed25519.PublicKey `json:"public_key"` //User public key. Optional, if empty, keypaire auto generated by wallet service
}

// WalletRequest common struct for transfer
type WalletRequest struct {
	Payload   string         `json:"payload"`
	Signature *SignatureBody `json:"signature"`
}

// WalletResponse common struct for wallet API response
type WalletResponse struct {
	Code           int        `json:"code"`
	Message        string     `json:"message"`
	Id             Identifier `json:"id"`
	Endpoint       string     `json:"endpoint"`
	KeyPair        *KeyPair   `json:"key_pair"`
	Created        int64      `json:"created"`
	TokenId        string     `json:"token_id"`
	TransactionIds []string   `json:"transaction_ids"`
}

type KeyPair struct {
	PrivateKey string `json:"private_key"` //base64 encoded ed25519 private key
	PublicKey  string `json:"public_key"`  //base64 encoded ed25519 public key
}

// TokenBalance ...
type CTokenBalance struct {
	Id     string `json:"id"`     //ctoken id
	Amount int64  `json:"amount"` //ctoken amount
}

// AssetBalance ...
type AssetBalance struct {
	Id     string `json:"id"`     //asset id
	Amount int64  `json:"amount"` //asset amount
	Name   string `json:"name"`   //asset name
	Status int    `json:"status"` //asset status
}

type WalletBalance struct {
	ColoredTokens map[string]*CTokenBalance `json:"colored_tokens"` //all the colored tokens in wallet
	DigitalAssets map[string]*AssetBalance  `json:"digital_assets"` //all the digital assets in wallet
}

type WalletInfo struct {
	Id       Identifier                 `json:"id"`
	Type     DidType                    `json:"type"`
	Endpoint DidEndpoint                `json:"endpoint"`
	Status   DidStatus                  `json:"status"`
	Created  int64                      `json:"created"`
	Updated  int64                      `json:"updated"`
	HDS      map[Identifier]*WalletInfo `json:"hds"`
}

// IWalletClient defines the behaviors implemented by wallet sdk
type IWalletClient interface {
	// Register is used to register user wallet.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	Register(http.Header, *RegisterWalletBody) (*WalletResponse, error)

	// RegisterSubWallet is used to register user subwallet.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	RegisterSubWallet(http.Header, *RegisterSubWalletBody) (*WalletResponse, error)

	// GetWalletBalance is used to get wallet balances.
	//
	GetWalletBalance(http.Header, Identifier) (*WalletBalance, error)

	// GetWalletInfo is used to get wallet base information.
	//
	GetWalletInfo(http.Header, Identifier) (*WalletInfo, error)

	// CreatePOE is used to create POE digital asset.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	// The signature is generated by caller using
	// 'github.com/arxanchain/sdk-go-common/crypto/tools/sign-util' tool.
	//
	CreatePOE(http.Header, *POEBody, *SignatureBody) (*WalletResponse, error)

	// CreatePOESign is used to create POE digital asset.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	// The signature is generated by SDK, need to pass the user private key to the SDK.
	//
	CreatePOESign(http.Header, *POEBody, *SignatureParam) (*WalletResponse, error)

	// UpdatePOE is used to update POE digital asset.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	// The signature is generated by caller using
	// 'github.com/arxanchain/sdk-go-common/crypto/tools/sign-util' tool.
	//
	UpdatePOE(http.Header, *POEBody, *SignatureBody) (*WalletResponse, error)

	// UpdatePOESign is used to update POE digital asset.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	// The signature is generated by SDK, need to pass the user private key to the SDK.
	//
	UpdatePOESign(http.Header, *POEBody, *SignatureParam) (*WalletResponse, error)

	// QueryPOE is used to query POE digital asset.
	//
	QueryPOE(http.Header, Identifier) (*POEPayload, error)

	// UploadPOEFile is used to upload file for specified POE digital asset
	//
	// poeID parameter is the POE digital asset ID pre-created using CreatePOE API.
	//
	// poeFile parameter is the path to file to be uploaded.
	//
	UploadPOEFile(header http.Header, poeID string, poeFile string) (*WalletResponse, error)

	// IssueCToken is used to issue colored token.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	// The signature is generated by caller using
	// 'github.com/arxanchain/sdk-go-common/crypto/tools/sign-util' tool.
	//
	IssueCToken(http.Header, *IssueBody, *SignatureBody) (*WalletResponse, error)

	// IssueCTokenSign is used to issue colored token.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	// The signature is generated by SDK, need to pass the user private key to the SDK.
	//
	IssueCTokenSign(http.Header, *IssueBody, *SignatureParam) (*WalletResponse, error)

	// IssueAsset is used to issue digital asset.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	// The signature is generated by caller using
	// 'github.com/arxanchain/sdk-go-common/crypto/tools/sign-util' tool.
	//
	IssueAsset(http.Header, *IssueAssetBody, *SignatureBody) (*WalletResponse, error)

	// IssueAssetSign is used to issue digital asset.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	// The signature is generated by SDK, need to pass the user private key to the SDK.
	//
	IssueAssetSign(http.Header, *IssueAssetBody, *SignatureParam) (*WalletResponse, error)

	// TransferCToken is used to transfer colored tokens from one user to another.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	// The signature is generated by caller using
	// 'github.com/arxanchain/sdk-go-common/crypto/tools/sign-util' tool.
	//
	TransferCToken(http.Header, *TransferCTokenBody, *SignatureBody) (*WalletResponse, error)

	// TransferCTokenSign is used to transfer colored tokens from one user to another.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	// The signature is generated by SDK, need to pass the user private key to the SDK.
	//
	TransferCTokenSign(http.Header, *TransferCTokenBody, *SignatureParam) (*WalletResponse, error)

	// TransferAsset is used to transfer assets from one user to another.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	// The signature is generated by caller using
	// 'github.com/arxanchain/sdk-go-common/crypto/tools/sign-util' tool.
	//
	TransferAsset(http.Header, *TransferAssetBody, *SignatureBody) (*WalletResponse, error)

	// TransferAssetSign is used to transfer assets from one user to another.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	// The signature is generated by SDK, need to pass the user private key to the SDK.
	//
	TransferAssetSign(http.Header, *TransferAssetBody, *SignatureParam) (*WalletResponse, error)

	// QueryTransactionLogs is used to query transaction logs.
	//
	// txType:
	// in: query income type transaction
	// out: query spending type transaction
	//
	QueryTransactionLogs(http.Header, Identifier, string) (TransactionLogs, error)
}

////////////////////////////////////////////////////////////////////////////////
// POE Client Structs

// POEBody POE request body structure definition
type POEBody struct {
	Id         Identifier `json:"id"`
	Name       string     `json:"name"`
	ParentId   Identifier `json:"parent_id"`
	Owner      Identifier `json:"owner"`
	ExpireTime int64      `json:"expire_time"`
	Hash       string     `json:"hash"`
	Metadata   []byte     `json:"metadata"`
}

// POEPayload POE query payload structure definition
type POEPayload struct {
	Id         Identifier `json:"id"`
	Name       string     `json:"name"`
	ParentId   Identifier `json:"parent_id"`
	Owner      Identifier `json:"owner"`
	ExpireTime int64      `json:"expire_time"`
	Hash       string     `json:"hash"`
	Metadata   []byte     `json:"metadata"`
	Created    int64      `json:"created"`
	Updated    int64      `json:"updated"`
	Status     DidStatus  `json:"status"`
}

// OffchainMetadata offchain storage metadata
type OffchainMetadata struct {
	Filename    string `json:"filename"`
	Endpoint    string `json:"endpoint"`
	StorageType string `json:"storageType"`
	ContentHash string `json:"contentHash"`
	Size        int    `json:"size"`
}

const (
	// OffchainPOEID poe did when invoke upload file api formdata param
	OffchainPOEID = "poe_id"
	// OffchainPOEFile poe binary file when invoke upload file api formdata param
	OffchainPOEFile = "poe_file"
)

// signature const for the formdata
const (
	// SignatureCreator issuer did
	SignatureCreator = "signature.creator"
	// SignatureCreated timestamp
	SignatureCreated = "signature.created"
	// SignatureNonce nonce
	SignatureNonce = "signature.nonce"
	// SignatureSignatureValue sign value
	SignatureSignatureValue = "signature.signatureValue"
)

type TransactionLogs map[string]*TransactionLog // key is endpoint

type TransactionLog struct {
	Utxo []*UTXO       `json:"utxo"` // unspent transaction output
	Stxo []*SpentTxOUT `json:"stxo"` // spent transaction output
}

type UTXO struct {
	// SourceTxDataHash the Bitcoin hash (double sha256) of
	// the given transaction
	SourceTxDataHash string `protobuf:"bytes,1,opt,name=sourceTxDataHash" json:"sourceTxDataHash,omitempty" `
	// Ix index of output array in the transaction
	Ix uint32 `protobuf:"varint,2,opt,name=ix" json:"ix,omitempty" `
	// ColoredToken ID
	CTokenId string `protobuf:"bytes,3,opt,name=cTokenId" json:"cTokenId,omitempty" `
	// ColorType
	CType int32 `protobuf:"varint,4,opt,name=cType" json:"cType,omitempty"`
	// token amount
	Value int64 `protobuf:"varint,4,opt,name=value" json:"value,omitempty"`
	// who will receive this txout
	Addr string `protobuf:"bytes,5,opt,name=addr" json:"addr,omitempty" `
	// until xx timestamp, any one cant spend the txout
	// -1 means no check
	Until int64 `protobuf:"varint,6,opt,name=until" json:"until,omitempty"`
	// script
	Script []byte `protobuf:"bytes,7,opt,name=script,proto3" json:"script,omitempty"`
	// CreatedAt
	CreatedAt *Timestamp `protobuf:"bytes,8,opt,name=createdAt" json:"createdAt,omitempty"`
	// Founder who created this tx
	Founder string `protobuf:"bytes,9,opt,name=founder" json:"founder,omitempty" `
	TxType  int32  `protobuf:"varint,10,opt,name=txType" json:"txType,omitempty"`
	// BCTxID blockchain transaction id
	BCTxID string `protobuf:"bytes,11,opt,name=bcTxID" json:"bcTxID,omitempty"`
}

// SpentTxOUT
type SpentTxOUT struct {
	// SourceTxDataHash the Bitcoin hash (double sha256) of
	// the given transaction
	SourceTxDataHash string `protobuf:"bytes,1,opt,name=sourceTxDataHash" json:"sourceTxDataHash,omitempty" `
	// Ix index of output array in the transaction
	Ix uint32 `protobuf:"varint,2,opt,name=ix" json:"ix,omitempty" `
	// ColoredToken ID
	CTokenId string `protobuf:"bytes,3,opt,name=cTokenId" json:"cTokenId,omitempty" `
	// ColorType
	CType int32 `protobuf:"varint,4,opt,name=cType" json:"cType,omitempty"`
	// token amount
	Value int64 `protobuf:"varint,4,opt,name=value" json:"value,omitempty"`
	// who will receive this txout
	Addr string `protobuf:"bytes,5,opt,name=addr" json:"addr,omitempty" `
	// until xx timestamp, any one cant spend the txout
	// -1 means no check
	Until int64 `protobuf:"varint,6,opt,name=until" json:"until,omitempty"`
	// script
	Script []byte `protobuf:"bytes,7,opt,name=script,proto3" json:"script,omitempty"`
	// CreatedAt
	CreatedAt *Timestamp `protobuf:"bytes,8,opt,name=createdAt" json:"createdAt,omitempty"`
	// SpentTxDataHash
	SpentTxDataHash string `protobuf:"bytes,9,opt,name=spentTxDataHash" json:"spentTxDataHash,omitempty" `
	// SpentAt ...
	SpentAt *Timestamp `protobuf:"bytes,10,opt,name=spentAt" json:"spentAt,omitempty"`
	// Founder who created this tx
	Founder string `protobuf:"bytes,11,opt,name=founder" json:"founder,omitempty"`
	TxType  int32  `protobuf:"varint,12,opt,name=txType" json:"txType,omitempty"`
	// BCTxID blockchain transaction id
	BCTxID string `protobuf:"bytes,13,opt,name=bcTxID" json:"bcTxID,omitempty"`
}

// Transaction Structs

type AXTUnit int64

const (
	ATOM     AXTUnit = 1
	MicroAXT         = 1000 * ATOM
	AXT              = 1000 * MicroAXT
)

// Colored Token Amount Structure
type TokenAmount struct {
	TokenId string `json:"token_id"`
	Amount  int64  `json:"amount"`
}

// Transaction Fee Structure
type Fee struct {
	Amount AXTUnit `json:"amount"`
}

// Issue Asset Request Structure
type IssueAssetBody struct {
	Issuer  string `json:"issuer"`
	Owner   string `json:"owner"`
	AssetId string `json:"asset_id"`
	Fee     *Fee   `json:"fee"`
}

// Issue Colored Token Request Structure
type IssueBody struct {
	Issuer  string `json:"issuer"`
	Owner   string `json:"owner"`
	AssetId string `json:"asset_id"`
	Amount  int64  `json:"amount"`
	Fee     *Fee   `json:"fee"`
}

// Transfer Colored Token Request Structure
type TransferCTokenBody struct {
	From    string         `json:"from"`
	To      string         `json:"to"`
	AssetId string         `json:"asset_id"`
	Tokens  []*TokenAmount `json:"tokens"`
	Fee     *Fee           `json:"fee"`
}

// Transfer to process Tx Request Structure
type ProcessTxBody struct {
	Txs []*wallet.TX `json:"txs"`
}

// Transfer Asset Request Structure
type TransferAssetBody struct {
	From   string   `json:"from"`
	To     string   `json:"to"`
	Assets []string `json:"assets"`
	Fee    *Fee     `json:"fee"`
}

// Timestamp Structure
type Timestamp struct {
	Seconds int64 `json:"seconds"`
	Nanos   int32 `json:"nanos"`
}

type IssueCTokenPrepareResponse struct {
	TokenId string       `json:"token_id"`
	Txs     []*wallet.TX `json:"txs"`
}
