# Khulnasoft Python SDK

The Khulnasoft Python SDK (khulnasoft) is the core package used when writing Khulnasoft programs in Python. It contains everything that you’ll need in order to interact with Khulnasoft resource providers and express infrastructure using Python code. Khulnasoft resource providers all depend on this library and express their resources in terms of the types defined in this module.

The Khulnasoft Python SDK requires a [supported version](https://devguide.python.org/versions/#versions) of Python.

note:
pip is required to install dependencies. If you installed Python from source, with an installer from [python.org](https://python.org/), or via [Homebrew](https://brew.sh/) you should already have pip. If Python is installed using your OS package manager, you may have to install pip separately, see [Installing pip/setuptools/wheel with Linux Package Managers](https://packaging.python.org/guides/installing-using-linux-tools/). For example, on Debian/Ubuntu you must run sudo apt install python3-venv python3-pip.

## Getting Started

The fastest way to get up and running is to choose from one of the following Getting Started guides:
-[aws](https://www.khulnasoft.com/docs/get-started/aws/?language=python)
-[microsoft azure](https://www.khulnasoft.com/docs/get-started/azure/?language=python)
-[google cloud](https://www.khulnasoft.com/docs/get-started/gcp/?language=python)
-[kubernetes](https://www.khulnasoft.com/docs/get-started/kubernetes/?language=python)

## Khulnasoft Programming Model

The Khulnasoft programming model defines the core concepts you will use when creating infrastructure as code programs using Khulnasoft. Architecture & Concepts describes these concepts with examples available in Python. These concepts are made available to you in the Khulnasoft SDK.

The Khulnasoft SDK is available to Python developers as a Pip package distributed on [PyPI](https://www.khulnasoft.com/docs/intro/languages/python/#pypi-packages) . To learn more, [refer to the Khulnasoft SDK Reference Guide](https://www.khulnasoft.com/docs/reference/pkg/python/khulnasoft/).

The Khulnasoft programming model includes a core concept of Input and Output values, which are used to track how outputs of one resource flow in as inputs to another resource. This concept is important to understand when getting started with Python and Khulnasoft, and the [Inputs and Outputs] (https://www.khulnasoft.com/docs/intro/concepts/inputs-outputs/)documentation is recommended to get a feel for how to work with this core part of Khulnasoft in common cases.


## [The Khulnasoft Python Resource Model](https://www.khulnasoft.com/docs/reference/pkg/python/khulnasoft/#the-khulnasoft-python-resource-model-1)

Like most languages usable with Khulnasoft, Khulnasoft represents cloud resources as classes and Python programs can instantiate those classes. All classes that can be instantiated to produce actual resources derive from the khulnasoft.Resource class.