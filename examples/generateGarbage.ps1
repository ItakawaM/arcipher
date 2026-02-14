$text = "This is a benchmark line for my CLI tool.`n" * 1000
$path = ".\garbage.txt"
# This loops until the file reaches 10GB
while ((Get-Item $path).Length -lt 10GB) {
    $text | Out-File -FilePath $path -Append -NoNewline
}