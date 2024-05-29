package types

// PackRgbppUnlock returns serialized hexadecimal string for the RGB++ unlock witness
//
// The inputsLen and outputsLen are the length of the CKB virtual TX inputs and outputs
// The default value of version is 0
func PackRgbppUnlock(btcTxWithoutWitness string, btcTxProof string, inputsLen, outputsLen uint8, version uint16) string {
	btxTxBytes := *PackBytes(HexToBytes(btcTxWithoutWitness))
	btcTxProofBytes := *PackBytes(HexToBytes(btcTxProof))
	extraData := *ExtraCommitmentDataFromSliceUnchecked([]byte{inputsLen, outputsLen})
	builder := NewRGBPPUnlockBuilder()
	builder.Version(*PackUint16(version))
	builder.ExtraData(extraData)
	builder.BtcTx(btxTxBytes)
	builder.BtcTxProof(btcTxProofBytes)
	rgbppUnlock := builder.Build()
	return BytesToHex(rgbppUnlock.AsSlice())
}
