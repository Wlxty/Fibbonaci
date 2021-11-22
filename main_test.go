package main

import (
  "Fibbonaci/Data"
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestFibbonaciWithSuccess(t *testing.T) {
  assert := assert.New(t)
  {
  	fibbonaci := Data.Fibbonaci(6)
  	assert.Equal(fibbonaci, []int{0,1,1,2,3,5}, "they should be equal")
  }
  {
  	fibbonaci := Data.Fibbonaci(1)
  	assert.Equal(fibbonaci, []int{0}, "they should be equal")
  }
}

func TestFibbonaciWithFailure(t *testing.T) {
	assert := assert.New(t)
	{
        	fibbonaci := Data.Fibbonaci(-1)
        	assert.Nil(fibbonaci)
	}

	{
                fibbonaci := Data.Fibbonaci(1)
                assert.NotEqual(fibbonaci, []int{1}, "they should not be equal")
        }

}

