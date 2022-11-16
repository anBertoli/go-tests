package main

/*
#cgo CFLAGS: -g -Wall -I/usr/local/include -I/usr/local/Cellar/python@3.10/3.10.8/Frameworks/Python.framework/Versions/3.10/include/python3.10
#cgo LDFLAGS: -L/usr/local/lib -L/usr/local/opt/python@3.10/Frameworks/Python.framework/Versions/3.10/lib/python3.10/config-3.10-darwin -lpython3.10
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <ctype.h>



#include "./c/glue.c"
*/
import "C"
import (
	"log"
	"unsafe"
)

func init() {
	C.init_python()
}

func passRowsThroughCGO(positions []*CorePosition) []*CorePosition {
	CPositions := make([]C.struct_CorePosition, len(positions))
	for i, pos := range positions {
		CPositions[i] = C.struct_CorePosition{
			RCID:                 C.long(pos.RCID),
			Company:              C.int(pos.Company),
			GranularityID:        C.ulong(pos.GranularityID),
			Seniority:            C.int(pos.Seniority),
			StartDate:            C.int(pos.StartDate),
			EndDate:              C.int(pos.EndDate),
			Weight:               C.float(pos.Weight),
			Multiplicator:        C.float(pos.Multiplicator),
			Inflation:            C.float(pos.Inflation),
			EstimatedUSLogSalary: C.float(pos.EstimatedUSLogSalary),
			FProb:                C.float(pos.FProb),
			MProb:                C.float(pos.MProb),
			WhiteProb:            C.float(pos.WhiteProb),
			BlackProb:            C.float(pos.BlackProb),
			HispanicProb:         C.float(pos.HispanicProb),
			NativeProb:           C.float(pos.NativeProb),
			ApiProb:              C.float(pos.ApiProb),
			MultipleProb:         C.float(pos.MultipleProb),
			Region:               C.int(pos.Region),
			Country:              C.int(pos.Country),
			State:                C.int(pos.State),
			Msa:                  C.int(pos.Msa),
			MappedRole:           C.int(pos.MappedRole),
			Soc6dTitle:           C.int(pos.Soc6dTitle),
			HighestDegree:        C.int(pos.HighestDegree),
			GenderR:              C.float(pos.GenderR),
			Gender:               C.int(pos.Gender),
			EthnicityR:           C.float(pos.EthnicityR),
			Ethnicity:            C.int(pos.Ethnicity),
		}
	}

	pyFunc := `
def update(row): 
	row['RCID'] = 4
	print(row)
	return row
`

	log.Println("Entering CGO.")
	var ret *C.struct_CorePosition = C.py_process(
		C.CString(pyFunc),
		&CPositions[0],
		C.int(len(CPositions)),
	)
	log.Println("Exiting CGO.")

	// TODO NOTE: in C code only RCID and weight are processed.
	p := unsafe.Pointer(ret)
	var data []C.struct_CorePosition = unsafe.Slice((*C.struct_CorePosition)(p), len(positions))
	log.Println("Data coming from Python: ")
	for i, d := range data {
		log.Printf("%d --> %+v\n", i, d)
	}

	// convert back
	positionsRet := make([]*CorePosition, 0, 10000)
	for _, x := range data {
		positionsRet = append(positionsRet, &CorePosition{
			RCID: int64(x.RCID),
		})
	}

	return positionsRet
}
