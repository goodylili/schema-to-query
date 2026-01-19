package sui

type AddressTransactionRelationship string

const AddressTransactionRelationship_SENT AddressTransactionRelationship = "SENT"
const AddressTransactionRelationship_AFFECTED AddressTransactionRelationship = "AFFECTED"

type ConsensusObjectCancellationReason string

const ConsensusObjectCancellationReason_CANCELLED_READ ConsensusObjectCancellationReason = "CANCELLED_READ"
const ConsensusObjectCancellationReason_CONGESTED ConsensusObjectCancellationReason = "CONGESTED"
const ConsensusObjectCancellationReason_RANDOMNESS_UNAVAILABLE ConsensusObjectCancellationReason = "RANDOMNESS_UNAVAILABLE"
const ConsensusObjectCancellationReason_UNKNOWN ConsensusObjectCancellationReason = "UNKNOWN"

type ExecutionStatus string

const ExecutionStatus_SUCCESS ExecutionStatus = "SUCCESS"
const ExecutionStatus_FAILURE ExecutionStatus = "FAILURE"

type MoveAbility string

const MoveAbility_COPY MoveAbility = "COPY"
const MoveAbility_DROP MoveAbility = "DROP"
const MoveAbility_KEY MoveAbility = "KEY"
const MoveAbility_STORE MoveAbility = "STORE"

type MoveVisibility string

const MoveVisibility_PUBLIC MoveVisibility = "PUBLIC"
const MoveVisibility_PRIVATE MoveVisibility = "PRIVATE"
const MoveVisibility_FRIEND MoveVisibility = "FRIEND"

type OwnerKind string

const OwnerKind_ADDRESS OwnerKind = "ADDRESS"
const OwnerKind_OBJECT OwnerKind = "OBJECT"
const OwnerKind_SHARED OwnerKind = "SHARED"
const OwnerKind_IMMUTABLE OwnerKind = "IMMUTABLE"

type RegulatedState string

const RegulatedState_REGULATED RegulatedState = "REGULATED"
const RegulatedState_UNREGULATED RegulatedState = "UNREGULATED"

type SupplyState string

const SupplyState_BURN_ONLY SupplyState = "BURN_ONLY"
const SupplyState_FIXED SupplyState = "FIXED"

type TransactionKindInput string

const TransactionKindInput_SYSTEM_TX TransactionKindInput = "SYSTEM_TX"
const TransactionKindInput_PROGRAMMABLE_TX TransactionKindInput = "PROGRAMMABLE_TX"

type ZkLoginIntentScope string

const ZkLoginIntentScope_TRANSACTION_DATA ZkLoginIntentScope = "TRANSACTION_DATA"
const ZkLoginIntentScope_PERSONAL_MESSAGE ZkLoginIntentScope = "PERSONAL_MESSAGE"

type AccumulatorRootCreateTransaction struct {
	_ bool `json:"_"`
}

type ActiveJwk struct {
	Alg string `json:"alg"`
	E string `json:"e"`
	Epoch *Epoch `json:"epoch"`
	Iss string `json:"iss"`
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	N string `json:"n"`
}

type ActiveJwkConnection struct {
	Edges []*ActiveJwkEdge `json:"edges"`
	Nodes []*ActiveJwk `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type ActiveJwkEdge struct {
	Cursor string `json:"cursor"`
	Node *ActiveJwk `json:"node"`
}

type Address struct {
	Address string `json:"address"`
	AsObject *Object `json:"asObject"`
	Balance *Balance `json:"balance"`
	Balances *BalanceConnection `json:"balances"`
	DefaultSuinsName string `json:"defaultSuinsName"`
	DynamicField *DynamicField `json:"dynamicField"`
	DynamicFields *DynamicFieldConnection `json:"dynamicFields"`
	DynamicObjectField *DynamicField `json:"dynamicObjectField"`
	MultiGetBalances []*Balance `json:"multiGetBalances"`
	MultiGetDynamicFields []*DynamicField `json:"multiGetDynamicFields"`
	MultiGetDynamicObjectFields []*DynamicField `json:"multiGetDynamicObjectFields"`
	Objects *MoveObjectConnection `json:"objects"`
	Transactions *TransactionConnection `json:"transactions"`
}

type AddressAliasStateCreateTransaction struct {
	_ bool `json:"_"`
}

type AddressOwner struct {
	Address *Address `json:"address"`
}

type AuthenticatorStateCreateTransaction struct {
	_ bool `json:"_"`
}

type AuthenticatorStateExpireTransaction struct {
	AuthenticatorObjInitialSharedVersion string `json:"authenticatorObjInitialSharedVersion"`
	MinEpoch *Epoch `json:"minEpoch"`
}

type AuthenticatorStateUpdateTransaction struct {
	AuthenticatorObjInitialSharedVersion string `json:"authenticatorObjInitialSharedVersion"`
	Epoch *Epoch `json:"epoch"`
	NewActiveJwks *ActiveJwkConnection `json:"newActiveJwks"`
	Round string `json:"round"`
}

type AvailableRange struct {
	First *Checkpoint `json:"first"`
	Last *Checkpoint `json:"last"`
}

type Balance struct {
	CoinType *MoveType `json:"coinType"`
	TotalBalance string `json:"totalBalance"`
}

type BalanceChange struct {
	Amount string `json:"amount"`
	CoinType *MoveType `json:"coinType"`
	Owner *Address `json:"owner"`
}

type BalanceChangeConnection struct {
	Edges []*BalanceChangeEdge `json:"edges"`
	Nodes []*BalanceChange `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type BalanceChangeEdge struct {
	Cursor string `json:"cursor"`
	Node *BalanceChange `json:"node"`
}

type BalanceConnection struct {
	Edges []*BalanceEdge `json:"edges"`
	Nodes []*Balance `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type BalanceEdge struct {
	Cursor string `json:"cursor"`
	Node *Balance `json:"node"`
}

type BridgeCommitteeInitTransaction struct {
	BridgeObjectVersion string `json:"bridgeObjectVersion"`
}

type BridgeStateCreateTransaction struct {
	ChainIdentifier string `json:"chainIdentifier"`
}

type ChangeEpochTransaction struct {
	ComputationCharge string `json:"computationCharge"`
	Epoch *Epoch `json:"epoch"`
	EpochStartTimestamp string `json:"epochStartTimestamp"`
	NonRefundableStorageFee string `json:"nonRefundableStorageFee"`
	ProtocolConfigs *ProtocolConfigs `json:"protocolConfigs"`
	StorageCharge string `json:"storageCharge"`
	StorageRebate string `json:"storageRebate"`
	SystemPackages *MovePackageConnection `json:"systemPackages"`
}

type Checkpoint struct {
	ArtifactsDigest string `json:"artifactsDigest"`
	ContentBcs string `json:"contentBcs"`
	ContentDigest string `json:"contentDigest"`
	Digest string `json:"digest"`
	Epoch *Epoch `json:"epoch"`
	NetworkTotalTransactions string `json:"networkTotalTransactions"`
	PreviousCheckpointDigest string `json:"previousCheckpointDigest"`
	Query *Query `json:"query"`
	RollingGasSummary *GasCostSummary `json:"rollingGasSummary"`
	SequenceNumber string `json:"sequenceNumber"`
	SummaryBcs string `json:"summaryBcs"`
	Timestamp string `json:"timestamp"`
	Transactions *TransactionConnection `json:"transactions"`
	ValidatorSignatures *ValidatorAggregatedSignature `json:"validatorSignatures"`
}

type CheckpointConnection struct {
	Edges []*CheckpointEdge `json:"edges"`
	Nodes []*Checkpoint `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type CheckpointEdge struct {
	Cursor string `json:"cursor"`
	Node *Checkpoint `json:"node"`
}

type CoinDenyListStateCreateTransaction struct {
	_ bool `json:"_"`
}

type CoinMetadata struct {
	Address string `json:"address"`
	AllowGlobalPause bool `json:"allowGlobalPause"`
	Balance *Balance `json:"balance"`
	Balances *BalanceConnection `json:"balances"`
	Contents *MoveValue `json:"contents"`
	Decimals int `json:"decimals"`
	DefaultSuinsName string `json:"defaultSuinsName"`
	DenyCap *MoveObject `json:"denyCap"`
	Description string `json:"description"`
	Digest string `json:"digest"`
	DynamicField *DynamicField `json:"dynamicField"`
	DynamicFields *DynamicFieldConnection `json:"dynamicFields"`
	DynamicObjectField *DynamicField `json:"dynamicObjectField"`
	HasPublicTransfer bool `json:"hasPublicTransfer"`
	IconUrl string `json:"iconUrl"`
	MoveObjectBcs string `json:"moveObjectBcs"`
	MultiGetBalances []*Balance `json:"multiGetBalances"`
	MultiGetDynamicFields []*DynamicField `json:"multiGetDynamicFields"`
	MultiGetDynamicObjectFields []*DynamicField `json:"multiGetDynamicObjectFields"`
	Name string `json:"name"`
	ObjectAt *Object `json:"objectAt"`
	ObjectBcs string `json:"objectBcs"`
	ObjectVersionsAfter *ObjectConnection `json:"objectVersionsAfter"`
	ObjectVersionsBefore *ObjectConnection `json:"objectVersionsBefore"`
	Objects *MoveObjectConnection `json:"objects"`
	Owner interface{} `json:"owner"`
	PreviousTransaction *Transaction `json:"previousTransaction"`
	ReceivedTransactions *TransactionConnection `json:"receivedTransactions"`
	RegulatedState RegulatedState `json:"regulatedState"`
	StorageRebate string `json:"storageRebate"`
	Supply string `json:"supply"`
	SupplyState SupplyState `json:"supplyState"`
	Symbol string `json:"symbol"`
	Version string `json:"version"`
}

type CoinRegistryCreateTransaction struct {
	_ bool `json:"_"`
}

type CommandConnection struct {
	Edges []*CommandEdge `json:"edges"`
	Nodes []interface{} `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type CommandEdge struct {
	Cursor string `json:"cursor"`
	Node interface{} `json:"node"`
}

type CommandOutput struct {
	Argument interface{} `json:"argument"`
	Value *MoveValue `json:"value"`
}

type CommandResult struct {
	MutatedReferences []*CommandOutput `json:"mutatedReferences"`
	ReturnValues []*CommandOutput `json:"returnValues"`
}

type ConsensusAddressOwner struct {
	Address *Address `json:"address"`
	StartVersion string `json:"startVersion"`
}

type ConsensusCommitPrologueTransaction struct {
	AdditionalStateDigest string `json:"additionalStateDigest"`
	CommitTimestamp string `json:"commitTimestamp"`
	ConsensusCommitDigest string `json:"consensusCommitDigest"`
	Epoch *Epoch `json:"epoch"`
	Round string `json:"round"`
	SubDagIndex string `json:"subDagIndex"`
}

type ConsensusObjectCancelled struct {
	Address string `json:"address"`
	CancellationReason ConsensusObjectCancellationReason `json:"cancellationReason"`
}

type ConsensusObjectRead struct {
	Object *Object `json:"object"`
}

type Display struct {
	Errors string `json:"errors"`
	Output string `json:"output"`
}

type DisplayRegistryCreateTransaction struct {
	_ bool `json:"_"`
}

type DynamicField struct {
	Address string `json:"address"`
	Balance *Balance `json:"balance"`
	Balances *BalanceConnection `json:"balances"`
	Contents *MoveValue `json:"contents"`
	DefaultSuinsName string `json:"defaultSuinsName"`
	Digest string `json:"digest"`
	DynamicField *DynamicField `json:"dynamicField"`
	DynamicFields *DynamicFieldConnection `json:"dynamicFields"`
	DynamicObjectField *DynamicField `json:"dynamicObjectField"`
	HasPublicTransfer bool `json:"hasPublicTransfer"`
	MoveObjectBcs string `json:"moveObjectBcs"`
	MultiGetBalances []*Balance `json:"multiGetBalances"`
	MultiGetDynamicFields []*DynamicField `json:"multiGetDynamicFields"`
	MultiGetDynamicObjectFields []*DynamicField `json:"multiGetDynamicObjectFields"`
	Name *MoveValue `json:"name"`
	ObjectAt *Object `json:"objectAt"`
	ObjectBcs string `json:"objectBcs"`
	ObjectVersionsAfter *ObjectConnection `json:"objectVersionsAfter"`
	ObjectVersionsBefore *ObjectConnection `json:"objectVersionsBefore"`
	Objects *MoveObjectConnection `json:"objects"`
	Owner interface{} `json:"owner"`
	PreviousTransaction *Transaction `json:"previousTransaction"`
	ReceivedTransactions *TransactionConnection `json:"receivedTransactions"`
	StorageRebate string `json:"storageRebate"`
	Value interface{} `json:"value"`
	Version string `json:"version"`
}

type DynamicFieldConnection struct {
	Edges []*DynamicFieldEdge `json:"edges"`
	Nodes []*DynamicField `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type DynamicFieldEdge struct {
	Cursor string `json:"cursor"`
	Node *DynamicField `json:"node"`
}

type EndOfEpochTransaction struct {
	Transactions *EndOfEpochTransactionKindConnection `json:"transactions"`
}

type EndOfEpochTransactionKindConnection struct {
	Edges []*EndOfEpochTransactionKindEdge `json:"edges"`
	Nodes []interface{} `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type EndOfEpochTransactionKindEdge struct {
	Cursor string `json:"cursor"`
	Node interface{} `json:"node"`
}

type Epoch struct {
	Checkpoints *CheckpointConnection `json:"checkpoints"`
	CoinDenyList *Object `json:"coinDenyList"`
	EndTimestamp string `json:"endTimestamp"`
	EpochId string `json:"epochId"`
	FundInflow string `json:"fundInflow"`
	FundOutflow string `json:"fundOutflow"`
	FundSize string `json:"fundSize"`
	LiveObjectSetDigest string `json:"liveObjectSetDigest"`
	NetInflow string `json:"netInflow"`
	ProtocolConfigs *ProtocolConfigs `json:"protocolConfigs"`
	ReferenceGasPrice string `json:"referenceGasPrice"`
	SafeMode *SafeMode `json:"safeMode"`
	StartTimestamp string `json:"startTimestamp"`
	StorageFund *StorageFund `json:"storageFund"`
	SystemPackages *MovePackageConnection `json:"systemPackages"`
	SystemParameters *SystemParameters `json:"systemParameters"`
	SystemStakeSubsidy *StakeSubsidy `json:"systemStakeSubsidy"`
	SystemStateVersion string `json:"systemStateVersion"`
	TotalCheckpoints string `json:"totalCheckpoints"`
	TotalGasFees string `json:"totalGasFees"`
	TotalStakeRewards string `json:"totalStakeRewards"`
	TotalStakeSubsidies string `json:"totalStakeSubsidies"`
	TotalTransactions string `json:"totalTransactions"`
	Transactions *TransactionConnection `json:"transactions"`
	ValidatorSet *ValidatorSet `json:"validatorSet"`
}

type EpochConnection struct {
	Edges []*EpochEdge `json:"edges"`
	Nodes []*Epoch `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type EpochEdge struct {
	Cursor string `json:"cursor"`
	Node *Epoch `json:"node"`
}

type Event struct {
	Contents *MoveValue `json:"contents"`
	EventBcs string `json:"eventBcs"`
	Sender *Address `json:"sender"`
	SequenceNumber string `json:"sequenceNumber"`
	Timestamp string `json:"timestamp"`
	Transaction *Transaction `json:"transaction"`
	TransactionModule *MoveModule `json:"transactionModule"`
}

type EventConnection struct {
	Edges []*EventEdge `json:"edges"`
	Nodes []*Event `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type EventEdge struct {
	Cursor string `json:"cursor"`
	Node *Event `json:"node"`
}

type ExecutionError struct {
	AbortCode string `json:"abortCode"`
	Constant string `json:"constant"`
	Function *MoveFunction `json:"function"`
	Identifier string `json:"identifier"`
	InstructionOffset int `json:"instructionOffset"`
	Message string `json:"message"`
	Module *MoveModule `json:"module"`
	SourceLineNumber int `json:"sourceLineNumber"`
}

type ExecutionResult struct {
	Effects *TransactionEffects `json:"effects"`
	Errors []string `json:"errors"`
}

type FeatureFlag struct {
	Key string `json:"key"`
	Value bool `json:"value"`
}

type GasCoin struct {
	_ bool `json:"_"`
}

type GasCostSummary struct {
	ComputationCost string `json:"computationCost"`
	NonRefundableStorageFee string `json:"nonRefundableStorageFee"`
	StorageCost string `json:"storageCost"`
	StorageRebate string `json:"storageRebate"`
}

type GasEffects struct {
	GasObject *Object `json:"gasObject"`
	GasSummary *GasCostSummary `json:"gasSummary"`
}

type GasInput struct {
	GasBudget string `json:"gasBudget"`
	GasPayment *ObjectConnection `json:"gasPayment"`
	GasPrice string `json:"gasPrice"`
	GasSponsor *Address `json:"gasSponsor"`
}

type GenesisTransaction struct {
	Objects *ObjectConnection `json:"objects"`
}

type Immutable struct {
	_ bool `json:"_"`
}

type Input struct {
	Ix int `json:"ix"`
}

type Linkage struct {
	OriginalId string `json:"originalId"`
	UpgradedId string `json:"upgradedId"`
	Version string `json:"version"`
}

type MakeMoveVecCommand struct {
	Elements []interface{} `json:"elements"`
	Type *MoveType `json:"type"`
}

type MergeCoinsCommand struct {
	Coin interface{} `json:"coin"`
	Coins []interface{} `json:"coins"`
}

type MoveCallCommand struct {
	Arguments []interface{} `json:"arguments"`
	Function *MoveFunction `json:"function"`
}

type MoveDatatype struct {
	Abilities []MoveAbility `json:"abilities"`
	AsMoveEnum *MoveEnum `json:"asMoveEnum"`
	AsMoveStruct *MoveStruct `json:"asMoveStruct"`
	Module *MoveModule `json:"module"`
	Name string `json:"name"`
	TypeParameters []*MoveDatatypeTypeParameter `json:"typeParameters"`
}

type MoveDatatypeConnection struct {
	Edges []*MoveDatatypeEdge `json:"edges"`
	Nodes []*MoveDatatype `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type MoveDatatypeEdge struct {
	Cursor string `json:"cursor"`
	Node *MoveDatatype `json:"node"`
}

type MoveDatatypeTypeParameter struct {
	Constraints []MoveAbility `json:"constraints"`
	IsPhantom bool `json:"isPhantom"`
}

type MoveEnum struct {
	Abilities []MoveAbility `json:"abilities"`
	Module *MoveModule `json:"module"`
	Name string `json:"name"`
	TypeParameters []*MoveDatatypeTypeParameter `json:"typeParameters"`
	Variants []*MoveEnumVariant `json:"variants"`
}

type MoveEnumConnection struct {
	Edges []*MoveEnumEdge `json:"edges"`
	Nodes []*MoveEnum `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type MoveEnumEdge struct {
	Cursor string `json:"cursor"`
	Node *MoveEnum `json:"node"`
}

type MoveEnumVariant struct {
	Fields []*MoveField `json:"fields"`
	Name string `json:"name"`
}

type MoveField struct {
	Name string `json:"name"`
	Type *OpenMoveType `json:"type"`
}

type MoveFunction struct {
	IsEntry bool `json:"isEntry"`
	Module *MoveModule `json:"module"`
	Name string `json:"name"`
	Parameters []*OpenMoveType `json:"parameters"`
	Return []*OpenMoveType `json:"return"`
	TypeParameters []*MoveFunctionTypeParameter `json:"typeParameters"`
	Visibility MoveVisibility `json:"visibility"`
}

type MoveFunctionConnection struct {
	Edges []*MoveFunctionEdge `json:"edges"`
	Nodes []*MoveFunction `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type MoveFunctionEdge struct {
	Cursor string `json:"cursor"`
	Node *MoveFunction `json:"node"`
}

type MoveFunctionTypeParameter struct {
	Constraints []MoveAbility `json:"constraints"`
}

type MoveModule struct {
	Bytes string `json:"bytes"`
	Datatype *MoveDatatype `json:"datatype"`
	Datatypes *MoveDatatypeConnection `json:"datatypes"`
	Disassembly string `json:"disassembly"`
	Enum *MoveEnum `json:"enum"`
	Enums *MoveEnumConnection `json:"enums"`
	FileFormatVersion int `json:"fileFormatVersion"`
	Friends *MoveModuleConnection `json:"friends"`
	Function *MoveFunction `json:"function"`
	Functions *MoveFunctionConnection `json:"functions"`
	Name string `json:"name"`
	Package *MovePackage `json:"package"`
	Struct *MoveStruct `json:"struct"`
	Structs *MoveStructConnection `json:"structs"`
}

type MoveModuleConnection struct {
	Edges []*MoveModuleEdge `json:"edges"`
	Nodes []*MoveModule `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type MoveModuleEdge struct {
	Cursor string `json:"cursor"`
	Node *MoveModule `json:"node"`
}

type MoveObject struct {
	Address string `json:"address"`
	AsCoinMetadata *CoinMetadata `json:"asCoinMetadata"`
	AsDynamicField *DynamicField `json:"asDynamicField"`
	Balance *Balance `json:"balance"`
	Balances *BalanceConnection `json:"balances"`
	Contents *MoveValue `json:"contents"`
	DefaultSuinsName string `json:"defaultSuinsName"`
	Digest string `json:"digest"`
	DynamicField *DynamicField `json:"dynamicField"`
	DynamicFields *DynamicFieldConnection `json:"dynamicFields"`
	DynamicObjectField *DynamicField `json:"dynamicObjectField"`
	HasPublicTransfer bool `json:"hasPublicTransfer"`
	MoveObjectBcs string `json:"moveObjectBcs"`
	MultiGetBalances []*Balance `json:"multiGetBalances"`
	MultiGetDynamicFields []*DynamicField `json:"multiGetDynamicFields"`
	MultiGetDynamicObjectFields []*DynamicField `json:"multiGetDynamicObjectFields"`
	ObjectAt *Object `json:"objectAt"`
	ObjectBcs string `json:"objectBcs"`
	ObjectVersionsAfter *ObjectConnection `json:"objectVersionsAfter"`
	ObjectVersionsBefore *ObjectConnection `json:"objectVersionsBefore"`
	Objects *MoveObjectConnection `json:"objects"`
	Owner interface{} `json:"owner"`
	PreviousTransaction *Transaction `json:"previousTransaction"`
	ReceivedTransactions *TransactionConnection `json:"receivedTransactions"`
	StorageRebate string `json:"storageRebate"`
	Version string `json:"version"`
}

type MoveObjectConnection struct {
	Edges []*MoveObjectEdge `json:"edges"`
	Nodes []*MoveObject `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type MoveObjectEdge struct {
	Cursor string `json:"cursor"`
	Node *MoveObject `json:"node"`
}

type MovePackage struct {
	Address string `json:"address"`
	Balance *Balance `json:"balance"`
	Balances *BalanceConnection `json:"balances"`
	DefaultSuinsName string `json:"defaultSuinsName"`
	Digest string `json:"digest"`
	Linkage []*Linkage `json:"linkage"`
	Module *MoveModule `json:"module"`
	ModuleBcs string `json:"moduleBcs"`
	Modules *MoveModuleConnection `json:"modules"`
	MultiGetBalances []*Balance `json:"multiGetBalances"`
	ObjectAt *Object `json:"objectAt"`
	ObjectBcs string `json:"objectBcs"`
	ObjectVersionsAfter *ObjectConnection `json:"objectVersionsAfter"`
	ObjectVersionsBefore *ObjectConnection `json:"objectVersionsBefore"`
	Objects *MoveObjectConnection `json:"objects"`
	Owner interface{} `json:"owner"`
	PackageAt *MovePackage `json:"packageAt"`
	PackageBcs string `json:"packageBcs"`
	PackageVersionsAfter *MovePackageConnection `json:"packageVersionsAfter"`
	PackageVersionsBefore *MovePackageConnection `json:"packageVersionsBefore"`
	PreviousTransaction *Transaction `json:"previousTransaction"`
	ReceivedTransactions *TransactionConnection `json:"receivedTransactions"`
	StorageRebate string `json:"storageRebate"`
	TypeOrigins []*TypeOrigin `json:"typeOrigins"`
	Version string `json:"version"`
}

type MovePackageConnection struct {
	Edges []*MovePackageEdge `json:"edges"`
	Nodes []*MovePackage `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type MovePackageEdge struct {
	Cursor string `json:"cursor"`
	Node *MovePackage `json:"node"`
}

type MoveStruct struct {
	Abilities []MoveAbility `json:"abilities"`
	Fields []*MoveField `json:"fields"`
	Module *MoveModule `json:"module"`
	Name string `json:"name"`
	TypeParameters []*MoveDatatypeTypeParameter `json:"typeParameters"`
}

type MoveStructConnection struct {
	Edges []*MoveStructEdge `json:"edges"`
	Nodes []*MoveStruct `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type MoveStructEdge struct {
	Cursor string `json:"cursor"`
	Node *MoveStruct `json:"node"`
}

type MoveType struct {
	Abilities []MoveAbility `json:"abilities"`
	Layout string `json:"layout"`
	Repr string `json:"repr"`
	Signature string `json:"signature"`
}

type MoveValue struct {
	Bcs string `json:"bcs"`
	Display *Display `json:"display"`
	Json string `json:"json"`
	Type *MoveType `json:"type"`
}

type MutateConsensusStreamEnded struct {
	Address string `json:"address"`
	SequenceNumber string `json:"sequenceNumber"`
}

type Mutation struct {
	ExecuteTransaction *ExecutionResult `json:"executeTransaction"`
}

type Object struct {
	Address string `json:"address"`
	AsMoveObject *MoveObject `json:"asMoveObject"`
	AsMovePackage *MovePackage `json:"asMovePackage"`
	Balance *Balance `json:"balance"`
	Balances *BalanceConnection `json:"balances"`
	DefaultSuinsName string `json:"defaultSuinsName"`
	Digest string `json:"digest"`
	DynamicField *DynamicField `json:"dynamicField"`
	DynamicFields *DynamicFieldConnection `json:"dynamicFields"`
	DynamicObjectField *DynamicField `json:"dynamicObjectField"`
	MultiGetBalances []*Balance `json:"multiGetBalances"`
	MultiGetDynamicFields []*DynamicField `json:"multiGetDynamicFields"`
	MultiGetDynamicObjectFields []*DynamicField `json:"multiGetDynamicObjectFields"`
	ObjectAt *Object `json:"objectAt"`
	ObjectBcs string `json:"objectBcs"`
	ObjectVersionsAfter *ObjectConnection `json:"objectVersionsAfter"`
	ObjectVersionsBefore *ObjectConnection `json:"objectVersionsBefore"`
	Objects *MoveObjectConnection `json:"objects"`
	Owner interface{} `json:"owner"`
	PreviousTransaction *Transaction `json:"previousTransaction"`
	ReceivedTransactions *TransactionConnection `json:"receivedTransactions"`
	StorageRebate string `json:"storageRebate"`
	Version string `json:"version"`
}

type ObjectChange struct {
	Address string `json:"address"`
	IdCreated bool `json:"idCreated"`
	IdDeleted bool `json:"idDeleted"`
	InputState *Object `json:"inputState"`
	OutputState *Object `json:"outputState"`
}

type ObjectChangeConnection struct {
	Edges []*ObjectChangeEdge `json:"edges"`
	Nodes []*ObjectChange `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type ObjectChangeEdge struct {
	Cursor string `json:"cursor"`
	Node *ObjectChange `json:"node"`
}

type ObjectConnection struct {
	Edges []*ObjectEdge `json:"edges"`
	Nodes []*Object `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type ObjectEdge struct {
	Cursor string `json:"cursor"`
	Node *Object `json:"node"`
}

type ObjectOwner struct {
	Address *Address `json:"address"`
}

type OpenMoveType struct {
	Repr string `json:"repr"`
	Signature string `json:"signature"`
}

type OtherCommand struct {
	_ bool `json:"_"`
}

type OwnedOrImmutable struct {
	Object *Object `json:"object"`
}

type PageInfo struct {
	EndCursor string `json:"endCursor"`
	HasNextPage bool `json:"hasNextPage"`
	HasPreviousPage bool `json:"hasPreviousPage"`
	StartCursor string `json:"startCursor"`
}

type PerEpochConfig struct {
	Object *Object `json:"object"`
}

type ProgrammableSystemTransaction struct {
	Commands *CommandConnection `json:"commands"`
	Inputs *TransactionInputConnection `json:"inputs"`
}

type ProgrammableTransaction struct {
	Commands *CommandConnection `json:"commands"`
	Inputs *TransactionInputConnection `json:"inputs"`
}

type ProtocolConfig struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

type ProtocolConfigs struct {
	Config *ProtocolConfig `json:"config"`
	Configs []*ProtocolConfig `json:"configs"`
	FeatureFlag *FeatureFlag `json:"featureFlag"`
	FeatureFlags []*FeatureFlag `json:"featureFlags"`
	ProtocolVersion string `json:"protocolVersion"`
}

type PublishCommand struct {
	Dependencies []string `json:"dependencies"`
	Modules []string `json:"modules"`
}

type Pure struct {
	Bytes string `json:"bytes"`
}

type Query struct {
	Address *Address `json:"address"`
	ChainIdentifier string `json:"chainIdentifier"`
	Checkpoint *Checkpoint `json:"checkpoint"`
	Checkpoints *CheckpointConnection `json:"checkpoints"`
	CoinMetadata *CoinMetadata `json:"coinMetadata"`
	Epoch *Epoch `json:"epoch"`
	Epochs *EpochConnection `json:"epochs"`
	Events *EventConnection `json:"events"`
	MultiGetCheckpoints []*Checkpoint `json:"multiGetCheckpoints"`
	MultiGetEpochs []*Epoch `json:"multiGetEpochs"`
	MultiGetObjects []*Object `json:"multiGetObjects"`
	MultiGetPackages []*MovePackage `json:"multiGetPackages"`
	MultiGetTransactionEffects []*TransactionEffects `json:"multiGetTransactionEffects"`
	MultiGetTransactions []*Transaction `json:"multiGetTransactions"`
	MultiGetTypes []*MoveType `json:"multiGetTypes"`
	Object *Object `json:"object"`
	ObjectVersions *ObjectConnection `json:"objectVersions"`
	Objects *ObjectConnection `json:"objects"`
	Package *MovePackage `json:"package"`
	PackageVersions *MovePackageConnection `json:"packageVersions"`
	Packages *MovePackageConnection `json:"packages"`
	ProtocolConfigs *ProtocolConfigs `json:"protocolConfigs"`
	ServiceConfig *ServiceConfig `json:"serviceConfig"`
	SimulateTransaction *SimulationResult `json:"simulateTransaction"`
	SuinsName *Address `json:"suinsName"`
	Transaction *Transaction `json:"transaction"`
	TransactionEffects *TransactionEffects `json:"transactionEffects"`
	Transactions *TransactionConnection `json:"transactions"`
	Type *MoveType `json:"type"`
	VerifyZkLoginSignature *ZkLoginVerifyResult `json:"verifyZkLoginSignature"`
}

type RandomnessStateCreateTransaction struct {
	_ bool `json:"_"`
}

type RandomnessStateUpdateTransaction struct {
	Epoch int `json:"epoch"`
	RandomBytes string `json:"randomBytes"`
	RandomnessObjInitialSharedVersion int `json:"randomnessObjInitialSharedVersion"`
	RandomnessRound int `json:"randomnessRound"`
}

type ReadConsensusStreamEnded struct {
	Address string `json:"address"`
	SequenceNumber string `json:"sequenceNumber"`
}

type Receiving struct {
	Object *Object `json:"object"`
}

type SafeMode struct {
	Enabled bool `json:"enabled"`
	GasSummary *GasCostSummary `json:"gasSummary"`
}

type ServiceConfig struct {
	AvailableRange *AvailableRange `json:"availableRange"`
	DefaultPageSize int `json:"defaultPageSize"`
	MaxDisassembledModuleSize int `json:"maxDisassembledModuleSize"`
	MaxDisplayFieldDepth int `json:"maxDisplayFieldDepth"`
	MaxDisplayOutputSize int `json:"maxDisplayOutputSize"`
	MaxMoveValueBound int `json:"maxMoveValueBound"`
	MaxMoveValueDepth int `json:"maxMoveValueDepth"`
	MaxMultiGetSize int `json:"maxMultiGetSize"`
	MaxOutputNodes int `json:"maxOutputNodes"`
	MaxPageSize int `json:"maxPageSize"`
	MaxQueryDepth int `json:"maxQueryDepth"`
	MaxQueryNodes int `json:"maxQueryNodes"`
	MaxQueryPayloadSize int `json:"maxQueryPayloadSize"`
	MaxTransactionPayloadSize int `json:"maxTransactionPayloadSize"`
	MaxTypeArgumentDepth int `json:"maxTypeArgumentDepth"`
	MaxTypeArgumentWidth int `json:"maxTypeArgumentWidth"`
	MaxTypeNodes int `json:"maxTypeNodes"`
	MutationTimeoutMs int `json:"mutationTimeoutMs"`
	QueryTimeoutMs int `json:"queryTimeoutMs"`
}

type Shared struct {
	InitialSharedVersion string `json:"initialSharedVersion"`
}

type SharedInput struct {
	Address string `json:"address"`
	InitialSharedVersion string `json:"initialSharedVersion"`
	Mutable bool `json:"mutable"`
}

type SimulationResult struct {
	Effects *TransactionEffects `json:"effects"`
	Error string `json:"error"`
	Outputs []*CommandResult `json:"outputs"`
}

type SplitCoinsCommand struct {
	Amounts []interface{} `json:"amounts"`
	Coin interface{} `json:"coin"`
}

type StakeSubsidy struct {
	Balance string `json:"balance"`
	CurrentDistributionAmount string `json:"currentDistributionAmount"`
	DecreaseRate int `json:"decreaseRate"`
	DistributionCounter int `json:"distributionCounter"`
	PeriodLength int `json:"periodLength"`
}

type StorageFund struct {
	NonRefundableBalance string `json:"nonRefundableBalance"`
	TotalObjectStorageRebates string `json:"totalObjectStorageRebates"`
}

type StoreExecutionTimeObservationsTransaction struct {
	_ bool `json:"_"`
}

type SystemParameters struct {
	DurationMs string `json:"durationMs"`
	MaxValidatorCount int `json:"maxValidatorCount"`
	MinValidatorCount int `json:"minValidatorCount"`
	MinValidatorJoiningStake string `json:"minValidatorJoiningStake"`
	StakeSubsidyStartEpoch string `json:"stakeSubsidyStartEpoch"`
	ValidatorLowStakeGracePeriod string `json:"validatorLowStakeGracePeriod"`
	ValidatorLowStakeThreshold string `json:"validatorLowStakeThreshold"`
	ValidatorVeryLowStakeThreshold string `json:"validatorVeryLowStakeThreshold"`
}

type Transaction struct {
	Digest string `json:"digest"`
	Effects *TransactionEffects `json:"effects"`
	Expiration *Epoch `json:"expiration"`
	GasInput *GasInput `json:"gasInput"`
	Kind interface{} `json:"kind"`
	Sender *Address `json:"sender"`
	Signatures []*UserSignature `json:"signatures"`
	TransactionBcs string `json:"transactionBcs"`
	TransactionJson string `json:"transactionJson"`
}

type TransactionConnection struct {
	Edges []*TransactionEdge `json:"edges"`
	Nodes []*Transaction `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type TransactionEdge struct {
	Cursor string `json:"cursor"`
	Node *Transaction `json:"node"`
}

type TransactionEffects struct {
	BalanceChanges *BalanceChangeConnection `json:"balanceChanges"`
	BalanceChangesJson string `json:"balanceChangesJson"`
	Checkpoint *Checkpoint `json:"checkpoint"`
	Dependencies *TransactionConnection `json:"dependencies"`
	Digest string `json:"digest"`
	EffectsBcs string `json:"effectsBcs"`
	EffectsDigest string `json:"effectsDigest"`
	EffectsJson string `json:"effectsJson"`
	Epoch *Epoch `json:"epoch"`
	Events *EventConnection `json:"events"`
	ExecutionError *ExecutionError `json:"executionError"`
	GasEffects *GasEffects `json:"gasEffects"`
	LamportVersion string `json:"lamportVersion"`
	ObjectChanges *ObjectChangeConnection `json:"objectChanges"`
	Status ExecutionStatus `json:"status"`
	Timestamp string `json:"timestamp"`
	Transaction *Transaction `json:"transaction"`
	UnchangedConsensusObjects *UnchangedConsensusObjectConnection `json:"unchangedConsensusObjects"`
}

type TransactionInputConnection struct {
	Edges []*TransactionInputEdge `json:"edges"`
	Nodes []interface{} `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type TransactionInputEdge struct {
	Cursor string `json:"cursor"`
	Node interface{} `json:"node"`
}

type TransferObjectsCommand struct {
	Address interface{} `json:"address"`
	Inputs []interface{} `json:"inputs"`
}

type TxResult struct {
	Cmd int `json:"cmd"`
	Ix int `json:"ix"`
}

type TypeOrigin struct {
	DefiningId string `json:"definingId"`
	Module string `json:"module"`
	Struct string `json:"struct"`
}

type UnchangedConsensusObjectConnection struct {
	Edges []*UnchangedConsensusObjectEdge `json:"edges"`
	Nodes []interface{} `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type UnchangedConsensusObjectEdge struct {
	Cursor string `json:"cursor"`
	Node interface{} `json:"node"`
}

type UpgradeCommand struct {
	CurrentPackage string `json:"currentPackage"`
	Dependencies []string `json:"dependencies"`
	Modules []string `json:"modules"`
	UpgradeTicket interface{} `json:"upgradeTicket"`
}

type UserSignature struct {
	SignatureBytes string `json:"signatureBytes"`
}

type Validator struct {
	Address string `json:"address"`
	AtRisk string `json:"atRisk"`
	Balance *Balance `json:"balance"`
	Balances *BalanceConnection `json:"balances"`
	CommissionRate int `json:"commissionRate"`
	Credentials *ValidatorCredentials `json:"credentials"`
	DefaultSuinsName string `json:"defaultSuinsName"`
	Description string `json:"description"`
	ExchangeRatesSize string `json:"exchangeRatesSize"`
	ExchangeRatesTable *Address `json:"exchangeRatesTable"`
	GasPrice string `json:"gasPrice"`
	ImageUrl string `json:"imageUrl"`
	MultiGetBalances []*Balance `json:"multiGetBalances"`
	Name string `json:"name"`
	NextEpochCommissionRate int `json:"nextEpochCommissionRate"`
	NextEpochCredentials *ValidatorCredentials `json:"nextEpochCredentials"`
	NextEpochGasPrice string `json:"nextEpochGasPrice"`
	NextEpochStake string `json:"nextEpochStake"`
	Objects *MoveObjectConnection `json:"objects"`
	OperationCap *MoveObject `json:"operationCap"`
	PendingPoolTokenWithdraw string `json:"pendingPoolTokenWithdraw"`
	PendingStake string `json:"pendingStake"`
	PendingTotalSuiWithdraw string `json:"pendingTotalSuiWithdraw"`
	PoolTokenBalance string `json:"poolTokenBalance"`
	ProjectUrl string `json:"projectUrl"`
	ReportRecords *ValidatorConnection `json:"reportRecords"`
	RewardsPool string `json:"rewardsPool"`
	StakingPoolActivationEpoch string `json:"stakingPoolActivationEpoch"`
	StakingPoolId string `json:"stakingPoolId"`
	StakingPoolSuiBalance string `json:"stakingPoolSuiBalance"`
	VotingPower int `json:"votingPower"`
}

type ValidatorAggregatedSignature struct {
	Epoch *Epoch `json:"epoch"`
	Signature string `json:"signature"`
	SignersMap []int `json:"signersMap"`
}

type ValidatorConnection struct {
	Edges []*ValidatorEdge `json:"edges"`
	Nodes []*Validator `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type ValidatorCredentials struct {
	NetAddress string `json:"netAddress"`
	NetworkPubKey string `json:"networkPubKey"`
	P2PAddress string `json:"p2PAddress"`
	PrimaryAddress string `json:"primaryAddress"`
	ProofOfPossession string `json:"proofOfPossession"`
	ProtocolPubKey string `json:"protocolPubKey"`
	WorkerAddress string `json:"workerAddress"`
	WorkerPubKey string `json:"workerPubKey"`
}

type ValidatorEdge struct {
	Cursor string `json:"cursor"`
	Node *Validator `json:"node"`
}

type ValidatorSet struct {
	ActiveValidators *ValidatorConnection `json:"activeValidators"`
	InactivePoolsId string `json:"inactivePoolsId"`
	InactivePoolsSize int `json:"inactivePoolsSize"`
	PendingActiveValidatorsId string `json:"pendingActiveValidatorsId"`
	PendingActiveValidatorsSize int `json:"pendingActiveValidatorsSize"`
	PendingRemovals []int `json:"pendingRemovals"`
	StakingPoolMappingsId string `json:"stakingPoolMappingsId"`
	StakingPoolMappingsSize int `json:"stakingPoolMappingsSize"`
	TotalStake string `json:"totalStake"`
	ValidatorCandidatesId string `json:"validatorCandidatesId"`
	ValidatorCandidatesSize int `json:"validatorCandidatesSize"`
}

type WriteAccumulatorStorageCostTransaction struct {
	_ bool `json:"_"`
}

type ZkLoginVerifyResult struct {
	Error string `json:"error"`
	Success bool `json:"success"`
}

type CheckpointFilter struct {
	AfterCheckpoint string `json:"afterCheckpoint"`
	AtCheckpoint string `json:"atCheckpoint"`
	AtEpoch string `json:"atEpoch"`
	BeforeCheckpoint string `json:"beforeCheckpoint"`
}

type DynamicFieldName struct {
	Bcs string `json:"bcs"`
	Type string `json:"type"`
}

type EventFilter struct {
	AfterCheckpoint string `json:"afterCheckpoint"`
	AtCheckpoint string `json:"atCheckpoint"`
	BeforeCheckpoint string `json:"beforeCheckpoint"`
	Module string `json:"module"`
	Sender string `json:"sender"`
	Type string `json:"type"`
}

type ObjectFilter struct {
	Owner string `json:"owner"`
	OwnerKind OwnerKind `json:"ownerKind"`
	Type string `json:"type"`
}

type ObjectKey struct {
	Address string `json:"address"`
	AtCheckpoint string `json:"atCheckpoint"`
	RootVersion string `json:"rootVersion"`
	Version string `json:"version"`
}

type PackageCheckpointFilter struct {
	AfterCheckpoint string `json:"afterCheckpoint"`
	BeforeCheckpoint string `json:"beforeCheckpoint"`
}

type PackageKey struct {
	Address string `json:"address"`
	AtCheckpoint string `json:"atCheckpoint"`
	Version string `json:"version"`
}

type TransactionFilter struct {
	AffectedAddress string `json:"affectedAddress"`
	AffectedObject string `json:"affectedObject"`
	AfterCheckpoint string `json:"afterCheckpoint"`
	AtCheckpoint string `json:"atCheckpoint"`
	BeforeCheckpoint string `json:"beforeCheckpoint"`
	Function string `json:"function"`
	Kind TransactionKindInput `json:"kind"`
	SentAddress string `json:"sentAddress"`
}

type VersionFilter struct {
	AfterVersion string `json:"afterVersion"`
	BeforeVersion string `json:"beforeVersion"`
}

