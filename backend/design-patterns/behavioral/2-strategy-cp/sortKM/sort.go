package sortKM

//kali ini kita coba menggunakan strategy pattern untuk menentukan sort
//ada dua sort yaitu ascending dan descending sort

type Strategy interface {
	Sort([]int)
}

type SortKM struct {
	Strategy Strategy
}

func (s *SortKM) Sort(array []int) {
	s.Strategy.Sort(array)
}

func (s *SortKM) SetStrategy(Strategy Strategy) {
	s.Strategy = Strategy
}