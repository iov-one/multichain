package acala

import (
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/renproject/multichain/api/address"
)

type addressDecoder struct{}

type addressEncoder struct{}

// NewAddressDecoder returns the default AddressDecoder for Substract chains. It
// uses the Bitcoin base58 alphabet to decode the string, and interprets the
// result as a 2-byte address type, 32-byte array, and 1-byte checksum.
func NewAddressDecoder() address.Decoder {
	return addressDecoder{}
}

func NewAddressEncoder() address.Encoder {
	return addressEncoder{}
}

// DecodeAddress the string using the Bitcoin base58 alphabet. If the string
// does not a 2-byte address type, 32-byte array, and 1-byte checksum, then an
// error is returned.
func (addressDecoder) DecodeAddress(addr address.Address) (address.RawAddress, error) {
	data := base58.Decode(string(addr))
	if len(data) != 35 {
		return address.RawAddress([]byte{}), fmt.Errorf("expected 35 bytes, got %v bytes", len(data))
	}
	return address.RawAddress(data), nil
}

func (addressEncoder) EncodeAddress(rawAddr address.RawAddress) (address.Address, error) {
	data := base58.Encode(rawAddr)
	return address.Address(data), nil
}
