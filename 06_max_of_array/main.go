package main

import (
	"fmt"
)

func main() {
	arrData := []int64{22377, -25602, 28181, -63400, -78260, -76569, -19943, 23247, -59298, 23896, 44547, 27960, 73970, 5421, -30161, -74303, -5491, 1262, -21066, -39744, -36479, -78483, 71507, -70474, 38700, 53178, 67350, -53749, 71045, -34953, -79099, 13422, 19445, 29081, 30022, 21185, 32511, -69920, -35568, 53212, 33975, -71020, 1173, 27946, 14400, 51011, 33642, -71090, -27726, -67424, -30835, 15794, -65908, -39327, 25320, 52792, -66149, 12670, 79043, -75104, 57717, 79943, 18318, -2838, 29024, -31660, -61652, -18464, -21581, -17221, -45251, -67605, -8242, 35921, 40340, -73841, 6933, -6018, -64932, 59206, 6558, -15767, -4999, 20650, 24905, -59679, -6558, 38755, 32991, -7514, 43651, 10708, -7570, -18031, -72129, -58545, 30309, -53782, 2990, -71271, 8996, 37738, -58877, -79245, -6340, 61462, -73086, -79407, -24555, -58018, 59798, 62002, 6214, -25200, 2652, -48881, -4880, 76094, 69874, -51889, -11420, 33525, 38819, 61009, -64505, 46689, -77536, 45803, 72906, 5454, 54531, 1903, -36807, 75654, 2658, 36852, 57116, 9571, 37445, -47439, 31552, 17244, -65437, -42233, 72043, 17215, -11115, -12837, 13309, -21240, 15274, -78110, -67714, -25907, 62899, -52220, -59217, 65363, 73583, -66310, -9183, 48115, 15592, 34010, 43769, -61749, -9138, 20885, 27821, -51693, 53446, -20626, 45551, 68009, 17140, 37594, 5224, -73974, -55242, -61466, -15215, 40032, -59576, -2929, -65874, -76677, 24850, -45092, 68685, 18434, -31403, -20497, -13451, 64189, -66487, -49682, -77560, 4374, 51202, 30261, 32681, 24649, -70365, -1767, 12658, 26775, -44173, -62117, 32800, -19415, -43584, -62414, -59383, -23161, 14656, -45258, -19838, -40493, -10350, -31153, 57940, 38247, 28349, -35511, 22436, 41861, -5193, 24876, -33764, -33991, -24863, 78917, 70658, -15228, -2850, 3316, -68453, 32976, 21198, 44347, -66438, 57614, 61932, -45822, -45546, -3412, -11080, 14615, 36095, 58569, 63461, 14035, 16816, 11811, 58524, -40747, -26328, -26669, 64128, 19908, 19340, -40734, 18825, 9998, 24037, -64024, -66685, 35584, 48952, 34513, -69, 62513, 12127, -18137, -63309, 46580, 58451, 5610, -18804, 14546, -15820, -35342, -51419, -79004, 56468, -72894, -39752, -49859, -19564, -55623, 50049, 79776, -16358, -11126, 9774, -72321, 4850, 23088, 43263, -26198}
	var min, max int64
	min = 9999999999
	max = -9999999999
	for _, val := range arrData {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	fmt.Println(max, "", min)
}
