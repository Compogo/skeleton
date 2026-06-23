// Package main — точка входа в приложение.
//
// Задача main — максимально простая:
//  1. Создать корневую команду CLI.
//  2. Выполнить её.
//
// Вся сложность вынесена в пакет cli.
// Это позволяет легко тестировать приложение и добавлять новые команды.
package main

import (
	"github.com/Compogo/sceleton/interface/cli"
)

func main() {
	// Создаём корневую команду.
	cmd, err := cli.NewRootCommand()
	if err != nil {
		panic(err)
	}

	// Выполняем команду.
	// При успешном выполнении cmd.Execute() возвращает nil.
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
