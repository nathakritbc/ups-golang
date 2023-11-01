package main

import (
	"fmt"

	"github.com/mdlayher/apcupsd"
)

func main() {
	// เชื่อมต่อกับ APCUPSD server ที่ใช้พอร์ต 3551 บน localhost
	// c, err := apcupsd.Dial("tcp", "localhost:3551")
	c, err := apcupsd.Dial("tcp", "127.0.0.1:3551")

	if err != nil {
		fmt.Printf("เกิดข้อผิดพลาดในการเชื่อมต่อ: %v\n", err)
		return
	}
	defer c.Close()

	// อ่านข้อมูลสถานะจากเครื่อง UPS
	status, err := c.Status()
	if err != nil {
		fmt.Printf("เกิดข้อผิดพลาดในการอ่านข้อมูลสถานะ: %v\n", err)
		return
	}

	// แสดงข้อมูลสถานะ
	fmt.Printf("Model: %s\n", status.Model)
	fmt.Printf("Status: %s\n", status.Status)
	fmt.Printf("Battery Charge: %.2f%%\n", status.BatteryChargePercent)
	fmt.Printf("Battery Voltage: %.2f V\n", status.BatteryVoltage)
	fmt.Printf("LineVoltage : %.2f V\n", status.LineVoltage)

	fmt.Printf("Load: %.2f%%\n", status.LoadPercent)

	fmt.Printf("status : %s\n", status)
}
