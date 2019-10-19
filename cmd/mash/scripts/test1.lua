print("hi from lua", 3)
print_error("something went terribly wrong", 42.42)

numargs = mash:NumArgs()
print("num-args:", numargs)
for i=0,numargs-1 do
    print("arg", i, mash:Arg(i))
end