@echo off
setlocal
for /f "delims=" %%i in ('node -e "console.log(require.resolve('@khulnasoft/khulnasoft/cmd/dynamic-provider'))"') do set KHULNASOFT_DYNAMIC_PROVIDER_SCRIPT_PATH=%%i
if DEFINED KHULNASOFT_DYNAMIC_PROVIDER_SCRIPT_PATH (
   @node "%KHULNASOFT_DYNAMIC_PROVIDER_SCRIPT_PATH%" %*
)
