// naive way to do it
func TestTiming(t *testing.T) {
	num := 9000000
	start := time.Now().Unix()
	for i := 0; i < num; i++ {
		header := "blahblahblah,countrycode=US,blahblahblah,blahblahblah,blahblahblah,blahblahblah,blahblahblah,blahblahblah,blahblahblah,blahblahblah"
		country, err := GetCountryCodeByIndex(header)
		if err != nil { ...
	}
	stop := time.Now().Unix()
	fmt.Printf("Index %v start=%v stop=%v diff=%v\n",
		num, start, stop, stop-start)
