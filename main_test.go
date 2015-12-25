package main
import (
	"testing"
	"fmt"
)

// สร้าง site จาก Site struct ได้ ถ้าชื่อไซต์ไม่ใช่ตามกำหนด ให้ Error
func Test_CreateSite(t *testing.T) {
	s := Site{ID:1, Name:"บริษัท นพดลพานิช จำกัด"}
	if s.Name != "บริษัท นพดลพานิช จำกัด" {
		t.Error("แสดงชื่อมาไม่ใช่ นพดลพานิช ตามที่ทดสอบ")
	}
	fmt.Printf("Test1.Site.ID: %d Site.Name: %s\n", s.ID, s.Name)
}

// ดึงข้อมูลจาก DB table Site ทั้งหมดมาเก็บใน Struct Site คาดว่าจำนวนรายการต้องครบ
func Test_LoadSite(t *testing.T)  {
	var s Site
	sites := s.LoadDB()
	for _, s = range sites {
		fmt.Printf("Test2.Site:ID: %v Name: %s\n", s.ID, s.Name)
	}
}
// สร้างตัวเชื่อมฐานข้อมูลใหม่ ที่เป็นอิสระโดยแต่ละฟังค์ชั่นจะใช้ SQlite3 หรือ  MySQL ก็ได้
func Test_NewDB(t *testing.T) {
	db, err := NewDB("./server.db")
	if err != nil {
		t.Error("NewDB:", err)
	}
	if err = db.Ping(); err != nil {
		t.Error(err)
	}
}