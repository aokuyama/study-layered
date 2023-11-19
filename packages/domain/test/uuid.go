package test

func GenUUIDStinrg(n int) string {
	if n == 1 {
		return "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1"
	}
	if n == 2 {
		return "d833a112-95e8-4042-ab02-ffde48bc874a"
	}
	if n == 3 {
		return "550e8400-e29b-41d4-a716-446655440000"
	}
	panic("undefined factory")
}
