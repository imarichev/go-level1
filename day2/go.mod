module example.com/day2

go 1.22.3

replace example.com/bottles => ../lib/bottles
replace example.com/readers => ../lib/readers
replace example.com/triangles => ../lib/triangles

require (
	example.com/bottles v0.0.0-00010101000000-000000000000 // indirect
	example.com/readers v0.0.0-00010101000000-000000000000 // indirect
	example.com/triangles v0.0.0-00010101000000-000000000000 // indirect
)
