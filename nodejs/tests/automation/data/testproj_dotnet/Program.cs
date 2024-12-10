// Copyright 2016-2023, KhulnaSoft, Ltd.  All rights reserved.

using System.Collections.Generic;
using Khulnasoft;

return await Deployment.RunAsync(() =>
{
    return new Dictionary<string, object>
    {
        {  "exp_static", "foo" },
    };
});