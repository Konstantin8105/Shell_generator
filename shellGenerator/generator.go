package shellGenerator

// Stiffiner - input data of stiffiners
type Stiffiner struct {
	Amount    uint    // unit - items. Amount of stiffiners on shell
	Height    float64 // unit - meter. Height of single stiffiner
	Precition float64 // unit - meter. Maximal distance between points
}

// ShellWithStiffiners - input data for shell with stiffiners
type ShellWithStiffiners struct {
	Sh Shell     // data of shell
	St Stiffiner // data of stiffiners
}
