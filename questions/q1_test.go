package main

import (
	"bytes"
	"strings"
	"testing"
)

//	Самый эффективный способ конкатенации строк
//	strings.Builder начиная с Go 1.10 самый эффективный способ.
//  До Go 1.10 самый эффективный способ bytes.Buffer

// BenchmarkOperatorConcat Конкатенация при помощи строкового оператора +
func BenchmarkOperatorConcat(b *testing.B) {
	var ts string

	for i := 0; i < b.N; i++ {
		ts += "a"
	}
	b.StopTimer()

	if s := strings.Repeat("a", b.N); ts != s {
		b.Errorf("unexpected result; got=%s, want=%s", ts, s)
	}
}

// BenchmarkBufferConcat Конкатенация при помощи bytes.Buffer
func BenchmarkBufferConcat(b *testing.B) {
	var buffer bytes.Buffer

	for i := 0; i < b.N; i++ {
		buffer.WriteString("a")
	}

	b.StopTimer()

	if s := strings.Repeat("a", b.N); buffer.String() != s {
		b.Errorf("unexpected result; got=%s, want=%s", buffer.String(), s)
	}
}

// BenchmarkBuilderConcat Конкатенация при помощи strings.Builder
func BenchmarkBuilderConcat(b *testing.B) {

	//	Builder используется для эффективного формирования строки используя Write метод.
	//	Копирование памяти минимально.
	//	Можно использовать сразу без инициализации.

	var sb strings.Builder

	for i := 0; i < b.N; i++ {
		sb.WriteString("a")
	}

	b.StopTimer()

	if s := strings.Repeat("a", b.N); sb.String() != s {
		b.Errorf("unexpected result; got=%s, want=%s", sb.String(), s)
	}
}

// BenchmarkCopy Конкатенация при помощи copy
func BenchmarkCopy(b *testing.B) {
	bs := make([]byte, b.N)
	bl := 0

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bl += copy(bs[bl:], "x")
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); string(bs) != s {
		b.Errorf("unexpected result; got=%s, want=%s", string(bs), s)
	}
}
