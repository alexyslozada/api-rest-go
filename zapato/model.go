package zapato

var storage Storage

func init() {
	storage = make(map[string]*Model)
}

type Model struct {
	Marca  string `json:"marca"`
	Precio int    `json:"precio"`
	Color  string `json:"color"`
}

type Storage map[string]*Model

func (s Storage) Create(m *Model) *Model {
	s[m.Marca] = m
	return s[m.Marca]
}

func (s Storage) GetAll() Storage {
	return s
}

func (s Storage) GetByMarca(m string) *Model {
	if v, ok := s[m]; ok {
		return v
	}

	return nil
}

func (s Storage) Delete(m string) {
	delete(s, m)
}

func (s Storage) Update(m string, z *Model) {
	s[m] = z
}
