package mailago

import "testing"

// TestFormatHostPortTest test the formatHostPort function
func TestFormatHostPortTest(t *testing.T) {
  type formatHostPortTest struct {
    host     string //input 1
    port     int    //input 2
    expected string // expected result
  }

  var formatHostPortTests = []formatHostPortTest{
    {"localhost", 3031, "localhost:3031"},
    {"localhost", 3030, "localhost:3030"},
    {"127.0.0.1", 8080, "127.0.0.1:8080"},
    {"web.domain.com", 80, "web.domain.com:80"},
  }

  for _, tt := range formatHostPortTests {
    actual := formatHostPort(tt.host, tt.port)
    if actual != tt.expected {
      t.Errorf("formatHostPort(%s, %d): expected %s, actual %v", tt.host, tt.port, tt.expected, actual)
    }
  }
}

// TestValidateEmailInput tests the email payload validator
func TestValidateEmailInput(t *testing.T) {
  testCases := []struct {
    Input     *EmailPayload
    ShouldErr bool
  }{
    {
      // Tests that it fails on missing FROM field in payload
      Input: &EmailPayload{
        To:      "user@gmail.com",
        Subject: "Subject",
        Body:    "Body",
      },
      ShouldErr: true,
    },
    {
      // Tests that it fails on missing TO field in payload
      Input: &EmailPayload{
        From:    "user@gmail.com",
        Subject: "Subject",
        Body:    "Body",
      },
      ShouldErr: true,
    },
    {
      // Tests that it fails on missing Subject field in payload
      Input: &EmailPayload{
        To:   "user@gmail.com",
        From: "user@gmail.com",
        Body: "Body",
      },
      ShouldErr: true,
    },
    {
      // Tests that it fails on missing Body field in payload
      Input: &EmailPayload{
        To:      "user@gmail.com",
        From:    "user@gmail.com",
        Subject: "Subject",
      },
      ShouldErr: true,
    },
    {
      // Tests that it fails on improper FROM email address
      Input: &EmailPayload{
        From:    "derp",
        To:      "user@gmail.com",
        Subject: "Subject",
        Body:    "Body",
      },
      ShouldErr: true,
    },
    {
      // Tests that it fails on improper TO email address
      Input: &EmailPayload{
        From:    "derp@gmail.com",
        To:      "user",
        Subject: "Subject",
        Body:    "Body",
      },
      ShouldErr: true,
    },
    // More test cases
  }
  for _, testCase := range testCases {
    err := validateEmailInput(testCase.Input)
    if testCase.ShouldErr && err == nil {
      t.Errorf("Expected validation error, but got none.")
    }
  }
}
