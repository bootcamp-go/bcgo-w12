package intro

func HeightCheck(height float64) string {
	if height < 150.0 {
		return "Can't ride the roller coaster"
	}
	return "Can ride the roller coaster"
}