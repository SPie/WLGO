package response

type Coordinates [2]float64

func (coordinates Coordinates) GetLatitude() (float64) {
    return coordinates[0]
}

func (coordinates Coordinates) GetLongitude() (float64) {
    return coordinates[1]
}
