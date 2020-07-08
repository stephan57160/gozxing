package oned

import (
	"github.com/stephan57160/gozxing"
	"strings"
)

// This object renders a CODE128 code as a {@link BitMatrix}.
// SGU : Developped from BOOMBULER development.

type code128BisEncoder struct{}

const abTable = " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_"
const bTable = abTable + "`abcdefghijklmnopqrstuvwxyz{|}~\u007F"
const aOnlyTable = "\u0000\u0001\u0002\u0003\u0004" + // NUL, SOH, STX, ETX, EOT
    "\u0005\u0006\u0007\u0008\u0009" + // ENQ, ACK, BEL, BS,  HT
    "\u000A\u000B\u000C\u000D\u000E" + // LF,  VT,  FF,  CR,  SO
    "\u000F\u0010\u0011\u0012\u0013" + // SI,  DLE, DC1, DC2, DC3
    "\u0014\u0015\u0016\u0017\u0018" + // DC4, NAK, SYN, ETB, CAN
    "\u0019\u001A\u001B\u001C\u001D" + // EM,  SUB, ESC, FS,  GS
    "\u001E\u001F" // RS,  US

    const aTable = abTable + aOnlyTable

var code128_EncodingTable = [107][]bool{
	[]bool{true, true, false, true, true, false, false, true, true, false, false},
	[]bool{true, true, false, false, true, true, false, true, true, false, false},
	[]bool{true, true, false, false, true, true, false, false, true, true, false},
	[]bool{true, false, false, true, false, false, true, true, false, false, false},
	[]bool{true, false, false, true, false, false, false, true, true, false, false},
	[]bool{true, false, false, false, true, false, false, true, true, false, false},
	[]bool{true, false, false, true, true, false, false, true, false, false, false},
	[]bool{true, false, false, true, true, false, false, false, true, false, false},
	[]bool{true, false, false, false, true, true, false, false, true, false, false},
	[]bool{true, true, false, false, true, false, false, true, false, false, false},
	[]bool{true, true, false, false, true, false, false, false, true, false, false},
	[]bool{true, true, false, false, false, true, false, false, true, false, false},
	[]bool{true, false, true, true, false, false, true, true, true, false, false},
	[]bool{true, false, false, true, true, false, true, true, true, false, false},
	[]bool{true, false, false, true, true, false, false, true, true, true, false},
	[]bool{true, false, true, true, true, false, false, true, true, false, false},
	[]bool{true, false, false, true, true, true, false, true, true, false, false},
	[]bool{true, false, false, true, true, true, false, false, true, true, false},
	[]bool{true, true, false, false, true, true, true, false, false, true, false},
	[]bool{true, true, false, false, true, false, true, true, true, false, false},
	[]bool{true, true, false, false, true, false, false, true, true, true, false},
	[]bool{true, true, false, true, true, true, false, false, true, false, false},
	[]bool{true, true, false, false, true, true, true, false, true, false, false},
	[]bool{true, true, true, false, true, true, false, true, true, true, false},
	[]bool{true, true, true, false, true, false, false, true, true, false, false},
	[]bool{true, true, true, false, false, true, false, true, true, false, false},
	[]bool{true, true, true, false, false, true, false, false, true, true, false},
	[]bool{true, true, true, false, true, true, false, false, true, false, false},
	[]bool{true, true, true, false, false, true, true, false, true, false, false},
	[]bool{true, true, true, false, false, true, true, false, false, true, false},
	[]bool{true, true, false, true, true, false, true, true, false, false, false},
	[]bool{true, true, false, true, true, false, false, false, true, true, false},
	[]bool{true, true, false, false, false, true, true, false, true, true, false},
	[]bool{true, false, true, false, false, false, true, true, false, false, false},
	[]bool{true, false, false, false, true, false, true, true, false, false, false},
	[]bool{true, false, false, false, true, false, false, false, true, true, false},
	[]bool{true, false, true, true, false, false, false, true, false, false, false},
	[]bool{true, false, false, false, true, true, false, true, false, false, false},
	[]bool{true, false, false, false, true, true, false, false, false, true, false},
	[]bool{true, true, false, true, false, false, false, true, false, false, false},
	[]bool{true, true, false, false, false, true, false, true, false, false, false},
	[]bool{true, true, false, false, false, true, false, false, false, true, false},
	[]bool{true, false, true, true, false, true, true, true, false, false, false},
	[]bool{true, false, true, true, false, false, false, true, true, true, false},
	[]bool{true, false, false, false, true, true, false, true, true, true, false},
	[]bool{true, false, true, true, true, false, true, true, false, false, false},
	[]bool{true, false, true, true, true, false, false, false, true, true, false},
	[]bool{true, false, false, false, true, true, true, false, true, true, false},
	[]bool{true, true, true, false, true, true, true, false, true, true, false},
	[]bool{true, true, false, true, false, false, false, true, true, true, false},
	[]bool{true, true, false, false, false, true, false, true, true, true, false},
	[]bool{true, true, false, true, true, true, false, true, false, false, false},
	[]bool{true, true, false, true, true, true, false, false, false, true, false},
	[]bool{true, true, false, true, true, true, false, true, true, true, false},
	[]bool{true, true, true, false, true, false, true, true, false, false, false},
	[]bool{true, true, true, false, true, false, false, false, true, true, false},
	[]bool{true, true, true, false, false, false, true, false, true, true, false},
	[]bool{true, true, true, false, true, true, false, true, false, false, false},
	[]bool{true, true, true, false, true, true, false, false, false, true, false},
	[]bool{true, true, true, false, false, false, true, true, false, true, false},
	[]bool{true, true, true, false, true, true, true, true, false, true, false},
	[]bool{true, true, false, false, true, false, false, false, false, true, false},
	[]bool{true, true, true, true, false, false, false, true, false, true, false},
	[]bool{true, false, true, false, false, true, true, false, false, false, false},
	[]bool{true, false, true, false, false, false, false, true, true, false, false},
	[]bool{true, false, false, true, false, true, true, false, false, false, false},
	[]bool{true, false, false, true, false, false, false, false, true, true, false},
	[]bool{true, false, false, false, false, true, false, true, true, false, false},
	[]bool{true, false, false, false, false, true, false, false, true, true, false},
	[]bool{true, false, true, true, false, false, true, false, false, false, false},
	[]bool{true, false, true, true, false, false, false, false, true, false, false},
	[]bool{true, false, false, true, true, false, true, false, false, false, false},
	[]bool{true, false, false, true, true, false, false, false, false, true, false},
	[]bool{true, false, false, false, false, true, true, false, true, false, false},
	[]bool{true, false, false, false, false, true, true, false, false, true, false},
	[]bool{true, true, false, false, false, false, true, false, false, true, false},
	[]bool{true, true, false, false, true, false, true, false, false, false, false},
	[]bool{true, true, true, true, false, true, true, true, false, true, false},
	[]bool{true, true, false, false, false, false, true, false, true, false, false},
	[]bool{true, false, false, false, true, true, true, true, false, true, false},
	[]bool{true, false, true, false, false, true, true, true, true, false, false},
	[]bool{true, false, false, true, false, true, true, true, true, false, false},
	[]bool{true, false, false, true, false, false, true, true, true, true, false},
	[]bool{true, false, true, true, true, true, false, false, true, false, false},
	[]bool{true, false, false, true, true, true, true, false, true, false, false},
	[]bool{true, false, false, true, true, true, true, false, false, true, false},
	[]bool{true, true, true, true, false, true, false, false, true, false, false},
	[]bool{true, true, true, true, false, false, true, false, true, false, false},
	[]bool{true, true, true, true, false, false, true, false, false, true, false},
	[]bool{true, true, false, true, true, false, true, true, true, true, false},
	[]bool{true, true, false, true, true, true, true, false, true, true, false},
	[]bool{true, true, true, true, false, true, true, false, true, true, false},
	[]bool{true, false, true, false, true, true, true, true, false, false, false},
	[]bool{true, false, true, false, false, false, true, true, true, true, false},
	[]bool{true, false, false, false, true, false, true, true, true, true, false},
	[]bool{true, false, true, true, true, true, false, true, false, false, false},
	[]bool{true, false, true, true, true, true, false, false, false, true, false},
	[]bool{true, true, true, true, false, true, false, true, false, false, false},
	[]bool{true, true, true, true, false, true, false, false, false, true, false},
	[]bool{true, false, true, true, true, false, true, true, true, true, false},
	[]bool{true, false, true, true, true, true, false, true, true, true, false},
	[]bool{true, true, true, false, true, false, true, true, true, true, false},
	[]bool{true, true, true, true, false, true, false, true, true, true, false},
	[]bool{true, true, false, true, false, false, false, false, true, false, false},
	[]bool{true, true, false, true, false, false, true, false, false, false, false},
	[]bool{true, true, false, true, false, false, true, true, true, false, false},
	[]bool{true, true, false, false, false, true, true, true, false, true, false, true, true},
}

func NewCode128BisWriter() gozxing.Writer {
	return NewOneDimensionalCodeWriter(code128BisEncoder{})
}

func (code128BisEncoder) getSupportedWriteFormats() gozxing.BarcodeFormats {
	return gozxing.BarcodeFormats{gozxing.BarcodeFormat_CODE_128}
}

func (code128BisEncoder) encode(input string) ([]bool, error) {
	contents := []rune(input)
	length := len(contents)
	if length < 1 || length > 80 {
		return nil, gozxing.NewWriterException("IllegalArgumentException: "+
		    "Contents length should be between 1 and 80 characters, but got %v", length)
	}

	idxList, err := getCodeIndexList(contents)
	if err != nil {
		return nil, err
	}

	result := NewCode128_BitList()
	sum := 0
	for i, idx := range idxList.GetBytes() {
		if i == 0 {
			sum = int(idx)
		} else {
			sum += i * int(idx)
		}
		result.AddBits(code128_EncodingTable[idx]...)
	}
	sum = sum % 103
	result.AddBits(code128_EncodingTable[sum]...)
	result.AddBits(code128_EncodingTable[code128CODE_STOP]...)

	boolArray := []bool{}
	capacity := result.GetCap()
	for i:=uint32(0); i<capacity; i++ {
		b := result.GetBit(i)
		boolArray = append(boolArray, b)
	}

	return boolArray, nil
}

func shouldUseCTable(nextRunes []rune, curEncoding byte) bool {
	requiredDigits := 4
	if curEncoding == code128CODE_START_C {
		requiredDigits = 2
	}
	if len(nextRunes) < requiredDigits {
		return false
	}
	for i := 0; i < requiredDigits; i++ {
		if i%2 == 0 && nextRunes[i] == code128ESCAPE_FNC_1 {
			requiredDigits++
			if len(nextRunes) < requiredDigits {
				return false
			}
			continue
		}
		if nextRunes[i] < '0' || nextRunes[i] > '9' {
			return false
		}
	}
	return true
}

func tableContainsRune(table string, r rune) bool {
	return strings.ContainsRune(table, r) || r == code128ESCAPE_FNC_1 || r == code128ESCAPE_FNC_2 || r == code128ESCAPE_FNC_3 || r == code128ESCAPE_FNC_4
}

func shouldUseATable(nextRunes []rune, curEncoding byte) bool {
	nextRune := nextRunes[0]
	if !tableContainsRune(bTable, nextRune) || curEncoding == code128CODE_START_A {
		return tableContainsRune(aTable, nextRune)
	}
	if curEncoding == 0 {
		for _, r := range nextRunes {
			if tableContainsRune(abTable, r) {
				continue
			}
			if strings.ContainsRune(aOnlyTable, r) {
				return true
			}
			break
		}
	}
	return false
}

func getCodeIndexList(content []rune) (*code128_BitList, error) {
	result := NewCode128_BitList()

	curEncoding := byte(0)
	for i := 0; i < len(content); i++ {
		if shouldUseCTable(content[i:], curEncoding) {
			if curEncoding != code128CODE_START_C {
				if curEncoding == byte(0) {
					result.AddByte(code128CODE_START_C)
				} else {
					result.AddByte(code128CODE_CODE_C)
				}
				curEncoding = code128CODE_START_C
			}
			if content[i] == code128ESCAPE_FNC_1 {
				result.AddByte(102)
			} else {
				idx := (content[i] - '0') * 10
				i++
				idx = idx + (content[i] - '0')
				result.AddByte(byte(idx))
			}
		} else if shouldUseATable(content[i:], curEncoding) {
			if curEncoding != code128CODE_START_A {
				if curEncoding == byte(0) {
					result.AddByte(code128CODE_START_A)
				} else {
					result.AddByte(code128CODE_CODE_A)
				}
				curEncoding = code128CODE_START_A
			}
			var idx int
			switch content[i] {
			case code128ESCAPE_FNC_1:
				idx = 102
				break
			case code128ESCAPE_FNC_2:
				idx = 97
				break
			case code128ESCAPE_FNC_3:
				idx = 96
				break
			case code128ESCAPE_FNC_4:
				idx = 101
				break
			default:
				idx = strings.IndexRune(aTable, content[i])
				break
			}
			if idx < 0 {
				return nil, gozxing.NewWriterException("Invalid index : %d", idx)
			}
			result.AddByte(byte(idx))
		} else {
			if curEncoding != code128CODE_START_B {
				if curEncoding == byte(0) {
					result.AddByte(code128CODE_START_B)
				} else {
					result.AddByte(code128CODE_CODE_B)
				}
				curEncoding = code128CODE_START_B
			}
			var idx int
			switch content[i] {
			case code128ESCAPE_FNC_1:
				idx = 102
				break
			case code128ESCAPE_FNC_2:
				idx = 97
				break
			case code128ESCAPE_FNC_3:
				idx = 96
				break
			case code128ESCAPE_FNC_4:
				idx = 100
				break
			default:
				idx = strings.IndexRune(bTable, content[i])
				break
			}

			if idx < 0 {
				return nil, gozxing.NewWriterException("invalid index : %d", idx)
			}
			result.AddByte(byte(idx))
		}
	}
	return result, nil
}
