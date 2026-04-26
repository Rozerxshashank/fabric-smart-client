/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

type mockCloser struct {
	called bool
	err    error
}

func (m *mockCloser) Close() error {
	m.called = true
	return m.err
}

func TestCloseMute(t *testing.T) {
	t.Parallel()
	// Test nil
	CloseMute(nil)

	// Test non-nil
	mc := &mockCloser{}
	CloseMute(mc)
	require.True(t, mc.called)

	// Test with error
	mc = &mockCloser{err: errors.New("close error")}
	CloseMute(mc)
	require.True(t, mc.called)
}

func TestIgnoreError(t *testing.T) {
	t.Parallel()
	// Note: t.Parallel() might interfere with stdout capture if we were doing that,
	// but here we just want to execute the code.
	IgnoreError(nil)
	IgnoreError(errors.New("some error"))
}

func TestIgnoreErrorFunc(t *testing.T) {
	t.Parallel()
	called := false
	IgnoreErrorFunc(func() error {
		called = true
		return errors.New("func error")
	})
	require.True(t, called)
}

func TestIgnoreErrorWithOneArg(t *testing.T) {
	t.Parallel()
	calledWith := ""
	IgnoreErrorWithOneArg(func(s string) error {
		calledWith = s
		return errors.New("arg error")
	}, "hello")
	require.Equal(t, "hello", calledWith)
}

func TestZero(t *testing.T) {
	t.Parallel()
	require.Equal(t, 0, Zero[int]())
	require.Equal(t, "", Zero[string]())
	require.Nil(t, Zero[*int]())
}

func TestMust(t *testing.T) {
	t.Parallel()
	require.NotPanics(t, func() { Must(nil) })
	require.Panics(t, func() { Must(errors.New("panic")) })
}

func TestMustGet(t *testing.T) {
	t.Parallel()
	require.Equal(t, 10, MustGet(10, nil))
	require.Panics(t, func() { MustGet(10, errors.New("panic")) })
}

func TestDefaultZero(t *testing.T) {
	t.Parallel()
	require.Equal(t, 0, DefaultZero[int](nil))
	require.Equal(t, 10, DefaultZero[int](10))
	require.Equal(t, 0, DefaultZero[int]("not int"))
}

func TestDefaultInt(t *testing.T) {
	t.Parallel()
	require.Equal(t, 5, DefaultInt(nil, 5))
	require.Equal(t, 10, DefaultInt(10, 5))
	require.Equal(t, 5, DefaultInt(0, 5))
	require.Equal(t, 5, DefaultInt("not int", 5))
}

func TestDefaultString(t *testing.T) {
	t.Parallel()
	require.Equal(t, "default", DefaultString(nil, "default"))
	require.Equal(t, "value", DefaultString("value", "default"))
	require.Equal(t, "default", DefaultString("", "default"))
	require.Equal(t, "default", DefaultString(10, "default"))
}

func TestIsNil(t *testing.T) {
	t.Parallel()
	require.True(t, IsNil[*int](nil))
	var m map[string]int
	require.True(t, IsNil(m))
	var s []int
	require.True(t, IsNil(s))
	var c chan int
	require.True(t, IsNil(c))
	var f func()
	require.True(t, IsNil(f))

	require.False(t, IsNil(10))
	require.False(t, IsNil("hello"))
	i := 10
	require.False(t, IsNil(&i))
}
