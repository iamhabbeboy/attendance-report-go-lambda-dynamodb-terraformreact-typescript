package main

type Firestorer struct {
}

func NewFirestore() *Firestorer {
	return &Firestorer{}
}

func (*Firestorer) Store(report Report) (interface{}, error) {
	return nil, nil
}

func (*Firestorer) Get(ID string) error {
	return nil
}
