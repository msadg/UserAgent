package entity

// UserAgents array
type Uas []string

// browsers json data
// type Browsers struct {
// 	Chrome  Uas `json:"chrome"`
// 	Opera   Uas `json:"opera"`
// 	Firefox Uas `json:"firefox"`
// 	Ie      Uas `json:"internetexplorer"`
// 	Safari  Uas `json:"safari"`
// }

// UserAgents Json Data
type UasJsonData struct {
	Bs map[string]Uas    `json:"browsers"`  // browsers
	Rs map[string]string `json:"randomize"` // randomize
}

// UaData
type UaData struct {
	Browsers map[string]Uas
	Names    []string
	Res      []string
}
