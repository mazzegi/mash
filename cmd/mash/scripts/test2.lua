
local sourceDir = "/home/marting/dev/go/default/bin"
local targetDir = fs:Pwd():Text().."/bin"

print("deploying from", sourceDir, "to", targetDir)

local files = {
    "blog",
    "direct",
    "echo",
    "fpdf"  
    }

mash:Try( fs:Rmall(targetDir), 20 )    
mash:Try( fs:Mkdir(targetDir), 21 )

if mash:Exec("foo"):Failed() then print_error("failed to exec foo") end

res = mash:Exec("foo")
if res:Failed() then
    print_error(res:ErrorText())
end

for i,f in ipairs(files) do
    mash:Try( fs:CopyFile(sourceDir.."/"..f, targetDir.."/"..f), 22 )
end
print("deploying from", sourceDir, "to", targetDir, " ... done")

