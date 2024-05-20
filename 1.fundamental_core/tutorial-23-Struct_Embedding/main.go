package main

import (
	"fmt"
)

// Definisikan struct dasar untuk kendaraan umum
type Vehicle struct {
    Brand string
    Model string
    Year  int
}

// Definisikan struct untuk mobil yang meng-embed struct Vehicle
type Car struct {
    Vehicle   // embedding struct Vehicle
    Doors     int
    IsElectric bool
}

// Definisikan struct untuk truk yang meng-embed struct Vehicle
type Truck struct {
    Vehicle   // embedding struct Vehicle
    Capacity  int // kapasitas dalam ton
    IsFourWheelDrive bool
}

func main() {
    // Buat instance dari struct Car
    car := Car{
        Vehicle: Vehicle{
            Brand: "Himalayan",
            Model: "2.0",
            Year:  2021,
        },
        Doors: 4,
        IsElectric: true,
    }

    // Buat instance dari struct Truck
    truck := Truck{
        Vehicle: Vehicle{
            Brand: "Ford",
            Model: "F-150",
            Year:  2020,
        },
        Capacity: 2, // kapasitas 2 ton
        IsFourWheelDrive: true,
    }

    // Tampilkan informasi tentang mobil
    fmt.Printf("Car: %s %s (%d), Doors: %d, Electric: %t\n", 
        car.Brand, car.Model, car.Year, car.Doors, car.IsElectric)
    
    // Tampilkan informasi tentang truk
    fmt.Printf("Truck: %s %s (%d), Capacity: %d ton, 4WD: %t\n", 
        truck.Brand, truck.Model, truck.Year, truck.Capacity, truck.IsFourWheelDrive)
}
