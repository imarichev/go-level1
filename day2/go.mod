module example.com/day2

go 1.22.3

replace example.com/bottles => ./bottles

replace example.com/readers => ./readers

require (
	example.com/bottles v0.0.0-00010101000000-000000000000 // indirect
	example.com/readers v0.0.0-00010101000000-000000000000 // indirect
)
