package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/alisdairrankine/chip8"
)

func main() {

	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err == nil {
			contents, err := ioutil.ReadAll(file)
			if err == nil {
				Program = contents
			}
		}

	}

	fmt.Println(chip8.DisassembleProgram(Program))

}

func run() {
	clock := time.Tick(time.Second / time.Duration(10))
	cpu := chip8.NewCPU(clock)

	//Load BootLoader
	cpu.LoadData(0x200, Program)

	//Load font
	cpu.LoadData(0, chip8.DefaultFont)
	cpu.Run()
}

var Program = []byte{
	0x12, 0x94, 0x62, 0x09, 0x64, 0x10, 0x83, 0x20, 0x22, 0x3E, 0x72, 0xFF, 0x32, 0xFF, 0x12, 0x06, 0x62, 0x0F, 0xA2, 0x84, 0xF0, 0x65, 0x8F, 0x00, 0xA2, 0x84, 0xF2, 0x1E, 0xF0, 0x65, 0xA2, 0x84, 0xF0, 0x55, 0xA2, 0x84, 0xF2, 0x1E, 0x80, 0xF0, 0xF0, 0x55, 0x72, 0xFF, 0x63, 0x00, 0x84, 0x20, 0x22, 0x3E, 0x32, 0x00, 0x12, 0x12, 0x00, 0xEE, 0x85, 0x60, 0x87, 0x00, 0x12, 0x68, 0xA2, 0x84, 0xF3, 0x1E, 0xF0, 0x65, 0x88, 0x00, 0x86, 0x3E, 0x8E, 0x40, 0x8E, 0x65, 0x3F, 0x01, 0x00, 0xEE, 0xA2, 0x84, 0xF6, 0x1E, 0xF1, 0x65, 0x85, 0x60, 0x75, 0x01, 0x87, 0x10, 0x8E, 0x10, 0x8E, 0x05, 0x3F, 0x01, 0x12, 0x38, 0x96, 0x40, 0x12, 0x38, 0x8E, 0x70, 0x8E, 0x87, 0x4F, 0x01, 0x00, 0xEE, 0xA2, 0x84, 0xF3, 0x1E, 0x80, 0x70, 0xF0, 0x55, 0xA2, 0x84, 0xF5, 0x1E, 0x80, 0x80, 0xF0, 0x55, 0x83, 0x50, 0x12, 0x46, 0x0E, 0x05, 0x0F, 0x06, 0x01, 0x03, 0x0A, 0x07, 0x00, 0x09, 0x0B, 0x04, 0x02, 0x0D, 0x08, 0x0C, 0x22, 0x02, 0xA2, 0x84, 0xFF, 0x65, 0x00, 0xEE,
}