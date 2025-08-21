package test

import (
	"PicNest/internal/utils"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	s, _ := utils.HashPassword("123456")
	exp := "$2a$10$zFwovhwDdb5X4NFBDOIo9.c5UD3sMZXpkcKBepBu7t3OCiJ.BQdRi"
	if err := bcrypt.CompareHashAndPassword([]byte(s), []byte("123456")); err == nil {
		t.Logf("HashPassword test passed: %s", "123456")
	} else {
		t.Errorf("HashPassword test failed: expected %s, got %s", exp, s)
	}
}
