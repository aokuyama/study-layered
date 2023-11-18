package test

func GenUUIDStinrg(n int) string {
	if n == 1 {
		return "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1"
	}
	if n == 2 {
		return "d833a112-95e8-4042-ab02-ffde48bc874a"
	}
	if n == 3 {
		return "a7ece1ec-f5f3-4af4-b716-995feb91ac51"
	}
	panic("undefined factory")
}
