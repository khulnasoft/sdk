@echo off

REM Parse the -virtualenv command line argument and populate `args` with all other arguments.
set khulnasoft_runtime_python_virtualenv=
set args=
:parse
if "%~1"=="" goto endparse
if "%~1"=="-virtualenv" (
    REM Get the value as a fully-qualified path.
    set "khulnasoft_runtime_python_virtualenv=%~f2"
    shift /1
) else (
    set args=%args% %~1
)
shift /1
goto parse
:endparse

if defined khulnasoft_runtime_python_virtualenv (
    REM If python exists in the virtual environment, set PATH and run it.
    if exist "%khulnasoft_runtime_python_virtualenv%\Scripts\python.exe" (
        REM Update PATH and unset PYTHONHOME.
        set "PATH=%khulnasoft_runtime_python_virtualenv%\Scripts;%PATH%"
        set PYTHONHOME=

        REM Run python from the virtual environment.
        "%khulnasoft_runtime_python_virtualenv%\Scripts\python.exe" -u -m khulnasoft.policy %args%
        exit /B
    ) else (
        echo The 'virtualenv' option in KhulnasoftPolicy.yaml is set to %khulnasoft_runtime_python_virtualenv%, but %khulnasoft_runtime_python_virtualenv% doesn't appear to be a virtual environment. 1>&2
        echo Run the following commands to create the virtual environment and install dependencies into it: 1>&2
        echo     1. python -m venv %khulnasoft_runtime_python_virtualenv% 1>&2
        echo     2. %khulnasoft_runtime_python_virtualenv%\Scripts\python.exe -m pip install --upgrade pip setuptools wheel 1>&2
        echo     3. %khulnasoft_runtime_python_virtualenv%\Scripts\python.exe -m pip install -r %cd%\requirements.txt 1>&2
        echo For more information see: https://www.khulnasoft.com/docs/intro/languages/python/#virtual-environments 1>&2
        exit 1
    )
) else (
    if defined KHULNASOFT_PYTHON_CMD (
        REM If KHULNASOFT_PYTHON_CMD is defined, run it.
        "%KHULNASOFT_PYTHON_CMD%" -u -m khulnasoft.policy %args%
    ) else (
        REM Otherwise, just run python. We use `python` instead of `python3` because Windows
        REM Python installers install only `python.exe` by default.
        @python -u -m khulnasoft.policy %args%
    )
)
