package wordclouds

import (
	"bytes"
	"fmt"
	"image/color"
	"image/png"
	"io"

	"github.com/golang/freetype/truetype"
)

type Option func(*Options) error

// Font to use.
func Font(file io.Reader) Option {
	return func(options *Options) error {
		buf := &bytes.Buffer{}
		_, err := buf.ReadFrom(file)
		if err != nil {
			return fmt.Errorf("could not read font file: %w", err)
		}

		f, err := truetype.Parse(buf.Bytes())
		if err != nil {
			return err
		}
		options.Font = f

		return nil
	}
}

// Logo to use.
func Logo(file io.Reader) Option {
	return func(options *Options) error {
		logo, err := png.Decode(file)
		if err != nil {
			return fmt.Errorf("could not read logo: %w", err)
		}
		options.Logo = logo
		return nil
	}
}

// Output file background color
func BackgroundColor(color color.Color) Option {
	return func(options *Options) error {
		options.BackgroundColor = color
		return nil
	}
}

// Colors to use for the words
func Colors(colors []color.Color) Option {
	return func(options *Options) error {
		options.Colors = colors
		return nil
	}
}

// Max font size
func FontMaxSize(max int) Option {
	return func(options *Options) error {
		options.FontMaxSize = max
		return nil
	}
}

// Min font size
func FontMinSize(min int) Option {
	return func(options *Options) error {
		options.FontMinSize = min
		return nil
	}
}

// A list of bounding boxes where words can not be placed.
// See Mask
func MaskBoxes(mask []*Box) Option {
	return func(options *Options) error {
		options.Mask = mask
		return nil
	}
}

func Width(w int) Option {
	return func(options *Options) error {
		options.Width = w
		return nil
	}
}

func Height(h int) Option {
	return func(options *Options) error {
		options.Height = h
		return nil
	}
}

// Place words randomly
func RandomPlacement(do bool) Option {
	return func(options *Options) error {
		options.RandomPlacement = do
		return nil
	}
}

// Draw bounding boxes around words
func Debug() Option {
	return func(options *Options) error {
		options.Debug = true
		return nil
	}
}
