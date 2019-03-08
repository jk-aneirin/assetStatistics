function Get-MacAddress { 
    $colItems = get-wmiobject -class "Win32_NetworkAdapterConfiguration" -computername $env:computername |Where{$_.IpEnabled -Match "True"}  
     
    foreach ($objItem in $colItems) {  
     
        $objItem | select MACAddress  
     
    } 
} 

function Get-MonitorSn {
    $Monitor = gwmi wmimonitorid -namespace root\wmi -ComputerName $env:computername

    $Monitor | %{
    $psObject = New-Object PSObject
    $psObject | Add-Member NoteProperty SerialNumber ""
    $psObject.SerialNumber = ($_.SerialNumberID -ne 0 | %{[char]$_}) -join ""
    $psObject
    }
}

function Get-HostSrvTag {
    wmic csproduct get identifyingnumber
}

$username=read-host "what's your name?"
$apt=read-host "Which department do you belong to?"
Get-HostSrvTag | Out-File $home\desktop\"$username.txt"
Get-MacAddress | Out-File $home\desktop\"$username.txt" -NoClobber -Append
Get-MonitorSn  | Out-File $home\desktop\"$username.txt" -NoClobber -Append


$sertag = (Get-Content $home\desktop\"$username.txt" -totalcount 3)[-1]
$macaddr = (Get-Content $home\desktop\"$username.txt" -totalcount 10)[-1]
$mtorsn1 = (Get-Content $home\desktop\"$username.txt" -totalcount 16)[-1]
$mtorsn2 = (Get-Content $home\desktop\"$username.txt" -totalcount 17)[-1]

if([String]::IsNullOrEmpty($mtorsn2))
{
	$mtorsn = $mtorsn1
}
else
{
	$mtorsn = $mtorsn1+"|"+$mtorsn2
}

$sertag = $sertag.Trim()
$data = @{name="$username";apt="$apt";sertag="$sertag";macaddr="$macaddr";mtorsn="$mtorsn"} | ConvertTo-Json -Compress
Invoke-WebRequest -usebasicparsing http://10.1.48.30:8080/assets/ -contenttype "application/json" -method post -body $data
