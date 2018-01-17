package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type point struct {
	x       float64
	y       float64
	cluster int
}

func kmpp(points []point, k int) {
	centers := kpp_seeds(points, k)

	//print_points(centers)

	for {
		for i := 0; i < len(centers); i++ {
			centers[i].x = 0
			centers[i].y = 0
			centers[i].cluster = 0
		}
		for i := 0; i < len(points); i++ {
			centers[points[i].cluster].x += points[i].x
			centers[points[i].cluster].y += points[i].y
		}

		for i := 0; i < len(centers); i++ {
			centers[i].x *= float64(1.0) / float64(len(centers))
			centers[i].y *= float64(1.0) / float64(len(centers))
		}

		var changes bool
		for i := 0; i < len(points); i++ {
			index, _ := nearest_distance(points[i], centers)
			if index != points[i].cluster {
				points[i].cluster = index
				changes = true

			}
		}
		if !changes {
			break
		}

	}
	print_points(points)
}

func generate_random_points(count int, radius int) []point {
	rand.Seed(time.Now().UnixNano())

	//	points := make([]point)
	var points []point

	for i := 0; i < count; i++ {
		r := rand.Float64() * float64(radius)
		var p point
		angle := rand.Float64() * 2 * math.Pi
		p.x = r * math.Cos(angle)
		p.y = r * math.Sin(angle)
		points = append(points, p)
		//points[i] = p
	}
	//print_points(points)
	return points
}

func kpp_seeds(points []point, k int) []point {
	// 初始化k个中心
	centers := make([]point, k)
	dist := make([]float64, len(points))

	for i := 0; i < k; i++ {
		sum := 0.0
		for j := 0; j < len(points); j++ {
			_, d := nearest_distance(points[j], centers)
			dist[j] = d
			sum += dist[j]
		}

		sum *= rand.Float64()
		for j := 0; j < len(points); j++ {
			sum -= dist[j]
			if sum <= 0 {
				centers[i] = points[j]
				points[j].cluster = i
				break
			}
		}
	}
	return centers
}

// 返回当前点距离cluster 中心最近的点
func nearest_distance(p point, center []point) (int, float64) {
	index := 0
	min_dist := sqrt_distance(p, center[0])
	for i := 0; i < len(center); i++ {
		d := sqrt_distance(p, center[i])
		if min_dist > d {
			min_dist = d
			index = i
		}
	}
	return index, min_dist
}

func sqrt_distance(p1 point, p2 point) float64 {
	distance := (p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)
	return math.Sqrt(distance)
}

func print_points(points []point) {
	fmt.Println("---begin---\n")
	for i := 0; i < len(points); i++ {
		fmt.Printf("(x=%f,y=%f,c=%d)\n", points[i].x, points[i].y, points[i].cluster)
	}
	fmt.Println("---end----\n")
}

func main() {
	var point_count, clusters int
	point_count = 10
	clusters = 2

	points := generate_random_points(point_count, 10)
	kmpp(points, clusters)

}
