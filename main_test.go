package main

import (
  "Fibbonaci/fibbonaci"
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestFibbonaciWithSuccess(t *testing.T) {
  assert := assert.New(t)
  {
  	array := fibbonaci.Fibbonaci(6)
  	assert.Equal(array, []int{0,1,1,2,3,5}, "they should be equal")
  }
  {
  	array := fibbonaci.Fibbonaci(1)
  	assert.Equal(array, []int{0}, "they should be equal")
  }
}

func TestFibbonaciWithFailure(t *testing.T) {
	assert := assert.New(t)
	{
        	array := fibbonaci.Fibbonaci(-1)
        	assert.Nil(array)
	}

	{
                array := fibbonaci.Fibbonaci(1)
                assert.NotEqual(array, []int{1}, "they should not be equal")
        }

}

