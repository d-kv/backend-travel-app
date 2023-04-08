package util

import (
	"strconv"
	"strings"
)

type Version struct {
	// TODO: use standard Major.Minor.Path versioning
	// TODO: Patch might be a string like rc
	Patch uint64
}

func (v Version) Less(otherV Version) bool {
	return v.Patch < otherV.Patch
}

func NewVersion(patch uint64) *Version {
	return &Version{
		Patch: patch,
	}
}

func NewVersionFromString(patchStr string) (*Version, error) {
	v := &Version{}
	if err := parseVersionFromString(v, patchStr); err != nil {
		return nil, err
	}
	return v, nil
}

func parseVersionFromString(v *Version, rawStr string) error {
	p, err := strconv.ParseUint(strings.TrimSpace(rawStr), base, bitSize)
	if err != nil {
		return err
	}

	v.Patch = p

	return nil
}
