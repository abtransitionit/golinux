package user

import (
	"os"
	"os/user"
	"runtime"
	"testing"
)

// Name: TestCanBeSudo
//
// Notes:
//
// - Since we cannot change the user's privileges during a test run. the test is simple
// - test must be played by a sudo user otherwise the test will fail. a root user or a non-sudo user are not a sudo user.
// - this test will only pass if the result from TestCanBeSudoAndIsNotRoot is true
func TestCanBeSudoAndIsNotRoot(t *testing.T) {
	// This function works only on Linux (cf. function's Notes)
	if runtime.GOOS == "darwin" {
		t.Skip("Skipping test on Darwin; this function is designed for Linux.")
	}
	// Test setup: Determine the expected result based on the current user.
	isRoot := os.Geteuid() == 0
	expectedResult := !isRoot // true for non-root, false for root

	// Run the function under test.
	obtainedResult, err := CanBeSudoAndIsNotRoot() // true if the current user is sudo

	// Assertion for expected error state.
	if err != nil {
		// we do not expect an error here
		t.Fatalf("unexpected error when checking for sudo privileges: %v", err)
	}

	// obtained is not expected
	if obtainedResult != expectedResult {
		currentUser, err := user.Current()
		if err != nil {
			// we do not expect an error here
			t.Fatalf("failed to get current user info: %v", err)
		}
		// If 'expectedResult' is true, it means the test was run by a sudo-user: the test must fail: anormal situation
		if expectedResult {
			t.Errorf("Test Failed: Expected obtainedResult to be true for non-root user %s, but got false.", currentUser.Username)
		} else {
			// If 'expectedResult' is false, it means the test was run by either a non-sudo or a root user: the test must fail: anormal situation
			t.Errorf("Test Failed: Expected obtainedResult to be false for root user %s, but got true.", currentUser.Username)
		}
	} else {
		// obtained is expected
		// If 'expectedResult' is true, the user is a non-root, sudo-user: the test must pass: normal situation
		if expectedResult {
			t.Logf("✅ Test Passed: User is not root, and obtainedResult correctly returned true.")
		} else {
			// If 'expectedResult' is false, the user is a root or a non-sudo user:: the test must pass: normal situation
			t.Logf("✅ Test Passed: User is root, and obtainedResult correctly returned false.")
		}
	}
}
