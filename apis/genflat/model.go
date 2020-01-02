package genflat

func (a *LoginRequestT) Byte() []byte {
	b := flatbuffers.NewBuilder(0)
	b.Finish(LoginRequestPack(b, a))
	return b.FinishedBytes()
}

func ByteLoginRequestT(b []byte) *LoginRequestT {
	return GetRootAsLoginRequest(b, 0).UnPack()
}

func LoginRequestTBuild(in *LoginRequestT) *flatbuffers.Builder {
	b := flatbuffers.NewBuilder(0)
	b.Reset()
	b.Finish(LoginRequestPack(b, in))
	return b
}
