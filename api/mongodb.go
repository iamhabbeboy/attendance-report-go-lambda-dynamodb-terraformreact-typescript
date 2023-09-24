package main

type Mongodber struct{}

func NewMongo() *Mongodber {
	return &Mongodber{}
}

func (a *Mongodber) Store(report Report) error {
	return nil
}

func (a *Mongodber) Get(ID string) error {
	return nil
}
