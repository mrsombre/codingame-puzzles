package main

type BFSNode struct {
	Cell *Cell

	Distance int
}

func BFS(grid *Grid) string {
	visited := make(map[int]bool)
	queue := make([]*BFSNode, 0, grid.Levels*grid.Width*grid.Height)

	visited[grid.Start.ID] = true
	queue = append(queue, &BFSNode{
		Cell:     grid.Start,
		Distance: 0,
	})

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Cell == grid.End {
			return intToStr(current.Distance)
		}

		for _, cell := range current.Cell.Directions {
			if visited[cell.ID] {
				continue
			}
			if cell.IsWall {
				continue
			}

			visited[cell.ID] = true
			node := &BFSNode{
				Cell:     cell,
				Distance: current.Distance + 1,
			}
			queue = append(queue, node)
		}
	}

	return "NO PATH"
}
