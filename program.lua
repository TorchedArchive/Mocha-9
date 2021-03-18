--rect(1, 1, 142, 1, 1)
x = 180 / 2
y = 120 / 2

function Brew()
	clear()
	plot(x, y, 1)
	
	if button(0) then
		y = y - 1
	end
	if button(1) then
		y = y + 1
	end
	if button(2) then
		x = x - 1
	end
	if button(3) then
		x = x + 1
	end
end
