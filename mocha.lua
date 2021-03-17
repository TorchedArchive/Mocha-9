-- Mocha-9's "bios" will setup the console for higher level usage,
-- add useful functions and other stuff a BIOS usually does i guess
-- For now it just sets up basic functions and calls a lua file

-- Define clear function
setfunc("clear", function ()
	for x = 143, 0, -1 do
		for y = 80, 0, -1 do
			plot(x, y, i1)
		end
	end
end)

clear()

setfunc("rect", function (x, y, w, h, color)
	local dx = x
	local dy = y
	
	for hi = h, 0, -1 do
		for wi = w, 0, -1 do
			plot(dx, dy, color)
			dx = dx + 1
		end
		dy = dy + 1
		dx = x
	end
end)


dofile('program.lua')
