package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("December 10th, 2018")

	input := []byte(`-30509,  41062/ 3, -4
-50990, -20297/ 5,  2
-50968,  51288/ 5, -5
-51003, -40754/ 5,  4
-51011,  10375/ 5, -1
 30840,  41058/-3, -4
-30546,  51284/ 3, -5
-10090,  41065/ 1, -4
-51024, -40755/ 5,  4
 30832, -40757/-3,  4
 20597, -20297/-2,  2
-50964,  51292/ 5, -5
-20303,  20608/ 2, -2
 10351,  51292/-1, -5
 30792,  41065/-3, -4
-51008,  30833/ 5, -3
 41023,  51287/-4, -5
 41067, -50987/-4,  5
-50966,  41060/ 5, -4
 30792, -30527/-3,  3
 51246,  10375/-5, -1
 20586,  30831/-2, -3
-40789, -10071/ 4,  1
 51248,  30833/-5, -3
-10071, -50987/ 1,  5
 10378,  30831/-1, -3
 20581, -50978/-2,  5
-10116, -50984/ 1,  5
-20335, -10073/ 2,  1
-50968, -30529/ 5,  3
 51267, -20299/-5,  2
-51016, -20306/ 5,  2
-10084,  41057/ 1, -4
 20589, -30526/-2,  3
 30849,  41065/-3, -4
 10383,  41065/-1, -4
 41059,  51286/-4, -5
 41022, -50987/-4,  5
 20600,  30838/-2, -3
-20335, -50981/ 2,  5
-20285,  30838/ 2, -3
-10073,  10377/ 1, -1
 41069, -50978/-4,  5
-30567, -50987/ 3,  5
-20291, -30524/ 2,  3
-10076,  51287/ 1, -5
-40788,  51287/ 4, -5
 30808,  10378/-3, -1
-30562, -40759/ 3,  4
-51016,  30833/ 5, -3
 20565, -20302/-2,  2
-51019,  41056/ 5, -4
 30802,  10378/-3, -1
 51307,  41059/-5, -4
 51265, -50982/-5,  5
 41080, -50980/-4,  5
 30813,  10378/-3, -1
 30809, -10078/-3,  1
 10346, -20299/-1,  2
-51016,  10381/ 5, -1
-10081, -20297/ 1,  2
 51265, -40755/-5,  4
 51303,  20602/-5, -2
-50992, -40755/ 5,  4
 41059,  51291/-4, -5
 10370, -30529/-1,  3
-30549, -10071/ 3,  1
-20308, -20297/ 2,  2
 10341,  51287/-1, -5
 41079, -40756/-4,  4
 10339,  51292/-1, -5
-20283,  30829/ 2, -3
 41067, -30530/-4,  3
-10105, -30531/ 1,  3
-40765, -30527/ 4,  3
 10356,  51288/-1, -5
 51278, -10073/-5,  1
-51016, -20298/ 5,  2
-40780, -30528/ 4,  3
-30522, -50980/ 3,  5
-40776,  51289/ 4, -5
 20597, -50984/-2,  5
-10076, -10079/ 1,  1
 51294, -50978/-5,  5
 41051,  30836/-4, -3
 41028, -10075/-4,  1
-50973,  10384/ 5, -1
 51246, -20304/-5,  2
 30843,  20611/-3, -2
-40797,  10378/ 4, -1
 51254,  20611/-5, -2
 41052,  51292/-4, -5
 10343, -20297/-1,  2
-30570,  20609/ 3, -2
-40753,  51284/ 4, -5
 20581,  41061/-2, -4
-30514, -20304/ 3,  2
 51246, -30527/-5,  3
 10378,  30836/-1, -3
 51290,  20603/-5, -2
 51247,  51292/-5, -5
-51008, -20299/ 5,  2
-30553,  30834/ 3, -3
 51257, -30526/-5,  3
 41060,  10380/-4, -1
-30511,  20602/ 3, -2
 10367, -50978/-1,  5
 41051,  41056/-4, -4
-40777,  20603/ 4, -2
 30853,  10380/-3, -1
-10112,  41065/ 1, -4
-40757,  20609/ 4, -2
 51303,  20606/-5, -2
-51022,  41065/ 5, -4
-10106,  10378/ 1, -1
 41024,  30829/-4, -3
-50965, -50987/ 5,  5
 51307, -50985/-5,  5
 51270, -30532/-5,  3
-51014,  30835/ 5, -3
 10354, -30526/-1,  3
-40786,  41058/ 4, -4
-10095, -30524/ 1,  3
-40797, -10072/ 4,  1
 20621,  20606/-2, -2
 10346, -10079/-1,  1
 20569, -20297/-2,  2
 51282,  41065/-5, -4
 10398, -40751/-1,  4
 20567, -20306/-2,  2
-20282, -20305/ 2,  2
-40795,  10379/ 4, -1
 51303,  51287/-5, -5
-10107,  41061/ 1, -4
 10338, -20305/-1,  2
 20584,  51283/-2, -5
-20306, -20297/ 2,  2
 41027,  10379/-4, -1
 41040, -30526/-4,  3
 51254, -20301/-5,  2
 30840, -10076/-3,  1
 10370,  10376/-1, -1
 30817,  51292/-3, -5
-50968, -30527/ 5,  3
 30812, -30528/-3,  3
-30557,  20611/ 3, -2
 20575,  41062/-2, -4
-40763, -10070/ 4,  1
 30851,  41060/-3, -4
 30795, -20297/-3,  2
 30848,  20605/-3, -2
-40757, -10078/ 4,  1
 10362, -10070/-1,  1
-51003, -10074/ 5,  1
 20581, -20298/-2,  2
 30792, -50981/-3,  5
-50991, -30524/ 5,  3
-20319,  41062/ 2, -4
-10072, -50986/ 1,  5
 30804,  41057/-3, -4
-20303, -50984/ 2,  5
 41047,  51292/-4, -5
 41038, -20306/-4,  2
-20284,  41065/ 2, -4
 41067, -20298/-4,  2
-30509,  51290/ 3, -5
-50965, -50983/ 5,  5
 20621,  30837/-2, -3
 30796, -10079/-3,  1
-50999, -40751/ 5,  4
-20339,  10379/ 2, -1
 20589, -10077/-2,  1
-40754, -30526/ 4,  3
 41027, -30530/-4,  3
 51278, -50979/-5,  5
 51304, -30533/-5,  3
-30533,  30838/ 3, -3
 10338,  51286/-1, -5
-40747,  51292/ 4, -5
-10071, -20297/ 1,  2
 41061, -50981/-4,  5
-30511, -50987/ 3,  5
-50964,  10375/ 5, -1
-30546,  41059/ 3, -4
 51248,  41065/-5, -4
 20589, -20306/-2,  2
 10338,  30837/-1, -3
 10338, -40753/-1,  4
-40789, -30531/ 4,  3
-40789, -50984/ 4,  5
 51294, -30528/-5,  3
-20303,  41063/ 2, -4
-10071,  10384/ 1, -1
-30529,  30833/ 3, -3
 41075, -20298/-4,  2
 41078, -30529/-4,  3
 51262,  30831/-5, -3
 20613,  20602/-2, -2
-50968,  41063/ 5, -4
-30512, -40756/ 3,  4
 51304, -50987/-5,  5
 30800, -20297/-3,  2
 10370, -20303/-1,  2
-20282,  30837/ 2, -3
-20335, -50986/ 2,  5
-50964, -30524/ 5,  3
-20319, -40756/ 2,  4
 10350,  51291/-1, -5
-51024,  41064/ 5, -4
-40748,  20611/ 4, -2
-51000,  30837/ 5, -3
-10097,  30834/ 1, -3
 20613, -20305/-2,  2
 51258,  41057/-5, -4
-40757, -50982/ 4,  5
 10388, -20297/-1,  2
-51003, -20297/ 5,  2
-10068,  41062/ 1, -4
 10359, -20304/-1,  2
-10073,  51290/ 1, -5
 30840, -20304/-3,  2
 30800,  30834/-3, -3
 10394,  30838/-1, -3
-40794, -20306/ 4,  2
-51016, -30528/ 5,  3
-10067,  10384/ 1, -1
 41067, -10071/-4,  1
-10076, -50982/ 1,  5
-40741,  10376/ 4, -1
 51247, -20302/-5,  2
-10076, -20305/ 1,  2
-50976, -20302/ 5,  2
 41036,  30830/-4, -3
-50976, -40758/ 5,  4
-10100, -40753/ 1,  4
-30530,  51288/ 3, -5
 20613,  51286/-2, -5
-20311,  10380/ 2, -1
 20570, -40760/-2,  4
 41040, -40756/-4,  4
-30530,  51292/ 3, -5
 20613,  10381/-2, -1
-20287, -10074/ 2,  1
 30849, -10079/-3,  1
 41027,  41056/-4, -4
 51287,  51287/-5, -5
-50971, -30524/ 5,  3
-40776, -50980/ 4,  5
 10395,  51283/-1, -5
 10354, -20300/-1,  2
 10378,  20603/-1, -2
 20597, -50985/-2,  5
-30570, -50987/ 3,  5
 20609,  20610/-2, -2
 41062, -50980/-4,  5
 41064, -10079/-4,  1
 30792, -30533/-3,  3
 10362, -30530/-1,  3
-20282,  10380/ 2, -1
 30853,  41057/-3, -4
 41027, -40756/-4,  4
-20322,  51285/ 2, -5
-30558, -30532/ 3,  3
 20623, -20302/-2,  2
 10395,  10379/-1, -1
 51295,  41065/-5, -4
-10116,  30833/ 1, -3
-40741,  51286/ 4, -5
 10362,  10380/-1, -1
 30816, -10071/-3,  1
 30803,  20604/-3, -2
 30834,  20605/-3, -2
-30534,  41065/ 3, -4
-50984,  20608/ 5, -2
 51294, -50981/-5,  5
 10371, -40751/-1,  4
 41059, -50978/-4,  5
-40797, -10078/ 4,  1
 10354,  20605/-1, -2
-50992, -30525/ 5,  3
-30554, -50983/ 3,  5
 30795, -30529/-3,  3
-10055, -30526/ 1,  3
-50992, -40759/ 5,  4
-20295, -40755/ 2,  4
-20322, -10071/ 2,  1
 41060,  51288/-4, -5
-51011,  30838/ 5, -3
 30819, -20297/-3,  2
 10359,  30837/-1, -3
 30801,  10379/-3, -1
 10394, -50982/-1,  5
-30546, -20303/ 3,  2
-30562, -20304/ 3,  2
-20339,  41056/ 2, -4
-50964, -20302/ 5,  2
 10343,  20611/-1, -2
 10370, -20301/-1,  2
-20343,  20606/ 2, -2
 30832,  10383/-3, -1
 20565,  20611/-2, -2
 30797,  30838/-3, -3
 41080,  30831/-4, -3
 20597, -10076/-2,  1
-50995,  41065/ 5, -4
-40793, -10079/ 4,  1
 10378, -30531/-1,  3
 30848,  10375/-3, -1
-10076, -40758/ 1,  4
-40736, -30530/ 4,  3
-40776, -20301/ 4,  2
-10068,  41056/ 1, -4
-30529,  10379/ 3, -1
-30570,  10377/ 3, -1
-20295, -10075/ 2,  1
 51247, -40760/-5,  4
-40773, -40755/ 4,  4
 30809, -50982/-3,  5
-30551,  51283/ 3, -5
-10116,  10376/ 1, -1
-10115, -40751/ 1,  4
-20311, -20298/ 2,  2
-10098,  10375/ 1, -1
 30849, -50978/-3,  5`)
	// 	input := []byte(` 9,  1/ 0,  2
	//  7,  0/-1,  0
	//  3, -2/-1,  1
	//  6, 10/-2, -1
	//  2, -4/ 2,  2
	// -6, 10/ 2, -2
	//  1,  8/ 1, -1
	//  1,  7/ 1,  0
	// -3, 11/ 1, -2
	//  7,  6/-1, -1
	// -2,  3/ 1,  0
	// -4,  3/ 2,  0
	// 10, -3/-1,  1
	//  5, 11/ 1, -2
	//  4,  7/ 0, -1
	//  8, -2/ 0,  1
	// 15,  0/-2,  0
	//  1,  6/ 1,  0
	//  8,  9/ 0, -1
	//  3,  3/-1,  1
	//  0,  5/ 0, -1
	// -2,  2/ 2,  0
	//  5, -2/ 1,  2
	//  1,  4/ 2,  1
	// -2,  7/ 2, -2
	//  3,  6/-1, -1
	//  5,  0/ 1,  0
	// -6,  0/ 2,  0
	//  5,  9/ 1, -2
	// 14,  7/-2,  0
	// -3,  6/ 2, -1`)

	// numSeconds := int64(3)
	work := make(chan bool)
	// for s := int64(0); s <= numSeconds; s++ {
	// go buildImageAndSave(string(input), 10220, work)
	// go buildImageAndSave(string(input), 10221, work)
	// go buildImageAndSave(string(input), 10222, work)
	// go buildImageAndSave(string(input), 10223, work)
	// go buildImageAndSave(string(input), 10224, work)
	// go buildImageAndSave(string(input), 10225, work)
	// go buildImageAndSave(string(input), 10226, work)
	go buildImageAndSave(string(input), 10227, work)
	// go buildImageAndSave(string(input), 10228, work)
	// go buildImageAndSave(string(input), 10229, work)
	// go buildImageAndSave(string(input), 10230, work)
	// go buildImageAndSave(string(input), 10231, work)
	// go buildImageAndSave(string(input), 10232, work)
	// go buildImageAndSave(string(input), 10233, work)
	// go buildImageAndSave(string(input), 10234, work)
	// go buildImageAndSave(string(input), 10235, work)
	// go buildImageAndSave(string(input), 10236, work)
	// go buildImageAndSave(string(input), 10237, work)
	// go buildImageAndSave(string(input), 10238, work)
	// go buildImageAndSave(string(input), 10239, work)
	// go buildImageAndSave(string(input), 10240, work)
	// go buildImageAndSave(string(input), 3, work)

	// <-work
	// <-work
	// <-work
	// <-work
	// <-work
	// <-work
	// <-work
	<-work
	// <-work
	// <-work
	// <-work
	// <-work
	// <-work
	// <-work
	// <-work
	// <-work
	// <-work
	// <-work
	// <-work
	// <-work
	// <-work
	// <-work

	// }
	fmt.Printf("Result (part 1): %v\n", 0)
}

func buildImageAndSave(input string, secs int, res chan bool) {
	lines := strings.Split(input, "\n")
	sky := map[int]map[int]string{}
	minX := 0
	maxX := 0
	minY := 0
	maxY := 0
	for i := 0; i < len(lines); i++ {
		lines[i] = strings.Replace(lines[i], " ", "", -1)
		posX, _ := strconv.ParseInt(strings.Split(strings.Split(lines[i], "/")[0], ",")[0], 10, 64)
		posY, _ := strconv.ParseInt(strings.Split(strings.Split(lines[i], "/")[0], ",")[1], 10, 64)

		velX, _ := strconv.ParseInt(strings.Split(strings.Split(lines[i], "/")[1], ",")[0], 10, 64)
		velY, _ := strconv.ParseInt(strings.Split(strings.Split(lines[i], "/")[1], ",")[1], 10, 64)

		posX += velX * int64(secs)
		posY += velY * int64(secs)

		if int(posX) < minX {
			minX = int(posX)
		}
		if int(posX) > maxX {
			maxX = int(posX)
		}
		if int(posY) < minY {
			minY = int(posY)
		}
		if int(posY) > maxY {
			maxY = int(posY)
		}

		// fmt.Printf("posX: %v\n", posX)
		// fmt.Printf("posY: %v\n", posY)
		// fmt.Printf("velX: %v\n", velX)
		// fmt.Printf("velY: %v\n", velY)

		if _, ok := sky[int(posX)]; !ok {
			sky[int(posX)] = map[int]string{}
		}
		sky[int(posX)][int(posY)] = "#"
	}

	// fmt.Printf("minX: %v maxX: %v minY: %v maxY: %v\n", minX, maxX, minY, maxY)
	// fmt.Printf("%+v\n", sky)

	// Because...
	maxX, maxY = maxY, maxX

	upleft := image.Point{0, 0}
	bottomright := image.Point{maxY + 1, maxX + 1}
	img := image.NewGray(image.Rectangle{upleft, bottomright})
	progress := 0
	newProgress := 0
	for j := minX; j <= maxX; j++ {
		for i := minY; i <= maxY; i++ {
			if sky[i][j] == "" {
				img.Set(i, j, color.White)
				// fmt.Printf(".")
			} else {
				img.Set(i, j, color.Black)
				// fmt.Printf("#")
			}
			newProgress = (((j - minX) * (maxY - minY)) * 100) / ((maxX - minX) * (maxY - minY))
			if newProgress != progress {
				// fmt.Printf("%d %d%%\n", secs, newProgress)
				progress = newProgress
			}
		}
		// fmt.Printf("\n")
	}
	f, err := os.Create(fmt.Sprintf("sky_at_%d.jpeg", secs))
	if err != nil {
		log.Fatal("can't create file:", err)
	}
	var opts jpeg.Options
	opts.Quality = 100
	err = jpeg.Encode(f, img, &opts)
	if err != nil {
		log.Fatal("can't save image:", err)
	}

	res <- true
}
