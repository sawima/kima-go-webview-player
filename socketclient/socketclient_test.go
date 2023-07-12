// package socketclient

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestReadSocket(t *testing.T) {
// 	type wants struct {
// 		want string
// 		reload bool
// 	}
// 	tests := []struct {
// 		name string
// 		args string
// 		want wants
// 	}{
// 		{"case1", "hello", wants{"xx",false}}, {"case2", "world", "http://localhost:8081/apps/VueSingleVideoTemplate/#/?appID=Z2OxE8gM6uswRo"}, {"case3", "james", "http://localhost:8081/apps/VueSingleVideoTemplate/#/?appID=Z2OxE8gM6uswRo"},
// 	}
// 	// for _, tt := range tests {
// 	// 	t.Run(tt.name, func(t *testing.T) {
// 	// 		if got:=ReadSocket(tt.args);got!=tt.want {

// 	// 		}
// 	// 	})
// 	// }
// 	for _, tt := range tests {
// 		assert.Equal(t, ReadSocket(tt.args), tt.want, tt.name)
// 		// t.Run(tt.name, func(t *testing.T) {
// 		// 	if got := Add(tt.args.x, tt.args.y); got != tt.want {
// 		// 		t.Errorf("Add() = %v, want %v", got, tt.want)
// 		// 	}
// 		// })
// 	}
// }
