package cookie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetName(t *testing.T) {
	testCases := []struct {
		desc               string
		cookieName         string
		backendName        string
		expectedCookieName string
		cookiePath         string
		expectedCookiePath string
	}{
		{
			desc:               "with backend name, without cookie name, without cookie path",
			cookieName:         "",
			backendName:        "/my/BACKEND-v1.0~rc1",
			expectedCookieName: "_5f7bc",
			cookiePath:         "",
			expectedCookiePath: "/",
		},
		{
			desc:               "without backend name, with cookie name, without cookie path",
			cookieName:         "/my/BACKEND-v1.0~rc1",
			backendName:        "",
			expectedCookieName: "_my_BACKEND-v1.0~rc1",
			cookiePath:         "",
			expectedCookiePath: "/",
		},
		{
			desc:               "without backend name, with cookie name, with cookie path",
			cookieName:         "/my/BACKEND-v1.0~rc1",
			backendName:        "",
			expectedCookieName: "_my_BACKEND-v1.0~rc1",
			cookiePath:         "/my",
			expectedCookiePath: "/my",
		},
		{
			desc:               "with backend name, with cookie name, with cookie path",
			cookieName:         "containous",
			backendName:        "treafik",
			expectedCookieName: "containous",
			cookiePath:         "/日本語",
			expectedCookiePath: "/",
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			cookieName := GetName(test.cookieName, test.backendName)

			assert.Equal(t, test.expectedCookieName, cookieName)

			cookiePath := GetPath(test.cookiePath)

			assert.Equal(t, test.expectedCookiePath, cookiePath)
		})
	}
}

func Test_sanitizeName(t *testing.T) {
	testCases := []struct {
		desc         string
		srcName      string
		expectedName string
	}{
		{
			desc:         "with /",
			srcName:      "/my/BACKEND-v1.0~rc1",
			expectedName: "_my_BACKEND-v1.0~rc1",
		},
		{
			desc:         "some chars",
			srcName:      "!#$%&'()*+-./:<=>?@[]^_`{|}~",
			expectedName: "!#$%&'__*+-._________^_`_|_~",
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			cookieName := sanitizeName(test.srcName)

			assert.Equal(t, test.expectedName, cookieName, "Cookie name")
		})
	}
}

func TestGenerateName(t *testing.T) {
	cookieName := GenerateName("containous")

	assert.Len(t, "_8a7bc", 6)
	assert.Equal(t, "_8a7bc", cookieName)
}

func Test_sanitizePath(t *testing.T) {
	testCases := []struct {
		desc         string
		srcPath      string
		expectedPath string
	}{
		{
			desc:         "non-empty",
			srcName:      "/my",
			expectedPath: "/my",
		},
		{
			desc:         "valid chars",
			srcPath:      " !\"#$%&'()*+,-./0123456789:<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~",
			expectedPath: " !\"#$%&'()*+,-./0123456789:<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~",
		},
		{
			desc:         "invalid ctl",
			srcPath:      "/a\x19z",
			expectedPath: "/",
		},
		{
			desc:         "invalid char",
			srcPath:      "/a;z",
			expectedPath: "/",
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			cookiePath := sanitizeName(test.srcPath)

			assert.Equal(t, test.expectedPath, cookiePath, "Cookie path")
		})
	}
}
