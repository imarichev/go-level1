module example.com/day4

go 1.22.3

replace example.com/invoice => ../lib/invoice

replace example.com/readers => ../lib/readers

require example.com/invoice v0.0.0-00010101000000-000000000000

require example.com/readers v0.0.0-00010101000000-000000000000 // indirect
