package wordclouds

import (
	"bufio"
	"encoding/json"
	"image/color"
	"image/png"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWordcloud_Draw(t *testing.T) {
	colorsRGBA := []color.RGBA{
		{0x1b, 0x1b, 0x1b, 0xff},
		{0x48, 0x48, 0x4B, 0xff},
		{0x59, 0x3a, 0xee, 0xff},
		{0x65, 0xCD, 0xFA, 0xff},
		{0x70, 0xD6, 0xBF, 0xff},
	}
	colors := make([]color.Color, 0)
	for _, c := range colorsRGBA {
		colors = append(colors, c)
	}
	// Load words
	f, err := os.Open("testdata/input.json")
	if err != nil {
		t.Error(err)
		return
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	dec := json.NewDecoder(reader)
	inputWords := make(map[string]int)
	err = dec.Decode(&inputWords)
	assert.NoError(t, err)

	t0 := time.Now()

	boxes, err := Mask(
		"testdata/mask.png",
		2048,
		2048,
		color.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 0,
		})
	if err != nil {
		t.Error(err)
		return
	}

	font, err := os.Open("testdata/Roboto-Regular.ttf")
	if err != nil {
		t.Error(err)
		return
	}
	defer font.Close()

	t.Logf("Mask loading took %v", time.Since(t0))
	t0 = time.Now()
	w, err := NewWordcloud(inputWords,
		Font(font),
		FontMaxSize(300),
		FontMinSize(30),
		Colors(colors),
		MaskBoxes(boxes),
		Height(2048),
		Width(2048))
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Wordcloud init took %v", time.Since(t0))
	t0 = time.Now()

	img := w.Draw()

	t.Logf("Drawing took %v", time.Since(t0))
	t0 = time.Now()

	outputFile, err := os.Create("test.png")
	assert.NoError(t, err)

	// Encode takes a writer interface and an image interface
	// We pass it the File and the RGBA
	png.Encode(outputFile, img)

	// Don't forget to close files
	outputFile.Close()
}
