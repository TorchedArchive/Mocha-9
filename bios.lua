-- Mocha-9's "bios" will setup the console for higher level usage,
-- add useful functions and other stuff a BIOS usually does i guess
-- For now it just sets up basic functions and calls a lua file

-- Define clear function
setfunc("clear", function ()
	for x = 143, 0, -1 do
		for y = 80, 0, -1 do
			plot(x, y, 0)
		end
	end
end)

clear()

dofile('hello.lua')
