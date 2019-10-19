
local sourceDir = "/home/marting/dev/go/default/bin"
local targetDir = fs:Pwd():Text().."/bin"

print("deploying from", sourceDir, "to", targetDir)

local files = {
    "blog",
    "direct",
    "echo",
    "fpdf"  
    }

mash:Try( fs:Mkdir(targetDir), 21 )

mash:Try( mash:Exec("foo"), 21 )

for i,f in ipairs(files) do
    mash:Try( fs:CopyFile(sourceDir.."/"..f, targetDir.."/"..f), 22 )
end
print("deploying from", sourceDir, "to", targetDir, " ... done")

