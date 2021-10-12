package buildvar

// -ldflags '-X "github.com/danbrakeley/actiontest/internal/buildvar.Version=${{ github.event.release.tag_name }}"'
var Version string

// -ldflags '-X "github.com/danbrakeley/actiontest/internal/buildvar.BuildTime=${{ github.event.release.tag_name }}"'
var BuildTime string

// -ldflags '-X "github.com/danbrakeley/actiontest/internal/buildvar.ReleaseName=${{ github.event.release.name }}"'
var RelaseName string

// -ldflags '-X "github.com/danbrakeley/actiontest/internal/buildvar.ReleaseDescription=${{ github.event.release.body }}"'
var ReleaseDescription string

// -ldflags '-X "github.com/danbrakeley/actiontest/internal/buildvar.ReleaseURL=${{ github.event.release.html_url }}"'
var ReleaseURL string
