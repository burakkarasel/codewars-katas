package kata

func HowMuchILoveYou(i int) string {
	return []string{"not at all", "I love you", "a little", "a lot", "passionately", "mady"}[i%6]
}

// obsolote
/* func HowMuchILoveYou(i int) (r string) {
	i %= 6
	switch i {
	case 1:
		r = "I love you"
	case 2:
		r = "a little"
	case 3:
		r = "a lot"
	case 4:
		r = "passionately"
	case 5:
		r = "madly"
	default:
		r = "not at all"
	}
	return
} */
