--rect(1, 1, 142, 1, 1)
num = 0

print(num, 1, 1, 1)

function Brew()
	if button(0) then
		clear()
		num = num + 1
		print(num, 1, 1, 1)
	end
end
