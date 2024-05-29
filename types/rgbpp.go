package types

func PackRgbppUnlock(btcTxWithoutWitness string, btcTxProof string, inputsLen, outputsLen uint8) string {
	btxTxBytes := *PackBytes(HexToBytes(btcTxWithoutWitness))
	btcTxProofBytes := *PackBytes((HexToBytes(btcTxProof)))
	extraData := *ExtraCommitmentDataFromSliceUnchecked([]byte{inputsLen, outputsLen})
	version := *Uint16FromSliceUnchecked([]byte{0, 0})
	builder := NewRGBPPUnlockBuilder()
	builder.Version(version)
	builder.ExtraData(extraData)
	builder.BtcTx(btxTxBytes)
	builder.BtcTxProof(btcTxProofBytes)
	rgbppUnlock := builder.Build()
	return BytesToHex(rgbppUnlock.AsSlice())
}
