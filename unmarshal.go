package main

import "encoding/json"

// UnmarshalJSON Convert the main JSON response into a
func (j *jsonData) UnmarshalJSON(data []byte) error {
	aux := &struct {
		Status int      `json:"status"`
		Data   []weapon `json:"data"`
	}{}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	j.Status = aux.Status

	j.Data = make([]weapon, len(aux.Data))
	for i, s := range aux.Data {
		j.Data[i] = s
	}

	return nil
}

// UnmarshalJSON Populate the skins array in each weapon
func (w *weapon) UnmarshalJSON(data []byte) error {
	aux := &struct {
		Name  string `json:"displayName"`
		Skins []skin `json:"skins"`
	}{}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	w.Name = aux.Name

	w.Skins = make([]skin, len(aux.Skins))
	for i, s := range aux.Skins {
		w.Skins[i] = s
	}

	return nil
}

// UnmarshalJSON Populate the chromas in the skins arrays
func (s *skin) UnmarshalJSON(data []byte) error {
	aux := &struct {
		UUID    string   `json:"uuid"`
		Name    string   `json:"displayName"`
		Icon    string   `json:"displayIcon"`
		Chromas []chroma `json:"chromas"`
	}{}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	s.UUID = aux.UUID
	s.Name = aux.Name
	s.Icon = aux.Icon

	s.Chromas = make([]chroma, len(aux.Chromas))
	for i, c := range aux.Chromas {
		s.Chromas[i] = c
	}

	return nil
}
