# Copyright 2016-2020, KhulnaSoft, Ltd.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import argparse
import os
import sys
import traceback
import runpy

import khulnasoft
import khulnasoft.runtime


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("engineAddress")
    parser.add_argument("program")
    parser.add_argument("--tracing", required=False, nargs=1, metavar="TRACING_FILE")
    parser.add_argument("--logtostderr", required=False, action="store_true")
    parser.add_argument("--logflow", required=False, action="store_true")
    args = parser.parse_args()
    program = args.program

    # If any config variables are present, parse and set them, so subsequent accesses are fast.
    config_env = khulnasoft.runtime.get_config_env()
    for k, v in config_env.items():
        khulnasoft.runtime.set_config(k, v)

    # Configure the runtime so that the user program hooks up to Khulnasoft as appropriate.
    if (
        "KHULNASOFT_PROJECT" in os.environ
        and "KHULNASOFT_STACK" in os.environ
        and "KHULNASOFT_DRY_RUN" in os.environ
    ):
        khulnasoft.runtime.configure(
            khulnasoft.runtime.Settings(
                project=os.environ["KHULNASOFT_PROJECT"],
                stack=os.environ["KHULNASOFT_STACK"],
                dry_run=os.environ["KHULNASOFT_DRY_RUN"] == "true",
                # KHULNASOFT_ORGANIZATION might not be set for diy backends
                organization=os.environ.get("KHULNASOFT_ORGANIZATION", "organization"),
            )
        )

    successful = False

    try:
        runpy.run_path(program, run_name="__main__")
        successful = True
    except Exception:  # noqa: BLE001 catch blind exception
        khulnasoft.log.error(
            "Program failed with an unhandled exception:\n" + traceback.format_exc()
        )
    finally:
        sys.stdout.flush()
        sys.stderr.flush()

    exit_code = 0 if successful else 1
    sys.exit(exit_code)


main()
