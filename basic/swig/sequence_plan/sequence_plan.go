/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (https://www.swig.org).
 * Version 4.2.0
 *
 * Do not make changes to this file unless you know what you are doing - modify
 * the SWIG interface file instead.
 * ----------------------------------------------------------------------------- */

// source: sequence_plan.i

package sequence_plan

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stddef.h>
#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef _gostring_ swig_type_1;
extern void _wrap_Swig_free_sequence_plan_4d4f36d573574039(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_sequence_plan_4d4f36d573574039(swig_intgo arg1);
extern swig_intgo _wrap_sequence_plan_sequence_plan_4d4f36d573574039(swig_type_1 arg1);
#undef intgo
#cgo LDFLAGS: -L/lib/x86_64-linux-gnu -larrow -lparquet -larrow_acero
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"

type _ unsafe.Pointer

var Swig_escape_always_false bool
var Swig_escape_val interface{}

type _swig_fnptr *byte
type _swig_memberptr *byte

func getSwigcptr(v interface{ Swigcptr() uintptr }) uintptr {
	if v == nil {
		return 0
	}
	return v.Swigcptr()
}

type _ sync.Mutex

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_sequence_plan_4d4f36d573574039(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_sequence_plan_4d4f36d573574039(C.swig_intgo(_swig_i_0)))
	return swig_r
}

func Sequence_plan(arg1 string) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_sequence_plan_sequence_plan_4d4f36d573574039(*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_0))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}
