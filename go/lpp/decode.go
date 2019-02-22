package lpp

import "errors"

func Decode(payload []byte, port int) (*Uplink, error) {
	up := make(Uplink)
	if port == 0 {
		return nil, errors.New("port 0 is reserved for LoRaWAN MAC commands")
	}

	err := up.UnmarshalBinary(payload)
	if err != nil {
		return nil, err
	}

	return &up, nil
}
