# Запуск Ubuntu для ДЗ (Windows, PowerShell).
#   powershell -ExecutionPolicy Bypass -File .\start.ps1 ivan          — первый запуск
#   powershell -ExecutionPolicy Bypass -File .\start.ps1               — вернуться в контейнер
#   powershell -ExecutionPolicy Bypass -File .\start.ps1 -Reset ivan   — начать с нуля
param([string]$Student, [switch]$Reset)

$Name = 'linux-quest'
$Image = 'ghcr.io/ycovich/linux-quest'

if ($Reset) {
    docker rm -f $Name *> $null
}

$Student = "$Student".ToLower()

# Контейнер уже есть: возвращаемся в него, но только если он создан под тем же именем
$Existing = docker container inspect -f '{{.Config.Labels.quest_student}}' $Name 2>$null
if ($LASTEXITCODE -eq 0) {
    if ($Existing -eq '<no value>') { $Existing = '' }
    if ($Student -and $Existing -and $Student -ne $Existing) {
        Write-Host "Контейнер уже создан для пользователя $Existing."
        Write-Host "Вернуться в него: .\start.ps1"
        Write-Host "Начать с нуля под именем ${Student}: .\start.ps1 -Reset $Student"
        exit 1
    }
    docker start -ai $Name
    exit
}

if ($Student -notmatch '^[a-z][a-z0-9_-]{0,30}$') {
    Write-Host 'Укажи своё имя латиницей, например:  .\start.ps1 ivan'
    exit 1
}

# Образ скачивается один раз, потом только обновляется.
# Если обновить не вышло (например, нет интернета), работаем со старым образом
docker image inspect $Image *> $null
if ($LASTEXITCODE -eq 0) {
    docker pull -q $Image *> $null
} else {
    Write-Host 'Скачиваю образ, это займёт пару минут...'
    docker pull $Image
    if ($LASTEXITCODE -ne 0) { exit 1 }
}
docker run -it --name $Name -h tms-lab --label quest_student=$Student -e STUDENT=$Student $Image
