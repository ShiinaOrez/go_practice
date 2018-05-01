package tempconv

func CToF (c C) F {
	return F(c*9/5+32)
}

func FToC (f F) C {
	return C((f-32)*5/9)
}

func CToK (c C) K {
	return K(c+273.15)
}

func FToK (f F) K {
	return K(FToC(F)+273.15)
}

func KToC (k K) C {
	return C(k-273.15)
}

func KToF (k K) F {
	return F(CToF(KToC(K)))
}