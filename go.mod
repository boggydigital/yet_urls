module github.com/boggydigital/yt_urls

go 1.16

require (
	github.com/boggydigital/gost v0.1.0
	github.com/boggydigital/match_node v0.1.1
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5
)

replace (
	github.com/boggydigital/gost => ../gost
)
