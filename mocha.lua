-- Mocha-9 on boot will setup the console for higher level usage;
-- add useful functions, and do other stuff a BIOS might do

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

setfunc("line", function (x1, y1, x2, y2, color) 
	local x = x2 - x1
	local y = y2 - y1
	local length = math.sqrt((x * x) + (y * y))
	local addx = x / length
	local addy = y / length
	local x = x1
	local y = y1

	for i = length, 0, -1 do
		plot(x, y, color)
		x = x + addx
		y = y + addy
	end
end)

dofile('program.lua')

