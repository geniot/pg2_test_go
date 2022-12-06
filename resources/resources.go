package resources

import "embed"

var (
	//go:embed media/*
	MEDIA_LIST embed.FS
)
