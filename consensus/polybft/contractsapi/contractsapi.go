// Code generated by scapi/gen. DO NOT EDIT.
package contractsapi

import (
	"math/big"

	"github.com/0xPolygon/polygon-edge/types"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/abi"
)

type Epoch struct {
	StartBlock *big.Int   `abi:"startBlock"`
	EndBlock   *big.Int   `abi:"endBlock"`
	EpochRoot  types.Hash `abi:"epochRoot"`
}

var EpochABIType = abi.MustNewType("tuple(uint256 startBlock,uint256 endBlock,bytes32 epochRoot)")

func (e *Epoch) EncodeAbi() ([]byte, error) {
	return EpochABIType.Encode(e)
}

func (e *Epoch) DecodeAbi(buf []byte) error {
	return decodeStruct(EpochABIType, buf, &e)
}

type UptimeData struct {
	Validator    types.Address `abi:"validator"`
	SignedBlocks *big.Int      `abi:"signedBlocks"`
}

var UptimeDataABIType = abi.MustNewType("tuple(address validator,uint256 signedBlocks)")

func (u *UptimeData) EncodeAbi() ([]byte, error) {
	return UptimeDataABIType.Encode(u)
}

func (u *UptimeData) DecodeAbi(buf []byte) error {
	return decodeStruct(UptimeDataABIType, buf, &u)
}

type Uptime struct {
	EpochID     *big.Int      `abi:"epochId"`
	UptimeData  []*UptimeData `abi:"uptimeData"`
	TotalBlocks *big.Int      `abi:"totalBlocks"`
}

var UptimeABIType = abi.MustNewType("tuple(uint256 epochId,tuple(address validator,uint256 signedBlocks)[] uptimeData,uint256 totalBlocks)")

func (u *Uptime) EncodeAbi() ([]byte, error) {
	return UptimeABIType.Encode(u)
}

func (u *Uptime) DecodeAbi(buf []byte) error {
	return decodeStruct(UptimeABIType, buf, &u)
}

type CommitEpochChildValidatorSetFn struct {
	ID     *big.Int `abi:"id"`
	Epoch  *Epoch   `abi:"epoch"`
	Uptime *Uptime  `abi:"uptime"`
}

func (c *CommitEpochChildValidatorSetFn) Sig() []byte {
	return ChildValidatorSet.Abi.Methods["commitEpoch"].ID()
}

func (c *CommitEpochChildValidatorSetFn) EncodeAbi() ([]byte, error) {
	return ChildValidatorSet.Abi.Methods["commitEpoch"].Encode(c)
}

func (c *CommitEpochChildValidatorSetFn) DecodeAbi(buf []byte) error {
	return decodeMethod(ChildValidatorSet.Abi.Methods["commitEpoch"], buf, c)
}

type InitStruct struct {
	EpochReward   *big.Int `abi:"epochReward"`
	MinStake      *big.Int `abi:"minStake"`
	MinDelegation *big.Int `abi:"minDelegation"`
	EpochSize     *big.Int `abi:"epochSize"`
}

var InitStructABIType = abi.MustNewType("tuple(uint256 epochReward,uint256 minStake,uint256 minDelegation,uint256 epochSize)")

func (i *InitStruct) EncodeAbi() ([]byte, error) {
	return InitStructABIType.Encode(i)
}

func (i *InitStruct) DecodeAbi(buf []byte) error {
	return decodeStruct(InitStructABIType, buf, &i)
}

type ValidatorInit struct {
	Addr      types.Address `abi:"addr"`
	Pubkey    [4]*big.Int   `abi:"pubkey"`
	Signature [2]*big.Int   `abi:"signature"`
	Stake     *big.Int      `abi:"stake"`
}

var ValidatorInitABIType = abi.MustNewType("tuple(address addr,uint256[4] pubkey,uint256[2] signature,uint256 stake)")

func (v *ValidatorInit) EncodeAbi() ([]byte, error) {
	return ValidatorInitABIType.Encode(v)
}

func (v *ValidatorInit) DecodeAbi(buf []byte) error {
	return decodeStruct(ValidatorInitABIType, buf, &v)
}

type InitializeChildValidatorSetFn struct {
	Init       *InitStruct      `abi:"init"`
	Validators []*ValidatorInit `abi:"validators"`
	NewBls     types.Address    `abi:"newBls"`
	Governance types.Address    `abi:"governance"`
}

func (i *InitializeChildValidatorSetFn) Sig() []byte {
	return ChildValidatorSet.Abi.Methods["initialize"].ID()
}

func (i *InitializeChildValidatorSetFn) EncodeAbi() ([]byte, error) {
	return ChildValidatorSet.Abi.Methods["initialize"].Encode(i)
}

func (i *InitializeChildValidatorSetFn) DecodeAbi(buf []byte) error {
	return decodeMethod(ChildValidatorSet.Abi.Methods["initialize"], buf, i)
}

type AddToWhitelistChildValidatorSetFn struct {
	WhitelistAddreses []ethgo.Address `abi:"whitelistAddreses"`
}

func (a *AddToWhitelistChildValidatorSetFn) Sig() []byte {
	return ChildValidatorSet.Abi.Methods["addToWhitelist"].ID()
}

func (a *AddToWhitelistChildValidatorSetFn) EncodeAbi() ([]byte, error) {
	return ChildValidatorSet.Abi.Methods["addToWhitelist"].Encode(a)
}

func (a *AddToWhitelistChildValidatorSetFn) DecodeAbi(buf []byte) error {
	return decodeMethod(ChildValidatorSet.Abi.Methods["addToWhitelist"], buf, a)
}

type RegisterChildValidatorSetFn struct {
	Signature [2]*big.Int `abi:"signature"`
	Pubkey    [4]*big.Int `abi:"pubkey"`
}

func (r *RegisterChildValidatorSetFn) Sig() []byte {
	return ChildValidatorSet.Abi.Methods["register"].ID()
}

func (r *RegisterChildValidatorSetFn) EncodeAbi() ([]byte, error) {
	return ChildValidatorSet.Abi.Methods["register"].Encode(r)
}

func (r *RegisterChildValidatorSetFn) DecodeAbi(buf []byte) error {
	return decodeMethod(ChildValidatorSet.Abi.Methods["register"], buf, r)
}

type NewValidatorEvent struct {
	Validator types.Address `abi:"validator"`
	BlsKey    [4]*big.Int   `abi:"blsKey"`
}

func (*NewValidatorEvent) Sig() ethgo.Hash {
	return ChildValidatorSet.Abi.Events["NewValidator"].ID()
}

func (*NewValidatorEvent) Encode(inputs interface{}) ([]byte, error) {
	return ChildValidatorSet.Abi.Events["NewValidator"].Inputs.Encode(inputs)
}

func (n *NewValidatorEvent) ParseLog(log *ethgo.Log) (bool, error) {
	if !ChildValidatorSet.Abi.Events["NewValidator"].Match(log) {
		return false, nil
	}

	return true, decodeEvent(ChildValidatorSet.Abi.Events["NewValidator"], log, n)
}

type StakedEvent struct {
	Validator types.Address `abi:"validator"`
	Amount    *big.Int      `abi:"amount"`
}

func (*StakedEvent) Sig() ethgo.Hash {
	return ChildValidatorSet.Abi.Events["Staked"].ID()
}

func (*StakedEvent) Encode(inputs interface{}) ([]byte, error) {
	return ChildValidatorSet.Abi.Events["Staked"].Inputs.Encode(inputs)
}

func (s *StakedEvent) ParseLog(log *ethgo.Log) (bool, error) {
	if !ChildValidatorSet.Abi.Events["Staked"].Match(log) {
		return false, nil
	}

	return true, decodeEvent(ChildValidatorSet.Abi.Events["Staked"], log, s)
}

type DelegatedEvent struct {
	Delegator types.Address `abi:"delegator"`
	Validator types.Address `abi:"validator"`
	Amount    *big.Int      `abi:"amount"`
}

func (*DelegatedEvent) Sig() ethgo.Hash {
	return ChildValidatorSet.Abi.Events["Delegated"].ID()
}

func (*DelegatedEvent) Encode(inputs interface{}) ([]byte, error) {
	return ChildValidatorSet.Abi.Events["Delegated"].Inputs.Encode(inputs)
}

func (d *DelegatedEvent) ParseLog(log *ethgo.Log) (bool, error) {
	if !ChildValidatorSet.Abi.Events["Delegated"].Match(log) {
		return false, nil
	}

	return true, decodeEvent(ChildValidatorSet.Abi.Events["Delegated"], log, d)
}

type UnstakedEvent struct {
	Validator types.Address `abi:"validator"`
	Amount    *big.Int      `abi:"amount"`
}

func (*UnstakedEvent) Sig() ethgo.Hash {
	return ChildValidatorSet.Abi.Events["Unstaked"].ID()
}

func (*UnstakedEvent) Encode(inputs interface{}) ([]byte, error) {
	return ChildValidatorSet.Abi.Events["Unstaked"].Inputs.Encode(inputs)
}

func (u *UnstakedEvent) ParseLog(log *ethgo.Log) (bool, error) {
	if !ChildValidatorSet.Abi.Events["Unstaked"].Match(log) {
		return false, nil
	}

	return true, decodeEvent(ChildValidatorSet.Abi.Events["Unstaked"], log, u)
}

type UndelegatedEvent struct {
	Delegator types.Address `abi:"delegator"`
	Validator types.Address `abi:"validator"`
	Amount    *big.Int      `abi:"amount"`
}

func (*UndelegatedEvent) Sig() ethgo.Hash {
	return ChildValidatorSet.Abi.Events["Undelegated"].ID()
}

func (*UndelegatedEvent) Encode(inputs interface{}) ([]byte, error) {
	return ChildValidatorSet.Abi.Events["Undelegated"].Inputs.Encode(inputs)
}

func (u *UndelegatedEvent) ParseLog(log *ethgo.Log) (bool, error) {
	if !ChildValidatorSet.Abi.Events["Undelegated"].Match(log) {
		return false, nil
	}

	return true, decodeEvent(ChildValidatorSet.Abi.Events["Undelegated"], log, u)
}

type AddedToWhitelistEvent struct {
	Validator types.Address `abi:"validator"`
}

func (*AddedToWhitelistEvent) Sig() ethgo.Hash {
	return ChildValidatorSet.Abi.Events["AddedToWhitelist"].ID()
}

func (*AddedToWhitelistEvent) Encode(inputs interface{}) ([]byte, error) {
	return ChildValidatorSet.Abi.Events["AddedToWhitelist"].Inputs.Encode(inputs)
}

func (a *AddedToWhitelistEvent) ParseLog(log *ethgo.Log) (bool, error) {
	if !ChildValidatorSet.Abi.Events["AddedToWhitelist"].Match(log) {
		return false, nil
	}

	return true, decodeEvent(ChildValidatorSet.Abi.Events["AddedToWhitelist"], log, a)
}

type WithdrawalEvent struct {
	Account types.Address `abi:"account"`
	To      types.Address `abi:"to"`
	Amount  *big.Int      `abi:"amount"`
}

func (*WithdrawalEvent) Sig() ethgo.Hash {
	return ChildValidatorSet.Abi.Events["Withdrawal"].ID()
}

func (*WithdrawalEvent) Encode(inputs interface{}) ([]byte, error) {
	return ChildValidatorSet.Abi.Events["Withdrawal"].Inputs.Encode(inputs)
}

func (w *WithdrawalEvent) ParseLog(log *ethgo.Log) (bool, error) {
	if !ChildValidatorSet.Abi.Events["Withdrawal"].Match(log) {
		return false, nil
	}

	return true, decodeEvent(ChildValidatorSet.Abi.Events["Withdrawal"], log, w)
}

type TransferEvent struct {
	From  types.Address `abi:"from"`
	To    types.Address `abi:"to"`
	Value *big.Int      `abi:"value"`
}

func (*TransferEvent) Sig() ethgo.Hash {
	return ChildValidatorSet.Abi.Events["Transfer"].ID()
}

func (*TransferEvent) Encode(inputs interface{}) ([]byte, error) {
	return ChildValidatorSet.Abi.Events["Transfer"].Inputs.Encode(inputs)
}

func (t *TransferEvent) ParseLog(log *ethgo.Log) (bool, error) {
	if !ChildValidatorSet.Abi.Events["Transfer"].Match(log) {
		return false, nil
	}

	return true, decodeEvent(ChildValidatorSet.Abi.Events["Transfer"], log, t)
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

type StateSyncedEvent struct {
	ID       *big.Int      `abi:"id"`
	Sender   types.Address `abi:"sender"`
	Receiver types.Address `abi:"receiver"`
	Data     []byte        `abi:"data"`
}

func (*StateSyncedEvent) Sig() ethgo.Hash {
	return StateSender.Abi.Events["StateSynced"].ID()
}

func (*StateSyncedEvent) Encode(inputs interface{}) ([]byte, error) {
	return StateSender.Abi.Events["StateSynced"].Inputs.Encode(inputs)
}

func (s *StateSyncedEvent) ParseLog(log *ethgo.Log) (bool, error) {
	if !StateSender.Abi.Events["StateSynced"].Match(log) {
		return false, nil
	}

	return true, decodeEvent(StateSender.Abi.Events["StateSynced"], log, s)
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
