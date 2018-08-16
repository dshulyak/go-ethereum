package les

type Enabler interface {
	Enable()
	Disable()
}

type FetcherApi struct {
	e Enabler
}

func (f FetcherApi) EnableSync() {
	f.e.Enable()
}

func (f FetcherApi) DisableSync() {
	f.e.Disable()
}
