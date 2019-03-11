package main

//
//import (
//	"github.com/stretchr/testify/assert"
//	"io/ioutil"
//	"testing"
//)
//
//func TestBrokenChain(t *testing.T) {
//	assert := assert.New(t)
//	tests := []struct {
//		name              string
//		writeDateYearDays int // year days to write to the test chain
//		writeDateYear     int // year to write to the test chain
//		expectBroken      bool
//	}{
//		{
//			name:              "today",
//			writeDateYearDays: today().yearDays,
//			writeDateYear:     today().year,
//			expectBroken:      false,
//		},
//		{
//			name:              "yesterday",
//			writeDateYearDays: yesterday().yearDays,
//			writeDateYear:     yesterday().year,
//			expectBroken:      false,
//		},
//		{
//			name:              "two_days_ago",
//			writeDateYearDays: yesterday().yearDays - 1,
//			writeDateYear:     yesterday().year,
//			expectBroken:      true,
//		},
//		{
//			name:              "invalid_year",
//			writeDateYearDays: today().yearDays,
//			writeDateYear:     2000,
//			expectBroken:      true,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tmp, err := ioutil.TempFile("", "")
//			assert.Nil(err)
//			assert.Nil(tmp.Close())
//
//			// use file as config database
//			db := dbFile{path: tmp.Name()}
//			assert.Nil(db.Load())
//
//			c := chain{&db}
//			assert.Nil(c.db.WriteLastDate(tt.writeDateYearDays, tt.writeDateYear))
//
//			// load db to pull file values into db struct as the chain references the db
//			assert.Nil(db.Load())
//
//			assert.EqualValues(tt.expectBroken, c.Broken())
//		})
//	}
//}
