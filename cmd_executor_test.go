package main

import "testing"

func TestGetSystemInfo(t *testing.T) {
	cmdr := NewCommander()
	info, err := cmdr.GetSystemInfo()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if info.Hostname == "" {
		t.Error("Expected hostname to be non-empty")
	}

	if info.IPAddress == "" {
		t.Error("Expected IP address to be non-empty")
	}
}

func BenchmarkGetSystemInfo(b *testing.B) {
	cmdr := NewCommander()
	for i := 0; i < b.N; i++ {
		_, err := cmdr.GetSystemInfo()
		if err != nil {
			b.Fatalf("Expected no error, got %v", err)
		}
	}
}
