package util

type Version struct {
	// TODO: use standard Major.Minor.Path versioning
	// TODO: Patch might be a string like rc
	Patch int32
}

func (v Version) Less(otherV Version) bool {
	return v.Patch > otherV.Patch
}

func NewVersion(patch int32) *Version {
	return &Version{
		Patch: patch,
	}
}
