package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/skip2/go-qrcode"
	"golang.design/x/clipboard"
)

// สร้างชื่อไฟล์จาก URL หรือ text
func GenerateFileName(input string) string {
	// ลบ protocol
	cleaned := strings.ReplaceAll(input, "https://", "")
	cleaned = strings.ReplaceAll(cleaned, "http://", "")
	cleaned = strings.ReplaceAll(cleaned, "www.", "")
	
	// ถ้าเป็น URL ใช้ domain name
	if strings.Contains(cleaned, ".") && strings.Contains(cleaned, "/") {
		parts := strings.Split(cleaned, "/")
		domain := parts[0]
		if strings.Contains(domain, ".") {
			domainParts := strings.Split(domain, ".")
			cleaned = domainParts[0]
		}
	} else if strings.Contains(cleaned, ".") {
		domainParts := strings.Split(cleaned, ".")
		cleaned = domainParts[0]
	} else {
		cleaned = input
	}
	
	// ทำความสะอาดชื่อไฟล์
	reg := regexp.MustCompile(`[^\w\-_]`)
	cleaned = reg.ReplaceAllString(cleaned, "_")
	
	// จำกัดความยาว
	if len(cleaned) > 20 {
		cleaned = cleaned[:20]
	}
	
	if cleaned == "" {
		cleaned = "qrcode"
	}
	
	return cleaned
}

// สร้าง QR Code และบันทึกไฟล์
func CreateAndSaveQR(input string) error {
	// Initialize clipboard
	err := clipboard.Init()
	if err != nil {
		return fmt.Errorf("ไม่สามารถเปิด clipboard: %v", err)
	}

	// สร้างโฟลเดอร์ qrcode
	qrDir := "qrcode"
	if err := os.MkdirAll(qrDir, 0755); err != nil {
		return fmt.Errorf("ไม่สามารถสร้างโฟลเดอร์: %v", err)
	}
	
	// สร้างชื่อไฟล์
	fileName := GenerateFileName(input) + ".png"
	filePath := filepath.Join(qrDir, fileName)
	
	// ตรวจสอบไฟล์ซ้ำ
	counter := 1
	for {
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			break
		}
		baseName := strings.TrimSuffix(fileName, ".png")
		newFileName := fmt.Sprintf("%s_%d.png", baseName, counter)
		filePath = filepath.Join(qrDir, newFileName)
		counter++
	}
	
	// สร้าง QR Code
	err = qrcode.WriteFile(input, qrcode.Medium, 256, filePath)
	if err != nil {
		return fmt.Errorf("ไม่สามารถสร้าง QR Code: %v", err)
	}
	
	// อ่านไฟล์รูป
	imageData, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("ไม่สามารถอ่านไฟล์รูป: %v", err)
	}
	
	// Copy รูปเข้า clipboard
	clipboard.Write(clipboard.FmtImage, imageData)
	
	return nil
}