local start = time:Now()

local sourceDir = fs:Pwd():Value().."/source"
local targetDir = fs:Pwd():Value().."/bin"

local files = {
    "blog",
    "direct",
    "echo",
    "dot"  
    }

print("preparing source: ", sourceDir)
mash:Try( fs:Mkdir(sourceDir), 21 )
for i,f in ipairs(files) do
    fs:WriteFile(sourceDir.."/"..f, f.."_data")
end

print("deploying from", sourceDir, "to", targetDir)

mash:Try( fs:RemoveAll(targetDir), 20 )    
mash:Try( fs:Mkdir(targetDir), 21 )

for i,f in ipairs(files) do
    mash:Try( fs:CopyFile(sourceDir.."/"..f, targetDir.."/"..f), 22 )
end
print("deploying from", sourceDir, "to", targetDir, " ... done")
local dur = time:Since(start)
print(time:FormatDuration(dur))

