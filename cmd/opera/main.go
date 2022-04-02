package main

import (
	"github.com/Fantom-foundation/go-opera/cmd/opera/launcher"
)

func main() {
	provideServices := launcher.NewHeavyGethServiceProvider([]string{"console", "--gcmode", "full"})
	provideServices(func(services launcher.HeavyGethServices) {
		launcher.ContinuouslyReportServiceStats(services.Stack, services.Ethereum)
	})
	// if err := launcher.Launch(os.Args); err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	os.Exit(1)
	// }
}
