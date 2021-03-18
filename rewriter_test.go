package rewriter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRewriter(t *testing.T) {
	rewriter := NewRewriter([]Rule{
		{"/old", "/new"},
	})

	assert.Equal(t, "/api", rewriter.MustRewrite("/api"))
	assert.Equal(t, "/new", rewriter.MustRewrite("/old"))
	assert.Equal(t, "/new/123", rewriter.MustRewrite("/new/123"))
}

func TestRewriterSPA(t *testing.T) {
	rewriter := NewRewriter([]Rule{
		{"/*", "/index.html"},
	})

	assert.Equal(t, "/index.html", rewriter.MustRewrite("/"))
	assert.Equal(t, "/index.html", rewriter.MustRewrite("/posts"))
	assert.Equal(t, "/index.html", rewriter.MustRewrite("/posts/123"))
}

func TestRewriterAPI(t *testing.T) {
	rewriter := NewRewriter([]Rule{
		{"/api/*", "https://api.site.com/$1"},
	})

	assert.Equal(t, "https://api.site.com/", rewriter.MustRewrite("/api/"))
	assert.Equal(t, "https://api.site.com/posts", rewriter.MustRewrite("/api/posts"))
	assert.Equal(t, "https://api.site.com/posts/123", rewriter.MustRewrite("/api/posts/123"))
}
