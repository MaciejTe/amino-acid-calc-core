package calculator

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
	"testing"
)

func PkgDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func TestNewIngredient(t *testing.T) {
	testCaseTable := []struct {
		inputIngredientID     string
		inputUsdaResponseBody []byte
		expectedResult        *Ingredient
		expectedError         error
	} {
		{
			inputIngredientID: "174608",
			expectedResult: &Ingredient{
				ID:            "174608",
				Description:   "Chicken breast, roll, oven-roasted",
				AminoAcids:    AminoAcids{
					Alanine:       0.873,
					Isoleucine:    0.711,
					Leucine:       1.048,
					Valine:        0.702,
					Phenylalanine: 0.562,
					Tryptophan:    0.16,
					Tyrosine:      0.461,
					Asparagine:    0.5,
					Cysteine:      0.5,
					Glutamine:     0.5,
					Methionine:    0.382,
					Serine:        0.52,
					Threonine:     0.597,
					Arginine:      0.928,
					Histidine:     0.419,
					Lysine:        1.167,
					AsparticAcid:  1.301,
					GlutamicAcid:  2.112,
					Glycine:       1.034,
					Proline:       0.75,
				},
				Carbohydrates: 1.79,
				Protein:       14.59,
				Fat:           7.65,
				Kcal:          134,
			},
			expectedError: nil,
		},
	}

	for _, testCase := range testCaseTable {
		inputUsdaResponseBody, err := ioutil.ReadFile(filepath.Join(PkgDir(),"testdata", t.Name()+".golden"))
		if err != nil {
			t.Fatalf("failed reading .golden: %s", err)
		}
		foodDetails, err := NewIngredient(testCase.inputIngredientID, inputUsdaResponseBody)
		assert.Equal(t, testCase.expectedResult, foodDetails, "Actual foodDetails is different than expected one")
		assert.Equal(t, testCase.expectedError, err, "Actual error is different than expected one")
	}
}

func TestNewIngredientNegative(t *testing.T) {
	testCaseTable := []struct {
		inputIngredientID     string
		inputUsdaResponseBody []byte
		expectedResult        *Ingredient
		expectedError         error
	} {
		{
			inputIngredientID: "",
			expectedResult: nil,
			inputUsdaResponseBody: []byte("some_response_body"),
			expectedError: errors.New("NewIngredient received empty ingredientID string"),
		},
		{
			inputIngredientID: "123",
			inputUsdaResponseBody: []byte("asd^(@^)*#*"),
			expectedResult: nil,
			expectedError: errors.New("cannot unmarshal JSON response: invalid character 'a' looking for beginning of value"),
		},
	}

	for _, testCase := range testCaseTable {
		foodDetails, err := NewIngredient(testCase.inputIngredientID, testCase.inputUsdaResponseBody)
		assert.Equal(t, testCase.expectedResult, foodDetails, "Actual foodDetails is different than expected one")
		assert.Equal(t, testCase.expectedError, err, "Actual error is different than expected one")
	}
}