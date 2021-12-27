package day16

import (
	"advent/rkutil"
	"encoding/hex"
	"math"
)

type BITSPacket struct {
	Version            uint32
	TypeID             uint32
	LengthTypeID       uint32
	LiteralValues      []uint32
	LengthOfSubPackets uint32
	NumberOfSubPackets uint32
	SubPackets         []*BITSPacket
}

func (packet BITSPacket) VersionSum() uint32 {
	result := packet.Version
	for _, sub := range packet.SubPackets {
		result += sub.VersionSum()
	}
	return result
}

func (packet BITSPacket) Value() uint64 {
	switch packet.TypeID {
	case PacketTypeSum:
		result := uint64(0)
		for _, sub := range packet.SubPackets {
			result += sub.Value()
		}
		return result
	case PacketTypeProduct:
		if len(packet.SubPackets) == 1 {
			return packet.SubPackets[0].Value()
		} else {
			result := uint64(1)
			for _, sub := range packet.SubPackets {
				result *= sub.Value()
			}
			return result
		}
	case PacketTypeMinimum:
		result := uint64(math.MaxUint64)
		for _, sub := range packet.SubPackets {
			result = rkutil.MinUint64(result, sub.Value())
		}
		return result
	case PacketTypeMaximum:
		result := uint64(0)
		for _, sub := range packet.SubPackets {
			result = rkutil.MaxUint64(result, sub.Value())
		}
		return result
	case PacketTypeLiteral:
		if len(packet.LiteralValues) == 0 || len(packet.LiteralValues) > 16 {
			rkutil.UnexpectedCodePath()
		}
		result := uint64(0)
		for _, literal := range packet.LiteralValues {
			if result != 0 {
				result = result<<4 | uint64(literal)
			} else {
				result = uint64(literal)
			}
		}
		return result
	case PacketTypeGt:
		rkutil.Ensure(len(packet.SubPackets) == 2, "2 packets")
		if packet.SubPackets[0].Value() > packet.SubPackets[1].Value() {
			return 1
		} else {
			return 0
		}
	case PacketTypeLt:
		rkutil.Ensure(len(packet.SubPackets) == 2, "2 packets")
		if packet.SubPackets[0].Value() < packet.SubPackets[1].Value() {
			return 1
		} else {
			return 0
		}
	case PacketTypeEqual:
		rkutil.Ensure(len(packet.SubPackets) == 2, "2 packets")
		if packet.SubPackets[0].Value() == packet.SubPackets[1].Value() {
			return 1
		} else {
			return 0
		}
	}
	rkutil.UnexpectedCodePath()
	return 0
}

func (packet BITSPacket) SubPacketCount() int {
	result := 0
	for _, sub := range packet.SubPackets {
		result++
		result += sub.SubPacketCount()
	}
	return result
}

const (
	PacketTypeSum     = 0
	PacketTypeProduct = 1
	PacketTypeMinimum = 2
	PacketTypeMaximum = 3
	PacketTypeLiteral = 4
	PacketTypeGt      = 5
	PacketTypeLt      = 6
	PacketTypeEqual   = 7
)

type Bytes []byte

func (bytes Bytes) Value(at int, len int) uint32 {
	rkutil.Ensure(len < 32, "unsupported length")
	var result uint32
	byteIndex := at / 8
	at = at % 8
	for len > 0 {
		bitsFromThisByte := rkutil.MinInt(8-at, len)
		shift := 8 - bitsFromThisByte - at
		mask := byte(0xff) >> byte(8-bitsFromThisByte)
		cur := bytes[byteIndex] >> shift
		cur &= mask
		result = result<<bitsFromThisByte | uint32(cur)
		len = len - bitsFromThisByte
		if len > 0 {
			at = 0
			byteIndex++
		}
	}
	return result
}

func readFirstLine(filename string) string {
	lines := rkutil.MustLines(filename)
	return lines[0]
}

type Reader struct {
	Bytes Bytes
	Idx   int
}

func (rdr *Reader) ReadPacketHeader() BITSPacket {
	packet := BITSPacket{}
	packet.Version = rdr.Bytes.Value(rdr.Idx, 3)
	rdr.Idx += 3
	packet.TypeID = rdr.Bytes.Value(rdr.Idx, 3)
	rdr.Idx += 3
	return packet
}

func (rdr *Reader) ReadLiteralPacket(packet *BITSPacket) int {
	lenRead := 0
	for {
		next := rdr.Bytes.Value(rdr.Idx, 1)
		rdr.Idx += 1
		packet.LiteralValues = append(packet.LiteralValues, rdr.Bytes.Value(rdr.Idx, 4))
		rdr.Idx += 4
		lenRead += 5
		if next == 0 {
			return lenRead
		}
	}
}

func (rdr *Reader) ReadOperatorPacket(packet *BITSPacket) int {
	startIdx := rdr.Idx
	packet.LengthTypeID = rdr.Bytes.Value(rdr.Idx, 1)
	rdr.Idx += 1
	if packet.LengthTypeID == 0 {
		packet.LengthOfSubPackets = rdr.Bytes.Value(rdr.Idx, 15)
		rdr.Idx += 15
		lenToRead := int(packet.LengthOfSubPackets)
		for lenToRead > 0 {
			subPacket, lenRead := rdr.ReadPacket()
			lenToRead -= lenRead
			packet.SubPackets = append(packet.SubPackets, &subPacket)
		}
	} else {
		packet.NumberOfSubPackets = rdr.Bytes.Value(rdr.Idx, 11)
		rdr.Idx += 11
		packetsToRead := int(packet.NumberOfSubPackets)
		for packetsToRead > 0 {
			subPacket, _ := rdr.ReadPacket()
			packet.SubPackets = append(packet.SubPackets, &subPacket)
			packetsToRead--
		}
	}
	return rdr.Idx - startIdx
}

func (rdr *Reader) ReadPacket() (BITSPacket, int) {
	cur := rdr.ReadPacketHeader()
	if cur.TypeID == PacketTypeLiteral {
		l := rdr.ReadLiteralPacket(&cur)
		return cur, l + 6
	} else {
		l := rdr.ReadOperatorPacket(&cur)
		return cur, l + 6
	}
}

func SolveEasierString(s string) int {
	bytes, err := hex.DecodeString(s)
	rkutil.Ensure(err == nil, "no error")
	rdr := Reader{}
	rdr.Bytes = bytes
	cur, _ := rdr.ReadPacket()
	return int(cur.VersionSum())
}

func SolveEasier() int {
	return SolveEasierString(readFirstLine("input.txt"))
}

func SolveHarderString(s string) int {
	bytes, err := hex.DecodeString(s)
	rkutil.Ensure(err == nil, "no error")
	rdr := Reader{}
	rdr.Bytes = bytes
	cur, _ := rdr.ReadPacket()
	return int(cur.Value())
}

func SolveHarder() int {
	return SolveHarderString(readFirstLine("input.txt"))
}
