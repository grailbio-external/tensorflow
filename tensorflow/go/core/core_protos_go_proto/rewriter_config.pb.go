// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/rewriter_config.proto

package core_protos_go_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	attr_value_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/attr_value_go_proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RewriterConfig_Toggle int32

const (
	RewriterConfig_DEFAULT RewriterConfig_Toggle = 0
	RewriterConfig_ON      RewriterConfig_Toggle = 1
	RewriterConfig_OFF     RewriterConfig_Toggle = 2
	// Enable some aggressive optimizations that use assumptions that TF graphs
	// may break. For example, assume the shape of a placeholder matches its
	// actual feed.
	RewriterConfig_AGGRESSIVE RewriterConfig_Toggle = 3
)

var RewriterConfig_Toggle_name = map[int32]string{
	0: "DEFAULT",
	1: "ON",
	2: "OFF",
	3: "AGGRESSIVE",
}

var RewriterConfig_Toggle_value = map[string]int32{
	"DEFAULT":    0,
	"ON":         1,
	"OFF":        2,
	"AGGRESSIVE": 3,
}

func (x RewriterConfig_Toggle) String() string {
	return proto.EnumName(RewriterConfig_Toggle_name, int32(x))
}

func (RewriterConfig_Toggle) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1dd7de60bf190bbb, []int{2, 0}
}

// Enum controlling the number of times to run optimizers. The default is to
// run them twice.
type RewriterConfig_NumIterationsType int32

const (
	RewriterConfig_DEFAULT_NUM_ITERS RewriterConfig_NumIterationsType = 0
	RewriterConfig_ONE               RewriterConfig_NumIterationsType = 1
	RewriterConfig_TWO               RewriterConfig_NumIterationsType = 2
)

var RewriterConfig_NumIterationsType_name = map[int32]string{
	0: "DEFAULT_NUM_ITERS",
	1: "ONE",
	2: "TWO",
}

var RewriterConfig_NumIterationsType_value = map[string]int32{
	"DEFAULT_NUM_ITERS": 0,
	"ONE":               1,
	"TWO":               2,
}

func (x RewriterConfig_NumIterationsType) String() string {
	return proto.EnumName(RewriterConfig_NumIterationsType_name, int32(x))
}

func (RewriterConfig_NumIterationsType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1dd7de60bf190bbb, []int{2, 1}
}

type RewriterConfig_MemOptType int32

const (
	// The default setting (SCHEDULING and SWAPPING HEURISTICS only)
	RewriterConfig_DEFAULT_MEM_OPT RewriterConfig_MemOptType = 0
	// Disabled in the meta-optimizer.
	RewriterConfig_NO_MEM_OPT RewriterConfig_MemOptType = 1
	// Driven by manual op-level annotations.
	RewriterConfig_MANUAL RewriterConfig_MemOptType = 2
	// Swapping heuristic will move a tensor from the GPU to the CPU and move
	// it back when needed to reduce peak memory usage.
	RewriterConfig_SWAPPING_HEURISTICS RewriterConfig_MemOptType = 4
	// Recomputation heuristics will recompute ops (such as Relu activation)
	// during backprop instead of storing them, reducing peak memory usage.
	RewriterConfig_RECOMPUTATION_HEURISTICS RewriterConfig_MemOptType = 5
	// Scheduling will split big ops such as AddN and try to enforce a schedule
	// of the new computations that decreases peak memory usage.
	RewriterConfig_SCHEDULING_HEURISTICS RewriterConfig_MemOptType = 6
	// Use any combination of swapping and recomputation heuristics.
	RewriterConfig_HEURISTICS RewriterConfig_MemOptType = 3
)

var RewriterConfig_MemOptType_name = map[int32]string{
	0: "DEFAULT_MEM_OPT",
	1: "NO_MEM_OPT",
	2: "MANUAL",
	4: "SWAPPING_HEURISTICS",
	5: "RECOMPUTATION_HEURISTICS",
	6: "SCHEDULING_HEURISTICS",
	3: "HEURISTICS",
}

var RewriterConfig_MemOptType_value = map[string]int32{
	"DEFAULT_MEM_OPT":          0,
	"NO_MEM_OPT":               1,
	"MANUAL":                   2,
	"SWAPPING_HEURISTICS":      4,
	"RECOMPUTATION_HEURISTICS": 5,
	"SCHEDULING_HEURISTICS":    6,
	"HEURISTICS":               3,
}

func (x RewriterConfig_MemOptType) String() string {
	return proto.EnumName(RewriterConfig_MemOptType_name, int32(x))
}

func (RewriterConfig_MemOptType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1dd7de60bf190bbb, []int{2, 2}
}

type AutoParallelOptions struct {
	Enable               bool     `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	NumReplicas          int32    `protobuf:"varint,2,opt,name=num_replicas,json=numReplicas,proto3" json:"num_replicas,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutoParallelOptions) Reset()         { *m = AutoParallelOptions{} }
func (m *AutoParallelOptions) String() string { return proto.CompactTextString(m) }
func (*AutoParallelOptions) ProtoMessage()    {}
func (*AutoParallelOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd7de60bf190bbb, []int{0}
}

func (m *AutoParallelOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoParallelOptions.Unmarshal(m, b)
}
func (m *AutoParallelOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoParallelOptions.Marshal(b, m, deterministic)
}
func (m *AutoParallelOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoParallelOptions.Merge(m, src)
}
func (m *AutoParallelOptions) XXX_Size() int {
	return xxx_messageInfo_AutoParallelOptions.Size(m)
}
func (m *AutoParallelOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoParallelOptions.DiscardUnknown(m)
}

var xxx_messageInfo_AutoParallelOptions proto.InternalMessageInfo

func (m *AutoParallelOptions) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *AutoParallelOptions) GetNumReplicas() int32 {
	if m != nil {
		return m.NumReplicas
	}
	return 0
}

type ScopedAllocatorOptions struct {
	// If present, only perform optimization for these ops.
	EnableOp             []string `protobuf:"bytes,1,rep,name=enable_op,json=enableOp,proto3" json:"enable_op,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScopedAllocatorOptions) Reset()         { *m = ScopedAllocatorOptions{} }
func (m *ScopedAllocatorOptions) String() string { return proto.CompactTextString(m) }
func (*ScopedAllocatorOptions) ProtoMessage()    {}
func (*ScopedAllocatorOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd7de60bf190bbb, []int{1}
}

func (m *ScopedAllocatorOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScopedAllocatorOptions.Unmarshal(m, b)
}
func (m *ScopedAllocatorOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScopedAllocatorOptions.Marshal(b, m, deterministic)
}
func (m *ScopedAllocatorOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScopedAllocatorOptions.Merge(m, src)
}
func (m *ScopedAllocatorOptions) XXX_Size() int {
	return xxx_messageInfo_ScopedAllocatorOptions.Size(m)
}
func (m *ScopedAllocatorOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ScopedAllocatorOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ScopedAllocatorOptions proto.InternalMessageInfo

func (m *ScopedAllocatorOptions) GetEnableOp() []string {
	if m != nil {
		return m.EnableOp
	}
	return nil
}

type RewriterConfig struct {
	// Optimize tensor layouts (default is ON)
	// e.g. This will try to use NCHW layout on GPU which is faster.
	LayoutOptimizer RewriterConfig_Toggle `protobuf:"varint,1,opt,name=layout_optimizer,json=layoutOptimizer,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"layout_optimizer,omitempty"`
	// Fold constants (default is ON)
	// Statically infer the value of tensors when possible, and materialize the
	// result using constants.
	ConstantFolding RewriterConfig_Toggle `protobuf:"varint,3,opt,name=constant_folding,json=constantFolding,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"constant_folding,omitempty"`
	// Shape optimizations (default is ON)
	// Simplify computations made on shapes.
	ShapeOptimization RewriterConfig_Toggle `protobuf:"varint,13,opt,name=shape_optimization,json=shapeOptimization,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"shape_optimization,omitempty"`
	// Remapping (default is ON)
	// Remap subgraphs onto more efficient implementations.
	Remapping RewriterConfig_Toggle `protobuf:"varint,14,opt,name=remapping,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"remapping,omitempty"`
	// Common subgraph elimination (default is ON)
	// e.g. Simplify arithmetic ops; merge ops with same value (like constants).
	CommonSubgraphElimination RewriterConfig_Toggle `protobuf:"varint,24,opt,name=common_subgraph_elimination,json=commonSubgraphElimination,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"common_subgraph_elimination,omitempty"`
	// Arithmetic optimizations (default is ON)
	// e.g. Simplify arithmetic ops; merge ops with same value (like constants).
	ArithmeticOptimization RewriterConfig_Toggle `protobuf:"varint,7,opt,name=arithmetic_optimization,json=arithmeticOptimization,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"arithmetic_optimization,omitempty"`
	// Control dependency optimizations (default is ON).
	// Remove redundant control dependencies, which may enable other optimization.
	DependencyOptimization RewriterConfig_Toggle `protobuf:"varint,8,opt,name=dependency_optimization,json=dependencyOptimization,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"dependency_optimization,omitempty"`
	// Loop optimizations (default is ON).
	LoopOptimization RewriterConfig_Toggle `protobuf:"varint,9,opt,name=loop_optimization,json=loopOptimization,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"loop_optimization,omitempty"`
	// Function optimizations (default is ON).
	FunctionOptimization RewriterConfig_Toggle `protobuf:"varint,10,opt,name=function_optimization,json=functionOptimization,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"function_optimization,omitempty"`
	// Strips debug-related nodes from the graph (off by default).
	DebugStripper RewriterConfig_Toggle `protobuf:"varint,11,opt,name=debug_stripper,json=debugStripper,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"debug_stripper,omitempty"`
	// If true, don't remove unnecessary ops from the graph
	DisableModelPruning bool `protobuf:"varint,2,opt,name=disable_model_pruning,json=disableModelPruning,proto3" json:"disable_model_pruning,omitempty"`
	// Try to allocate some independent Op outputs contiguously in order to
	// merge or eliminate downstream Ops (off by default).
	ScopedAllocatorOptimization RewriterConfig_Toggle `protobuf:"varint,15,opt,name=scoped_allocator_optimization,json=scopedAllocatorOptimization,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"scoped_allocator_optimization,omitempty"`
	// Force small ops onto the CPU (default is OFF).
	PinToHostOptimization RewriterConfig_Toggle `protobuf:"varint,18,opt,name=pin_to_host_optimization,json=pinToHostOptimization,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"pin_to_host_optimization,omitempty"`
	// Enable the swap of kernel implementations based on the device placement
	// (default is ON).
	ImplementationSelector RewriterConfig_Toggle `protobuf:"varint,22,opt,name=implementation_selector,json=implementationSelector,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"implementation_selector,omitempty"`
	// Optimize data types for CUDA (default is OFF).
	// This will try to use float16 on GPU which is faster.
	// Note that this can change the numerical stability of the graph and may
	// require the use of loss scaling to maintain model convergence.
	AutoMixedPrecision RewriterConfig_Toggle `protobuf:"varint,23,opt,name=auto_mixed_precision,json=autoMixedPrecision,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"auto_mixed_precision,omitempty"`
	// Optimize data types for MKL (default is OFF).
	// This will try to use bfloat16 on CPUs, which is faster.
	// Note that this can change the numerical stability of the graph.
	AutoMixedPrecisionMkl RewriterConfig_Toggle `protobuf:"varint,25,opt,name=auto_mixed_precision_mkl,json=autoMixedPrecisionMkl,proto3,enum=tensorflow.RewriterConfig_Toggle" json:"auto_mixed_precision_mkl,omitempty"`
	// Disable the entire meta optimizer (off by default).
	DisableMetaOptimizer bool `protobuf:"varint,19,opt,name=disable_meta_optimizer,json=disableMetaOptimizer,proto3" json:"disable_meta_optimizer,omitempty"`
	// Controls how many times we run the optimizers in meta optimizer (default
	// is once).
	MetaOptimizerIterations RewriterConfig_NumIterationsType `protobuf:"varint,12,opt,name=meta_optimizer_iterations,json=metaOptimizerIterations,proto3,enum=tensorflow.RewriterConfig_NumIterationsType" json:"meta_optimizer_iterations,omitempty"`
	// The minimum number of nodes in a graph to optimizer. For smaller graphs,
	// optimization is skipped.
	// 0 means the system picks an appropriate number.
	// < 0 means do not skip optimization.
	MinGraphNodes int32 `protobuf:"varint,17,opt,name=min_graph_nodes,json=minGraphNodes,proto3" json:"min_graph_nodes,omitempty"`
	// Configures memory optimization passes through the meta-optimizer. Has no
	// effect on manually requested memory optimization passes in the optimizers
	// field.
	MemoryOptimization RewriterConfig_MemOptType `protobuf:"varint,4,opt,name=memory_optimization,json=memoryOptimization,proto3,enum=tensorflow.RewriterConfig_MemOptType" json:"memory_optimization,omitempty"`
	// A node name scope for node names which are valid outputs of recomputations.
	// Inputs to nodes that match this scope may be recomputed (subject either to
	// manual annotation of those input nodes or to manual annotation and
	// heuristics depending on memory_optimization), but the nodes themselves will
	// not be recomputed. This matches any sub-scopes as well, meaning the scope
	// can appear not just as a top-level scope. For example, if the value is
	// "gradients/", the default, it will match node name "gradients/foo",
	// "foo/gradients/bar", but not "foo_gradients/"
	MemoryOptimizerTargetNodeNameScope string `protobuf:"bytes,6,opt,name=memory_optimizer_target_node_name_scope,json=memoryOptimizerTargetNodeNameScope,proto3" json:"memory_optimizer_target_node_name_scope,omitempty"`
	// Maximum number of milliseconds to spend optimizing a single graph before
	// timing out. If equal to 0 the system picks a default (currently 5 minutes).
	// If less than 0 the optimizer will never time out.
	MetaOptimizerTimeoutMs int64 `protobuf:"varint,20,opt,name=meta_optimizer_timeout_ms,json=metaOptimizerTimeoutMs,proto3" json:"meta_optimizer_timeout_ms,omitempty"`
	// Configures AutoParallel optimization passes either through the
	// meta-optimizer or when manually specified through the optimizers field.
	AutoParallel *AutoParallelOptions `protobuf:"bytes,5,opt,name=auto_parallel,json=autoParallel,proto3" json:"auto_parallel,omitempty"`
	// If true, any optimization pass failing will cause the MetaOptimizer to
	// stop with an error. By default - or when set to false, failing passes are
	// skipped silently.
	FailOnOptimizerErrors bool                    `protobuf:"varint,21,opt,name=fail_on_optimizer_errors,json=failOnOptimizerErrors,proto3" json:"fail_on_optimizer_errors,omitempty"`
	ScopedAllocatorOpts   *ScopedAllocatorOptions `protobuf:"bytes,16,opt,name=scoped_allocator_opts,json=scopedAllocatorOpts,proto3" json:"scoped_allocator_opts,omitempty"`
	// If non-empty, will use this as an alternative way to specify a list of
	// optimizations to turn on and the order of the optimizations (replacing the
	// meta-optimizer).
	//
	// Of the RewriterConfig options, only the AutoParallel configuration options
	// (the auto_parallel field) apply to manually requested optimization passes
	// ("autoparallel"). Memory optimization passes ("memory") invoked here are
	// not configurable (in contrast to memory optimization passes through the
	// meta-optimizer) and act only on manual op annotations.
	//
	// Custom optimizers (see custom_optimizers) that are not part of this
	// schedule will be run after - in the order that they were specified.
	Optimizers []string `protobuf:"bytes,100,rep,name=optimizers,proto3" json:"optimizers,omitempty"`
	// list of CustomGraphOptimizers to apply.
	CustomOptimizers []*RewriterConfig_CustomGraphOptimizer `protobuf:"bytes,200,rep,name=custom_optimizers,json=customOptimizers,proto3" json:"custom_optimizers,omitempty"`
	// VerifierConfig specifying the verifiers to be run after every optimizer.
	InterOptimizerVerifierConfig *VerifierConfig `protobuf:"bytes,300,opt,name=inter_optimizer_verifier_config,json=interOptimizerVerifierConfig,proto3" json:"inter_optimizer_verifier_config,omitempty"`
	// VerifierConfig specifying the verifiers to be run at the end, after all
	// optimizers have run.
	PostOptimizationVerifierConfig *VerifierConfig `protobuf:"bytes,301,opt,name=post_optimization_verifier_config,json=postOptimizationVerifierConfig,proto3" json:"post_optimization_verifier_config,omitempty"`
	XXX_NoUnkeyedLiteral           struct{}        `json:"-"`
	XXX_unrecognized               []byte          `json:"-"`
	XXX_sizecache                  int32           `json:"-"`
}

func (m *RewriterConfig) Reset()         { *m = RewriterConfig{} }
func (m *RewriterConfig) String() string { return proto.CompactTextString(m) }
func (*RewriterConfig) ProtoMessage()    {}
func (*RewriterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd7de60bf190bbb, []int{2}
}

func (m *RewriterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RewriterConfig.Unmarshal(m, b)
}
func (m *RewriterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RewriterConfig.Marshal(b, m, deterministic)
}
func (m *RewriterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewriterConfig.Merge(m, src)
}
func (m *RewriterConfig) XXX_Size() int {
	return xxx_messageInfo_RewriterConfig.Size(m)
}
func (m *RewriterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RewriterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RewriterConfig proto.InternalMessageInfo

func (m *RewriterConfig) GetLayoutOptimizer() RewriterConfig_Toggle {
	if m != nil {
		return m.LayoutOptimizer
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetConstantFolding() RewriterConfig_Toggle {
	if m != nil {
		return m.ConstantFolding
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetShapeOptimization() RewriterConfig_Toggle {
	if m != nil {
		return m.ShapeOptimization
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetRemapping() RewriterConfig_Toggle {
	if m != nil {
		return m.Remapping
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetCommonSubgraphElimination() RewriterConfig_Toggle {
	if m != nil {
		return m.CommonSubgraphElimination
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetArithmeticOptimization() RewriterConfig_Toggle {
	if m != nil {
		return m.ArithmeticOptimization
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetDependencyOptimization() RewriterConfig_Toggle {
	if m != nil {
		return m.DependencyOptimization
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetLoopOptimization() RewriterConfig_Toggle {
	if m != nil {
		return m.LoopOptimization
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetFunctionOptimization() RewriterConfig_Toggle {
	if m != nil {
		return m.FunctionOptimization
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetDebugStripper() RewriterConfig_Toggle {
	if m != nil {
		return m.DebugStripper
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetDisableModelPruning() bool {
	if m != nil {
		return m.DisableModelPruning
	}
	return false
}

func (m *RewriterConfig) GetScopedAllocatorOptimization() RewriterConfig_Toggle {
	if m != nil {
		return m.ScopedAllocatorOptimization
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetPinToHostOptimization() RewriterConfig_Toggle {
	if m != nil {
		return m.PinToHostOptimization
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetImplementationSelector() RewriterConfig_Toggle {
	if m != nil {
		return m.ImplementationSelector
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetAutoMixedPrecision() RewriterConfig_Toggle {
	if m != nil {
		return m.AutoMixedPrecision
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetAutoMixedPrecisionMkl() RewriterConfig_Toggle {
	if m != nil {
		return m.AutoMixedPrecisionMkl
	}
	return RewriterConfig_DEFAULT
}

func (m *RewriterConfig) GetDisableMetaOptimizer() bool {
	if m != nil {
		return m.DisableMetaOptimizer
	}
	return false
}

func (m *RewriterConfig) GetMetaOptimizerIterations() RewriterConfig_NumIterationsType {
	if m != nil {
		return m.MetaOptimizerIterations
	}
	return RewriterConfig_DEFAULT_NUM_ITERS
}

func (m *RewriterConfig) GetMinGraphNodes() int32 {
	if m != nil {
		return m.MinGraphNodes
	}
	return 0
}

func (m *RewriterConfig) GetMemoryOptimization() RewriterConfig_MemOptType {
	if m != nil {
		return m.MemoryOptimization
	}
	return RewriterConfig_DEFAULT_MEM_OPT
}

func (m *RewriterConfig) GetMemoryOptimizerTargetNodeNameScope() string {
	if m != nil {
		return m.MemoryOptimizerTargetNodeNameScope
	}
	return ""
}

func (m *RewriterConfig) GetMetaOptimizerTimeoutMs() int64 {
	if m != nil {
		return m.MetaOptimizerTimeoutMs
	}
	return 0
}

func (m *RewriterConfig) GetAutoParallel() *AutoParallelOptions {
	if m != nil {
		return m.AutoParallel
	}
	return nil
}

func (m *RewriterConfig) GetFailOnOptimizerErrors() bool {
	if m != nil {
		return m.FailOnOptimizerErrors
	}
	return false
}

func (m *RewriterConfig) GetScopedAllocatorOpts() *ScopedAllocatorOptions {
	if m != nil {
		return m.ScopedAllocatorOpts
	}
	return nil
}

func (m *RewriterConfig) GetOptimizers() []string {
	if m != nil {
		return m.Optimizers
	}
	return nil
}

func (m *RewriterConfig) GetCustomOptimizers() []*RewriterConfig_CustomGraphOptimizer {
	if m != nil {
		return m.CustomOptimizers
	}
	return nil
}

func (m *RewriterConfig) GetInterOptimizerVerifierConfig() *VerifierConfig {
	if m != nil {
		return m.InterOptimizerVerifierConfig
	}
	return nil
}

func (m *RewriterConfig) GetPostOptimizationVerifierConfig() *VerifierConfig {
	if m != nil {
		return m.PostOptimizationVerifierConfig
	}
	return nil
}

// Message to describe custom graph optimizer and its parameters
type RewriterConfig_CustomGraphOptimizer struct {
	Name                 string                                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ParameterMap         map[string]*attr_value_go_proto.AttrValue `protobuf:"bytes,2,rep,name=parameter_map,json=parameterMap,proto3" json:"parameter_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *RewriterConfig_CustomGraphOptimizer) Reset()         { *m = RewriterConfig_CustomGraphOptimizer{} }
func (m *RewriterConfig_CustomGraphOptimizer) String() string { return proto.CompactTextString(m) }
func (*RewriterConfig_CustomGraphOptimizer) ProtoMessage()    {}
func (*RewriterConfig_CustomGraphOptimizer) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dd7de60bf190bbb, []int{2, 0}
}

func (m *RewriterConfig_CustomGraphOptimizer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RewriterConfig_CustomGraphOptimizer.Unmarshal(m, b)
}
func (m *RewriterConfig_CustomGraphOptimizer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RewriterConfig_CustomGraphOptimizer.Marshal(b, m, deterministic)
}
func (m *RewriterConfig_CustomGraphOptimizer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewriterConfig_CustomGraphOptimizer.Merge(m, src)
}
func (m *RewriterConfig_CustomGraphOptimizer) XXX_Size() int {
	return xxx_messageInfo_RewriterConfig_CustomGraphOptimizer.Size(m)
}
func (m *RewriterConfig_CustomGraphOptimizer) XXX_DiscardUnknown() {
	xxx_messageInfo_RewriterConfig_CustomGraphOptimizer.DiscardUnknown(m)
}

var xxx_messageInfo_RewriterConfig_CustomGraphOptimizer proto.InternalMessageInfo

func (m *RewriterConfig_CustomGraphOptimizer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RewriterConfig_CustomGraphOptimizer) GetParameterMap() map[string]*attr_value_go_proto.AttrValue {
	if m != nil {
		return m.ParameterMap
	}
	return nil
}

func init() {
	proto.RegisterEnum("tensorflow.RewriterConfig_Toggle", RewriterConfig_Toggle_name, RewriterConfig_Toggle_value)
	proto.RegisterEnum("tensorflow.RewriterConfig_NumIterationsType", RewriterConfig_NumIterationsType_name, RewriterConfig_NumIterationsType_value)
	proto.RegisterEnum("tensorflow.RewriterConfig_MemOptType", RewriterConfig_MemOptType_name, RewriterConfig_MemOptType_value)
	proto.RegisterType((*AutoParallelOptions)(nil), "tensorflow.AutoParallelOptions")
	proto.RegisterType((*ScopedAllocatorOptions)(nil), "tensorflow.ScopedAllocatorOptions")
	proto.RegisterType((*RewriterConfig)(nil), "tensorflow.RewriterConfig")
	proto.RegisterType((*RewriterConfig_CustomGraphOptimizer)(nil), "tensorflow.RewriterConfig.CustomGraphOptimizer")
	proto.RegisterMapType((map[string]*attr_value_go_proto.AttrValue)(nil), "tensorflow.RewriterConfig.CustomGraphOptimizer.ParameterMapEntry")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/rewriter_config.proto", fileDescriptor_1dd7de60bf190bbb)
}

var fileDescriptor_1dd7de60bf190bbb = []byte{
	// 1267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x97, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0xc7, 0x2b, 0xa7, 0x49, 0x9b, 0x93, 0x2f, 0x9b, 0x8e, 0x1d, 0x25, 0xed, 0xda, 0xd4, 0xc0,
	0xb6, 0x60, 0x1b, 0x1c, 0x20, 0xfb, 0x1e, 0x06, 0x0c, 0x6e, 0xea, 0x24, 0x06, 0x62, 0xcb, 0x90,
	0x9d, 0x14, 0x28, 0x30, 0x10, 0xb4, 0x4c, 0x3b, 0x42, 0x44, 0x91, 0xa0, 0xa8, 0x76, 0xd9, 0xcd,
	0x1e, 0x65, 0x97, 0xbb, 0xd9, 0xde, 0x63, 0xcf, 0xb2, 0x27, 0xd8, 0xe5, 0x40, 0xca, 0x1f, 0x92,
	0x9d, 0x75, 0xda, 0x8d, 0x21, 0xf3, 0xf0, 0xfc, 0xf8, 0x3f, 0x47, 0x87, 0x87, 0x14, 0xd4, 0x15,
	0x0d, 0x23, 0x2e, 0x47, 0x01, 0x7f, 0x77, 0xec, 0x71, 0x49, 0x8f, 0x85, 0xe4, 0x8a, 0x0f, 0xe2,
	0xd1, 0xb1, 0xa4, 0xef, 0xa4, 0xaf, 0xa8, 0xc4, 0x1e, 0x0f, 0x47, 0xfe, 0xb8, 0x6e, 0x0c, 0x08,
	0xe6, 0xf3, 0x0f, 0x3e, 0x59, 0xf4, 0x1d, 0x49, 0xc2, 0xe8, 0x3b, 0x2e, 0x6f, 0x8f, 0x89, 0x52,
	0x12, 0xbf, 0x25, 0x41, 0x4c, 0x13, 0xbf, 0x83, 0x7f, 0x5f, 0xe7, 0x2d, 0x95, 0xfe, 0xc8, 0x5f,
	0x58, 0xa7, 0xd6, 0x85, 0x72, 0x23, 0x56, 0xbc, 0x4b, 0x24, 0x09, 0x02, 0x1a, 0x38, 0x42, 0xf9,
	0x3c, 0x8c, 0x50, 0x15, 0xd6, 0x68, 0x48, 0x06, 0x01, 0xb5, 0xad, 0x43, 0xeb, 0xe8, 0xb1, 0x3b,
	0xf9, 0x87, 0x5e, 0xc0, 0x66, 0x18, 0x33, 0x2c, 0xa9, 0x08, 0x7c, 0x8f, 0x44, 0x76, 0xe1, 0xd0,
	0x3a, 0x5a, 0x75, 0x37, 0xc2, 0x98, 0xb9, 0x93, 0xa1, 0xda, 0x97, 0x50, 0xed, 0x79, 0x5c, 0xd0,
	0x61, 0x23, 0x08, 0xb8, 0x47, 0x14, 0x97, 0x53, 0xe8, 0x13, 0x58, 0x4f, 0x30, 0x98, 0x0b, 0xdb,
	0x3a, 0x5c, 0x39, 0x5a, 0x77, 0x1f, 0x27, 0x03, 0x8e, 0xa8, 0xfd, 0x56, 0x85, 0x6d, 0x77, 0x92,
	0x8a, 0x53, 0xa3, 0x10, 0x5d, 0x42, 0x31, 0x20, 0x77, 0x3c, 0x56, 0x98, 0x0b, 0xe5, 0x33, 0xff,
	0x67, 0x2a, 0x8d, 0x9c, 0xed, 0x93, 0x17, 0xa9, 0x30, 0xeb, 0x59, 0xaf, 0x7a, 0x9f, 0x8f, 0xc7,
	0x01, 0x75, 0x77, 0x12, 0x57, 0x67, 0xea, 0xa9, 0x69, 0x1e, 0x0f, 0x23, 0x45, 0x42, 0x85, 0x47,
	0x3c, 0x18, 0xfa, 0xe1, 0xd8, 0x5e, 0xc9, 0x4d, 0x9b, 0xba, 0x9e, 0x25, 0x9e, 0xa8, 0x0b, 0x28,
	0xba, 0x21, 0x82, 0x4e, 0xa5, 0x11, 0x1d, 0xa2, 0xbd, 0x95, 0x97, 0x57, 0x32, 0xce, 0x4e, 0xca,
	0x17, 0xfd, 0x00, 0xeb, 0x92, 0x32, 0x22, 0x84, 0x16, 0xb6, 0x9d, 0x17, 0x34, 0xf7, 0x41, 0x04,
	0x9e, 0x78, 0x9c, 0x31, 0x1e, 0xe2, 0x28, 0x1e, 0x8c, 0x25, 0x11, 0x37, 0x98, 0x06, 0x3e, 0xf3,
	0xc3, 0x44, 0x9b, 0x9d, 0x17, 0xb9, 0x9f, 0x50, 0x7a, 0x13, 0x48, 0x73, 0xce, 0x40, 0x6f, 0x60,
	0x8f, 0x48, 0x5f, 0xdd, 0x30, 0xaa, 0x7c, 0x2f, 0x1b, 0xfa, 0xa3, 0xbc, 0xf8, 0xea, 0x9c, 0x90,
	0x89, 0xff, 0x0d, 0xec, 0x0d, 0xa9, 0xa0, 0xe1, 0x90, 0x86, 0xde, 0x5d, 0x96, 0xfd, 0x38, 0x37,
	0x7b, 0x4e, 0xc8, 0xb0, 0x3b, 0x50, 0x0a, 0x38, 0x17, 0x59, 0xea, 0x7a, 0x5e, 0x6a, 0x51, 0xfb,
	0x66, 0x78, 0xd7, 0x50, 0x19, 0xc5, 0xa1, 0xa7, 0x9f, 0xb3, 0x4c, 0xc8, 0xcb, 0xdc, 0x9d, 0xfa,
	0x67, 0xb8, 0x17, 0xb0, 0x3d, 0xa4, 0x83, 0x78, 0x8c, 0x23, 0x25, 0x7d, 0x21, 0xa8, 0xb4, 0x37,
	0xf2, 0x02, 0xb7, 0x8c, 0x63, 0x6f, 0xe2, 0x87, 0x4e, 0xa0, 0x32, 0xf4, 0x23, 0xb3, 0xd9, 0x18,
	0x1f, 0xd2, 0x00, 0x0b, 0x19, 0x87, 0xba, 0xb2, 0x0a, 0x66, 0x3f, 0x97, 0x27, 0xc6, 0xb6, 0xb6,
	0x75, 0x13, 0x13, 0xa2, 0xf0, 0x41, 0x64, 0x76, 0x2e, 0x26, 0xd3, 0xad, 0x9b, 0x8d, 0x6e, 0x27,
	0xaf, 0x98, 0x27, 0xd1, 0x72, 0x07, 0x48, 0xbd, 0x68, 0x5b, 0xf8, 0x21, 0x56, 0x1c, 0xdf, 0xf0,
	0x48, 0x65, 0x57, 0x40, 0x79, 0x57, 0xa8, 0x08, 0x3f, 0xec, 0xf3, 0x0b, 0x1e, 0xa9, 0xc5, 0x22,
	0xf2, 0x99, 0x08, 0x28, 0xa3, 0xa1, 0x32, 0x23, 0x38, 0xa2, 0x01, 0xf5, 0x14, 0x97, 0x76, 0x35,
	0x77, 0x11, 0x65, 0x09, 0xbd, 0x09, 0x00, 0xf5, 0x60, 0x97, 0xc4, 0x8a, 0x63, 0xe6, 0xff, 0x44,
	0x87, 0x58, 0x48, 0xea, 0xf9, 0x91, 0xd6, 0xbc, 0x97, 0x17, 0x8c, 0xb4, 0x7b, 0x5b, 0x7b, 0x77,
	0xa7, 0xce, 0x3a, 0x19, 0xf7, 0x41, 0x31, 0xbb, 0x0d, 0xec, 0xfd, 0xdc, 0xc9, 0x58, 0x06, 0xb7,
	0x6f, 0x03, 0xf4, 0x05, 0x54, 0x67, 0x35, 0x40, 0x15, 0x49, 0x75, 0xd1, 0xb2, 0x29, 0x82, 0xdd,
	0x69, 0x11, 0x50, 0x45, 0xe6, 0x7d, 0xf2, 0x06, 0xf6, 0xb3, 0xb3, 0xb1, 0x5e, 0xd0, 0xa4, 0x22,
	0xb2, 0x37, 0x8d, 0xa4, 0xcf, 0xde, 0x23, 0xa9, 0x13, 0xb3, 0xd6, 0x6c, 0x7e, 0xff, 0x4e, 0x50,
	0x77, 0x8f, 0xa5, 0xf9, 0x73, 0x23, 0xfa, 0x08, 0x76, 0x98, 0x1f, 0xe2, 0xa4, 0x55, 0x85, 0x7c,
	0x48, 0x23, 0xbb, 0x64, 0xce, 0x93, 0x2d, 0xe6, 0x87, 0xe7, 0x7a, 0xb4, 0xa3, 0x07, 0xd1, 0x35,
	0x94, 0x19, 0x65, 0x5c, 0x2e, 0x74, 0x85, 0x87, 0x46, 0xcb, 0x87, 0xef, 0xd1, 0xd2, 0xa6, 0xcc,
	0x11, 0xca, 0x88, 0x40, 0x09, 0x21, 0x53, 0x2c, 0x3d, 0xf8, 0x38, 0xcb, 0xa5, 0x12, 0x2b, 0x22,
	0xc7, 0x54, 0x19, 0x35, 0x38, 0x24, 0x8c, 0x62, 0x53, 0xc9, 0xf6, 0xda, 0xa1, 0x75, 0xb4, 0xee,
	0xd6, 0x32, 0x10, 0x2a, 0xfb, 0x66, 0xb2, 0x16, 0xd9, 0x21, 0x8c, 0x9a, 0x53, 0x0f, 0x7d, 0xbb,
	0x94, 0x3e, 0xe5, 0x33, 0xaa, 0x0f, 0x31, 0x16, 0xd9, 0xbb, 0x87, 0xd6, 0xd1, 0x8a, 0x5b, 0xcd,
	0x24, 0xa4, 0x9f, 0x98, 0xdb, 0x11, 0x7a, 0x05, 0x5b, 0xa6, 0x16, 0xc4, 0xe4, 0x30, 0xb6, 0x57,
	0x0f, 0xad, 0xa3, 0x8d, 0x93, 0xe7, 0xe9, 0x08, 0xef, 0x39, 0xac, 0xdd, 0x4d, 0x92, 0x1a, 0x44,
	0x5f, 0x83, 0x3d, 0x22, 0x7e, 0x80, 0xe7, 0xad, 0x89, 0x4a, 0x4c, 0xa5, 0xe4, 0x32, 0xb2, 0x2b,
	0xe6, 0xbd, 0x57, 0xb4, 0xdd, 0x09, 0x67, 0x0a, 0x9a, 0xc6, 0xa8, 0x9b, 0xda, 0x7d, 0xdb, 0x3f,
	0xb2, 0x8b, 0x46, 0x46, 0x2d, 0x2d, 0xe3, 0xfe, 0x13, 0xde, 0x2d, 0x2f, 0xef, 0xfb, 0x08, 0x3d,
	0x03, 0x98, 0x09, 0x89, 0xec, 0xa1, 0x39, 0xf7, 0x53, 0x23, 0xe8, 0x47, 0x28, 0x79, 0x71, 0xa4,
	0x38, 0xc3, 0xa9, 0x69, 0x7f, 0xea, 0xfb, 0xc1, 0xc6, 0xc9, 0xf1, 0x7b, 0xde, 0xee, 0xa9, 0x71,
	0x32, 0x75, 0x32, 0x0b, 0xc5, 0x2d, 0x26, 0x28, 0x67, 0x8e, 0x1f, 0xc0, 0x73, 0x3f, 0xd4, 0xf7,
	0xab, 0x79, 0x36, 0x16, 0xae, 0x42, 0xf6, 0xef, 0x05, 0x13, 0x61, 0xfa, 0xf2, 0x54, 0xbf, 0x9e,
	0xcc, 0x49, 0x16, 0x73, 0x9f, 0x1a, 0xc6, 0x0c, 0x9b, 0xb5, 0xa2, 0x11, 0xbc, 0x10, 0x8b, 0xbd,
	0x6c, 0x69, 0x95, 0x3f, 0xfe, 0x7b, 0x95, 0x67, 0x62, 0xa1, 0xa1, 0x65, 0xed, 0x07, 0x7f, 0x59,
	0xb0, 0x7b, 0x5f, 0xd8, 0x08, 0xc1, 0x43, 0x5d, 0xad, 0xe6, 0x7a, 0xb4, 0xee, 0x9a, 0x67, 0x34,
	0x82, 0x2d, 0x5d, 0x49, 0x8c, 0xea, 0xe0, 0x19, 0x11, 0x76, 0xc1, 0xa4, 0xb4, 0xf1, 0x3f, 0x53,
	0x5a, 0xef, 0x4e, 0x21, 0x6d, 0x22, 0x9a, 0xa1, 0x92, 0x77, 0xee, 0xa6, 0x48, 0x0d, 0x1d, 0x5c,
	0x43, 0x69, 0x69, 0x0a, 0x2a, 0xc2, 0xca, 0x2d, 0xbd, 0x9b, 0xe8, 0xd1, 0x8f, 0xe8, 0x53, 0x58,
	0x35, 0x17, 0x55, 0x3b, 0x49, 0x43, 0x25, 0x53, 0xd5, 0x4a, 0xc9, 0x6b, 0x6d, 0x74, 0x93, 0x39,
	0xdf, 0x15, 0xbe, 0xb1, 0x6a, 0x5f, 0xc1, 0x5a, 0xd2, 0xdf, 0xd0, 0x06, 0x3c, 0x7a, 0xd5, 0x3c,
	0x6b, 0x5c, 0x5d, 0xf6, 0x8b, 0x0f, 0xd0, 0x1a, 0x14, 0x9c, 0x4e, 0xd1, 0x42, 0x8f, 0x60, 0xc5,
	0x39, 0x3b, 0x2b, 0x16, 0xd0, 0x36, 0x40, 0xe3, 0xfc, 0xdc, 0x6d, 0xf6, 0x7a, 0xad, 0xeb, 0x66,
	0x71, 0xa5, 0xf6, 0x3d, 0x94, 0x96, 0x9a, 0x10, 0xaa, 0x40, 0x69, 0x82, 0xc0, 0x9d, 0xab, 0x36,
	0x6e, 0xf5, 0x9b, 0x6e, 0xaf, 0xf8, 0xc0, 0x40, 0x3a, 0xcd, 0x84, 0xd6, 0x7f, 0xed, 0x14, 0x0b,
	0xb5, 0x5f, 0x2d, 0x80, 0x79, 0xdf, 0x40, 0x65, 0xd8, 0x99, 0xfa, 0xb5, 0x9b, 0x6d, 0xec, 0x74,
	0xb5, 0x84, 0x6d, 0x80, 0x8e, 0x33, 0xfb, 0x6f, 0x21, 0x80, 0xb5, 0x76, 0xa3, 0x73, 0xd5, 0xb8,
	0x2c, 0x16, 0xd0, 0x1e, 0x94, 0x7b, 0xaf, 0x1b, 0xdd, 0x6e, 0xab, 0x73, 0x8e, 0x2f, 0x9a, 0x57,
	0x6e, 0xab, 0xd7, 0x6f, 0x9d, 0xf6, 0x8a, 0x0f, 0xd1, 0x53, 0xb0, 0xdd, 0xe6, 0xa9, 0xd3, 0xee,
	0x5e, 0xf5, 0x1b, 0xfd, 0x96, 0xd3, 0x49, 0x5b, 0x57, 0xd1, 0x3e, 0x54, 0x7a, 0xa7, 0x17, 0xcd,
	0x57, 0x57, 0x97, 0x0b, 0x8e, 0x6b, 0x7a, 0xb5, 0xd4, 0xff, 0x95, 0x97, 0xbf, 0x80, 0xcd, 0xe5,
	0x38, 0x9d, 0xbe, 0xd9, 0xf7, 0xc0, 0xcb, 0xdd, 0xec, 0x0b, 0xed, 0xea, 0x3b, 0x7e, 0xd4, 0xb5,
	0xde, 0x5c, 0x8c, 0x7d, 0x75, 0x13, 0x0f, 0xea, 0x1e, 0x67, 0xc7, 0xa9, 0x2f, 0x84, 0xfb, 0x1f,
	0xc7, 0x3c, 0xf9, 0x74, 0xd0, 0x3f, 0xd8, 0x7c, 0x26, 0x44, 0x78, 0xcc, 0x93, 0xa7, 0xbf, 0x2d,
	0x6b, 0xb0, 0x66, 0x9e, 0x3e, 0xff, 0x27, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xf1, 0x10, 0xde, 0xd4,
	0x0c, 0x00, 0x00,
}
