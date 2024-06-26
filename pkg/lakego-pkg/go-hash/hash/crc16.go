package hash

import (
    "github.com/deatil/go-hash/crc16/x25"
    "github.com/deatil/go-hash/crc16/modbus"
)

// CRC16_X25 哈希值
func (this Hash) CRC16_X25() Hash {
    data := x25.Checksum(string(this.data))

    this.data, this.Error = this.HexDecode(data)

    return this
}

// CRC16_Modbus 哈希值
func (this Hash) CRC16_Modbus() Hash {
    data := modbus.Checksum(string(this.data))

    this.data, this.Error = this.HexDecode(data)

    return this
}
