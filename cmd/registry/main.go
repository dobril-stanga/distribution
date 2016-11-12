package main

import (
	_ "net/http/pprof"

	"github.com/dobril-stanga/distribution/registry"
	_ "github.com/dobril-stanga/distribution/registry/auth/htpasswd"
	_ "github.com/dobril-stanga/distribution/registry/auth/silly"
	_ "github.com/dobril-stanga/distribution/registry/auth/token"
	_ "github.com/dobril-stanga/distribution/registry/proxy"
	_ "github.com/dobril-stanga/distribution/registry/storage/driver/azure"
	_ "github.com/dobril-stanga/distribution/registry/storage/driver/filesystem"
	_ "github.com/dobril-stanga/distribution/registry/storage/driver/gcs"
	_ "github.com/dobril-stanga/distribution/registry/storage/driver/inmemory"
	_ "github.com/dobril-stanga/distribution/registry/storage/driver/middleware/cloudfront"
	_ "github.com/dobril-stanga/distribution/registry/storage/driver/middleware/redirect"
	_ "github.com/dobril-stanga/distribution/registry/storage/driver/oss"
	_ "github.com/dobril-stanga/distribution/registry/storage/driver/s3-aws"
	_ "github.com/dobril-stanga/distribution/registry/storage/driver/s3-goamz"
	_ "github.com/dobril-stanga/distribution/registry/storage/driver/swift"
)

func main() {
	registry.RootCmd.Execute()
}
