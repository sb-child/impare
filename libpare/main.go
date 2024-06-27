package main

/*
// Generated by rust2go. Please DO NOT edit this C part manually.

#include <stdarg.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

typedef struct RemoveTaskRequestRef {
  uint64_t id;
} RemoveTaskRequestRef;

typedef struct RemoveTaskResponseRef {
  bool success;
} RemoveTaskResponseRef;

typedef struct CreateTaskResponseRef {
  uint64_t id;
} CreateTaskResponseRef;

typedef struct ListRef {
  const void *ptr;
  uintptr_t len;
} ListRef;

typedef struct StringRef {
  const uint8_t *ptr;
  uintptr_t len;
} StringRef;

typedef struct CreateTaskRequestRef {
  uintptr_t data_shards;
  uintptr_t parity_shards;
} CreateTaskRequestRef;

// hack from: https://stackoverflow.com/a/69904977
__attribute__((weak))
inline void SysApi_create_task_cb(const void *f_ptr, struct CreateTaskResponseRef resp, const void *slot) {
((void (*)(struct CreateTaskResponseRef, const void*))f_ptr)(resp, slot);
}

// hack from: https://stackoverflow.com/a/69904977
__attribute__((weak))
inline void SysApi_remove_task_cb(const void *f_ptr, struct RemoveTaskResponseRef resp, const void *slot) {
((void (*)(struct RemoveTaskResponseRef, const void*))f_ptr)(resp, slot);
}
*/
import "C"
import (
	"runtime"
	"unsafe"
)

var SysApiImpl SysApi

type SysApi interface {
	create_task(req CreateTaskRequest) CreateTaskResponse
	remove_task(req RemoveTaskRequest) RemoveTaskResponse
}

//export CSysApi_create_task
func CSysApi_create_task(req C.CreateTaskRequestRef, slot *C.void, cb *C.void) {
	resp := SysApiImpl.create_task(newCreateTaskRequest(req))
	resp_ref, buffer := cvt_ref(cntCreateTaskResponse, refCreateTaskResponse)(&resp)
	C.SysApi_create_task_cb(unsafe.Pointer(cb), resp_ref, unsafe.Pointer(slot))
	runtime.KeepAlive(resp)
	runtime.KeepAlive(buffer)
}

//export CSysApi_remove_task
func CSysApi_remove_task(req C.RemoveTaskRequestRef, slot *C.void, cb *C.void) {
	resp := SysApiImpl.remove_task(newRemoveTaskRequest(req))
	resp_ref, buffer := cvt_ref(cntRemoveTaskResponse, refRemoveTaskResponse)(&resp)
	C.SysApi_remove_task_cb(unsafe.Pointer(cb), resp_ref, unsafe.Pointer(slot))
	runtime.KeepAlive(resp)
	runtime.KeepAlive(buffer)
}

func newString(s_ref C.StringRef) string {
	return unsafe.String((*byte)(unsafe.Pointer(s_ref.ptr)), s_ref.len)
}
func refString(s *string, _buffer *[]byte) C.StringRef {
	return C.StringRef{
		ptr: (*C.uint8_t)(unsafe.StringData(*s)),
		len: C.uintptr_t(len(*s)),
	}
}

func cntString(s *string, cnt *uint) [0]C.StringRef { return [0]C.StringRef{} }
func new_list_mapper[T1, T2 any](f func(T1) T2) func(C.ListRef) []T2 {
	return func(x C.ListRef) []T2 {
		input := unsafe.Slice((*T1)(unsafe.Pointer(x.ptr)), x.len)
		output := make([]T2, len(input))
		for i, v := range input {
			output[i] = f(v)
		}
		return output
	}
}
func new_list_mapper_primitive[T1, T2 any](f func(T1) T2) func(C.ListRef) []T2 {
	return func(x C.ListRef) []T2 {
		return unsafe.Slice((*T2)(unsafe.Pointer(x.ptr)), x.len)
	}
}

// only handle non-primitive type T
func cnt_list_mapper[T, R any](f func(s *T, cnt *uint) [0]R) func(s *[]T, cnt *uint) [0]C.ListRef {
	return func(s *[]T, cnt *uint) [0]C.ListRef {
		for _, v := range *s {
			f(&v, cnt)
		}
		*cnt += uint(len(*s)) * size_of[R]()
		return [0]C.ListRef{}
	}
}

// only handle primitive type T
func cnt_list_mapper_primitive[T, R any](f func(s *T, cnt *uint) [0]R) func(s *[]T, cnt *uint) [0]C.ListRef {
	return func(s *[]T, cnt *uint) [0]C.ListRef { return [0]C.ListRef{} }
}

// only handle non-primitive type T
func ref_list_mapper[T, R any](f func(s *T, buffer *[]byte) R) func(s *[]T, buffer *[]byte) C.ListRef {
	return func(s *[]T, buffer *[]byte) C.ListRef {
		if len(*buffer) == 0 {
			return C.ListRef{
				ptr: unsafe.Pointer(nil),
				len: C.uintptr_t(len(*s)),
			}
		}
		ret := C.ListRef{
			ptr: unsafe.Pointer(&(*buffer)[0]),
			len: C.uintptr_t(len(*s)),
		}
		children_bytes := int(size_of[R]()) * len(*s)
		children := (*buffer)[:children_bytes]
		*buffer = (*buffer)[children_bytes:]
		for _, v := range *s {
			child := f(&v, buffer)
			len := unsafe.Sizeof(child)
			copy(children, unsafe.Slice((*byte)(unsafe.Pointer(&child)), len))
			children = children[len:]
		}
		return ret
	}
}

// only handle primitive type T
func ref_list_mapper_primitive[T, R any](f func(s *T, buffer *[]byte) R) func(s *[]T, buffer *[]byte) C.ListRef {
	return func(s *[]T, buffer *[]byte) C.ListRef {
		if len(*s) == 0 {
			return C.ListRef{
				ptr: unsafe.Pointer(nil),
				len: C.uintptr_t(0),
			}
		}
		return C.ListRef{
			ptr: unsafe.Pointer(&(*s)[0]),
			len: C.uintptr_t(len(*s)),
		}
	}
}
func size_of[T any]() uint {
	var t T
	return uint(unsafe.Sizeof(t))
}
func cvt_ref[R, CR any](cnt_f func(s *R, cnt *uint) [0]CR, ref_f func(p *R, buffer *[]byte) CR) func(p *R) (CR, []byte) {
	return func(p *R) (CR, []byte) {
		var cnt uint
		cnt_f(p, &cnt)
		buffer := make([]byte, cnt)
		return ref_f(p, &buffer), buffer
	}
}

func newC_uint8_t(n C.uint8_t) uint8    { return uint8(n) }
func newC_uint16_t(n C.uint16_t) uint16 { return uint16(n) }
func newC_uint32_t(n C.uint32_t) uint32 { return uint32(n) }
func newC_uint64_t(n C.uint64_t) uint64 { return uint64(n) }
func newC_int8_t(n C.int8_t) int8       { return int8(n) }
func newC_int16_t(n C.int16_t) int16    { return int16(n) }
func newC_int32_t(n C.int32_t) int32    { return int32(n) }
func newC_int64_t(n C.int64_t) int64    { return int64(n) }
func newC_bool(n C.bool) bool           { return bool(n) }
func newC_uintptr_t(n C.uintptr_t) uint { return uint(n) }
func newC_intptr_t(n C.intptr_t) int    { return int(n) }
func newC_float(n C.float) float32      { return float32(n) }
func newC_double(n C.double) float64    { return float64(n) }

func cntC_uint8_t(s *uint8, cnt *uint) [0]C.uint8_t    { return [0]C.uint8_t{} }
func cntC_uint16_t(s *uint16, cnt *uint) [0]C.uint16_t { return [0]C.uint16_t{} }
func cntC_uint32_t(s *uint32, cnt *uint) [0]C.uint32_t { return [0]C.uint32_t{} }
func cntC_uint64_t(s *uint64, cnt *uint) [0]C.uint64_t { return [0]C.uint64_t{} }
func cntC_int8_t(s *int8, cnt *uint) [0]C.int8_t       { return [0]C.int8_t{} }
func cntC_int16_t(s *int16, cnt *uint) [0]C.int16_t    { return [0]C.int16_t{} }
func cntC_int32_t(s *int32, cnt *uint) [0]C.int32_t    { return [0]C.int32_t{} }
func cntC_int64_t(s *int64, cnt *uint) [0]C.int64_t    { return [0]C.int64_t{} }
func cntC_bool(s *bool, cnt *uint) [0]C.bool           { return [0]C.bool{} }
func cntC_uintptr_t(s *uint, cnt *uint) [0]C.uintptr_t { return [0]C.uintptr_t{} }
func cntC_intptr_t(s *int, cnt *uint) [0]C.intptr_t    { return [0]C.intptr_t{} }
func cntC_float(s *float32, cnt *uint) [0]C.float      { return [0]C.float{} }
func cntC_double(s *float64, cnt *uint) [0]C.double    { return [0]C.double{} }

func refC_uint8_t(p *uint8, buffer *[]byte) C.uint8_t    { return C.uint8_t(*p) }
func refC_uint16_t(p *uint16, buffer *[]byte) C.uint16_t { return C.uint16_t(*p) }
func refC_uint32_t(p *uint32, buffer *[]byte) C.uint32_t { return C.uint32_t(*p) }
func refC_uint64_t(p *uint64, buffer *[]byte) C.uint64_t { return C.uint64_t(*p) }
func refC_int8_t(p *int8, buffer *[]byte) C.int8_t       { return C.int8_t(*p) }
func refC_int16_t(p *int16, buffer *[]byte) C.int16_t    { return C.int16_t(*p) }
func refC_int32_t(p *int32, buffer *[]byte) C.int32_t    { return C.int32_t(*p) }
func refC_int64_t(p *int64, buffer *[]byte) C.int64_t    { return C.int64_t(*p) }
func refC_bool(p *bool, buffer *[]byte) C.bool           { return C.bool(*p) }
func refC_uintptr_t(p *uint, buffer *[]byte) C.uintptr_t { return C.uintptr_t(*p) }
func refC_intptr_t(p *int, buffer *[]byte) C.intptr_t    { return C.intptr_t(*p) }
func refC_float(p *float32, buffer *[]byte) C.float      { return C.float(*p) }
func refC_double(p *float64, buffer *[]byte) C.double    { return C.double(*p) }

type CreateTaskRequest struct {
	data_shards   uint
	parity_shards uint
}

func newCreateTaskRequest(p C.CreateTaskRequestRef) CreateTaskRequest {
	return CreateTaskRequest{
		data_shards:   newC_uintptr_t(p.data_shards),
		parity_shards: newC_uintptr_t(p.parity_shards),
	}
}
func cntCreateTaskRequest(s *CreateTaskRequest, cnt *uint) [0]C.CreateTaskRequestRef {
	return [0]C.CreateTaskRequestRef{}
}
func refCreateTaskRequest(p *CreateTaskRequest, buffer *[]byte) C.CreateTaskRequestRef {
	return C.CreateTaskRequestRef{
		data_shards:   refC_uintptr_t(&p.data_shards, buffer),
		parity_shards: refC_uintptr_t(&p.parity_shards, buffer),
	}
}

type CreateTaskResponse struct {
	id uint64
}

func newCreateTaskResponse(p C.CreateTaskResponseRef) CreateTaskResponse {
	return CreateTaskResponse{
		id: newC_uint64_t(p.id),
	}
}
func cntCreateTaskResponse(s *CreateTaskResponse, cnt *uint) [0]C.CreateTaskResponseRef {
	return [0]C.CreateTaskResponseRef{}
}
func refCreateTaskResponse(p *CreateTaskResponse, buffer *[]byte) C.CreateTaskResponseRef {
	return C.CreateTaskResponseRef{
		id: refC_uint64_t(&p.id, buffer),
	}
}

type RemoveTaskRequest struct {
	id uint64
}

func newRemoveTaskRequest(p C.RemoveTaskRequestRef) RemoveTaskRequest {
	return RemoveTaskRequest{
		id: newC_uint64_t(p.id),
	}
}
func cntRemoveTaskRequest(s *RemoveTaskRequest, cnt *uint) [0]C.RemoveTaskRequestRef {
	return [0]C.RemoveTaskRequestRef{}
}
func refRemoveTaskRequest(p *RemoveTaskRequest, buffer *[]byte) C.RemoveTaskRequestRef {
	return C.RemoveTaskRequestRef{
		id: refC_uint64_t(&p.id, buffer),
	}
}

type RemoveTaskResponse struct {
	success bool
}

func newRemoveTaskResponse(p C.RemoveTaskResponseRef) RemoveTaskResponse {
	return RemoveTaskResponse{
		success: newC_bool(p.success),
	}
}
func cntRemoveTaskResponse(s *RemoveTaskResponse, cnt *uint) [0]C.RemoveTaskResponseRef {
	return [0]C.RemoveTaskResponseRef{}
}
func refRemoveTaskResponse(p *RemoveTaskResponse, buffer *[]byte) C.RemoveTaskResponseRef {
	return C.RemoveTaskResponseRef{
		success: refC_bool(&p.success, buffer),
	}
}
func main() {}
