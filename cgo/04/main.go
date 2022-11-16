package main

func main() {
	passRowsThroughCGO([]*CorePosition{
		{RCID: 3, Weight: 2.56},
		{RCID: 9, Weight: 99.1},
		{RCID: 99, Weight: 213.213},
	})
}

type CorePosition struct {
	RCID                 int64
	Company              int
	GranularityID        uint64
	Seniority            int16
	StartDate            int16
	EndDate              int16
	Weight               float32
	Multiplicator        float32
	Inflation            float32
	EstimatedUSLogSalary float32
	FProb                float32
	MProb                float32
	WhiteProb            float32
	BlackProb            float32
	HispanicProb         float32
	NativeProb           float32
	ApiProb              float32
	MultipleProb         float32
	Region               int16
	Country              int16
	State                int16
	Msa                  int16
	MappedRole           int16
	Soc6dTitle           int16
	HighestDegree        int16
	GenderR              float32
	Gender               int16
	EthnicityR           float32
	Ethnicity            int16
}
