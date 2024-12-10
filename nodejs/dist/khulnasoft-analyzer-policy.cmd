@echo off
setlocal
for /f "delims=" %%i in ('node -e "console.log(require.resolve('@khulnasoft/khulnasoft/cmd/run-policy-pack'))"') do set KHULNASOFT_RUN_SCRIPT_PATH=%%i
if DEFINED KHULNASOFT_RUN_SCRIPT_PATH (
   @node "%KHULNASOFT_RUN_SCRIPT_PATH%" %*
)
