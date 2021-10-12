package buildvar

// -ldflags "-X 'github.com/danbrakeley/actiontest/internal/buildvar.Version=vX.Y.Z'"
var Version string

// -ldflags "-X 'github.com/danbrakeley/actiontest/internal/buildvar.BuildTime=$(date -Iseconds)'"
var BuildTime string
