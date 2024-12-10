@echo off

if defined KHULNASOFT_RUNTIME_VIRTUALENV (
    REM If python exists in the virtual environment, set PATH and run it.
    if exist "%KHULNASOFT_RUNTIME_VIRTUALENV%\Scripts\python.exe" (
        REM Update PATH and unset PYTHONHOME.
        set "PATH=%KHULNASOFT_RUNTIME_VIRTUALENV%\Scripts;%PATH%"
        set PYTHONHOME=

        REM Run python from the virtual environment.
        "%KHULNASOFT_RUNTIME_VIRTUALENV%\Scripts\python.exe" -u -m khulnasoft.dynamic %*
        exit /B
    ) else (
        echo The 'virtualenv' option in Khulnasoft.yaml is set to %KHULNASOFT_RUNTIME_VIRTUALENV%, but %KHULNASOFT_RUNTIME_VIRTUALENV% doesn't appear to be a virtual environment. 1>&2
        echo Run the following commands to create the virtual environment and install dependencies into it: 1>&2
        echo     1. python -m venv %KHULNASOFT_RUNTIME_VIRTUALENV% 1>&2
        echo     2. %KHULNASOFT_RUNTIME_VIRTUALENV%\Scripts\python.exe -m pip install --upgrade pip setuptools wheel 1>&2
        echo     3. %KHULNASOFT_RUNTIME_VIRTUALENV%\Scripts\python.exe -m pip install -r %cd%\requirements.txt 1>&2
        echo For more information see: https://www.khulnasoft.com/docs/intro/languages/python/#virtual-environments 1>&2
        exit 1
    )
) else (
    if defined KHULNASOFT_PYTHON_CMD (
        REM If KHULNASOFT_PYTHON_CMD is defined, run it.
        "%KHULNASOFT_PYTHON_CMD%" -u -m khulnasoft.dynamic %*
    ) else (
        REM Otherwise, just run python. We use `python` instead of `python3` because Windows
        REM Python installers install only `python.exe` by default.
        @python -u -m khulnasoft.dynamic %*
    )
)
