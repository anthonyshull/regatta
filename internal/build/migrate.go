//+build mage

package main

import (
	"context"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Migrate mg.Namespace

func (Migrate) Up() error {
	err := sh.Run("journey", "--url", "cassandra://127.0.0.1:9042/regatta", "--path", "../migrations", "migrate", "up")
	if err != nil {
		return err
	}
	return nil
}

func (Migrate) Down() error {
	err := sh.Run("journey", "--url", "cassandra://127.0.0.1:9042/regatta", "--path", "../migrations", "migrate", "down")
	if err != nil {
		return err
	}
	return nil
}

func (Migrate) Create(ctx context.Context) error {
	name := os.Getenv("MIGRATION")
	if name == "" {
		return nil
	}
	err := sh.Run("journey", "--url", "cassandra://127.0.0.1:9042/regatta", "--path", "../migrations", "create", name)
	if err != nil {
		return err
	}
	return nil
}