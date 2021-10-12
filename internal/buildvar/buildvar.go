package buildvar

// -ldflags '-X "github.com/danbrakeley/actiontest/internal/buildvar.Version=${{ github.event.release.tag_name }}"'
var Version string

// -ldflags '-X "github.com/danbrakeley/actiontest/internal/buildvar.BuildTime=${{ github.event.release.created_at }}"'
var BuildTime string

// -ldflags '-X "github.com/danbrakeley/actiontest/internal/buildvar.ReleaseDescription=${{ github.event.release.body }}"'
var ReleaseNotes string

// -ldflags '-X "github.com/danbrakeley/actiontest/internal/buildvar.ReleaseURL=${{ github.event.release.html_url }}"'
var ReleaseURL string
