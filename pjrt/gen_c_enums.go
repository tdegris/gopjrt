package pjrt

/***** File generated by ./cmd/pjrt_codegen, don't edit it directly. *****/

// PJRT_Extension_Type is mapping of the corresponded C enum defined in pjrt_c_api.h.
type PJRT_Extension_Type int

const (
	// PJRT_Extension_Type_Gpu_Custom_Call is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_Extension_Type_Gpu_Custom_Call PJRT_Extension_Type = 0

	// PJRT_Extension_Type_Profiler is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_Extension_Type_Profiler PJRT_Extension_Type = 1

	// PJRT_Extension_Type_Custom_Partitioner is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_Extension_Type_Custom_Partitioner PJRT_Extension_Type = 2

	// PJRT_Extension_Type_Stream is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_Extension_Type_Stream PJRT_Extension_Type = 3

	// PJRT_Extension_Type_Layouts is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_Extension_Type_Layouts PJRT_Extension_Type = 4

	// PJRT_Extension_Type_FFI is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_Extension_Type_FFI PJRT_Extension_Type = 5
)

// PJRT_Error_Code is mapping of the corresponded C enum defined in pjrt_c_api.h.
type PJRT_Error_Code int

const (
	// PJRT_Error_Code_CANCELLED is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_Error_Code_CANCELLED PJRT_Error_Code = 1

	// PJRT_Error_Code_UNKNOWN is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_Error_Code_UNKNOWN PJRT_Error_Code = 2

	// PJRT_Error_Code_INVALID_ARGUMENT is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_Error_Code_INVALID_ARGUMENT PJRT_Error_Code = 3

	// PJRT_Error_Code_DEADLINE_EXCEEDED is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_Error_Code_DEADLINE_EXCEEDED PJRT_Error_Code = 4

	// PJRT_Error_Code_NOT_FOUND is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_Error_Code_NOT_FOUND PJRT_Error_Code = 5

	// PJRT_Error_Code_ALREADY_EXISTS is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_Error_Code_ALREADY_EXISTS PJRT_Error_Code = 6

	// PJRT_Error_Code_PERMISSION_DENIED is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_Error_Code_PERMISSION_DENIED PJRT_Error_Code = 7

	// PJRT_Error_Code_RESOURCE_EXHAUSTED is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_Error_Code_RESOURCE_EXHAUSTED PJRT_Error_Code = 8

	// PJRT_Error_Code_FAILED_PRECONDITION is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_Error_Code_FAILED_PRECONDITION PJRT_Error_Code = 9

	// PJRT_Error_Code_ABORTED is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_Error_Code_ABORTED PJRT_Error_Code = 10

	// PJRT_Error_Code_OUT_OF_RANGE is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_Error_Code_OUT_OF_RANGE PJRT_Error_Code = 11

	// PJRT_Error_Code_UNIMPLEMENTED is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_Error_Code_UNIMPLEMENTED PJRT_Error_Code = 12

	// PJRT_Error_Code_INTERNAL is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_Error_Code_INTERNAL PJRT_Error_Code = 13

	// PJRT_Error_Code_UNAVAILABLE is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_Error_Code_UNAVAILABLE PJRT_Error_Code = 14

	// PJRT_Error_Code_DATA_LOSS is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_Error_Code_DATA_LOSS PJRT_Error_Code = 15

	// PJRT_Error_Code_UNAUTHENTICATED is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_Error_Code_UNAUTHENTICATED PJRT_Error_Code = 16
)

// PJRT_NamedValue_Type is mapping of the corresponded C enum defined in pjrt_c_api.h.
type PJRT_NamedValue_Type int

const (
	// PJRT_NamedValue_kString is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_NamedValue_kString PJRT_NamedValue_Type = 0

	// PJRT_NamedValue_kInt64 is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_NamedValue_kInt64 PJRT_NamedValue_Type = 1

	// PJRT_NamedValue_kInt64List is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_NamedValue_kInt64List PJRT_NamedValue_Type = 2

	// PJRT_NamedValue_kFloat is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_NamedValue_kFloat PJRT_NamedValue_Type = 3

	// PJRT_NamedValue_kBool is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_NamedValue_kBool PJRT_NamedValue_Type = 4
)

// PJRT_HostBufferSemantics is mapping of the corresponded C enum defined in pjrt_c_api.h.
type PJRT_HostBufferSemantics int

const (
	// PJRT_HostBufferSemantics_kImmutableOnlyDuringCall is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	// The runtime may not hold references to `data` after the call to
	// `PJRT_Client_BufferFromHostBuffer` completes. The caller promises that
	// `data` is immutable and will not be freed only for the duration of the
	// PJRT_Client_BufferFromHostBuffer call.
	PJRT_HostBufferSemantics_kImmutableOnlyDuringCall PJRT_HostBufferSemantics = 0

	// PJRT_HostBufferSemantics_kImmutableUntilTransferCompletes is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	// The runtime may hold onto `data` after the call to
	// `PJRT_Client_BufferFromHostBuffer`
	// returns while the runtime completes a transfer to the device. The caller
	// promises not to mutate or free `data` until the transfer completes, at
	// which point `done_with_host_buffer` will be triggered.
	PJRT_HostBufferSemantics_kImmutableUntilTransferCompletes PJRT_HostBufferSemantics = 1

	// PJRT_HostBufferSemantics_kImmutableZeroCopy is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	// The PjRtBuffer may alias `data` internally and the runtime may use the
	// `data` contents as long as the buffer is alive. The runtime promises not
	// to mutate contents of the buffer (i.e. it will not use it for aliased
	// output buffers). The caller promises to keep `data` alive and not to mutate
	// its contents as long as the buffer is alive; to notify the caller that the
	// buffer may be freed, the runtime will call `done_with_host_buffer` when the
	// PjRtBuffer is freed.
	PJRT_HostBufferSemantics_kImmutableZeroCopy PJRT_HostBufferSemantics = 2

	// PJRT_HostBufferSemantics_kMutableZeroCopy is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	// The PjRtBuffer may alias `data` internally and the runtime may use the
	// `data` contents as long as the buffer is alive. The runtime is allowed
	// to mutate contents of the buffer (i.e. use it for aliased output
	// buffers). The caller promises to keep `data` alive and not to mutate its
	// contents as long as the buffer is alive (otherwise it could be a data
	// race with the runtime); to notify the caller that the buffer may be
	// freed, the runtime will call `on_done_with_host_buffer` when the
	// PjRtBuffer is freed. On non-CPU platforms this acts identically to
	// kImmutableUntilTransferCompletes.
	PJRT_HostBufferSemantics_kMutableZeroCopy PJRT_HostBufferSemantics = 3
)

// PJRT_Buffer_MemoryLayout_Type is mapping of the corresponded C enum defined in pjrt_c_api.h.
type PJRT_Buffer_MemoryLayout_Type int

const (
	// PJRT_Buffer_MemoryLayout_Type_Tiled is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_Buffer_MemoryLayout_Type_Tiled PJRT_Buffer_MemoryLayout_Type = 0

	// PJRT_Buffer_MemoryLayout_Type_Strides is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h.
	PJRT_Buffer_MemoryLayout_Type_Strides PJRT_Buffer_MemoryLayout_Type = 1
)
