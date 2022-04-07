package piperenv

type Artifact struct {
	Kind string `json:"kind"`
	Path string `json:"path"`
}

type Artifacts []Artifact

func (a Artifacts) FindByKind(kind string) Artifacts {
	var filtered Artifacts

	for _, artifact := range a {
		if artifact.Kind == kind {
			filtered = append(filtered, artifact)
		}
	}

	return filtered
}
