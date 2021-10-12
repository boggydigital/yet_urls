package yt_urls

import "strings"

const (
	opCuBrace = "{"
	clCuBrace = "}"
)

func extractJsonObject(data string) string {
	fi, li := strings.Index(data, opCuBrace), strings.LastIndex(data, clCuBrace)
	return data[fi : li+1]
}
