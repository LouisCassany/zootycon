package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg" // Support decoding JPEGs
	"image/png"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
)

type Point struct{ X, Y int }

func main() {
	// 1. Define Flags
	inputPtr := flag.String("i", "", "Input image path (Required)")
	outputPtr := flag.String("o", "", "Output folder path (Required)")
	sizePtr := flag.String("s", "", "Target size 'WIDTHxHEIGHT' (e.g. 64x64)")
	tolerancePtr := flag.Int("t", 20, "Background color tolerance (0-255)")
	minPtr := flag.Int("m", 10, "Minimum pixel size for icons")

	// 2. Customize the Help Menu
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Icon pack - A tool to split and clean sprite sheets\n\n")
		fmt.Fprintf(os.Stderr, "USAGE:\n")
		fmt.Fprintf(os.Stderr, "  iconpack -i <input> -o <output_dir> [options]\n\n")
		fmt.Fprintf(os.Stderr, "OPTIONS:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nDETAILED TIPS:\n")
		fmt.Fprintf(os.Stderr, "  -t (Tolerance):\n")
		fmt.Fprintf(os.Stderr, "     Higher values (30-50) are better for JPEGs with 'fuzzy' white backgrounds.\n")
		fmt.Fprintf(os.Stderr, "     Lower values (5-10) are better for clean, pixel-perfect PNGs.\n\n")
		fmt.Fprintf(os.Stderr, "  -m (Min Size) - THE NOISE FILTER:\n")
		fmt.Fprintf(os.Stderr, "     This flag prevents the tool from saving tiny 'garbage' files.\n")
		fmt.Fprintf(os.Stderr, "     • Use -m 15: If you are getting tiny specks or dots saved as files.\n")
		fmt.Fprintf(os.Stderr, "     • Use -m 2:  If you have very small, thin icons that are being skipped.\n")
		fmt.Fprintf(os.Stderr, "     • Tip: Usually, setting this to 10-20%% of your expected icon size is ideal.\n\n")
		fmt.Fprintf(os.Stderr, "EXAMPLES:\n")
		fmt.Fprintf(os.Stderr, "  # Basic split and crop:\n")
		fmt.Fprintf(os.Stderr, "  iconpack -i sheet.png -o ./out\n\n")
		fmt.Fprintf(os.Stderr, "  # Split and resize icons to 64x64 squares, ignoring noise under 20px:\n")
		fmt.Fprintf(os.Stderr, "  iconpack -i sheet.jpg -o ./out -s 64x64 -m 20 -t 40\n")
	}

	flag.Parse()

	// 3. Validation
	if *inputPtr == "" || *outputPtr == "" {
		flag.Usage()
		return
	}

	// Load Image
	file, err := os.Open(*inputPtr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer file.Close()

	src, _, err := image.Decode(file)
	if err != nil {
		fmt.Printf("Error decoding image: %v\n", err)
		return
	}

	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	// Convert to mutable RGBA
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.Draw(img, img.Bounds(), src, bounds.Min, draw.Src)

	fmt.Println("Processing background...")
	floodFillAlpha(img, *tolerancePtr)

	fmt.Println("Splitting icons...")
	visited := make([]bool, w*h)
	iconCount := 0

	os.MkdirAll(*outputPtr, 0755)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			idx := y*w + x
			_, _, _, a := img.At(x, y).RGBA()

			// If pixel is not transparent and not visited
			if a > 0 && !visited[idx] {
				points, minP, maxP := findIsland(img, x, y, visited)

				iconW := (maxP.X - minP.X) + 1
				iconH := (maxP.Y - minP.Y) + 1

				if iconW < *minPtr || iconH < *minPtr {
					continue
				}

				// Crop the icon
				iconImg := image.NewRGBA(image.Rect(0, 0, iconW, iconH))
				for _, p := range points {
					iconImg.Set(p.X-minP.X, p.Y-minP.Y, img.At(p.X, p.Y))
				}

				var finalImg image.Image = iconImg

				// Resize and Center if requested
				if *sizePtr != "" {
					tw, th := parseSize(*sizePtr)
					finalImg = resizeAndCenter(iconImg, tw, th)
				}

				// Save
				outName := fmt.Sprintf("icon_%03d.png", iconCount)
				saveImage(filepath.Join(*outputPtr, outName), finalImg)
				fmt.Printf("  Saved %s (%dx%d)\n", outName, iconW, iconH)
				iconCount++
			}
		}
	}
	fmt.Printf("Finished! Extracted %d icons.\n", iconCount)
}

func floodFillAlpha(img *image.RGBA, tolerance int) {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	seed := img.At(0, 0)
	sr, sg, sb, _ := seed.RGBA()

	queue := []Point{{0, 0}}
	visited := make([]bool, w*h)
	visited[0] = true

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		cr, cg, cb, _ := img.At(p.X, p.Y).RGBA()
		if colorDist(sr, sg, sb, cr, cg, cb) <= float64(tolerance) {
			img.Set(p.X, p.Y, color.Transparent)

			for _, d := range []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				nx, ny := p.X+d.X, p.Y+d.Y
				if nx >= 0 && nx < w && ny >= 0 && ny < h && !visited[ny*w+nx] {
					visited[ny*w+nx] = true
					queue = append(queue, Point{nx, ny})
				}
			}
		}
	}
}

func findIsland(img *image.RGBA, startX, startY int, globalVisited []bool) ([]Point, Point, Point) {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	queue := []Point{{startX, startY}}
	globalVisited[startY*w+startX] = true

	var points []Point
	minP := Point{startX, startY}
	maxP := Point{startX, startY}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		points = append(points, p)

		// Corrected boundary tracking logic
		if p.X < minP.X {
			minP.X = p.X
		}
		if p.Y < minP.Y {
			minP.Y = p.Y
		}
		if p.X > maxP.X {
			maxP.X = p.X
		}
		if p.Y > maxP.Y {
			maxP.Y = p.Y
		}

		// Check 8-way neighbors (including diagonals) to keep icons connected
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				if dx == 0 && dy == 0 {
					continue
				}
				nx, ny := p.X+dx, p.Y+dy
				if nx >= 0 && nx < w && ny >= 0 && ny < h {
					idx := ny*w + nx
					_, _, _, a := img.At(nx, ny).RGBA()
					if a > 0 && !globalVisited[idx] {
						globalVisited[idx] = true
						queue = append(queue, Point{nx, ny})
					}
				}
			}
		}
	}
	return points, minP, maxP
}

func resizeAndCenter(src *image.RGBA, tw, th int) image.Image {
	// Fit maintaining aspect ratio
	resized := imaging.Fit(src, tw, th, imaging.Lanczos)

	dst := image.NewRGBA(image.Rect(0, 0, tw, th))

	bx := resized.Bounds()
	pos := image.Pt((tw-bx.Dx())/2, (th-bx.Dy())/2)

	draw.Draw(dst, image.Rectangle{pos, pos.Add(bx.Size())}, resized, image.Point{}, draw.Over)
	return dst
}

func colorDist(r1, g1, b1, r2, g2, b2 uint32) float64 {
	// Convert 0-65535 to 0-255
	dr := float64(r1>>8) - float64(r2>>8)
	dg := float64(g1>>8) - float64(g2>>8)
	db := float64(b1>>8) - float64(b2>>8)
	return math.Sqrt(dr*dr + dg*dg + db*db)
}

func parseSize(s string) (int, int) {
	parts := strings.Split(strings.ToLower(s), "x")
	if len(parts) != 2 {
		return 0, 0
	}
	w, _ := strconv.Atoi(parts[0])
	h, _ := strconv.Atoi(parts[1])
	return w, h
}

func saveImage(path string, img image.Image) {
	f, err := os.Create(path)
	if err != nil {
		return
	}
	defer f.Close()
	png.Encode(f, img)
}
