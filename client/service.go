package client

func (GetURL) ProvinceService() map[string]int {
	Provinces := make(map[string]int)

	for _, item := range ClientInfo.List {
		// check if the item/element exist in the duplicate_frequency map

		_, exist := Provinces[item.Province]

		if exist {
			Provinces[item.Province] += 1 // increase counter by 1 if already in the map
		} else {
			Provinces[item.Province] = 1 // else start counting from 1
		}
	}
	return Provinces
}

type Age struct {
	Age0to30  int `json:"0-30"`
	Age31to60 int `json:"31-60"`
	Age61up   int `json:"61+"`
	AgeNA     int `json:"N/A"`
}

func (GetURL) AgeService() Age {
	AgeGroup := Age{
		Age0to30:  0,
		Age31to60: 0,
		Age61up:   0,
		AgeNA:     0,
	}
	for _, item := range ClientInfo.List {
		if item.Age > 60 {
			AgeGroup.Age61up += 1
		} else if item.Age > 30 {
			AgeGroup.Age31to60 += 1
		} else if item.Age > 0 {
			AgeGroup.Age0to30 += 1
		} else {
			AgeGroup.AgeNA += 1
		}
	}
	return AgeGroup
}
