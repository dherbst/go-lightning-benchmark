import "testing"

// This is the right way to do it
func BenchmarkGetCountryCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		header := "blahblahblah,countrycode=US,blahblahblah,blahblahblah,blahblahblah,blahblahblah,blahblahblah"
		country, err := GetCountryCodeByRegexp(header)
		if err != nil {
			fmt.Printf("Error!")
		}
		if *country != "US" {
			fmt.Printf("didn't get US, got %v", country)
		}
 	}
}
