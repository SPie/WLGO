package response

import (
    "testing"
    "encoding/json"
)

func TestCoordinates(t *testing.T) {
    t.Log("Should create Coordinates")
        {
        var coordinates Coordinates = [2]float64{10.10, 11.11}

        if coordinates[0] != 10.10 {
            t.Fatal("Latitude should be 10.10")
        }
        if coordinates[1] != 11.11 {
            t.Fatal("Longitude should be 11.11")
        }

        t.Log("Correct Coordinates")
    }
}

func TestParseCoordinatesFromJson(t *testing.T) {
    jsonString := "[10.10, 11.11]"
    var coordinates Coordinates
    err := json.Unmarshal([]byte(jsonString), &coordinates)
    t.Log("Should create Coordinates from JSOn with latitude and longitude")
    {
        if err != nil {
            t.Fatalf("Should be able to decode. Error: %s", err)
        }

        if coordinates[0] != 10.10 {
            t.Fatal("Latitude should be 10.10")
        }
        if coordinates[1] != 11.11 {
            t.Fatal("Longitude should be 11.11")
        }

        t.Log("Coordinates successfully parsed")
    }
}

func TestParseCoordinatesFromEmptyJsonString(t *testing.T) {
    var coordinates Coordinates
    err := json.Unmarshal([]byte("[]"), &coordinates)
    t.Log("Should create empty Coordinates")
    {
        if err != nil {
            t.Fatalf("Should be able to unmarshal. Error: %s", err)
        }
        if coordinates[0] != 0 {
            t.Fatal("Latitude should be empty")
        }
        if coordinates[1] != 0 {
            t.Fatal("Longitude should be empty")
        }
        t.Log("Parsed empty Coordinates")
    }
}

func TestGetLatitude(t *testing.T) {
    var coordinates Coordinates = [2]float64{10.10, 11.11}
    t.Log("Should get Latitude")
    {
        if coordinates.GetLatitude() != 10.10 {
            t.Fatal("Latitude should be 10.10")
        }
    }
}

func TestGetLongitude(t *testing.T) {
    var coordinates Coordinates = [2]float64{10.10, 11.11}
    t.Log("Should get Longitude")
    {
        if coordinates.GetLongitude() != 11.11 {
            t.Fatal("Longitude should be 11.11")
        }
    }
}

func TestGetLatitudeWithEmptyCoordinates(t *testing.T) {
    var coordinates Coordinates = [2]float64{}
    t.Log("Latitude should be empty")
    {
        if coordinates.GetLatitude() != 0 {
            t.Fatal("Latitude is not empty")
        }
    }
}

func TestGetLongitudeWithEmptyCoordinates(t *testing.T) {
    var coordinates Coordinates = [2]float64{}
    t.Log("Longitude should be empty")
    {
        if coordinates.GetLongitude() != 0 {
            t.Fatal("Longitude is not empty")
        }
    }

}
