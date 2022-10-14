package main

import (
	"testing"

	"github.com/Ribeiro/go-password-cracker/pass"
	"github.com/stretchr/testify/assert"
)

func TestHash_superman(t *testing.T) {
	expected := pass.CrackSha1Hash("18c28604dd31094a8d69dae60f1bcd347f1afc5a", false)
	assert.Equal(t, expected, "superman", "password is equal to superman")
}
func TestSalt_superman(t *testing.T) {
	expected := pass.CrackSha1Hash("53d8b3dc9d39f0184144674e310185e41a87ffd5", true)
	assert.Equal(t, expected, "superman", "password is equal to superman")
}

func TestHash_NOT_FOUND(t *testing.T) {
	expected := pass.CrackSha1Hash("03810a46a2c1a0eae58d9332f01c32bdcec9a01a", false)
	assert.Equal(t, expected, "PASSWORD NOT IN DATABASE", "password is not in database")
}
