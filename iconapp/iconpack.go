package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
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
	inputPtr := flag.String("i", "", "Input image path (Required)")
	outputPtr := flag.String("o", "", "Output path (File for single, Folder for split)")
	sizePtr := flag.String("s", "", "Target size 'WIDTHxHEIGHT' (e.g. 300x300)")
	tolerancePtr := flag.Int("t", 20, "Background color tolerance (0-255)")
	minPtr := flag.Int("m", 10, "Minimum pixel dimension for icons (width or height)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Icon pack - Sprite Sheet Extractor\n\n")
		fmt.Fprintf(os.Stderr, "USAGE:\n")
		fmt.Fprintf(os.Stderr, "  Split sheet:    iconpack -i sheet.png -o ./out_folder -s 256x256\n")
		fmt.Fprintf(os.Stderr, "  Single image:   iconpack -i icon.png -o ./processed.png -s 512x512\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *inputPtr == "" || *outputPtr == "" {
		flag.Usage()
		return
	}

	outExt := filepath.Ext(*outputPtr)
	isSingleMode := outExt != ""

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
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.Draw(img, img.Bounds(), src, bounds.Min, draw.Src)

	fmt.Println("Step 1: Removing background and denoising...")
	floodFillAlpha(img, *tolerancePtr)

	// NEW: Remove alpha noise (stray pixels with very low opacity)
	denoise(img)

	if isSingleMode {
		fmt.Printf("Step 2: Processing single image to '%s'...\n", *outputPtr)
		os.MkdirAll(filepath.Dir(*outputPtr), 0755)
		finalImg := image.Image(tightCrop(img))
		if *sizePtr != "" {
			tw, th := parseSize(*sizePtr)
			finalImg = resizeAndCenter(finalImg.(*image.RGBA), tw, th)
		}
		saveImage(*outputPtr, finalImg)
		fmt.Println("Finished!")
		return
	}

	fmt.Printf("Step 2: Splitting icons into folder '%s'...\n", *outputPtr)
	os.MkdirAll(*outputPtr, 0755)

	visited := make([]bool, w*h)
	iconCount := 0

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			idx := y*w + x
			_, _, _, a := img.At(x, y).RGBA()

			// Start finding an island only if pixel is significantly visible
			if a > 5000 && !visited[idx] {
				points, minP, maxP := findIsland(img, x, y, visited)

				iconW := (maxP.X - minP.X) + 1
				iconH := (maxP.Y - minP.Y) + 1

				// Strictly filter out small clumps
				// If the total area is tiny OR both dimensions are smaller than min
				if (iconW < *minPtr && iconH < *minPtr) || len(points) < (*minPtr*2) {
					continue
				}

				tempImg := image.NewRGBA(image.Rect(0, 0, iconW, iconH))
				for _, p := range points {
					tempImg.Set(p.X-minP.X, p.Y-minP.Y, img.At(p.X, p.Y))
				}

				tightImg := tightCrop(tempImg)
				var finalImg image.Image = tightImg

				if *sizePtr != "" {
					tw, th := parseSize(*sizePtr)
					finalImg = resizeAndCenter(tightImg, tw, th)
				}

				outName := fmt.Sprintf("icon_%03d.png", iconCount)
				saveImage(filepath.Join(*outputPtr, outName), finalImg)
				fmt.Printf("  Saved %s (%dx%d raw)\n", outName, tightImg.Bounds().Dx(), tightImg.Bounds().Dy())
				iconCount++
			}
		}
	}
	fmt.Printf("\nFinished! Extracted %d icons.\n", iconCount)
}

// denoise removes very faint pixels that often cause "clumps"
func denoise(img *image.RGBA) {
	for i := 0; i < len(img.Pix); i += 4 {
		// If alpha is less than ~5% (12/255), make it fully transparent
		if img.Pix[i+3] < 12 {
			img.Pix[i] = 0
			img.Pix[i+1] = 0
			img.Pix[i+2] = 0
			img.Pix[i+3] = 0
		}
	}
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

		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				if dx == 0 && dy == 0 {
					continue
				}
				nx, ny := p.X+dx, p.Y+dy
				if nx >= 0 && nx < w && ny >= 0 && ny < h {
					idx := ny*w + nx
					_, _, _, a := img.At(nx, ny).RGBA()
					// Threshold 5000 (~7% alpha) to ignore noise clumps
					if a > 5000 && !globalVisited[idx] {
						globalVisited[idx] = true
						queue = append(queue, Point{nx, ny})
					}
				}
			}
		}
	}
	return points, minP, maxP
}

func tightCrop(img *image.RGBA) *image.RGBA {
	bounds := img.Bounds()
	minX, minY, maxX, maxY := bounds.Max.X, bounds.Max.Y, bounds.Min.X, bounds.Min.Y

	hasPixels := false
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			_, _, _, a := img.At(x, y).RGBA()
			if a > 5000 { // Match denoise threshold
				hasPixels = true
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}

	if !hasPixels {
		return img
	}

	rect := image.Rect(minX, minY, maxX+1, maxY+1)
	sub := img.SubImage(rect).(*image.RGBA)
	out := image.NewRGBA(image.Rect(0, 0, rect.Dx(), rect.Dy()))
	draw.Draw(out, out.Bounds(), sub, rect.Min, draw.Src)
	return out
}

func resizeAndCenter(src *image.RGBA, tw, th int) image.Image {
	resized := imaging.Fit(src, tw, th, imaging.Lanczos)
	dst := image.NewRGBA(image.Rect(0, 0, tw, th))
	bx := resized.Bounds()
	startX := (tw - bx.Dx()) / 2
	startY := (th - bx.Dy()) / 2
	draw.Draw(dst, image.Rect(startX, startY, startX+bx.Dx(), startY+bx.Dy()), resized, image.Point{0, 0}, draw.Src)
	return dst
}

func colorDist(r1, g1, b1, r2, g2, b2 uint32) float64 {
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
