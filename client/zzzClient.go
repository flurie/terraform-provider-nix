package client

type StorePath string

type Derivation string

type DerivatiopnDetail struct {
	Outputs map[StorePath]struct {
		Path StorePath `json:"path"`
	} `json:"outputs"`
	InputSrcs []StorePath                 `json:"inputSrcs"`
	InputDrvs map[Derivation][]Derivation `json:"inputDrvs"`
	// TODO: this should probably be an enum derived from the nix source
	System  string            `json:"system"`
	Builder StorePath         `json:"builder"`
	Args    []string          `json:"args"`
	Env     map[string]string `json:"env"`
}

type ShowDerivation map[Derivation]DerivatiopnDetail

func (d ShowDerivation) getStorePaths() []StorePath {
	var o []StorePath
	for _, drv := range d {
		for _, output := range drv.Outputs {
			o = append(o, output.Path)
		}
	}
	return o
}
