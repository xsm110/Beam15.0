// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by starcgen. DO NOT EDIT.
// File: stats.shims.go

package stats

import (
	"reflect"

	// Library imports
	"github.com/apache/beam/sdks/go/pkg/beam/core/runtime"
	"github.com/apache/beam/sdks/go/pkg/beam/core/typex"
	"github.com/apache/beam/sdks/go/pkg/beam/core/util/reflectx"
)

func init() {
	runtime.RegisterFunction(mapFn)
	runtime.RegisterFunction(maxFloat32Fn)
	runtime.RegisterFunction(maxFloat64Fn)
	runtime.RegisterFunction(maxInt16Fn)
	runtime.RegisterFunction(maxInt32Fn)
	runtime.RegisterFunction(maxInt64Fn)
	runtime.RegisterFunction(maxInt8Fn)
	runtime.RegisterFunction(maxIntFn)
	runtime.RegisterFunction(maxUint16Fn)
	runtime.RegisterFunction(maxUint32Fn)
	runtime.RegisterFunction(maxUint64Fn)
	runtime.RegisterFunction(maxUint8Fn)
	runtime.RegisterFunction(maxUintFn)
	runtime.RegisterFunction(minFloat32Fn)
	runtime.RegisterFunction(minFloat64Fn)
	runtime.RegisterFunction(minInt16Fn)
	runtime.RegisterFunction(minInt32Fn)
	runtime.RegisterFunction(minInt64Fn)
	runtime.RegisterFunction(minInt8Fn)
	runtime.RegisterFunction(minIntFn)
	runtime.RegisterFunction(minUint16Fn)
	runtime.RegisterFunction(minUint32Fn)
	runtime.RegisterFunction(minUint64Fn)
	runtime.RegisterFunction(minUint8Fn)
	runtime.RegisterFunction(minUintFn)
	runtime.RegisterFunction(sumFloat32Fn)
	runtime.RegisterFunction(sumFloat64Fn)
	runtime.RegisterFunction(sumInt16Fn)
	runtime.RegisterFunction(sumInt32Fn)
	runtime.RegisterFunction(sumInt64Fn)
	runtime.RegisterFunction(sumInt8Fn)
	runtime.RegisterFunction(sumIntFn)
	runtime.RegisterFunction(sumUint16Fn)
	runtime.RegisterFunction(sumUint32Fn)
	runtime.RegisterFunction(sumUint64Fn)
	runtime.RegisterFunction(sumUint8Fn)
	runtime.RegisterFunction(sumUintFn)
	runtime.RegisterType(reflect.TypeOf((*meanAccum)(nil)).Elem())
	runtime.RegisterType(reflect.TypeOf((*meanFn)(nil)).Elem())
	runtime.RegisterType(reflect.TypeOf((*typex.T)(nil)).Elem())
	reflectx.RegisterStructWrapper(reflect.TypeOf((*meanFn)(nil)).Elem(), wrapMakerMeanFn)
	reflectx.RegisterFunc(reflect.TypeOf((*func(float32,float32) (float32))(nil)).Elem(), funcMakerFloat32Float32ГFloat32)
	reflectx.RegisterFunc(reflect.TypeOf((*func(float64,float64) (float64))(nil)).Elem(), funcMakerFloat64Float64ГFloat64)
	reflectx.RegisterFunc(reflect.TypeOf((*func(int16,int16) (int16))(nil)).Elem(), funcMakerInt16Int16ГInt16)
	reflectx.RegisterFunc(reflect.TypeOf((*func(int32,int32) (int32))(nil)).Elem(), funcMakerInt32Int32ГInt32)
	reflectx.RegisterFunc(reflect.TypeOf((*func(int64,int64) (int64))(nil)).Elem(), funcMakerInt64Int64ГInt64)
	reflectx.RegisterFunc(reflect.TypeOf((*func(int8,int8) (int8))(nil)).Elem(), funcMakerInt8Int8ГInt8)
	reflectx.RegisterFunc(reflect.TypeOf((*func(int,int) (int))(nil)).Elem(), funcMakerIntIntГInt)
	reflectx.RegisterFunc(reflect.TypeOf((*func(meanAccum,typex.T) (meanAccum))(nil)).Elem(), funcMakerMeanAccumTypex۰TГMeanAccum)
	reflectx.RegisterFunc(reflect.TypeOf((*func(meanAccum) (float64))(nil)).Elem(), funcMakerMeanAccumГFloat64)
	reflectx.RegisterFunc(reflect.TypeOf((*func([]meanAccum) (meanAccum))(nil)).Elem(), funcMakerSliceofMeanAccumГMeanAccum)
	reflectx.RegisterFunc(reflect.TypeOf((*func(typex.T) (typex.T,int))(nil)).Elem(), funcMakerTypex۰TГTypex۰TInt)
	reflectx.RegisterFunc(reflect.TypeOf((*func(uint16,uint16) (uint16))(nil)).Elem(), funcMakerUint16Uint16ГUint16)
	reflectx.RegisterFunc(reflect.TypeOf((*func(uint32,uint32) (uint32))(nil)).Elem(), funcMakerUint32Uint32ГUint32)
	reflectx.RegisterFunc(reflect.TypeOf((*func(uint64,uint64) (uint64))(nil)).Elem(), funcMakerUint64Uint64ГUint64)
	reflectx.RegisterFunc(reflect.TypeOf((*func(uint8,uint8) (uint8))(nil)).Elem(), funcMakerUint8Uint8ГUint8)
	reflectx.RegisterFunc(reflect.TypeOf((*func(uint,uint) (uint))(nil)).Elem(), funcMakerUintUintГUint)
	reflectx.RegisterFunc(reflect.TypeOf((*func() (meanAccum))(nil)).Elem(), funcMakerГMeanAccum)
}

func wrapMakerMeanFn(fn interface{}) map[string]reflectx.Func {
	dfn := fn.(*meanFn)
	return map[string]reflectx.Func{
		"AddInput": reflectx.MakeFunc(func(a0 meanAccum, a1 typex.T) (meanAccum) { return dfn.AddInput(a0, a1) }),
		"CreateAccumulator": reflectx.MakeFunc(func() (meanAccum) { return dfn.CreateAccumulator() }),
		"ExtractOutput": reflectx.MakeFunc(func(a0 meanAccum) (float64) { return dfn.ExtractOutput(a0) }),
		"MergeAccumulators": reflectx.MakeFunc(func(a0 []meanAccum) (meanAccum) { return dfn.MergeAccumulators(a0) }),
	}
}

type callerFloat32Float32ГFloat32 struct {
	fn func(float32,float32) (float32)
}

func funcMakerFloat32Float32ГFloat32(fn interface{}) reflectx.Func {
	f := fn.(func(float32,float32) (float32))
	return &callerFloat32Float32ГFloat32{fn: f}
}

func (c *callerFloat32Float32ГFloat32) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerFloat32Float32ГFloat32) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerFloat32Float32ГFloat32) Call(args []interface{}) []interface{} {
	out0 := c.fn(args[0].(float32), args[1].(float32))
	return []interface{}{out0}
}

func (c *callerFloat32Float32ГFloat32) Call2x1(arg0, arg1 interface{}) (interface{}) {
	return c.fn(arg0.(float32), arg1.(float32))
}

type callerFloat64Float64ГFloat64 struct {
	fn func(float64,float64) (float64)
}

func funcMakerFloat64Float64ГFloat64(fn interface{}) reflectx.Func {
	f := fn.(func(float64,float64) (float64))
	return &callerFloat64Float64ГFloat64{fn: f}
}

func (c *callerFloat64Float64ГFloat64) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerFloat64Float64ГFloat64) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerFloat64Float64ГFloat64) Call(args []interface{}) []interface{} {
	out0 := c.fn(args[0].(float64), args[1].(float64))
	return []interface{}{out0}
}

func (c *callerFloat64Float64ГFloat64) Call2x1(arg0, arg1 interface{}) (interface{}) {
	return c.fn(arg0.(float64), arg1.(float64))
}

type callerInt16Int16ГInt16 struct {
	fn func(int16,int16) (int16)
}

func funcMakerInt16Int16ГInt16(fn interface{}) reflectx.Func {
	f := fn.(func(int16,int16) (int16))
	return &callerInt16Int16ГInt16{fn: f}
}

func (c *callerInt16Int16ГInt16) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerInt16Int16ГInt16) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerInt16Int16ГInt16) Call(args []interface{}) []interface{} {
	out0 := c.fn(args[0].(int16), args[1].(int16))
	return []interface{}{out0}
}

func (c *callerInt16Int16ГInt16) Call2x1(arg0, arg1 interface{}) (interface{}) {
	return c.fn(arg0.(int16), arg1.(int16))
}

type callerInt32Int32ГInt32 struct {
	fn func(int32,int32) (int32)
}

func funcMakerInt32Int32ГInt32(fn interface{}) reflectx.Func {
	f := fn.(func(int32,int32) (int32))
	return &callerInt32Int32ГInt32{fn: f}
}

func (c *callerInt32Int32ГInt32) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerInt32Int32ГInt32) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerInt32Int32ГInt32) Call(args []interface{}) []interface{} {
	out0 := c.fn(args[0].(int32), args[1].(int32))
	return []interface{}{out0}
}

func (c *callerInt32Int32ГInt32) Call2x1(arg0, arg1 interface{}) (interface{}) {
	return c.fn(arg0.(int32), arg1.(int32))
}

type callerInt64Int64ГInt64 struct {
	fn func(int64,int64) (int64)
}

func funcMakerInt64Int64ГInt64(fn interface{}) reflectx.Func {
	f := fn.(func(int64,int64) (int64))
	return &callerInt64Int64ГInt64{fn: f}
}

func (c *callerInt64Int64ГInt64) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerInt64Int64ГInt64) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerInt64Int64ГInt64) Call(args []interface{}) []interface{} {
	out0 := c.fn(args[0].(int64), args[1].(int64))
	return []interface{}{out0}
}

func (c *callerInt64Int64ГInt64) Call2x1(arg0, arg1 interface{}) (interface{}) {
	return c.fn(arg0.(int64), arg1.(int64))
}

type callerInt8Int8ГInt8 struct {
	fn func(int8,int8) (int8)
}

func funcMakerInt8Int8ГInt8(fn interface{}) reflectx.Func {
	f := fn.(func(int8,int8) (int8))
	return &callerInt8Int8ГInt8{fn: f}
}

func (c *callerInt8Int8ГInt8) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerInt8Int8ГInt8) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerInt8Int8ГInt8) Call(args []interface{}) []interface{} {
	out0 := c.fn(args[0].(int8), args[1].(int8))
	return []interface{}{out0}
}

func (c *callerInt8Int8ГInt8) Call2x1(arg0, arg1 interface{}) (interface{}) {
	return c.fn(arg0.(int8), arg1.(int8))
}

type callerIntIntГInt struct {
	fn func(int,int) (int)
}

func funcMakerIntIntГInt(fn interface{}) reflectx.Func {
	f := fn.(func(int,int) (int))
	return &callerIntIntГInt{fn: f}
}

func (c *callerIntIntГInt) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerIntIntГInt) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerIntIntГInt) Call(args []interface{}) []interface{} {
	out0 := c.fn(args[0].(int), args[1].(int))
	return []interface{}{out0}
}

func (c *callerIntIntГInt) Call2x1(arg0, arg1 interface{}) (interface{}) {
	return c.fn(arg0.(int), arg1.(int))
}

type callerMeanAccumTypex۰TГMeanAccum struct {
	fn func(meanAccum,typex.T) (meanAccum)
}

func funcMakerMeanAccumTypex۰TГMeanAccum(fn interface{}) reflectx.Func {
	f := fn.(func(meanAccum,typex.T) (meanAccum))
	return &callerMeanAccumTypex۰TГMeanAccum{fn: f}
}

func (c *callerMeanAccumTypex۰TГMeanAccum) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerMeanAccumTypex۰TГMeanAccum) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerMeanAccumTypex۰TГMeanAccum) Call(args []interface{}) []interface{} {
	out0 := c.fn(args[0].(meanAccum), args[1].(typex.T))
	return []interface{}{out0}
}

func (c *callerMeanAccumTypex۰TГMeanAccum) Call2x1(arg0, arg1 interface{}) (interface{}) {
	return c.fn(arg0.(meanAccum), arg1.(typex.T))
}

type callerMeanAccumГFloat64 struct {
	fn func(meanAccum) (float64)
}

func funcMakerMeanAccumГFloat64(fn interface{}) reflectx.Func {
	f := fn.(func(meanAccum) (float64))
	return &callerMeanAccumГFloat64{fn: f}
}

func (c *callerMeanAccumГFloat64) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerMeanAccumГFloat64) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerMeanAccumГFloat64) Call(args []interface{}) []interface{} {
	out0 := c.fn(args[0].(meanAccum))
	return []interface{}{out0}
}

func (c *callerMeanAccumГFloat64) Call1x1(arg0 interface{}) (interface{}) {
	return c.fn(arg0.(meanAccum))
}

type callerSliceofMeanAccumГMeanAccum struct {
	fn func([]meanAccum) (meanAccum)
}

func funcMakerSliceofMeanAccumГMeanAccum(fn interface{}) reflectx.Func {
	f := fn.(func([]meanAccum) (meanAccum))
	return &callerSliceofMeanAccumГMeanAccum{fn: f}
}

func (c *callerSliceofMeanAccumГMeanAccum) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerSliceofMeanAccumГMeanAccum) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerSliceofMeanAccumГMeanAccum) Call(args []interface{}) []interface{} {
	out0 := c.fn(args[0].([]meanAccum))
	return []interface{}{out0}
}

func (c *callerSliceofMeanAccumГMeanAccum) Call1x1(arg0 interface{}) (interface{}) {
	return c.fn(arg0.([]meanAccum))
}

type callerTypex۰TГTypex۰TInt struct {
	fn func(typex.T) (typex.T,int)
}

func funcMakerTypex۰TГTypex۰TInt(fn interface{}) reflectx.Func {
	f := fn.(func(typex.T) (typex.T,int))
	return &callerTypex۰TГTypex۰TInt{fn: f}
}

func (c *callerTypex۰TГTypex۰TInt) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerTypex۰TГTypex۰TInt) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerTypex۰TГTypex۰TInt) Call(args []interface{}) []interface{} {
	out0, out1 := c.fn(args[0].(typex.T))
	return []interface{}{out0, out1}
}

func (c *callerTypex۰TГTypex۰TInt) Call1x2(arg0 interface{}) (interface{}, interface{}) {
	return c.fn(arg0.(typex.T))
}

type callerUint16Uint16ГUint16 struct {
	fn func(uint16,uint16) (uint16)
}

func funcMakerUint16Uint16ГUint16(fn interface{}) reflectx.Func {
	f := fn.(func(uint16,uint16) (uint16))
	return &callerUint16Uint16ГUint16{fn: f}
}

func (c *callerUint16Uint16ГUint16) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerUint16Uint16ГUint16) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerUint16Uint16ГUint16) Call(args []interface{}) []interface{} {
	out0 := c.fn(args[0].(uint16), args[1].(uint16))
	return []interface{}{out0}
}

func (c *callerUint16Uint16ГUint16) Call2x1(arg0, arg1 interface{}) (interface{}) {
	return c.fn(arg0.(uint16), arg1.(uint16))
}

type callerUint32Uint32ГUint32 struct {
	fn func(uint32,uint32) (uint32)
}

func funcMakerUint32Uint32ГUint32(fn interface{}) reflectx.Func {
	f := fn.(func(uint32,uint32) (uint32))
	return &callerUint32Uint32ГUint32{fn: f}
}

func (c *callerUint32Uint32ГUint32) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerUint32Uint32ГUint32) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerUint32Uint32ГUint32) Call(args []interface{}) []interface{} {
	out0 := c.fn(args[0].(uint32), args[1].(uint32))
	return []interface{}{out0}
}

func (c *callerUint32Uint32ГUint32) Call2x1(arg0, arg1 interface{}) (interface{}) {
	return c.fn(arg0.(uint32), arg1.(uint32))
}

type callerUint64Uint64ГUint64 struct {
	fn func(uint64,uint64) (uint64)
}

func funcMakerUint64Uint64ГUint64(fn interface{}) reflectx.Func {
	f := fn.(func(uint64,uint64) (uint64))
	return &callerUint64Uint64ГUint64{fn: f}
}

func (c *callerUint64Uint64ГUint64) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerUint64Uint64ГUint64) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerUint64Uint64ГUint64) Call(args []interface{}) []interface{} {
	out0 := c.fn(args[0].(uint64), args[1].(uint64))
	return []interface{}{out0}
}

func (c *callerUint64Uint64ГUint64) Call2x1(arg0, arg1 interface{}) (interface{}) {
	return c.fn(arg0.(uint64), arg1.(uint64))
}

type callerUint8Uint8ГUint8 struct {
	fn func(uint8,uint8) (uint8)
}

func funcMakerUint8Uint8ГUint8(fn interface{}) reflectx.Func {
	f := fn.(func(uint8,uint8) (uint8))
	return &callerUint8Uint8ГUint8{fn: f}
}

func (c *callerUint8Uint8ГUint8) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerUint8Uint8ГUint8) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerUint8Uint8ГUint8) Call(args []interface{}) []interface{} {
	out0 := c.fn(args[0].(uint8), args[1].(uint8))
	return []interface{}{out0}
}

func (c *callerUint8Uint8ГUint8) Call2x1(arg0, arg1 interface{}) (interface{}) {
	return c.fn(arg0.(uint8), arg1.(uint8))
}

type callerUintUintГUint struct {
	fn func(uint,uint) (uint)
}

func funcMakerUintUintГUint(fn interface{}) reflectx.Func {
	f := fn.(func(uint,uint) (uint))
	return &callerUintUintГUint{fn: f}
}

func (c *callerUintUintГUint) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerUintUintГUint) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerUintUintГUint) Call(args []interface{}) []interface{} {
	out0 := c.fn(args[0].(uint), args[1].(uint))
	return []interface{}{out0}
}

func (c *callerUintUintГUint) Call2x1(arg0, arg1 interface{}) (interface{}) {
	return c.fn(arg0.(uint), arg1.(uint))
}

type callerГMeanAccum struct {
	fn func() (meanAccum)
}

func funcMakerГMeanAccum(fn interface{}) reflectx.Func {
	f := fn.(func() (meanAccum))
	return &callerГMeanAccum{fn: f}
}

func (c *callerГMeanAccum) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerГMeanAccum) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerГMeanAccum) Call(args []interface{}) []interface{} {
	out0 := c.fn()
	return []interface{}{out0}
}

func (c *callerГMeanAccum) Call0x1() (interface{}) {
	return c.fn()
}


// DO NOT MODIFY: GENERATED CODE