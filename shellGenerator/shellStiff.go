package shellGenerator

import "fmt"

// Stiffiner - input data of stiffiners
type Stiffiner struct {
	Amount    int     // unit - items. Amount of stiffiners on shell
	Height    float64 // unit - meter. Height of single stiffiner
	Precition float64 // unit - meter. Maximal distance between points
}

func (s Stiffiner) check() (err error) {
	switch {
	case s.Amount < 2:
		return fmt.Errorf("Amount of stiffiners cannot be less 2")
	case s.Height <= 0:
		return fmt.Errorf("Height of stiffiners cannot be less or equal zero")
	case s.Precition <= 0:
		return fmt.Errorf("Precition of stiffiners cannot be less or equal zero")
	}
	return nil
}

// ShellWithStiffiners - input data for shell with stiffiners
type ShellWithStiffiners struct {
	shell      Shell       // data of shell
	stiffiners []Stiffiner // data of stiffiners
}

// AddShell - add shell
func (s *ShellWithStiffiners) AddShell(sh Shell) (err error) {
	if err = sh.check(); err != nil {
		return err
	}
	s.shell = sh
	return nil
}

// AddStiffiners - add stiffiners on shell
func (s *ShellWithStiffiners) AddStiffiners(st Stiffiner) (err error) {
	if err = st.check(); err != nil {
		return err
	}
	s.stiffiners = append(s.stiffiners, st)
	return nil
}

// Generate mesh of shell
func (s ShellWithStiffiners) Generate(offset bool) (mesh Mesh, err error) {
	if offset {
		return s.generateWithOffset()
	}
	return s.generateWithoutOffset()
}
