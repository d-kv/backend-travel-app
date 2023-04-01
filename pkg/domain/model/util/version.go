package util

type Version struct {
	// TODO: use standard Major.Minor.Path versioning
	Patch int32
}

func (v Version) Less(otherV Version) bool {
	return v.Patch > otherV.Patch
}
