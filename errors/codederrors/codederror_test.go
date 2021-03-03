package codederrors_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Alter17Ego/generic-app/errors/codederrors"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateErrorWithRightCode(t *testing.T) {
	errorCode := "internal.app.error"
	errorMessage := "something went wrong"
	codedError := codederrors.New(errorCode, errorMessage)
	assert.Equal(t, errorCode, codedError.Code)
}

func TestShouldCreateErrorWithRightMessage(t *testing.T) {
	errorCode := "internal.app.error"
	errorMessage := "something went wrong"
	codedError := codederrors.New(errorCode, errorMessage)
	assert.Equal(t, errorMessage, codedError.Message)
}

func TestCodedErrorShouldReturnErrorMessageAndCode(t *testing.T) {
	errorCode := "internal.app.error"
	errorMessage := "something went wrong"
	expectedError := fmt.Sprintf("Error : %s - %s", errorCode, errorMessage)
	codedError := codederrors.New(errorCode, errorMessage)
	assert.Equal(t, expectedError, codedError.Error())
}

func TestShouldWrapExistingErrorWithCodedError(t *testing.T) {
	errorCode := "authentication.error"
	errorMessage := "invalid username or password"
	err := errors.New("Bad Credential")
	codedError := codederrors.Wrap(err, errorCode, errorMessage)
	assert.Equal(t, err, codedError.Err)
}
