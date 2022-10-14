package txtax

type ByTimestamp []Transaction

func (a ByTimestamp) Len() int { return len(a) }
func (a ByTimestamp) Less(i, j int) bool {
	return a[i].TimeStamp < a[j].TimeStamp
}
func (a ByTimestamp) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
