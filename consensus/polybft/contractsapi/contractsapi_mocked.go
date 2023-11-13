// H_MODIFY: Code is copied from the latest upstream version of Edge. We need it to remove compile errors
package contractsapi

import (
	"math/big"

	"github.com/0xPolygon/polygon-edge/types"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/abi"
)

type CheckpointMetadata struct {
	BlockHash               types.Hash `abi:"blockHash"`
	BlockRound              *big.Int   `abi:"blockRound"`
	CurrentValidatorSetHash types.Hash `abi:"currentValidatorSetHash"`
}

var CheckpointMetadataABIType = abi.MustNewType("tuple(bytes32 blockHash,uint256 blockRound,bytes32 currentValidatorSetHash)")

func (c *CheckpointMetadata) EncodeAbi() ([]byte, error) {
	return CheckpointMetadataABIType.Encode(c)
}

func (c *CheckpointMetadata) DecodeAbi(buf []byte) error {
	return decodeStruct(CheckpointMetadataABIType, buf, &c)
}

type Checkpoint struct {
	Epoch       *big.Int   `abi:"epoch"`
	BlockNumber *big.Int   `abi:"blockNumber"`
	EventRoot   types.Hash `abi:"eventRoot"`
}

var CheckpointABIType = abi.MustNewType("tuple(uint256 epoch,uint256 blockNumber,bytes32 eventRoot)")

func (c *Checkpoint) EncodeAbi() ([]byte, error) {
	return CheckpointABIType.Encode(c)
}

func (c *Checkpoint) DecodeAbi(buf []byte) error {
	return decodeStruct(CheckpointABIType, buf, &c)
}

type Validator struct {
	Address     types.Address `abi:"_address"`
	BlsKey      [4]*big.Int   `abi:"blsKey"`
	VotingPower *big.Int      `abi:"votingPower"`
}

var ValidatorABIType = abi.MustNewType("tuple(address _address,uint256[4] blsKey,uint256 votingPower)")

func (v *Validator) EncodeAbi() ([]byte, error) {
	return ValidatorABIType.Encode(v)
}

func (v *Validator) DecodeAbi(buf []byte) error {
	return decodeStruct(ValidatorABIType, buf, &v)
}

type SubmitCheckpointManagerFn struct {
	CheckpointMetadata *CheckpointMetadata `abi:"checkpointMetadata"`
	Checkpoint         *Checkpoint         `abi:"checkpoint"`
	Signature          [2]*big.Int         `abi:"signature"`
	NewValidatorSet    []*Validator        `abi:"newValidatorSet"`
	Bitmap             []byte              `abi:"bitmap"`
}

func (s *SubmitCheckpointManagerFn) Sig() []byte {
	return CheckpointManager.Abi.Methods["submit"].ID()
}

func (s *SubmitCheckpointManagerFn) EncodeAbi() ([]byte, error) {
	return CheckpointManager.Abi.Methods["submit"].Encode(s)
}

func (s *SubmitCheckpointManagerFn) DecodeAbi(buf []byte) error {
	return decodeMethod(CheckpointManager.Abi.Methods["submit"], buf, s)
}

type InitializeCheckpointManagerFn struct {
	NewBls          types.Address `abi:"newBls"`
	NewBn256G2      types.Address `abi:"newBn256G2"`
	ChainID_        *big.Int      `abi:"chainId_"`
	NewValidatorSet []*Validator  `abi:"newValidatorSet"`
}

func (i *InitializeCheckpointManagerFn) Sig() []byte {
	return CheckpointManager.Abi.Methods["initialize"].ID()
}

func (i *InitializeCheckpointManagerFn) EncodeAbi() ([]byte, error) {
	return CheckpointManager.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeCheckpointManagerFn) DecodeAbi(buf []byte) error {
	return decodeMethod(CheckpointManager.Abi.Methods["initialize"], buf, i)
}

type GetCheckpointBlockCheckpointManagerFn struct {
	BlockNumber *big.Int `abi:"blockNumber"`
}

func (g *GetCheckpointBlockCheckpointManagerFn) Sig() []byte {
	return CheckpointManager.Abi.Methods["getCheckpointBlock"].ID()
}

func (g *GetCheckpointBlockCheckpointManagerFn) EncodeAbi() ([]byte, error) {
	return CheckpointManager.Abi.Methods["getCheckpointBlock"].Encode(g)
}

func (g *GetCheckpointBlockCheckpointManagerFn) DecodeAbi(buf []byte) error {
	return decodeMethod(CheckpointManager.Abi.Methods["getCheckpointBlock"], buf, g)
}

type InitializeChildERC20PredicateFn struct {
	NewL2StateSender          types.Address `abi:"newL2StateSender"`
	NewStateReceiver          types.Address `abi:"newStateReceiver"`
	NewRootERC20Predicate     types.Address `abi:"newRootERC20Predicate"`
	NewChildTokenTemplate     types.Address `abi:"newChildTokenTemplate"`
	NewNativeTokenRootAddress types.Address `abi:"newNativeTokenRootAddress"`
}

func (i *InitializeChildERC20PredicateFn) Sig() []byte {
	return ChildERC20Predicate.Abi.Methods["initialize"].ID()
}

func (i *InitializeChildERC20PredicateFn) EncodeAbi() ([]byte, error) {
	return ChildERC20Predicate.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeChildERC20PredicateFn) DecodeAbi(buf []byte) error {
	return decodeMethod(ChildERC20Predicate.Abi.Methods["initialize"], buf, i)
}

type WithdrawToChildERC20PredicateFn struct {
	ChildToken types.Address `abi:"childToken"`
	Receiver   types.Address `abi:"receiver"`
	Amount     *big.Int      `abi:"amount"`
}

func (w *WithdrawToChildERC20PredicateFn) Sig() []byte {
	return ChildERC20Predicate.Abi.Methods["withdrawTo"].ID()
}

func (w *WithdrawToChildERC20PredicateFn) EncodeAbi() ([]byte, error) {
	return ChildERC20Predicate.Abi.Methods["withdrawTo"].Encode(w)
}

func (w *WithdrawToChildERC20PredicateFn) DecodeAbi(buf []byte) error {
	return decodeMethod(ChildERC20Predicate.Abi.Methods["withdrawTo"], buf, w)
}

type InitializeChildERC20PredicateACLFn struct {
	NewL2StateSender          types.Address `abi:"newL2StateSender"`
	NewStateReceiver          types.Address `abi:"newStateReceiver"`
	NewRootERC20Predicate     types.Address `abi:"newRootERC20Predicate"`
	NewChildTokenTemplate     types.Address `abi:"newChildTokenTemplate"`
	NewNativeTokenRootAddress types.Address `abi:"newNativeTokenRootAddress"`
	NewUseAllowList           bool          `abi:"newUseAllowList"`
	NewUseBlockList           bool          `abi:"newUseBlockList"`
	NewOwner                  types.Address `abi:"newOwner"`
}

func (i *InitializeChildERC20PredicateACLFn) Sig() []byte {
	return ChildERC20PredicateACL.Abi.Methods["initialize"].ID()
}

func (i *InitializeChildERC20PredicateACLFn) EncodeAbi() ([]byte, error) {
	return ChildERC20PredicateACL.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeChildERC20PredicateACLFn) DecodeAbi(buf []byte) error {
	return decodeMethod(ChildERC20PredicateACL.Abi.Methods["initialize"], buf, i)
}

type WithdrawToChildERC20PredicateACLFn struct {
	ChildToken types.Address `abi:"childToken"`
	Receiver   types.Address `abi:"receiver"`
	Amount     *big.Int      `abi:"amount"`
}

func (w *WithdrawToChildERC20PredicateACLFn) Sig() []byte {
	return ChildERC20PredicateACL.Abi.Methods["withdrawTo"].ID()
}

func (w *WithdrawToChildERC20PredicateACLFn) EncodeAbi() ([]byte, error) {
	return ChildERC20PredicateACL.Abi.Methods["withdrawTo"].Encode(w)
}

func (w *WithdrawToChildERC20PredicateACLFn) DecodeAbi(buf []byte) error {
	return decodeMethod(ChildERC20PredicateACL.Abi.Methods["withdrawTo"], buf, w)
}

type InitializeRootMintableERC20PredicateFn struct {
	NewL2StateSender       types.Address `abi:"newL2StateSender"`
	NewStateReceiver       types.Address `abi:"newStateReceiver"`
	NewChildERC20Predicate types.Address `abi:"newChildERC20Predicate"`
	NewChildTokenTemplate  types.Address `abi:"newChildTokenTemplate"`
}

func (i *InitializeRootMintableERC20PredicateFn) Sig() []byte {
	return RootMintableERC20Predicate.Abi.Methods["initialize"].ID()
}

func (i *InitializeRootMintableERC20PredicateFn) EncodeAbi() ([]byte, error) {
	return RootMintableERC20Predicate.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeRootMintableERC20PredicateFn) DecodeAbi(buf []byte) error {
	return decodeMethod(RootMintableERC20Predicate.Abi.Methods["initialize"], buf, i)
}

type InitializeRootMintableERC20PredicateACLFn struct {
	NewL2StateSender       types.Address `abi:"newL2StateSender"`
	NewStateReceiver       types.Address `abi:"newStateReceiver"`
	NewChildERC20Predicate types.Address `abi:"newChildERC20Predicate"`
	NewChildTokenTemplate  types.Address `abi:"newChildTokenTemplate"`
	NewUseAllowList        bool          `abi:"newUseAllowList"`
	NewUseBlockList        bool          `abi:"newUseBlockList"`
	NewOwner               types.Address `abi:"newOwner"`
}

func (i *InitializeRootMintableERC20PredicateACLFn) Sig() []byte {
	return RootMintableERC20PredicateACL.Abi.Methods["initialize"].ID()
}

func (i *InitializeRootMintableERC20PredicateACLFn) EncodeAbi() ([]byte, error) {
	return RootMintableERC20PredicateACL.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeRootMintableERC20PredicateACLFn) DecodeAbi(buf []byte) error {
	return decodeMethod(RootMintableERC20PredicateACL.Abi.Methods["initialize"], buf, i)
}

type InitializeNativeERC20Fn struct {
	Predicate_   types.Address `abi:"predicate_"`
	RootToken_   types.Address `abi:"rootToken_"`
	Name_        string        `abi:"name_"`
	Symbol_      string        `abi:"symbol_"`
	Decimals_    uint8         `abi:"decimals_"`
	TokenSupply_ *big.Int      `abi:"tokenSupply_"`
}

func (i *InitializeNativeERC20Fn) Sig() []byte {
	return NativeERC20.Abi.Methods["initialize"].ID()
}

func (i *InitializeNativeERC20Fn) EncodeAbi() ([]byte, error) {
	return NativeERC20.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeNativeERC20Fn) DecodeAbi(buf []byte) error {
	return decodeMethod(NativeERC20.Abi.Methods["initialize"], buf, i)
}

type InitializeNativeERC20MintableFn struct {
	Predicate_   types.Address `abi:"predicate_"`
	Owner_       types.Address `abi:"owner_"`
	RootToken_   types.Address `abi:"rootToken_"`
	Name_        string        `abi:"name_"`
	Symbol_      string        `abi:"symbol_"`
	Decimals_    uint8         `abi:"decimals_"`
	TokenSupply_ *big.Int      `abi:"tokenSupply_"`
}

func (i *InitializeNativeERC20MintableFn) Sig() []byte {
	return NativeERC20Mintable.Abi.Methods["initialize"].ID()
}

func (i *InitializeNativeERC20MintableFn) EncodeAbi() ([]byte, error) {
	return NativeERC20Mintable.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeNativeERC20MintableFn) DecodeAbi(buf []byte) error {
	return decodeMethod(NativeERC20Mintable.Abi.Methods["initialize"], buf, i)
}

type InitializeRootERC20PredicateFn struct {
	NewStateSender         types.Address `abi:"newStateSender"`
	NewExitHelper          types.Address `abi:"newExitHelper"`
	NewChildERC20Predicate types.Address `abi:"newChildERC20Predicate"`
	NewChildTokenTemplate  types.Address `abi:"newChildTokenTemplate"`
	NativeTokenRootAddress types.Address `abi:"nativeTokenRootAddress"`
}

func (i *InitializeRootERC20PredicateFn) Sig() []byte {
	return RootERC20Predicate.Abi.Methods["initialize"].ID()
}

func (i *InitializeRootERC20PredicateFn) EncodeAbi() ([]byte, error) {
	return RootERC20Predicate.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeRootERC20PredicateFn) DecodeAbi(buf []byte) error {
	return decodeMethod(RootERC20Predicate.Abi.Methods["initialize"], buf, i)
}

type DepositToRootERC20PredicateFn struct {
	RootToken types.Address `abi:"rootToken"`
	Receiver  types.Address `abi:"receiver"`
	Amount    *big.Int      `abi:"amount"`
}

func (d *DepositToRootERC20PredicateFn) Sig() []byte {
	return RootERC20Predicate.Abi.Methods["depositTo"].ID()
}

func (d *DepositToRootERC20PredicateFn) EncodeAbi() ([]byte, error) {
	return RootERC20Predicate.Abi.Methods["depositTo"].Encode(d)
}

func (d *DepositToRootERC20PredicateFn) DecodeAbi(buf []byte) error {
	return decodeMethod(RootERC20Predicate.Abi.Methods["depositTo"], buf, d)
}

type TokenMappedEvent struct {
	RootToken  types.Address `abi:"rootToken"`
	ChildToken types.Address `abi:"childToken"`
}

func (*TokenMappedEvent) Sig() ethgo.Hash {
	return RootERC20Predicate.Abi.Events["TokenMapped"].ID()
}

func (*TokenMappedEvent) Encode(inputs interface{}) ([]byte, error) {
	return RootERC20Predicate.Abi.Events["TokenMapped"].Inputs.Encode(inputs)
}

func (t *TokenMappedEvent) ParseLog(log *ethgo.Log) (bool, error) {
	if !RootERC20Predicate.Abi.Events["TokenMapped"].Match(log) {
		return false, nil
	}

	return true, decodeEvent(RootERC20Predicate.Abi.Events["TokenMapped"], log, t)
}

type InitializeChildMintableERC20PredicateFn struct {
	NewStateSender        types.Address `abi:"newStateSender"`
	NewExitHelper         types.Address `abi:"newExitHelper"`
	NewRootERC20Predicate types.Address `abi:"newRootERC20Predicate"`
	NewChildTokenTemplate types.Address `abi:"newChildTokenTemplate"`
}

func (i *InitializeChildMintableERC20PredicateFn) Sig() []byte {
	return ChildMintableERC20Predicate.Abi.Methods["initialize"].ID()
}

func (i *InitializeChildMintableERC20PredicateFn) EncodeAbi() ([]byte, error) {
	return ChildMintableERC20Predicate.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeChildMintableERC20PredicateFn) DecodeAbi(buf []byte) error {
	return decodeMethod(ChildMintableERC20Predicate.Abi.Methods["initialize"], buf, i)
}

type MintableTokenMappedEvent struct {
	RootToken  types.Address `abi:"rootToken"`
	ChildToken types.Address `abi:"childToken"`
}

func (*MintableTokenMappedEvent) Sig() ethgo.Hash {
	return ChildMintableERC20Predicate.Abi.Events["MintableTokenMapped"].ID()
}

func (*MintableTokenMappedEvent) Encode(inputs interface{}) ([]byte, error) {
	return ChildMintableERC20Predicate.Abi.Events["MintableTokenMapped"].Inputs.Encode(inputs)
}

func (m *MintableTokenMappedEvent) ParseLog(log *ethgo.Log) (bool, error) {
	if !ChildMintableERC20Predicate.Abi.Events["MintableTokenMapped"].Match(log) {
		return false, nil
	}

	return true, decodeEvent(ChildMintableERC20Predicate.Abi.Events["MintableTokenMapped"], log, m)
}

type BalanceOfRootERC20Fn struct {
	Account types.Address `abi:"account"`
}

func (b *BalanceOfRootERC20Fn) Sig() []byte {
	return RootERC20.Abi.Methods["balanceOf"].ID()
}

func (b *BalanceOfRootERC20Fn) EncodeAbi() ([]byte, error) {
	return RootERC20.Abi.Methods["balanceOf"].Encode(b)
}

func (b *BalanceOfRootERC20Fn) DecodeAbi(buf []byte) error {
	return decodeMethod(RootERC20.Abi.Methods["balanceOf"], buf, b)
}

type ApproveRootERC20Fn struct {
	Spender types.Address `abi:"spender"`
	Amount  *big.Int      `abi:"amount"`
}

func (a *ApproveRootERC20Fn) Sig() []byte {
	return RootERC20.Abi.Methods["approve"].ID()
}

func (a *ApproveRootERC20Fn) EncodeAbi() ([]byte, error) {
	return RootERC20.Abi.Methods["approve"].Encode(a)
}

func (a *ApproveRootERC20Fn) DecodeAbi(buf []byte) error {
	return decodeMethod(RootERC20.Abi.Methods["approve"], buf, a)
}

type MintRootERC20Fn struct {
	To     types.Address `abi:"to"`
	Amount *big.Int      `abi:"amount"`
}

func (m *MintRootERC20Fn) Sig() []byte {
	return RootERC20.Abi.Methods["mint"].ID()
}

func (m *MintRootERC20Fn) EncodeAbi() ([]byte, error) {
	return RootERC20.Abi.Methods["mint"].Encode(m)
}

func (m *MintRootERC20Fn) DecodeAbi(buf []byte) error {
	return decodeMethod(RootERC20.Abi.Methods["mint"], buf, m)
}

type InitializeRootERC1155PredicateFn struct {
	NewStateSender           types.Address `abi:"newStateSender"`
	NewExitHelper            types.Address `abi:"newExitHelper"`
	NewChildERC1155Predicate types.Address `abi:"newChildERC1155Predicate"`
	NewChildTokenTemplate    types.Address `abi:"newChildTokenTemplate"`
}

func (i *InitializeRootERC1155PredicateFn) Sig() []byte {
	return RootERC1155Predicate.Abi.Methods["initialize"].ID()
}

func (i *InitializeRootERC1155PredicateFn) EncodeAbi() ([]byte, error) {
	return RootERC1155Predicate.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeRootERC1155PredicateFn) DecodeAbi(buf []byte) error {
	return decodeMethod(RootERC1155Predicate.Abi.Methods["initialize"], buf, i)
}

type DepositBatchRootERC1155PredicateFn struct {
	RootToken types.Address   `abi:"rootToken"`
	Receivers []ethgo.Address `abi:"receivers"`
	TokenIDs  []*big.Int      `abi:"tokenIds"`
	Amounts   []*big.Int      `abi:"amounts"`
}

func (d *DepositBatchRootERC1155PredicateFn) Sig() []byte {
	return RootERC1155Predicate.Abi.Methods["depositBatch"].ID()
}

func (d *DepositBatchRootERC1155PredicateFn) EncodeAbi() ([]byte, error) {
	return RootERC1155Predicate.Abi.Methods["depositBatch"].Encode(d)
}

func (d *DepositBatchRootERC1155PredicateFn) DecodeAbi(buf []byte) error {
	return decodeMethod(RootERC1155Predicate.Abi.Methods["depositBatch"], buf, d)
}

type InitializeChildMintableERC1155PredicateFn struct {
	NewStateSender          types.Address `abi:"newStateSender"`
	NewExitHelper           types.Address `abi:"newExitHelper"`
	NewRootERC1155Predicate types.Address `abi:"newRootERC1155Predicate"`
	NewChildTokenTemplate   types.Address `abi:"newChildTokenTemplate"`
}

func (i *InitializeChildMintableERC1155PredicateFn) Sig() []byte {
	return ChildMintableERC1155Predicate.Abi.Methods["initialize"].ID()
}

func (i *InitializeChildMintableERC1155PredicateFn) EncodeAbi() ([]byte, error) {
	return ChildMintableERC1155Predicate.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeChildMintableERC1155PredicateFn) DecodeAbi(buf []byte) error {
	return decodeMethod(ChildMintableERC1155Predicate.Abi.Methods["initialize"], buf, i)
}

type SetApprovalForAllRootERC1155Fn struct {
	Operator types.Address `abi:"operator"`
	Approved bool          `abi:"approved"`
}

func (s *SetApprovalForAllRootERC1155Fn) Sig() []byte {
	return RootERC1155.Abi.Methods["setApprovalForAll"].ID()
}

func (s *SetApprovalForAllRootERC1155Fn) EncodeAbi() ([]byte, error) {
	return RootERC1155.Abi.Methods["setApprovalForAll"].Encode(s)
}

func (s *SetApprovalForAllRootERC1155Fn) DecodeAbi(buf []byte) error {
	return decodeMethod(RootERC1155.Abi.Methods["setApprovalForAll"], buf, s)
}

type MintBatchRootERC1155Fn struct {
	To      types.Address `abi:"to"`
	IDs     []*big.Int    `abi:"ids"`
	Amounts []*big.Int    `abi:"amounts"`
	Data    []byte        `abi:"data"`
}

func (m *MintBatchRootERC1155Fn) Sig() []byte {
	return RootERC1155.Abi.Methods["mintBatch"].ID()
}

func (m *MintBatchRootERC1155Fn) EncodeAbi() ([]byte, error) {
	return RootERC1155.Abi.Methods["mintBatch"].Encode(m)
}

func (m *MintBatchRootERC1155Fn) DecodeAbi(buf []byte) error {
	return decodeMethod(RootERC1155.Abi.Methods["mintBatch"], buf, m)
}

type BalanceOfRootERC1155Fn struct {
	Account types.Address `abi:"account"`
	ID      *big.Int      `abi:"id"`
}

func (b *BalanceOfRootERC1155Fn) Sig() []byte {
	return RootERC1155.Abi.Methods["balanceOf"].ID()
}

func (b *BalanceOfRootERC1155Fn) EncodeAbi() ([]byte, error) {
	return RootERC1155.Abi.Methods["balanceOf"].Encode(b)
}

func (b *BalanceOfRootERC1155Fn) DecodeAbi(buf []byte) error {
	return decodeMethod(RootERC1155.Abi.Methods["balanceOf"], buf, b)
}

type InitializeChildERC1155PredicateFn struct {
	NewL2StateSender        types.Address `abi:"newL2StateSender"`
	NewStateReceiver        types.Address `abi:"newStateReceiver"`
	NewRootERC1155Predicate types.Address `abi:"newRootERC1155Predicate"`
	NewChildTokenTemplate   types.Address `abi:"newChildTokenTemplate"`
}

func (i *InitializeChildERC1155PredicateFn) Sig() []byte {
	return ChildERC1155Predicate.Abi.Methods["initialize"].ID()
}

func (i *InitializeChildERC1155PredicateFn) EncodeAbi() ([]byte, error) {
	return ChildERC1155Predicate.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeChildERC1155PredicateFn) DecodeAbi(buf []byte) error {
	return decodeMethod(ChildERC1155Predicate.Abi.Methods["initialize"], buf, i)
}

type WithdrawBatchChildERC1155PredicateFn struct {
	ChildToken types.Address   `abi:"childToken"`
	Receivers  []ethgo.Address `abi:"receivers"`
	TokenIDs   []*big.Int      `abi:"tokenIds"`
	Amounts    []*big.Int      `abi:"amounts"`
}

func (w *WithdrawBatchChildERC1155PredicateFn) Sig() []byte {
	return ChildERC1155Predicate.Abi.Methods["withdrawBatch"].ID()
}

func (w *WithdrawBatchChildERC1155PredicateFn) EncodeAbi() ([]byte, error) {
	return ChildERC1155Predicate.Abi.Methods["withdrawBatch"].Encode(w)
}

func (w *WithdrawBatchChildERC1155PredicateFn) DecodeAbi(buf []byte) error {
	return decodeMethod(ChildERC1155Predicate.Abi.Methods["withdrawBatch"], buf, w)
}

type InitializeChildERC1155PredicateACLFn struct {
	NewL2StateSender        types.Address `abi:"newL2StateSender"`
	NewStateReceiver        types.Address `abi:"newStateReceiver"`
	NewRootERC1155Predicate types.Address `abi:"newRootERC1155Predicate"`
	NewChildTokenTemplate   types.Address `abi:"newChildTokenTemplate"`
	NewUseAllowList         bool          `abi:"newUseAllowList"`
	NewUseBlockList         bool          `abi:"newUseBlockList"`
	NewOwner                types.Address `abi:"newOwner"`
}

func (i *InitializeChildERC1155PredicateACLFn) Sig() []byte {
	return ChildERC1155PredicateACL.Abi.Methods["initialize"].ID()
}

func (i *InitializeChildERC1155PredicateACLFn) EncodeAbi() ([]byte, error) {
	return ChildERC1155PredicateACL.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeChildERC1155PredicateACLFn) DecodeAbi(buf []byte) error {
	return decodeMethod(ChildERC1155PredicateACL.Abi.Methods["initialize"], buf, i)
}

type WithdrawBatchChildERC1155PredicateACLFn struct {
	ChildToken types.Address   `abi:"childToken"`
	Receivers  []ethgo.Address `abi:"receivers"`
	TokenIDs   []*big.Int      `abi:"tokenIds"`
	Amounts    []*big.Int      `abi:"amounts"`
}

func (w *WithdrawBatchChildERC1155PredicateACLFn) Sig() []byte {
	return ChildERC1155PredicateACL.Abi.Methods["withdrawBatch"].ID()
}

func (w *WithdrawBatchChildERC1155PredicateACLFn) EncodeAbi() ([]byte, error) {
	return ChildERC1155PredicateACL.Abi.Methods["withdrawBatch"].Encode(w)
}

func (w *WithdrawBatchChildERC1155PredicateACLFn) DecodeAbi(buf []byte) error {
	return decodeMethod(ChildERC1155PredicateACL.Abi.Methods["withdrawBatch"], buf, w)
}

type InitializeRootMintableERC1155PredicateFn struct {
	NewL2StateSender         types.Address `abi:"newL2StateSender"`
	NewStateReceiver         types.Address `abi:"newStateReceiver"`
	NewChildERC1155Predicate types.Address `abi:"newChildERC1155Predicate"`
	NewChildTokenTemplate    types.Address `abi:"newChildTokenTemplate"`
}

func (i *InitializeRootMintableERC1155PredicateFn) Sig() []byte {
	return RootMintableERC1155Predicate.Abi.Methods["initialize"].ID()
}

func (i *InitializeRootMintableERC1155PredicateFn) EncodeAbi() ([]byte, error) {
	return RootMintableERC1155Predicate.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeRootMintableERC1155PredicateFn) DecodeAbi(buf []byte) error {
	return decodeMethod(RootMintableERC1155Predicate.Abi.Methods["initialize"], buf, i)
}

type InitializeRootMintableERC1155PredicateACLFn struct {
	NewL2StateSender         types.Address `abi:"newL2StateSender"`
	NewStateReceiver         types.Address `abi:"newStateReceiver"`
	NewChildERC1155Predicate types.Address `abi:"newChildERC1155Predicate"`
	NewChildTokenTemplate    types.Address `abi:"newChildTokenTemplate"`
	NewUseAllowList          bool          `abi:"newUseAllowList"`
	NewUseBlockList          bool          `abi:"newUseBlockList"`
	NewOwner                 types.Address `abi:"newOwner"`
}

func (i *InitializeRootMintableERC1155PredicateACLFn) Sig() []byte {
	return RootMintableERC1155PredicateACL.Abi.Methods["initialize"].ID()
}

func (i *InitializeRootMintableERC1155PredicateACLFn) EncodeAbi() ([]byte, error) {
	return RootMintableERC1155PredicateACL.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeRootMintableERC1155PredicateACLFn) DecodeAbi(buf []byte) error {
	return decodeMethod(RootMintableERC1155PredicateACL.Abi.Methods["initialize"], buf, i)
}

type L2MintableTokenMappedEvent struct {
	RootToken  types.Address `abi:"rootToken"`
	ChildToken types.Address `abi:"childToken"`
}

func (*L2MintableTokenMappedEvent) Sig() ethgo.Hash {
	return RootMintableERC1155PredicateACL.Abi.Events["L2MintableTokenMapped"].ID()
}

func (*L2MintableTokenMappedEvent) Encode(inputs interface{}) ([]byte, error) {
	return RootMintableERC1155PredicateACL.Abi.Events["L2MintableTokenMapped"].Inputs.Encode(inputs)
}

func (l *L2MintableTokenMappedEvent) ParseLog(log *ethgo.Log) (bool, error) {
	if !RootMintableERC1155PredicateACL.Abi.Events["L2MintableTokenMapped"].Match(log) {
		return false, nil
	}

	return true, decodeEvent(RootMintableERC1155PredicateACL.Abi.Events["L2MintableTokenMapped"], log, l)
}

type InitializeChildERC1155Fn struct {
	RootToken_ types.Address `abi:"rootToken_"`
	Uri_       string        `abi:"uri_"`
}

func (i *InitializeChildERC1155Fn) Sig() []byte {
	return ChildERC1155.Abi.Methods["initialize"].ID()
}

func (i *InitializeChildERC1155Fn) EncodeAbi() ([]byte, error) {
	return ChildERC1155.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeChildERC1155Fn) DecodeAbi(buf []byte) error {
	return decodeMethod(ChildERC1155.Abi.Methods["initialize"], buf, i)
}

type BalanceOfChildERC1155Fn struct {
	Account types.Address `abi:"account"`
	ID      *big.Int      `abi:"id"`
}

func (b *BalanceOfChildERC1155Fn) Sig() []byte {
	return ChildERC1155.Abi.Methods["balanceOf"].ID()
}

func (b *BalanceOfChildERC1155Fn) EncodeAbi() ([]byte, error) {
	return ChildERC1155.Abi.Methods["balanceOf"].Encode(b)
}

func (b *BalanceOfChildERC1155Fn) DecodeAbi(buf []byte) error {
	return decodeMethod(ChildERC1155.Abi.Methods["balanceOf"], buf, b)
}

type InitializeRootERC721PredicateFn struct {
	NewStateSender          types.Address `abi:"newStateSender"`
	NewExitHelper           types.Address `abi:"newExitHelper"`
	NewChildERC721Predicate types.Address `abi:"newChildERC721Predicate"`
	NewChildTokenTemplate   types.Address `abi:"newChildTokenTemplate"`
}

func (i *InitializeRootERC721PredicateFn) Sig() []byte {
	return RootERC721Predicate.Abi.Methods["initialize"].ID()
}

func (i *InitializeRootERC721PredicateFn) EncodeAbi() ([]byte, error) {
	return RootERC721Predicate.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeRootERC721PredicateFn) DecodeAbi(buf []byte) error {
	return decodeMethod(RootERC721Predicate.Abi.Methods["initialize"], buf, i)
}

type DepositBatchRootERC721PredicateFn struct {
	RootToken types.Address   `abi:"rootToken"`
	Receivers []ethgo.Address `abi:"receivers"`
	TokenIDs  []*big.Int      `abi:"tokenIds"`
}

func (d *DepositBatchRootERC721PredicateFn) Sig() []byte {
	return RootERC721Predicate.Abi.Methods["depositBatch"].ID()
}

func (d *DepositBatchRootERC721PredicateFn) EncodeAbi() ([]byte, error) {
	return RootERC721Predicate.Abi.Methods["depositBatch"].Encode(d)
}

func (d *DepositBatchRootERC721PredicateFn) DecodeAbi(buf []byte) error {
	return decodeMethod(RootERC721Predicate.Abi.Methods["depositBatch"], buf, d)
}

type InitializeChildMintableERC721PredicateFn struct {
	NewStateSender         types.Address `abi:"newStateSender"`
	NewExitHelper          types.Address `abi:"newExitHelper"`
	NewRootERC721Predicate types.Address `abi:"newRootERC721Predicate"`
	NewChildTokenTemplate  types.Address `abi:"newChildTokenTemplate"`
}

func (i *InitializeChildMintableERC721PredicateFn) Sig() []byte {
	return ChildMintableERC721Predicate.Abi.Methods["initialize"].ID()
}

func (i *InitializeChildMintableERC721PredicateFn) EncodeAbi() ([]byte, error) {
	return ChildMintableERC721Predicate.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeChildMintableERC721PredicateFn) DecodeAbi(buf []byte) error {
	return decodeMethod(ChildMintableERC721Predicate.Abi.Methods["initialize"], buf, i)
}

type SetApprovalForAllRootERC721Fn struct {
	Operator types.Address `abi:"operator"`
	Approved bool          `abi:"approved"`
}

func (s *SetApprovalForAllRootERC721Fn) Sig() []byte {
	return RootERC721.Abi.Methods["setApprovalForAll"].ID()
}

func (s *SetApprovalForAllRootERC721Fn) EncodeAbi() ([]byte, error) {
	return RootERC721.Abi.Methods["setApprovalForAll"].Encode(s)
}

func (s *SetApprovalForAllRootERC721Fn) DecodeAbi(buf []byte) error {
	return decodeMethod(RootERC721.Abi.Methods["setApprovalForAll"], buf, s)
}

type MintRootERC721Fn struct {
	To types.Address `abi:"to"`
}

func (m *MintRootERC721Fn) Sig() []byte {
	return RootERC721.Abi.Methods["mint"].ID()
}

func (m *MintRootERC721Fn) EncodeAbi() ([]byte, error) {
	return RootERC721.Abi.Methods["mint"].Encode(m)
}

func (m *MintRootERC721Fn) DecodeAbi(buf []byte) error {
	return decodeMethod(RootERC721.Abi.Methods["mint"], buf, m)
}

type InitializeChildERC721PredicateFn struct {
	NewL2StateSender       types.Address `abi:"newL2StateSender"`
	NewStateReceiver       types.Address `abi:"newStateReceiver"`
	NewRootERC721Predicate types.Address `abi:"newRootERC721Predicate"`
	NewChildTokenTemplate  types.Address `abi:"newChildTokenTemplate"`
}

func (i *InitializeChildERC721PredicateFn) Sig() []byte {
	return ChildERC721Predicate.Abi.Methods["initialize"].ID()
}

func (i *InitializeChildERC721PredicateFn) EncodeAbi() ([]byte, error) {
	return ChildERC721Predicate.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeChildERC721PredicateFn) DecodeAbi(buf []byte) error {
	return decodeMethod(ChildERC721Predicate.Abi.Methods["initialize"], buf, i)
}

type WithdrawBatchChildERC721PredicateFn struct {
	ChildToken types.Address   `abi:"childToken"`
	Receivers  []ethgo.Address `abi:"receivers"`
	TokenIDs   []*big.Int      `abi:"tokenIds"`
}

func (w *WithdrawBatchChildERC721PredicateFn) Sig() []byte {
	return ChildERC721Predicate.Abi.Methods["withdrawBatch"].ID()
}

func (w *WithdrawBatchChildERC721PredicateFn) EncodeAbi() ([]byte, error) {
	return ChildERC721Predicate.Abi.Methods["withdrawBatch"].Encode(w)
}

func (w *WithdrawBatchChildERC721PredicateFn) DecodeAbi(buf []byte) error {
	return decodeMethod(ChildERC721Predicate.Abi.Methods["withdrawBatch"], buf, w)
}

type InitializeChildERC721PredicateACLFn struct {
	NewL2StateSender       types.Address `abi:"newL2StateSender"`
	NewStateReceiver       types.Address `abi:"newStateReceiver"`
	NewRootERC721Predicate types.Address `abi:"newRootERC721Predicate"`
	NewChildTokenTemplate  types.Address `abi:"newChildTokenTemplate"`
	NewUseAllowList        bool          `abi:"newUseAllowList"`
	NewUseBlockList        bool          `abi:"newUseBlockList"`
	NewOwner               types.Address `abi:"newOwner"`
}

func (i *InitializeChildERC721PredicateACLFn) Sig() []byte {
	return ChildERC721PredicateACL.Abi.Methods["initialize"].ID()
}

func (i *InitializeChildERC721PredicateACLFn) EncodeAbi() ([]byte, error) {
	return ChildERC721PredicateACL.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeChildERC721PredicateACLFn) DecodeAbi(buf []byte) error {
	return decodeMethod(ChildERC721PredicateACL.Abi.Methods["initialize"], buf, i)
}

type WithdrawBatchChildERC721PredicateACLFn struct {
	ChildToken types.Address   `abi:"childToken"`
	Receivers  []ethgo.Address `abi:"receivers"`
	TokenIDs   []*big.Int      `abi:"tokenIds"`
}

func (w *WithdrawBatchChildERC721PredicateACLFn) Sig() []byte {
	return ChildERC721PredicateACL.Abi.Methods["withdrawBatch"].ID()
}

func (w *WithdrawBatchChildERC721PredicateACLFn) EncodeAbi() ([]byte, error) {
	return ChildERC721PredicateACL.Abi.Methods["withdrawBatch"].Encode(w)
}

func (w *WithdrawBatchChildERC721PredicateACLFn) DecodeAbi(buf []byte) error {
	return decodeMethod(ChildERC721PredicateACL.Abi.Methods["withdrawBatch"], buf, w)
}

type InitializeRootMintableERC721PredicateFn struct {
	NewL2StateSender        types.Address `abi:"newL2StateSender"`
	NewStateReceiver        types.Address `abi:"newStateReceiver"`
	NewChildERC721Predicate types.Address `abi:"newChildERC721Predicate"`
	NewChildTokenTemplate   types.Address `abi:"newChildTokenTemplate"`
}

func (i *InitializeRootMintableERC721PredicateFn) Sig() []byte {
	return RootMintableERC721Predicate.Abi.Methods["initialize"].ID()
}

func (i *InitializeRootMintableERC721PredicateFn) EncodeAbi() ([]byte, error) {
	return RootMintableERC721Predicate.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeRootMintableERC721PredicateFn) DecodeAbi(buf []byte) error {
	return decodeMethod(RootMintableERC721Predicate.Abi.Methods["initialize"], buf, i)
}

type InitializeRootMintableERC721PredicateACLFn struct {
	NewL2StateSender        types.Address `abi:"newL2StateSender"`
	NewStateReceiver        types.Address `abi:"newStateReceiver"`
	NewChildERC721Predicate types.Address `abi:"newChildERC721Predicate"`
	NewChildTokenTemplate   types.Address `abi:"newChildTokenTemplate"`
	NewUseAllowList         bool          `abi:"newUseAllowList"`
	NewUseBlockList         bool          `abi:"newUseBlockList"`
	NewOwner                types.Address `abi:"newOwner"`
}

func (i *InitializeRootMintableERC721PredicateACLFn) Sig() []byte {
	return RootMintableERC721PredicateACL.Abi.Methods["initialize"].ID()
}

func (i *InitializeRootMintableERC721PredicateACLFn) EncodeAbi() ([]byte, error) {
	return RootMintableERC721PredicateACL.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeRootMintableERC721PredicateACLFn) DecodeAbi(buf []byte) error {
	return decodeMethod(RootMintableERC721PredicateACL.Abi.Methods["initialize"], buf, i)
}

type InitializeChildERC721Fn struct {
	RootToken_ types.Address `abi:"rootToken_"`
	Name_      string        `abi:"name_"`
	Symbol_    string        `abi:"symbol_"`
}

func (i *InitializeChildERC721Fn) Sig() []byte {
	return ChildERC721.Abi.Methods["initialize"].ID()
}

func (i *InitializeChildERC721Fn) EncodeAbi() ([]byte, error) {
	return ChildERC721.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeChildERC721Fn) DecodeAbi(buf []byte) error {
	return decodeMethod(ChildERC721.Abi.Methods["initialize"], buf, i)
}

type OwnerOfChildERC721Fn struct {
	TokenID *big.Int `abi:"tokenId"`
}

func (o *OwnerOfChildERC721Fn) Sig() []byte {
	return ChildERC721.Abi.Methods["ownerOf"].ID()
}

func (o *OwnerOfChildERC721Fn) EncodeAbi() ([]byte, error) {
	return ChildERC721.Abi.Methods["ownerOf"].Encode(o)
}

func (o *OwnerOfChildERC721Fn) DecodeAbi(buf []byte) error {
	return decodeMethod(ChildERC721.Abi.Methods["ownerOf"], buf, o)
}

type InitializeCustomSupernetManagerFn struct {
	NewStakeManager      types.Address `abi:"newStakeManager"`
	NewBls               types.Address `abi:"newBls"`
	NewStateSender       types.Address `abi:"newStateSender"`
	NewMatic             types.Address `abi:"newMatic"`
	NewChildValidatorSet types.Address `abi:"newChildValidatorSet"`
	NewExitHelper        types.Address `abi:"newExitHelper"`
	NewDomain            string        `abi:"newDomain"`
}

func (i *InitializeCustomSupernetManagerFn) Sig() []byte {
	return CustomSupernetManager.Abi.Methods["initialize"].ID()
}

func (i *InitializeCustomSupernetManagerFn) EncodeAbi() ([]byte, error) {
	return CustomSupernetManager.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeCustomSupernetManagerFn) DecodeAbi(buf []byte) error {
	return decodeMethod(CustomSupernetManager.Abi.Methods["initialize"], buf, i)
}

type WhitelistValidatorsCustomSupernetManagerFn struct {
	Validators_ []ethgo.Address `abi:"validators_"`
}

func (w *WhitelistValidatorsCustomSupernetManagerFn) Sig() []byte {
	return CustomSupernetManager.Abi.Methods["whitelistValidators"].ID()
}

func (w *WhitelistValidatorsCustomSupernetManagerFn) EncodeAbi() ([]byte, error) {
	return CustomSupernetManager.Abi.Methods["whitelistValidators"].Encode(w)
}

func (w *WhitelistValidatorsCustomSupernetManagerFn) DecodeAbi(buf []byte) error {
	return decodeMethod(CustomSupernetManager.Abi.Methods["whitelistValidators"], buf, w)
}

type RegisterCustomSupernetManagerFn struct {
	Signature [2]*big.Int `abi:"signature"`
	Pubkey    [4]*big.Int `abi:"pubkey"`
}

func (r *RegisterCustomSupernetManagerFn) Sig() []byte {
	return CustomSupernetManager.Abi.Methods["register"].ID()
}

func (r *RegisterCustomSupernetManagerFn) EncodeAbi() ([]byte, error) {
	return CustomSupernetManager.Abi.Methods["register"].Encode(r)
}

func (r *RegisterCustomSupernetManagerFn) DecodeAbi(buf []byte) error {
	return decodeMethod(CustomSupernetManager.Abi.Methods["register"], buf, r)
}

type GetValidatorCustomSupernetManagerFn struct {
	Validator_ types.Address `abi:"validator_"`
}

func (g *GetValidatorCustomSupernetManagerFn) Sig() []byte {
	return CustomSupernetManager.Abi.Methods["getValidator"].ID()
}

func (g *GetValidatorCustomSupernetManagerFn) EncodeAbi() ([]byte, error) {
	return CustomSupernetManager.Abi.Methods["getValidator"].Encode(g)
}

func (g *GetValidatorCustomSupernetManagerFn) DecodeAbi(buf []byte) error {
	return decodeMethod(CustomSupernetManager.Abi.Methods["getValidator"], buf, g)
}

type ValidatorRegisteredEvent struct {
	Validator types.Address `abi:"validator"`
	BlsKey    [4]*big.Int   `abi:"blsKey"`
}

func (*ValidatorRegisteredEvent) Sig() ethgo.Hash {
	return CustomSupernetManager.Abi.Events["ValidatorRegistered"].ID()
}

func (*ValidatorRegisteredEvent) Encode(inputs interface{}) ([]byte, error) {
	return CustomSupernetManager.Abi.Events["ValidatorRegistered"].Inputs.Encode(inputs)
}

func (v *ValidatorRegisteredEvent) ParseLog(log *ethgo.Log) (bool, error) {
	if !CustomSupernetManager.Abi.Events["ValidatorRegistered"].Match(log) {
		return false, nil
	}

	return true, decodeEvent(CustomSupernetManager.Abi.Events["ValidatorRegistered"], log, v)
}

type InitializeStakeManagerFn struct {
	NewMatic types.Address `abi:"newMatic"`
}

func (i *InitializeStakeManagerFn) Sig() []byte {
	return StakeManager.Abi.Methods["initialize"].ID()
}

func (i *InitializeStakeManagerFn) EncodeAbi() ([]byte, error) {
	return StakeManager.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeStakeManagerFn) DecodeAbi(buf []byte) error {
	return decodeMethod(StakeManager.Abi.Methods["initialize"], buf, i)
}

type RegisterChildChainStakeManagerFn struct {
	Manager types.Address `abi:"manager"`
}

func (r *RegisterChildChainStakeManagerFn) Sig() []byte {
	return StakeManager.Abi.Methods["registerChildChain"].ID()
}

func (r *RegisterChildChainStakeManagerFn) EncodeAbi() ([]byte, error) {
	return StakeManager.Abi.Methods["registerChildChain"].Encode(r)
}

func (r *RegisterChildChainStakeManagerFn) DecodeAbi(buf []byte) error {
	return decodeMethod(StakeManager.Abi.Methods["registerChildChain"], buf, r)
}

type StakeForStakeManagerFn struct {
	ID     *big.Int `abi:"id"`
	Amount *big.Int `abi:"amount"`
}

func (s *StakeForStakeManagerFn) Sig() []byte {
	return StakeManager.Abi.Methods["stakeFor"].ID()
}

func (s *StakeForStakeManagerFn) EncodeAbi() ([]byte, error) {
	return StakeManager.Abi.Methods["stakeFor"].Encode(s)
}

func (s *StakeForStakeManagerFn) DecodeAbi(buf []byte) error {
	return decodeMethod(StakeManager.Abi.Methods["stakeFor"], buf, s)
}

type ReleaseStakeOfStakeManagerFn struct {
	Validator types.Address `abi:"validator"`
	Amount    *big.Int      `abi:"amount"`
}

func (r *ReleaseStakeOfStakeManagerFn) Sig() []byte {
	return StakeManager.Abi.Methods["releaseStakeOf"].ID()
}

func (r *ReleaseStakeOfStakeManagerFn) EncodeAbi() ([]byte, error) {
	return StakeManager.Abi.Methods["releaseStakeOf"].Encode(r)
}

func (r *ReleaseStakeOfStakeManagerFn) DecodeAbi(buf []byte) error {
	return decodeMethod(StakeManager.Abi.Methods["releaseStakeOf"], buf, r)
}

type WithdrawStakeStakeManagerFn struct {
	To     types.Address `abi:"to"`
	Amount *big.Int      `abi:"amount"`
}

func (w *WithdrawStakeStakeManagerFn) Sig() []byte {
	return StakeManager.Abi.Methods["withdrawStake"].ID()
}

func (w *WithdrawStakeStakeManagerFn) EncodeAbi() ([]byte, error) {
	return StakeManager.Abi.Methods["withdrawStake"].Encode(w)
}

func (w *WithdrawStakeStakeManagerFn) DecodeAbi(buf []byte) error {
	return decodeMethod(StakeManager.Abi.Methods["withdrawStake"], buf, w)
}

type StakeOfStakeManagerFn struct {
	Validator types.Address `abi:"validator"`
	ID        *big.Int      `abi:"id"`
}

func (s *StakeOfStakeManagerFn) Sig() []byte {
	return StakeManager.Abi.Methods["stakeOf"].ID()
}

func (s *StakeOfStakeManagerFn) EncodeAbi() ([]byte, error) {
	return StakeManager.Abi.Methods["stakeOf"].Encode(s)
}

func (s *StakeOfStakeManagerFn) DecodeAbi(buf []byte) error {
	return decodeMethod(StakeManager.Abi.Methods["stakeOf"], buf, s)
}

type ChildManagerRegisteredEvent struct {
	ID      *big.Int      `abi:"id"`
	Manager types.Address `abi:"manager"`
}

func (*ChildManagerRegisteredEvent) Sig() ethgo.Hash {
	return StakeManager.Abi.Events["ChildManagerRegistered"].ID()
}

func (*ChildManagerRegisteredEvent) Encode(inputs interface{}) ([]byte, error) {
	return StakeManager.Abi.Events["ChildManagerRegistered"].Inputs.Encode(inputs)
}

func (c *ChildManagerRegisteredEvent) ParseLog(log *ethgo.Log) (bool, error) {
	if !StakeManager.Abi.Events["ChildManagerRegistered"].Match(log) {
		return false, nil
	}

	return true, decodeEvent(StakeManager.Abi.Events["ChildManagerRegistered"], log, c)
}

type StakeAddedEvent struct {
	ID        *big.Int      `abi:"id"`
	Validator types.Address `abi:"validator"`
	Amount    *big.Int      `abi:"amount"`
}

func (*StakeAddedEvent) Sig() ethgo.Hash {
	return StakeManager.Abi.Events["StakeAdded"].ID()
}

func (*StakeAddedEvent) Encode(inputs interface{}) ([]byte, error) {
	return StakeManager.Abi.Events["StakeAdded"].Inputs.Encode(inputs)
}

func (s *StakeAddedEvent) ParseLog(log *ethgo.Log) (bool, error) {
	if !StakeManager.Abi.Events["StakeAdded"].Match(log) {
		return false, nil
	}

	return true, decodeEvent(StakeManager.Abi.Events["StakeAdded"], log, s)
}

type StakeWithdrawnEvent struct {
	Validator types.Address `abi:"validator"`
	Recipient types.Address `abi:"recipient"`
	Amount    *big.Int      `abi:"amount"`
}

func (*StakeWithdrawnEvent) Sig() ethgo.Hash {
	return StakeManager.Abi.Events["StakeWithdrawn"].ID()
}

func (*StakeWithdrawnEvent) Encode(inputs interface{}) ([]byte, error) {
	return StakeManager.Abi.Events["StakeWithdrawn"].Inputs.Encode(inputs)
}

func (s *StakeWithdrawnEvent) ParseLog(log *ethgo.Log) (bool, error) {
	if !StakeManager.Abi.Events["StakeWithdrawn"].Match(log) {
		return false, nil
	}

	return true, decodeEvent(StakeManager.Abi.Events["StakeWithdrawn"], log, s)
}

type UnstakeValidatorSetFn struct {
	Amount *big.Int `abi:"amount"`
}

func (u *UnstakeValidatorSetFn) Sig() []byte {
	return ValidatorSet.Abi.Methods["unstake"].ID()
}

func (u *UnstakeValidatorSetFn) EncodeAbi() ([]byte, error) {
	return ValidatorSet.Abi.Methods["unstake"].Encode(u)
}

func (u *UnstakeValidatorSetFn) DecodeAbi(buf []byte) error {
	return decodeMethod(ValidatorSet.Abi.Methods["unstake"], buf, u)
}

type WithdrawalRegisteredEvent struct {
	Account types.Address `abi:"account"`
	Amount  *big.Int      `abi:"amount"`
}

func (*WithdrawalRegisteredEvent) Sig() ethgo.Hash {
	return ValidatorSet.Abi.Events["WithdrawalRegistered"].ID()
}

func (*WithdrawalRegisteredEvent) Encode(inputs interface{}) ([]byte, error) {
	return ValidatorSet.Abi.Events["WithdrawalRegistered"].Inputs.Encode(inputs)
}

func (w *WithdrawalRegisteredEvent) ParseLog(log *ethgo.Log) (bool, error) {
	if !ValidatorSet.Abi.Events["WithdrawalRegistered"].Match(log) {
		return false, nil
	}

	return true, decodeEvent(ValidatorSet.Abi.Events["WithdrawalRegistered"], log, w)
}

type InitializeRewardPoolFn struct {
	NewRewardToken  types.Address `abi:"newRewardToken"`
	NewRewardWallet types.Address `abi:"newRewardWallet"`
	NewValidatorSet types.Address `abi:"newValidatorSet"`
	NewBaseReward   *big.Int      `abi:"newBaseReward"`
}

func (i *InitializeRewardPoolFn) Sig() []byte {
	return RewardPool.Abi.Methods["initialize"].ID()
}

func (i *InitializeRewardPoolFn) EncodeAbi() ([]byte, error) {
	return RewardPool.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeRewardPoolFn) DecodeAbi(buf []byte) error {
	return decodeMethod(RewardPool.Abi.Methods["initialize"], buf, i)
}

type DistributeRewardForRewardPoolFn struct {
	EpochID *big.Int  `abi:"epochId"`
	Uptime  []*Uptime `abi:"uptime"`
}

func (d *DistributeRewardForRewardPoolFn) Sig() []byte {
	return RewardPool.Abi.Methods["distributeRewardFor"].ID()
}

func (d *DistributeRewardForRewardPoolFn) EncodeAbi() ([]byte, error) {
	return RewardPool.Abi.Methods["distributeRewardFor"].Encode(d)
}

func (d *DistributeRewardForRewardPoolFn) DecodeAbi(buf []byte) error {
	return decodeMethod(RewardPool.Abi.Methods["distributeRewardFor"], buf, d)
}

type InitializeEIP1559BurnFn struct {
	NewChildERC20Predicate types.Address `abi:"newChildERC20Predicate"`
	NewBurnDestination     types.Address `abi:"newBurnDestination"`
}

func (i *InitializeEIP1559BurnFn) Sig() []byte {
	return EIP1559Burn.Abi.Methods["initialize"].ID()
}

func (i *InitializeEIP1559BurnFn) EncodeAbi() ([]byte, error) {
	return EIP1559Burn.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeEIP1559BurnFn) DecodeAbi(buf []byte) error {
	return decodeMethod(EIP1559Burn.Abi.Methods["initialize"], buf, i)
}

type L2StateSyncedEvent struct {
	ID       *big.Int      `abi:"id"`
	Sender   types.Address `abi:"sender"`
	Receiver types.Address `abi:"receiver"`
	Data     []byte        `abi:"data"`
}

func (*L2StateSyncedEvent) Sig() ethgo.Hash {
	return L2StateSender.Abi.Events["L2StateSynced"].ID()
}

func (*L2StateSyncedEvent) Encode(inputs interface{}) ([]byte, error) {
	return L2StateSender.Abi.Events["L2StateSynced"].Inputs.Encode(inputs)
}

func (l *L2StateSyncedEvent) ParseLog(log *ethgo.Log) (bool, error) {
	if !L2StateSender.Abi.Events["L2StateSynced"].Match(log) {
		return false, nil
	}

	return true, decodeEvent(L2StateSender.Abi.Events["L2StateSynced"], log, l)
}

type StateSyncCommitment struct {
	StartID *big.Int   `abi:"startId"`
	EndID   *big.Int   `abi:"endId"`
	Root    types.Hash `abi:"root"`
}

var StateSyncCommitmentABIType = abi.MustNewType("tuple(uint256 startId,uint256 endId,bytes32 root)")

func (s *StateSyncCommitment) EncodeAbi() ([]byte, error) {
	return StateSyncCommitmentABIType.Encode(s)
}

func (s *StateSyncCommitment) DecodeAbi(buf []byte) error {
	return decodeStruct(StateSyncCommitmentABIType, buf, &s)
}

type CommitStateReceiverFn struct {
	Commitment *StateSyncCommitment `abi:"commitment"`
	Signature  []byte               `abi:"signature"`
	Bitmap     []byte               `abi:"bitmap"`
}

func (c *CommitStateReceiverFn) Sig() []byte {
	return StateReceiver.Abi.Methods["commit"].ID()
}

func (c *CommitStateReceiverFn) EncodeAbi() ([]byte, error) {
	return StateReceiver.Abi.Methods["commit"].Encode(c)
}

func (c *CommitStateReceiverFn) DecodeAbi(buf []byte) error {
	return decodeMethod(StateReceiver.Abi.Methods["commit"], buf, c)
}

type NewCommitmentEvent struct {
	StartID *big.Int   `abi:"startId"`
	EndID   *big.Int   `abi:"endId"`
	Root    types.Hash `abi:"root"`
}

func (*NewCommitmentEvent) Sig() ethgo.Hash {
	return StateReceiver.Abi.Events["NewCommitment"].ID()
}

func (*NewCommitmentEvent) Encode(inputs interface{}) ([]byte, error) {
	return StateReceiver.Abi.Events["NewCommitment"].Inputs.Encode(inputs)
}

func (n *NewCommitmentEvent) ParseLog(log *ethgo.Log) (bool, error) {
	if !StateReceiver.Abi.Events["NewCommitment"].Match(log) {
		return false, nil
	}

	return true, decodeEvent(StateReceiver.Abi.Events["NewCommitment"], log, n)
}

type StateSync struct {
	ID       *big.Int      `abi:"id"`
	Sender   types.Address `abi:"sender"`
	Receiver types.Address `abi:"receiver"`
	Data     []byte        `abi:"data"`
}

var StateSyncABIType = abi.MustNewType("tuple(uint256 id,address sender,address receiver,bytes data)")

func (s *StateSync) EncodeAbi() ([]byte, error) {
	return StateSyncABIType.Encode(s)
}

func (s *StateSync) DecodeAbi(buf []byte) error {
	return decodeStruct(StateSyncABIType, buf, &s)
}

type ExecuteStateReceiverFn struct {
	Proof []types.Hash `abi:"proof"`
	Obj   *StateSync   `abi:"obj"`
}

func (e *ExecuteStateReceiverFn) Sig() []byte {
	return StateReceiver.Abi.Methods["execute"].ID()
}

func (e *ExecuteStateReceiverFn) EncodeAbi() ([]byte, error) {
	return StateReceiver.Abi.Methods["execute"].Encode(e)
}

func (e *ExecuteStateReceiverFn) DecodeAbi(buf []byte) error {
	return decodeMethod(StateReceiver.Abi.Methods["execute"], buf, e)
}

type StateSyncResultEvent struct {
	Counter *big.Int `abi:"counter"`
	Status  bool     `abi:"status"`
	Message []byte   `abi:"message"`
}

func (*StateSyncResultEvent) Sig() ethgo.Hash {
	return StateReceiver.Abi.Events["StateSyncResult"].ID()
}

func (*StateSyncResultEvent) Encode(inputs interface{}) ([]byte, error) {
	return StateReceiver.Abi.Events["StateSyncResult"].Inputs.Encode(inputs)
}

func (s *StateSyncResultEvent) ParseLog(log *ethgo.Log) (bool, error) {
	if !StateReceiver.Abi.Events["StateSyncResult"].Match(log) {
		return false, nil
	}

	return true, decodeEvent(StateReceiver.Abi.Events["StateSyncResult"], log, s)
}

type SyncStateStateSenderFn struct {
	Receiver types.Address `abi:"receiver"`
	Data     []byte        `abi:"data"`
}

func (s *SyncStateStateSenderFn) Sig() []byte {
	return StateSender.Abi.Methods["syncState"].ID()
}

func (s *SyncStateStateSenderFn) EncodeAbi() ([]byte, error) {
	return StateSender.Abi.Methods["syncState"].Encode(s)
}

func (s *SyncStateStateSenderFn) DecodeAbi(buf []byte) error {
	return decodeMethod(StateSender.Abi.Methods["syncState"], buf, s)
}

type ExitExitHelperFn struct {
	BlockNumber  *big.Int     `abi:"blockNumber"`
	LeafIndex    *big.Int     `abi:"leafIndex"`
	UnhashedLeaf []byte       `abi:"unhashedLeaf"`
	Proof        []types.Hash `abi:"proof"`
}

func (e *ExitExitHelperFn) Sig() []byte {
	return ExitHelper.Abi.Methods["exit"].ID()
}

func (e *ExitExitHelperFn) EncodeAbi() ([]byte, error) {
	return ExitHelper.Abi.Methods["exit"].Encode(e)
}

func (e *ExitExitHelperFn) DecodeAbi(buf []byte) error {
	return decodeMethod(ExitHelper.Abi.Methods["exit"], buf, e)
}

type InitializeExitHelperFn struct {
	NewCheckpointManager types.Address `abi:"newCheckpointManager"`
}

func (i *InitializeExitHelperFn) Sig() []byte {
	return ExitHelper.Abi.Methods["initialize"].ID()
}

func (i *InitializeExitHelperFn) EncodeAbi() ([]byte, error) {
	return ExitHelper.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeExitHelperFn) DecodeAbi(buf []byte) error {
	return decodeMethod(ExitHelper.Abi.Methods["initialize"], buf, i)
}
