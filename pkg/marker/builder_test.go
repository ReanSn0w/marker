package marker_test

import (
	"bytes"
	"context"
	"log"
	"testing"

	"github.com/ReanSn0w/marker/pkg/marker"
)

func TestBuilder(t *testing.T) {
	page := marker.NewPage(
		marker.Title(nil, "Test"),
		marker.P(marker.Attributes(
			marker.SetStyle("color: red;"),
			marker.Set("some", ""),
		),
			"Hello, World!",
		),
	)

	var buf bytes.Buffer
	page.WritePage(context.Background(), &buf)

	if buf.String() != `<!DOCTYPE html>
<html>
<head>
<title>
Test
</title>
</head>
<body>
<p style="color: red;" some>
Hello, world!
</p>
</body>
</html>` {
		log.Println(buf.String())
		t.Fatal("Test failed")
	}
}
