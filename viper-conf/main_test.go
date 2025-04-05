package main

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestGetPort(t *testing.T) {
	t.Run("Should return the port when PORT is set", func(t *testing.T) {
		expectedPort := "8080"
		viper.Set("PORT", expectedPort)

		port := getPort()

		if port != expectedPort {